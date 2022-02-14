package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gomodule/redigo/redis"
)

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

		jsonFile, err := os.Open("../src/assets/session-1234.json")
		if err != nil {
			fmt.Println(err)
		}
		defer jsonFile.Close()

		//byteValue, _ := ioutil.ReadAll(jsonFile)
		byteValue := getBoard()
		fmt.Println(byteValue)
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

func getBoard() []byte {
	conn, _ := redis.Dial("tcp", "redis-12016.c74.us-east-1-4.ec2.cloud.redislabs.com:12016", redis.DialPassword("RpMBNyGMzig5cFTifZhTtLcCLWkMbGvV"))
	r, _ := redis.String(conn.Do("JSON.GET", "Session1234"))
	//fmt.Print(r)
	rByte := []byte(r)
	return rByte
	//var bo []tile
	//json.Unmarshal(rByte, &bo)
	//boByte := []byte(bo)
	//fmt.Print(boByte)
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
	//router := mux.NewRouter()

	//var boardPlacements []tile
	//json.Unmarshal([]byte(byteValue), &boardPlacements)

	//fmt.Println(boardPlacements)

	fmt.Println("Server starting")

	fs := http.FileServer(http.Dir("dist"))
	http.Handle("/", cors(fs))
	http.Handle("/chess", cors(fs))
	http.HandleFunc("/board", doBoard)
	http.Handle("/assets", http.FileServer(http.Dir("src")))

	getBoard()

	if err := http.ListenAndServe(":8082", nil); err != nil {
		log.Fatal(err)
	}

}
