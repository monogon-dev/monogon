package cluster

import (
	"fmt"
	"os"
	"strings"
)

func getDMIBoardName() (string, error) {
	b, err := os.ReadFile("/sys/devices/virtual/dmi/id/board_name")
	if err != nil {
		return "", fmt.Errorf("could not read board name: %w", err)
	}
	return strings.TrimRight(string(b), "\n"), nil
}

func isGCPInstance(boardName string) bool {
	return boardName == "Google Compute Engine"
}
