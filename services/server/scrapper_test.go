package server

import (
	"testing"

	"github.com/kirill-a-belov/test-sg/models"
)

var testData = []*models.Product{
	{
		URL:   "https://www.amazon.co.uk/gp/product/1509836071",
		Title: "The Fat-Loss Plan: 100 Quick and Easy Recipes with Workouts",
		Price: "8.49",
		Image: "https://images-na.ssl-images-amazon.com/images/I/71FL39V4zCL.jpg",
	},
	{
		URL:   "https://www.amazon.co.uk/dp/0008212244",
		Title: "Guilt: The Sunday Times best selling psychological thriller that you need to read in 2018",
		Price: "3.75",
		Image: "https://images-na.ssl-images-amazon.com/images/I/51Me%2BsYQ2gL.jpg",
	},
}

func TestScrapURL(t *testing.T) {
	for _, td := range testData {
		res, err := ScrapURL(td.URL)
		if err != nil {
			t.Error(
				"For input: ", td,
				"got error", err,
			)
			continue
		}

		if res.Title != td.Title ||
			res.Price != td.Price ||
			res.Image != td.Image ||
			res.URL != td.URL {

			t.Error(
				"For input: ", td,
				"got not equal", res,
			)
		}
	}

}
