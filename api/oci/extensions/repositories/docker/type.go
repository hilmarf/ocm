package docker

import (
	"context"

	"ocm.software/ocm/api/credentials"
	"ocm.software/ocm/api/oci/cpi"
	"ocm.software/ocm/api/utils"
	"ocm.software/ocm/api/utils/runtime"
)

const (
	Type   = "DockerDaemon"
	TypeV1 = Type + runtime.VersionSeparator + "v1"
)

func init() {
	cpi.RegisterRepositoryType(cpi.NewRepositoryType[*RepositorySpec](Type))
	cpi.RegisterRepositoryType(cpi.NewRepositoryType[*RepositorySpec](TypeV1))
}

// RepositorySpec describes an OCI registry interface backed by an oci registry.
type RepositorySpec struct {
	runtime.ObjectVersionedType `json:",inline"`
	DockerHost                  string `json:"dockerHost,omitempty"`
}

// NewRepositorySpec creates a new RepositorySpec for an optional host.
func NewRepositorySpec(host ...string) *RepositorySpec {
	return &RepositorySpec{
		ObjectVersionedType: runtime.NewVersionedTypedObject(Type),
		DockerHost:          utils.Optional(host...),
	}
}

func (a *RepositorySpec) GetType() string {
	return Type
}

func (a *RepositorySpec) Name() string {
	return Type
}

func (a *RepositorySpec) UniformRepositorySpec() *cpi.UniformRepositorySpec {
	return cpi.UniformRepositorySpecForHostURL(Type, a.DockerHost)
}

func (a *RepositorySpec) Repository(ctx cpi.Context, creds credentials.Credentials) (cpi.Repository, error) {
	return NewRepository(ctx, a)
}

func (a *RepositorySpec) Validate(ctx cpi.Context, creds credentials.Credentials, usageContext ...credentials.UsageContext) error {
	client, err := newDockerClient(a.DockerHost)
	if err != nil {
		return err
	}

	_, err = client.Ping(context.Background())
	return err
}
