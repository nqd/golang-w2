package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreateSecret(t *testing.T) {
	server := httptest.NewServer(NewRouter())
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

	var resBody interface{}
	json.NewDecoder(res.Body).Decode(&resBody)
	assert.Contains(t, resBody, "hash")
	assert.Contains(t, resBody, "secretText")
	assert.Contains(t, resBody, "createdAt")
	assert.Contains(t, resBody, "expiresAt")
	assert.Contains(t, resBody, "remainingViews")
}
