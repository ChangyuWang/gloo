// Code generated by solo-kit. DO NOT EDIT.

package v1

import (
	"github.com/solo-io/solo-kit/pkg/utils/hashutils"
	"go.uber.org/zap"
)

type DiscoverySnapshot struct {
	Secrets   SecretsByNamespace
	Upstreams UpstreamsByNamespace
}

func (s DiscoverySnapshot) Clone() DiscoverySnapshot {
	return DiscoverySnapshot{
		Secrets:   s.Secrets.Clone(),
		Upstreams: s.Upstreams.Clone(),
	}
}

func (s DiscoverySnapshot) Hash() uint64 {
	return hashutils.HashAll(
		s.hashSecrets(),
		s.hashUpstreams(),
	)
}

func (s DiscoverySnapshot) hashSecrets() uint64 {
	return hashutils.HashAll(s.Secrets.List().AsInterfaces()...)
}

func (s DiscoverySnapshot) hashUpstreams() uint64 {
	return hashutils.HashAll(s.Upstreams.List().AsInterfaces()...)
}

func (s DiscoverySnapshot) HashFields() []zap.Field {
	var fields []zap.Field
	fields = append(fields, zap.Uint64("secrets", s.hashSecrets()))
	fields = append(fields, zap.Uint64("upstreams", s.hashUpstreams()))

	return append(fields, zap.Uint64("snapshotHash", s.Hash()))
}
