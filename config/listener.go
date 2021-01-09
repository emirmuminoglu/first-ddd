package config

import (
	"net"
	"os"
)

func NewListener() (net.Listener, error) {
	network := os.Getenv("LN_NETWORK")
	addr := os.Getenv("LN_ADDR")

	return net.Listen(network, addr)
}
