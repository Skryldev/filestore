package filestore

import "errors"

// Domain-level errors (framework agnostic)

var (
	// Validation errors
	ErrInvalidFileSize  = errors.New("invalid file size")
	ErrInvalidFileName  = errors.New("invalid file name")
	ErrInvalidMimeType  = errors.New("invalid mime type")
	ErrFileTooLarge     = errors.New("file exceeds allowed size")

	// Storage errors
	ErrFileNotFound     = errors.New("file not found")
	ErrFileAlreadyExists = errors.New("file already exists")
	ErrStorageUnavailable = errors.New("storage service unavailable")

	// Security errors
	ErrUnauthorizedAccess = errors.New("unauthorized file access")

	// Versioning
	ErrVersionNotFound = errors.New("file version not found")
)
