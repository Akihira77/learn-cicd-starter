package auth

import (
	"net/http"
	"strings"
	"testing"
)

func TestGetAPIKey(t *testing.T) {

	tests := map[string]struct {
		key       string
		value     string
		expect    string
		expectErr string
	}{
		"valid":             {key: "Authorization", value: "ApiKey secret-token", expect: "secret-token", expectErr: ""},
		"wrong value":       {key: "Authorization", value: "", expectErr: "no authorization header included"},
		"wrong key":         {key: "Authorization", value: "secret-token", expectErr: "malformed authorization header"},
		"invalid key&value": {key: "", value: "", expectErr: "no authorization header included"},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			header := http.Header{}
			header.Add(tc.key, tc.value)

			got, err := GetAPIKey(header)
			if err == nil {
				if tc.expect != got {
					t.Errorf("Unexpected: TestGetAPIKey:%v\n", err)
					return
				}
			} else {
				msg := err.Error()
				if strings.Contains(msg, tc.expectErr) {
					return
				}

				t.Errorf("Unexpected: TestGetAPIKey:%v\n", err)
			}
		})
	}
}
