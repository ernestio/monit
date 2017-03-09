/* This Source Code Form is subject to the terms of the Mozilla Public
 * License, v. 2.0. If a copy of the MPL was not distributed with this
 * file, You can obtain one at http://mozilla.org/MPL/2.0/. */

package main

import "strings"

// Instance : ...
type Instance struct {
}

// Handle : ...
func (n *Instance) Handle(subject string, component interface{}, lines []Message) []Message {
	parts := strings.Split(subject, ".")
	subject = parts[0] + "." + parts[1]
	switch subject {
	case "instance.create":
		lines = n.getSingleDetail(component, "Instance created")
	case "instance.update":
		lines = n.getSingleDetail(component, "Instance udpated")
	case "instance.delete":
		lines = n.getSingleDetail(component, "Instance delete")
	case "instance.find":
		lines = n.getSingleDetail(component, "Instance find")
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
