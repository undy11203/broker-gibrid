package main

import (
    "context"
    "fmt"
    "os"
    "os/exec"
    "os/signal"
    "strconv"
    "syscall"
    "time"
)

func main() {
    ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
    defer stop()

    fmt.Println("Starting broker server...")

    // Get current process PID
    goServerPID := os.Getpid()
    fmt.Printf("Go server PID: %d\n", goServerPID)

    // Start core broker process with Go server PID as argument
    coreCmd := exec.Command("./core_broker", strconv.Itoa(goServerPID))
    coreCmd.Stdout = os.Stdout
    coreCmd.Stderr = os.Stderr

    err := coreCmd.Start()
    if err != nil {
        fmt.Printf("Failed to start core broker: %v\n", err)
        return
    }
    fmt.Printf("Core broker started with PID %d\n", coreCmd.Process.Pid)

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

    // Stop core broker process
    if coreCmd.Process != nil {
        fmt.Println("Stopping core broker process...")
        coreCmd.Process.Signal(syscall.SIGTERM)
        done := make(chan error)
        go func() {
            done <- coreCmd.Wait()
        }()
        select {
        case <-shutdownCtx.Done():
            fmt.Println("Timeout waiting for core broker to stop, killing...")
            coreCmd.Process.Kill()
        case err := <-done:
            if err != nil {
                fmt.Printf("Core broker exited with error: %v\n", err)
            } else {
                fmt.Println("Core broker stopped gracefully")
            }
        }
    }

    fmt.Println("Server stopped gracefully")
}
