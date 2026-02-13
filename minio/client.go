package minioadapter

import (
	"context"
	"io"
	"path/filepath"
	"time"

	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"

	"github.com/Skryldev/filestore"
	"github.com/Skryldev/filestore/utils"
)

type MinioStorage struct {
	client *minio.Client
	cfg    *filestore.Config
	logger filestore.Logger
}

func New(cfg *filestore.Config, logger filestore.Logger) (*MinioStorage, error) {
	client, err := minio.New(cfg.Endpoint, &minio.Options{
		Creds:  credentials.NewStaticV4(cfg.AccessKey, cfg.SecretKey, ""),
		Secure: cfg.UseSSL,
	})
	if err != nil {
		return nil, err
	}

	return &MinioStorage{
		client: client,
		cfg:    cfg,
		logger: logger,
	}, nil
}

func (m *MinioStorage) Upload(
	ctx context.Context,
	reader io.Reader,
	size int64,
	originalName string,
	opts filestore.UploadOptions,
) (*filestore.FileInfo, error) {

	if size <= 0 {
		return nil, filestore.ErrInvalidFileSize
	}

	storageKey := utils.GenerateStorageKey()
	publicID := utils.GeneratePublicID()

	objectName := storageKey

	uploadOpts := minio.PutObjectOptions{
		ContentType: opts.ContentType,
	}

	info, err := m.client.PutObject(ctx, m.cfg.Bucket, objectName, reader, size, uploadOpts)
	if err != nil {
		return nil, err
	}

	m.logger.Info(ctx, "file uploaded", map[string]interface{}{
		"public_id": publicID,
		"object":    objectName,
	})

	return &filestore.FileInfo{
		ID:        publicID,
		Name:      filepath.Base(originalName),
		Size:      size,
		MimeType:  opts.ContentType,
		ETag:      info.ETag,
		VersionID: info.VersionID,
		CreatedAt: time.Now().UTC(),
	}, nil
}

