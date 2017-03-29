/* This Source Code Form is subject to the terms of the Mozilla Public
 * License, v. 2.0. If a copy of the MPL was not distributed with this
 * file, You can obtain one at http://mozilla.org/MPL/2.0/. */

package main

import (
	"encoding/json"
	"log"
	"strings"

	"github.com/nats-io/nats"
)

type component map[string]interface{}

func (m *component) getServiceID() string {
	id, ok := (*m)["service"].(string)
	if ok {
		return id
	}
	return ""
}

func (m *component) getServicePart() string {
	pieces := strings.Split(m.getServiceID(), "-")
	return pieces[len(pieces)-1]
}

func (m *component) getFoundComponents() []component {
	var c []component

	components, ok := (*m)["components"].([]interface{})
	if ok {
		for _, x := range components {
			c = append(c, x.(map[string]interface{}))
		}
	}

	return c
}

func genericHandler(msg *nats.Msg) {
	var msgLines []Message
	var c component

	if err := json.Unmarshal(msg.Data, &c); err != nil {
		return
	}

	parts := strings.Split(msg.Subject, ".")
	component := parts[0]

	switch component {
	case "ebs_volumes", "ebs_volume":
		var nt EBSVolume
		msgLines = nt.Handle(msg.Subject, c, msgLines)
	case "instances", "instance":
		var nt Instance
		msgLines = nt.Handle(msg.Subject, c, msgLines)
	case "networks", "network":
		var nt Network
		msgLines = nt.Handle(msg.Subject, c, msgLines)
	case "firewalls", "firewall":
		var nt Firewall
		msgLines = nt.Handle(msg.Subject, c, msgLines)
	case "nats", "nat":
		var nt Nat
		msgLines = nt.Handle(msg.Subject, c, msgLines)
	case "routers", "router":
		var nt Router
		msgLines = nt.Handle(msg.Subject, c, msgLines)
	case "vpcs", "vpc":
		var nt Vpc
		msgLines = nt.Handle(msg.Subject, c, msgLines)
	case "elbs", "elb":
		var nt ELB
		msgLines = nt.Handle(msg.Subject, c, msgLines)
	case "s3s", "s3":
		var nt S3Bucket
		msgLines = nt.Handle(msg.Subject, c, msgLines)
	case "rds_clusters", "rds_cluster":
		var nt RDSCluster
		msgLines = nt.Handle(msg.Subject, c, msgLines)
	case "rds_instances", "rds_instance":
		var nt RDSInstance
		msgLines = nt.Handle(msg.Subject, c, msgLines)
	default:
		log.Println("unsupported: " + msg.Subject)
	}
	for _, v := range msgLines {
		publishMessage(c.getServicePart(), &v)
	}
}

func genericErrorMessageHandler(components []interface{}, cType, cAction string) (lines []Message) {
	for _, c := range components {
		component := c.(map[string]interface{})
		if component["status"].(string) == "errored" {
			name := component["name"].(string)
			msg := component["error"].(string)
			msg = strings.Replace(msg, ":", " -", -1)
			line := cType + " " + name + " " + cAction + " failed with: \n" + msg
			lines = append(lines, Message{Body: line, Level: "ERROR"})
		}
	}

	return lines
}
