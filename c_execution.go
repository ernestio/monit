/* This Source Code Form is subject to the terms of the Mozilla Public
 * License, v. 2.0. If a copy of the MPL was not distributed with this
 * file, You can obtain one at http://mozilla.org/MPL/2.0/. */

package main

type Execution struct {
}

func (n *Execution) Handle(subject string, components []interface{}, lines []Message) []Message {
	switch subject {
	case "executions.create":
		lines = append(lines, Message{Body: "Running executions:", Level: "INFO"})
	case "executions.create.done":
		lines = n.getDetails(components)
		return append(lines, Message{Body: "Executions ran", Level: "INFO"})
	case "executions.create.error":
		lines = n.getDetails(components)
		return append(lines, Message{Body: "Executions failed", Level: "INFO"})
	}
	return lines
}

func (n *Execution) getDetails(components []interface{}) (lines []Message) {
	for _, v := range components {
		r := v.(map[string]interface{})
		name := r["name"].(string)
		status := r["status"].(string)
		lines = append(lines, Message{Body: " - " + name, Level: ""})
		lines = append(lines, Message{Body: "   Status    : " + status, Level: ""})
		if status == "errored" {
			err := r["error_message"].(string)
			lines = append(lines, Message{Body: "   Error     : " + err, Level: "ERROR"})
		}
	}

	return lines
}
