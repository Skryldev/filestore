package filestore

import (
	"context"
	"io"
	"time"
)

type FileInfo struct {
	ID        string    // ULID (public)
	Name      string    // original filename
	Size      int64
	MimeType  string
	ETag      string
	VersionID string
	CreatedAt time.Time
}

type UploadOptions struct {
	ContentType string
	Metadata    map[string]string
}

type ListOptions struct {
	Prefix string
	Suffix string
}

type Storage interface {
	Upload(ctx context.Context, reader io.Reader, size int64, originalName string, opts UploadOptions) (*FileInfo, error)
	Download(ctx context.Context, id string) (io.ReadCloser, *FileInfo, error)
	Delete(ctx context.Context, id string) error
	PresignedURL(ctx context.Context, id string, expiry time.Duration) (string, error)
	List(ctx context.Context, opts ListOptions) ([]FileInfo, error)
}
