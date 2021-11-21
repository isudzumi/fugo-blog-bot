package fugoblog

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"

	"github.com/dghubble/oauth1"
	"github.com/dghubble/sling"
)

const twitterAPIv2 = "https://api.twitter.com/2/"

var (
	consumerKey       = os.Getenv("TWITTER_OAUTH_CONSUMER_KEY")
	consumerSecret    = os.Getenv("TWITTER_OAUTH_CONSUMER_SECRET")
	accessToken       = os.Getenv("TWITTER_ACCESS_TOKEN")
	accessTokenSecret = os.Getenv("TWITTER_ACCESS_TOKEN_SECRET")
)

type TweetsService struct {
	sling *sling.Sling
}

type TweetParam struct {
	Text string `json:"text"`
}

type TweetData struct {
	Id   string `json:"id"`
	Text string `json:"text"`
}

type TweetResponse struct {
	Data *TweetData `json:"data"`
}

type APIError struct {
	Title  string `json:"title"`
	Type   string `json:"type"`
	Status int    `json:"status"`
	Detail string `json:"detail"`
}

func (e APIError) Error() string {
	if e.Status > 400 {
		return fmt.Sprintf("twitter: %d %q %s", e.Status, e.Title, e.Detail)
	}
	return ""
}

func newTweetsService(sling *sling.Sling) *TweetsService {
	return &TweetsService{
		sling: sling.Path("tweets"),
	}
}

func (s *TweetsService) Tweet(text string) (*TweetResponse, *http.Response, error) {
	params := &TweetParam{
		Text: text,
	}
	tweet := new(TweetResponse)
	apiError := new(APIError)
	resp, err := s.sling.New().Post("").BodyJSON(params).Receive(tweet, apiError)
	if err != nil {
		return nil, nil, err
	}
	if apiError.Status > 400 {
		return nil, nil, apiError
	}
	return tweet, resp, nil
}

type Client struct {
	sling  *sling.Sling
	Tweets *TweetsService
}

func NewTwitterClient(httpClient *http.Client) *Client {
	base := sling.New().Client(httpClient).Base(twitterAPIv2)
	return &Client{
		sling:  base,
		Tweets: newTweetsService(base.New()),
	}
}

func Tweet(message string) (*TweetResponse, error) {
	config := oauth1.NewConfig(consumerKey, consumerSecret)
	token := oauth1.NewToken(accessToken, accessTokenSecret)
	httpClient := config.Client(oauth1.NoContext, token)
	client := NewTwitterClient(httpClient)

	tweet, resp, err := client.Tweets.Tweet(message)

	if err != nil {
		return nil, err
	}

	defer func() {
		io.Copy(ioutil.Discard, resp.Body)
		resp.Body.Close()
	}()

	return tweet, nil
}
