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
	w.Header().Set("Content-Type", "application/json")
	switch r.Method {
	case "GET":
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
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

func main() {

	//var boardPlacements []tile
	//json.Unmarshal([]byte(byteValue), &boardPlacements)

	//fmt.Println(boardPlacements)

	fmt.Println("Hello")
	fileServer := http.FileServer(http.Dir("src"))
	http.Handle("/", fileServer)
	http.HandleFunc("/board", doBoard)

	if err := http.ListenAndServe(":8082", nil); err != nil {
		log.Fatal(err)
	}

}
