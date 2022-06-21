package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"strconv"
)

type pagination struct {
	currentPage int
	totalPages  int
	boundaries  int
	around      int
}

func getPagination(bounderies int, currentPage int, totalPages int, around int) (result string) {
	earlyElipse := false
	lateElipse := false

	for i := 1; i <= totalPages; i++ {
		if i == currentPage ||
			((i < currentPage && i >= currentPage-around) || (i > currentPage && i <= currentPage+around)) ||
			(i <= bounderies || i > totalPages-bounderies) {
			result += strconv.Itoa(i) + " "
			continue
		}

		if !earlyElipse && i > bounderies && i < currentPage {
			result += "... "
			earlyElipse = true
			continue
		}

		if !lateElipse && i > currentPage && i <= totalPages-bounderies {
			result += "... "
			lateElipse = true
			continue
		}

	}

	return result
}

func viewHandler(w http.ResponseWriter, r *http.Request) {
	queryParams := r.URL.Query()
	paginations, err := validateInput(queryParams)

	if err != nil {
		errorValidator(w, "Error!")
		return
	}

	result := getPagination(paginations.boundaries, paginations.currentPage, paginations.totalPages, paginations.around)

	ok(w, result)
}

func validateInput(queryParams url.Values) (pagination, error) {
	boundaries, err := strconv.Atoi(queryParams.Get("boundaries"))
	if err != nil {
		return pagination{}, err
	}
	currentPage, err := strconv.Atoi(queryParams.Get("current_page"))
	if err != nil {
		return pagination{}, err
	}
	totalPages, err := strconv.Atoi(queryParams.Get("total_pages"))
	if err != nil {
		return pagination{}, err
	}
	around, err := strconv.Atoi(queryParams.Get("around"))
	if err != nil {
		return pagination{}, err
	}

	if boundaries < 0 || currentPage < 0 || totalPages < 0 || around < 0 {
		return pagination{}, fmt.Errorf("invalid input")
	}

	return pagination{
		boundaries:  boundaries,
		currentPage: currentPage,
		totalPages:  totalPages,
		around:      around,
	}, nil
}

func ok(w http.ResponseWriter, result string) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	json.NewEncoder(w).Encode(result)
}

func errorValidator(w http.ResponseWriter, result string) {
	w.WriteHeader(http.StatusBadRequest)

	json.NewEncoder(w).Encode(result)
}

func main() {
	http.HandleFunc("/", viewHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
