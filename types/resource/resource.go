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

// ResourceError wraps an error with a Resource
type ResourceError struct {
	Err      error
	Resource *resource.Resource
}

// Unwrap yields the underlying error for a ResourceError.
func (e *ResourceError) Unwrap() error { return e.wrappederr }

func (e *ResourceError) Error() string {
	return fmt.Sprintf("%v with %v", e.Err, e.Resource)
}
