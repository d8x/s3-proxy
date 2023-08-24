package providers

import (
	"context"
	"io"

	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
)

const MinioProvider ProviderType = "minio"

func init() {
	SupportedProviders[MinioProvider] = struct{}{}
}

type Minio struct {
	bucket string
	client *minio.Client
}

func NewMinio(endpoint, bucket, accessKey, secretKey string) (*Minio, error) {
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

	return &Minio{
		bucket: bucket,
		client: client,
	}, nil
}

func (m *Minio) GetObjectStat(ctx context.Context, object string) (*ObjectStat, error) {
	stat, err := m.client.StatObject(ctx, m.bucket, object, minio.StatObjectOptions{})
	if err != nil {
		return nil, err
	}

	return &ObjectStat{
		ContentType: stat.ContentType,
		Size:        stat.Size,
	}, nil
}

func (m *Minio) GetObject(ctx context.Context, object string) (io.ReadCloser, error) {
	return m.client.GetObject(ctx, m.bucket, object, minio.GetObjectOptions{})
}
