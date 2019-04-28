package handlers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"

	"github.com/stretchr/testify/assert"
)

var server = httptest.NewServer(NewRouter())
var hash string

func TestCreateSecret(t *testing.T) {
	uri := fmt.Sprintf("%s/v1/secret", server.URL)

	data := url.Values{}
	data.Set("secret", `foo`)
	data.Add("expireAfterViews", `2`)
	data.Add("expireAfter", `1`)

	req, err := http.NewRequest("POST", uri, bytes.NewBufferString(data.Encode()))
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded; param=value")
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
	assert.Equal(t, resBody.SecretText, "foo")
	assert.EqualValues(t, resBody.RemainingViews, 1)
	assert.NotEmpty(t, resBody.Hash)
	assert.NotEmpty(t, resBody.CreatedAt)
	assert.NotEmpty(t, resBody.ExpiresAt)
}
