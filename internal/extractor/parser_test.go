package extractor //nolint: testpackage // testing unexported functions

import (
	"reflect"
	"testing"
)

type TestType1 string

const (
	TestType1Value1 TestType1 = "1"
	TestType1Value2 TestType1 = "2"
	TestValue1      string    = "3"
	TestValue2      int       = 4
)

type TestType2 int

const (
	TestType2Value1 TestType2 = iota
	TestType2Value2
	TestType2Value3
	TestType2Value4
	testType2Value5
)

type TestType3 float64

const (
	TestType3Value1 TestType3 = 1.0
)

type TestType4 bool

const TestType4Value1 TestType4 = true

const (
	TestValue3 float64 = 2.0
)

const TestValue4 bool = true

const (
	TestValue5 float64 = 3.0
	TestValue6 float64 = 4.0
)

//nolint:funlen // huge test cases
func Test_parseEnumsFromFile(t *testing.T) {
	t.Parallel()

	type args struct {
		filePath              string
		targetObjectTypeNames []constantObjectTypeName
		includeNonExported    bool
	}
	tests := []struct {
		name    string
		args    args
		want    enums
		wantErr bool
	}{
		{
			name: "1",
			args: args{
				filePath:              "./parser_test.go",
				targetObjectTypeNames: nil,
				includeNonExported:    true,
			},
			want: enums{
				"TestType1": []enum{
					{
						Name:  "TestType1Value1",
						Value: "",
					},
					{
						Name:  "TestType1Value2",
						Value: "",
					},
				},
				"TestType2": []enum{
					{
						Name:  "TestType2Value1",
						Value: "iota",
					},
					{
						Name:  "TestType2Value2",
						Value: "",
					},
					{
						Name:  "TestType2Value3",
						Value: "",
					},
					{
						Name:  "TestType2Value4",
						Value: "",
					},
					{
						Name:  "testType2Value5",
						Value: "",
					},
				},
				"TestType3": []enum{
					{
						Name:  "TestType3Value1",
						Value: "",
					},
				},
				"TestType4": []enum{
					{
						Name:  "TestType4Value1",
						Value: "true",
					},
				},
			},
			wantErr: false,
		},
		{
			name: "2",
			args: args{
				filePath:              "./parser_test.go",
				includeNonExported:    false,
				targetObjectTypeNames: []constantObjectTypeName{},
			},
			want: enums{
				"TestType1": []enum{
					{
						Name:  "TestType1Value1",
						Value: "",
					},
					{
						Name:  "TestType1Value2",
						Value: "",
					},
				},
				"TestType2": []enum{
					{
						Name:  "TestType2Value1",
						Value: "iota",
					},
					{
						Name:  "TestType2Value2",
						Value: "",
					},
					{
						Name:  "TestType2Value3",
						Value: "",
					},
					{
						Name:  "TestType2Value4",
						Value: "",
					},
				},
				"TestType3": []enum{
					{
						Name:  "TestType3Value1",
						Value: "",
					},
				},
				"TestType4": []enum{
					{
						Name:  "TestType4Value1",
						Value: "true",
					},
				},
			},
			wantErr: false,
		},
		{
			name: "3",
			args: args{
				filePath:              "./parser_test.go",
				includeNonExported:    false,
				targetObjectTypeNames: []constantObjectTypeName{"TestType3"},
			},
			want: enums{
				"TestType3": []enum{
					{
						Name:  "TestType3Value1",
						Value: "",
					},
				},
			},
			wantErr: false,
		},
		{
			name: "4",
			args: args{
				filePath:              "./parser_test.go",
				includeNonExported:    false,
				targetObjectTypeNames: []constantObjectTypeName{"TestType5"},
			},
			want:    enums{},
			wantErr: false,
		},
		{
			name: "5",
			args: args{
				filePath:           "./parser_test1.go",
				includeNonExported: false,
			},
			wantErr: true,
		},
	}
	for _, testCase := range tests {
		testCase := testCase
		t.Run(
			testCase.name,
			func(t *testing.T) {
				t.Parallel()

				got, err := parseEnumsFromFile(
					testCase.args.filePath,
					testCase.args.targetObjectTypeNames,
					testCase.args.includeNonExported,
				)
				if (err != nil) != testCase.wantErr {
					t.Errorf("parseEnumsFromFile() error = %v, wantErr %v", err, testCase.wantErr)

					return
				}
				if !reflect.DeepEqual(got, testCase.want) {
					t.Errorf("parseEnumsFromFile() got = %v, want %v", got, testCase.want)
				}
			},
		)
	}
}
