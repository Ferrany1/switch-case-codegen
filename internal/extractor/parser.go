package extractor

import (
	"go/ast"
	"go/parser"
	"go/token"

	"github.com/rotisserie/eris"
)

func parseEnumsFromFile(
	filePath string,
	targetObjectTypeNames []constantObjectTypeName,
	includeNonExported bool,
) (enums, error) {
	fileNode, err := parser.ParseFile(token.NewFileSet(), filePath, nil, parser.ParseComments)
	if err != nil {
		return nil, eris.Wrapf(err, "parsing file %s", filePath)
	}

	var objects = make(enums)
	for i := range fileNode.Decls {
		declaration, ok := fileNode.Decls[i].(*ast.GenDecl)
		if !ok {
			continue
		}

		if declaration.Tok != token.CONST {
			continue
		}

		for key, value := range extractEnumsFromDeclaration(declaration, includeNonExported) {
			if insideTargetObjectTypeNames(key, targetObjectTypeNames) {
				objects[key] = value
			}
		}
	}

	return objects, nil
}

func extractEnumsFromDeclaration(declaration *ast.GenDecl, includeNonExported bool) enums {
	var objects = make(enums)
	for i := range declaration.Specs {
		constantDeclarationSpecification, constantDeclarationSpecificationOk := declaration.Specs[i].(*ast.ValueSpec)
		if !constantDeclarationSpecificationOk {
			continue
		}

		objectTypeName, e := extractEnumFromSpecification(
			constantDeclarationSpecification,
			objects.constantObjectTypeNameIfConsistsOfIota,
			includeNonExported,
		)

		if objectTypeName == "" {
			continue
		} else if len(e.Name) == 0 {
			continue
		}

		objects[objectTypeName] = append(objects[objectTypeName], e)
	}

	return objects
}

func extractEnumFromSpecification(
	constantDeclarationSpecification *ast.ValueSpec,
	objectTypeNameOnIotaFunction func() constantObjectTypeName,
	includeNonExported bool,
) (constantObjectTypeName, enum) {
	name, objectTypeName, value := extractConstantsBlock(constantDeclarationSpecification, includeNonExported)
	if objectTypeName == "" {
		objectTypeName = objectTypeNameOnIotaFunction()
	}

	return objectTypeName, enum{Name: name, Value: value}
}

func insideTargetObjectTypeNames(
	objectTypeName constantObjectTypeName,
	targetObjectTypeNames []constantObjectTypeName,
) bool {
	if len(targetObjectTypeNames) == 0 {
		return true
	}

	for i := range targetObjectTypeNames {
		if targetObjectTypeNames[i] == objectTypeName {
			return true
		}
	}

	return false
}
