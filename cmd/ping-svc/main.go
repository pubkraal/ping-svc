package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/ping", func(w http.ResponseWriter, r *http.Request) {
		b, err := json.Marshal(r.Header)
		if err != nil {
			fmt.Fprint(w, "huge error")
			return
		}

		_, err = w.Write(b)
		if err != nil {
			log.Println("Error!", err)
		}
	})

	log.Fatal(http.ListenAndServe(":8080", nil))
}
