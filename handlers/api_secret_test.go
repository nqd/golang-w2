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
	var server *httptest.Server = httptest.NewServer(NewRouter())
	url := fmt.Sprintf("%s/v1/secret", server.URL)

	body, err := json.Marshal(map[string]interface{}{
		"secret":           "foo",
		"expireAfterViews": 2,
		"expireAfter":      1,
	})
	if err != nil {
		t.Fatal(err)
	}

	// Create a request to pass to our handler. We don't have any query parameters for now, so we'll
	// pass 'nil' as the third parameter.
	req, err := http.NewRequest("POST", url, strings.NewReader(string(body)))
	if err != nil {
		t.Fatal(err)
	}
	res, err := http.DefaultClient.Do(req)
	assert.Nil(t, err)
	assert.Equal(t, res.StatusCode, 200)

}
