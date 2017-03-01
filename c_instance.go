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
