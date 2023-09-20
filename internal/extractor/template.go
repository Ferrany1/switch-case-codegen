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

func executeRawSwitchCaseTemplate(writer io.Writer, packageName string, objects enums) error {
	rawSwitchCaseTemplate, err := template.New("rawSwitchCaseTemplate").
		Parse(fmt.Sprintf(switchCaseRangeOverEnumsTemplate, switchCaseRangeOverEnumValuesTemplate))
	if err != nil {
		return eris.Wrap(err, "parsing raw switch case template")
	}

	return executeTemplate(writer, rawSwitchCaseTemplate, packageName, objects)
}

func executeFunctionSwitchCaseTemplate(writer io.Writer, packageName string, objects enums) error {
	functionSwitchCaseTemplate, err := template.New("functionSwitchCaseTemplate").
		Parse(fmt.Sprintf(switchCaseRangeOverEnumsTemplate, switchCaseFunctionTemplate))
	if err != nil {
		return eris.Wrap(err, "parsing raw switch case template")
	}

	return executeTemplate(writer, functionSwitchCaseTemplate, packageName, objects)
}

func executeTemplate(writer io.Writer, t *template.Template, packageName string, objects enums) error {
	type templateData struct {
		Objects     enums
		PackageName string
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
