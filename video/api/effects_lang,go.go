package api

type Effect struct {
	title string
	url   string
}

var effects = map[string][]Effect{
	"nl_NL": {
		{title: "Wave", url: "wave"},
		{title: "Sparkle", url: "sparkle"},
		{title: "Mond", url: "mond"},
		{title: "Fill", url: "fill"},
		{title: "Diamond", url: "diamond"},
		{title: "Circle", url: "circle"},
		{title: "Bars", url: "bars"},
		{title: "Clear", url: "clear"},
		{title: "Snake", url: "snake"},
		{title: "Clock", url: "clock"},
	},
	"af_ZA": {
		{title: "Golf", url: "wave"},
		{title: "Vonkel", url: "sparkle"},
		{title: "Bek", url: "mond"},
		{title: "Vullen", url: "fill"},
		{title: "Diamant", url: "diamond"},
		{title: "Sirkel", url: "circle"},
		{title: "Stafies", url: "bars"},
		{title: "Skoon", url: "clear"},
		{title: "Slang", url: "snake"},
		{title: "Horlosie", url: "clock"},
	},
}
