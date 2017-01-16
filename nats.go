/* This Source Code Form is subject to the terms of the Mozilla Public
 * License, v. 2.0. If a copy of the MPL was not distributed with this
 * file, You can obtain one at http://mozilla.org/MPL/2.0/. */

package main

import (
	"github.com/nats-io/nats"
	"github.com/r3labs/sse"
	"log"
	"time"
)

func natsHandler(msg *nats.Msg) {
	var notification Notification
	if err := processNotification(&notification, msg); err != nil {
		return
	}

	switch msg.Subject {
	case "monitor.user":
		// Publish messages to subscribers
		for _, nm := range notification.Messages {
			publishMessage(notification.getServiceID(), &nm)
		}
	case "service.create", "service.delete":
		var handler Service
		// Create a new stream
		log.Println("Creating stream for", notification.getServiceID())
		s.CreateStream(notification.getServiceID())
		lines := handler.Handle(msg.Subject, notification.Messages)
		for _, nm := range lines {
			publishMessage(notification.getServiceID(), &nm)
		}
	case "service.create.done", "service.create.error", "service.delete.done", "service.delete.error", "service.import.done", "service.import.error":
		var handler Service
		lines := handler.Handle(msg.Subject, notification.Messages)
		for _, nm := range lines {
			publishMessage(notification.getServiceID(), &nm)
		}
		time.Sleep(10 * time.Millisecond)
		// Remove a new stream when the build completes
		log.Println("Closing stream for", notification.getServiceID())
		go func(s *sse.Server) {
			s.RemoveStream(notification.getServiceID())
		}(s)
	}
}
