package main

import (
	"log"
	"os"

	"github.com/rotisserie/eris"
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"

	"github.com/Ferrany1/switch-case-codegen/internal/extractor"
)

func main() {
	cmd := &cobra.Command{}
	cmd.AddCommand(generateSwitchCases())
	pflag.BoolP("verbose", "v", false, "verbose error output")

	if err := cmd.Execute(); err != nil {
		verbose, _ := cmd.Flags().GetBool("verbose")
		log.Println(eris.ToString(err, verbose))
		os.Exit(1)
	}
}

func generateSwitchCases() *cobra.Command {
	cmd := &cobra.Command{
		Use: "enum",
		RunE: func(cmd *cobra.Command, args []string) error {
			filePath, _ := cmd.PersistentFlags().GetString("file_path")
			packageName, _ := cmd.Flags().GetString("package_name")
			templateType, _ := cmd.Flags().GetString("template_type")
			var output = os.Stdout
			outputFilePath, _ := cmd.Flags().GetString("output")
			if len(outputFilePath) > 0 {
				f, err := os.OpenFile(outputFilePath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
				if err != nil {
					return eris.Wrapf(err, "opening output: %s", outputFilePath)
				}

				output = f
			}
			types, _ := cmd.Flags().GetStringSlice("types")
			includeNonExported, _ := cmd.Flags().GetBool("include_exported")

			return eris.Wrap(
				extractor.GenerateSwitchCases(
					filePath,
					packageName,
					extractor.TemplateType(templateType),
					output,
					extractor.NewConstantObjectTypeNamesFromStringSlice(types),
					includeNonExported,
				),
				"generating switch case",
			)
		},
	}

	cmd.PersistentFlags().StringP("file_path", "p", "", "path to file with enums")
	cmd.Flags().StringP("package_name", "n", "", "generated import package name (default:empty)")
	cmd.Flags().StringP(
		"template_type",
		"t",
		string(extractor.RawTemplateType),
		"template to generate. options: raw|function (default:raw)",
	)
	cmd.Flags().StringP("output", "o", "", "output file (default:stdout)")
	cmd.Flags().StringSlice("types", []string{}, "types to generate (default:all)")
	cmd.Flags().BoolP("include_exported", "i", true, "include non exported enums (default:false)")

	return cmd
}
