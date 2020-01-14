package sniffer

import "context"
import "time"
import "github.com/ipfs-search/ipfs-search-2/types/resource"

// Interface for sniffer
type Interface interface {
	Start(context.Context, <-chan resource.SniffedResource) (<-chan resource.ResourceError, error)
}

// SniffedResource represents a resource as it's sniffed from the dweb
type SniffedResource struct {
	*Resource
	time.Time
	Peer string
}
