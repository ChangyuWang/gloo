// Code generated by solo-kit. DO NOT EDIT.

package v1

import (
	"github.com/solo-io/solo-kit/pkg/utils/hashutils"
	"go.uber.org/zap"
)

type ApiSnapshot struct {
	Proxies   ProxiesByNamespace
	Artifacts ArtifactsByNamespace
	Endpoints EndpointsByNamespace
	Secrets   SecretsByNamespace
	Upstreams UpstreamsByNamespace
}

func (s ApiSnapshot) Clone() ApiSnapshot {
	return ApiSnapshot{
		Proxies:   s.Proxies.Clone(),
		Artifacts: s.Artifacts.Clone(),
		Endpoints: s.Endpoints.Clone(),
		Secrets:   s.Secrets.Clone(),
		Upstreams: s.Upstreams.Clone(),
	}
}

func (s ApiSnapshot) Hash() uint64 {
	return hashutils.HashAll(
		s.hashProxies(),
		s.hashArtifacts(),
		s.hashEndpoints(),
		s.hashSecrets(),
		s.hashUpstreams(),
	)
}

func (s ApiSnapshot) hashProxies() uint64 {
	return hashutils.HashAll(s.Proxies.List().AsInterfaces()...)
}

func (s ApiSnapshot) hashArtifacts() uint64 {
	return hashutils.HashAll(s.Artifacts.List().AsInterfaces()...)
}

func (s ApiSnapshot) hashEndpoints() uint64 {
	return hashutils.HashAll(s.Endpoints.List().AsInterfaces()...)
}

func (s ApiSnapshot) hashSecrets() uint64 {
	return hashutils.HashAll(s.Secrets.List().AsInterfaces()...)
}

func (s ApiSnapshot) hashUpstreams() uint64 {
	return hashutils.HashAll(s.Upstreams.List().AsInterfaces()...)
}

func (s ApiSnapshot) HashFields() []zap.Field {
	var fields []zap.Field
	fields = append(fields, zap.Uint64("proxies", s.hashProxies()))
	fields = append(fields, zap.Uint64("artifacts", s.hashArtifacts()))
	fields = append(fields, zap.Uint64("endpoints", s.hashEndpoints()))
	fields = append(fields, zap.Uint64("secrets", s.hashSecrets()))
	fields = append(fields, zap.Uint64("upstreams", s.hashUpstreams()))

	return append(fields, zap.Uint64("snapshotHash", s.Hash()))
}
