package main

import (
    "context"
    "fmt"
    "os"
    "os/signal"
    "syscall"
    "time"
)

func main() {
    ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
    defer stop()

    fmt.Println("Starting broker server...")

    // Run server in a goroutine or main logic here
    go func() {
        // Placeholder for server start logic
        fmt.Println("Server is running...")
        <-ctx.Done()
    }()

    // Wait for interrupt signal
    <-ctx.Done()
    fmt.Println("Shutdown signal received")

    // Graceful shutdown logic here
    shutdownCtx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
    defer cancel()

    // Simulate cleanup work
    select {
    case <-time.After(2 * time.Second):
        fmt.Println("Cleanup completed")
    case <-shutdownCtx.Done():
        fmt.Println("Shutdown timeout reached")
    }

    fmt.Println("Server stopped gracefully")
}
