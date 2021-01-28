package main

import (
	"reflect"
	"testing"
)

func Test_getAppEnvKey(t *testing.T) {
	type args struct {
		osEnv string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "Can get correct app env key",
			args: args{
				osEnv: "APPENV_PORT",
			},
			want: "PORT",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getAppEnvKey(tt.args.osEnv); got != tt.want {
				t.Errorf("getAppEnvKey() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_getAppEnv(t *testing.T) {
	type args struct {
		environs []string
	}
	tests := []struct {
		name        string
		args        args
		wantAppEnvs []appEnv
	}{
		{
			name: "Can get app env",
			args: args{
				environs: []string{
					"OS=xxx",
					"APPENV_PORT=3214",
					"APPENV_HOST=localhost",
				},
			},
			wantAppEnvs: []appEnv{
				{
					key:   "PORT",
					value: "3214",
				},
				{
					key:   "HOST",
					value: "localhost",
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotAppEnvs := getAppEnv(tt.args.environs); !reflect.DeepEqual(gotAppEnvs, tt.wantAppEnvs) {
				t.Errorf("getAppEnv() = %v, want %v", gotAppEnvs, tt.wantAppEnvs)
			}
		})
	}
}
