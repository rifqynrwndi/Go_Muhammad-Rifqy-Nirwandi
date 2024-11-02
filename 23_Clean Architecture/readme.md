# Clean Architecture

## Pengantar

Package `context` di Go berfungsi membawa data, sinyal pembatalan, atau nilai lainnya dalam permintaan API, terutama untuk pengaturan waktu eksekusi, penghentian, atau komunikasi data.

## Implementasi Context

- **Root Context**: Dibuat menggunakan fungsi `background()` sebagai root context untuk operasi lainnya.
- **Context with Value**: Menambahkan data ke dalam root context, sering digunakan di middleware.
- **Context with Cancellation**: Menghentikan request ke database jika pengguna membatalkan permintaan.

## Migrasi dari MVC ke Clean Architecture

### Sebelum Migrasi

Migrasi ke clean architecture bertujuan agar kode lebih modular, scalable, dan mudah dipelihara.

### Langkah Migrasi

Tiga opsi migrasi dari MVC ke clean architecture:
1. Pisahkan dependensi, tetapi pertahankan desain.
2. Pindahkan kode ke layer tertentu.
3. Ubah desain dan pisahkan dependensi.

Pada contoh ini, dipilih opsi kedua dengan memindahkan kode ke layer-layer **usecase** dan **repository**.

### Struktur Kode

- **Controller**: Berinteraksi dengan user (interface layer).
- **Repository**: Koneksi langsung ke database (interface layer).
- **Usecase**: Berisi logika bisnis.

### Hasil

Struktur kode menjadi lebih rapi dan mudah dipelihara.