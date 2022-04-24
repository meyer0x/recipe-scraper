package main

import (
	"recipe-scraper/models"
	"recipe-scraper/scraping/marmiton"
)

func main() {
	models.InsertManyRecipes(marmiton.InitScraping(3)) // insert in MongoDB all recipe (no duplicate) from marmiton and scrape N pages (3)
}
