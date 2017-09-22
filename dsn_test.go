package mgdsn

import "testing"

func TestDSN(t *testing.T) {
	type testCase struct {
		in, out string
	}

	for _, v := range []testCase{
		{"api_key=as=d", "api_key=as=d"},
		{"domain=abc  api_key=asd1 public_api_key=asdf  ", "domain=abc api_key=asd1 public_api_key=asdf"},
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
