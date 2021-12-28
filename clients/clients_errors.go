package clients

import "fmt"

type ClientError struct {
	ClientName string
	Cause      error
}

func (c *ClientError) Error() string {
	if c.Cause == nil {
		return fmt.Sprintf("error while consuming client %s . No root cause", c.ClientName)
	}
	return fmt.Sprintf("error while consuming client %s , root cause : %v \n", c.ClientName, c.Cause.Error())
}

func NewClientError(clientName string, err error) error {
	return &ClientError{
		ClientName: clientName,
		Cause:      err,
	}
}
