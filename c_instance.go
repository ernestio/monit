/* This Source Code Form is subject to the terms of the Mozilla Public
 * License, v. 2.0. If a copy of the MPL was not distributed with this
 * file, You can obtain one at http://mozilla.org/MPL/2.0/. */

package main

type Instance struct {
}

func (n *Instance) Handle(subject string, components []interface{}, lines []Message) []Message {
	switch subject {
	case "instances.create":
		lines = append(lines, Message{Body: "Creating instances:", Level: "INFO"})
	case "instances.create.done":
		lines = n.getDetails(components)
		return append(lines, Message{Body: "Instances successfully created", Level: "INFO"})
	case "instances.create.error":
		lines = n.getDetails(components)
		return append(lines, Message{Body: "Instances creation failed", Level: "INFO"})
	case "instances.update":
		return append(lines, Message{Body: "Updating instances:", Level: "INFO"})
	case "instances.update.done":
		lines = n.getDetails(components)
		return append(lines, Message{Body: "Instances successfully updated", Level: "INFO"})
	case "instances.update.error":
		lines = n.getDetails(components)
		return append(lines, Message{Body: "Instances modification failed", Level: "INFO"})
	case "instances.delete":
		return append(lines, Message{Body: "Deleting instances:", Level: "INFO"})
	case "instances.delete.done":
		lines = n.getDetails(components)
		return append(lines, Message{Body: "Instances deleted", Level: "INFO"})
	case "instances.delete.error":
		lines = n.getDetails(components)
		return append(lines, Message{Body: "Instances deletion failed", Level: "INFO"})
	}
	return lines
}

func (n *Instance) getDetails(components []interface{}) (lines []Message) {
	for _, v := range components {
		r := v.(map[string]interface{})
		ip, _ := r["ip"].(string)
		name, _ := r["name"].(string)
		status, _ := r["status"].(string)
		lines = append(lines, Message{Body: " - " + name, Level: ""})
		lines = append(lines, Message{Body: "   IP        : " + ip, Level: ""})
		public_ip, _ := r["public_ip"].(string)
		if public_ip != "" {
			lines = append(lines, Message{Body: "   PUBLIC IP : " + public_ip, Level: ""})
		}
		id, _ := r["instance_aws_id"].(string)
		if id != "" {
			lines = append(lines, Message{Body: "   AWS ID    : " + id, Level: ""})
		}
		lines = append(lines, Message{Body: "   Status    : " + status, Level: ""})
		if status == "errored" {
			err, _ := r["error_message"].(string)
			lines = append(lines, Message{Body: "   Error     : " + err, Level: "ERROR"})
		}
	}

	return lines
}
