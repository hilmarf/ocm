package app

import (
	"github.com/spf13/cobra"

	"github.com/open-component-model/ocm/cmds/helminstaller/app/driver"
	"github.com/open-component-model/ocm/cmds/helminstaller/app/driver/helm"
	"github.com/open-component-model/ocm/pkg/contexts/clictx"
	"github.com/open-component-model/ocm/pkg/toi/support"
)

func NewCliCommand(ctx clictx.Context, d driver.Driver) *cobra.Command {
	if d == nil {
		d = helm.New()
	}
	return support.NewCLICommand(ctx.OCMContext(), "helmbootstrapper", New(d))
}
