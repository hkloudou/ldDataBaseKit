package ldDataBaseKit_test

import (
	"log"
	"testing"
	"time"

	"github.com/hkloudou/ldDataBaseKit"
)

func Test_config(t *testing.T) {
	log.Println(ldDataBaseKit.Err())
}

func Test_time(t *testing.T) {
	log.Println(time.ParseDuration("2s"))
}
