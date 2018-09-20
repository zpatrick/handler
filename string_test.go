package handler

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestString(t *testing.T) {
	handler := String(http.StatusTeapot, "foo")
	recorder := httptest.NewRecorder()
	handler.ServeHTTP(recorder, nil)

	assert.Equal(t, http.StatusTeapot, recorder.Code)
	assert.Equal(t, "foo", recorder.Body.String())
	assert.Equal(t, http.Header{}, recorder.HeaderMap)
}

func TestStringf(t *testing.T) {
	handler := Stringf(http.StatusTeapot, "foo %s %d", "bar", 1)
	recorder := httptest.NewRecorder()
	handler.ServeHTTP(recorder, nil)

	assert.Equal(t, http.StatusTeapot, recorder.Code)
	assert.Equal(t, "foo bar 1", recorder.Body.String())
	assert.Equal(t, http.Header{}, recorder.HeaderMap)
}

func TestStringWithHeader(t *testing.T) {
	header := http.Header{
		"K1": []string{"k1v1"},
		"K2": []string{"k2v1", "k2v2"},
	}

	handler := StringWithHeader(http.StatusTeapot, "foo", header)
	recorder := httptest.NewRecorder()
	handler.ServeHTTP(recorder, nil)

	assert.Equal(t, http.StatusTeapot, recorder.Code)
	assert.Equal(t, "foo", recorder.Body.String())
	assert.Equal(t, header, recorder.HeaderMap)
}
