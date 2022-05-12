package utils

import (
	"github.com/oklog/ulid"
	"math/rand"
	"time"
)

func GenerateUlid() string {
	t := time.Now()
	entropy := rand.New(rand.NewSource(t.UnixNano()))
	id := ulid.MustNew(ulid.Timestamp(t), entropy)
	return id.String()
}
