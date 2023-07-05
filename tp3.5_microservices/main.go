package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

const localPort = "80"

func main() {
	http.HandleFunc("/", actAsAProxy)
	http.HandleFunc("/favicon.ico", func(w http.ResponseWriter, request *http.Request) {})
	fmt.Println("Server Web started on port", localPort)
	log.Fatal(http.ListenAndServe(":"+localPort, nil))
}

func actAsAProxy(w http.ResponseWriter, r *http.Request) {
	res, _ := callFakeAPI(r.URL.Path)
	println(res)
	println(string(res))
	fmt.Fprint(w, string(res))
}

func callFakeAPI(path string) ([]byte, error) {
	const serverURL = "api"
	const serverPort = 8080

	requestURL := fmt.Sprintf("http://%s:%d%s", serverURL, serverPort, path)
	println("calling: ", requestURL)
	res, err := http.Get(requestURL)
	if err != nil {
		fmt.Printf("error making http request: %s\n", err)
	}

	fmt.Printf("client: got response!\n")
	fmt.Printf("client: status code: %d\n", res.StatusCode)

	resBody, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Printf("client: could not read response body: %s\n", err)
	}
	fmt.Printf("client: response body: %s\n", resBody)
	return resBody, nil
}
