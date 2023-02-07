package main

func main() {
	fw := FileWriter{}
	logger := NewLogger(fw)
	logger.Log("Sample message")
}
