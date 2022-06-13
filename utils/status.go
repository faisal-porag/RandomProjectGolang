package utils

func IsStatusSuccess(code int) bool {
	return code >= 200 && code < 300
}
