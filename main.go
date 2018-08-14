/* This Source Code Form is subject to the terms of the Mozilla Public
 * License, v. 2.0. If a copy of the MPL was not distributed with this
 * file, You can obtain one at http://mozilla.org/MPL/2.0/. */

package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/nats-io/go-nats"
	"github.com/r3labs/broadcast"
)

var nc *nats.Conn
var bc *broadcast.Server
var host string
var port string
var secret string
var err error

func main() {
	setup()
	defer nc.Close()

	// Create new SSE server
	bc = broadcast.New()
	bc.AutoStream = true
	defer bc.Close()

	// Create new HTTP Server and add the route handler
	mux := http.NewServeMux()
	mux.HandleFunc("/events", handler)

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
