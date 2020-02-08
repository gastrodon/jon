package main

import (
	"github.com/gastrodon/jon/jon"

	"github.com/turnage/graw/reddit"
)

func main() {
	var client reddit.Bot
	var err error
	client, err = reddit.NewBotFromAgentFile("agent", 0)
	if err != nil {
		panic(err)
	}

	var kill_chan chan bool = make(chan bool)
	var err_chan chan error = make(chan error)
	go jon.HandleComments("garfieldminusgarfield+garfieldminusjon", client, kill_chan, err_chan)
	go jon.LogErrs(err_chan)

	select {
	case <-kill_chan:
		return
	}
}
