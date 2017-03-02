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

var n *nats.Conn
var s *sse.Server
var host string
var port string
var secret string
var err error

func main() {
	setup()
	defer n.Close()

	// Create new SSE server
	s = sse.New()
	s.AutoStream = true
	s.EncodeBase64 = true
	defer s.Close()

	// Create new HTTP Server and add the route handler
	mux := http.NewServeMux()
	mux.HandleFunc("/events", authMiddleware)

	// Start nats handler, subscribe to all events related with the monitor
	_, err = n.Subscribe("monitor.user", natsHandler)
	if err != nil {
		log.Println(err)
		return
	}
	_, err = n.Subscribe("service.create", natsHandler)
	if err != nil {
		log.Println(err)
		return
	}
	_, err = n.Subscribe("service.delete", natsHandler)
	if err != nil {
		log.Println(err)
		return
	}
	_, err = n.Subscribe("service.create.done", natsHandler)
	if err != nil {
		log.Println(err)
		return
	}
	_, err = n.Subscribe("service.create.error", natsHandler)
	if err != nil {
		log.Println(err)
		return
	}
	_, err = n.Subscribe("service.delete.done", natsHandler)
	if err != nil {
		log.Println(err)
		return
	}
	_, err = n.Subscribe("service.delete.error", natsHandler)
	if err != nil {
		log.Println(err)
		return
	}
	_, err = n.Subscribe("service.import.done", natsHandler)
	if err != nil {
		log.Println(err)
		return
	}
	_, err = n.Subscribe("service.import.error", natsHandler)
	if err != nil {
		log.Println(err)
		return
	}

	_, err = n.Subscribe("*.*", genericHandler)
	if err != nil {
		log.Println(err)
		return
	}
	_, err = n.Subscribe("*.*.*", genericHandler)
	if err != nil {
		log.Println(err)
		return
	}

	// Start Listening
	addr := fmt.Sprintf("%s:%s", host, port)
	_ = http.ListenAndServe(addr, mux)
}
