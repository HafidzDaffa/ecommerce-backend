#!/bin/bash

# Script to stop backend

PID_FILE=".backend.pid"

# Function to kill process and its children
kill_process_tree() {
    local pid=$1
    local children=$(pgrep -P "$pid" 2>/dev/null)
    
    # Kill children first
    for child in $children; do
        kill_process_tree "$child"
    done
    
    # Kill the process itself
    if ps -p "$pid" > /dev/null 2>&1; then
        kill "$pid" 2>/dev/null
    fi
}

if [ ! -f "$PID_FILE" ]; then
    echo "Backend is not running (no PID file found)"
    # Try to kill any running backend process
    pkill -f "cmd/api/main.go" 2>/dev/null && echo "✅ Stopped orphan backend process"
    # Also kill any "main" binary on port 8080
    PORT_PID=$(lsof -ti:8080 2>/dev/null)
    if [ ! -z "$PORT_PID" ]; then
        kill "$PORT_PID" 2>/dev/null && echo "✅ Stopped process on port 8080"
    fi
    exit 0
fi

PID=$(cat "$PID_FILE")

if ps -p "$PID" > /dev/null 2>&1; then
    echo "Stopping backend (PID: $PID)..."
    
    # Kill process tree (parent + children)
    kill_process_tree "$PID"
    
    # Wait for process to stop (max 5 seconds)
    for i in {1..5}; do
        if ! ps -p "$PID" > /dev/null 2>&1; then
            break
        fi
        sleep 1
    done
    
    # Force kill if still running
    if ps -p "$PID" > /dev/null 2>&1; then
        echo "Force stopping backend..."
        kill -9 "$PID" 2>/dev/null
    fi
    
    rm -f "$PID_FILE"
    echo "✅ Backend stopped"
else
    echo "Backend process not found (PID: $PID)"
    rm -f "$PID_FILE"
fi

# Clean up any orphan processes
pkill -f "cmd/api/main.go" 2>/dev/null

# Clean up process on port 8080
PORT_PID=$(lsof -ti:8080 2>/dev/null)
if [ ! -z "$PORT_PID" ]; then
    kill "$PORT_PID" 2>/dev/null
fi
