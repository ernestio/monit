/* This Source Code Form is subject to the terms of the Mozilla Public
 * License, v. 2.0. If a copy of the MPL was not distributed with this
 * file, You can obtain one at http://mozilla.org/MPL/2.0/. */

package main

import (
	"fmt"
	"net/http"

	"github.com/nats-io/nats"
	"github.com/r3labs/sse"
)

var n *nats.Conn
var s *sse.Server
var host string
var port string
var secret string

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
	n.Subscribe("monitor.user", natsHandler)
	n.Subscribe("service.create", natsHandler)
	n.Subscribe("service.delete", natsHandler)
	n.Subscribe("service.create.done", natsHandler)
	n.Subscribe("service.create.error", natsHandler)
	n.Subscribe("service.delete.done", natsHandler)
	n.Subscribe("service.delete.error", natsHandler)
	n.Subscribe("service.import.done", natsHandler)
	n.Subscribe("service.import.error", natsHandler)

	n.Subscribe("*.*", genericHandler)
	n.Subscribe("*.*.*", genericHandler)

	// Start Listening
	addr := fmt.Sprintf("%s:%s", host, port)
	http.ListenAndServe(addr, mux)
}
