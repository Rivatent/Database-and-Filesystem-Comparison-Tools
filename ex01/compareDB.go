package main

import (
	"comparing/shared"
	"errors"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

const (
	jsonExt = ".json"
	xmlExt  = ".xml"
)

func main() {
	oldDatabaseFilepath, newDatabaseFilepath, oldDatabaseExtension, newDatabaseExtension, error := CheckInputCompareDB()
	if error != nil {
		log.Fatal(error)
	}

	oldDatabaseFile, error := os.ReadFile(oldDatabaseFilepath)
	if error != nil {
		log.Fatal(error)
	}

	newDatabaseFile, error := os.ReadFile(newDatabaseFilepath)
	if error != nil {
		log.Fatal(error)
	}

	oldRecipeBook := shared.RecipeBook{}
	newRecipeBook := shared.RecipeBook{}

	switch oldDatabaseExtension {
	case jsonExt:
		oldRecipeBook = shared.Reader(shared.JSONbytes(oldDatabaseFile))
	case xmlExt:
		oldRecipeBook = shared.Reader(shared.XMLbytes(oldDatabaseFile))
	}

	switch newDatabaseExtension {
	case jsonExt:
		newRecipeBook = shared.Reader(shared.JSONbytes(newDatabaseFile))
	case xmlExt:
		newRecipeBook = shared.Reader(shared.XMLbytes(newDatabaseFile))
	}

	CompareCakes(os.Stdout, oldRecipeBook, newRecipeBook)
}

func CompareCakes(out io.Writer, oldRecipeBook, newRecipeBook shared.RecipeBook) {
	oldCakes := map[string]shared.Cake{}
	newCakes := map[string]shared.Cake{}

	for _, cake := range oldRecipeBook.Cakes {
		oldCakes[cake.Name] = cake
	}

	for _, cake := range newRecipeBook.Cakes {
		newCakes[cake.Name] = cake
	}

	for cakeName := range newCakes {
		if _, exist := oldCakes[cakeName]; !exist {
			fmt.Fprintf(out, "ADDED cake \"%s\"\n", cakeName)
		}
	}

	for cakeName := range oldCakes {
		if _, exist := newCakes[cakeName]; !exist {
			fmt.Fprintf(out, "REMOVED cake \"%s\"\n", cakeName)
		}
	}

	for cakeName, oldCake := range oldCakes {
		if newCake, exist := newCakes[cakeName]; exist {
			if oldCake.Time != newCake.Time {
				fmt.Fprintf(out, "CHANGED cooking time for cake \"%s\" - \"%s\" instead of \"%s\"\n", cakeName, newCake.Time, oldCake.Time)
			}
			CompareIngredients(out, oldCake, newCake)
		}
	}
}

func CheckInputCompareDB() (oldDatabaseFilepath, newDatabaseFilepath, oldDatabaseExtension, newDatabaseExtension string, err error) {
	oldDatabaseFlag := flag.String("old", "", "usage: ./compareDB [-old] oldDatabaseFilepath [-new] newDatabaseFilepath")
	newDatabaseFlag := flag.String("new", "", "usage: ./compareDB [-old] oldDatabaseFilepath [-new] newDatabaseFilepath")
	flag.Parse()
	if flag.NArg() != 0 {
		err = errors.New("usage: ./compareDB [-old] oldDatabaseFilepath [-new] newDatabaseFilepath")
		return
	}
	oldDatabaseFilepath = *oldDatabaseFlag
	newDatabaseFilepath = *newDatabaseFlag

	switch {
	case oldDatabaseFilepath == "":
		err = errors.New("filename is required")
		return
	case strings.HasSuffix(oldDatabaseFilepath, jsonExt):
		oldDatabaseExtension = jsonExt
	case strings.HasSuffix(oldDatabaseFilepath, xmlExt):
		oldDatabaseExtension = xmlExt
	default:
		err = errors.New("unsupported file extension")
		return
	}

	switch {
	case newDatabaseFilepath == "":
		err = errors.New("filename is required")
		return
	case strings.HasSuffix(newDatabaseFilepath, jsonExt):
		newDatabaseExtension = jsonExt
	case strings.HasSuffix(newDatabaseFilepath, xmlExt):
		newDatabaseExtension = xmlExt
	default:
		err = errors.New("unsupported file extension")
		return
	}

	return
}

func CompareIngredients(out io.Writer, oldCake, newCake shared.Cake) {
	oldIngredients := map[string]shared.Ingredient{}
	newIngredients := map[string]shared.Ingredient{}

	for _, ingredient := range oldCake.Ingredients {
		oldIngredients[ingredient.Name] = ingredient
	}

	for _, ingredient := range newCake.Ingredients {
		newIngredients[ingredient.Name] = ingredient
	}

	// for add
	for ingredientName := range newIngredients {
		if _, exist := oldIngredients[ingredientName]; !exist {
			fmt.Fprintf(out, "ADDED ingredient \"%s\" for cake \"%s\"\n", ingredientName, oldCake.Name)
		}
	}
	// for remove
	for ingredientName := range oldIngredients {
		if _, exist := newIngredients[ingredientName]; !exist {
			fmt.Fprintf(out, "REMOVED ingredient \"%s\" for cake \"%s\"\n", ingredientName, oldCake.Name)
		}
	}
	// for change
	for ingredientName, oldIngredient := range oldIngredients {
		if _, exist := newIngredients[ingredientName]; exist {
			if oldIngredient.Unit != newIngredients[ingredientName].Unit {
				if newIngredients[ingredientName].Unit == "" && oldIngredient.Unit != "" {
					fmt.Fprintf(out, "REMOVED unit \"%s\" for ingredient \"%s\" for cake \"%s\"\n", oldIngredient.Unit, ingredientName, oldCake.Name)
				} else if oldIngredient.Unit == "" && newIngredients[ingredientName].Unit != "" {
					fmt.Fprintf(out, "ADDED unit \"%s\" for ingredient \"%s\" for cake \"%s\"\n", newIngredients[ingredientName].Unit, ingredientName, oldCake.Name)
				} else {
					fmt.Fprintf(out, "CHANGED unit for ingredient \"%s\" for cake \"%s\" - \"%s\" instead of \"%s\"\n", ingredientName, oldCake.Name, newIngredients[ingredientName].Unit, oldIngredient.Unit)
				}
			}
			if oldIngredient.Count != newIngredients[ingredientName].Count {
				fmt.Fprintf(out, "CHANGED unit count for ingredient \"%s\" for cake \"%s\" - \"%s\" instead of \"%s\"\n", ingredientName, oldCake.Name, newIngredients[ingredientName].Count, oldIngredient.Count)
			}
		}
	}
}
