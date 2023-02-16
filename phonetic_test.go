package phonetic

import (
	"testing"
)

func TestConvertBytes(t *testing.T) {
	want := "tree fower fife"
	got := string(NATO.ConvertBytes([]byte("345"), false))
	if want != got {
		t.Fatalf("want:%q got:%q", want, got)
	}

	got = string(NATO.ConvertBytes([]byte("345"), true))
	want = "tree foewhur fife"
	if want != got {
		t.Fatalf("want:%q got:%q", want, got)
	}
}

func TestConvert(t *testing.T) {
	want := "xray tango"
	got := NATO.Convert("xt", false)
	if want != got {
		t.Fatalf("want:%q got:%q", want, got)
	}
}
