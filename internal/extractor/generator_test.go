package extractor //nolint: testpackage // testing unexported functions

import (
	"bytes"
	"testing"
)

//nolint:funlen // test function
func TestGenerateSwitchCases(t *testing.T) {
	t.Parallel()

	type args struct {
		filePath              string
		packageName           string
		templateType          TemplateType
		targetObjectTypeNames []constantObjectTypeName
		includeNonExported    bool
	}
	tests := []struct {
		name       string
		args       args
		wantOutput string
		wantErr    bool
	}{
		{
			name: "1",
			args: args{
				filePath:              "./parser_test.go",
				packageName:           "test_package",
				templateType:          RawTemplateType,
				targetObjectTypeNames: nil,
				includeNonExported:    true,
			},
			wantOutput: `	switch test_package.input { 
	case test_package.TestType1Value1:
		// TODO: add case expression
	case test_package.TestType1Value2:
		// TODO: add case expression
	}

	switch test_package.input { 
	case test_package.TestType2Value1:
		// TODO: add case expression
	case test_package.TestType2Value2:
		// TODO: add case expression
	case test_package.TestType2Value3:
		// TODO: add case expression
	case test_package.TestType2Value4:
		// TODO: add case expression
	case test_package.testType2Value5:
		// TODO: add case expression
	}

	switch test_package.input { 
	case test_package.TestType3Value1:
		// TODO: add case expression
	}

	switch test_package.input { 
	case test_package.TestType4Value1:
		// TODO: add case expression
	}

`,
			wantErr: false,
		},
		{
			name: "2",
			args: args{
				filePath:              "./parser_test.go",
				packageName:           "",
				templateType:          FunctionTemplateType,
				targetObjectTypeNames: []constantObjectTypeName{"TestType4"},
				includeNonExported:    true,
			},
			wantOutput: `func blob(input TestType4) {
	switch input { 
	case TestType4Value1:
		// TODO: add case expression
	}
}

`,
			wantErr: false,
		},
		{
			name: "3",
			args: args{
				filePath:              "./parser_test.go",
				packageName:           "",
				templateType:          "test",
				targetObjectTypeNames: nil,
				includeNonExported:    false,
			},
			wantOutput: "",
			wantErr:    true,
		},
		{
			name: "4",
			args: args{
				filePath:              "./parser_test1.go",
				packageName:           "",
				templateType:          RawTemplateType,
				targetObjectTypeNames: nil,
				includeNonExported:    false,
			},
			wantOutput: "",
			wantErr:    true,
		},
	}
	for _, testCase := range tests {
		testCase := testCase

		t.Run(
			testCase.name,
			func(t *testing.T) {
				t.Parallel()

				output := &bytes.Buffer{}
				err := GenerateSwitchCases(
					testCase.args.filePath,
					testCase.args.packageName,
					testCase.args.templateType,
					output,
					testCase.args.targetObjectTypeNames,
					testCase.args.includeNonExported,
				)
				if (err != nil) != testCase.wantErr {
					t.Errorf("GenerateSwitchCases() error = %v, wantErr %v", err, testCase.wantErr)
					return
				}
				if gotOutput := output.String(); gotOutput != testCase.wantOutput {
					t.Errorf("GenerateSwitchCases() gotOutput = %v, want %v", gotOutput, testCase.wantOutput)
				}
			},
		)
	}
}
