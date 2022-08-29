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
						n1 := assertData(conn, parse(i))
						n2 := assertData(conn, parse(i))
						if n1 == "" || n2 == "" {
							break
						}
						db.add(n1, n2)
						return "Success"

					case "get":
						key := assertData(conn, parse(i))
						if key == "" {
							break
						}
						return db.get(key)

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
