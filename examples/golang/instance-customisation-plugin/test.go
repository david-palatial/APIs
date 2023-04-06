package main

import(
    "context"
    "log"
    "time"
    //"github.com/golang/protobuf/proto"
    "github.com/scalablepixelstreaming/apis/pkg/customisation"
    "google.golang.org/grpc"
    //"unicode/utf8"
    //b64 "encoding/base64"
)

func main() {
	// The default port for the custom plugin for kubernetes to listen on

	// This port is important as it will be the port you refer to when adding this as a plugin
	// to the Scalable Pixel Streaming Application (e.g. my-plugin.default.svc.cluster.local:55774)

	const PLUGIN_PORT = 55774

	conn, err := grpc.Dial("10.145.106.71:55774", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Failed to connect: %v", err)
	}
	defer conn.Close()

    	request := &customisation.UpdateRuntimeOptionsRequest{
        	Instance: "sps-application-e5ee3447-3fd1-4d47-b9dc-f1648cb8ea77-55967wghfs",
        	RuntimeOptions: &customisation.RuntimeOptions{
            		Resolution: &customisation.RuntimeOptions_Resolution{
                		X: 720,
                		Y: 855,
            		},
        	},
    	}

	// Create a client instance using the connection.
	client := customisation.NewInstanceCustomisationPluginClient(conn)

	// Call the UpdateRuntimeOptions function of the client and pass the request object as an argument.
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	response, err := client.UpdateRuntimeOptions(ctx, request)
	if err != nil {
		log.Fatalf("Failed to update runtime options: %v", err)
	}

	log.Printf("Response: %v", response)
}
