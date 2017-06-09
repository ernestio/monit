package main

import (
	"log"
	"time"

	"github.com/r3labs/sse"
)

func publishEvent(id string, data []byte) {

	//	id := svr.getID()

	//	data, err := json.Marshal(svr)
	//	if err != nil {
	//		panic(err)
	//	}

	//fmt.Printf("mydata = %+v\n", string(data))

	// Create a new stream
	log.Println("Creating stream for", id)
	ss.CreateStream(id)

	ss.Publish(id, data)

	time.Sleep(10 * time.Millisecond)
	// Remove a new stream when the build completes
	log.Println("Closing stream for", id)
	go func(s *sse.Server) {
		ss.RemoveStream(id)
	}(ss)

	//	publishMessage(notification.getServiceID(), &nm)
}
