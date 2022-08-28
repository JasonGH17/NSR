package main

const (
	HOST = "localhost"
	PORT = "1766"
)

func main() {
	db := DB{make(map[string]string)}
	var controller = DBC(&db)
	TCP(HOST, PORT, controller)
}
