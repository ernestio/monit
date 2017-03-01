/* This Source Code Form is subject to the terms of the Mozilla Public
 * License, v. 2.0. If a copy of the MPL was not distributed with this
 * file, You can obtain one at http://mozilla.org/MPL/2.0/. */

package main

// EBSVolume : ...
type EBSVolume struct {
}

// Handle : ...
func (n *EBSVolume) Handle(subject string, components []interface{}, lines []Message) []Message {
	switch subject {
	case "ebs_volumes.create":
		return append(lines, Message{Body: "Creating ebs volumes:", Level: "INFO"})
	case "ebs_volumes.create.done":
		lines = n.getDetails(components)
		return append(lines, Message{Body: "EBS volumes successfully created", Level: "INFO"})
	case "ebs_volumes.create.error":
		lines = n.getDetails(components)
		return append(lines, Message{Body: "EBS volumes creation failed", Level: "INFO"})
	case "ebs_volumes.delete":
		return append(lines, Message{Body: "Deleting ebs volumes:", Level: "INFO"})
	case "ebs_volumes.delete.done":
		lines = n.getDetails(components)
		return append(lines, Message{Body: "EBS volumes deleted", Level: "INFO"})
	case "ebs_volumes.delete.error":
		lines = n.getDetails(components)
		return append(lines, Message{Body: "EBS volumes deletion failed", Level: "INFO"})
	case "ebs_volumes.find":
		return append(lines, Message{Body: "Importing ebs volumes:", Level: "INFO"})
	case "ebs_volumes.find.done":
		lines = n.getDetails(components)
		return append(lines, Message{Body: "EBS volumes successfully imported", Level: "INFO"})
	case "ebs_volumes.find.error":
		lines = n.getDetails(components)
		return append(lines, Message{Body: "EBS volumes import failed", Level: "INFO"})

	case "ebs_volume.create.done", "ebs_volume.create.error":
		lines = n.getSingleDetail(components, "Created EBS volume ")
	case "ebs_volume.delete.done", "ebs_volume.delete.error":
		lines = n.getSingleDetail(components, "Deleted EBS volume ")
	case "ebs_volume.find.done", "ebs_volume.find.error":
		lines = n.getSingleDetail(components, "Found EBS volume ")
	}
	return lines
}

func (n *EBSVolume) getDetails(components []interface{}) (lines []Message) {
	for _, v := range components {
		for _, l := range n.getSingleDetail(v, "") {
			lines = append(lines, l)
		}
	}
	return lines
}

func (n *EBSVolume) getSingleDetail(v interface{}, prefix string) (lines []Message) {
	r := v.(map[string]interface{})
	name, _ := r["name"].(string)
	if prefix != "" {
		name = prefix + " " + name
	}
	status, _ := r["status"].(string)
	lines = append(lines, Message{Body: " - " + name, Level: ""})
	id, _ := r["volume_aws_id"].(string)
	if id != "" {
		lines = append(lines, Message{Body: "   AWS ID : " + id, Level: ""})
	}
	lines = append(lines, Message{Body: "   Status : " + status, Level: ""})
	if status == "errored" {
		err, _ := r["error"].(string)
		lines = append(lines, Message{Body: "   Error     : " + err, Level: "ERROR"})
	}
	return lines
}
