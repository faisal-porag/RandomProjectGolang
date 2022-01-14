package HashAlgorithm

import (
	"crypto/hmac"
	"crypto/rand"
	"crypto/sha256"
	"crypto/sha512"
	"encoding/base64"
	"fmt"
	"io"
)

var secretKey = "4234kxzjcjj3nxnxbcvsjfj"

// Generate a salt string with 16 bytes of crypto/rand data.
func generateSalt() string {
	randomBytes := make([]byte, 16)
	_, err := rand.Read(randomBytes)
	if err != nil {
		return ""
	}
	return base64.URLEncoding.EncodeToString(randomBytes)
}

func GenerateHashAlgorithmV1() {
	message := "Today web engineering has modern apps adhere to what is known as a single-page app (SPA) model."
	salt := generateSalt()
	fmt.Println("Message: " + message)
	fmt.Println("\nSalt: " + salt)

	hash := hmac.New(sha256.New, []byte(secretKey))
	_, err := io.WriteString(hash, message+salt)
	if err != nil {
		return
	}
	fmt.Printf("\nHMAC-Sha256: %x", hash.Sum(nil))

	hash = hmac.New(sha512.New, []byte(secretKey))
	_, err1 := io.WriteString(hash, message+salt)
	if err1 != nil {
		return
	}
	fmt.Printf("\n\nHMAC-sha512: %x", hash.Sum(nil))
}

/* OUTPUT:
Message: Today web engineering has modern apps adhere to what is known as a sing
le-page app (SPA) model.

Salt: iWk9q-tQgWQTnqDgdoxaXQ==

HMAC-Sha256: b158c5a1bbcdac3cf87fe761030828cb5811b0a6fdfa6366c7bdfddba6391728

HMAC-sha512: e350ca7f0349c2b16a410f224b1ad0c8fc9319708b1dd2be9e83a53b3d4b93d9dd1
f0637ea27641edcfac3d3196795d9889778bd4894ad332ba643d0735aa089
*/
