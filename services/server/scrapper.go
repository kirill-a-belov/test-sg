package server

import (
	"errors"
	"fmt"
	"net/http"
	"regexp"
	"strings"

	"github.com/PuerkitoBio/goquery"
	"github.com/kirill-a-belov/test-sg/models"
	uuid2 "github.com/nu7hatch/gouuid"
	"golang.org/x/net/html"
)

var priceRegExp = regexp.MustCompile(`[0-9.]+`)
var imgUrlRegExp = regexp.MustCompile(`(?mU)https://.*.jpg`)

// To small to create new package
func ScrapURL(URL string) (*models.Product, error) {
	res, err := http.Get(URL)
	if err != nil {
		return nil, fmt.Errorf("error from ScrapURL: %v", err)
	}
	defer res.Body.Close()
	if res.Request == nil {
		return nil, errors.New("error from ScrapURL:nil request")
	}

	// Parse the HTML into nodes
	root, err := html.Parse(res.Body)
	if err != nil {
		return nil, fmt.Errorf("error from ScrapURL: %v", err)
	}

	page := goquery.NewDocumentFromNode(root)
	title := page.Find("#productTitle").Text()
	price := page.Find(".a-size-medium.a-color-price.offer-price.a-text-normal").Contents().Text()
	price = priceRegExp.FindString(price)
	image := page.Find("#leftCol").Contents().Text()
	image = imgUrlRegExp.FindString(image)
	inStock := page.Find("#availability").Text()

	uuid, err := uuid2.NewV4()

	return &models.Product{
		PID:       uuid.String(),
		URL:       URL,
		Title:     title,
		Price:     price,
		Image:     image,
		IsInStock: strings.Contains(inStock, "In stock"),
	}, nil
}
