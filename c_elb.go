/* This Source Code Form is subject to the terms of the Mozilla Public
 * License, v. 2.0. If a copy of the MPL was not distributed with this
 * file, You can obtain one at http://mozilla.org/MPL/2.0/. */

package main

// ELB : ...
type ELB struct {
}

// Handle : ...
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
	case "elbs.find":
		lines = append(lines, Message{Body: "Importing ELBs:", Level: "INFO"})
	case "elbs.find.done":
		lines = n.getDetails(components)
		return append(lines, Message{Body: "ELBs imported", Level: "INFO"})
	case "elbs.find.error":
		lines = n.getDetails(components)
		return append(lines, Message{Body: "ELBs import failed", Level: "INFO"})

	case "elb.create.done", "elb.create.error":
		lines = n.getSingleDetail(components, "Created ELB")
	case "elb.update.done", "elb.update.error":
		lines = n.getSingleDetail(components, "Updated ELB")
	case "elb.delete.done", "elb.delete.error":
		lines = n.getSingleDetail(components, "Deleted ELB")
	case "elb.find.done", "elb.find.error":
		lines = n.getSingleDetail(components, "Found ELB")

	}
	return lines
}

func (n *ELB) getDetails(components []interface{}) (lines []Message) {
	for _, v := range components {
		for _, l := range n.getSingleDetail(v, "") {
			lines = append(lines, l)
		}
	}

	return lines
}

func (n *ELB) getSingleDetail(v interface{}, prefix string) (lines []Message) {
	r := v.(map[string]interface{})
	name, _ := r["name"].(string)
	if prefix != "" {
		name = prefix + " " + name
	}
	status, _ := r["status"].(string)
	lines = append(lines, Message{Body: " - " + name, Level: ""})
	lines = append(lines, Message{Body: "   Status    : " + status, Level: ""})
	if r["dns_name"] != nil {
		dnsName, _ := r["dns_name"].(string)
		if dnsName != "" {
			lines = append(lines, Message{Body: "   DNS    : " + dnsName, Level: ""})
		}
	}
	lines = append(lines)
	if status == "errored" {
		err, _ := r["error"].(string)
		lines = append(lines, Message{Body: "   Error     : " + err, Level: "ERROR"})
	}

	return lines
}
