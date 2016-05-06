package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"runtime"

	"github.com/ChimeraCoder/anaconda"
)

func hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "hello, world from %s", runtime.Version())
}

func helloWorld(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello, this is a simple twitter app written in Go :)")
}

func bindListenServe() {
	bind := fmt.Sprintf("0.0.0.0:8080")

	err := http.ListenAndServe(bind, nil)
	if err != nil {
		log.Fatal(err)
	}
}

func getTweet(w http.ResponseWriter, r *http.Request) {
	anaconda.SetConsumerKey(os.Getenv("CONSUMER_KEY"))
	anaconda.SetConsumerSecret(os.Getenv("CONSUMER_SECRET"))
	api := anaconda.NewTwitterApi(os.Getenv("ACCESS_TOKEN"), os.Getenv("ACCESS_TOKEN_SECRET"))

	searchResult, err := api.GetSearch("#"+r.URL.Path[7:], nil)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Fprintf(w, "Recent tweets that contain #%s:\n", r.URL.Path[7:])
	for _, tweet := range searchResult.Statuses {
		if tweet.Text != "" {
			//fmt.Fprint(w, tweet.User.Name+" @"+tweet.User.ScreenName+"\n")
			fmt.Fprint(w, tweet.Text+"\n")
		}
	}
}

func main() {
	http.HandleFunc("/hello", helloWorld)
	http.HandleFunc("/", hello)
	http.HandleFunc("/tweet/", getTweet)
	bindListenServe()
}
