package clients

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type IUsersClient interface {
	GetUser(id int) (UserClientResponse, error)
}

type UsersClient struct {
	Client     *http.Client
	BaseDomain string
}

// GetUser fetch information of user
func (u UsersClient) GetUser(id int) (UserClientResponse, error) {
	userClientResponse := UserClientResponse{}
	response, err := u.Client.Get(fmt.Sprintf("%s/users/%d", u.BaseDomain, id))
	if err != nil {
		return userClientResponse, NewClientError("users", fmt.Errorf("error while consuming API : %v", err))
	}
	if response.StatusCode < 200 || response.StatusCode >= 400 {
		return userClientResponse, NewClientError("users", fmt.Errorf("response unsuccessfull : %v", err))
	}
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return userClientResponse, NewClientError("users", fmt.Errorf("cannot read response body : %v", err))
	}
	if err := json.Unmarshal(body, &userClientResponse); err != nil {
		return userClientResponse, NewClientError("users", fmt.Errorf("cannot unmarshal repsonse : %v \n with body %s", err, string(body)))
	}
	return userClientResponse, nil
}

type UserClientResponse struct {
	ID       int                   `json:"id,omitempty"`
	Name     string                `json:"name,omitempty"`
	Username string                `json:"username,omitempty"`
	Email    string                `json:"email,omitempty"`
	Address  AddressClientResponse `json:"address"`
	Phone    string                `json:"phone,omitempty"`
	Website  string                `json:"website,omitempty"`
	Company  CompanyClientResponse `json:"company"`
}

type AddressClientResponse struct {
	Street  string            `json:"street,omitempty"`
	Suite   string            `json:"suite,omitempty"`
	City    string            `json:"city,omitempty"`
	Zipcode string            `json:"zipcode,omitempty"`
	Geo     GeoClientResponse `json:"geo"`
}

type GeoClientResponse struct {
	Latitude  string `json:"lat,omitempty"`
	Longitude string `json:"lng,omitempty"`
}

type CompanyClientResponse struct {
	Name        string `json:"name,omitempty"`
	CatchPhrase string `json:"catchPhrase,omitempty"`
	BS          string `json:"bs,omitempty"`
}
