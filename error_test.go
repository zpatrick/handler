package handler

import (
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestError(t *testing.T) {
	handler := Error(http.StatusTeapot, errors.New("foo"))
	recorder := httptest.NewRecorder()
	handler.ServeHTTP(recorder, nil)

	assert.Equal(t, http.StatusTeapot, recorder.Code)
	assert.Equal(t, "foo", recorder.Body.String())
	assert.Equal(t, http.Header{}, recorder.HeaderMap)
}

func TestErrorf(t *testing.T) {
	handler := Errorf(http.StatusTeapot, "foo %s %d", "bar", 1)
	recorder := httptest.NewRecorder()
	handler.ServeHTTP(recorder, nil)

	assert.Equal(t, http.StatusTeapot, recorder.Code)
	assert.Equal(t, "foo bar 1", recorder.Body.String())
	assert.Equal(t, http.Header{}, recorder.HeaderMap)
}

func TestErrorWithHeader(t *testing.T) {
	header := http.Header{
		"K1": []string{"k1v1"},
		"K2": []string{"k2v1", "k2v2"},
	}

	handler := ErrorWithHeader(http.StatusTeapot, errors.New("foo"), header)
	recorder := httptest.NewRecorder()
	handler.ServeHTTP(recorder, nil)

	assert.Equal(t, http.StatusTeapot, recorder.Code)
	assert.Equal(t, "foo", recorder.Body.String())
	assert.Equal(t, header, recorder.HeaderMap)
}
