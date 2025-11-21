#!/bin/bash

# Script to start backend in background

PID_FILE=".backend.pid"
LOG_FILE="backend.log"

# Check if backend is already running
if [ -f "$PID_FILE" ]; then
    PID=$(cat "$PID_FILE")
    if ps -p "$PID" > /dev/null 2>&1; then
        echo "Backend is already running (PID: $PID)"
        exit 0
    else
        rm -f "$PID_FILE"
    fi
fi

echo "Starting backend server..."

# Start backend in background and save PID
nohup go run cmd/api/main.go > "$LOG_FILE" 2>&1 &
BACKEND_PID=$!

# Save PID to file
echo "$BACKEND_PID" > "$PID_FILE"

# Wait a moment to check if it started successfully
sleep 2

if ps -p "$BACKEND_PID" > /dev/null 2>&1; then
    echo "✅ Backend started successfully (PID: $BACKEND_PID)"
    echo "📝 Logs: tail -f $LOG_FILE"
    echo "🌐 API: http://localhost:8080"
else
    echo "❌ Failed to start backend"
    rm -f "$PID_FILE"
    exit 1
fi
