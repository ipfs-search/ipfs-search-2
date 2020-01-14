package crawler

import "context"
import "github.com/ipfs-search/ipfs-search-2/types/resource"

// Interface represents a crawler for resources on the dweb.
type Interface interface {
	Start(context.Context, <-chan resource.Resource) (<-chan resource.ResourceError, error)
}
