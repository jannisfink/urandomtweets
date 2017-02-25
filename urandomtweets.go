package main

import (
	"fmt"
	"github.com/jannisfink/urandomtweets/randomsources/wikipedia"
	"math/rand"
	"time"
	"github.com/jannisfink/urandomtweets/randomsources"
	"github.com/jannisfink/urandomtweets/twitter"
	"github.com/jannisfink/urandomtweets/config"
)

var randomWordGenerators = []func() randomsources.RandomInformation{
	wikipedia.SelectRandomWikipediaArticle,
}

func main() {
	config.LoadConfiguration()

	rand.Seed(time.Now().UnixNano())

	randomFunc := randomWordGenerators[rand.Intn(len(randomWordGenerators))]

	r := randomFunc()

	fmt.Println(r.Title)
	fmt.Println(r.Url)

	twitterManager := twitter.TwitterManager{}
	twitterManager.Connect()

	twitterManager.Tweet(r)
}
