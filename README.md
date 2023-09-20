[![codecov](https://codecov.io/gh/Ferrany1/switch-case-codegen/graph/badge.svg?token=UY1SM9YT0U)](https://codecov.io/gh/Ferrany1/switch-case-codegen)
# Switch Case Codegen

## Overview:
Tool for generating switch cases for enums.

## Install:
```
go install github.com/Ferrany1/switch-case-codegen/cmd/switch-case-codegen
```

## Usage:
### Cli:
```
switch-case-codegen enum [flags]

Flags:
-p, --file_path string       path to file with enums
-h, --help                   help for enum
-i, --include_exported       include non exported enums (default:false) (default true)
-o, --output string          output file (default:stdout)
-n, --package_name string    generated import package name (default:empty)
-t, --template_type string   template to generate. options: raw|function (default:raw) (default "raw")
--types strings          types to generate (default:all)

Global Flags:
-v, --verbose   verbose error output
```
### Go file:
```
//go:generate switch-case-codegen enum -p ./enums.go -t function -n test_package
```

## Example:

### Function:
enum.go:
```golang
//go:generate switch-case-codegen enum -p ./enum.go -t function -n test_package --types TestType2

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
```
output:
```golang
func blob(input test_package.TestType2) {
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
}
```

### Raw:
enum.go:
```golang
//go:generate switch-case-codegen enum -p ./enum.go -n test_package --types TestType2

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
```
output:
```golang
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
```