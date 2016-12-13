/* This Source Code Form is subject to the terms of the Mozilla Public
 * License, v. 2.0. If a copy of the MPL was not distributed with this
 * file, You can obtain one at http://mozilla.org/MPL/2.0/. */

package main

type EBSVolume struct {
}

func (v *EBSVolume) Handle(subject string, components []interface{}, lines []Message) []Message {
	switch subject {
	case "ebs_volumes.create":
		return append(lines, Message{Body: "Creating ebs volumes:", Level: "INFO"})
	case "ebs_volumes.create.done":
		lines = v.getDetails(components)
		return append(lines, Message{Body: "EBS volumes successfully created", Level: "INFO"})
	case "ebs_volumes.create.error":
		lines = v.getDetails(components)
		return append(lines, Message{Body: "EBS volumes creation failed", Level: "INFO"})
	case "ebs_volumes.delete":
		return append(lines, Message{Body: "Deleting ebs volumes:", Level: "INFO"})
	case "ebs_volumes.delete.done":
		lines = v.getDetails(components)
		return append(lines, Message{Body: "EBS volumes deleted", Level: "INFO"})
	case "ebs_volumes.delete.error":
		lines = v.getDetails(components)
		return append(lines, Message{Body: "EBS volumes deletion failed", Level: "INFO"})
	}
	return lines
}

func (v *EBSVolume) getDetails(components []interface{}) (lines []Message) {
	for _, v := range components {
		r := v.(map[string]interface{})
		name, _ := r["name"].(string)
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
	}
	return lines
}
