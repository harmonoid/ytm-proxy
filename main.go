/// This file is a part of Harmonoid (https://github.com/harmonoid/harmonoid).
///
/// Copyright © 2020-2022, Hitesh Kumar Saini <saini123hitesh@gmail.com>.
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
	"os"
)

func main() {
	fmt.Println("Starting Harmonoid proxy...")

	r := mux.NewRouter()
	r.HandleFunc("/music", func(w http.ResponseWriter, r *http.Request) {
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

		url := fmt.Sprintf("https://music.youtube.com%s", r.URL.Query().Get("url"))

		if r.Method == "GET" {
			fmt.Printf("Serving GET request to %s using headers %s\n", url, request.Headers)

			get, err := request.Get(url)
			if err != nil {
				w.Write([]byte(err.Error()))
				return
			}
			res, err := gzip.NewReader(get.Body)
			if err != nil {
				fmt.Println(err.Error())
				w.Write([]byte(err.Error()))
				return
			}
			r, err := io.ReadAll(res)
			if err != nil {
				fmt.Println(err.Error())
				w.Write([]byte(err.Error()))
				return
			}
			w.WriteHeader(get.StatusCode)
			w.Write(r)
		} else if r.Method == "POST" {
			b, err := io.ReadAll(r.Body)
			if err != nil {
				w.Write([]byte(err.Error()))
				return
			}
			request.SetBody(string(b))

			fmt.Printf("Serving POST request to %s using headers %s and body %s\n", url, request.Headers, string(b))

			post, err := request.Post(url)
			if err != nil {
				fmt.Println(err.Error())
				w.Write([]byte(err.Error()))
				return
			}
			res, err := gzip.NewReader(post.Body)
			if err != nil {
				fmt.Println(err.Error())
				w.Write([]byte(err.Error()))
				return
			}
			r, err := io.ReadAll(res)
			if err != nil {
				fmt.Println(err.Error())
				w.Write([]byte(err.Error()))
				return
			}
			w.WriteHeader(post.StatusCode)
			w.Write(r)
		}
	})

	err := http.ListenAndServe("0.0.0.0:80", r)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
}
