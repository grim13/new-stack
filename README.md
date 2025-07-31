
# Go API & Next.js Full-Stack Project

Selamat datang di repositori proyek full-stack ini. Repositori ini berisi kode untuk API backend yang dibuat dengan **Go (Gin)** dan aplikasi frontend yang dibuat dengan **Next.js** dan **Shadcn UI**.

Proyek ini dirancang sebagai fondasi yang kokoh dan dapat diskalakan untuk aplikasi web modern, dengan fokus pada praktik terbaik seperti dependency injection (repository pattern) di backend dan otentikasi yang aman di frontend.

---

## ✨ Fitur Utama

### Backend (Go)

- Arsitektur RESTful API dengan Gin.
- Otentikasi aman menggunakan JWT dengan algoritma RS256.
- Login bisa menggunakan Username atau Email.
- Sistem peran (roles) pengguna yang dapat diperluas.
- Interaksi database yang bersih menggunakan GORM dan Repository Pattern.
- Struktur proyek yang modular dan rapi (routes, handlers, middleware, dll).
- Hot-Reloading untuk pengembangan cepat dengan Air.

### Frontend (Next.js)

- UI modern dan aksesibel dibangun dengan Shadcn UI dan Tailwind CSS.
- Manajemen otentikasi sisi klien dan server dengan Next-Auth.
- Halaman yang dilindungi (protected routes) dan halaman publik.
- Struktur berbasis komponen yang mudah dikelola.

---

## 🛠️ Tumpukan Teknologi (Tech Stack)

| Bagian    | Teknologi                                                                 |
|-----------|--------------------------------------------------------------------------|
| Backend   | Go, Gin, GORM, PostgreSQL, golang-jwt/jwt, bcrypt, Air                   |
| Frontend  | Next.js, React, TypeScript, Tailwind CSS, Shadcn UI, Next-Auth, Zod      |
| Database  | PostgreSQL                                                               |

---

## 🚀 Memulai Proyek

### Prasyarat

Pastikan Anda telah menginstal perangkat lunak berikut di mesin Anda:

- Go (versi 1.18+ direkomendasikan)
- PostgreSQL
- Node.js (versi 18.0+)
- Air (`go install github.com/cosmtrek/air@latest`)
- openssl (biasanya sudah terinstal di macOS/Linux)

---

### 1. Backend Setup

**Clone repositori:**

```bash
git clone [URL_REPOSITORI_ANDA]
cd [NAMA_FOLDER_PROYEK]
```

**Konfigurasi Environment Backend:**

- Salin file `.env.example` menjadi `.env`.
- Sesuaikan nilai variabel di dalam file `.env`, terutama koneksi database Anda.

Contoh `.env`:

```env
# Database Connection
DB_DSN="host=localhost user=postgres password=yourpassword dbname=gogin port=5432 sslmode=disable"

# Server Port
SERVER_PORT=8080

# JWT Key Files
JWT_PRIVATE_KEY_FILE=private.pem
JWT_PUBLIC_KEY_FILE=public.pem
```

**Hasilkan Kunci JWT (RS256):**

Jalankan perintah ini di terminal untuk membuat file `private.pem` dan `public.pem`:

```bash
# Hasilkan Private Key
openssl genrsa -out private.pem 2048

# Ekstrak Public Key
openssl rsa -in private.pem -pubout -out public.pem
```

**Instal dependensi Go:**

```bash
go mod tidy
```

**Jalankan Server Backend:**

Gunakan Air untuk menjalankan server dengan hot-reloading:

```bash
air
```

Server backend sekarang berjalan di [http://localhost:8080](http://localhost:8080).

---

### 2. Frontend Setup

- Pindah ke direktori frontend (asumsi berada di dalam folder `frontend`).

**Konfigurasi Environment Frontend:**

- Salin file `.env.local.example` menjadi `.env.local`.
- Isi variabel yang diperlukan. `NEXTAUTH_SECRET` bisa dihasilkan dengan `openssl rand -base64 32`.

Contoh `.env.local`:

```env
# URL API backend Anda
NEXT_PUBLIC_API_URL=http://localhost:8080/api

# Secret untuk Next-Auth
NEXTAUTH_URL=http://localhost:3000
NEXTAUTH_SECRET=isisendiri_dengan_kunci_rahasia_anda
```

**Instal dependensi Node.js:**

```bash
npm install
# atau
yarn install
```

**Jalankan Server Frontend:**

```bash
npm run dev
# atau
yarn dev
```

Aplikasi frontend sekarang berjalan di [http://localhost:3000](http://localhost:3000).

---

## 📚 Endpoints API

| Method | Endpoint             | Deskripsi                           | Memerlukan Auth      | Body Contoh                                                      |
|--------|----------------------|-------------------------------------|----------------------|------------------------------------------------------------------|
| POST   | `/api/auth/register` | Mendaftarkan pengguna baru.         | Tidak                | `{"name": "...", "username": "...", "email": "...", "password": "..."}` |
| POST   | `/api/auth/login`    | Login pengguna dengan email/username| Tidak                | `{"identifier": "...", "password": "..."}`                       |
| GET    | `/api/users/profile` | Mendapatkan profil pengguna yang login | Ya (Bearer Token) | -                                                                |

---

## 📜 Lisensi

Proyek ini dilisensikan di bawah Lisensi MIT. Lihat file [LICENSE](LICENSE) untuk detail lebih lanjut.
