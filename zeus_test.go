package zeus

import (
	"reflect"
	"testing"

	"gorm.io/gorm"
)

func TestBootstrap(t *testing.T) {
	type args struct {
		database string
	}
	tests := []struct {
		name    string
		args    args
		want    *gorm.DB
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Bootstrap(tt.args.database)
			if (err != nil) != tt.wantErr {
				t.Errorf("Bootstrap() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Bootstrap() = %v, want %v", got, tt.want)
			}
		})
	}
}
