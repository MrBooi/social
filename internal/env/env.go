package env

import (
	"os"
)

func GetString(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}

//func getInt(key string, fallback int) int {
//	valStr, ok := os.LookupEnv(key)
//	if !ok {
//		return fallback
//	}
//
//	val, err := strconv.Atoi(valStr)
//	if err != nil {
//		return fallback
//	}
//
//	return val
//}
