package helpers

import (
	b64 "encoding/base64"
)

func UUIDToURLBase64(uuid []byte) string {
	uEnc := b64.URLEncoding.EncodeToString(uuid)
	return uEnc
}

func Base64ToUUID(uuid string) ([]byte, error) {
	uDec, err := b64.URLEncoding.DecodeString(uuid)
	return uDec, err
}
