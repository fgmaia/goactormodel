package actormodel_test

import (
	"fmt"
	"testing"

	"github.com/fgmaia/goactormodel/actormodel"
)

func TestServer(t *testing.T) {
	server := actormodel.NewServer()

	for i := 0; i < 10; i++ {
		player := &actormodel.Player{
			Name: fmt.Sprintf("player_%d", i),
		}

		go server.HandleNewPlayer(player)
	}
}
