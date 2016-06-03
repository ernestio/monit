/* This Source Code Form is subject to the terms of the Mozilla Public
 * License, v. 2.0. If a copy of the MPL was not distributed with this
 * file, You can obtain one at http://mozilla.org/MPL/2.0/. */

package main

import (
	"time"

	"github.com/nats-io/nats"
	"github.com/r3labs/sse"
)

func natsHandler(msg *nats.Msg) {
	notification, err := processNotification(msg)
	if err != nil {
		return
	}

	switch msg.Subject {
	case "monitor.user":
		// Publish messages to subscribers
		for _, nm := range notification.Messages {
			publishMessage(notification.getServiceID(), &nm)
		}
	case "service.create", "service.delete":
		// Create a new stream
		s.CreateStream(notification.getServiceID())
	case "service.create.done", "service.create.error", "service.delete.done", "service.delete.error":
		// Remove a new stream when the build completes
		go func(s *sse.Server) {
			// Notifications appear out of order, wait for all notifications to come through before closing
			time.Sleep(time.Second)
			s.RemoveStream(notification.getServiceID())
		}(s)

	}
}
