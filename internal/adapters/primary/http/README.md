# HTTP Adapters (Primary/Input)

This directory contains HTTP handlers that receive requests from external sources.

## Purpose
- Handle HTTP requests using Fiber framework
- Transform HTTP requests into domain operations
- Call appropriate services (input ports)
- Transform domain responses into HTTP responses

## Structure
```
http/
├── handlers/       # HTTP handlers for each domain
│   ├── user_handler.go
│   ├── product_handler.go
│   └── order_handler.go
├── middleware/     # Custom middleware
└── dto/           # Data Transfer Objects for HTTP
```
