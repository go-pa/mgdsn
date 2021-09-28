package mgdsn

import (
	"flag"
	"testing"
)

func TestDSN(t *testing.T) {
	type testCase struct {
		in, out string
	}
	for _, v := range []testCase{
		{"api_key=as=d", "api_key=as=d"},
		{"domain=abc  api_key=asd1 public_api_key=asdf  ", "domain=abc api_key=asd1 public_api_key=asdf"},
		{"domain=abc api_koo=1", ""},
		{"domain=abc  api_key=asd1 public_api_key=asdf  api_koo=1", ""},
		{"domain= api_key=1 public_api_key=1", "api_key=1 public_api_key=1"},
	} {
		var d DSN
		err := d.Set(v.in)
		if v.out == "" && err == nil {
			t.Fatal("should have failed", v.in, d.String())
		}
	}
}

func TestNil(t *testing.T) {
	var d *DSN
	if d.String() != "" {
		t.Fail()
	}
}

func TestFlag(t *testing.T) {
	fs := flag.NewFlagSet("test", flag.ContinueOnError)
	var dsn DSN
	fs.Var(&dsn, "testdsn", "testtest")
	err := fs.Parse([]string{"-testdsn", "domain=1 api_key=2 public_api_key=3"})
	if err != nil {
		t.Fatal(err)
	}
	if dsn.Domain != "1" {
		t.Fatal("domain", dsn.Domain)
	}
	if dsn.APIKey != "2" {
		t.Fatal("api key", dsn.APIKey)
	}
	if dsn.PublicAPIKey != "3" {
		t.Fatal("pub pi key", dsn.PublicAPIKey)
	}
	fs.VisitAll(func(f *flag.Flag) {
		v := f.Value.(flag.Getter).Get()
		if v.(*DSN).Domain != "1" {
			t.Fatal(v)
		}
	})
}
