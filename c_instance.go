/* This Source Code Form is subject to the terms of the Mozilla Public
 * License, v. 2.0. If a copy of the MPL was not distributed with this
 * file, You can obtain one at http://mozilla.org/MPL/2.0/. */

package main

// Instance : ...
type Instance struct {
}

// Handle : ...
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
	case "instances.find":
		lines = append(lines, Message{Body: "Importing instances:", Level: "INFO"})
	case "instances.find.done":
		lines = n.getDetails(components)
		return append(lines, Message{Body: "Instances successfully imported", Level: "INFO"})
	case "instances.find.error":
		lines = n.getDetails(components)
		return append(lines, Message{Body: "Instances import failed", Level: "INFO"})

	case "instance.create.done", "instance.create.error":
		lines = n.getSingleDetail(components, "Instance created")
	case "instance.update.done", "instance.update.error":
		lines = n.getSingleDetail(components, "Instance udpated")
	case "instance.delete.done", "instance.delete.error":
		lines = n.getSingleDetail(components, "Instance delete")
	case "instance.find.done", "instance.find.error":
		lines = n.getSingleDetail(components, "Instance find")
	}
	return lines
}

func (n *Instance) getDetails(components []interface{}) (lines []Message) {
	for _, v := range components {
		for _, l := range n.getSingleDetail(v, "") {
			lines = append(lines, l)
		}
	}

	return lines
}

func (n *Instance) getSingleDetail(v interface{}, prefix string) (lines []Message) {
	r := v.(map[string]interface{})
	ip, _ := r["ip"].(string)
	name, _ := r["name"].(string)
	if prefix != "" {
		name = prefix + " " + name
	}
	status, _ := r["status"].(string)
	lines = append(lines, Message{Body: " - " + name, Level: ""})
	lines = append(lines, Message{Body: "   IP        : " + ip, Level: ""})
	publicIP, _ := r["public_ip"].(string)
	if publicIP != "" {
		lines = append(lines, Message{Body: "   PUBLIC IP : " + publicIP, Level: ""})
	}
	id, _ := r["instance_aws_id"].(string)
	if id != "" {
		lines = append(lines, Message{Body: "   AWS ID    : " + id, Level: ""})
	}
	lines = append(lines, Message{Body: "   Status    : " + status, Level: ""})
	if status == "errored" {
		err, _ := r["error"].(string)
		lines = append(lines, Message{Body: "   Error     : " + err, Level: "ERROR"})
	}
	return lines
}
