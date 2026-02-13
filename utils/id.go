package utils

import (
	"github.com/google/uuid"
	"github.com/oklog/ulid/v2"
	"time"
	mathrand "math/rand"
)

func GenerateStorageKey() string {
	return uuid.New().String() // internal storage name
}

func GeneratePublicID() string {
	t := time.Now().UTC()
	entropy := ulid.Monotonic(mathrand.New(mathrand.NewSource(t.UnixNano())), 0)
	return ulid.MustNew(ulid.Timestamp(t), entropy).String()
}
