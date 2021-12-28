package services

import (
	"log"
	"sync"
	"tech-assessment/clients"
)

type IUserService interface {
	GetUserPosts(id int) (UserInfo, error)
}

type UserService struct {
	UserClient clients.IUsersClient
	PostClient clients.IPostsClient
}

type UserServiceResponse struct {
	ID       int      `json:"id,omitempty"`
	UserInfo UserInfo `json:"userInfo"`
	Posts    []Post   `json:"posts,omitempty"`
}

type UserInfo struct {
	Name     string `json:"name,omitempty"`
	Username string `json:"username,omitempty"`
	Email    string `json:"email,omitempty"`
}

type Post struct {
	ID    int    `json:"id,omitempty"`
	Title string `json:"title,omitempty"`
	Body  string `json:"body,omitempty"`
}

// GetUserPosts calls the clients in order to construct merge user information
func (u UserService) GetUserPosts(id int) (UserServiceResponse, error) {
	var err error
	var wg sync.WaitGroup

	var userData clients.UserClientResponse
	var postData []clients.PostClientResponse

	wg.Add(1)
	go func() {
		defer wg.Done()
		userData, err = u.UserClient.GetUser(id)
		if err != nil {
			log.Println("cannot fetch user client information")
		}
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		postData, err = u.PostClient.GetPostsForUser(id)
		if err != nil {
			log.Println("cannot fetch posts client information")
		}
	}()

	wg.Wait()
	return UserInfoBuilder(userData, postData), nil
}

func UserInfoBuilder(userData clients.UserClientResponse, posts []clients.PostClientResponse) UserServiceResponse {
	userInfo := UserServiceResponse{}
	userInfo.ID = userData.ID
	userInfo.UserInfo.Name = userData.Name
	userInfo.UserInfo.Email = userData.Email
	userInfo.UserInfo.Username = userData.Username
	userInfo.Posts = make([]Post, 0)
	for _, post := range posts {
		userInfo.Posts = append(userInfo.Posts, Post{
			ID:    post.ID,
			Title: post.Title,
			Body:  post.Body,
		})
	}
	return userInfo
}
