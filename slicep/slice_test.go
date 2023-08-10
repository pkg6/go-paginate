package slicep

import (
	"github.com/pkg6/go-paginate"
	"reflect"
	"testing"
)

func TestSliceAdapter(t *testing.T) {
	type args struct {
		source any
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
			if got := SliceAdapter(tt.args.source); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SliceAdapter() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSlice_Length(t *testing.T) {
	type fields struct {
		src any
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
			s := Slice{
				src: tt.fields.src,
			}
			got, err := s.Length()
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

func TestSlice_Slice(t *testing.T) {
	type fields struct {
		src any
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
			s := Slice{
				src: tt.fields.src,
			}
			if err := s.Slice(tt.args.offset, tt.args.length, tt.args.dest); (err != nil) != tt.wantErr {
				t.Errorf("Slice() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_isPtr(t *testing.T) {
	type args struct {
		data any
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isPtr(tt.args.data); got != tt.want {
				t.Errorf("isPtr() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_isSlice(t *testing.T) {
	type args struct {
		data any
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isSlice(tt.args.data); got != tt.want {
				t.Errorf("isSlice() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_makeSlice(t *testing.T) {
	type args struct {
		data   interface{}
		length int
		cap    int
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
			if err := makeSlice(tt.args.data, tt.args.length, tt.args.cap); (err != nil) != tt.wantErr {
				t.Errorf("makeSlice() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
