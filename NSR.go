package main

const (
	HOST = "localhost"
	PORT = "1766"
)

func main() {
	db := initDB()
	var controller = DBC(&db)
	TCP(HOST, PORT, controller)
}
