package wikipedia

import (
	"github.com/jannisfink/urandomtweets/randomsources"
	"io/ioutil"
	"log"
	"net/http"
	"regexp"
)

const wikipediaRandomPageRedirect = "https://en.wikipedia.org/wiki/Special:Random"

var wikipediaHashtags = []string{
	"wikipedia",
	"random",
}

func SelectRandomWikipediaArticle() randomsources.RandomInformation {
	resp, err := http.Get(wikipediaRandomPageRedirect)
	if err != nil {
		log.Fatal("error during http get")
	}
	body := resp.Body
	readBody, err := ioutil.ReadAll(body)
	if err != nil {
		log.Fatal("error during body read")
	}

	stringBody := string(readBody)
	titleRegex := regexp.MustCompile("<title>(.*) - Wikipedia</title>")

	temp := titleRegex.FindString(stringBody)
	title := titleRegex.ReplaceAllString(temp, `$1`)
	return randomsources.RandomInformation{
		Title:    title,
		Url:      resp.Request.URL.String(),
		HashTags: wikipediaHashtags,
	}
}
