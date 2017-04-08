package helpers

import (
	b64 "encoding/base64"
)

func EncodeUUID(uuid []byte) string {
	uEnc := b64.RawURLEncoding.EncodeToString(uuid)
	return uEnc
}

func DecodeUUID(uuid string) ([]byte, error) {
	uDec, err := b64.RawURLEncoding.DecodeString(uuid)
	return uDec, err
}
