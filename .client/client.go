package main

import (
	"bufio"
	"context"
	"flag"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"

	gRPC "github.com/mbjnitu/diSysExam2021/proto"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

// Same principle as in client. Flags allows for user specific arguments/values
var clientsName = flag.String("name", "default", "Senders name")
var serverPort = flag.String("server", "5400", "Tcp server")

var server gRPC.TemplateClient  //the server
var ServerConn *grpc.ClientConn //the server connection

func main() {
	//parse flag/arguments
	flag.Parse()

	fmt.Println("--- CLIENT APP ---")

	//log to file instead of console
	//f := setLog()
	//defer f.Close()

	//connect to server and close the connection when program closes
	fmt.Println("--- join Server ---")
	ConnectToServer()
	defer ServerConn.Close()

	//start the biding
	parseInput()
}

// connect to server
func ConnectToServer() {

	//dial options
	//the server is not using TLS, so we use insecure credentials
	//(should be fine for local testing but not in the real world)
	var opts []grpc.DialOption
	opts = append(opts, grpc.WithBlock(), grpc.WithTransportCredentials(insecure.NewCredentials()))

	//dial the server, with the flag "server", to get a connection to it
	log.Printf("client %s: Attempts to dial on port %s\n", *clientsName, *serverPort)
	conn, err := grpc.Dial(fmt.Sprintf(":%s", *serverPort), opts...)
	if err != nil {
		log.Printf("Fail to Dial : %v", err)
		return
	}

	// makes a client from the server connection and saves the connection
	// and prints rather or not the connection was is READY
	server = gRPC.NewTemplateClient(conn)
	ServerConn = conn
	log.Println("the connection is: ", conn.GetState().String())
}

func parseInput() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("In order to put or get, type 'put:x,y' or 'get:x'")
	fmt.Println("--------------------")

	//Infinite loop to listen for clients input.
	for {
		fmt.Print("-> ")

		//Read input into var input and any errors into err
		input, err := reader.ReadString('\n')
		if err != nil {
			log.Fatal(err)
		}
		input = strings.TrimSpace(input) //Trim input

		if !conReady(server) {
			log.Printf("Client %s: something was wrong with the connection to the server :(", *clientsName)
			continue
		}

		str := strings.Split(input, ":")[1]

		if strings.Contains(input, "put") {
			key, errK := strconv.Atoi(strings.Split(str, ",")[0])
			val, errV := strconv.Atoi(strings.Split(str, ",")[1])
			if errK == nil && errV == nil {
				Put(int64(key), int64(val))
			} else {
				fmt.Println("PUT: You formatted the command wrongly")
			}

		} else if strings.Contains(input, "get") {
			key, errK := strconv.Atoi(strings.Split(str, ",")[0])
			if errK == nil {
				Get(int64(key))
			} else {
				fmt.Println("GET: You formatted the command wrongly")
			}
		}
	}
}

func Put(key int64, val int64) {
	//create amount type
	kvp := &gRPC.KeyValPair{
		Key: key,
		Val: val,
	}

	//Make gRPC call to server with amount, and recieve acknowlegdement back.
	putAck, err := server.Put(context.Background(), kvp)
	if err != nil {
		log.Printf("Client %s: PUT: no response from the server, attempting to reconnect", *clientsName)
		log.Println(err)
	}

	fmt.Printf("Put: %d\n", putAck.Response)
}

func Get(in_key int64) {
	//create amount type
	key := &gRPC.Key{
		Key: in_key,
	}

	//Make gRPC call to server with amount, and recieve acknowlegdement back.
	getAck, err := server.Get(context.Background(), key)
	if err != nil {
		log.Printf("Client %s: GET: no response from the server, attempting to reconnect", *clientsName)
		log.Println(err)
	}

	fmt.Printf("Get: %d\n", getAck.Response)
}

// Function which returns a true boolean if the connection to the server is ready, and false if it's not.
func conReady(s gRPC.TemplateClient) bool {
	return ServerConn.GetState().String() == "READY"
}

// sets the logger to use a log.txt file instead of the console
func setLog() *os.File {
	f, err := os.OpenFile("log.txt", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalf("error opening file: %v", err)
	}
	log.SetOutput(f)
	return f
}
