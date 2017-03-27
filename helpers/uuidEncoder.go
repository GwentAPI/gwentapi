package helpers

import (
	b64 "encoding/base64"
)

func UUIDToURLBase64(uuid string) string {
	uEnc := b64.URLEncoding.EncodeToString([]byte(uuid))
	return uEnc
}

func Base64ToUUID(uuid string) (string, error) {
	uDec, err := b64.URLEncoding.DecodeString(uuid)
	return string(uDec), err
}
