package indexer

import "context"
import "github.com/ipfs-search/ipfs-search-2/types/resource"

// Interface represents an Indexer for channels of resource data.
type Interface interface {
	Index(context.Context, resource.Data) error
}

// IndexChannel indexes a channel of resource data in a blocking way, returning error on failure.
func IndexChannel(ctx context.Context, i Interface, datac <-chan resource.Data, errc chan<- error) error {
	for {
		select {
		case <-ctx.Done():
			// Context closed, return
			return ctx.Err()

		case data := <-datac:
			err := i.Index(ctx, data)

			if err != nil {
				errc <- &resource.Error{
					Err:      err,
					Resource: data.Resource,
				}
			}
		}
	}

}
