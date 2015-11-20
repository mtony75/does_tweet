package main

import (
	"fmt"
	"bufio"
	"os"
)

const TWEET_LIMIT = 140

func valid_tweet(user_tweet string){
	if len(user_tweet) <= TWEET_LIMIT {
		fmt.Println("This is a valid Tweet")
	} else {
		fmt.Printf("This is not a valid Tweet. You are %d characters over the Twitter limit.", len(user_tweet)-TWEET_LIMIT)
}
}

func main() {
	
	fmt.Println("Type in a Tweet")

	in := bufio.NewReader(os.Stdin)
	tweet, _ := in.ReadString('\n')
	fmt.Printf("%s", tweet)
	valid_tweet(tweet)
	
	
	
}