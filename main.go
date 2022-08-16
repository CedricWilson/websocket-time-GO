package main

import (
	"flag"
	"log"
	"net/http"
	"time"
	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return  true
   },
}

var addr = flag.String("addr", "localhost:8080", "http service address")

func main() {
	http.HandleFunc("/echo", echo)
	http.HandleFunc("/time", getTime)

	log.Fatal(http.ListenAndServe(*addr, nil))
}

func echo(w http.ResponseWriter, r *http.Request) {
	c, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Print("upgrade:", err)
		return
	}
	defer c.Close()
	for {
		mt, message, err := c.ReadMessage()
		if err != nil {
			log.Println("read:", err)
			break
		}
		log.Printf("recv: %s", message)
		err = c.WriteMessage(mt, message)
		if err != nil {
			log.Println("write:", err)
			break
		}
	}
}
func getTime(w http.ResponseWriter, r *http.Request) {
	c, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Print("upgrade:", err)
		return
	}
	defer c.Close()
	for {

		i := 0

		for i < 10 {
			c.WriteMessage(1, []byte(time.Now().Format("2006-01-02 15:04:05")))
			//
			time.Sleep(1 * time.Second)

			i++
		}

	}
}
