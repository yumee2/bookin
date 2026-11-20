package httpapi

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGetHouses(t *testing.T) {
	router := NewRouter()
	req := httptest.NewRequest(http.MethodGet, "/houses", nil)
	rec := httptest.NewRecorder()

	router.ServeHTTP(rec, req)

	if rec.Code != http.StatusOK {
		t.Fatalf("status = %d, want %d; body = %s",
			rec.Code, http.StatusOK, rec.Body.String())
	}

	var got []house
	if err := json.Unmarshal(rec.Body.Bytes(), &got); err != nil {
		t.Fatalf("decode response: %v", err)
	}

	want := []house{
		{ID: 1, Name: "Big house", Capacity: 8},
		{ID: 2, Name: "Small house", Capacity: 4},
	}

	if len(got) != len(want) {
		t.Fatalf("got %d houses, want %d", len(got), len(want))
	}

	for i := range want {
		if got[i] != want[i] {
			t.Errorf("house[%d] = %+v, want %+v", i, got[i], want[i])
		}
	}
}
