package commands

import (
	"time"

	"github.com/foxtrottwist/pokego/client"
)

type Writer = func(string) error

type config struct {
	client   client.Client
	next     *string
	previous *string
	write    Writer
}

func NewConfig(writer Writer, timeout, interval time.Duration) *config {
	return &config{
		client: client.New(timeout, interval),
		write:  writer,
	}
}
