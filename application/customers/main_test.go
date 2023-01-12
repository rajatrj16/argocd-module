package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGetCustomers(t *testing.T) {
	t.Run("Testing the status of the request", func(t *testing.T) {
		// creating a request to pass to our handler
		req := httptest.NewRequest(http.MethodGet, "localhost:3000", nil)
		rr := httptest.NewRecorder()
		handler := http.HandlerFunc(getCustomers)

		handler.ServeHTTP(rr, req)

		// check the status code is what we expect.
		if status := rr.Code; status != http.StatusOK {
			t.Errorf("handler returned wrong status code")
		}
	})

	t.Run("Testing the expected output", func(t *testing.T) {
		request, _ := http.NewRequest(http.MethodGet, "localhost:3000", nil)
		response := httptest.NewRecorder()

		getCustomers(response, request)

		got := response.Body.String()
		want := "[{\"age\":75,\"id\":1,\"name\":\"Jewel Schaefer\"},{\"age\":26,\"id\":2,\"name\":\"Raleigh Larson\"},{\"age\":50,\"id\":3,\"name\":\"Eloise Senger\"},{\"age\":29,\"id\":4,\"name\":\"Moshe Zieme\"},{\"age\":45,\"id\":5,\"name\":\"Filiberto Lubowitz\"},{\"age\":32,\"id\":6,\"name\":\"Ms.Kadin Kling\"},{\"age\":28,\"id\":7,\"name\":\"Jennyfer Bergstrom\"},{\"age\":27,\"id\":8,\"name\":\"Candelario Rutherford\"},{\"age\":43,\"id\":9,\"name\":\"Kenyatta Flatley\"},{\"age\":27,\"id\":10,\"name\":\"Rajat Jain\"}]\n"

		if got != want {
			t.Errorf("got %q, want %q", got, want)
		}
	})

}
