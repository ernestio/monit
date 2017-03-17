/* This Source Code Form is subject to the terms of the Mozilla Public
 * License, v. 2.0. If a copy of the MPL was not distributed with this
 * file, You can obtain one at http://mozilla.org/MPL/2.0/. */

package main

import "strings"

// Router : ...
type Router struct {
}

// Handle : ...
func (n *Router) Handle(subject string, c component, lines []Message) []Message {
	parts := strings.Split(subject, ".")
	subject = parts[0] + "." + parts[1]
	switch subject {

	case "router.create":
		lines = n.getSingleDetail(c, "Created router")
	case "routers.delete":
		lines = n.getSingleDetail(c, "Deleted router")
	}
	return lines
}

func (n *Router) getSingleDetail(c component, prefix string) (lines []Message) {
	name, _ := c["name"].(string)
	if prefix != "" {
		name = prefix + " " + name
	}
	status, _ := c["_state"].(string)
	level := "INFO"
	if status == "errored" {
		level = "ERROR"
	}
	ip, _ := c["ip"].(string)
	lines = append(lines, Message{Body: " " + name, Level: level})
	lines = append(lines, Message{Body: "   IP        : " + ip, Level: ""})
	lines = append(lines, Message{Body: "   Status    : " + status, Level: ""})
	if status == "errored" {
		err, _ := c["error"].(string)
		lines = append(lines, Message{Body: "   Error     : " + err, Level: ""})
	}

	return lines
}
