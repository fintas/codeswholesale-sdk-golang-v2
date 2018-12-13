package codeswholesalesdkgolangv2

import (
	"reflect"
	"testing"
)

func TestSDK_GetProductDescription(t *testing.T) {
	sdk := &SDK{
		clientID:        "ff72ce315d1259e822f47d87d02d261e",                             // Sandbox credentials
		clientSecret:    "$2a$10$E2jVWDADFA5gh6zlRVcrlOOX01Q/HJoT6hXuDMJxek.YEo.lkO2T6", // Sandbox credentials
		clientSignature: "-",
		live:            false, //false,
	}
	type args struct {
		productID string
	}
	tests := []struct {
		name    string
		sdk     *SDK
		args    args
		want    *ProductDescription
		wantErr bool
	}{
		{
			"Success",
			sdk,
			args{
				"33e3e81d-2b78-475a-8886-9848116f5133",
			},
			nil,
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.sdk.GetProductDescription(tt.args.productID)
			if (err != nil) != tt.wantErr {
				t.Errorf("SDK.GetProductDescription() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SDK.GetProductDescription() = %v, want %v", got, tt.want)
			}
		})
	}
}
