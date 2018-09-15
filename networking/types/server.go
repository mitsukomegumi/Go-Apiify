package types

import (
	"errors"
	"fmt"
	"net/http"
	"reflect"
	"strings"

	"github.com/mitsukomegumi/Go-Rpcify/types"
)

// Server - networking type used for rpc server
type Server struct {
	Endpoint    string             `json:"endpoint"`    // Used for http access to server
	Environment *types.Environment `json:"environment"` // Used for call access
}

// NewServer - initialize new instance of Server struct
func NewServer(endpoint string) *Server {
	env := types.NewEnvironment(endpoint) // Initialize environment

	server := Server{endpoint, env} // Initialize server

	return &server // Return initialized server
}

// StartServer - start serving requests to server
func (server *Server) StartServer() error { // TODO: finished server start method
	if reflect.ValueOf(server).IsNil() { // Check for nil server
		return errors.New("nil server") // Return found error
	}

	http.HandleFunc("/", server.HandleRequest) // Handle request

	http.ListenAndServe(":8080", nil) // Serve requests

	return nil // No error occurred, return nil
}

// HandleRequest - attempt to handle request
func (server *Server) HandleRequest(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("\n-- SERVER -- found request %s\n\n", r.URL.Path) // Log request

	endpoint := strings.Split(r.URL.Path, "/")[len(strings.Split(r.URL.Path, "/"))-1] // Split endpoint

	call, err := server.Environment.SearchCallEndpoints(endpoint) // Query call

	if err != nil { // Check for errors
		stack, err := server.Environment.SearchStackEndpoints(endpoint) // Query stack

		if err != nil { // Check for errors
			fmt.Fprintf(w, err.Error()) // Log error
		}

		server.handleStack(stack, w, r) // Handle stack
	} else {
		server.handleCall(call, w, r) // Handle call
	}
}

func (server *Server) handleCall(call *types.Call, w http.ResponseWriter, r *http.Request) {
	fmt.Printf("running call %s\n", call.Endpoint) // Log running call

	output, err := call.Run() // Run call

	if err != nil { // Check for errors
		fmt.Println(err.Error() + "\n") // Log error
		fmt.Fprintf(w, err.Error())     // Log error
	} else {
		fmt.Println(output + "\n") // Log output
		fmt.Fprintf(w, output)     // Log success
	}
}

func (server *Server) handleStack(stack *types.Stack, w http.ResponseWriter, r *http.Request) {
	fmt.Printf("running stack %s\n", stack.Endpoint) // Log running call

	output, err := stack.Run() // Run stack

	if err != nil { // Check for errors
		fmt.Println(err.Error())    // Log error
		fmt.Fprintf(w, err.Error()) // Log error
	} else {
		fmt.Println(output)    // Log output
		fmt.Fprintf(w, output) // Log success
	}
}
