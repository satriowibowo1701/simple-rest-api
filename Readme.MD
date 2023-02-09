# Simple REST API

Membangun Backend REST API menggunakan Go


## Software yang harus diinstall
- Go

```
# Routes
GET      /api/users
GET     /api/user/{Name}

```

## Konfigurasi

```
Jika
# Setup Environment Variable
Note : Ganti PORT dengan PORT yang ingin dijalankan tanpa tanda kurung
-Windows : Buka CMD lalu ketikan "SET PORT=(PORT)"
-Linux/Unix : Buka CMD lalu ketikan "EXPORT PORT=(PORT)"

# Run
go run main.go

```

## Contoh penggunaan
- Kamu bisa menggunakan POSTMAN untuk mencoba Endpoints yang tersedia, misalkan untuk mencari semua data user menggunakan METHOD GET dengan Endpoint http://localhost:(port)/api/users dan menambahkan header  X-API-Key dengan value satrio  
- Send Request
