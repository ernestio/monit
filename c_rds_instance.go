/* This Source Code Form is subject to the terms of the Mozilla Public
 * License, v. 2.0. If a copy of the MPL was not distributed with this
 * file, You can obtain one at http://mozilla.org/MPL/2.0/. */

package main

// RDSInstance : ...
type RDSInstance struct {
}

// Handle : ...
func (n *RDSInstance) Handle(subject string, components []interface{}, lines []Message) []Message {
	switch subject {
	case "rds_instances.create":
		lines = append(lines, Message{Body: "Creating rds instances:", Level: "INFO"})
	case "rds_instances.create.done":
		lines = n.getDetails(components)
		return append(lines, Message{Body: "RDS instances created", Level: "INFO"})
	case "rds_instances.create.error":
		lines = n.getDetails(components)
		return append(lines, Message{Body: "RDS instances creation failed", Level: "INFO"})
	case "rds_instances.update":
		lines = append(lines, Message{Body: "Updating rds instances:", Level: "INFO"})
	case "rds_instances.update.done":
		lines = n.getDetails(components)
		return append(lines, Message{Body: "RDS instances modified", Level: "INFO"})
	case "rds_instances.update.error":
		lines = n.getDetails(components)
		return append(lines, Message{Body: "RDS instances modification failed", Level: "INFO"})
	case "rds_instances.delete":
		return append(lines, Message{Body: "Deleting rds instances:", Level: "INFO"})
	case "rds_instances.delete.done":
		lines = n.getDetails(components)
		return append(lines, Message{Body: "RDS instances deleted", Level: "INFO"})
	case "rds_instances.delete.error":
		lines = n.getDetails(components)
		return append(lines, Message{Body: "RDS instances deletion failed", Level: "INFO"})
	case "rds_instances.find":
		lines = append(lines, Message{Body: "Importing rds instances:", Level: "INFO"})
	case "rds_instances.find.done":
		lines = n.getDetails(components)
		return append(lines, Message{Body: "RDS instances imported", Level: "INFO"})
	case "rds_instances.find.error":
		lines = n.getDetails(components)
		return append(lines, Message{Body: "RDS instances import failed", Level: "INFO"})

	case "rds_instance.create.done", "rds_instance.create.error":
		lines = n.getSingleDetail(components, "RDS instance created")
	case "rds_instance.udpate.done", "rds_instance.update.error":
		lines = n.getSingleDetail(components, "RDS instance updated")
	case "rds_instance.delete.done", "rds_instance.delete.error":
		lines = n.getSingleDetail(components, "RDS instance deleted")
	case "rds_instance.find.done", "rds_instance.find.error":
		lines = n.getSingleDetail(components, "RDS instance found")
	}
	return lines
}

func (n *RDSInstance) getDetails(components []interface{}) (lines []Message) {
	for _, v := range components {
		for _, l := range n.getSingleDetail(v, "") {
			lines = append(lines, l)
		}
	}

	return lines
}

func (n *RDSInstance) getSingleDetail(v interface{}, prefix string) (lines []Message) {
	r := v.(map[string]interface{})
	name, _ := r["name"].(string)
	if prefix != "" {
		name = prefix + " " + name
	}
	engine, _ := r["engine"].(string)
	cluster, _ := r["cluster"].(string)
	endpoint, _ := r["endpoint"].(string)
	status, _ := r["status"].(string)
	lines = append(lines, Message{Body: " - " + name, Level: ""})
	lines = append(lines, Message{Body: "   Engine    : " + engine, Level: ""})
	lines = append(lines, Message{Body: "   Cluster   : " + cluster, Level: ""})
	lines = append(lines, Message{Body: "   Endpoint  : " + endpoint, Level: ""})
	if status == "errored" {
		err, _ := r["error"].(string)
		lines = append(lines, Message{Body: "   Error     : " + err, Level: "ERROR"})
	}

	return lines
}
