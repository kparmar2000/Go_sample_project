package pointer

import "fmt"

type Message struct {
	Recipient string
	Text      string
}

// Don't touch above this line

func SendMessage(m Message) {
	fmt.Printf("To: %v\n", m.Recipient)
	fmt.Printf("Message: %v\n", m.Text)
}

// Don't touch below this line

func Test(recipient string, text string) {
	m := Message{Recipient: recipient, Text: text}
	SendMessage(m)
	fmt.Println("=====================================")
}

func main() {
	Test("Lane", "Textio is getting better everyday!")
	Test("Allan", "This pointer stuff is weird...")
	Test("Tiffany", "What time will you be home for dinner?")
}
