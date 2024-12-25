package utils

import (
	"github.com/google/uuid"
)

func GetUuid() uuid.UUID {
	return uuid.New()
}
