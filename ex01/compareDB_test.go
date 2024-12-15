package main

import (
	"bytes"
	"comparing/shared"
	"encoding/xml"
	"testing"
)

func TestCompareIngredients(t *testing.T) {
	t.Run("test for adding ingredients", func(t *testing.T) {
		buffer := &bytes.Buffer{}
		oldRecipeBook := shared.RecipeBook{
			Cakes: []shared.Cake{
				{
					Name: "Red Velvet Strawberry Cake",
					Time: "45 min",
					Ingredients: []shared.Ingredient{
						{"Flour", "2", "cups"},
						{"Strawberries", "8", ""},
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

		newRecipeBook := shared.RecipeBook{
			Cakes: []shared.Cake{
				{
					Name: "Red Velvet Strawberry Cake",
					Time: "45 min",
					Ingredients: []shared.Ingredient{
						{"Flour", "2", "cups"},
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

		CompareIngredients(buffer, oldRecipeBook.Cakes[0], newRecipeBook.Cakes[0])

		got := buffer.String()
		want := "ADDED ingredient \"Coffee Beans\" for cake \"Red Velvet Strawberry Cake\"\n"

		if got != want {
			t.Errorf("\ngot %s \nwant %s", got, want)
		}
	})
	t.Run("test for adding ingredients", func(t *testing.T) {
		buffer := &bytes.Buffer{}
		oldRecipeBook := shared.RecipeBook{
			Cakes: []shared.Cake{
				{
					Name: "Red Velvet Strawberry Cake",
					Time: "45 min",
					Ingredients: []shared.Ingredient{
						{"Flour", "2", "cups"},
						{"Strawberries", "8", ""},
						{"Coffee Beans", "2.5", "tablespoons"},
						{"Cinnamon", "1", ""},
						{"Vanilla extract", "2", "tablespoon"},
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

		newRecipeBook := shared.RecipeBook{
			Cakes: []shared.Cake{
				{
					Name: "Red Velvet Strawberry Cake",
					Time: "45 min",
					Ingredients: []shared.Ingredient{
						{"Flour", "2", "cups"},
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

		CompareIngredients(buffer, oldRecipeBook.Cakes[0], newRecipeBook.Cakes[0])

		got := buffer.String()
		want := "REMOVED ingredient \"Vanilla extract\" for cake \"Red Velvet Strawberry Cake\"\n"

		if got != want {
			t.Errorf("\ngot %s \nwant %s", got, want)
		}
	})

	t.Run("test for changing unit for ingredient", func(t *testing.T) {
		buffer := &bytes.Buffer{}
		oldRecipeBook := shared.RecipeBook{
			Cakes: []shared.Cake{
				{
					Name: "Red Velvet Strawberry Cake",
					Time: "45 min",
					Ingredients: []shared.Ingredient{
						{"Flour", "2", "cups"},
						{"Strawberries", "8", ""},
						{"Coffee Beans", "2.5", "tablespoons"},
						{"Cinnamon", "1", ""},
						{"Vanilla extract", "2", "tablespoon"},
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

		newRecipeBook := shared.RecipeBook{
			Cakes: []shared.Cake{
				{
					Name: "Red Velvet Strawberry Cake",
					Time: "45 min",
					Ingredients: []shared.Ingredient{
						{"Flour", "2", "mugs"},
						{"Strawberries", "8", ""},
						{"Coffee Beans", "2.5", "tablespoons"},
						{"Cinnamon", "1", ""},
						{"Vanilla extract", "2", "tablespoon"},
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

		CompareIngredients(buffer, oldRecipeBook.Cakes[0], newRecipeBook.Cakes[0])

		got := buffer.String()
		want := "CHANGED unit for ingredient \"Flour\" for cake \"Red Velvet Strawberry Cake\" - \"mugs\" instead of \"cups\"\n"

		if got != want {
			t.Errorf("\ngot %s \nwant %s", got, want)
		}
	})

	t.Run("test for changing unit count for ingredient", func(t *testing.T) {
		buffer := &bytes.Buffer{}
		oldRecipeBook := shared.RecipeBook{
			Cakes: []shared.Cake{
				{
					Name: "Red Velvet Strawberry Cake",
					Time: "45 min",
					Ingredients: []shared.Ingredient{
						{"Flour", "2", "cups"},
						{"Strawberries", "7", ""},
						{"Coffee Beans", "2.5", "tablespoons"},
						{"Cinnamon", "1", ""},
						{"Vanilla extract", "2", "tablespoon"},
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

		newRecipeBook := shared.RecipeBook{
			Cakes: []shared.Cake{
				{
					Name: "Red Velvet Strawberry Cake",
					Time: "45 min",
					Ingredients: []shared.Ingredient{
						{"Flour", "2", "cups"},
						{"Strawberries", "8", ""},
						{"Coffee Beans", "2.5", "tablespoons"},
						{"Cinnamon", "1", ""},
						{"Vanilla extract", "2", "tablespoon"},
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

		CompareIngredients(buffer, oldRecipeBook.Cakes[0], newRecipeBook.Cakes[0])

		got := buffer.String()
		want := "CHANGED unit count for ingredient \"Strawberries\" for cake \"Red Velvet Strawberry Cake\" - \"8\" instead of \"7\"\n"

		if got != want {
			t.Errorf("\ngot %s \nwant %s", got, want)
		}
	})

	t.Run("test for removed unit for ingredient", func(t *testing.T) {
		buffer := &bytes.Buffer{}
		oldRecipeBook := shared.RecipeBook{
			Cakes: []shared.Cake{
				{
					Name: "Red Velvet Strawberry Cake",
					Time: "45 min",
					Ingredients: []shared.Ingredient{
						{"Flour", "2", "cups"},
						{"Strawberries", "8", ""},
						{"Coffee Beans", "2.5", "tablespoons"},
						{"Cinnamon", "1", "pieces"},
						{"Vanilla extract", "2", "tablespoon"},
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

		newRecipeBook := shared.RecipeBook{
			Cakes: []shared.Cake{
				{
					Name: "Red Velvet Strawberry Cake",
					Time: "45 min",
					Ingredients: []shared.Ingredient{
						{"Flour", "2", "cups"},
						{"Strawberries", "8", ""},
						{"Coffee Beans", "2.5", "tablespoons"},
						{"Cinnamon", "1", ""},
						{"Vanilla extract", "2", "tablespoon"},
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

		CompareIngredients(buffer, oldRecipeBook.Cakes[0], newRecipeBook.Cakes[0])

		got := buffer.String()
		want := "REMOVED unit \"pieces\" for ingredient \"Cinnamon\" for cake \"Red Velvet Strawberry Cake\"\n"

		if got != want {
			t.Errorf("\ngot %s \nwant %s", got, want)
		}
	})

	t.Run("test for added unit for ingredient", func(t *testing.T) {
		buffer := &bytes.Buffer{}
		oldRecipeBook := shared.RecipeBook{
			Cakes: []shared.Cake{
				{
					Name: "Red Velvet Strawberry Cake",
					Time: "45 min",
					Ingredients: []shared.Ingredient{
						{"Flour", "2", "cups"},
						{"Strawberries", "8", ""},
						{"Coffee Beans", "2.5", "tablespoons"},
						{"Cinnamon", "1", ""},
						{"Vanilla extract", "2", "tablespoon"},
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

		newRecipeBook := shared.RecipeBook{
			Cakes: []shared.Cake{
				{
					Name: "Red Velvet Strawberry Cake",
					Time: "45 min",
					Ingredients: []shared.Ingredient{
						{"Flour", "2", "cups"},
						{"Strawberries", "8", ""},
						{"Coffee Beans", "2.5", "tablespoons"},
						{"Cinnamon", "1", "pieces"},
						{"Vanilla extract", "2", "tablespoon"},
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

		CompareIngredients(buffer, oldRecipeBook.Cakes[0], newRecipeBook.Cakes[0])

		got := buffer.String()
		want := "ADDED unit \"pieces\" for ingredient \"Cinnamon\" for cake \"Red Velvet Strawberry Cake\"\n"

		if got != want {
			t.Errorf("\ngot %s \nwant %s", got, want)
		}
	})
}

func TestCompareCakes(t *testing.T) {

	t.Run("test for added cake", func(t *testing.T) {
		buffer := &bytes.Buffer{}
		oldRecipeBook := shared.RecipeBook{
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

		newRecipeBook := shared.RecipeBook{
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

		CompareCakes(buffer, oldRecipeBook, newRecipeBook)

		got := buffer.String()
		want := "ADDED cake \"Moonshine Muffin\"\n"

		if got != want {
			t.Errorf("\ngot %s \nwant %s", got, want)
		}
	})

	t.Run("test for removed cake", func(t *testing.T) {
		buffer := &bytes.Buffer{}
		oldRecipeBook := shared.RecipeBook{
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

		newRecipeBook := shared.RecipeBook{
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
			},
		}

		CompareCakes(buffer, oldRecipeBook, newRecipeBook)

		got := buffer.String()
		want := "REMOVED cake \"Blueberry Muffin Cake\"\n"

		if got != want {
			t.Errorf("\ngot %s \nwant %s", got, want)
		}
	})

	t.Run("test for removed cake", func(t *testing.T) {
		buffer := &bytes.Buffer{}
		oldRecipeBook := shared.RecipeBook{
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
			},
		}

		newRecipeBook := shared.RecipeBook{
			XMLName: xml.Name{Local: "recipes"},
			Cakes: []shared.Cake{
				{
					Name: "Red Velvet Strawberry Cake",
					Time: "45 min",
					Ingredients: []shared.Ingredient{
						{Name: "Flour", Count: "3", Unit: "cups"},
						{Name: "Vanilla extract", Count: "1.5", Unit: "tablespoons"},
						{Name: "Strawberries", Count: "7", Unit: ""},
						{Name: "Cinnamon", Count: "1", Unit: "pieces"},
					},
				},
			},
		}

		CompareCakes(buffer, oldRecipeBook, newRecipeBook)

		got := buffer.String()
		want := "CHANGED cooking time for cake \"Red Velvet Strawberry Cake\" - \"45 min\" instead of \"40 min\"\n"

		if got != want {
			t.Errorf("\ngot %s \nwant %s", got, want)
		}
	})
}

func TestAllTheCompareFunctions(t *testing.T) {
	buffer := &bytes.Buffer{}
	oldRecipeBook := shared.RecipeBook{
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

	newRecipeBook := shared.RecipeBook{
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

	CompareCakes(buffer, oldRecipeBook, newRecipeBook)

	got := buffer.String()
	want := `ADDED cake "Moonshine Muffin"
REMOVED cake "Blueberry Muffin Cake"
CHANGED cooking time for cake "Red Velvet Strawberry Cake" - "45 min" instead of "40 min"
ADDED ingredient "Coffee Beans" for cake "Red Velvet Strawberry Cake"
REMOVED ingredient "Vanilla extract" for cake "Red Velvet Strawberry Cake"
CHANGED unit for ingredient "Flour" for cake "Red Velvet Strawberry Cake" - "mugs" instead of "cups"
CHANGED unit count for ingredient "Flour" for cake "Red Velvet Strawberry Cake" - "2" instead of "3"
CHANGED unit count for ingredient "Strawberries" for cake "Red Velvet Strawberry Cake" - "8" instead of "7"
REMOVED unit "pieces" for ingredient "Cinnamon" for cake "Red Velvet Strawberry Cake"
`

	if got != want {
		t.Errorf("\ngot %s \nwant %s", got, want)
	}
}
