/* This Source Code Form is subject to the terms of the Mozilla Public
 * License, v. 2.0. If a copy of the MPL was not distributed with this
 * file, You can obtain one at http://mozilla.org/MPL/2.0/. */

package main

type Router struct {
}

func (n *Router) Handle(subject string, components []interface{}, lines []Message) []Message {
	switch subject {
	case "routers.create":
		lines = append(lines, Message{Body: "Creating routers:", Level: "INFO"})
	case "routers.create.done":
		lines = n.getDetails(components)
		return append(lines, Message{Body: "Routers created", Level: "INFO"})
	case "routers.create.error":
		lines = n.getDetails(components)
		return append(lines, Message{Body: "Routers creation failed", Level: "INFO"})
	case "routers.delete":
		return append(lines, Message{Body: "Deleting routers:", Level: "INFO"})
	case "routers.delete.done":
		return append(lines, Message{Body: "Routers deleted", Level: "INFO"})
	case "routers.delete.error":
		lines = n.getDetails(components)
		return append(lines, Message{Body: "Routers deletion failed", Level: "INFO"})
	}
	return lines
}

func (n *Router) getDetails(components []interface{}) (lines []Message) {
	for _, v := range components {
		r := v.(map[string]interface{})
		name, _ := r["name"].(string)
		status, _ := r["status"].(string)
		ip, _ := r["ip"].(string)
		lines = append(lines, Message{Body: " - " + name, Level: ""})
		lines = append(lines, Message{Body: "   IP        : " + ip, Level: ""})
		lines = append(lines, Message{Body: "   Status    : " + status, Level: ""})
		if status == "errored" {
			err, _ := r["error_message"].(string)
			lines = append(lines, Message{Body: "   Error     : " + err, Level: "ERROR"})
		}
	}

	return lines
}
