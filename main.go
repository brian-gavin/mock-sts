package main

import (
	"encoding/xml"
	"log"
	"net/http"

	"github.com/google/uuid"
)

type GetCallerIdentityResponse struct {
	Result GetCallerIdentityResult `xml:"GetCallerIdentityResult"`
	Meta   ResponseMetadata        `xml:"ResponseMetadata"`
}

type GetCallerIdentityResult struct {
	Arn     string
	UserId  string
	Account string
}

type ResponseMetadata struct {
	RequestId string
}

func main() {
	const addr = ":8088"
	log.Printf("listen on %s\n", addr)
	response, _ := xml.Marshal(GetCallerIdentityResponse{
		Result: GetCallerIdentityResult{
			Arn:     "arn:aws:iam::1234567:test2",
			UserId:  "test2",
			Account: "2345678",
		},
		Meta: ResponseMetadata{
			RequestId: uuid.New().String(),
		},
	})
	log.Println(string(response))
	http.ListenAndServe(addr, http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Printf("Request:\n")
		r.Header.Write(log.Writer())
		defer r.Body.Close()
		w.Write(response)
	}))
}
