package main

const (
	HOST = "localhost"
	PORT = "1766"
)

func main() {
	db := initDB()
	var controllers = DBC(&db)
	TCP(HOST, PORT, controllers)
}
