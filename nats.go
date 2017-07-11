/* This Source Code Form is subject to the terms of the Mozilla Public
 * License, v. 2.0. If a copy of the MPL was not distributed with this
 * file, You can obtain one at http://mozilla.org/MPL/2.0/. */

package main

import (
	"github.com/nats-io/nats"
	"github.com/r3labs/pattern"
)

func natsHandler(msg *nats.Msg) {
	var services = []string{
		"service.create",
		"service.create.*",
		"service.delete",
		"service.delete.*",
		"service.import",
		"service.import.*",
	}
	var components = []string{
		"*.create.*",
		"*.create.*.*",
		"*.update.*",
		"*.update.*.*",
		"*.delete.*",
		"*.delete.*.*",
		"*.find.*",
		"*.find.*.*",
	}

	switch {
	case pattern.Match(msg.Subject, services...):
		processService(msg)
	case pattern.Match(msg.Subject, components...):
		processComponent(msg)
	}
}
