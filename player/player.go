
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

package player

// #cgo pkg-config: mpv
// #include <mpv/client.h>
import "C"

import "fmt"
import "log"
import "strconv"
import "sync"

import "github.com/vaughan0/go-ini"

import "github.com/ghedo/grooved/library"
import "github.com/ghedo/grooved/notify"
import "github.com/ghedo/grooved/util"

type Event byte
type Status byte

const (
    StatusPlaying Status = iota
    StatusPaused
    StatusStopped
)

func (s Status) String() string {
    switch s {
    case StatusPlaying:
        return "play"

    case StatusPaused:
        return "pause"

    case StatusStopped:
        return "stop"
    }

    return "invalid"
}

type Player struct {
    handle *C.mpv_handle
    Status Status

    library string
    notify  bool
    started bool

    Verbose bool

    HandleStatusChange func()
    HandleTrackChange  func()
    HandleTracksChange func()
    HandleVolumeChange func()

    Wait sync.WaitGroup
}

func (p *Player) ChangeStatus(status Status) {
    p.Status = status