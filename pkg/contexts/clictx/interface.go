package clictx

import (
	"github.com/open-component-model/ocm/pkg/contexts/clictx/internal"
)

type (
	Context = internal.Context
	OCI     = internal.OCI
	OCM     = internal.OCM
)

func DefaultContext() Context {
	return internal.DefaultContext
}
