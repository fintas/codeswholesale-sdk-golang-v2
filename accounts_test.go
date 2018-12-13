package codeswholesalesdkgolangv2

import (
	"reflect"
	"testing"
)

func TestSDK_GetAccount(t *testing.T) {
	sdk := &SDK{
		clientID:        "ff72ce315d1259e822f47d87d02d261e",                             // Sandbox credentials
		clientSecret:    "$2a$10$E2jVWDADFA5gh6zlRVcrlOOX01Q/HJoT6hXuDMJxek.YEo.lkO2T6", // Sandbox credentials
		clientSignature: "-",
		live:            true, //false,
	}

	tests := []struct {
		name    string
		sdk     *SDK
		want    *Account
		wantErr bool
	}{
		{
			"Success",
			sdk,
			nil,
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.sdk.GetAccount()
			if (err != nil) != tt.wantErr {
				t.Errorf("SDK.GetAccount() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SDK.GetAccount() = %v, want %v", got, tt.want)
			}
		})
	}
}
