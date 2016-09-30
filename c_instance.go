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
		return append(lines, Message{Body: "Instances successfully created", Level: "INFO"})
	case "instances.update":
		return append(lines, Message{Body: "Updating instances:", Level: "INFO"})
	case "instances.update.done":
		lines = n.getDetails(components)
		return append(lines, Message{Body: "Instances successfully updated", Level: "INFO"})
	case "instances.update.error":
	case "instances.delete":
		return append(lines, Message{Body: "Deleting instances", Level: "INFO"})
	case "instances.delete.done":
		return append(lines, Message{Body: "Instances deleted", Level: "INFO"})
	case "instances.delete.error":
	}
	return lines
}

func (n *Instance) getDetails(components []interface{}) (lines []Message) {
	for _, v := range components {
		r := v.(map[string]interface{})
		ip := r["ip"].(string)
		name := r["name"].(string)
		status := r["status"].(string)
		lines = append(lines, Message{Body: " - " + name, Level: ""})
		lines = append(lines, Message{Body: "   IP        : " + ip, Level: ""})
		public_ip := r["public_ip"].(string)
		if public_ip != "" {
			lines = append(lines, Message{Body: "   PUBLIC IP : " + public_ip, Level: ""})
		}
		id := r["instance_aws_id"].(string)
		if id != "" {
			lines = append(lines, Message{Body: "   AWS ID    : " + id, Level: ""})
		}
		lines = append(lines, Message{Body: "   Status    : " + status, Level: ""})
	}

	return lines
}
