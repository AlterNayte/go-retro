package main

import (
	"flag"
	"fmt"
	"go-retro/generator"
	"log"
	"os"
)

func main() {
	//var catsAPI = goretro.NewCatFactsAPIClient("https://cat-fact.herokuapp.com")
	//data, derr := catsAPI.Facts()
	//if derr != nil {
	//	log.Fatalf("Error getting facts: %v", derr)
	//}
	//fmt.Printf("Facts: %v\n", data)
	//
	//var httpbinapi = goretro.NewHttpAPIClient("https://httpbin.org")
	//bindata, derr := httpbinapi.PostExample(httpbin.PostInput{Boots: 2})
	//if derr != nil {
	//	log.Fatalf("Error posting example: %v", derr)
	//}
	//fmt.Printf("Post example: %v\n", bindata)

	var outputDir string
	var workingDir string

	flag.StringVar(&outputDir, "output", "goretro/generated", "Directory to output generated client")
	flag.StringVar(&workingDir, "dir", ".", "Working directory to search for .go files")
	flag.Parse()

	if _, err := os.Stat(workingDir); os.IsNotExist(err) {
		log.Fatalf("Working directory does not exist: %s", workingDir)
	}

	err := generator.Generate(outputDir, workingDir)
	if err != nil {
		log.Fatalf("Error generating client: %v", err)
	}

	fmt.Printf("Client generated successfully in %s\n", outputDir)

}
