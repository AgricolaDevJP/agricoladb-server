package test

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"testing"

	"github.com/99designs/gqlgen/graphql"
	"github.com/stretchr/testify/assert"
)

const (
	playgroundURL = "http://localhost:8080"
	graphqlURL    = "http://localhost:8080/query"
)

func TestE2EPlayGround(t *testing.T) {
	res, err := http.Get(playgroundURL)
	assert.Nil(t, err)
	assert.Equal(t, res.StatusCode, 200, "status code should be 200")
}

type RequestBody struct {
	Query string `json:"query"`
}

func TestE2ECardsQuery(t *testing.T) {
	query := `
		query {
			cards(where: {revisionID: 1}, first: 20) {
				edges {
					node {
						literalID
						nameJa
					}
				}
			}
		}
	`
	reqBodyReader, _ := generateReqBodyReader(query)
	res, err := http.Post(graphqlURL, "application/json", reqBodyReader)
	assert.Nil(t, err)
	defer res.Body.Close()
	assert.Equal(t, res.StatusCode, 200, "status code should be 200")

	resRawBody, _ := io.ReadAll(res.Body)
	var resBody graphql.Response
	err = json.Unmarshal(resRawBody, &resBody)
	assert.Nil(t, err)

	var resBodyData []byte
	err = resBody.Data.UnmarshalJSON(resBodyData)
	fmt.Printf("resBody.Data: %v\n", string(resBodyData))
	assert.Nil(t, err)
}

func generateReqBodyReader(query string) (*bytes.Reader, error) {
	reqBody := &RequestBody{
		Query: query,
	}
	reqBodyJson, err := json.Marshal(reqBody)
	if err != nil {
		return nil, err
	}
	reqBodyReader := bytes.NewReader(reqBodyJson)
	return reqBodyReader, nil
}
