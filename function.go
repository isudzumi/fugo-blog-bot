package fugoblog

import (
	"log"
	"net/http"
)

const tweetText = "フゴブロ更新！？"

func Handler(_ http.ResponseWriter, _ *http.Request) {
	r, err := CheckIfRSSUpdated()
	if err != nil {
		log.Fatalln(err)
	}
	if !r {
		return
	}

	res, err := Tweet(tweetText)

	if err != nil {
		log.Fatalf("Failed to tweet: %v\n", err)
		return
	}

	log.Println(res.Text)
}
