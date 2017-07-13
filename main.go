package main

import (
	"encoding/json"
	"log"
	"net/http"
	"os"

	"github.com/robfig/cron"
)

// Rooms gather all monitoring process
var Rooms []RoomWatcher

// The pool for all the websocket client
var pool *SocketPool

func main() {

	// We grab the webserver port from the environment and if none we set it at
	// 8080 for local testing purpose
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	// We create one roomwatcher per room we need to monitor
	for i := 111; i <= 115; i++ {
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
		body, _ := json.Marshal(w.LastObservation)
		pool.broadcast <- body
	}
}
