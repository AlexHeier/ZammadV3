package global

import (
	"fmt"
	"strings"
	"time"
)

// logitnes bare en loading screen. må kalles med go forann slik at programmet kan fortsette.

var IsLoading bool

func LoadingScreen() {
	i := 0
	ClearScreen()
	fmt.Print("Loading \033[?25l")
	defer fmt.Print("\033[?25h") //gjemmer couirsen imens den loader
	for IsLoading {
		fmt.Printf("\rLoading %s", strings.Repeat(" ", 3)) // skriver bare over loading linjen
		fmt.Printf("\rLoading %s", strings.Repeat(".", (i%4)))
		time.Sleep(500 * time.Millisecond)
		i++
	}
}
