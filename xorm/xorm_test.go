package xorm

import (
	"github.com/pkg6/go-paginate"
	"reflect"
	"testing"
	"xorm.io/xorm"
)

func TestAdapter(t *testing.T) {
	type args struct {
		session *xorm.Session
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
			if got := Adapter(tt.args.session); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Adapter() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestXORM_Length(t *testing.T) {
	type fields struct {
		session *xorm.Session
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
			x := XORM{
				session: tt.fields.session,
			}
			got, err := x.Length()
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

func TestXORM_Slice(t *testing.T) {
	type fields struct {
		session *xorm.Session
	}
	type args struct {
		offset int64
		length int64
		data   any
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
			x := XORM{
				session: tt.fields.session,
			}
			if err := x.Slice(tt.args.offset, tt.args.length, tt.args.data); (err != nil) != tt.wantErr {
				t.Errorf("Slice() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
