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

	// Subscribe to service events
	serviceSubjects := []string{
		"service.create",
		"service.delete",
		"service.import",
	}

	for _, s := range serviceSubjects {
		_, err = nc.Subscribe(s, serviceHandler)
		if err != nil {
			log.Println(err)
			return
		}
	}

	// Subscribe to component events
	componentSubjects := []string{
		"*.create.*",
		"*.create.*.*",
		"*.update.*",
		"*.update.*.*",
		"*.delete.*",
		"*.delete.*.*",
		"*.find.*",
		"*.find.*.*",
	}

	for _, s := range componentSubjects {
		_, err = nc.Subscribe(s, componentHandler)
		if err != nil {
			log.Println(err)
			return
		}
	}

	// Start Listening
	addr := fmt.Sprintf("%s:%s", host, port)
	_ = http.ListenAndServe(addr, mux)
}
