package nrk

import (
	"fmt"
	"strings"
	"time"
)

func PrintResponse(respnonse any, title string) {
	fmt.Println("===============================================")
	fmt.Printf("%s Information:\n", title)
	fmt.Println("===============================================")

	switch r := respnonse.(type) {
	case Repo:
		fmt.Printf("%-12s %s\n", "Name:", r.Name)
		fmt.Printf("%-12s %s\n", "URL:", r.URL)
		fmt.Printf("%-12s %s\n", "Description:", r.Description)
		fmt.Printf("%-12s %s\n", "License:", r.License.Name)
	case Commit:
		msg := r.Commit.Message
		title := strings.SplitN(msg, "\n", 2)[0]
		fmt.Printf("%-12s %s\n", "Author:", r.Commit.Author.Name)
		fmt.Printf("%-12s %s\n", "Date:", r.Commit.Author.Date.Format(time.RFC1123))
		fmt.Printf("%-12s %s\n", "Message:", title)
	default:
		fmt.Println("Unknown response type")
	}

	fmt.Println("===============================================")
}
