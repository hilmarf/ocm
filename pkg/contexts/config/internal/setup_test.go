package internal_test

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"github.com/open-component-model/ocm/pkg/contexts/config"
	"github.com/open-component-model/ocm/pkg/contexts/config/internal"
)

var _ = Describe("setup", func() {
	It("creates initial", func() {
		Expect(len(config.DefaultContext().ConfigTypes().KnownTypeNames())).To(Equal(6))
		Expect(len(internal.DefaultConfigTypeScheme.KnownTypeNames())).To(Equal(6))
	})
})
