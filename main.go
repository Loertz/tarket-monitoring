package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/robfig/cron"
)

func main3() {
	t := GetTarketService()
	result, _ := t.GetRoomInfos(111)

	fmt.Print(result)

}

var Rooms []RoomWatcher
var pool *SocketPool

func main() {

	port := os.Getenv("PORT")

	if port == "" {
		port = "8080"
	}

	for i := 111; i <= 116; i++ {
		w := NewRoomWatcher(i)
		Rooms = append(Rooms, w)
	}
	// Create and start the socket pool that will manage all clients
	pool = NewSocketPool()
	go pool.run()
	log.Print("Websocket pool started ===")

	// Create the cron that will be responsible for the rate limiting of
	// the refresh process on the Tarket API
	c := cron.New()
	c.AddFunc("@every 60s", refreshDataFromAPI)
	c.Start()
	log.Print("Cron started ===")

	// Start the webserver that will respond to the interface query and websocket
	fs := http.FileServer(http.Dir("static"))
	http.Handle("/", http.FileServer(http.Dir("./static")))
	http.Handle("/static/", http.StripPrefix("/static/", fs))
	http.HandleFunc("/receive", func(w http.ResponseWriter, r *http.Request) {
		serveWs(pool, w, r)
	})

	log.Print("Web server started ===")
	err := http.ListenAndServe(":"+port, nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}

}

func refreshDataFromAPI() {
	log.Print("Cron runned")
	for _, w := range Rooms {
		w.UpdateDataInStore()
		body, _ := json.Marshal(w)
		pool.broadcast <- body
	}
}

func serveHome(res http.ResponseWriter, req *http.Request) {
}
