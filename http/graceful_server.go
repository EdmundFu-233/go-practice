package main

import (
    "context"
    "fmt"
    "net/http"
    "os/signal"
    "syscall"
)

func main() {
    ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
    defer stop()

    mux := http.NewServeMux()
    mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        fmt.Fprintf(w, "Hello, World!")
    })

    server := &http.Server{Addr: ":8080", Handler: mux}

    go func() {
        <-ctx.Done()
        fmt.Println("Shutting down...")
        server.Shutdown(context.Background())
    }()

    fmt.Println("Server on :8080")
    if err := server.ListenAndServe(); err != http.ErrServerClosed {
        fmt.Println("Error:", err)
    }
    fmt.Println("Server stopped gracefully")
}
