package indexer

import "context"
import "github.com/ipfs-search/ipfs-search-2/types/resource"

// Interface represents an Indexer for channels of resource data.
type Interface interface {
	Start(context.Context, <-chan resoure.Data) (<-chan resource.ResourceError, error)
}
