package clients

import (
	"github.com/jarcoal/httpmock"
	"net/http"
	"testing"
)

func TestPostClient_GetPostsForUser(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()
	postClient := PostsClient{http.DefaultClient, "https://fake.com"}
	httpmock.RegisterResponder("GET", "https://fake.com/posts?userId=1",
		httpmock.NewStringResponder(200, `
	[
		{
			"userId": 1,
			"id": 1,
			"title": "title of the user’s post",
			"body": "body of the user’s post"
		}
	]
	`))
	postResponse, err := postClient.GetPostsForUser(1)
	if err != nil {
		t.Fatal("unexpected error")
	}
	if len(postResponse) != 1 {
		t.Fatal("expect just one element on response")
	}
	if postResponse[0].ID != 1 {
		t.Fatal("incorrect bind of id")
	}
	if postResponse[0].Title != "title of the user’s post" {
		t.Fatal("incorrect bind of title")
	}
	if postResponse[0].Body != "body of the user’s post" {
		t.Fatal("incorrect bind of body")
	}
}
