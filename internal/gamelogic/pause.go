package gamelogic

import (
	"fmt"

	"github.com/sirtaylor88/learn-pub-sub-starter/internal/routing"
)

func (gs *GameState) HandlePause(ps routing.PlayingState) {
	defer fmt.Println("------------------------")
	fmt.Println()
	if ps.IsPaused {
		fmt.Println("==== Pause Detected ====")
		gs.pauseGame()
	} else {
		fmt.Println("==== Resume Detected ====")
		gs.resumeGame()
	}
}
