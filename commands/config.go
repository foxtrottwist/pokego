package commands

import (
	"time"

	"github.com/foxtrottwist/pokego/client"
)

type printer = func(string) error

type config struct {
	client   client.Client
	next     *string
	previous *string
	print    printer
}

func NewConfig(p printer, timeout, interval time.Duration) *config {
	return &config{
		client: client.New(timeout, interval),
		print:  p,
	}
}
