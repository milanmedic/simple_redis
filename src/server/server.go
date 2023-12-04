package server

import (
	"fmt"
	"io"
	"net"

	"simple_redis.com/m/src/utils"
)

type Server struct {
	listener   net.Listener
	connection net.Conn
}

/*
Initializes a new socket and starts listening to requests.
*/
func CreateServer(host, port string) *Server {
	listener, err := net.Listen("tcp", host+":"+port)
	if err != nil {
		utils.Exit(1, fmt.Sprintf("Error while establishing socket: %s:%s\n", host, port))
	}

	connection, err := listener.Accept()
	if err != nil {
		utils.Exit(1, fmt.Sprintf("Error while accepting requests on: %s:%s\n", host, port))
	}

	fmt.Printf("Started listening to connections on %s:%s\n", host, port)

	return &Server{listener: listener, connection: connection}
}

/*
Reads incoming data from the socket.
*/
func (s *Server) Read(buffer []byte) error {
	_, err := s.connection.Read(buffer)
	if err != nil {
		if err == io.EOF {
			s.close(nil)
		}
		customErr := fmt.Errorf(fmt.Sprintf("Error while reading from the client.\n Message: %s", err.Error()))
		s.close(customErr)
	}

	return nil
}

/*
Writes back data to the client as an array of bytes.
*/
func (s *Server) Write(buffer string) error {
	_, err := s.connection.Write([]byte(buffer))
	if err != nil {
		customErr := fmt.Errorf(fmt.Sprintf("Error while writting back to the client.\n Message:%s", err.Error()))
		s.close(customErr)
	}

	return nil
}

/*
Listener and connection cleanup for the server.
*/
func (s *Server) close(cause error) {
	if cause != nil {
		utils.Exit(1, cause.Error())
	}

	err := s.connection.Close()
	if err != nil {
		utils.Exit(1, fmt.Sprintf("Error while closing the connection.\n Message: %s", err.Error()))
	}

	err = s.listener.Close()
	if err != nil {
		utils.Exit(1, fmt.Sprintf("Error while closing the listener.\n Message: %s", err.Error()))
	}
	utils.Exit(0, "Server shutdown gracefuly")
}
