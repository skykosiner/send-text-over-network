package typer

import (
	"fmt"
	"log"
	"os/exec"
)

func TypeMeDaddy(text string) {
	cmd := fmt.Sprintf("xdotool type --delay 100 '%s'", text)

	typeMeDaddy := exec.Command("bash", "-c", cmd)

	_, err := typeMeDaddy.Output()

	if err != nil {
		log.Fatal("Error typing", err)
	}
}
