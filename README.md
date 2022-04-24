# recipe-scraper
Recipe Scraper is a library allowing to scrape the recipes of big site such as marmiton, tasty and to feed a database.


## Installation

Use golang to install recipe-scraper.

```bash
brew install go
```
```bash
go mod download
```
## Usage

```go
package main

import (
	"recipe-scraper/models"
	"recipe-scraper/scraping/marmiton"
)

func main() {
    // insert in MongoDB all recipe (no duplicate) from marmiton and scrape N pages (3).
	models.InsertManyRecipes(marmiton.InitScraping(3)) 
}

```

## Contributing
Pull requests are welcome. For major changes, please open an issue first to discuss what you would like to change.

## Scraping is illegal, Please, for education purpose only

## License
[MIT](https://choosealicense.com/licenses/mit/)
