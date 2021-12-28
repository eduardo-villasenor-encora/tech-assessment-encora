package clients

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type IPostsClient interface {
	GetPostsForUser(id int) ([]PostClientResponse, error)
}

type PostsClient struct {
	Client     *http.Client
	BaseDomain string
}

// GetPostsForUser fetch information of posts related to a given user
func (p PostsClient) GetPostsForUser(id int) ([]PostClientResponse, error) {
	postClientResponse := make([]PostClientResponse, 0)
	response, err := p.Client.Get(fmt.Sprintf("%s/posts?userId=%d", p.BaseDomain, id))
	if err != nil {
		return postClientResponse, NewClientError("posts", fmt.Errorf("error while consuming API : %v", err))
	}
	if response.StatusCode < 200 || response.StatusCode >= 400 {
		return postClientResponse, NewClientError("posts", fmt.Errorf("response unsuccessfull , with status %d : %v", response.StatusCode, err))
	}
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return postClientResponse, NewClientError("posts", fmt.Errorf("cannot read response body : %v", err))
	}
	if err := json.Unmarshal(body, &postClientResponse); err != nil {
		return postClientResponse, NewClientError("posts", fmt.Errorf("cannot unmarshal repsonse : %v \n with body %s", err, string(body)))
	}
	return postClientResponse, nil
}

type PostClientResponse struct {
	UserID int    `json:"userId,omitempty"`
	ID     int    `json:"id,omitempty"`
	Title  string `json:"title,omitempty"`
	Body   string `json:"body,omitempty"`
}
