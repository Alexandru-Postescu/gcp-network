package config

import (
	"log"
	"testing"
)

func Test_getJSONPath(t *testing.T) {
	str, err := getJSONPath()
	log.Println(str)
	if err != nil {
		t.Errorf("err:%v", err)
	}
	if str != "C:/Users/APostescu/Downloads/gcpnetwork-349117-21f2f3fa3c28.json" {
		t.Error("nope")
	}
}
