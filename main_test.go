package main


import (
	"bytes"
	"fmt"
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"net/http"
	"testing"
)

func TestGetContents(t *testing.T) {

	// using real client
	client := MockClient{}
	parser := NewParser(&client)

	// get full contents of example.com html
	html, err := parser.GetContents("http://example.com/", 10)
	if err != nil {
		fmt.Println("Error: ", err)
		return
	}

	assert.NotNil(t, html)
	assert.EqualValues(t, "<!doctype ", html)

}


func (c *MockClient) Get(url string) (resp *http.Response, err error) {

	body := "<html></html>"

	r := ioutil.NopCloser(bytes.NewReader([]byte(body)))

	return &http.Response{
		StatusCode: 200,
		Body:       r,
	}, nil


}