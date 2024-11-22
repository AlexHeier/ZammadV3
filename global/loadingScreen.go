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
	for IsLoading {
		fmt.Printf("\r%s", strings.Repeat(" ", 20)) // skriver bare over loading linjen
		fmt.Printf("\rLoading %s", strings.Repeat(".", (i%4)))
		time.Sleep(500 * time.Millisecond)
		i++
	}
}
