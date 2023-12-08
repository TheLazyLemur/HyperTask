package main

import (
	"net/http"
	"strconv"
)

func formValueAsInt(r *http.Request, name string) (int, error) {
	value := r.FormValue(name)
	if value == "" {
		return 0, nil
	}

	return strconv.Atoi(value)
}
