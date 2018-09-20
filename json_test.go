package handler

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

type Product struct {
	ProductID string `json:"product_id"`
	Quantity  int    `json:"quantity"`
}

func TestJSON(t *testing.T) {
	p := Product{"p1", 3}
	handler := JSON(http.StatusTeapot, p)
	recorder := httptest.NewRecorder()
	handler.ServeHTTP(recorder, nil)

	assert.Equal(t, http.StatusTeapot, recorder.Code)
	assert.Equal(t, http.Header{"Content-Type": []string{"application/json"}}, recorder.HeaderMap)

	var result Product
	if err := json.Unmarshal(recorder.Body.Bytes(), &result); err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, p, result)
}

func TestJSONWithHeader(t *testing.T) {
	header := http.Header{
		"K1": []string{"k1v1"},
		"K2": []string{"k2v1", "k2v2"},
	}

	p := Product{"p1", 3}
	handler := JSONWithHeader(http.StatusTeapot, p, header)
	recorder := httptest.NewRecorder()
	handler.ServeHTTP(recorder, nil)

	assert.Equal(t, http.StatusTeapot, recorder.Code)
	assert.Equal(t, header, recorder.HeaderMap)

	var result Product
	if err := json.Unmarshal(recorder.Body.Bytes(), &result); err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, p, result)
}
