package twitter

import (
	"github.com/dghubble/oauth1"
	"github.com/dghubble/go-twitter/twitter"
	"log"
	"github.com/jannisfink/urandomtweets/randomsources"
	"fmt"
	"github.com/jannisfink/urandomtweets/config"
)

type TwitterManager struct {
	client *twitter.Client
}

func (t *TwitterManager) Connect() {
	c := config.GetConfiguration()
	conf := oauth1.NewConfig(c.Twitter.ConsumerKey, c.Twitter.ConsumerSecret)
	token := oauth1.NewToken(c.Twitter.AccessToken, c.Twitter.AccessSecret)

	httpClient := conf.Client(oauth1.NoContext, token)
	client := twitter.NewClient(httpClient)

	t.client = client

}

func (t *TwitterManager) Tweet(info randomsources.RandomInformation) {
	var hashtags = ""

	for _, hashtag := range info.HashTags {
		hashtags = fmt.Sprintf("%s #%s", hashtags, hashtag)
	}

	message := fmt.Sprintf("%s %s %s", info.Title, info.Url, hashtags)

	_, _, err := t.client.Statuses.Update(message, &twitter.StatusUpdateParams{})
	if err != nil {
		log.Fatalf("%v", err)
	}
}
