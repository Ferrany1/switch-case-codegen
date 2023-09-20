package extractor //nolint: testpackage // testing unexported functions

import (
	"reflect"
	"testing"
)

func Test_enums_constantObjectTypeNameIfConsistsOfIota(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name string
		e    enums
		want constantObjectTypeName
	}{
		{
			name: "1",
			e:    enums{},
			want: "",
		},
		{
			name: "2",
			e:    enums{constantObjectTypeName("test"): []enum{}},
			want: "",
		},
		{
			name: "3",
			e:    enums{constantObjectTypeName("test"): []enum{{Name: "test1", Value: "iota"}}},
			want: "test",
		},
	}

	for _, testCase := range tests {
		testCase := testCase
		t.Run(
			testCase.name,
			func(t *testing.T) {
				t.Parallel()
				if got := testCase.e.constantObjectTypeNameIfConsistsOfIota(); got != testCase.want {
					t.Errorf("constantObjectTypeNameIfConsistsOfIota() = %v, want %v", got, testCase.want)
				}
			},
		)
	}
}

func TestNewConstantObjectTypeNamesFromStringSlice(t *testing.T) {
	t.Parallel()

	type args struct {
		input []string
	}
	tests := []struct {
		name string
		args args
		want []constantObjectTypeName
	}{
		{
			name: "1",
			args: args{
				input: []string{"test"},
			},
			want: []constantObjectTypeName{"test"},
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(
			tt.name,
			func(t *testing.T) {
				t.Parallel()

				if got := NewConstantObjectTypeNamesFromStringSlice(tt.args.input); !reflect.DeepEqual(got, tt.want) {
					t.Errorf("NewConstantObjectTypeNamesFromStringSlice() = %v, want %v", got, tt.want)
				}
			},
		)
	}
}
