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
	case "s3s.create":
		lines = append(lines, Message{Body: "Creating s3 buckets:", Level: "INFO"})
	case "s3s.create.done":
		lines = n.getDetails(components)
		return append(lines, Message{Body: "S3 buckets created", Level: "INFO"})
	case "s3s.create.error":
		lines = n.getDetails(components)
		return append(lines, Message{Body: "S3 buckets creation failed", Level: "INFO"})
	case "s3s.update":
		lines = append(lines, Message{Body: "Updating s3 buckets:", Level: "INFO"})
	case "s3s.update.done":
		lines = n.getDetails(components)
		return append(lines, Message{Body: "S3 buckets modified", Level: "INFO"})
	case "s3s.update.error":
		lines = n.getDetails(components)
		return append(lines, Message{Body: "S3 buckets modification failed", Level: "INFO"})
	case "s3s.delete":
		return append(lines, Message{Body: "Deleting s3 buckets:", Level: "INFO"})
	case "s3s.delete.done":
		lines = n.getDetails(components)
		return append(lines, Message{Body: "S3 buckets deleted", Level: "INFO"})
	case "s3s.delete.error":
		lines = n.getDetails(components)
		return append(lines, Message{Body: "S3 buckets deletion failed", Level: "INFO"})
	case "s3s.find":
		lines = append(lines, Message{Body: "Importing s3 buckets:", Level: "INFO"})
	case "s3s.find.done":
		lines = n.getDetails(components)
		return append(lines, Message{Body: "S3 buckets imported", Level: "INFO"})
	case "s3s.find.error":
		lines = n.getDetails(components)
		return append(lines, Message{Body: "S3 buckets import failed", Level: "INFO"})

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

func (n *S3Bucket) getDetails(components []interface{}) (lines []Message) {
	for _, v := range components {
		for _, l := range n.getSingleDetail(v, "") {
			lines = append(lines, l)
		}
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
