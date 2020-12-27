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

package bus

import "fmt"

import "github.com/godbus/dbus"
import "github.com/godbus/dbus/introspect"
import "github.com/godbus/dbus/prop"

import "github.com/ghedo/grooved/player"

const bus_introspection = `
<?xml version="1.0" encoding="UTF-8"?>
<node>
  <interface name="io.github.ghedo.grooved.Player">
    <property name="PlaybackStatus" type="s" access="read">
    </property>

    <property name="LoopStatus" type="s" access="readwrite">
    </property>

    <property name="TrackMetadata" type="a{ss}" access="read">
    </property>

    <property name="TrackPath" type="s" access="read">
    </property>

    <property name="TrackLength" type="d" access="read">
    </property>

    <method name="TrackPosition">
      <arg direction="out" name="position" type="d"/>
      <arg direction="out" name="percent" type="d"/>
    </method>

    <property name="TrackTitle" type="s" access="read">
    </property>

    <property name="Tracks" type="as" access="read">
    </property>

    <property name="Volume" type="d" access="readwrite">
    </property>

    <method name="Play">
    </method>

    <method name="Pause">
    </method>

    <method name="Toggle">
    </method>

    <method name="Next">
    </method>

    <method name="Prev">
    </method>

    <method name="Stop">
    </method>

    <method name="Seek">
      <arg