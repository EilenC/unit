package gee_cache

import "gee-cache/geecachepb"

// PeerPicker is the interface that must be implemented to locate
type PeerPicker interface {
	PickPeer(key string) (peer PeerGetter, ok bool)
}

// PeerGetter is the interface that must be implemented by a peer.
type PeerGetter interface {
	// Get(group string, key string) ([]byte, error)
	Get(in *geecachepb.Request, out *geecachepb.Response) error
}
