package objectStorage

import (
	"context"
	"github.com/minio/minio-go/v7/pkg/credentials"
	"sync"

	"github.com/minio/minio-go/v7"
)

var instance *minio.Client
var once sync.Once

func MustGetInstance(ctx context.Context, endpoint, password, login string) *minio.Client {
	once.Do(func() {
		instance = mustNewMinio(ctx, endpoint, password, login)
	})

	return instance
}

func mustNewMinio(_ context.Context, endpoint, password, login string) *minio.Client {
	client, err := minio.New(endpoint, &minio.Options{
		Creds:  credentials.NewStaticV4(login, password, ""),
		Secure: false,
	})
	if err != nil {
		panic("MINIO DOESN'T CONNECT")
	}

	if client.IsOffline() {
		panic("MINIO IS OFFLINE")
	}

	return client
}
