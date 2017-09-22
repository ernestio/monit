/* This Source Code Form is subject to the terms of the Mozilla Public
 * License, v. 2.0. If a copy of the MPL was not distributed with this
 * file, You can obtain one at http://mozilla.org/MPL/2.0/. */

package main

import (
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/nats-io/nats"
	"github.com/r3labs/sse"
	. "github.com/smartystreets/goconvey/convey"
)

func wait(ch chan *sse.Event, duration time.Duration) (*sse.Event, error) {
	var err error
	var msg *sse.Event

	select {
	case event := <-ch:
		msg = event
	case <-time.After(duration):
		err = errors.New("timeout")
	}
	return msg, err
}

func TestMain(t *testing.T) {

	Convey("Given a new server", t, func() {
		// New Server
		ss = sse.New()
		defer ss.Close()

		mux := http.NewServeMux()
		mux.HandleFunc("/events", ss.HTTPHandler)
		hs := httptest.NewServer(mux)
		url := hs.URL + "/events"

		Convey("When listening for NATS messages", func() {
			createEvents := []string{"build.create", "build.delete", "build.import"}
			for _, event := range createEvents {
				Convey("On receiving "+event, func() {
					msg := nats.Msg{Subject: event, Data: []byte(`{"id": "test"}`)}
					natsHandler(&msg)

					time.Sleep(time.Millisecond * 10)

					Convey("It should create a stream for the service", func() {
						So(ss.StreamExists("test"), ShouldBeTrue)
					})

				})
			}

			deleteEvents := []string{"build.create.done", "build.delete.done", "build.import.done", "build.create.error", "build.delete.error", "build.import.error"}
			for _, event := range deleteEvents {
				ss.CreateStream("test")
				time.Sleep(time.Millisecond * 10)

				Convey("On receiving "+event, func() {
					msg := nats.Msg{Subject: event, Data: []byte(`{"id": "test"}`)}
					natsHandler(&msg)

					time.Sleep(time.Millisecond * 1500)

					Convey("It should remove the services stream", func() {
						So(ss.StreamExists("test"), ShouldBeFalse)
					})

				})
			}

			Convey("On receiving an unknown message", func() {
				// Clean server
				ss.RemoveStream("test")
				time.Sleep(time.Millisecond * 10)

				msg := nats.Msg{Subject: "test.event", Data: []byte(`{"id": "test"}`)}
				natsHandler(&msg)

				Convey("It should not create a stream", func() {
					So(ss.StreamExists("test"), ShouldBeFalse)
				})
			})

			Convey("When receiving component event network.create.aws.done", func() {
				testEvent := `{"service": "test", "name": "network"}`
				msg := nats.Msg{Subject: "network.create.aws.done", Data: []byte(testEvent)}

				Convey("And a stream exists", func() {
					rcv := make(chan *sse.Event)
					cl := sse.NewClient(url)

					ss.CreateStream("test")
					time.Sleep(time.Millisecond * 10)

					go func() {
						_ = cl.SubscribeChan("test", rcv)
					}()
					time.Sleep(time.Millisecond * 10)

					Convey("It should publish a message to the stream", func() {
						natsHandler(&msg)

						for {
							event, err := wait(rcv, time.Millisecond*100)
							So(err, ShouldBeNil)

							if len(event.Data) > 0 {
								So(string(event.Data), ShouldEqual, `{"_component_id":"","_subject":"network.create.aws.done","_component":"","_state":"","_action":"","_provider":"","name":"network","service":"test"}`)
								break
							}
						}
					})
				})

				ss.RemoveStream("test")

				Convey("And a stream doesn't exist", func() {
					rcv := make(chan *sse.Event)
					cl := sse.NewClient(url)

					time.Sleep(time.Millisecond * 10)

					Convey("It should error when connecting to the stream", func() {
						err := cl.SubscribeChan("test", rcv)
						So(err, ShouldNotBeNil)
					})
				})
			})

		})
	})

}
