package pnet

import (
	"errors"
	"os"

	iconn "github.com/libp2p/go-libp2p-interface-conn"
)

const envKey = "LIBP2P_FORCE_PNET"

// Setting this variable to true or setting LIBP2P_FORCE_PNET environment variable
// to true will make libp2p to require private network protector.
// If no network protector is provided and this variable is set to true libp2p will
// refuse to connect.
var ForcePrivateNetwork bool = false

func init() {
	ForcePrivateNetwork = os.Getenv(envKey) == "1"
}

var ErrNotInPrivateNetwork = errors.New("private network was not configured but" +
	" is enforced by the environment")

// This interface is a way for private network implementation to be transparent in
// libp2p. It is created by implementation and use by libp2p-conn to secure connections
// so they can be only established with selected number of peers.
type Protector interface {
	Protect(iconn.Conn) (iconn.Conn, error)
}
