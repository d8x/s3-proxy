package providers

import (
	"context"
	"io"

	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
)

const ScalewayProvider ProviderType = "scaleway"

func init() {
	SupportedProviders[ScalewayProvider] = struct{}{}
}

type Scaleway struct {
	client *minio.Client
	bucket string
}

func NewScaleway(endpoint, bucket, accessKey, secretKey string) (*Scaleway, error) {
	options := &minio.Options{
		Creds: credentials.NewStaticV4(accessKey, secretKey, ""),
	}
	if accessKey == "" || secretKey == "" {
		options.Creds = nil
	}

	client, err := minio.New(endpoint, options)
	if err != nil {
		return nil, err
	}

	return &Scaleway{
		client: client,
		bucket: bucket,
	}, nil
}

func (s *Scaleway) GetObjectStat(ctx context.Context, object string) (*ObjectStat, error) {
	stat, err := s.client.StatObject(ctx, s.bucket, object, minio.StatObjectOptions{})
	if err != nil {
		return nil, err
	}

	return &ObjectStat{
		ContentType: stat.ContentType,
		Size:        stat.Size,
	}, nil
}

func (s *Scaleway) GetObject(ctx context.Context, object string) (io.ReadCloser, error) {
	return s.client.GetObject(ctx, s.bucket, object, minio.GetObjectOptions{})
}
