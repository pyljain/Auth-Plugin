package pkce

import (
	"crypto/sha256"
	"encoding/base64"
)

func genSha256(codeVerifier string) []byte {

	hashFn := sha256.New()

	hashFn.Write([]byte(codeVerifier))
	b := hashFn.Sum(nil)

	return b
}

func GenerateCodeChallenge(codeVerifier string) string {
	sha := genSha256(codeVerifier)
	return base64.RawURLEncoding.EncodeToString(sha)
}
