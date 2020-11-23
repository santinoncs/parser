package main

import (
  "net/http"
  "fmt"
  "io/ioutil"
)

// Parser struct
type Parser struct {
	//*http.Client  // we'll need to replace this with an interface to be able to mock it
	client ClientInterface
}

type MockClient struct {}

type ClientInterface interface {
       Get(string) (*http.Response, error)
}

// Constructor
func NewParser(client ClientInterface) *Parser {
  return &Parser{client}
}

// Returns contents of the page located at url, limiting them to limit characters.
// If limit is set to 0, return full contents without truncating.
func (p Parser) GetContents(url string, limit int) (string, error) {
  p.client = &http.Client{}
  resp, err := p.client.Get(url)
  if err != nil {
    return "", err
  }
  defer resp.Body.Close()

  if resp.StatusCode != http.StatusOK {
      return "", fmt.Errorf("Http code returned: %d", resp.StatusCode)
  }

  data, err := ioutil.ReadAll(resp.Body)
  if err != nil {
    return "", err
  }

  str := string(data)
  if limit > 0 {
    str = str[:limit]
  }

  return str, nil
}
