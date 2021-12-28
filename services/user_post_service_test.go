package services

import (
	"tech-assessment/clients"
	"testing"
)

func TestUserService_GetUserPosts(t *testing.T) {
	userService := UserService{
		UserClient: MockUserClient{},
		PostClient: MockPostClient{},
	}
	response, err := userService.GetUserPosts(1)
	if err != nil {
		t.Fatal("unexpected error")
	}
	if response.ID < 1 {
		t.Fatal("id should be greater than 1")
	}
	if len(response.Posts) != 2 {
		t.Fatal("expected 2 posts")
	}
	if response.UserInfo.Name != "Eduardo" {
		t.Fatal("expected userInfo.name  is Eduardo")
	}
	if response.UserInfo.Username != "edu.marquez" {
		t.Fatal("expected userInfo.name  is edu.marquez")
	}
	if response.UserInfo.Email != "eduardo.villasenor@fake.com" {
		t.Fatal("expected userInfo.Email is ")
	}
}

type MockUserClient struct {
}

func (m MockUserClient) GetUser(id int) (clients.UserClientResponse, error) {
	return clients.UserClientResponse{
		ID:       1,
		Name:     "Eduardo",
		Username: "edu.marquez",
		Email:    "eduardo.villasenor@fake.com",
		Address: clients.AddressClientResponse{
			Street:  "fake street",
			Suite:   "fake suite",
			City:    "city",
			Zipcode: "5490",
			Geo: clients.GeoClientResponse{
				Latitude:  "39.0",
				Longitude: "-30.20",
			},
		},
		Phone:   "10498513459 X5490",
		Website: "www.google.com",
		Company: clients.CompanyClientResponse{
			Name:        "Fake Company",
			CatchPhrase: "fake clients",
			BS:          "who knows",
		},
	}, nil
}

type MockPostClient struct {
}

func (m MockPostClient) GetPostsForUser(id int) ([]clients.PostClientResponse, error) {
	return []clients.PostClientResponse{
		{
			ID:    1,
			Title: "Fake title 1",
			Body:  "Body",
		},
		{
			ID:    2,
			Title: "Fake title 2",
			Body:  "Body",
		},
	}, nil
}
