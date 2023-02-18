
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
// #include <stdlib.h>
// #include <mpv/client.h>
//
// char **alloc_array(int size) {
//     return calloc(sizeof(char*), size);
// }
//
// void set_array(char **a, int i, char* s) {
//     a[i] = s;
// }
//
// int64_t node_get_int64(mpv_node *node) { return node -> u.int64; }
// double  node_get_double(mpv_node *node) { return node -> u.double_; }
// char   *node_get_string(mpv_node *node) { return node -> u.string; }
//
// int64_t node_get_map_len(mpv_node *node) {
//     return node -> u.list -> num;
// }
//
// char *node_get_map_key(mpv_node *node, int64_t i) {
//   return node -> u.list -> keys[i];
// }
//
// mpv_node *node_get_map_val(mpv_node *node, int64_t i) {
//   return &node -> u.list -> values[i];
// }
import "C"

import "fmt"
import "unsafe"

const (
    FormatNone      C.mpv_format = 0
    FormatString                 = 1
    FormatFlag                   = 3
    FormatInt64                  = 4
    FormatDouble                 = 5
    FormatNode                   = 6
    FormatNodeArray              = 7
    FormatNodeMap                = 8
)

func ErrorString(err C.int) error {
    err_str := C.mpv_error_string(err)
    return fmt.Errorf(C.GoString(err_str))
}

func (p *Player) SetOptionString(name string, value string) error {
    cname := C.CString(name)
    defer C.free(unsafe.Pointer(cname))

    cvalue := C.CString(value)
    defer C.free(unsafe.Pointer(cvalue))

    err := C.mpv_set_option_string(p.handle, cname, cvalue)
    if err != 0 {
        return ErrorString(err)
    }

    return nil
}

func (p *Player) SetProperty(name string, value interface{}) error {
    var data unsafe.Pointer
    var format C.mpv_format

    cname := C.CString(name)
    defer C.free(unsafe.Pointer(cname))

    switch value.(type) {
    case bool:
        format  = FormatFlag
        cvalue := value.(bool)
        data    = unsafe.Pointer(&cvalue)

    case int64:
        format  = FormatInt64
        cvalue := value.(int64)
        data    = unsafe.Pointer(&cvalue)

    case float64:
        format  = FormatDouble
        cvalue := value.(float64)
        data    = unsafe.Pointer(&cvalue)

    case string:
        format  = FormatString
        cvalue := C.CString(value.(string))
        data    = unsafe.Pointer(&cvalue)

        defer C.free(unsafe.Pointer(cvalue))
    }

    err := C.mpv_set_property(p.handle, cname, format, data)
    if err != 0 {
        return ErrorString(err)
    }

    return nil
}

func (p *Player) GetProperty(name string) (interface{}, error) {
    var node C.mpv_node

    cname := C.CString(name)
    defer C.free(unsafe.Pointer(cname))

    err := C.mpv_get_property(
        p.handle, cname, FormatNode, unsafe.Pointer(&node),
    )
    if err != 0 {
        return nil, ErrorString(err)
    }
    defer C.mpv_free_node_contents(&node)

    return node_to_go(&node)