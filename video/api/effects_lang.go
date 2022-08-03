package api

type Effect struct {
	Title string `json:"title"`
	Url   string `json:"url"`
}

var effects = map[string][]Effect{
	"nl_NL": {
		{Title: "Wave", Url: "wave"},
		{Title: "Sparkle", Url: "sparkle"},
		{Title: "Mond", Url: "mond"},
		{Title: "Fill", Url: "fill"},
		{Title: "Diamond", Url: "diamond"},
		{Title: "Circle", Url: "circle"},
		{Title: "Bars", Url: "bars"},
		{Title: "Clear", Url: "clear"},
		{Title: "Snake", Url: "snake"},
		{Title: "Clock", Url: "clock"},
	},
	"af_ZA": {
		{Title: "Golf", Url: "wave"},
		{Title: "Vonkel", Url: "sparkle"},
		{Title: "Bek", Url: "mond"},
		{Title: "Vullen", Url: "fill"},
		{Title: "Diamant", Url: "diamond"},
		{Title: "Sirkel", Url: "circle"},
		{Title: "Stafies", Url: "bars"},
		{Title: "Skoon", Url: "clear"},
		{Title: "Slang", Url: "snake"},
		{Title: "Horlosie", Url: "clock"},
	},
}

var activations = map[string][]Effect{
	"nl_NL": {
		{Title: "Linear", Url: "linear"},
		{Title: "Sine", Url: "sine"},
		{Title: "Smoothstep", Url: "smoothstep"},
		{Title: "Smootherstep", Url: "smootherstep"},
		{Title: "Quadratic", Url: "quadratic"},
		{Title: "Cubic", Url: "cubic"},
		{Title: "TruncatedLinear", Url: "truncatedlinear"},
		{Title: "MoreTruncatedLinear", Url: "moretruncatedlinear"},
	},
}
