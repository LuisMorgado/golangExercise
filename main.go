package main

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"
)

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

	boundaries, _ := strconv.Atoi(queryParams.Get("boundaries"))
	currentPage, _ := strconv.Atoi(queryParams.Get("current_page"))
	totalPages, _ := strconv.Atoi(queryParams.Get("total_pages"))
	around, _ := strconv.Atoi(queryParams.Get("around"))

	if boundaries < 0 || currentPage < 0 || totalPages < 0 || around < 0 {
		return
	}

	result := getPagination(boundaries, currentPage, totalPages, around)

	ok(w, result)
}

func ok(w http.ResponseWriter, result string) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	json.NewEncoder(w).Encode(result)
}

func main() {
	http.HandleFunc("/", viewHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
