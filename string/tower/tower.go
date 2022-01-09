package tower

import "fmt"

func BuildTower(height int) []string {
	lines := make([]string, 0)
	buildTop(height, &lines)

	buildBody(height, &lines)
	buildBottom(height, &lines)
	for _, line := range lines {
		fmt.Println(line)
	}
	return lines
}

func buildBottom(height int, lines *[]string) []string {
	*lines = append(*lines, repeat("-", (height+1)*2+1))
	return *lines
}

func buildBody(height int, lines *[]string) []string {
	for i := height - 1; i >= 0; i-- {
		value := height - i
		padding := i + 1
		line := repeat(" ", padding) + repeat("*", value) + "|" + repeat("*", value)
		*lines = append(*lines, line)
	}
	return *lines
}

func buildTop(height int, lines *[]string) []string {
	*lines = append(*lines, repeat(" ", height+1)+"|")
	return *lines
}

func repeat(character string, length int) string {
	result := ""
	for i := 0; i < length; i++ {
		result += character
	}
	return result
}
