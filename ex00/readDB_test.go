package main

import (
	"comparing/shared"
	"encoding/json"
	"encoding/xml"
	"reflect"
	"strings"
	"testing"
)

var xmlData = `<recipes>
    <cake>
        <name>Red Velvet Strawberry Cake</name>
        <stovetime>40 min</stovetime>
        <ingredients>
            <item>
                <itemname>Flour</itemname>
                <itemcount>3</itemcount>
                <itemunit>cups</itemunit>
            </item>
            <item>
                <itemname>Vanilla extract</itemname>
                <itemcount>1.5</itemcount>
                <itemunit>tablespoons</itemunit>
            </item>
            <item>
                <itemname>Strawberries</itemname>
                <itemcount>7</itemcount>
                <itemunit></itemunit>
            </item>
            <item>
                <itemname>Cinnamon</itemname>
                <itemcount>1</itemcount>
                <itemunit>pieces</itemunit>
            </item>
        </ingredients>
    </cake>
    <cake>
        <name>Blueberry Muffin Cake</name>
        <stovetime>30 min</stovetime>
        <ingredients>
            <item>
                <itemname>Baking powder</itemname>
                <itemcount>3</itemcount>
                <itemunit>teaspoons</itemunit>
            </item>
            <item>
                <itemname>Brown sugar</itemname>
                <itemcount>0.5</itemcount>
                <itemunit>cup</itemunit>
            </item>
            <item>
                <itemname>Blueberries</itemname>
                <itemcount>1</itemcount>
                <itemunit>cup</itemunit>
            </item>
        </ingredients>
    </cake>
</recipes>
`
var jsonData = `{
  "cake": [
    {
      "name": "Red Velvet Strawberry Cake",
      "time": "45 min",
      "ingredients": [
        {
          "ingredient_name": "Flour",
          "ingredient_count": "2",
          "ingredient_unit": "mugs"
        },
        {
          "ingredient_name": "Strawberries",
          "ingredient_count": "8"
        },
        {
          "ingredient_name": "Coffee Beans",
          "ingredient_count": "2.5",
          "ingredient_unit": "tablespoons"
        },
        {
          "ingredient_name": "Cinnamon",
          "ingredient_count": "1"
        }
      ]
    },
    {
      "name": "Moonshine Muffin",
      "time": "30 min",
      "ingredients": [
        {
          "ingredient_name": "Brown sugar",
          "ingredient_count": "1",
          "ingredient_unit": "mug"
        },
        {
          "ingredient_name": "Blueberries",
          "ingredient_count": "1",
          "ingredient_unit": "mug"
        }
      ]
    }
  ]
}
`

func removeWhitespace(s string) string {
	s = strings.ReplaceAll(s, " ", "")
	s = strings.ReplaceAll(s, "\t", "")
	s = strings.ReplaceAll(s, "\n", "")
	s = strings.ReplaceAll(s, "\r", "")
	return s
}

func TestStructureForJSON(t *testing.T) {
	want := jsonData

	rb := shared.RecipeBook{}
	if err := json.Unmarshal([]byte(jsonData), &rb); err != nil {
		panic(err)
	}
	got, _ := json.MarshalIndent(rb, "", "    ")

	sGot := removeWhitespace(string(got))
	sWant := removeWhitespace(string(want))

	if string(sGot) != sWant {
		t.Errorf("got %s want %s", got, want)
	}
}

func TestReadFromJSON(t *testing.T) {
	jData := shared.JSONbytes(jsonData)

	got := jData.Read()
	want := shared.RecipeBook{
		Cakes: []shared.Cake{
			{
				Name: "Red Velvet Strawberry Cake",
				Time: "45 min",
				Ingredients: []shared.Ingredient{
					{"Flour", "2", "mugs"},
					{"Strawberries", "8", ""},
					{"Coffee Beans", "2.5", "tablespoons"},
					{"Cinnamon", "1", ""},
				},
			},
			{
				Name: "Moonshine Muffin",
				Time: "30 min",
				Ingredients: []shared.Ingredient{
					{"Brown sugar", "1", "mug"},
					{"Blueberries", "1", "mug"},
				},
			},
		},
	}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %+v, want %+v", got, want)
	}
}

func TestJSONReaderInterface(t *testing.T) {
	jData := shared.JSONbytes(jsonData)
	got := shared.Reader(&jData)
	want := shared.RecipeBook{
		Cakes: []shared.Cake{
			{
				Name: "Red Velvet Strawberry Cake",
				Time: "45 min",
				Ingredients: []shared.Ingredient{
					{"Flour", "2", "mugs"},
					{"Strawberries", "8", ""},
					{"Coffee Beans", "2.5", "tablespoons"},
					{"Cinnamon", "1", ""},
				},
			},
			{
				Name: "Moonshine Muffin",
				Time: "30 min",
				Ingredients: []shared.Ingredient{
					{"Brown sugar", "1", "mug"},
					{"Blueberries", "1", "mug"},
				},
			},
		},
	}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %+v, want %+v", got, want)
	}
}

func TestStructureForXML(t *testing.T) {
	want := xmlData

	rb := shared.RecipeBook{}
	if err := xml.Unmarshal([]byte(xmlData), &rb); err != nil {
		panic(err)
	}
	got, _ := xml.MarshalIndent(rb, "", "    ")

	sGot := removeWhitespace(string(got))
	sWant := removeWhitespace(string(want))

	if string(sGot) != sWant {
		t.Errorf("got %s want %s", got, want)
	}
}

func TestReadFromXML(t *testing.T) {
	xData := shared.XMLbytes(xmlData)

	got := xData.Read()
	want := shared.RecipeBook{
		XMLName: xml.Name{Local: "recipes"},
		Cakes: []shared.Cake{
			{
				Name: "Red Velvet Strawberry Cake",
				Time: "40 min",
				Ingredients: []shared.Ingredient{
					{Name: "Flour", Count: "3", Unit: "cups"},
					{Name: "Vanilla extract", Count: "1.5", Unit: "tablespoons"},
					{Name: "Strawberries", Count: "7", Unit: ""},
					{Name: "Cinnamon", Count: "1", Unit: "pieces"},
				},
			},
			{
				Name: "Blueberry Muffin Cake",
				Time: "30 min",
				Ingredients: []shared.Ingredient{
					{Name: "Baking powder", Count: "3", Unit: "teaspoons"},
					{Name: "Brown sugar", Count: "0.5", Unit: "cup"},
					{Name: "Blueberries", Count: "1", Unit: "cup"},
				},
			},
		},
	}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %+v, want %+v", got, want)
	}
}

func TestXMLReaderInterface(t *testing.T) {
	xData := shared.XMLbytes(xmlData)
	got := shared.Reader(&xData)
	want := shared.RecipeBook{
		XMLName: xml.Name{Local: "recipes"},
		Cakes: []shared.Cake{
			{
				Name: "Red Velvet Strawberry Cake",
				Time: "40 min",
				Ingredients: []shared.Ingredient{
					{Name: "Flour", Count: "3", Unit: "cups"},
					{Name: "Vanilla extract", Count: "1.5", Unit: "tablespoons"},
					{Name: "Strawberries", Count: "7", Unit: ""},
					{Name: "Cinnamon", Count: "1", Unit: "pieces"},
				},
			},
			{
				Name: "Blueberry Muffin Cake",
				Time: "30 min",
				Ingredients: []shared.Ingredient{
					{Name: "Baking powder", Count: "3", Unit: "teaspoons"},
					{Name: "Brown sugar", Count: "0.5", Unit: "cup"},
					{Name: "Blueberries", Count: "1", Unit: "cup"},
				},
			},
		},
	}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %+v, want %+v", got, want)
	}
}
