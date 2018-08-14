/* This Source Code Form is subject to the terms of the Mozilla Public
 * License, v. 2.0. If a copy of the MPL was not distributed with this
 * file, You can obtain one at http://mozilla.org/MPL/2.0/. */

package main

import (
	"github.com/nats-io/go-nats"
	"github.com/r3labs/pattern"
)

func natsHandler(msg *nats.Msg) {
	var buildss = []string{
		"build.create",
		"build.create.*",
		"build.delete",
		"build.delete.*",
		"build.import",
		"build.import.*",
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
	case pattern.Match(msg.Subject, buildss...):
		processBuild(msg)
	case pattern.Match(msg.Subject, components...):
		processComponent(msg)
	}
}
