package extractor //nolint: testpackage // testing unexported functions

import (
	"bytes"
	"testing"
)

//nolint:dupl,funlen // FP // huge test cases
func Test_ExecuteRawSwitchCaseTemplate(t *testing.T) {
	t.Parallel()

	type args struct {
		objects     enums
		packageName string
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{
			name: "1",
			args: args{
				objects:     enums{constantObjectTypeName("test"): []enum{{Name: "test1", Value: "iota"}}},
				packageName: "",
			},
			want: `	switch input { 
	case test1:
		// TODO: add case expression
	}

`,
			wantErr: false,
		},
		{
			name: "2",
			args: args{
				objects:     enums{constantObjectTypeName("test"): []enum{{Name: "test1", Value: "iota"}}},
				packageName: "testPackage",
			},
			want: `	switch testPackage.input { 
	case testPackage.test1:
		// TODO: add case expression
	}

`,
			wantErr: false,
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(
			tt.name,
			func(t *testing.T) {
				t.Parallel()

				buf := &bytes.Buffer{}

				err := executeTemplate(buf, RawTemplateType, tt.args.packageName, tt.args.objects)
				if (err != nil) != tt.wantErr {
					t.Errorf("ExecuteRawSwitchCaseTemplate() error = %v, wantErr %v", err, tt.wantErr)

					return
				}
				if buf.String() != tt.want {
					t.Errorf("executeTemplate() got = %v, want %v", buf.String(), tt.want)
				}
			},
		)
	}
}

//nolint:dupl,funlen // FP // huge test cases
func Test_ExecuteFunctionSwitchCaseTemplate(t *testing.T) {
	t.Parallel()

	type args struct {
		objects     enums
		packageName string
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{
			name: "1",
			args: args{
				objects:     enums{constantObjectTypeName("test"): []enum{{Name: "test1", Value: "iota"}}},
				packageName: "",
			},
			want: `func blob(input test) {
	switch input { 
	case test1:
		// TODO: add case expression
	}
}

`,
			wantErr: false,
		},
		{
			name: "2",
			args: args{
				objects:     enums{constantObjectTypeName("test"): []enum{{Name: "test1", Value: "iota"}}},
				packageName: "testPackage",
			},
			want: `func blob(input testPackage.test) {
	switch testPackage.input { 
	case testPackage.test1:
		// TODO: add case expression
	}
}

`,
			wantErr: false,
		},
	}

	for _, testCase := range tests {
		testCase := testCase
		t.Run(
			testCase.name,
			func(t *testing.T) {
				t.Parallel()

				buf := &bytes.Buffer{}

				err := executeTemplate(buf, FunctionTemplateType, testCase.args.packageName, testCase.args.objects)
				if (err != nil) != testCase.wantErr {
					t.Errorf("ExecuteFunctionSwitchCaseTemplate() error = %v, wantErr %v", err, testCase.wantErr)

					return
				}
				if buf.String() != testCase.want {
					t.Errorf("executeTemplate() got = %v, want %v", buf.String(), testCase.want)
				}
			},
		)
	}
}

func Test_createTemplate(t *testing.T) {
	t.Parallel()

	type args struct {
		templateType TemplateType
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "raw",
			args: args{
				templateType: RawTemplateType,
			},
			wantErr: false,
		},
		{
			name: "function",
			args: args{
				templateType: FunctionTemplateType,
			},
			wantErr: false,
		},
		{
			name: "not valid",
			args: args{
				templateType: testNotValid,
			},
			wantErr: true,
		},
		{
			name: "unknown",
			args: args{
				templateType: "",
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		tt := tt

		t.Run(
			tt.name,
			func(t *testing.T) {
				t.Parallel()

				_, err := createTemplate(tt.args.templateType)
				if (err != nil) != tt.wantErr {
					t.Errorf("createTemplate() error = %v, wantErr %v", err, tt.wantErr)
					return
				}
			},
		)
	}
}
