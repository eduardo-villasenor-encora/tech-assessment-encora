package clients

import (
	"github.com/jarcoal/httpmock"
	"net/http"
	"testing"
)

func TestUserClient_GetUsers(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()
	userClient := UsersClient{http.DefaultClient, "https://fake.com"}
	httpmock.RegisterResponder("GET", "https://fake.com/users/1",
		httpmock.NewStringResponder(200, `
		{
		  "id": 1,
		  "name": "Leanne Graham",
		  "username": "Bret",
		  "email": "Sincere@april.biz",
		  "address": {
			"street": "Kulas Light",
			"suite": "Apt. 556",
			"city": "Gwenborough",
			"zipcode": "92998-3874",
			"geo": {
			  "lat": "-37.3159",
			  "lng": "81.1496"
			}
		  },
		  "phone": "1-770-736-8031 x56442",
		  "website": "hildegard.org",
		  "company": {
			"name": "Romaguera-Crona",
			"catchPhrase": "Multi-layered client-server neural-net",
			"bs": "harness real-time e-markets"
		  }
		}
	`))
	userResponse, err := userClient.GetUser(1)
	if err != nil {
		t.Fatal("unexpected error")
	}
	if userResponse.ID != 1 {
		t.Fatal("incorrect id")
	}
	if userResponse.Name != "Leanne Graham" {
		t.Fatal("incorrect id")
	}
	if userResponse.Username != "Bret" {
		t.Fatal("incorrect id")
	}
	if userResponse.Email != "Sincere@april.biz" {
		t.Fatal("incorrect id")
	}
}
