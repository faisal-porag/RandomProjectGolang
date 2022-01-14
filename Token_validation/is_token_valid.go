package Token_validation

import (
	"encoding/base64"
	"encoding/json"
	"strings"
	"time"
)

// JWT TOKEN (IS JWT TOKEN VALID)

func IsTokenValid(token string) bool {
	if token == "" {
		return false
	}

	tokenMiddleSplit := strings.Split(token, ".")
	if len(tokenMiddleSplit) != 3 {
		return false
	}

	tokenMiddlePayload := tokenMiddleSplit[1]
	tokenMiddle, err := base64.RawStdEncoding.DecodeString(tokenMiddlePayload)
	if err != nil {
		return false
	}

	var tokenMiddleMap map[string]int64
	err = json.Unmarshal(tokenMiddle, &tokenMiddleMap)

	exp := tokenMiddleMap["exp"]

	// expire token `TokenExpiryThreshold` seconds earlier than exp to avoid expiry during request
	adjustedExp := exp - int64(time.Duration(4).Seconds())
	expTime := time.Unix(adjustedExp, 0)
	now := time.Now()

	return expTime.After(now)
}
