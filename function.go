package fugoblog

import (
	"log"
	"net/http"
)

const tweetText = "フゴブロ更新！？"

func Handler(_ http.ResponseWriter, _ *http.Request) {
	r, err := CheckIfRSSUpdated()
	if err != nil {
		log.Printf(`{ "message": "%v", "severity": "error" }`, err)
	}
	if !r {
		return
	}

	res, err := Tweet(tweetText)

	if err != nil {
		log.Printf(`{ "message": "%v", "severity": "error" }`, err)
		return
	}

	log.Printf(`{ "message": "Successfully tweeted: %s", "severity": "info" }`, res.Text)
}
