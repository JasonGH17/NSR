package main

const (
	HOST = "localhost"
	PORT = "1766"
)

func main() {
	db := DB{}
	_ = db
	// Add DB commands interface
	TCP(HOST, PORT)
}
