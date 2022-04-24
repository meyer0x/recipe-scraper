package utils

import (
	"recipe-scraper/models"
	"strings"
)

func FormatRecipe(recipe models.Recipe) models.Recipe {
	return models.Recipe{
		Name:          formatName(recipe.Name),
		ServingPerson: recipe.ServingPerson,
		Ingredients:   formatIngredients(recipe.Ingredients),
		URL:           formatURL(recipe.URL),
	}
}

func formatName(name string) string {
	return strings.TrimSpace(name)
}

func formatURL(url string) string {
	return url
}

func formatIngredients(ingredients []string) []string {
	for i := 0; i < len(ingredients); i++ {
		ingredients[i] = strings.TrimSpace(ingredients[i])
	}

	return ingredients
}
