/* This Source Code Form is subject to the terms of the Mozilla Public
 * License, v. 2.0. If a copy of the MPL was not distributed with this
 * file, You can obtain one at http://mozilla.org/MPL/2.0/. */

package main

// Vpc : ...
type Vpc struct {
}

// Handle : ...
func (n *Vpc) Handle(subject string, components []interface{}, lines []Message) []Message {
	switch subject {
	case "vpcs.create":
		lines = append(lines, Message{Body: "Creating Vpc:", Level: "INFO"})
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
	case "vpcs.find":
		lines = append(lines, Message{Body: "Importing Vpc:", Level: "INFO"})
	case "vpcs.find.done":
		lines = n.getDetails(components)
		return append(lines, Message{Body: "Vpc imported", Level: "INFO"})
	case "vpcs.find.error":
		lines = n.getDetails(components)
		return append(lines, Message{Body: "Vpc import failed", Level: "INFO"})

	case "vpc.create.done", "vpc.create.error":
		lines = n.getSingleDetail(components, "VPC created")
	case "vpc.update.done", "vpc.update.error":
		lines = n.getSingleDetail(components, "VPC udpated")
	case "vpc.delete.done", "vpc.delete.error":
		lines = n.getSingleDetail(components, "VPC deleted")
	case "vpc.find.done", "vpc.find.error":
		lines = n.getSingleDetail(components, "VPC Found")
	}
	return lines
}

func (n *Vpc) getDetails(components []interface{}) (lines []Message) {
	for _, v := range components {
		for _, l := range n.getSingleDetail(v, "") {
			lines = append(lines, l)
		}
	}

	return lines
}

func (n *Vpc) getSingleDetail(v interface{}, prefix string) (lines []Message) {
	r := v.(map[string]interface{})
	id, _ := r["vpc_id"].(string)
	if prefix != "" {
		id = prefix + " " + id
	}
	subnet, _ := r["vpc_subnet"].(string)
	status, _ := r["status"].(string)
	lines = append(lines, Message{Body: " - " + id, Level: ""})
	lines = append(lines, Message{Body: "   Subnet    : " + subnet, Level: ""})
	lines = append(lines, Message{Body: "   Status    : " + status, Level: ""})
	if status == "errored" {
		err, _ := r["error"].(string)
		lines = append(lines, Message{Body: "   Error     : " + err, Level: "ERROR"})
	}
	return lines
}
