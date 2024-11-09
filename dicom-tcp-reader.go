package main

import (
    "context"
    "github.com/traefik/traefik/tree/master/pkg/middlewares"
    "net"
    "strings"
)

type dicomRouter struct{}

func (r *dicomRouter) Serve(ctx context.Context, request *middlewares.Request) {
    conn := request.Conn
    buffer := make([]byte, 256) // Buffer size based on where the AE Title is in the PDU
    n, err := conn.Read(buffer)
    if err != nil {
        // Handle error
        return
    }

    // Assuming AE Title is at bytes 20-28
    calledAET := string(buffer[20:28])
    calledAET = strings.TrimSpace(calledAET)
    calledAET = strings.Trim(calledAET, "\x00")

    // Add the AE Title to the request headers for routing decisions in Traefik
    request.SetHeader("X-Dicom-AET", calledAET)
    request.Next()
}

// Factory function to create middleware instance
func NewDicomRouter() *dicomRouter {
    return &dicomRouter{}
}
