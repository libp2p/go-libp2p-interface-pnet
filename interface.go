package pnet

import (
	"errors"
	"os"

	iconn "github.com/libp2p/go-libp2p-interface-conn"
)

const envKey = "LIBP2P_FORCE_PNET"

var ErrNotInPrivateNetwork = errors.New("private network was not configured")

type Protector interface {
	Protect(iconn.Conn) (iconn.Conn, error)
}

func ShouldForcePrivateNetwork() bool {
	v := os.Getenv(envKey)
	return v == "1"
}
