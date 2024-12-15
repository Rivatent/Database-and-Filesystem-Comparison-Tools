package main

import (
	"comparing/shared"
	"encoding/json"
	"encoding/xml"
	"fmt"
	"log"
	"os"
)

const (
	jsonExt = ".json"
	xmlExt  = ".xml"
)

func main() {
	filename, fileExtension, err := shared.CheckInputReadDB()
	if err != nil {
		log.Fatal(err)
	}

	fileData, err := os.ReadFile(filename)
	if err != nil {
		log.Fatal(err)
	}

	if fileExtension == jsonExt && !json.Valid(fileData) {
		log.Fatalf("invalid .json file: %v", filename)
	}

	recipeBook := shared.RecipeBook{}
	if fileExtension == ".json" {
		recipeBook = shared.Reader(shared.JSONbytes(fileData))
	} else if fileExtension == ".xml" {
		recipeBook = shared.Reader(shared.XMLbytes(fileData))
	}

	PrintAnotherFormat(fileExtension, recipeBook)
}

func PrintAnotherFormat(fileExtension string, recipeBook shared.RecipeBook) {
	switch fileExtension {
	case jsonExt:
		xmlFormat, err := xml.MarshalIndent(recipeBook, "", "    ")
		if err != nil {
			log.Fatal(err)
		}
		fmt.Print(string(xmlFormat))
	case xmlExt:
		jsonFormat, err := json.MarshalIndent(recipeBook, "", "    ")
		if err != nil {
			log.Fatal(err)
		}
		fmt.Print(string(jsonFormat))
	}
	fmt.Println()
}
