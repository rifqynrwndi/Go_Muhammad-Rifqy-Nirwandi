# Middleware

## Apa itu Middleware?

**Middleware** adalah entitas yang terhubung dengan proses pengolahan permintaan dan respons pada server. Middleware memungkinkan kita untuk mengimplementasikan berbagai fungsionalitas di sekitar permintaan HTTP yang masuk ke server dan respons yang keluar.

## Implementasi Middleware

### Alur Kerja Middleware
HTTP Request -> LoggingMiddleware, SessionMiddleware, AuthMiddleware, CustomMiddleware -> API Server -> LoggingMiddleware, SessionMiddleware, AuthMiddleware, CustomMiddleware -> API Server -> HTTP Response


## Jenis Middleware di Echo

### Contoh Middleware Pihak Ketiga
- Negroni
- Echo
- Interpose
- Alice

### Middleware yang Tersedia di Echo
- **Basic Auth**: Untuk otentikasi dasar.
- **Body Dump**: Mencetak isi body permintaan.
- **Body Limit**: Mengatur batas ukuran body.
- **CORS**: Mengatur Cross-Origin Resource Sharing.
- **CSRF**: Mengamankan aplikasi dari serangan Cross-Site Request Forgery.
- **Casbin Auth**: Mengimplementasikan otorisasi.
- **Gzip**: Mengompresi respons untuk menghemat bandwidth.
- **JWT**: Mengelola autentikasi menggunakan JSON Web Token.
- **Key Auth**: Menggunakan kunci untuk otentikasi.
- **Logger**: Mencatat permintaan untuk analisis.
- **Method Override**: Mengizinkan penggunaan metode HTTP yang diubah.
- **Proxy**: Meneruskan permintaan ke server lain.
- **Recover**: Menangani panic dan mengembalikan kesalahan.
- **Redirect**: Mengalihkan permintaan ke URL lain.
- **Request ID**: Menyisipkan ID unik pada setiap permintaan.
- **Rewrite**: Mengubah URL sebelum diteruskan ke handler.
- **Secure**: Menambahkan header keamanan pada respons.
- **Session**: Mengelola sesi pengguna.
- **Static**: Menyajikan file statis.

## Implementasi Middleware

### 1. Log Middleware
**Log Middleware** mencatat informasi dari permintaan HTTP sebagai jejak untuk analisis.

#### Langkah Implementasi:
- Buat folder `middleware`.
- Implementasikan `LogMiddleware` untuk mencatat detail permintaan.

### 2. Auth Middleware
**Auth Middleware** digunakan untuk identifikasi pengguna dan mengamankan data.

#### Jenis Otentikasi:
- **Basic Authentication**: Memerlukan username dan password dalam header permintaan.
- **JSON Web Token (JWT)**: Menggunakan token untuk otentikasi.

#### Langkah Implementasi:
- Buat `AuthMiddleware`.
- Implementasikan middleware dalam controller yang relevan.

### 3. JWT Middleware
**JWT Middleware** digunakan untuk mengelola autentikasi berbasis token.

#### Langkah Implementasi:
- Buat konstanta untuk kunci JWT.
- Implementasikan JWT Middleware.
- Terapkan middleware pada controller login dan detail.

## Kesimpulan

Middleware adalah komponen penting dalam pengembangan aplikasi web, memungkinkan kita untuk menangani berbagai aspek seperti keamanan, logging, dan pengaturan respons dengan lebih efisien. Dengan menggunakan middleware di Echo, pengembang dapat membuat aplikasi yang lebih terstruktur dan mudah dipelihara.

