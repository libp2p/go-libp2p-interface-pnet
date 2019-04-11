package ipnet

import moved "github.com/libp2p/go-libp2p/skel/pnet"

// Deprecated: use github.com/libp2p/go-libp2p/skel/pnet.EnvKey instead.
const EnvKey = moved.EnvKey

// Deprecated: it's not possible to alias this var, as it's copied by value.
// Warning: it's not possible to alias variables in Go. Setting a value here may have no effect.
var ForcePrivateNetwork = moved.ForcePrivateNetwork

// Deprecated: use github.com/libp2p/go-libp2p/skel/pnet.ErrNotInPrivateNetwork instead.
var ErrNotInPrivateNetwork = moved.ErrNotInPrivateNetwork

// Deprecated: use github.com/libp2p/go-libp2p/skel/pnet.Protector instead.
type Protector = moved.Protector

// Deprecated: use github.com/libp2p/go-libp2p/skel/pnet.Error instead.
type PNetError = moved.Error

// Deprecated: use github.com/libp2p/go-libp2p/skel/pnet.NewError instead.
var NewError = moved.NewError

// Deprecated: use github.com/libp2p/go-libp2p/skel/pnet.IsPNetError instead.
var IsPNetError = moved.IsPNetError
