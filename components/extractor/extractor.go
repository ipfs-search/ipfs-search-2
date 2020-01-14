package crawler

import "context"
import "github.com/ipfs-search/ipfs-search-2/types/resource"

// Interface represents an extractor for (meta)data from resources on the dweb.
type Interface interface {
	Start(context.Context, <-chan resource.Resource) (<-chan resource.Data, <-chan resource.ResourceError, error)
}
