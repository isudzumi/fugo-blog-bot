package main

import (
	"fmt"
	"log"
	"os"

	"github.com/dghubble/go-twitter/twitter"
	"github.com/dghubble/oauth1"
)

var (
    consumerKey  = os.Getenv("TWITTER_OAUTH_CONSUMER_KEY")
    consumerSecret = os.Getenv("TWITTER_OAUTH_CONSUMER_SECRET")
    accessToken = os.Getenv("TWITTER_ACCESS_TOKEN")
    accessTokenSecret = os.Getenv("TWITTER_ACCESS_TOKEN_SECRET")
)

func tweet(message string) {
    config := oauth1.NewConfig(consumerKey, consumerSecret)
    token := oauth1.NewToken(accessToken, accessTokenSecret)
    httpClient := config.Client(oauth1.NoContext, token)
    client := twitter.NewClient(httpClient)

    tweet, resp, err := client.Statuses.Update(message, nil)

    if err != nil {
        log.Fatal(err)
    }

    defer resp.Body.Close()

    fmt.Println(tweet.Text)
}

