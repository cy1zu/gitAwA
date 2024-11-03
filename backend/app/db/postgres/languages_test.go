package postgres

import (
	"reflect"
	"testing"
)

func TestGetLanguages(t *testing.T) {
	type args struct {
		recType string
		id      int64
	}
	tests := []struct {
		name    string
		args    args
		want    map[string]int64
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := GetLanguages(tt.args.recType, tt.args.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetLanguages() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetLanguages() got = %v, want %v", got, tt.want)
			}
		})
	}
}
