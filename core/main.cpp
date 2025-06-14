#include <iostream>
#include <csignal>
#include <atomic>
#include <chrono>
#include <thread>
#include <cstdlib>
#include <sstream>

#ifdef _WIN32
#include <windows.h>
#else
#include <sys/types.h>
#include <signal.h>
#include <unistd.h>
#endif

std::atomic<bool> running(true);

#ifdef _WIN32
DWORD go_server_pid = 0;
#else
pid_t go_server_pid = -1;
#endif

void notify_go_server() {
#ifdef _WIN32
    if (go_server_pid != 0) {
        std::cout << "Notifying Go server (PID " << go_server_pid << ") about shutdown." << std::endl;
        // Open process handle
        HANDLE hProcess = OpenProcess(PROCESS_TERMINATE, FALSE, go_server_pid);
        if (hProcess != NULL) {
            // Send CTRL_BREAK_EVENT to console process group
            // Note: This requires the processes to be in the same console group
            // Alternatively, terminate the process forcibly:
            if (!GenerateConsoleCtrlEvent(CTRL_BREAK_EVENT, go_server_pid)) {
                std::cout << "Failed to send CTRL_BREAK_EVENT." << std::endl;
            }
            CloseHandle(hProcess);
        } else {
            std::cout << "Failed to open Go server process handle." << std::endl;
        }
    } else {
        std::cout << "Go server PID not set, cannot notify." << std::endl;
    }
#else
    if (go_server_pid > 0) {
        std::cout << "Notifying Go server (PID " << go_server_pid << ") about shutdown." << std::endl;
        kill(go_server_pid, SIGTERM);
    } else {
        std::cout << "Go server PID not set, cannot notify." << std::endl;
    }
#endif
}

void signal_handler(int signal) {
    std::cout << "Signal " << signal << " received, shutting down..." << std::endl;
    running = false;
    notify_go_server();
}

int main(int argc, char* argv[]) {
    if (argc > 1) {
#ifdef _WIN32
        go_server_pid = std::stoul(argv[1]);
#else
        std::istringstream iss(argv[1]);
        iss >> go_server_pid;
#endif
        std::cout << "Received Go server PID: " << go_server_pid << std::endl;
    } else {
        std::cout << "No Go server PID provided." << std::endl;
    }

    std::signal(SIGINT, signal_handler);
    std::signal(SIGTERM, signal_handler);

    std::cout << "Core broker started. Press Ctrl+C to stop." << std::endl;

    while (running) {
        // Main broker loop placeholder
        std::this_thread::sleep_for(std::chrono::seconds(1));
    }

    std::cout << "Core broker stopped gracefully." << std::endl;
    return 0;
}
