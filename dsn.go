package mgdsn

import (
	"fmt"
	"strings"

	"gopkg.in/mailgun/mailgun-go.v1"
)

// DSN is a datasource like configuration for mailgun client configuration. It
// implements the flag.Value intervface so it can be used with the flag
// package.
type DSN struct {
	Domain       string
	APIKey       string
	PublicAPIKey string
}

const usageFmt = `"domain=xxx api_key=xxx public_api_key=xxx"`

// Usage can be used with flag.Var(&dsn, "mailgun", mgdsn.Usage)
const Usage = `Mailgun DSN: ` + usageFmt

// Mailgun returns a configured mailgun instance
func (d *DSN) Mailgun() (mailgun.Mailgun, error) {
	if d.Domain == "" || d.APIKey == "" || d.PublicAPIKey == "" {
		return nil, fmt.Errorf("fields are missing from '%s', expected format is '%s'", d.String(), usageFmt)
	}
	return mailgun.NewMailgun(d.Domain, d.APIKey, d.PublicAPIKey), nil
}

// String implements Flag.Value
func (d *DSN) String() string {
	if d == nil {
		return ""
	}
	var res []string
	if d.Domain != "" {
		res = append(res, "domain="+d.Domain)
	}
	if d.APIKey != "" {
		res = append(res, "api_key="+d.APIKey)
	}
	if d.PublicAPIKey != "" {
		res = append(res, "public_api_key="+d.PublicAPIKey)
	}
	return strings.Join(res, " ")
}

// Set implements flag.Value
func (d *DSN) Set(value string) error {
	for _, v := range strings.Split(value, " ") {

		kv := strings.SplitN(v, "=", 2)
		if len(kv) < 2 {
			return fmt.Errorf("not a valid k=v expression: %s", v)
		}
		switch kv[0] {
		case "domain":
			d.Domain = kv[1]
		case "api_key":
			d.APIKey = kv[1]
		case "public_api_key":
			d.PublicAPIKey = kv[1]
		default:
			return fmt.Errorf("not a known key: %s", kv[0])
		}
	}
	return nil
}

// Get implemets flag.Getter
func (d *DSN) Get() interface{} {
	return d
}
