// Copyright 2022 SAP SE or an SAP affiliate company. All rights reserved. This file is licensed under the Apache Software License, v. 2 except as noted otherwise in the LICENSE file
//
//  Licensed under the Apache License, Version 2.0 (the "License");
//  you may not use this file except in compliance with the License.
//  You may obtain a copy of the License at
//
//       http://www.apache.org/licenses/LICENSE-2.0
//
//  Unless required by applicable law or agreed to in writing, software
//  distributed under the License is distributed on an "AS IS" BASIS,
//  WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//  See the License for the specific language governing permissions and
//  limitations under the License.

package get_test

import (
	"bytes"

	. "github.com/gardener/ocm/cmds/ocm/testhelper"
	"github.com/gardener/ocm/cmds/ocm/testhelper/builder"
	"github.com/gardener/ocm/pkg/common/accessio"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

const CA = "/tmp/ca"
const CTF = "/tmp/ctf"
const VERSION = "v1"
const COMP = "test.de/x"
const COMP2 = "test.de/y"
const COMP3 = "test.de/z"
const PROVIDER = "mandelsoft"

var _ = Describe("Test Environment", func() {
	var env *builder.Builder

	BeforeEach(func() {
		env = builder.NewBuilder(NewTestEnv())

	})

	AfterEach(func() {
		env.Cleanup()
	})

	It("lists single reference in component archive", func() {
		env.ComponentArchive(CA, accessio.FormatDirectory, COMP, VERSION, func() {
			env.Provider(PROVIDER)
			env.Reference("test", COMP2, VERSION)
			env.Reference("withid", COMP3, VERSION, func() {
				env.ExtraIdentity("id", "test")
			})
		})

		buf := bytes.NewBuffer(nil)
		Expect(env.CatchOutput(buf).Execute("get", "references", "-o", "wide", CA)).To(Succeed())
		Expect("\n" + buf.String()).To(Equal(
			`
NAME   COMPONENT VERSION   IDENTITY
test   v1        test.de/y "name"="test"
withid v1        test.de/z "id"="test","name"="withid"
`))
	})

	It("lists single reference in component archive", func() {
		env.OCMCommonTransport(CTF, accessio.FormatDirectory, func() {
			env.ComponentVersion(COMP2, VERSION, func() {
				env.Provider(PROVIDER)
				env.Reference("withid", COMP3, VERSION, func() {
					env.ExtraIdentity("id", "test")
				})
			})
			env.ComponentVersion(COMP3, VERSION, func() {
				env.Provider(PROVIDER)
			})
		})
		env.ComponentArchive(CA, accessio.FormatDirectory, COMP, VERSION, func() {
			env.Provider(PROVIDER)
			env.Reference("test", COMP2, VERSION)
		})

		buf := bytes.NewBuffer(nil)
		Expect(env.CatchOutput(buf).Execute("get", "references", "--lookup", CTF, "-c", "-o", "wide", CA)).To(Succeed())
		Expect("\n" + buf.String()).To(Equal(
			`
REFERENCEPATH              NAME   COMPONENT VERSION   IDENTITY
test.de/x:v1               test   v1        test.de/y "name"="test"
test.de/x:v1->test.de/y:v1 withid v1        test.de/z "id"="test","name"="withid"
`))
	})

})