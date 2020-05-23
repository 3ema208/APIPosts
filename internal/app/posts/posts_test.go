package posts

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAPIPosts_handlerPost(t *testing.T) {
	conf := &Config{}
	apiPosts := New(conf)
	rec := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodGet, "/posts", nil)
	apiPosts.handlePosts().ServeHTTP(rec, req)
	assert.Equal(t, rec.Body.String(), "This Posts")
}
