package main

import (
	"fmt"
	"log"
	"net/http"

	"rsc.io/qr"
)

func serve(qrs []*qr.Code) {
	page := "<html><body>"
	mux := http.NewServeMux()
	for i := range qrs {
		page += fmt.Sprintf(
			`<img src="/image/%v.png" width="300px">`, i)
		func(j int, img *qr.Code) {
			mux.HandleFunc(
				fmt.Sprintf("/image/%v.png", j),
				func(w http.ResponseWriter, r *http.Request) {
					w.Header().Set("Content-Type", "image/png")
					_, _ = w.Write(img.PNG())
				})
		}(i, qrs[i])
	}
	page += "</body></html>"
	mux.HandleFunc(
		"/",
		func(w http.ResponseWriter, r *http.Request) {
			_, _ = w.Write([]byte(page))
		})
	fmt.Println("serving QR codes at http://localhost:3000")
	log.Fatal(http.ListenAndServe(":3000", mux))
}
