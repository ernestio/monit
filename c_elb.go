/* This Source Code Form is subject to the terms of the Mozilla Public
 * License, v. 2.0. If a copy of the MPL was not distributed with this
 * file, You can obtain one at http://mozilla.org/MPL/2.0/. */

package main

type ELB struct {
}

func (n *ELB) Handle(subject string, components []interface{}, lines []Message) []Message {
	switch subject {
	case "elbs.create":
		lines = append(lines, Message{Body: "Creating ELBs:", Level: "INFO"})
	case "elbs.create.done":
		lines = n.getDetails(components)
		return append(lines, Message{Body: "ELBs created", Level: "INFO"})
	case "elbs.create.error":
		lines = n.getDetails(components)
		return append(lines, Message{Body: "ELBs creation failed", Level: "INFO"})
	case "elbs.update":
		return append(lines, Message{Body: "Updating ELBs:", Level: "INFO"})
	case "elbs.update.done":
		lines = n.getDetails(components)
		return append(lines, Message{Body: "ELBs updated", Level: "INFO"})
	case "elbs.update.error":
		lines = n.getDetails(components)
		return append(lines, Message{Body: "ELBs modification failed", Level: "INFO"})
	case "elbs.delete":
		return append(lines, Message{Body: "Deleting ELBs:", Level: "INFO"})
	case "elbs.delete.done":
		lines = n.getDetails(components)
		return append(lines, Message{Body: "ELBs deleted", Level: "INFO"})
	case "elbs.delete.error":
		lines = n.getDetails(components)
		return append(lines, Message{Body: "ELBs deletion failed", Level: "INFO"})
	}
	return lines
}

func (n *ELB) getDetails(components []interface{}) (lines []Message) {
	for _, v := range components {
		r := v.(map[string]interface{})
		name := r["elb_name"].(string)
		status := r["status"].(string)
		lines = append(lines, Message{Body: " - " + name, Level: ""})
		lines = append(lines, Message{Body: "   Status    : " + status, Level: ""})
		if r["elb_dns_name"] != nil {
			dnsName := r["elb_dns_name"].(string)
			lines = append(lines, Message{Body: "   DNS    : " + dnsName, Level: ""})
		}
		lines = append(lines)
		if status == "errored" {
			err := r["error_message"].(string)
			lines = append(lines, Message{Body: "   Error     : " + err, Level: "ERROR"})
		}
	}

	return lines
}
