package game_logic

import (
	"encoding/json"
	"log"
	"math/rand"
	"net/http"
	"time"

	"github.com/gorilla/websocket"
	"github.com/triberraar/go-battleship/internal/messages"
)

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	}} // use default options

type Client struct {
	Conn       *websocket.Conn
	Battleship *Battleship
	Send       chan interface{}
}

func (c *Client) ReadPump() {
	defer c.Conn.Close()
	for {
		_, message, err := c.Conn.ReadMessage()
		if err != nil {
			log.Print("error: %v", err)
		}
		c.Battleship.messages <- message
	}
}

func (c *Client) WritePump() {
	for {
		select {
		case message := <-c.Send:
			c.Conn.WriteJSON(message)
		}
	}
}

func ServeWs(w http.ResponseWriter, r *http.Request) {
	c, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Print("upgrade:", err)
		return
	}
	client := &Client{Conn: c}
	battleship := NewBattleship(client)

	go client.WritePump()
	go client.ReadPump()
	go battleship.Run()
}

type ship struct {
	x        int
	y        int
	size     int
	vertical bool
	hits     int
}

// func newShip4() ship {
// 	return ship{
// 		size: 4,
// 	}
// }

// func newShip3() ship {
// 	return ship{
// 		size: 3,
// 	}
// }

// func newShip2() ship {
// 	return ship{
// 		size: 2,
// 	}
// }

// func newShip1() ship {
// 	return ship{
// 		size: 1,
// 	}
// }

func newShip(size int) ship {
	return ship{
		size: size,
	}
}

func (s *ship) hit() {
	s.hits++
}

func (s *ship) isDestroyed() bool {
	return s.hits == s.size
}

type tile struct {
	status string
	ship   *ship
}

func (t *tile) hasShip() bool {
	return t.ship != nil
}

type Battleship struct {
	messages chan []byte
	client   *Client

	board     [][]tile
	dimension int
	ships     [6]ship
	victory   bool
}

func (bs *Battleship) Run() {
	for {
		select {
		case message := <-bs.messages:
			bm := messages.BaseMessage{}
			json.Unmarshal(message, &bm)
			if bm.Type == "PLAY" {
				bs.newBoard()
				var shipSizes [len(bs.ships)]int
				for i := 0; i < len(bs.ships); i++ {
					shipSizes[i] = bs.ships[i].size
				}
				bs.client.Send <- messages.NewBoardMessage(shipSizes[:])
			} else if bm.Type == "FIRE" && !bs.victory {
				fm := messages.FireMessage{}
				json.Unmarshal(message, &fm)
				if bs.board[fm.Coordinate.X][fm.Coordinate.Y].status == "fired" {
					continue
				}
				if bs.board[fm.Coordinate.X][fm.Coordinate.Y].hasShip() {

					bs.board[fm.Coordinate.X][fm.Coordinate.Y].status = "fired"
					bs.board[fm.Coordinate.X][fm.Coordinate.Y].ship.hit()
					if bs.board[fm.Coordinate.X][fm.Coordinate.Y].ship.isDestroyed() {
						coordinate := messages.Coordinate{X: bs.board[fm.Coordinate.X][fm.Coordinate.Y].ship.x, Y: bs.board[fm.Coordinate.X][fm.Coordinate.Y].ship.y}
						bs.client.Send <- messages.NewShipDestroyedMessage(coordinate, bs.board[fm.Coordinate.X][fm.Coordinate.Y].ship.size, bs.board[fm.Coordinate.X][fm.Coordinate.Y].ship.vertical)
					} else {
						bs.client.Send <- messages.NewHitMessage(fm.Coordinate)
					}
					if bs.hasVictory() {
						bs.victory = true
						bs.client.Send <- messages.NewVictoryMessage()
						// select {
						// case <-time.After(10 * time.Second):
						// 	bs.newBoard()
						// 	c.WriteJSON(messages.NewRestartMessage())
						// }
					}
				} else {
					bs.client.Send <- messages.NewMissMessage(fm.Coordinate)
					bs.board[fm.Coordinate.X][fm.Coordinate.Y].status = "fired"
				}
			}
		}
	}
}

func (bs *Battleship) sendMessage() {

}

func NewBattleship(client *Client) *Battleship {
	return &Battleship{
		dimension: 10,
		victory:   false,
		messages:  make(chan []byte),
		client:    client,
	}
}

// func RunBattleship(c *websocket.Conn) {
// 	bs := battleship{c: c, dimension: 10, victory: false}

// 	for {
// 		bm := messages.BaseMessage{}
// 		_, message, _ := c.ReadMessage()
// 		json.Unmarshal(message, &bm)
// 		if bm.Type == "PLAY" {
// 			bs.newBoard()
// 			var shipSizes [len(bs.ships)]int
// 			for i := 0; i < len(bs.ships); i++ {
// 				shipSizes[i] = bs.ships[i].size
// 			}
// 			c.WriteJSON(messages.NewBoardMessage(shipSizes[:]))
// 		} else if bm.Type == "FIRE" && !bs.victory {
// 			fm := messages.FireMessage{}
// 			json.Unmarshal(message, &fm)
// 			if bs.board[fm.Coordinate.X][fm.Coordinate.Y].status == "fired" {
// 				continue
// 			}
// 			if bs.board[fm.Coordinate.X][fm.Coordinate.Y].hasShip() {

// 				bs.board[fm.Coordinate.X][fm.Coordinate.Y].status = "fired"
// 				bs.board[fm.Coordinate.X][fm.Coordinate.Y].ship.hit()
// 				if bs.board[fm.Coordinate.X][fm.Coordinate.Y].ship.isDestroyed() {
// 					coordinate := messages.Coordinate{X: bs.board[fm.Coordinate.X][fm.Coordinate.Y].ship.x, Y: bs.board[fm.Coordinate.X][fm.Coordinate.Y].ship.y}
// 					c.WriteJSON(messages.NewShipDestroyedMessage(coordinate, bs.board[fm.Coordinate.X][fm.Coordinate.Y].ship.size, bs.board[fm.Coordinate.X][fm.Coordinate.Y].ship.vertical))
// 				} else {
// 					c.WriteJSON(messages.NewHitMessage(fm.Coordinate))
// 				}
// 				if bs.hasVictory() {
// 					bs.victory = true
// 					c.WriteJSON(messages.NewVictoryMessage())
// 					// select {
// 					// case <-time.After(10 * time.Second):
// 					// 	bs.newBoard()
// 					// 	c.WriteJSON(messages.NewRestartMessage())
// 					// }
// 				}
// 			} else {
// 				c.WriteJSON(messages.NewMissMessage(fm.Coordinate))
// 				bs.board[fm.Coordinate.X][fm.Coordinate.Y].status = "fired"
// 			}
// 		}
// 	}
// }

func (b *Battleship) newBoard() {
	b.board = make([][]tile, b.dimension)
	for i := 0; i < b.dimension; i++ {
		b.board[i] = make([]tile, b.dimension)
	}
	for i := 0; i < b.dimension; i++ {
		for j := 0; j < b.dimension; j++ {
			b.board[i][j].status = "sea"
		}
	}

	for i := 0; i < len(b.ships); i++ {
		s1 := rand.NewSource(time.Now().UnixNano())
		r1 := rand.New(s1)
		size := r1.Intn(3)
		b.ships[i] = newShip(size + 2)
	}

	b.victory = false

	for i := 0; i < len(b.ships); i++ {
		b.generateShip(&b.ships[i])
	}
	log.Print("Generated ships \n")
	for i := 0; i < len(b.ships); i++ {
		log.Printf("ship %+v\n", b.ships[i])
	}
}

func (b *Battleship) hasVictory() bool {
	result := true
	for _, s := range b.ships {
		result = result && s.isDestroyed()
	}
	return result
}

func (b *Battleship) generateShip(s *ship) {
	vertical, x, y := randomPlace(b.dimension, s.size)

	for !b.isValidPlacement(s.size, vertical, x, y) {
		vertical, x, y = randomPlace(b.dimension, s.size)
	}
	s.x = x
	s.y = y
	s.vertical = vertical
	if vertical {
		for i := 0; i < s.size; i++ {
			b.board[x][y+i].status = "ship"
			b.board[x][y+i].ship = s
		}
	} else {
		for i := 0; i < s.size; i++ {
			b.board[x+i][y].status = "ship"
			b.board[x+i][y].ship = s
		}

	}
}

func randomPlace(dimension int, size int) (bool, int, int) {
	s1 := rand.NewSource(time.Now().UnixNano())
	r1 := rand.New(s1)
	vertical := r1.Intn(100) < 50
	var x, y int
	if vertical {
		x = r1.Intn(dimension)
		y = r1.Intn(dimension - size)
	} else {
		x = r1.Intn(dimension - size)
		y = r1.Intn(dimension)
	}

	return vertical, x, y
}

func (b *Battleship) isValidPlacement(size int, vertical bool, x int, y int) bool {
	if vertical {
		if y-1 >= 0 && b.board[x][y-1].status != "sea" {
			return false
		}
		for i := 0; i < size; i++ {
			if b.board[x][y+i].status != "sea" {
				return false
			}
			if x-1 >= 0 && b.board[x-1][y+i].status != "sea" {
				return false
			}
			if x+1 < b.dimension && b.board[x+1][y+i].status != "sea" {
				return false
			}
		}
		if y+size < b.dimension && b.board[x][y+size].status != "sea" {
			return false
		}
	} else {
		if x-1 >= 0 && b.board[x-1][y].status != "sea" {
			return false
		}
		for i := 0; i < size; i++ {
			if b.board[x+i][y].status != "sea" {
				return false
			}
			if y-1 > 0 && b.board[x+i][y-1].status != "sea" {
				return false
			}
			if y+1 < b.dimension && b.board[x+i][y+1].status != "sea" {
				return false
			}
		}
		if x+size < b.dimension && b.board[x+size][y].status != "sea" {
			return false
		}

	}
	return true
}
