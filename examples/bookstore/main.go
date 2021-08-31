package main

import (
	"fmt"
	"github.com/MattLaidlaw/godb-go-driver/pkg/driver"
	"log"
	"net/http"
	"strconv"
)

var client *driver.Client

func BookHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPost:
		_ = r.ParseForm()
		fmt.Println(r.Form)
		insertedCount, _ := client.Set(r.Form.Get("author"), r.Form.Get("title"))
		_, _ = w.Write([]byte(strconv.Itoa(insertedCount)))
	case http.MethodGet:
		_ = r.ParseForm()
		fmt.Println(r.Form)
		//val, _ := client.Get(r.Form.Get("author"))
		val, _ := client.Get("asimov")
		_, _ = w.Write([]byte(val))
	case http.MethodDelete:
		_ = r.ParseForm()
		fmt.Println(r.Form)
		deletedCount, _ := client.Del(r.Form.Get("author"))
		_, _ = w.Write([]byte(strconv.Itoa(deletedCount)))
	default:
		w.Write([]byte("404 not found"))
	}
}

func main() {

	var err error
	client, err = driver.NewClient("localhost:6342")
	if err != nil {
		log.Fatalln(err)
	}

	http.HandleFunc("/books", BookHandler)
	log.Println("== listening on port 8080")
	log.Println(http.ListenAndServe(":8080", nil))

}
