package extractor

import (
	"go/ast"
	"unicode"
)

type (
	enum struct {
		Name  constantName
		Value constantValue
	}
	enums map[constantObjectTypeName][]enum
)

func (e enums) constantObjectTypeNameIfConsistsOfIota() constantObjectTypeName {
	if len(e) != 1 {
		return ""
	}

	for key := range e {
		if len(e[key]) > 0 && e[key][0].Value == "iota" {
			return key
		}

		break
	}

	return ""
}

type (
	constantName           string
	constantObjectTypeName string
	constantValue          string
)

//nolint:revive // exporter to convert outside of internal
func NewConstantObjectTypeNamesFromStringSlice(input []string) []constantObjectTypeName {
	result := make([]constantObjectTypeName, len(input))
	for i := range input {
		result[i] = constantObjectTypeName(input[i])
	}

	return result
}

func extractConstantsBlock(
	constantDeclarationSpecification *ast.ValueSpec,
	includeNonExported bool,
) (constantName, constantObjectTypeName, constantValue) {
	var (
		objectName     string
		objectTypeName string
		objectValue    string
	)

	if len(constantDeclarationSpecification.Names) == 1 {
		objectName = constantDeclarationSpecification.Names[0].Name
	}

	if len(objectName) == 0 || (unicode.IsLower([]rune(objectName)[0]) && !includeNonExported) {
		return "", "", ""
	}

	if len(constantDeclarationSpecification.Values) == 1 {
		value, valueOk := constantDeclarationSpecification.Values[0].(*ast.Ident)
		if valueOk {
			objectValue = value.Name
		}
	}

	constantTypeObject, ok := constantDeclarationSpecification.Type.(*ast.Ident)
	if ok && constantTypeObject.Obj == nil {
		return "", "", ""
	} else if ok {
		objectTypeName = constantTypeObject.Name
	}

	return constantName(objectName), constantObjectTypeName(objectTypeName), constantValue(objectValue)
}
