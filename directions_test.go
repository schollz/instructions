package instructions

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
		[]string{"Place racks in upper and lower thirds of oven; preheat to 375°. Whisk flour, salt, and baking soda in a small bowl; set aside.", "Cook ½ cup (1 stick; 113 g) butter in a large saucepan over medium heat, swirling often and scraping bottom of pan with a heatproof rubber spatula, until butter foams, then browns, about 4 minutes. Transfer butter to a large heatproof bowl and let cool 1 minute. Cut remaining ¼ cup (½ stick; 56 g) butter into small pieces and add to brown butter (it should start to melt but not foam and sizzle, so test with one piece before adding the rest).", "Once butter is melted, add both sugars and whisk, breaking up any clumps, until sugar is incorporated and no lumps remain. Add egg and egg yolks and whisk until sugar dissolves and mixture is smooth, about 30 seconds. Whisk in vanilla. Using rubber spatula, fold reserved dry ingredients into butter mixture just until no dry spots remain, then fold in chocolate (the dough will be soft but should hold its shape once scooped; if it slumps or oozes after being scooped, stir dough back together several times and let rest 5–10 minutes until scoops hold their shape as the flour hydrates).", "Using a 1½-oz. scoop (3 Tbsp.), portion out 16 balls of dough and divide between 2 parchment-lined rimmed baking sheets. Bake cookies, rotating sheets if cookies are browning very unevenly (otherwise, just leave them alone), until deep golden brown and firm around the edges, 8–10 minutes. Let cool on baking sheets.", "do ahead: react-text: 204 Cookies can be made 3 days ahead. Store airtight at room temperature. /react-text"},
	},
	{
		"https://pinchofyum.com/the-best-soft-chocolate-chip-cookies",
		[]string{"Preheat the oven to 350 degrees. Microwave the butter for about 40 seconds to just barely melt it. It shouldn’t be hot – but it should be almost entirely in liquid form.", "Using a stand mixer or electric beaters, beat the butter with the sugars until creamy. Add the vanilla and the egg; beat on low speed until just incorporated – 10-15 seconds or so (if you beat the egg for too long, the cookies will be stiff).", "Add the flour, baking soda, and salt. Mix until crumbles form. Use your hands to press the crumbles together into a dough. It should form one large ball that is easy to handle (right at the stage between “wet” dough and “dry” dough). Add the chocolate chips and incorporate with your hands.", "Roll the dough into 12 large balls (or 9 for HUGELY awesome cookies) and place on a cookie sheet. Bake for 9-11 minutes until the cookies look puffy and dry and just barely golden. warning, friends: do not overbake. This advice is probably written on every cookie recipe everywhere, but this is essential for keeping the cookies soft. Take them out even if they look like they’re not done yet (see picture in the post). They’ll be pale and puffy.", "Let them cool on the pan for a good 30 minutes or so (I mean, okay, eat four or five but then let the rest of them cool). They will sink down and turn into these dense, buttery, soft cookies that are the best in all the land. These should stay soft for many days if kept in an airtight container. I also like to freeze them."},
	},
	{
		"https://www.modernhoney.com/the-best-chocolate-chip-cookies/",
		[]string{"Preheat oven to 400 degrees (if not chilling the dough). In a large mixing bowl, cream butter, brown sugar, and sugar for 4 minutes until light and fluffy.", "Add eggs and vanilla. Mix for 1 minute longer.", "Stir in flour, cornstarch, baking soda, and salt. Mix just until combined. Fold in chocolate chips.", "If time permits, wrap the dough tightly in plastic wrap and chill for 24 hours. If not, scoop cookie dough onto baking sheets. I suggest using parchment paper or Silpat silicone baking sheets on light-colored baking sheets.", "Bake for 9-11 minutes or until the edges just begin to turn a light golden color.\u00a0 Remove from oven and let set for 5 minutes before removing from the cookie sheet."},
	},
	{
		"https://laurenslatest.com/actually-perfect-chocolate-chip-cookies/",
		[]string{"Stir melted butter, brown sugar and granulated sugar until well combined.", "Add in egg, egg yolk and vanilla. Mix until lighter in color.", "Add in the flour, baking soda and salt. Scrape the sides of the bowl and the bottom really well to ensure a smooth, well stirred batter.", "Hand stir the chocolate chips into the batter and let it sit 20 minutes to let the flour soak into the rest of the batter. (The warm melted butter will help with this.)", "Preheat oven to 325. Line two baking sheets with parchment paper and scoop cookie dough onto prepared pans, using a two tablespoon cookie scoop.", "Bake 8-9 minutes, rotating sheets half way through baking. When you pull your cookies out of the oven, they will looked cooked around the edges and undercooked in the center.", "Leave the cookies on the hot baking pans for 5-7 minutes or until you can remove them without falling apart. Place onto cooling racks and cool to room temperature before storing in air tight containers."},
	},
	{
		"https://www.allrecipes.com/recipe/10813/best-chocolate-chip-cookies/",
		[]string{"Preheat oven to 350 degrees F (175 degrees C).", "Cream together the butter, white sugar, and brown sugar until smooth. Beat in the eggs one at a time, then stir in the vanilla. Dissolve baking soda in hot water.  Add to batter along with salt. Stir in flour, chocolate chips, and nuts. Drop by large spoonfuls onto ungreased pans.", "Bake for about 10 minutes in the preheated oven, or until edges are nicely browned."},
	},
	{
		"https://joyfoodsunshine.com/the-most-amazing-chocolate-chip-cookies/",
		[]string{"Preheat oven to 375 degrees F. Line a baking pan with parchment paper and set aside.", "In a separate bowl mix flour, baking soda, salt, baking powder. Set aside.", "Cream together butter and sugars until combined.", "beat\u00a0in eggs and vanilla until fluffy.", "Mix in the dry ingredients until combined.", "Add 12 oz package of chocolate chips and mix well.", "Roll 2-3 TBS (depending on how large you like your cookies) of dough at a time into balls and place them evenly spaced on your prepared cookie sheets. (alternately, use a small cookie scoop to make your cookies)!", "Bake in preheated oven for approximately 8-10\u00a0minutes. Take them out when they are just barely starting to turn brown.", "Let them sit on the baking pan for 2 minutes before removing to cooling rack."},
	},
	{
		"https://cakebycourtney.com/soft-chewy-chocolate-chip-cookies/",
		[]string{"Preheat your oven to 350 degrees F. Line baking sheets with parchment paper.", "In the bowl of an electric mixer fitted with the paddle attachment beat the butter until smooth. Add the sugars and beat until light and fluffy, about 2 minutes. Scrape down the sides of the bowl.", "Add the eggs one at time\u00a0and mix until well incorporated. Scrape down the sides of the bowl. Add the vanilla and mix for about 30 seconds.", "With the mixer on low speed, add the flour, baking soda, baking powder and salt. Mix until combined.", "stir in the chocolate chips.", "Scoop 2-inch balls of dough onto the cookie sheet. Bake for 12 minutes. The cookies are done when the edges are slightly golden brown. The top may look a little soft. That's ok! The cookies will continue to bake a little while they cool."},
	},
	{
		"https://www.foodnetwork.com/recipes/dave-lieberman/noodle-kugel-recipe-1946564",
		[]string{"preheat oven to 375 degrees f.", "Boil the noodles in salted water for about 4 minutes. Strain noodles from water. In a large mixing bowl, combine noodles with remaining ingredients and pour into a greased, approximately 9-by-13-inch baking dish.", "Bake until custard is set and top is golden brown, about 30 to 45 minutes."},
	},
	{
		"https://cooking.nytimes.com/recipes/12320-apple-pie",
		[]string{"Melt butter in a large saute pan set over medium-high heat and add apples to the pan. Stir to coat fruit with butter and cook, stirring occasionally. Meanwhile, whisk together the spices, salt and .75 cup sugar, and sprinkle this over the pan, stirring to combine. Lower heat and cook until apples have started to soften, approximately 5 to 7 minutes. Sprinkle the flour and cornstarch over the apples and continue to cook, stirring occasionally, another 3 to 5 minutes. Remove pan from heat, add cider vinegar, stir and scrape fruit mixture into a bowl and allow to cool completely. (The fruit mixture will cool faster if spread out on a rimmed baking sheet.)", "Place a large baking sheet on the middle rack of oven and preheat to 425. Remove one disc of dough from the refrigerator and, using a pin, roll it out on a lightly floured surface until it is roughly 12 inches in diameter. Fit this crust into a 9-inch pie plate, trimming it to leave a .5-inch overhang. Place this plate, with the dough, in the freezer.", "Roll out the remaining dough on a lightly floured surface until it is roughly 10 or 11 inches in diameter.", "Remove pie crust from freezer and put the cooled pie filling into it. Cover with remaining dough. Press the edges together, trim the excess, then crimp the edges with the tines of a fork. Using a sharp knife, cut three or four steam vents in the top of the crust. Lightly brush the top of the pie with egg wash and sprinkle with remaining tablespoon of sugar.", "Place pie in oven and bake on hot baking sheet for 20 minutes, then reduce temperature to 375. Continue to cook until the interior is bubbling and the crust is golden brown, about 30 to 40 minutes more. Remove and allow to cool on a windowsill or kitchen rack, about two hours."},
	},
	{
		"https://www.bonappetit.com/recipe/bas-best-chocolate-chip-cookies",
		[]string{"Place racks in upper and lower thirds of oven; preheat to 375°. Whisk flour, salt, and baking soda in a small bowl; set aside.", "Cook ½ cup (1 stick; 113 g) butter in a large saucepan over medium heat, swirling often and scraping bottom of pan with a heatproof rubber spatula, until butter foams, then browns, about 4 minutes. Transfer butter to a large heatproof bowl and let cool 1 minute. Cut remaining ¼ cup (½ stick; 56 g) butter into small pieces and add to brown butter (it should start to melt but not foam and sizzle, so test with one piece before adding the rest).", "Once butter is melted, add both sugars and whisk, breaking up any clumps, until sugar is incorporated and no lumps remain. Add egg and egg yolks and whisk until sugar dissolves and mixture is smooth, about 30 seconds. Whisk in vanilla. Using rubber spatula, fold reserved dry ingredients into butter mixture just until no dry spots remain, then fold in chocolate (the dough will be soft but should hold its shape once scooped; if it slumps or oozes after being scooped, stir dough back together several times and let rest 5–10 minutes until scoops hold their shape as the flour hydrates).", "Using a 1½-oz. scoop (3 Tbsp.), portion out 16 balls of dough and divide between 2 parchment-lined rimmed baking sheets. Bake cookies, rotating sheets if cookies are browning very unevenly (otherwise, just leave them alone), until deep golden brown and firm around the edges, 8–10 minutes. Let cool on baking sheets.", "do ahead: react-text: 204 Cookies can be made 3 days ahead. Store airtight at room temperature. /react-text"},
	},
	{
		"https://www.ricardocuisine.com/en/recipes/4874-chewy-chocolate-chip-cookies-the-best",
		[]string{"With the rack in the middle position, preheat the oven to 375°F (190°C). Line two or three baking sheets with parchment paper.", "In a bowl, combine the flour, baking soda and salt.", "In another bowl, combine the butter and both sugars with a wooden spoon. Add the eggs and stir until smooth. Stir in the dry ingredients and chocolate. Cover and refrigerate for 1 hour or overnight.", "Using a 3-tbsp (45 ml) ice cream scoop, spoon five to six balls of dough on each baking sheet, spacing them out evenly.", "Bake one sheet at a time for 8 to 9 minutes or until lightly browned all over. They will still be very soft in the centre. Cool completely on the baking sheet."},
	},
}

func TestTable(t *testing.T) {
	log.SetLevel("info")
	for i, t0 := range ts {
		if i != len(ts)-1 {
			//continue
		}
		log.Info(i, t0.URL)
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
