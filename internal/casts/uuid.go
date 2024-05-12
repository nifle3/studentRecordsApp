package casts

import (
	"github.com/google/uuid"
)

func UuidToString(uuId uuid.UUID) (string, error) {
	idBytes, err := uuId.MarshalText()
	if err != nil {
		return "", err
	}

	return string(idBytes), nil
}

func StringToUuid(uuidString string) (uuid.UUID, error) {
	var uuId uuid.UUID
	var err error
	if uuId, err = uuid.FromBytes([]byte(uuidString)); err != nil {
		uuId, err = uuid.Parse(uuidString)
		if err != nil {
			return uuid.Nil, err
		}
	}

	return uuId, err
}
