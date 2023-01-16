package actormodel

import "fmt"

/*
 go lang actor model exemple
*/

type Player struct {
	Name string
}

type GameState struct {
	players []*Player

	msgch chan any
}

func NewGameState() *GameState {
	g := &GameState{
		players: []*Player{},
		msgch:   make(chan any, 10),
	}

	go g.receive()

	return g
}

func (g *GameState) Send(msg any) {
	g.msgch <- msg
}

func (g *GameState) receive() {
	for msg := range g.msgch {
		g.handleMessage(msg)
	}
}

func (g *GameState) handleMessage(message any) {
	switch msg := message.(type) {
	case *Player:
		g.addPlayer(msg)
	default:
		fmt.Println("invalida message")
	}
}

func (g *GameState) addPlayer(p *Player) {
	g.players = append(g.players, p)
	fmt.Println("adding player: ", p.Name)
}
