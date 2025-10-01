package templates

import "fmt"

// GetCLIAppTemplates returns templates for CLI application projects
func GetCLIAppTemplates(framework string, data TemplateData) []Template {
	templates := GetCommonTemplates(data)
	
	templates = append(templates, []Template{
		{
			Path:    fmt.Sprintf("cmd/%s/main.go", data.ProjectName),
			Content: GetCLIMainTemplate(framework, data),
		},
		{
			Path:    "internal/cli/root.go",
			Content: GetCLIRootTemplate(framework, data),
		},
		{
			Path:    "internal/cli/version.go",
			Content: GetCLIVersionTemplate(framework, data),
		},
	}...)
	
	return templates
}

func GetCLIMainTemplate(framework string, data TemplateData) string {
	switch framework {
	case "cobra":
		return fmt.Sprintf(`package main

import (
	"%s/internal/cli"
)

func main() {
	cli.Execute()
}
`, data.ModuleName)
	default:
		return fmt.Sprintf(`package main

import (
	"fmt"
	"os"

	"%s/internal/cli"
)

func main() {
	if err := cli.Run(os.Args); err != nil {
		fmt.Fprintf(os.Stderr, "Error: %%v\n", err)
		os.Exit(1)
	}
}
`, data.ModuleName)
	}
}

func GetCLIRootTemplate(framework string, data TemplateData) string {
	switch framework {
	case "cobra":
		return fmt.Sprintf(`package cli

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "%s",
	Short: "%s",
	Long:  `+"`%s`"+`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Hello from %s!")
		fmt.Println("Use --help for available commands")
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintf(os.Stderr, "Error: %%v\n", err)
		os.Exit(1)
	}
}

func init() {
	rootCmd.AddCommand(versionCmd)
}
`, data.ProjectName, data.Description, data.Description, data.ProjectName)
	default:
		return fmt.Sprintf(`package cli

import (
	"flag"
	"fmt"
)

func Run(args []string) error {
	fs := flag.NewFlagSet("%s", flag.ExitOnError)
	
	var (
		version = fs.Bool("version", false, "Show version information")
		help    = fs.Bool("help", false, "Show help information")
	)
	
	if err := fs.Parse(args[1:]); err != nil {
		return err
	}
	
	if *help {
		fmt.Printf("%%s - %%s\n\n", "%s", "%s")
		fmt.Println("Usage:")
		fs.PrintDefaults()
		return nil
	}
	
	if *version {
		fmt.Println(GetVersion())
		return nil
	}
	
	fmt.Println("Hello from %s!")
	return nil
}
`, data.ProjectName, data.ProjectName, data.Description, data.ProjectName)
	}
}

func GetCLIVersionTemplate(framework string, data TemplateData) string {
	switch framework {
	case "cobra":
		return `package cli

import (
	"fmt"

	"github.com/spf13/cobra"
)

var Version = "1.0.0"

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the version number",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("` + data.ProjectName + ` v%s\n", Version)
	},
}
`
	default:
		return `package cli

import "fmt"

var Version = "1.0.0"

func GetVersion() string {
	return fmt.Sprintf("` + data.ProjectName + ` v%s", Version)
}
`
	}
}