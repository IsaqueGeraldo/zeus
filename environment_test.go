package zeus

import (
	"reflect"
	"testing"
)

func TestGetenv(t *testing.T) {
	type args struct {
		key string
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Getenv(tt.args.key)
			if (err != nil) != tt.wantErr {
				t.Errorf("Getenv() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Getenv() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSetenv(t *testing.T) {
	type args struct {
		key   string
		value string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := Setenv(tt.args.key, tt.args.value); (err != nil) != tt.wantErr {
				t.Errorf("Setenv() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestEnviron(t *testing.T) {
	tests := []struct {
		name    string
		want    []Environment
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Environ()
			if (err != nil) != tt.wantErr {
				t.Errorf("Environ() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Environ() = %v, want %v", got, tt.want)
			}
		})
	}
}
