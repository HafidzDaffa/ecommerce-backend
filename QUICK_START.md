# Quick Start Guide

Panduan cepat untuk menjalankan backend E-Commerce.

## 🚀 Start Backend & Database

Jalankan PostgreSQL dan backend server sekaligus:

```bash
make docker-up
```

Output:
```
🚀 Starting all services...
1️⃣  Starting PostgreSQL...
⏳ Waiting for PostgreSQL to be ready...
2️⃣  Starting backend server...
✅ Backend started successfully (PID: xxxxx)

✅ All services are running!
📊 PostgreSQL: localhost:5432
🌐 Backend API: http://localhost:8080
```

Backend akan berjalan di background. Anda bisa tutup terminal dan backend tetap running.

## 🛑 Stop Backend & Database

Stop semua services:

```bash
make docker-down
```

Output:
```
🛑 Stopping all services...
✅ Backend stopped
✅ All services stopped
```

## 📊 Cek Status Services

Lihat status PostgreSQL dan backend:

```bash
make docker-status
```

Output:
```
📊 Service Status:
===================

PostgreSQL:
  ✅ Running (healthy)

Backend:
  ✅ Running (PID: xxxxx)
```

## 📝 Lihat Logs

**Backend logs:**
```bash
make backend-logs
```

**PostgreSQL logs:**
```bash
make docker-logs
```

## 🔄 Restart Services

Restart PostgreSQL dan backend:

```bash
make docker-restart
```

## 🧪 Test API

Setelah services running, test API:

```bash
# Health check
curl http://localhost:8080/api/health

# API v1
curl http://localhost:8080/api/v1/
```

## 📚 Command Lengkap

| Command | Deskripsi |
|---------|-----------|
| `make docker-up` | Start PostgreSQL + backend |
| `make docker-down` | Stop backend + PostgreSQL |
| `make docker-status` | Lihat status services |
| `make docker-restart` | Restart semua services |
| `make backend-logs` | Lihat backend logs |
| `make docker-logs` | Lihat PostgreSQL logs |
| `make run` | Run backend di foreground (untuk development) |
| `make migrate-up` | Run database migrations |
| `make seed` | Run database seeders |
| `make help` | Lihat semua command |

## 🔧 Development Mode

Jika ingin run backend di foreground (auto-reload saat development):

```bash
# PostgreSQL tetap di Docker
docker compose up -d postgres

# Run backend di terminal (foreground)
make run
```

Tekan `Ctrl+C` untuk stop backend.

## ⚠️ Troubleshooting

**Port 8080 sudah digunakan:**
```bash
# Cek process yang menggunakan port 8080
lsof -i :8080

# Stop backend
make docker-down
```

**PostgreSQL tidak connect:**
```bash
# Restart PostgreSQL
docker compose restart postgres

# Cek logs
make docker-logs
```

**Backend tidak start:**
```bash
# Lihat error di logs
make backend-logs

# Atau lihat file log langsung
cat backend.log
```

## 📍 Endpoint Info

- **Backend API:** http://localhost:8080
- **PostgreSQL:** localhost:5432
- **Health Check:** http://localhost:8080/api/health
- **API v1:** http://localhost:8080/api/v1/

## 🎯 First Time Setup

Jika pertama kali setup:

```bash
# 1. Setup environment
make env

# 2. Start services
make docker-up

# 3. Run migrations
make migrate-up

# 4. Run seeders
make seed

# 5. Test API
curl http://localhost:8080/api/health
```

Done! Backend siap digunakan 🎉
