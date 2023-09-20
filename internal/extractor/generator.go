package extractor

import (
	"io"

	"github.com/rotisserie/eris"
)

func GenerateSwitchCases(
	filePath,
	packageName string,
	templateType TemplateType,
	output io.Writer,
	targetObjectTypeNames []constantObjectTypeName,
	includeNonExported bool,
) error {
	if len(packageName) > 0 {
		includeNonExported = true
	}

	var executeTemplateFunction func(writer io.Writer, packageName string, objects enums) error
	switch templateType {
	case RawTemplateType:
		executeTemplateFunction = executeRawSwitchCaseTemplate
	case FunctionTemplateType:
		executeTemplateFunction = executeFunctionSwitchCaseTemplate
	default:
		return eris.New("invalid template")
	}

	e, err := parseEnumsFromFile(filePath, targetObjectTypeNames, includeNonExported)
	if err != nil {
		return eris.Wrap(err, "parsing enums from file")
	}

	return executeTemplateFunction(output, packageName, e)
}
