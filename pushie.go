package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

// This section sets up variables & objects I'll be using in some functions.

// Creates a Message type.

type Message struct {
	Channel string
	Data    map[string]interface{}
}

// Creates a Websocket type.

type Websocket struct {
	id int
}

// Creates some dummy socket objects.

var Socket1 = Websocket{
	id: 1,
}

var Socket2 = Websocket{
	id: 2,
}

var Socket3 = Websocket{
	id: 3,
}

// Creates send method for Websocket.

func (s Websocket) Send(message []byte) error {
	return nil
}

// This is the lookup function that searches a temporarily hardcoded hash and calls Send for all channels.

func ChannelLookup(currentmap map[string][]Websocket, channel string) []Websocket {

	target_sockets, ok := currentmap[channel]
	if !ok {
		return []Websocket{}
	}

	return target_sockets
}

// This section contains the functions called by each route.

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	//fmt.Fprintf(w, "Hi there, I love cheese!")
	io.WriteString(w, `{"alive": true}`)
}

func PublishHandler(w http.ResponseWriter, req *http.Request) {

	// Creates channels lookup hash. Eventually this will be done dynamically.

	var channelsHash = make(map[string][]Websocket)
	channelsHash["bunno"] = []Websocket{Socket1, Socket2}
	channelsHash["quadrupal"] = []Websocket{Socket1, Socket2, Socket3}
	channelsHash["thisisfine"] = []Websocket{}

	decoder := json.NewDecoder(req.Body)
	for {
		var m Message
		if err := decoder.Decode(&m); err == io.EOF {
			break
		} else if err != nil {
			log.Fatal(err)
		}

		fmt.Println("PublishHandler says: The request is asking for channel ", m.Channel)
		// get sockets for the channel from hash, currently hardcoded

		sockets := ChannelLookup(channelsHash, m.Channel)
		fmt.Println("PublishHandler says: The channel info is ", m.Channel, " ", m.Data)
		fmt.Println("PublishHandler says: The sockets listening on that channel are ", sockets)

	}

}

// Below are objects & functions for PublishHandler until it can read live messages.

func main() {

	r := mux.NewRouter()

	r.HandleFunc("/", HomeHandler)
	r.HandleFunc("/publish", PublishHandler)
	//r.HandleFunc("/user/kick/:uid", KickHandler)
	//r.HandleFunc("/user/logout/:authtoken", LogoutHandler)
	//r.HandleFunc("/user/channel/add/:channel/:uid", AddUserChannelHandler)
	//r.HandleFunc("/user/channel/remove/:channel/:uid", RemoveUserChannelHandler)
	//r.HandleFunc("channel/add/:channel", AddChannelHandler)
	//r.HandleFunc("health/check", HealthCheckHandler)
	//r.HandleFunc("channel/check/:channel", CheckChannelHandler)
	//r.HandleFunc("channel/remove/:channel", RemoveChannelHandler)
	//r.HandleFunc("user/presence-list/:uid/:uidList", SetUserPresenceHandler)
	//r.HandleFunc("debug/toggle", DebugToggleHandler)
	//r.HandleFunc("content/token/users", GetContentTokenHandler)
	//r.HandleFunc("content/token", SetContentTokenHandler)
	//r.HandleFunc("content/token/message", PublishMessageToContentChannelHandler)
	//r.HandleFunc("authtoken/channel/add/:channel/:authToken", AddAuthTokenChannelHandler)
	//r.HandleFunc("authtoken/channel/remove/:channel/:authToken", RemoveAuthTokenChannelHandler)

	http.Handle("/", r)

	http.ListenAndServe(":8080", r)
}
