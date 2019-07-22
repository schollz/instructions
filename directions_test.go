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
		[]string{"Place racks in upper and lower thirds of oven; preheat to 375°. Whisk flour, salt, and baking soda in a small bowl; set aside.", "Cook ½ cup (1 stick; 113 g) butter in a large saucepan over medium heat, swirling often and scraping bottom of pan with a heatproof rubber spatula, until butter foams, then browns, about 4 minutes. Transfer butter to a large heatproof bowl and let cool 1 minute. Cut remaining ¼ cup (½ stick; 56 g) butter into small pieces and add to brown butter (it should start to melt but not foam and sizzle, so test with one piece before adding the rest).", "Once butter is melted, add both sugars and whisk, breaking up any clumps, until sugar is incorporated and no lumps remain. Add egg and egg yolks and whisk until sugar dissolves and mixture is smooth, about 30 seconds. Whisk in vanilla. Using rubber spatula, fold reserved dry ingredients into butter mixture just until no dry spots remain, then fold in chocolate (the dough will be soft but should hold its shape once scooped; if it slumps or oozes after being scooped, stir dough back together several times and let rest 5–10 minutes until scoops hold their shape as the flour hydrates).", "Using a 1½-oz. scoop (3 Tbsp.), portion out 16 balls of dough and divide between 2 parchment-lined rimmed baking sheets. Bake cookies, rotating sheets if cookies are browning very unevenly (otherwise, just leave them alone), until deep golden brown and firm around the edges, 8–10 minutes. Let cool on baking sheets.", "Do Ahead: react-text: 204 Cookies can be made 3 days ahead. Store airtight at room temperature. /react-text"},
	},
	{
		"https://pinchofyum.com/the-best-soft-chocolate-chip-cookies",
		[]string{"Preheat the oven to 350 degrees. Microwave the butter for about 40 seconds to just barely melt it. It shouldn’t be hot – but it should be almost entirely in liquid form.", "Using a stand mixer or electric beaters, beat the butter with the sugars until creamy. Add the vanilla and the egg; beat on low speed until just incorporated – 10-15 seconds or so (if you beat the egg for too long, the cookies will be stiff).", "Add the flour, baking soda, and salt. Mix until crumbles form. Use your hands to press the crumbles together into a dough. It should form one large ball that is easy to handle (right at the stage between “wet” dough and “dry” dough). Add the chocolate chips and incorporate with your hands.", "Roll the dough into 12 large balls (or 9 for HUGELY awesome cookies) and place on a cookie sheet. Bake for 9-11 minutes until the cookies look puffy and dry and just barely golden. Warning, friends: DO NOT OVERBAKE. This advice is probably written on every cookie recipe everywhere, but this is essential for keeping the cookies soft. Take them out even if they look like they’re not done yet (see picture in the post). They’ll be pale and puffy.", "Let them cool on the pan for a good 30 minutes or so (I mean, okay, eat four or five but then let the rest of them cool). They will sink down and turn into these dense, buttery, soft cookies that are the best in all the land. These should stay soft for many days if kept in an airtight container. I also like to freeze them."},
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
	log.SetLevel("info")
	for _, t0 := range ts[:3] {
		fileToGet := t0.URL
		fileToGet = strings.TrimPrefix(fileToGet, "https://")
		if string(fileToGet[len(fileToGet)-1]) == "/" {
			fileToGet += "index.html"
		}
		fileToGet = path.Join("testing", "sites", fileToGet)
		log.Info(fileToGet)
		b, _ := ioutil.ReadFile(fileToGet)
		r, err := Parse(string(b))
		assert.Nil(t, err)
		assert.Equal(t, t0.Directions, r)
	}
}
