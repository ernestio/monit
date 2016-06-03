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
		s = sse.New()
		defer s.Close()

		mux := http.NewServeMux()
		mux.HandleFunc("/events", s.HTTPHandler)
		hs := httptest.NewServer(mux)
		url := hs.URL + "/events"

		Convey("When listening for NATS messages", func() {
			createEvents := []string{"service.create", "service.delete"}
			for _, event := range createEvents {
				Convey("On receiving "+event, func() {
					msg := nats.Msg{Subject: event, Data: []byte(`{"service": "test"}`)}
					natsHandler(&msg)

					time.Sleep(time.Millisecond * 10)

					Convey("It should create a stream for the service", func() {
						So(s.StreamExists("test"), ShouldBeTrue)
					})

				})
			}

			deleteEvents := []string{"service.create.done", "service.delete.done", "service.create.error", "service.delete.error"}
			for _, event := range deleteEvents {
				s.CreateStream("test")
				time.Sleep(time.Millisecond * 10)

				Convey("On receiving "+event, func() {
					msg := nats.Msg{Subject: event, Data: []byte(`{"service": "test"}`)}
					natsHandler(&msg)

					time.Sleep(time.Millisecond * 1500)

					Convey("It should remove the services stream", func() {
						So(s.StreamExists("test"), ShouldBeFalse)
					})

				})
			}

			Convey("On receiving an unknown message", func() {
				// Clean server
				s.RemoveStream("test")
				time.Sleep(time.Millisecond * 10)

				msg := nats.Msg{Subject: "test.event", Data: []byte(`{"service": "test"}`)}
				natsHandler(&msg)

				Convey("It should not create a stream", func() {
					So(s.StreamExists("test"), ShouldBeFalse)
				})
			})

			Convey("When receiving monitor.user", func() {
				testEvent := `{"service": "test", "messages":[{"body": "test", "color": "blue"}]}`
				msg := nats.Msg{Subject: "monitor.user", Data: []byte(testEvent)}

				Convey("And a stream exists", func() {
					rcv := make(chan *sse.Event)
					cl := sse.NewClient(url)

					s.CreateStream("test")
					time.Sleep(time.Millisecond * 10)

					go cl.SubscribeChan("test", rcv)
					time.Sleep(time.Millisecond * 10)

					Convey("It should publish a message to the stream", func() {
						natsHandler(&msg)

						event, err := wait(rcv, time.Millisecond*100)

						So(err, ShouldBeNil)
						So(string(event.Data), ShouldEqual, `{"body":"test","level":""}`)
					})
				})

				s.RemoveStream("test")

				Convey("And a stream doesn't exist", func() {
					rcv := make(chan *sse.Event)
					cl := sse.NewClient(url)

					time.Sleep(time.Millisecond * 10)

					go cl.SubscribeChan("test", rcv)
					time.Sleep(time.Millisecond * 10)

					Convey("It should not publish a message to the stream", func() {
						natsHandler(&msg)
						event, err := wait(rcv, time.Millisecond*100)

						So(err, ShouldBeNil)
						if err != nil {
							So(string(event.Data), ShouldNotEqual, `{"body":"test","color":"blue"}`)
						}
					})
				})
			})

		})
	})

}
