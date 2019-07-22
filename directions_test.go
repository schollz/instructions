package directions

import (
	"io/ioutil"
	"path"
	"strings"
	"testing"

	log "github.com/schollz/logger"
	"github.com/stretchr/testify/assert"
)

type URLDirections struct {
	URL        string
	Directions []string
}

var ts = []URLDirections{
	{
		"https://www.bonappetit.com/recipe/bas-best-chocolate-chip-cookies",
		[]string{},
	},
	{
		"https://pinchofyum.com/the-best-soft-chocolate-chip-cookies",
		[]string{},
	},
	{
		"https://www.modernhoney.com/the-best-chocolate-chip-cookies/",
		[]string{},
	},
	{
		"https://laurenslatest.com/actually-perfect-chocolate-chip-cookies/",
		[]string{},
	},
	{
		"https://www.allrecipes.com/recipe/10813/best-chocolate-chip-cookies/",
		[]string{},
	},
	{
		"https://joyfoodsunshine.com/the-most-amazing-chocolate-chip-cookies/",
		[]string{},
	},
	{
		"https://cakebycourtney.com/soft-chewy-chocolate-chip-cookies/",
		[]string{},
	},
	{
		"https://www.foodnetwork.com/recipes/dave-lieberman/noodle-kugel-recipe-1946564",
		[]string{},
	},
	{
		"https://cooking.nytimes.com/recipes/12320-apple-pie",
		[]string{},
	},
	{
		"https://www.bonappetit.com/recipe/bas-best-chocolate-chip-cookies",
		[]string{},
	},
}

func TestTable(t *testing.T) {
	log.SetLevel("error")
	for _, t0 := range ts {
		fileToGet := t0.URL
		fileToGet = strings.TrimPrefix(fileToGet, "https://")
		if string(fileToGet[len(fileToGet)-1]) == "/" {
			fileToGet += "index.html"
		}
		fileToGet = path.Join("testing", "sites", fileToGet)
		b, _ := ioutil.ReadFile(fileToGet)
		r, err := Parse(string(b))
		assert.Nil(t, err)
		assert.Equal(t, t0.Directions, r)
	}
}
