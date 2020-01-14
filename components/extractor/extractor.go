package crawler

import "context"
import "github.com/ipfs-search/ipfs-search-2/types/resource"

// Interface represents an extractor for (meta)data from resources on the dweb.
type Interface interface {
	Extract(context.Context, resource.Resource) (resource.Data, error)
}

// ExtractChannel indexes a channel of resource data in a blocking way, returning error on failure.
func ExtractChannel(ctx context.Context, i Interface, resourcec <-chan resource.Resource, datac chan<- resource.Data, errc chan<- error) error {
	for {
		select {
		case <-ctx.Done():
			// Context closed, return
			return ctx.Err()

		case r := <-resourcec:
			data, err := i.Extract(ctx, r)

			if err != nil {
				errc <- &resource.Error{
					Err:      err,
					Resource: &r,
				}
			}

			datac <- data
		}
	}

}
