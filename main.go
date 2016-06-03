/* This Source Code Form is subject to the terms of the Mozilla Public
 * License, v. 2.0. If a copy of the MPL was not distributed with this
 * file, You can obtain one at http://mozilla.org/MPL/2.0/. */

package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"gopkg.in/redis.v3"

	"github.com/nats-io/nats"
	"github.com/r3labs/sse"
)

var s *sse.Server
var db *redis.Client

func main() {
	// Open Nats connection
	n, err := nats.Connect(os.Getenv("NATS_URI"))
	if err != nil {
		log.Println("Could not connect to nats")
		return
	}
	defer n.Close()

	redisCfg := redisConfig{}
	msg, err := n.Request("config.get.redis", []byte(""), 1*time.Second)
	if err != nil {
		panic("Cant get redis config")
	}
	json.Unmarshal(msg.Data, &redisCfg)

	// Create new SSE server
	s = sse.New()
	s.AutoStream = true
	defer s.Close()

	// Open DB connection
	db = redis.NewClient(&redis.Options{
		Addr:     redisCfg.Host,
		Password: redisCfg.Password,
		DB:       redisCfg.DB,
	})

	// Create new HTTP Server and add the route handler
	mux := http.NewServeMux()
	mux.HandleFunc("/events", authMiddleware)

	// Start nats handler, subscribe to all events
	n.Subscribe(">", natsHandler)

	monitorCfg := monitorConfig{}
	msg, err = n.Request("config.get.monitor", []byte(""), 1*time.Second)
	if err != nil {
		panic("Can't get monitor config")
	}
	json.Unmarshal(msg.Data, &monitorCfg)

	// Start Listening
	addr := fmt.Sprintf("%s:%s", monitorCfg.Host, monitorCfg.Port)
	http.ListenAndServe(addr, mux)
}
