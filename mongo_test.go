package lddatabasekit_test

import (
	"log"
	"testing"
	"time"

	lddatabasekit "github.com/hkloudou/ldDataBaseKit"
)

func Test_config(t *testing.T) {
	log.Println(lddatabasekit.Err())
}

func Test_time(t *testing.T) {
	log.Println(time.ParseDuration("2s"))
}
