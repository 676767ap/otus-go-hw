package util

import (
	"reflect"
	"testing"
)

type testStruct struct {
	ID   int32
	Name string
}

type differentStruct struct {
	Doc float32
}

func TestMappingStructure(t *testing.T) {
	structure := &testStruct{
		ID:   1,
		Name: "name",
	}
	type args struct {
		input  interface{}
		output interface{}
	}
	tests := []struct {
		name    string
		args    args
		want    interface{}
		wantErr bool
	}{
		{
			name: "Test mapping",
			args: args{
				input:  structure,
				output: &testStruct{},
			},
			want:    structure,
			wantErr: false,
		},
		{
			name: "Test mapping error",
			args: args{
				input:  structure,
				output: &differentStruct{},
			},
			want:    &differentStruct{},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := MappingStructure(tt.args.input, tt.args.output); (err != nil) != tt.wantErr {
				t.Errorf("MappingStructure() error = %v, wantErr %v", err, tt.wantErr)
			}
			if !reflect.DeepEqual(tt.args.output, tt.want) {
				t.Errorf("TestMappingStructure = %v, want %v", tt.args.output, tt.want)
			}
		})
	}
}
