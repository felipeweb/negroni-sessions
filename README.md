# negroni-sessions 
[![GoDoc](https://godoc.org/github.com/felipeweb/negroni-sessions?status.svg)](http://godoc.org/github.com/felipeweb/negroni-sessions)
[![Build Status](https://travis-ci.org/felipeweb/negroni-sessions.svg?branch=master)](https://travis-ci.org/felipeweb/negroni-sessions)

Negroni middleware/handler for easy session management.

## Usage

~~~ go
package main

import (
  "github.com/urfave/negroni"
  "github.com/felipeweb/negroni-sessions"
  "github.com/felipeweb/negroni-sessions/cookiestore"
  "net/http"
)

func main() {
  n := negroni.Classic()

  store := cookiestore.New([]byte("secret123"))  
  n.Use(sessions.Sessions("my_session", store))

  mux := http.NewServeMux()
  mux.HandleFunc("/", func(w http.ResponseWriter, req *http.Request) {
    session := sessions.GetSession(req)
    session.Set("hello", "world")
  })

  n.UseHandler(mux)
  n.Run(":3000")
}

~~~

A fork of [https://github.com/GoIncremental/negroni-sessions](https://github.com/GoIncremental/negroni-sessions)
