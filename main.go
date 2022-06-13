package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"
)

type Pagination struct {
	CurrentPage int
	TotalPages  int
	Boundaries  int
	Around      int
}

func getQueryParameters(parameters string) map[string]int {
	var parametersKV = map[string]int{}
	if parameters == "" {
		return parametersKV
	}

	splitedParameters := strings.Split(parameters, "&")

	for i := 0; i < len(splitedParameters); i++ {
		splitedItem := strings.Split(splitedParameters[i], "=")
		stringValue, _ := strconv.ParseInt(splitedItem[1], 10, 0)
		parametersKV[strings.ToLower(splitedItem[0])] = int(stringValue)
	}

	return parametersKV
}

func getPagination(bounderies int, currentPage int, totalPages int, around int) string {
	result := ""
	earlyElipse := false
	lateElipse := false

	for i := 1; i <= totalPages; i++ {

		condition := []bool{
			i == currentPage,
			i < currentPage && i >= currentPage-around,
			i > currentPage && i <= currentPage+around,
			i <= bounderies || i > totalPages-bounderies,
		}

		fmt.Println(condition)

		if i == currentPage ||
			((i < currentPage && i >= currentPage-around) || (i > currentPage && i <= currentPage+around)) ||
			(i <= bounderies || i > totalPages-bounderies) {
			result += fmt.Sprintf("%d", i) + " "
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
	queryParams := getQueryParameters(r.URL.RawQuery)
	result := getPagination(queryParams["boundaries"], queryParams["current_page"], queryParams["total_pages"], queryParams["around"])

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
