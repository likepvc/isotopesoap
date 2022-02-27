
/*
 * Groovy music player daemon.
 *
 * Copyright (c) 2014, Alessandro Ghedini
 * All rights reserved.
 *
 * Redistribution and use in source and binary forms, with or without
 * modification, are permitted provided that the following conditions are
 * met:
 *
 *     * Redistributions of source code must retain the above copyright
 *       notice, this list of conditions and the following disclaimer.
 *
 *     * Redistributions in binary form must reproduce the above copyright
 *       notice, this list of conditions and the following disclaimer in the
 *       documentation and/or other materials provided with the distribution.
 *
 * THIS SOFTWARE IS PROVIDED BY THE COPYRIGHT HOLDERS AND CONTRIBUTORS "AS
 * IS" AND ANY EXPRESS OR IMPLIED WARRANTIES, INCLUDING, BUT NOT LIMITED TO,
 * THE IMPLIED WARRANTIES OF MERCHANTABILITY AND FITNESS FOR A PARTICULAR
 * PURPOSE ARE DISCLAIMED. IN NO EVENT SHALL THE COPYRIGHT HOLDER OR
 * CONTRIBUTORS BE LIABLE FOR ANY DIRECT, INDIRECT, INCIDENTAL, SPECIAL,
 * EXEMPLARY, OR CONSEQUENTIAL DAMAGES (INCLUDING, BUT NOT LIMITED TO,
 * PROCUREMENT OF SUBSTITUTE GOODS OR SERVICES; LOSS OF USE, DATA, OR
 * PROFITS; OR BUSINESS INTERRUPTION) HOWEVER CAUSED AND ON ANY THEORY OF
 * LIABILITY, WHETHER IN CONTRACT, STRICT LIABILITY, OR TORT (INCLUDING
 * NEGLIGENCE OR OTHERWISE) ARISING IN ANY WAY OUT OF THE USE OF THIS
 * SOFTWARE, EVEN IF ADVISED OF THE POSSIBILITY OF SUCH DAMAGE.
 */

package main

import "log"
import "path/filepath"
import "strconv"
import "time"
import "os"

import "github.com/docopt/docopt-go"
import "github.com/godbus/dbus"

const bus_name = "io.github.ghedo.grooved"
const bus_path = "/io/github/ghedo/grooved"
const bus_interface = "io.github.ghedo.grooved.Player"

func main() {
    log.SetFlags(0)

    usage := `groovectl [options] <command> [args...]

Usage:
  groovectl play
  groovectl pause
  groovectl toggle
  groovectl next
  groovectl prev
  groovectl stop
  groovectl add <track> ...
  groovectl [--append] load <file>
  groovectl save <file>
  groovectl goto <index>
  groovectl rm <index>
  groovectl ls
  groovectl status
  groovectl seek <seconds>
  groovectl loop (none|track|list|force)
  groovectl quit

Commands:
  play                           Unpause the player.
  pause                          Pause the player.
  toggle                         Toggle the player's pause status.
  next                           Skip to next track.
  prev                           Skip to previous track.
  stop                           Stop playback and clear tracklist.
  add <track> ...                Append tracks to the player's tracklist.
  [--append] load <file>         Load a playlist file.
  save <file>                    Save the tracklist to a playlist file.
  goto <index>                   Skip to a specific track in the tracklist.
  rm <index>                     Remove a track from the tracklist.
  ls                             Show the tracklist.
  status                         Show the status of the player.
  seek <seconds>                 Seek by a relative amount of seconds.
  loop (none|track|list|force)   Set the player's loop mode.
  quit                           Shutdown the player.

Options:
  -h, --help                         Show the program's help message and exit.`

    args, err := docopt.Parse(usage, nil, true, "", true)
    if err != nil {
        log.Fatalf("Invalid arguments: %s", err)
    }

    conn, err := dbus.SessionBus()
    if err != nil {