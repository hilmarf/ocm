package elemhdlr

import (
	"github.com/Masterminds/semver/v3"

	"github.com/open-component-model/ocm/cmds/ocm/commands/ocmcmds/common/handlers/comphdlr"
	"github.com/open-component-model/ocm/cmds/ocm/commands/ocmcmds/common/options/lookupoption"
	"github.com/open-component-model/ocm/cmds/ocm/commands/ocmcmds/common/options/versionconstraintsoption"
	"github.com/open-component-model/ocm/cmds/ocm/pkg/options"
	"github.com/open-component-model/ocm/pkg/contexts/ocm"
)

type Option interface {
	ApplyToElemHandler(handler *TypeHandler)
}

type Options []Option

func (o Options) ApplyToElemHandler(handler *TypeHandler) {
	for _, e := range o {
		e.ApplyToElemHandler(handler)
	}
}

func OptionsFor(o options.OptionSetProvider) Options {
	var hopts []Option
	constr := versionconstraintsoption.From(o)
	if len(constr.Constraints) > 0 {
		hopts = append(hopts, WithVersionConstraints(constr.Constraints))
	}
	if constr.Latest {
		hopts = append(hopts, LatestOnly())
	}
	if lookup := lookupoption.From(o); lookup != nil {
		hopts = append(hopts, Resolver(lookup))
	}
	return hopts
}

////////////////////////////////////////////////////////////////////////////////

type forceEmpty struct {
	flag bool
}

func (o forceEmpty) ApplyToElemHandler(handler *TypeHandler) {
	handler.forceEmpty = o.flag
}

func ForceEmpty(b bool) Option {
	return forceEmpty{b}
}

////////////////////////////////////////////////////////////////////////////////

type filter struct {
	filter ElementFilter
}

func (o filter) ApplyToElemHandler(handler *TypeHandler) {
	handler.filter = o.filter
}

func WithFilter(fi ElementFilter) Option {
	return filter{fi}
}

////////////////////////////////////////////////////////////////////////////////

type compoption = comphdlr.Option

type compoptionwrapper struct {
	compoption
}

func (o compoptionwrapper) ApplyToElemHandler(handler *TypeHandler) {
}

func WithVersionConstraints(c []*semver.Constraints) Option {
	return compoptionwrapper{comphdlr.WithVersionConstraints(c)}
}

func LatestOnly(b ...bool) Option {
	return compoptionwrapper{comphdlr.LatestOnly(b...)}
}

func Resolver(r ocm.ComponentVersionResolver) Option {
	return compoptionwrapper{comphdlr.Resolver(r)}
}
