package providers

import (
	"context"
	"io"
)

type ProviderType string

var SupportedProviders = map[ProviderType]struct{}{}

func IsSupported(provider ProviderType) bool {
	_, ok := SupportedProviders[provider]
	return ok
}
func GetSupportedProviders() []ProviderType {
	var providers []ProviderType
	for provider := range SupportedProviders {
		providers = append(providers, provider)
	}
	return providers
}

type ObjectStat struct {
	ContentType string
	Size        int64
}

type Provider interface {
	GetObjectStat(ctx context.Context, object string) (*ObjectStat, error)
	GetObject(ctx context.Context, object string) (io.ReadCloser, error)
}
