/// This file is a part of Harmonoid (https://github.com/harmonoid/harmonoid).
///
/// Copyright © 2022, Harmonoid <saini123hitesh@gmail.com>.
/// All rights reserved.
///
/// Use of this source code is governed by the End-User License Agreement for Harmonoid that can be found in the EULA.txt file in Harmonoid's repository (https://github.com/harmonoid/harmonoid).
///

package main

import (
	"compress/gzip"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/imroc/req/v3"
	"io"
	"net/http"
	"net/url"
	"os"
)

func main() {
	fmt.Println("Starting Harmonoid proxy...")

	r := mux.NewRouter()
	r.HandleFunc("/{type}/youtubei/v1/{endpoint:.*}", func(w http.ResponseWriter, r *http.Request) {
		client := req.C()
		client.DisableAutoReadResponse()
		request := client.R()

		// Loop over header names
		for name, values := range r.Header {
			// Loop over all values for the name.
			for _, value := range values {
				request.SetHeader(name, value)
			}
		}

		var uri = ""

		if mux.Vars(r)["type"] == "music" {
			uri = fmt.Sprintf("https://music.youtube.com/youtubei/v1/%s", mux.Vars(r)["endpoint"])
		} else if mux.Vars(r)["type"] == "youtube" {
			uri = fmt.Sprintf("https://www.youtube.com/youtubei/v1/%s", mux.Vars(r)["endpoint"])
		}

		if uri == "" {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte("Invalid request"))
			return
		}

		var response *req.Response
		var err error

		parse, err := url.Parse(uri)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte("Could not parse URI"))
			return
		}

		q := parse.Query()

		for name, values := range r.URL.Query() {
			for _, value := range values {
				q.Add(name, value)
			}
		}

		parse.RawQuery = q.Encode()

		uri = parse.String()

		if r.Method == "GET" {
			fmt.Printf("Serving GET request to %s using headers %s\n", uri, request.Headers)

			response, err = request.Get(uri)
		} else if r.Method == "POST" {
			var b []byte
			b, err = io.ReadAll(r.Body)
			if err != nil {
				w.Write([]byte(err.Error()))
				return
			}
			request.SetBody(string(b))

			fmt.Printf("Serving POST request to %s using headers %s and body %s\n", uri, request.Headers, string(b))

			response, err = request.Post(uri)
		}
		if err != nil {
			w.Write([]byte(err.Error()))
			return
		}
		res, err := gzip.NewReader(response.Body)
		if err != nil {
			fmt.Println(err.Error())
			w.Write([]byte(err.Error()))
			return
		}
		responseBody, err := io.ReadAll(res)
		if err != nil {
			fmt.Println(err.Error())
			w.Write([]byte(err.Error()))
			return
		}
		w.WriteHeader(response.StatusCode)
		w.Write(responseBody)
	})

        port := "80"
        if os.Getenv("PORT") != "" {
                port = os.Getenv("PORT")
        }

	err := http.ListenAndServe(":" + port, r)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
}
