# Unit Testing di Golang

## DEFINISI

### PENGUJIAN PERANGKAT LUNAK
Proses menganalisis suatu item perangkat lunak untuk mendeteksi perbedaan antara kondisi yang ada dan yang diinginkan (yaitu, cacat) serta mengevaluasi fitur dari item perangkat lunak tersebut.

## TUJUAN

### Mencegah Regresi, Meningkatkan Tingkat Kepercayaan dalam Refactoring, Meningkatkan Desain Kode, Dokumentasi Kode, Memperluas Tim

## LEVEL

### UI
(Uji End To End) Menguji interaksi secara keseluruhan melalui Antarmuka Pengguna.

### INTEGRASI
Menguji modul atau sub-sistem tertentu melalui API.

### UNIT
Menguji bagian terkecil yang dapat diuji (satu operasi logis) dari sebuah aplikasi melalui metode.

## FRAMEWORK

### Berdasarkan bahasa pemrograman Anda, Anda HARUS MEMILIH framework pengujian unit.
Framework menyediakan alat dan struktur yang diperlukan untuk menjalankan pengujian secara efisien.

### Daftar framework pengujian unit
[Wikipedia List](https://en.wikipedia.org/wiki/List_of_unit_testing_frameworks): Go: Testify, JS: Mocha, Jest

## STRUKTUR

### Pola Umum:
1. Pusatkan file pengujian Anda di dalam folder `tests`.
2. Simpan file pengujian bersama dengan file produksi.

#### File Pengujian: 
Koleksi Suite Pengujian -> Suite Pengujian: Koleksi Kasus Pengujian -> Fixture Pengujian: Setup & Teardown -> Kasus Pengujian: Input, Proses, Output

## PELAKSANA

### Alat yang menjalankan file pengujian. Gunakan mode watch, pilih pelaksana yang dapat menjalankan dengan cepat.

## MOCKING

### Kasus pengujian Anda perlu INDEPENDEN.

### ANDA TIDAK PERLU MENGUJII:
#### Modul pihak ketiga, Basis data, API pihak ketiga, atau sistem eksternal lainnya, Kelas objek yang telah Anda uji di tempat lain.

### Anda perlu membuat OBJEK PALSU
(n) objek palsu yang mensimulasikan perilaku objek nyata.

## PANGANAN

### Pengembang perlu memastikan apakah mereka telah membuat cukup pengujian.

### Alat pengukuran mencakup analisis kode aplikasi saat pengujian sedang dijalankan.

### LAPORAN PANGANAN:
Fungsi, Pernyataan, Cabang, Baris. FORMAT LAPORAN: CLI, XML, HTML, LCOV

### Gunakan format yang dapat dibaca oleh alat seperti SonarQube (analis kode statis).

## LANGKAH

### BUAT FILE PENGUJIAN BARU DI GO!
#### 1. Nama akhiran library_test.go (misalnya, user_test.go)
#### 2. Lokasi File:
   a. folder yang sama, paket yang sama
   b. folder yang sama, paket yang berbeda

### TULIS FUNGSI PENGUJIAN
#### 1. Nama: Test & Huruf Kapital
#### 2. Harus memiliki tanda tangan `func TestXxx(t *testing.T)`

### PAHAMI APA YANG TELAH ANDA BUAT!

#### LEBIH DARI SATU FUNGSI PENGUJIAN! IKUTI ATURAN PENGUJIAN DI GO.

### JALANKAN PENGUJIAN:
```bash
$ go test ./lib/calculate -cover
```

### JALANKAN DENGAN LAPORAN PANGANAN:
```bash
$ go test ./lib/calculate -coverprofile=cover.out && go tool cover -html=cover.out
```
### Uji Fungsi REST API
Konfigurasi, Kontroler DB, userControllers

### TULIS FUNGSI PENGUJIAN UNTUK ECHO:
- Setup Echo (Metode, Parameter, dll.)
- Panggil fungsi REST API di paket kontroler
- Buat fungsi pengujian TestUserControllers

### JALANKAN DENGAN LAPORAN PANGANAN UNTUK SEMUA PAKET:
```bash
$ go test -v -coverpkg=./... -coverprofile=profile.cov ./.. 
$ go tool cover -func profile.cov
```