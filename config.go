/* This Source Code Form is subject to the terms of the Mozilla Public
 * License, v. 2.0. If a copy of the MPL was not distributed with this
 * file, You can obtain one at http://mozilla.org/MPL/2.0/. */

package main

type monitorConfig struct {
	Host string `json:"host"`
	Port string `json:"port"`
}

type redisConfig struct {
	Host     string `json:"addr"`
	Password string `json:"password"`
	DB       int64  `json:"DB"`
}
