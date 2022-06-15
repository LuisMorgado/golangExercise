package main

import "testing"

func TestGetPagination(t *testing.T) {
	if getPagination(1, 4, 5, 0) != "1 ... 4 5 " {
		t.Error("Expected a string like this: 1 ... 4 5 ")
	}

	if getPagination(2, 4, 10, 2) != "1 2 3 4 5 6 ... 9 10 " {
		t.Error("Expected a string like this: 1 2 3 4 5 6 ... 9 10 ")
	}
}
