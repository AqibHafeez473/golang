package wordcount
import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// WordCount reads a string from user and returns words slice and count
func WordCount() ([]string, int) {
	fmt.Println("Enter a string:")
	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')
	input = strings.TrimSpace(input)

	words := strings.Fields(input)
	wordCount := len(words)

	return words, wordCount
}

// Optional: helper to print results
func PrintWords(words []string, count int) {
	fmt.Println("Words array:", words)
	fmt.Println("Total words:", count)
}



