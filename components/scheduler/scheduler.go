package scheduler

import "context"
import "github.com/ipfs-search/ipfs-search-2/types/resource"

// Interface represents a scheduler based on a number of resource input channels
type Interface interface {
	Start(context.Context, ...<-chan resoure.SniffedResource) (<-chan resource.SniffedResource, <-chan error, error)
}
