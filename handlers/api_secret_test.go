package handlers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

var server = httptest.NewServer(NewRouter())
var hash string

func TestCreateSecret(t *testing.T) {
	url := fmt.Sprintf("%s/v1/secret", server.URL)

	body, err := json.Marshal(map[string]interface{}{
		"secret":           "foo",
		"expireAfterViews": 2,
		"expireAfter":      1,
	})
	assert.Nil(t, err)

	// Create a request to pass to our handler. We don't have any query parameters for now, so we'll
	// pass 'nil' as the third parameter.
	req, err := http.NewRequest("POST", url, strings.NewReader(string(body)))
	if err != nil {
		assert.Nil(t, err)
	}
	res, err := http.DefaultClient.Do(req)
	assert.Nil(t, err)

	defer res.Body.Close()
	assert.Equal(t, res.StatusCode, 200)

	var resBody Secret
	json.NewDecoder(res.Body).Decode(&resBody)
	assert.Equal(t, resBody.SecretText, "foo")
	assert.EqualValues(t, resBody.RemainingViews, 2)
	assert.NotEmpty(t, resBody.Hash)
	assert.NotEmpty(t, resBody.CreatedAt)
	assert.NotEmpty(t, resBody.ExpiresAt)
	hash = resBody.Hash
}

func TestGetSecret(t *testing.T) {
	url := fmt.Sprintf("%s/v1/secret/%s", server.URL, hash)

	// Create a request to pass to our handler. We don't have any query parameters for now, so we'll
	// pass 'nil' as the third parameter.
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		assert.Nil(t, err)
	}
	res, err := http.DefaultClient.Do(req)
	assert.Nil(t, err)

	defer res.Body.Close()
	assert.Equal(t, res.StatusCode, 200)

	var resBody Secret
	json.NewDecoder(res.Body).Decode(&resBody)
	log.Println(resBody)
	assert.Equal(t, resBody.SecretText, "foo")
	assert.EqualValues(t, resBody.RemainingViews, 1)
	assert.NotEmpty(t, resBody.Hash)
	assert.NotEmpty(t, resBody.CreatedAt)
	assert.NotEmpty(t, resBody.ExpiresAt)
}
