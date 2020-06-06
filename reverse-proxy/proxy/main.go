package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
)

func main()  {
	var url = &url.URL{
		Scheme:     "http",
		Opaque:     "",
		User:       nil,
		Host:       "localhost:10001",
		Path:       "",
		RawPath:    "",
		ForceQuery: false,
		RawQuery:   "",
		Fragment:   "",
	}

	if false {
		log.Fatal(http.ListenAndServe(":10000", httputil.NewSingleHostReverseProxy(url)))
	} else {
		http.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
			fmt.Fprintf(writer, "through proxy:")
			//fmt.Fprintf(writer, "%v", request)
			//fmt.Printf("%s\n", url.String())
			req, _ := http.NewRequest(request.Method, url.String(), nil)

			for key, values := range request.Header {
				for _, value := range values {
					req.Header.Add(key, value)
				}
			}
			client := &http.Client{}
			resp, err := client.Do(req)
			if err != nil {
				log.Fatal(err)
			}

			b, err := ioutil.ReadAll(resp.Body)
			if err != nil {
				log.Fatal(err)
			}

			fmt.Fprintf(writer, "%s", b)

		})

		log.Fatal(http.ListenAndServe(":10000", nil))

	}


}
