package adapter

import (
	"github.com/pkg6/go-paginate"
	"gorm.io/gorm"
	"reflect"
	"testing"
)

func TestGORMAdapter(t *testing.T) {
	type args struct {
		db *gorm.DB
	}
	tests := []struct {
		name string
		args args
		want paginate.IAdapter
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GORMAdapter(tt.args.db); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GORMAdapter() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGORM_Length(t *testing.T) {
	type fields struct {
		db *gorm.DB
	}
	tests := []struct {
		name    string
		fields  fields
		want    int64
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a := GORM{
				db: tt.fields.db,
			}
			got, err := a.Length()
			if (err != nil) != tt.wantErr {
				t.Errorf("Length() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Length() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGORM_Slice(t *testing.T) {
	type fields struct {
		db *gorm.DB
	}
	type args struct {
		offset int64
		length int64
		dest   any
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a := GORM{
				db: tt.fields.db,
			}
			if err := a.Slice(tt.args.offset, tt.args.length, tt.args.dest); (err != nil) != tt.wantErr {
				t.Errorf("Slice() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
