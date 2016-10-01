package env

import (
	"os"
	"testing"
)

func TestGetString(t *testing.T) {
	name := "string"
	want := "default"
	s := GetString(name, want)
	if s != want {
		t.Errorf("got: %v, want: %v", s, want)
	}

	want = "hoge"
	if err := os.Setenv(name, want); err != nil {
		t.Errorf("set env error: %v", err)
	}

	s = GetString(name, "default")
	if s != want {
		t.Errorf("got: %v, want: %v", s, want)
	}
}

func TestGetInt(t *testing.T) {
	name := "int"
	want := 0
	i := GetInt(name, want)
	if i != want {
		t.Errorf("got: %v, want: %v", i, want)
	}

	if err := os.Setenv(name, "1"); err != nil {
		t.Errorf("set env error: %v", err)
	}

	want = 1
	i = GetInt(name, 0)
	if i != want {
		t.Errorf("got: %v, want: %v", i, want)
	}

	if err := os.Setenv(name, "hoge"); err != nil {
		t.Errorf("set env error: %v", err)
	}

	i = GetInt(name, 0)
	if i != 0 {
		t.Errorf("got: %v, want: %v", i, 0)
	}
}

func TestGetInt64(t *testing.T) {
	name := "int64"
	want := int64(0)
	i := GetInt64(name, want)
	if i != want {
		t.Errorf("got: %v, want: %v", i, want)
	}

	if err := os.Setenv(name, "1"); err != nil {
		t.Errorf("set env error: %v", err)
	}

	want = int64(1)
	i = GetInt64(name, 0)
	if i != want {
		t.Errorf("got: %v, want: %v", i, want)
	}

	if err := os.Setenv(name, "hoge"); err != nil {
		t.Errorf("set env error: %v", err)
	}

	i = GetInt64(name, 0)
	if i != 0 {
		t.Errorf("got: %v, want: %v", i, 0)
	}
}

func TestGetUint64(t *testing.T) {
	name := "uint64"
	want := uint64(0)
	i := GetUint64(name, want)
	if i != want {
		t.Errorf("got: %v, want: %v", i, want)
	}

	if err := os.Setenv(name, "1"); err != nil {
		t.Errorf("set env error: %v", err)
	}

	want = uint64(1)
	i = GetUint64(name, 0)
	if i != want {
		t.Errorf("got: %v, want: %v", i, want)
	}

	if err := os.Setenv(name, "hoge"); err != nil {
		t.Errorf("set env error: %v", err)
	}

	i = GetUint64(name, 0)
	if i != 0 {
		t.Errorf("got: %v, want: %v", i, 0)
	}
}

func TestGetFloat(t *testing.T) {
	name := "float64"
	want := float64(0)
	f := GetFloat(name, want)
	if f != want {
		t.Errorf("got: %v, want: %v", f, want)
	}

	if err := os.Setenv(name, "10.1"); err != nil {
		t.Errorf("set env error: %v", err)
	}

	want = float64(10.1)
	f = GetFloat(name, 0)
	if f != want {
		t.Errorf("got: %v, want: %v", f, want)
	}

	if err := os.Setenv(name, "hoge"); err != nil {
		t.Errorf("set env error: %v", err)
	}

	f = GetFloat(name, 0)
	if f != 0 {
		t.Errorf("got: %v, want: %v", f, 0)
	}
}

func TestGetBool(t *testing.T) {
	name := "bool"
	want := false
	b := GetBool(name, want)
	if b != want {
		t.Errorf("got: %v, want: %v", b, want)
	}

	if err := os.Setenv(name, "true"); err != nil {
		t.Errorf("set env error: %v", err)
	}

	want = true
	b = GetBool(name, false)
	if b != want {
		t.Errorf("got: %v, want: %v", b, want)
	}

	if err := os.Setenv(name, "hoge"); err != nil {
		t.Errorf("set env error: %v", err)
	}

	b = GetBool(name, false)
	if b != false {
		t.Errorf("got: %v, want: %v", b, false)
	}
}
