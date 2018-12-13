package codeswholesalesdkgolangv2

import (
	"reflect"
	"testing"
)

func TestSDK_GetCode(t *testing.T) {
	sdk := &SDK{
		clientID:        "1cd7ede60da5f9359ec368343ab49e11",                             //"ff72ce315d1259e822f47d87d02d261e",                             // Sandbox credentials
		clientSecret:    "$2a$10$T58Tbz/hgPrghjUDbRIKV.KRspoV6/6QcTV158mITvFu6JBYvlhSq", //"$2a$10$E2jVWDADFA5gh6zlRVcrlOOX01Q/HJoT6hXuDMJxek.YEo.lkO2T6", // Sandbox credentials
		clientSignature: "-",
		live:            true, //false,
	}

	type args struct {
		codeID string
	}
	tests := []struct {
		name    string
		sdk     *SDK
		args    args
		want    *Code
		wantErr bool
	}{
		{
			"Success",
			sdk,
			args{"895890232-2513"},
			nil,
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.sdk.GetCode(tt.args.codeID)
			if (err != nil) != tt.wantErr {
				t.Errorf("SDK.GetCode() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SDK.GetCode() = %v, want %v", got, tt.want)
			}
		})
	}
}
