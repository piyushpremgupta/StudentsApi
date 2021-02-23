package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
	"path/filepath"
	"strconv"
)

func GetStudents(w http.ResponseWriter, r *http.Request) {
	Students := map[int]map[string]string{1: map[string]string{"Name": "Piyush", "Address": "Pune", "Age": "31"}, 2: map[string]string{"Name": "Sagar", "Address": "Pune", "Age": "32"}}
	id := filepath.Base(r.URL.Path)
	fmt.Println(id)
	if id == "getStudents" {
		json.NewEncoder(w).Encode(&Students)
	} else {
		id, _ := strconv.Atoi(id)
		json.NewEncoder(w).Encode(Students[id])
	}
}
