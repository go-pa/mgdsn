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

func (m *DSN) MG() (mailgun.Mailgun, error) {
	if m.Domain == "" || m.APIKey == "" || m.PublicAPIKey == "" {
		return nil, fmt.Errorf("fields are missing: %s", m.String())
	}
	return mailgun.NewMailgun(m.Domain, m.APIKey, m.PublicAPIKey), nil
}

// Flag.Value
func (e *DSN) String() string {
	if e == nil {
		return ""
	}
	var res []string
	if e.Domain != "" {
		res = append(res, "domain="+e.Domain)
	}
	if e.APIKey != "" {
		res = append(res, "api_key="+e.APIKey)
	}
	if e.PublicAPIKey != "" {
		res = append(res, "public_api_key="+e.PublicAPIKey)

	}
	return strings.Join(res, " ")
}

func (e *DSN) Set(value string) error {
	for _, v := range strings.Split(value, " ") {

		kv := strings.SplitN(v, "=", 2)
		if len(kv) < 2 {
			return fmt.Errorf("not a valid k=v expression: %s", v)
		}
		switch kv[0] {
		case "domain":
			e.Domain = kv[1]
		case "api_key":
			e.APIKey = kv[1]
		case "public_api_key":
			e.PublicAPIKey = kv[1]
		default:
			return fmt.Errorf("not a known key: %s", kv[0])
		}

	}
	return nil
}
