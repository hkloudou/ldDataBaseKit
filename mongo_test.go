package lddatabasekit_test

/*
 * @Author: hkloudou
 * @Github: https://github.com/hkloudou/
 * @LastEditors: 卢教(aven) hkloudou@gmail.com
 * @Date: 2018-07-07 02:45:59
 * @LastEditTime: 2019-03-02 19:49:16
 */

import (
	"log"
	"testing"
	"time"

	lddatabasekit "github.com/hkloudou/lddatabasekit"
)

func Test_config(t *testing.T) {
	log.Println(lddatabasekit.Err())
}

func Test_time(t *testing.T) {
	log.Println(time.ParseDuration("2s"))
}
