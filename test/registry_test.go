package test

import (
	"fmt"
	"golang.org/x/sys/windows/registry"
	"testing"
)

func TestRegistry(t *testing.T) {
	key, err := registry.OpenKey(registry.CURRENT_USER, "SOFTWARE\\Microsoft\\Windows\\CurrentVersion\\Internet Settings", registry.ALL_ACCESS)

	if err != nil {
		t.Fatal(err)
	}

	fmt.Println(key.GetStringValue("ProxyServer"))
}
