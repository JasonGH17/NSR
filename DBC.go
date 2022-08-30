package main

import "net"

func assertData(conn net.Conn, str string) string {
	if str == "" {
		conn.Write([]byte("Invalid query"))
		conn.Close()
	}
	return str
}

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

func DBC(db *DB) func(net.Conn, []byte) {
	return func(conn net.Conn, buff []byte) {
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
						collection := assertData(conn, parse(i))
						key := assertData(conn, parse(i))
						value := assertData(conn, parse(i))
						if collection == "" || key == "" || value == "" {
							break
						}
						db.add(collection, key, value)
						return "Success"

					case "get":
						collection := assertData(conn, parse(i))
						key := assertData(conn, parse(i))
						if key == "" {
							break
						}
						return db.get(collection, key)

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
