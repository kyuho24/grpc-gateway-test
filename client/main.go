package main

import (
    "context"
    "flag"
    "log"
    "time"

    "google.golang.org/grpc"
    "google.golang.org/grpc/credentials/insecure"
    pb "grpc-gateway-test/gen/go/yourservice"
)

const (
    defaultName = "Kyuho"
)

var (
    addr = flag.String("addr", "localhost:9090", "the address to connect to")
    name = flag.String("value", defaultName, "Name to greet")
)

func main() {
    flag.Parse()
    // Set up a connection to the server.
    conn, err := grpc.Dial(*addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
    if err != nil {
        log.Fatalf("did not connect: %v", err)
    }
    defer conn.Close()
    c := pb.NewYourServiceClient(conn)

    // Contact the server and print out its response.
    ctx, cancel := context.WithTimeout(context.Background(), time.Second)
    defer cancel()
    r, err := c.Echo(ctx, &pb.StringMessage{Value: *name})
    if err != nil {
        log.Fatalf("could not Echo: %v", err)
    }
    log.Printf("Echo: %s", r.GetValue())
}
