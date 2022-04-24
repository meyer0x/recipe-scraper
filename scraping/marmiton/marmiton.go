package marmiton

import (
	"fmt"
	"recipe-scraper/models"
	"recipe-scraper/utils"
	"strconv"

	"github.com/gocolly/colly"
)

//InitScraping get all recipes from 0 to N pages and format.
func InitScraping(pages int) []models.Recipe {

	var recipes []models.Recipe

	urls := getURLRecipesFromUrl("https://www.marmiton.org/recettes?page=", pages)

	for i := 0; i < len(urls); i++ {
		recipe := utils.FormatRecipe(getRecipeFromUrl(urls[i]))
		recipes = append(recipes, recipe)
	}

	return recipes
}

func getRecipeFromUrl(url string) models.Recipe {

	recipe := models.Recipe{}

	c := colly.NewCollector()

	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting ", r.URL)
	})

	c.OnResponse(func(r *colly.Response) {
		fmt.Println("Status ", r.StatusCode)
	})

	c.OnHTML("h1.itJBWW", func(h *colly.HTMLElement) {
		recipe.Name = h.Text
	})

	c.OnHTML("div.ghZzUe", func(h *colly.HTMLElement) {
		h.ForEach("div.MuiGrid-item div.fLWRho", func(i int, x *colly.HTMLElement) {
			recipe.Ingredients = append(recipe.Ingredients, x.Text)
			recipe.URL = h.Request.URL.String()
		})
	})

	c.OnHTML("span.hYSrSW", func(h *colly.HTMLElement) {

		n, e := strconv.ParseInt(h.Text, 10, 64)
		if e != nil {
			panic("ServingPerson is not a number")
		}

		recipe.ServingPerson = int(n)
	})

	c.Visit(url)
	return recipe
}

func getURLRecipesFromUrl(url string, pages int) []string {

	var min, max int = 1, pages

	var recipesUrls []string

	c := colly.NewCollector()

	c.OnRequest(func(r *colly.Request) {
		fmt.Print("Visiting ", r.URL, "  ")
	})

	c.OnResponse(func(r *colly.Response) {
		fmt.Println("Status", r.StatusCode)
	})

	c.OnHTML("a.recipe-card-link", func(h *colly.HTMLElement) {
		link := h.Attr("href")
		recipesUrls = append(recipesUrls, link)
	})

	for i := min; i < max+1; i++ {
		c.Visit(url + strconv.FormatInt(int64(i), 10))
	}

	fmt.Println("Totaly Found ", len(recipesUrls), " urls")

	return recipesUrls
}
