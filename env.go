package lddatabasekit

import (
	"os"
)

func getEnv(key, def string) string {
	if tmp := os.Getenv(key); tmp != "" {
		return tmp
	}
	return def
}
