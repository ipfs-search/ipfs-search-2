package queue

import "github.com/ipfs-search/ipfs-search-2/types/resource"

// Queue represents a generic queue type for resource objects.
type Queue interface {
	Add(Resource) Error
	Pop() Resource
}

// New returns a new Queue object with given identifier.
func New(id string) Queue {
}
