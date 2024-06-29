package utils

import (
	"context"
	"edusync/internal/config"
	"edusync/pkg/logger"

	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
)

func SetupMinioConnection(cfg *config.Config, log *logger.Logger) (*minio.Client, error) {
	ctx := context.Background()

	// Initialize minio client object.
	minioClient, errInit := minio.New(cfg.MinioEndpoint, &minio.Options{
		Creds:  credentials.NewStaticV4(cfg.MinioAccessKeyId, cfg.MinioSecretKey, ""),
		Secure: cfg.MinioUseSSL,
	})
	if errInit != nil {
		log.Fatal(errInit)
	}

	err := minioClient.MakeBucket(ctx, cfg.MinioBucketName, minio.MakeBucketOptions{Region: cfg.MinioLocation})
	if err != nil {
		exists, errBucketExists := minioClient.BucketExists(ctx, cfg.MinioBucketName)
		if errBucketExists != nil && !exists {
			log.Fatal(err)
		}
	}

	return minioClient, errInit
}
