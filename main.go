package main

import "simple_redis.com/m/src/server"

func main() {
	srv := server.CreateServer("localhost", "6379")

	for {
		buffer := make([]byte, 1024)
		srv.Read(buffer)

		srv.Write("+OK\r\n")
	}
}
