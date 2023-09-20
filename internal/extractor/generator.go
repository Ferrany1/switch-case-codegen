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

	e, err := parseEnumsFromFile(filePath, targetObjectTypeNames, includeNonExported)
	if err != nil {
		return eris.Wrap(err, "parsing enums from file")
	}

	return executeTemplate(output, templateType, packageName, e)
}
