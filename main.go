package main

import (
	"flag"
	"fmt"
	"go-retro/generator"
	"log"
	"os"
)

var (
	version = "dev"
	commit  = "none"
	date    = "unknown"
)

func main() {
	var outputDir string
	var workingDir string
	var versionFlag bool

	flag.StringVar(&outputDir, "output", "go-retro/generated", "Directory to output generated client")
	flag.StringVar(&workingDir, "dir", ".", "Working directory to search for .go files")
	flag.BoolVar(&versionFlag, "v", false, "Prints the version of the program")
	flag.Parse()

	if versionFlag {
		fmt.Println("GoRetro Version:", version)
		os.Exit(0)
	}

	if _, err := os.Stat(workingDir); os.IsNotExist(err) {
		log.Fatalf("Working directory does not exist: %s", workingDir)
	}

	err := generator.Generate(outputDir, workingDir)
	if err != nil {
		log.Fatalf("Error generating client: %v", err)
	}

	fmt.Printf("Client generated successfully in %s\n", outputDir)

}
