package v1

// ResourceReference describes re resource identity relative to an (aggregation)
// component version.
type ResourceReference struct {
	Resource      Identity   `json:"resource,omitempty"`
	ReferencePath []Identity `json:"referencePath,omitempty"`
}

func NewResourceRef(id Identity, path ...Identity) ResourceReference {
	return ResourceReference{Resource: id, ReferencePath: path}
}

func NewNestedResourceRef(id Identity, path []Identity) ResourceReference {
	return ResourceReference{Resource: id, ReferencePath: path}
}

func (r *ResourceReference) String() string {
	s := r.Resource.String()

	for i := 1; i <= len(r.ReferencePath); i++ {
		s += "@" + r.ReferencePath[len(r.ReferencePath)-i].String()
	}
	return s
}
