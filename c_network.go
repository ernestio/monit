/* This Source Code Form is subject to the terms of the Mozilla Public
 * License, v. 2.0. If a copy of the MPL was not distributed with this
 * file, You can obtain one at http://mozilla.org/MPL/2.0/. */

package main

// Network : ...
type Network struct {
}

// Handle : ...
func (n *Network) Handle(subject string, components []interface{}, lines []Message) []Message {
	switch subject {
	case "networks.create":
		return append(lines, Message{Body: "Creating networks:", Level: "INFO"})
	case "networks.create.done":
		lines = n.getDetails(components)
		return append(lines, Message{Body: "Networks successfully created", Level: "INFO"})
	case "networks.create.error":
		lines = n.getDetails(components)
		return append(lines, Message{Body: "Networks creation failed", Level: "INFO"})
	case "networks.delete":
		return append(lines, Message{Body: "Deleting networks:", Level: "INFO"})
	case "networks.delete.done":
		lines = n.getDetails(components)
		return append(lines, Message{Body: "Networks deleted", Level: "INFO"})
	case "networks.delete.error":
		lines = n.getDetails(components)
		return append(lines, Message{Body: "Networks deletion failed", Level: "INFO"})
	case "networks.find":
		return append(lines, Message{Body: "Importing networks:", Level: "INFO"})
	case "networks.find.done":
		lines = n.getDetails(components)
		return append(lines, Message{Body: "Networks successfully imported", Level: "INFO"})
	case "networks.find.error":
		lines = n.getDetails(components)
		return append(lines, Message{Body: "Networks import failed", Level: "INFO"})

	case "network.create.done", "network.create.error":
		lines = n.getSingleDetail(components, "Network created")
	case "network.delete.done", "network.delete.error":
		lines = n.getSingleDetail(components, "Network deleted")
	case "network.find.done", "network.find.error":
		lines = n.getSingleDetail(components, "Network found")
	}
	return lines
}

func (n *Network) getDetails(components []interface{}) (lines []Message) {
	for _, v := range components {
		for _, l := range n.getSingleDetail(v, "") {
			lines = append(lines, l)
		}
	}
	return lines
}

func (n *Network) getSingleDetail(v interface{}, prefix string) (lines []Message) {
	r := v.(map[string]interface{})
	ip, _ := r["range"].(string)
	name, _ := r["name"].(string)
	if prefix != "" {
		name = prefix + " " + name
	}
	status, _ := r["status"].(string)
	lines = append(lines, Message{Body: " - " + name, Level: ""})
	lines = append(lines, Message{Body: "   IP     : " + ip, Level: ""})
	id, _ := r["network_aws_id"].(string)
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
