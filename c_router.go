/* This Source Code Form is subject to the terms of the Mozilla Public
 * License, v. 2.0. If a copy of the MPL was not distributed with this
 * file, You can obtain one at http://mozilla.org/MPL/2.0/. */

package main

// Router : ...
type Router struct {
}

// Handle : ...
func (n *Router) Handle(subject string, components []interface{}, lines []Message) []Message {
	switch subject {

	case "router.create.done", "router.create.error":
		lines = n.getSingleDetail(components, "Created router")
	case "router.delete.done", "router.delete.error":
		lines = n.getSingleDetail(components, "Deleted router")
	}
	return lines
}

func (n *Router) getSingleDetail(v interface{}, prefix string) (lines []Message) {
	r := v.(map[string]interface{})
	name, _ := r["name"].(string)
	if prefix != "" {
		name = prefix + " " + name
	}
	status, _ := r["status"].(string)
	ip, _ := r["ip"].(string)
	lines = append(lines, Message{Body: " - " + name, Level: ""})
	lines = append(lines, Message{Body: "   IP        : " + ip, Level: ""})
	lines = append(lines, Message{Body: "   Status    : " + status, Level: ""})
	if status == "errored" {
		err, _ := r["error"].(string)
		lines = append(lines, Message{Body: "   Error     : " + err, Level: "ERROR"})
	}

	return lines
}
