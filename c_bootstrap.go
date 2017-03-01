/* This Source Code Form is subject to the terms of the Mozilla Public
 * License, v. 2.0. If a copy of the MPL was not distributed with this
 * file, You can obtain one at http://mozilla.org/MPL/2.0/. */

package main

// Bootstrap : ..
type Bootstrap struct {
}

// Handle : ..
func (n *Bootstrap) Handle(subject string, components []interface{}, lines []Message) []Message {
	switch subject {
	case "bootstraps.create":
		lines = append(lines, Message{Body: "Running bootstraps:", Level: "INFO"})
	case "bootstraps.create.done":
		lines = n.getDetails(components)
		return append(lines, Message{Body: "Bootstrap ran", Level: "INFO"})
	case "bootstraps.create.error":
		lines = n.getDetails(components)
		return append(lines, Message{Body: "Bootstrap failed", Level: "INFO"})

	case "bootstrap.create.done", "bootstrap.create.error":
		lines = n.getSingleDetail(components, "Bootstrap ran")
	}

	return lines
}

func (n *Bootstrap) getDetails(components []interface{}) (lines []Message) {
	for _, v := range components {
		for _, l := range n.getSingleDetail(v, "") {
			lines = append(lines, l)
		}
	}

	return lines
}

func (n *Bootstrap) getSingleDetail(v interface{}, prefix string) (lines []Message) {
	r := v.(map[string]interface{})
	name, _ := r["name"].(string)
	if prefix != "" {
		name = prefix + " " + name
	}
	status, _ := r["status"].(string)
	lines = append(lines, Message{Body: " - " + name, Level: ""})
	lines = append(lines, Message{Body: "   Status    : " + status, Level: ""})
	if status == "errored" {
		err, _ := r["error"].(string)
		lines = append(lines, Message{Body: "   Error     : " + err, Level: "ERROR"})
	}
	return lines
}
