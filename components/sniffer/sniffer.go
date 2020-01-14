package sniffer

import "context"
import "github.com/ipfs-search/ipfs-search-2/types/resource"

// Interface for sniffer
type Interface interface {
	Start(context.Context, <-chan resource.Resource, <-chan resource.ResourceError, error)
}
