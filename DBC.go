package main

import (
	"fmt"
	"net"
)

type Controller func(net.Conn, []byte)

func parseString(query string, cursor *int, delim rune) string {
	stack := []rune{}

	*cursor++
	for ; *cursor < len(query); *cursor++ {
		char := rune(query[*cursor])

		if char == delim {
			*cursor++
			return string(stack)
		}
		stack = append(stack, char)
	}

	return ""
}

func createCtrl(database *Database) Controller {
	return func(conn net.Conn, buff []byte) {
		// Data Assertion
		var assertData = func(str string) string {
			if str == "" {
				conn.Write([]byte("Invalid query"))
				conn.Close()
			}
			return str
		}

		query := string(buff) + " "

		var parse func(i *int) string
		parse = func(i *int) string {
			if *i != 0 {
				*i++
			}

			stack := []rune{}
			for ; *i <= len(query); *i++ {
				char := rune(query[*i])

				if char != ' ' {
					if char == '"' || char == rune("'"[0]) {
						return parseString(query, i, char)
					}

					stack = append(stack, char)
				} else {
					switch string(stack) {
					case "add":
						collection := assertData(parse(i))
						key := assertData(parse(i))
						value := assertData(parse(i))
						if collection == "" || key == "" || value == "" {
							break
						}
						database.add(collection, key, value)
						return "Success"

					case "get":
						collection := assertData(parse(i))
						key := assertData(parse(i))
						if key == "" {
							break
						}
						return database.get(collection, key)

					default:
						return string(stack)
					}
				}
			}

			return ""
		}

		cursor := 0
		conn.Write([]byte(parse(&cursor)))
	}
}

func globalCtrl(db *DB, controllers *map[string]Controller) Controller {
	return func(conn net.Conn, buff []byte) {
		var assertData = func(str string) string {
			if str == "" {
				conn.Write([]byte("Invalid query"))
				conn.Close()
			}
			return str
		}

		query := string(buff) + " "

		var parse func(i *int) string
		parse = func(i *int) string {
			if *i != 0 {
				*i++
			}

			stack := []rune{}
			for ; *i <= len(query); *i++ {
				char := rune(query[*i])

				if char != ' ' {
					if char == '"' || char == rune("'"[0]) {
						return parseString(query, i, char)
					}

					stack = append(stack, char)
				} else {
					switch string(stack) {
					case "createDB":
						database := assertData(parse(i))
						collection := assertData(parse(i))
						if database == "" || collection == "" {
							break
						}
						db.createDB(database, collection)

						(*controllers)[database] = createCtrl(&db.databases[db.names[database]])

						return fmt.Sprintf("Created New Database: %s\nCreated New Collection: %s\n", database, collection)

					default:
						return string(stack)
					}
				}
			}
			return ""
		}

		cursor := 0
		conn.Write([]byte(parse(&cursor)))
	}
}

func DBC(db *DB) *map[string]Controller {
	controllers := make(map[string]Controller)

	controllers["NSR"] = globalCtrl(db, &controllers)

	for _, database := range db.databases {
		controllers[database.name] = createCtrl(&database)
	}

	return &controllers
}
