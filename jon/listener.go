package jon

import (
	"github.com/turnage/graw/reddit"
	"github.com/turnage/graw/streams"

	"fmt"
	"log"
)

func dispatchReply(client reddit.Bot, comment *reddit.Comment, err_chan chan<- error) {
	var reply string = ReplyFor(comment.Body)
	var err error

	if reply == "" {
		err = fmt.Errorf("%s is valid but spawned no reply", comment.Body)
	} else {
		err = client.Reply(comment.Name, ReplyFor(comment.Body))
	}

	if err != nil {
		err_chan <- err
	}
}

func HandleComments(subs string, client reddit.Bot, kill_chan chan bool, err_chan chan<- error) {
	var stream <-chan *reddit.Comment
	var err error
	stream, err = streams.SubredditComments(client, kill_chan, err_chan, subs)
	if err != nil {
		panic(err)
	}

	var comment *reddit.Comment
	for comment = range stream {
		if Valid(comment.Body) {
			go dispatchReply(client, comment, err_chan)
		}
	}
}

func LogErrs(err_chan chan error) {
	var err error
	for err = range err_chan {
		log.Println(err)
	}
}
