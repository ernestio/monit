/* This Source Code Form is subject to the terms of the Mozilla Public
 * License, v. 2.0. If a copy of the MPL was not distributed with this
 * file, You can obtain one at http://mozilla.org/MPL/2.0/. */

package main

// Execution : ...
type Execution struct {
}

// Handle : ...
func (n *Execution) Handle(subject string, components []interface{}, lines []Message) []Message {
	switch subject {
	case "execution.create.done", "execution.create.error":
		lines = n.getSingleDetail(components, "Ran execution")
	}
	return lines
}

func (n *Execution) getSingleDetail(v interface{}, prefix string) (lines []Message) {
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
