//////////////////////////////////////////////////////////////////////
//
// Given is a producer-consumer scenario, where a producer reads in
// tweets from a mockstream and a consumer is processing the
// data. Your task is to change the code so that the producer as well
// as the consumer can run concurrently
//

package main

import (
	"fmt"
	"time"
)

func producer(stream Stream, tweetBuffer chan *Tweet) {
	for {
		tweet, err := stream.Next()
		if err == ErrEOF {
			close(tweetBuffer)
			return
		}

		tweetBuffer <- tweet
	}
}

func consumer(tweetBuffer chan *Tweet) {
	for t := range tweetBuffer {
		if t.IsTalkingAboutGo() {
			fmt.Println(t.Username, "\ttweets about golang")
		} else {
			fmt.Println(t.Username, "\tdoes not tweet about golang")
		}
	}
}

func main() {
	start := time.Now()
	stream := GetMockStream()
	tweetBuffer := make(chan *Tweet)

	// Producer
	go func() {
		producer(stream, tweetBuffer)
	}()

	// Consumer
	consumer(tweetBuffer)

	fmt.Printf("Process took %s\n", time.Since(start))
}
