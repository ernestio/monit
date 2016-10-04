/* This Source Code Form is subject to the terms of the Mozilla Public
 * License, v. 2.0. If a copy of the MPL was not distributed with this
 * file, You can obtain one at http://mozilla.org/MPL/2.0/. */

package main

type Vpc struct {
}

func (n *Vpc) Handle(subject string, components []interface{}, lines []Message) []Message {
	switch subject {
	case "vpcs.create":
		lines = append(lines, Message{Body: "Creating firewall:", Level: "INFO"})
	case "vpcs.create.done":
		lines = n.getDetails(components)
		return append(lines, Message{Body: "Vpc created", Level: "INFO"})
	case "vpcs.create.error":
		lines = n.getDetails(components)
		return append(lines, Message{Body: "Vpc creation failed", Level: "INFO"})
	case "vpcs.delete":
		return append(lines, Message{Body: "Deleting vpcs:", Level: "INFO"})
	case "vpcs.delete.done":
		return append(lines, Message{Body: "Vpc deleted", Level: "INFO"})
	case "vpcs.delete.error":
		lines = n.getDetails(components)
		return append(lines, Message{Body: "Vpc deletion failed", Level: "INFO"})
	}
	return lines
}

func (n *Vpc) getDetails(components []interface{}) (lines []Message) {
	for _, v := range components {
		r := v.(map[string]interface{})
		id := r["vpc_id"].(string)
		subnet := r["vpc_subnet"].(string)
		status := r["status"].(string)
		lines = append(lines, Message{Body: " - " + id, Level: ""})
		lines = append(lines, Message{Body: " - Subnet    : " + subnet, Level: ""})
		lines = append(lines, Message{Body: "   Status    : " + status, Level: ""})
		if status == "errored" {
			err := r["error_message"].(string)
			lines = append(lines, Message{Body: "   Error     : " + err, Level: "ERROR"})
		}
	}

	return lines
}
