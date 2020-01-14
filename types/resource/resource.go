package resource

import "fmt"

// Resource represents a resource on the dweb.
type Resource struct {
	scheme, cid string
}

// New returns a new Resource given a scheme and cid.
func New(scheme, cid string) Resource {
	return &Resource{
		scheme: scheme,
		cid:    cid,
	}
}

// Hash generates a unique hash for the resource
func (r *Resource) Hash() string {
	return fmt.Sprintf("%s-%s", r.scheme, r.cid)
}
