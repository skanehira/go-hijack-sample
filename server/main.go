package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		hj, ok := w.(http.Hijacker)
		if !ok {
			http.Error(w, "is not hijacker", http.StatusInternalServerError)
			return
		}

		conn, bufrw, err := hj.Hijack()

		if err != nil {
			http.Error(w, "cannot get hijack", http.StatusInternalServerError)
			return
		}
		defer conn.Close()

		bufrw.WriteString("now tcp")
		bufrw.Flush()

		s, err := bufrw.ReadString('\n')
		if err != nil {
			log.Printf("error read string :%v", err)
			return
		}

		fmt.Fprintf(bufrw, "you said: %q\nBye.\n", s)
		bufrw.Flush()
	})

	log.Fatal(http.ListenAndServe(":80", nil))
}
