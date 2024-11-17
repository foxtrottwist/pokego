package commands

import (
	"time"

	"github.com/foxtrottwist/pokego/client"
)

type config struct {
	client.Client
	next     *string
	previous *string
}

func NewConfig(timeout, interval time.Duration) *config {
	return &config{
		Client: client.New(timeout, interval),
	}
}
