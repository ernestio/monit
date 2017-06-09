/* This Source Code Form is subject to the terms of the Mozilla Public
 * License, v. 2.0. If a copy of the MPL was not distributed with this
 * file, You can obtain one at http://mozilla.org/MPL/2.0/. */

package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/nats-io/nats"
	"github.com/r3labs/sse"
)

var nc *nats.Conn
var ss *sse.Server
var host string
var port string
var secret string
var err error

func main() {
	setup()
	defer nc.Close()

	// Create new SSE server
	ss = sse.New()
	ss.AutoStream = true
	ss.EncodeBase64 = true
	defer ss.Close()

	// Create new HTTP Server and add the route handler
	mux := http.NewServeMux()
	mux.HandleFunc("/events", authMiddleware)

	// Subscribe to subjects
	_, err = nc.Subscribe(">", natsHandler)
	if err != nil {
		log.Println(err)
		return
	}

	// Start Listening
	addr := fmt.Sprintf("%s:%s", host, port)
	_ = http.ListenAndServe(addr, mux)
}
