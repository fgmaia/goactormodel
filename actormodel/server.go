package actormodel

type Server struct {
	gameState *GameState
}

func NewServer() *Server {
	return &Server{
		gameState: NewGameState(),
	}
}

func (s *Server) HandleNewPlayer(player *Player) error {
	s.gameState.Send(player)
	return nil
}
