package handler

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBytes(t *testing.T) {
	handler := Bytes(http.StatusTeapot, []byte("foo"))
	recorder := httptest.NewRecorder()
	handler.ServeHTTP(recorder, nil)

	assert.Equal(t, http.StatusTeapot, recorder.Code)
	assert.Equal(t, "foo", recorder.Body.String())
	assert.Equal(t, http.Header{}, recorder.HeaderMap)
}

func TestBytesWithHeader(t *testing.T) {
	header := http.Header{
		"K1": []string{"k1v1"},
		"K2": []string{"k2v1", "k2v2"},
	}

	handler := BytesWithHeader(http.StatusTeapot, []byte("foo"), header)
	recorder := httptest.NewRecorder()
	handler.ServeHTTP(recorder, nil)

	assert.Equal(t, http.StatusTeapot, recorder.Code)
	assert.Equal(t, "foo", recorder.Body.String())
	assert.Equal(t, header, recorder.HeaderMap)
}
