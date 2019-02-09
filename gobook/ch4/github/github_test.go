package github

import (
	"fmt"
	"log"
	"testing"
)

func TestSearchIssues(t *testing.T) {
	terms := []string{"repo:mattermost/mattermost-server", "is:Open", "label:Go"}
	result, err := SearchIssues(terms)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%d issues:\n", result.TotalCount)
	for _, item := range result.Items {
		fmt.Printf("#%-5d %10.8s %.120s\n",
			item.Number, item.User.Login, item.Title)
	}
}
