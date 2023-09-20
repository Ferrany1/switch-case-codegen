package extractor

import (
	"fmt"
	"io"
	"text/template"

	"github.com/rotisserie/eris"
)

type TemplateType string

const (
	RawTemplateType      TemplateType = "raw"
	FunctionTemplateType TemplateType = "function"
	testNotValid         TemplateType = "not valid"
)

const (
	switchCaseRangeOverEnumsTemplate = "{{ range $key, $value := .Objects }}" +
		"%s\n" +
		"{{ end }}"
	switchCaseFunctionTemplate = "func blob(input " +
		"{{ if $.PackageName }}{{ $.PackageName }}.{{ end }}" + "{{ $key }}) {\n" +
		switchCaseRangeOverEnumValuesTemplate +
		"}\n"
	switchCaseRangeOverEnumValuesTemplate = "\tswitch {{ if $.PackageName }}{{ $.PackageName }}.{{ end }}input { " +
		"{{ range $value }}\n" +
		"\tcase {{ if $.PackageName }}{{ $.PackageName }}.{{ end }}{{ .Name }}:\n" +
		"\t\t// TODO: add case expression" +
		"{{ end }}\n" +
		"\t}\n"
)

func createTemplate(templateType TemplateType) (*template.Template, error) {
	var (
		name string
		text string
	)

	switch templateType {
	case RawTemplateType:
		name = "rawSwitchCaseTemplate"
		text = fmt.Sprintf(switchCaseRangeOverEnumsTemplate, switchCaseRangeOverEnumValuesTemplate)
	case FunctionTemplateType:
		name = "functionSwitchCaseTemplate"
		text = fmt.Sprintf(switchCaseRangeOverEnumsTemplate, switchCaseFunctionTemplate)
	case testNotValid:
		name = ""
		text = "{{{ . }}"
	default:
		return nil, eris.Errorf("not valid template type: %s", templateType)
	}

	t, err := template.New(name).Parse(text)
	if err != nil {
		return nil, eris.Wrap(err, "parsing template")
	}

	return t, nil
}

func executeTemplate(writer io.Writer, templateType TemplateType, packageName string, objects enums) error {
	type templateData struct {
		Objects     enums
		PackageName string
	}

	t, err := createTemplate(templateType)
	if err != nil {
		return eris.Wrap(err, "creating template")
	}

	return eris.Wrap(
		t.Execute(
			writer,
			templateData{
				Objects:     objects,
				PackageName: packageName,
			},
		),
		"executing template",
	)
}
