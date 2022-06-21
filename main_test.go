package main

import (
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"
)

func TestGetPagination(t *testing.T) {
	testCases := []struct {
		boundaries  int
		currentPage int
		totalPages  int
		around      int
		expected    string
	}{
		{
			boundaries:  1,
			currentPage: 4,
			totalPages:  5,
			around:      0,
			expected:    "1 ... 4 5 ",
		},
		{
			boundaries:  2,
			currentPage: 4,
			totalPages:  10,
			around:      2,
			expected:    "1 2 3 4 5 6 ... 9 10 ",
		},
	}

	for _, test := range testCases {
		result := getPagination(test.boundaries, test.currentPage, test.totalPages, test.around)
		if result != test.expected {
			t.Error("expected something else")
		}
	}
}

func TestViewHandler(t *testing.T) {
	testCases := []struct {
		boundaries  int
		currentPage int
		totalPages  int
		around      int
		expected    string
	}{
		{
			boundaries:  1,
			currentPage: 4,
			totalPages:  5,
			around:      0,
			expected:    "1 ... 4 5 ",
		},
		{
			boundaries:  2,
			currentPage: 4,
			totalPages:  10,
			around:      2,
			expected:    "1 2 3 4 5 6 ... 9 10 ",
		},
	}

	for _, test := range testCases {
		bdr := strconv.Itoa(test.boundaries)
		tpages := strconv.Itoa(test.totalPages)
		arnd := strconv.Itoa(test.around)
		cpages := strconv.Itoa(test.currentPage)

		req, err := http.NewRequest("GET", "localhost:8080?boundaries="+bdr+"&total_pages="+tpages+"&around="+arnd+"&current_page="+cpages, nil)

		if err != nil {
			t.Error(err)
			return
		}

		rec := httptest.NewRecorder()
		viewHandler(rec, req)

		res := rec.Result()

		if res.StatusCode != 200 {
			t.Error()
			return
		}
	}

}
