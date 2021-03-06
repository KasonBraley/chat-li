package socket

// Hub maintains the set of active clients and broadcasts messages to the
// clients.
type Hub struct {
	broadcast  chan []byte
	register   chan *Client
	unregister chan *Client
	clients    map[*Client]bool
	rooms      map[*Room]bool
}

// Create and return a new Hub instance.
func NewHub() *Hub {
	return &Hub{
		broadcast:  make(chan []byte),
		register:   make(chan *Client),
		unregister: make(chan *Client),
		clients:    make(map[*Client]bool),
		rooms:      make(map[*Room]bool),
	}
}

func (h *Hub) Run() {
	for {
		select {
		case client := <-h.register:
			h.clients[client] = true
		case client := <-h.unregister:
			if _, ok := h.clients[client]; ok {
				delete(h.clients, client)
				close(client.send)
			}
		case message := <-h.broadcast:
			for client := range h.clients {
				select {
				case client.send <- message:
				default:
					close(client.send)
					delete(h.clients, client)
				}
			}
		}
	}
}

// Creates a new Room, runs the Room, and adds the Room to the Hub.
func (h *Hub) createRoom(name string, private bool) *Room {
	room := newRoom(name, private)
	go room.runRoom()
	h.rooms[room] = true

	return room
}

func (h *Hub) findRoomByName(name string) *Room {
	var foundRoom *Room
	for room := range h.rooms {
		if room.Name == name {
			foundRoom = room
			break
		}
	}

	return foundRoom
}

func (h *Hub) findRoomByID(ID string) *Room {
	var foundRoom *Room
	for room := range h.rooms {
		if room.ID.String() == ID {
			foundRoom = room
			break
		}
	}

	return foundRoom
}

// Finds one Client matching an ID and returns the Client if found.
func (h *Hub) findClientByID(ID string) *Client {
	var foundClient *Client
	for client := range h.clients {
		if client.ID.String() == ID {
			foundClient = client
			break
		}
	}

	return foundClient
}

// Finds multiple Clients matching an ID and returns the a slice of Clients if found.
func (h *Hub) findClientsByID(ID string) []*Client {
	var foundClients []*Client
	for client := range h.clients {
		if client.ID.String() == ID {
			foundClients = append(foundClients, client)
		}
	}

	return foundClients
}

func (h *Hub) listOnlineClients(client *Client) {
	for existingClient := range h.clients {
		message := &Message{
			Action: UserJoinedAction,
			Sender: existingClient,
		}
		client.send <- message.encode()
	}
}

// func (h *Hub) findUserByID(ID string) models.User {
// 	var foundUser models.User
// 	for _, client := range h.users {
// 		if client.GetId() == ID {
// 			foundUser = client
// 			break
// 		}
// 	}
//
// 	return foundUser
// }

// func (h *Hub) notifyClientLeft(client *Client) {
// 	message := &Message{
// 		Action: UserLeftAction,
// 		Sender: client,
// 	}
//
// 	h.broadcastToClients(message.encode())
// }
