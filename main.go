package main

import "log"

const tweetText = "フゴブロ更新！？"

func main() {
	r, err := CheckIfRSSUpdated()
	if err != nil {
		log.Fatalln(err)
	}
	if !r {
		return
	}
	Tweet(tweetText)
}
