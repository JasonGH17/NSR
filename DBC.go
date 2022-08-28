package main

import "net"

func DBC(db *DB) func(net.Conn, []byte) {
	return func(conn net.Conn, buff []byte) {
		query := string(buff)

		var parse func(i *int) string
		parse = func(i *int) string {
			stack := []rune{}
			for ; *i < len(query); *i++ {
				char := rune(query[*i])

				if char != ' ' {
					stack = append(stack, char)
				} else {
					switch string(stack) {
					case "add":
						*i++
						n1 := parse(i)
						*i++
						n2 := parse(i)
						db.add(n1, n2)
						return "Success"

					case "get":
						*i++
						return db.get(parse(i))

					default:
						return string(stack)
					}
					stack = []rune{}
				}
			}

			return ""
		}

		cursor := 0
		conn.Write([]byte(parse(&cursor)))
	}
}
