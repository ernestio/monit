/* This Source Code Form is subject to the terms of the Mozilla Public
 * License, v. 2.0. If a copy of the MPL was not distributed with this
 * file, You can obtain one at http://mozilla.org/MPL/2.0/. */

package main

// S3Bucket : ...
type S3Bucket struct {
}

// Handle : ...
func (n *S3Bucket) Handle(subject string, components []interface{}, lines []Message) []Message {
	switch subject {
	case "s3.create.done", "s3.create.error":
		lines = n.getSingleDetail(components, "S3 bucket created")
	case "s3.update.done", "s3.update.error":
		lines = n.getSingleDetail(components, "S3 bucket updated")
	case "s3.delete.done", "s3.delete.error":
		lines = n.getSingleDetail(components, "S3 bucket deleted")
	case "s3.find.done", "s3.find.error":
		lines = n.getSingleDetail(components, "S3 bucket imported")
	}
	return lines
}

func (n *S3Bucket) getSingleDetail(v interface{}, prefix string) (lines []Message) {
	r := v.(map[string]interface{})
	name, _ := r["name"].(string)
	if prefix != "" {
		name = prefix + " " + name
	}
	acl, _ := r["acl"].(string)
	if acl == "" {
		acl = "by grantees"
	}
	status, _ := r["status"].(string)
	lines = append(lines, Message{Body: " - " + name, Level: ""})
	lines = append(lines, Message{Body: "   ACL       : " + acl, Level: ""})
	lines = append(lines, Message{Body: "   Status    : " + status, Level: ""})
	if status == "errored" {
		err, _ := r["error"].(string)
		lines = append(lines, Message{Body: "   Error     : " + err, Level: "ERROR"})
	}
	return lines
}
