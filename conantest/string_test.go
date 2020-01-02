package conantest

import (
	"strings"
	"testing"
)

func conanString(str, pre, suf string) (bool, bool) {
	startsWith := strings.HasPrefix(str, pre) // true
	endsWith := strings.HasSuffix(str, suf)   // true
	return startsWith, endsWith
}

func TestConanString(t *testing.T) {
	pre, suf := conanString("conan25216", "a", "216")
	if pre != true {
		t.Error("prefix failed")
	}
	if suf != true {
		t.Error("suf failed")
	}
}
