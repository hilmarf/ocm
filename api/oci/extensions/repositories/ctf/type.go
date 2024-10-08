package ctf

import (
	"strings"

	"ocm.software/ocm/api/credentials"
	"ocm.software/ocm/api/datacontext/attrs/vfsattr"
	"ocm.software/ocm/api/oci/cpi"
	"ocm.software/ocm/api/utils/accessio"
	"ocm.software/ocm/api/utils/accessobj"
	"ocm.software/ocm/api/utils/runtime"
)

const (
	Type   = cpi.CommonTransportFormat
	TypeV1 = Type + runtime.VersionSeparator + "v1"
)

func init() {
	cpi.RegisterRepositoryType(cpi.NewRepositoryType[*RepositorySpec](Type))
	cpi.RegisterRepositoryType(cpi.NewRepositoryType[*RepositorySpec](TypeV1))
}

// RepositorySpec describes an OCI registry interface backed by an oci registry.
type RepositorySpec struct {
	runtime.ObjectVersionedType `json:",inline"`
	accessio.StandardOptions    `json:",inline"`

	// FilePath is the file for the repository in the filesystem..
	FilePath string `json:"filePath"`
	// AccessMode can be set to request readonly access or creation
	AccessMode accessobj.AccessMode `json:"accessMode,omitempty"`
}

var _ cpi.RepositorySpec = (*RepositorySpec)(nil)

var _ cpi.IntermediateRepositorySpecAspect = (*RepositorySpec)(nil)

// NewRepositorySpec creates a new RepositorySpec.
func NewRepositorySpec(mode accessobj.AccessMode, filePath string, opts ...accessio.Option) (*RepositorySpec, error) {
	o, err := accessio.AccessOptions(nil, opts...)
	if err != nil {
		return nil, err
	}
	if o.GetFileFormat() == nil {
		for _, v := range SupportedFormats() {
			if strings.HasSuffix(filePath, "."+v.String()) {
				o.SetFileFormat(v)
				break
			}
		}
	}
	o.Default()
	return &RepositorySpec{
		ObjectVersionedType: runtime.NewVersionedTypedObject(Type),
		FilePath:            filePath,
		StandardOptions:     *o.(*accessio.StandardOptions),
		AccessMode:          mode,
	}, nil
}

func (a *RepositorySpec) IsIntermediate() bool {
	return true
}

func (a *RepositorySpec) GetType() string {
	return Type
}

func (s *RepositorySpec) Name() string {
	return s.FilePath
}

func (s *RepositorySpec) UniformRepositorySpec() *cpi.UniformRepositorySpec {
	u := &cpi.UniformRepositorySpec{
		Type: Type,
		Info: s.FilePath,
	}
	return u
}

func (a *RepositorySpec) Repository(ctx cpi.Context, creds credentials.Credentials) (cpi.Repository, error) {
	opts := a.StandardOptions
	opts.Default(vfsattr.Get(ctx))

	return Open(ctx, a.AccessMode, a.FilePath, 0o700, &opts)
}

func (a *RepositorySpec) Validate(ctx cpi.Context, creds credentials.Credentials, context ...credentials.UsageContext) error {
	opts := a.StandardOptions
	opts.Default(vfsattr.Get(ctx))

	return accessobj.ValidateDescriptor(accessObjectInfo, a.FilePath, opts.GetPathFileSystem())
}
