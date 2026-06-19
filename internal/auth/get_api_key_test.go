package auth

import (
	"net/http"
	"testing"
)

func TestGetAPIKey(t *testing.T) {
	type test struct {
		headerKey string
		headerValue string
		want string
	}

	tests := []test{
		{headerKey: "Authorization", headerValue: "ApiKey Test", want: "Test"},
		{headerKey: "aaaaaaaa", headerValue: "aaaaaa", want: ""},
		{headerKey: "Authorization", headerValue: "aKey", want: ""},
	}

	for _,tc := range tests {
		header := http.Header{}
		header.Add(tc.headerKey, tc.headerValue)
		got, _ := GetAPIKey(header)

		if got != tc.want {
			t.Error("Errored with" + got)
		}

	}
}