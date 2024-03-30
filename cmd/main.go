package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	imglocorg "github.com/timotewb/go-tools/cmd/img-loc-org"
	shared "github.com/timotewb/go-tools/cmd/shared"
)

func main() {

	var toolName string
	var apiKey string
	var help bool
	var inDir string
	var outDir string
	// Define CLI flags in shrot and long form
	flag.StringVar(&toolName, "t", "", "Name of tool to run (shorthand)")
	flag.StringVar(&toolName, "tool", "", "Name of tool to run")
	flag.StringVar(&inDir, "i", "", "")
	flag.StringVar(&inDir, "in-dir", "", "")
	flag.StringVar(&outDir, "o", "", "")
	flag.StringVar(&outDir, "out-dir", "", "")
	flag.StringVar(&apiKey, "k", "", "API key (shorthand)")
	flag.StringVar(&apiKey, "api-key", "", "API key")
	flag.BoolVar(&help, "h", false, "Show usage instructions (shorthand)")
	flag.BoolVar(&help, "help", false, "Show usage instructions")
	flag.Usage = func() {
		fmt.Fprintln(os.Stderr, "----------------------------------------------------------------------------------------")
		fmt.Fprintf(os.Stderr, "Usage of %s:\n\n", os.Args[0])
		fmt.Fprintln(os.Stderr, "Parameters:")
		fmt.Fprintln(os.Stderr, "  -t\t\tstring\t\tName of tool to run, required if not passing -h\n  --tool")
		fmt.Fprintln(os.Stderr, "\n  -h\t\tboolean\t\tShow usage instructions, required if not passing -t\n  --help")
		fmt.Fprintln(os.Stderr, "  -i\t\tstring\t\tInput directory\n  --in-dir")
		fmt.Fprintln(os.Stderr, "  -k\t\tstring\t\tAPI key (string)\n  --api-key")
		fmt.Fprintln(os.Stderr, "  -o\t\tstring\t\tOutput directory\n  --out-dir")
		fmt.Fprintln(os.Stderr, "\n----------------------------------------------------------------------------------------")
	}
	flag.Parse()

	// Print the Help docuemntation to the terminal if user passes help flag
	if help {
		flag.Usage()
		return
	}

	// Specify valid tool names
	validToolNames := []string{"img-loc-org"}
	if shared.Contains(validToolNames, toolName) {
		if toolName == "img-loc-org" {
			err := imglocorg.App(help, apiKey, inDir, outDir)
			if err != nil {
				log.Fatal(err)
			}
		}
	} else {
		log.Fatal("tool name not valid or no tool name provided.")
	}
}
