package shared

import (
	"encoding/json"
	"encoding/xml"
	"errors"
	"flag"
	"log"
	"strings"
)

const (
	jsonExt = ".json"
	xmlExt  = ".xml"
)

type Ingredient struct {
	Name  string `json:"ingredient_name" xml:"itemname"`
	Count string `json:"ingredient_count" xml:"itemcount"`
	Unit  string `json:"ingredient_unit,omitempty"  xml:"itemunit"`
}

type Cake struct {
	Name        string       `json:"name" xml:"name"`
	Time        string       `json:"time" xml:"stovetime"`
	Ingredients []Ingredient `json:"ingredients" xml:"ingredients>item"`
}

type RecipeBook struct {
	XMLName xml.Name `json:"-" xml:"recipes"`
	Cakes   []Cake   `json:"cake" xml:"cake"`
}

func Reader(d DBReader) RecipeBook {
	return d.Read()
}

type DBReader interface {
	Read() RecipeBook
}

type JSONbytes []byte
type XMLbytes []byte

func (j JSONbytes) Read() RecipeBook {
	res := RecipeBook{}
	if err := json.Unmarshal([]byte(j), &res); err != nil {
		log.Fatal(err)
	}
	return res
}

func (x XMLbytes) Read() RecipeBook {
	res := RecipeBook{}
	if err := xml.Unmarshal([]byte(x), &res); err != nil {
		log.Fatal(err)
	}
	return res
}

func CheckInputReadDB() (filename, fileExtension string, err error) {
	filenameFlag := flag.String("f", "", "Usage: [-f] filename")
	flag.Parse()

	if flag.NArg() != 0 {
		err = errors.New("unexpected arguments.\nUsage: ./readDB [-f] filename")
		return
	}

	filename = *filenameFlag

	switch {
	case filename == "":
		return "", "", errors.New("filename is required")
	case strings.HasSuffix(filename, jsonExt):
		return filename, jsonExt, nil
	case strings.HasSuffix(filename, xmlExt):
		return filename, xmlExt, nil
	default:
		return "", "", errors.New("unsupported file extension")
	}
}
