package crawler

import "context"
import "fmt"
import "github.com/ipfs-search/ipfs-search-2/types/resource"
import "github.com/ipfs-search/ipfs-search-2/components/extractor"
import "github.com/ipfs-search/ipfs-search-2/components/indexer"

// Crawler extracts data from resources and indexes them.
type Crawler struct {
	s *sniffer.Interface
	e *extractor.Interface
	i *indexer.Interface
}

// Start (blocking) crawling channel of resources, returning error if starting fails and ResourceError during crawling.
func (c *Crawler) Start(ctx context.Context, resources <-chan resource.Resource, errc chan<- error) error {
	crawlerContext, cancel := context.WithCancel(ctx)

	// Start the extractor
	data, extractorErrors, err := e.Start(crawlerContext, resources)
	if err != nil {
		return err
	}

	// Start the indexer
	indexerErrors, err := i.Start(crawlerContext, data)
	if err != nil {
		// Cancel context, stopping extractor
		cancel()

		return err
	}

	// Start crawling
	go func() {
		for {
			select {
			case <-crawlerContext.Done():
				// Context closed, return
				return crawlerContext.Err()

			case err := <-extractorErrors:
				// Non-fatal extractor error
				errc <- fmt.Errorf("Extractor error: %w", err)

			case err := <-indexerErrors:
				// Non-fatal indexer error
				errc <- fmt.Errorf("Indexing error: %w", err)
			}
		}
	}()
}

// New returns a new Crawler.
func New(extractor *extractor.Interface, indexer *indexer.Interface) *Crawler {
	return &Crawler{
		e: extractor,
		i: indexer,
	}
}
