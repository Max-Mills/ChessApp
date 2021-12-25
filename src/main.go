package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

type board struct {
	Board []tile
}

type tile struct {
	Location []int  `json:"loc"`
	Piece    string `json:"piece"`
	Player   string `json:"player"`
}

func doBoard(w http.ResponseWriter, r *http.Request) {
	//allowedHeaders := "Accept, Content-Type, Content-Length, Accept-Encoding, Authorization, X-CSRF-Token"
	switch r.Method {
	case "GET":
		w.Header().Set("Content-Type", "application/json")
		w.Header().Set("Access-Control-Allow-Origin", "*")

		c := w.Header().Get("Set-Cookie")
		c += "; SameSite=lax"
		w.Header().Set("Set-Cookie", c)

		//w.Header().Set("Access-Control-Allow-Headers", allowedHeaders)
		//w.Header().Set("Access-Control-Expose-Headers", "Authorization")
		jsonFile, err := os.Open("src/assets/session-1234.json")
		if err != nil {
			fmt.Println(err)
		}
		defer jsonFile.Close()

		byteValue, _ := ioutil.ReadAll(jsonFile)
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(byteValue))
	case "POST":
		w.WriteHeader(http.StatusCreated)
		w.Write([]byte(`{"message": "POST method requested"}`))
	default:
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte(`{"message": "Can't find method requested"}`))
	}
}

func cors(fs http.Handler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		c := w.Header().Get("Set-Cookie")
		c = "SameSite=None; Secure"
		w.Header().Set("Set-Cookie", c)

		fs.ServeHTTP(w, r)
	}
}

func main() {

	//var boardPlacements []tile
	//json.Unmarshal([]byte(byteValue), &boardPlacements)

	//fmt.Println(boardPlacements)

	fmt.Println("Hello")

	fs := http.FileServer(http.Dir("dist"))
	http.Handle("/", cors(fs))
	http.HandleFunc("/board", doBoard)
	http.Handle("/assets", http.FileServer(http.Dir("src")))

	if err := http.ListenAndServe(":8082", nil); err != nil {
		log.Fatal(err)
	}

}
