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
}

func TestTable(t *testing.T) {
	log.SetLevel("trace")
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
