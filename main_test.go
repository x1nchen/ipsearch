package main

import (
	"net"
	"reflect"
	"testing"
)

func Test_formatResData(t *testing.T) {
	type args struct {
		src string
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{
			"test formatResData",
			args{
				`["中国","天津","天津","","鹏博士"]`,
			},
			[]string{"中国", "天津", "天津", "鹏博士"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := formatResData(tt.args.src); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("formatResData() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_lookupIP(t *testing.T) {
	// mock net.LookupIP
	mockNetLookupIP := LookupIP
	defer func() { LookupIP = mockNetLookupIP }()

	LookupIP = func(domain string) ([]net.IP, error) {
		return []net.IP{[]byte{8, 8, 8, 8}}, nil
	}

	type args struct {
		domain string
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{name: "test correct lookup", args: args{"test.host"}, want: []string{"8.8.8.8"}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := lookupIP(tt.args.domain); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("lookupIP() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_parseInputParam(t *testing.T) {
	type args struct {
		input []string
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{name: "test param lack", args: args{[]string{}}, want: "", wantErr: true},
		{
			"test param exceed",
			args{
				[]string{"1", "2"},
			},
			"",
			true,
		},
		{
			"test one param",
			args{
				[]string{"1"},
			},
			"1",
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := parseInputParam(tt.args.input)
			if (err != nil) != tt.wantErr {
				t.Errorf("parseInputParam() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("parseInputParam() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_parseInputToTargetIPS(t *testing.T) {
	type args struct {
		input string
	}
	tests := []struct {
		name    string
		args    args
		want    []string
		wantErr bool
	}{
		// TODO
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := parseInputToTargetIPS(tt.args.input)
			if (err != nil) != tt.wantErr {
				t.Errorf("parseInputToTargetIPS() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("parseInputToTargetIPS() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_isIP(t *testing.T) {
	type args struct {
		ip string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{name: "test correct IP 0", args: args{"255.255.255.255"}, want: true},
		{name: "test correct IP 1", args: args{"0.0.0.0"}, want: true},
		{name: "test incorrect IP 0", args: args{"0.0"}, want: false},
		{name: "test incorrect IP 1", args: args{"#"}, want: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isIP(tt.args.ip); got != tt.want {
				t.Errorf("isIP() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_isURL(t *testing.T) {
	type args struct {
		raw string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{name: "test correct domain in general", args: args{"google.com"}, want: true},
		{name: "test correct domain with sub-domain", args: args{"api.google.com"}, want: true},
		{name: "test correct domain with protocal", args: args{"https://golang.org/pkg/net/url/#URL"}, want: true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isURL(tt.args.raw); got != tt.want {
				t.Errorf("isDomain() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_requestAPI(t *testing.T) {
	type args struct {
		ip string
	}
	tests := []struct {
		name string
		args args
	}{
		{name: "118.28.8.8", args: args{ip: "118.28.8.8"}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			requestAPI(tt.args.ip)
		})
	}
}
