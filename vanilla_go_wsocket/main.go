package main

import (
	"log"
	"net/http"
)

func main(){
	setupAPI()
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func setupAPI(){
	manager := NewManager()
	http.Handle("/", http.FileServer(http.Dir("./frontend")))
	http.HandleFunc("/ws", manager.serveWS)
}


// Manager
// var (
// 	websocketUpgrader = websocket.Upgrader{
// 		ReadBufferSize: 1024,
// 		WriteBufferSize: 1024,
// 	}
// )

// type Manager struct {

// }

// func NewManager() *Manager {
// 	return &Manager{}
// }

// func (m *Manager) serveWS(w http.ResponseWriter, r *http.Request){
// 	log.Println("new connection")
// 	conn, err := websocketUpgrader.Upgrade(w, r, nil)
// 	if err != nil{
// 		log.Println(err)
// 		return
// 	}
// 	conn.Close()
// }


