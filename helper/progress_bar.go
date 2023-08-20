package helper

import (
	"fmt"
	"math"
	"strings"
)

func PrintProgressBar(iteration, total int, prefix, suffix string, length int, fill string) {
	percent := float64(iteration) / float64(total)
	filledLength := int(length * iteration / total)
	end := ">"

	if iteration == total {
		end = "="
	}
	bar := strings.Repeat(fill, filledLength) + end + strings.Repeat("-", (length-filledLength))
	fmt.Printf("\r%s [%s] %d%% %s", prefix, bar, int(math.Floor(percent * 100)), suffix)
	if iteration == total {
		fmt.Println()
	}
}
