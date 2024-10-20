# RESTful API

## Apa itu API?
API (Application Programming Interface) adalah sekumpulan aturan dan mekanisme yang memungkinkan aplikasi perangkat lunak untuk saling berkomunikasi. API menyediakan cara untuk mengambil data atau fungsionalitas dari aplikasi lain tanpa harus memahami secara mendalam bagaimana aplikasi tersebut bekerja.

## REST (Representational State Transfer)
REST adalah salah satu arsitektur API yang paling populer. RESTful API bekerja dengan menggunakan metode HTTP standar seperti GET, POST, PUT, dan DELETE untuk melakukan operasi pada resource. API berbasis REST menggunakan URL untuk mengidentifikasi resource dan format respons seperti JSON atau XML.

## JSON (JavaScript Object Notation)
JSON adalah format pertukaran data yang ringan dan mudah dibaca oleh manusia. JSON sering digunakan sebagai format data untuk RESTful API karena kesederhanaannya dalam penulisan dan pemrosesan. Berikut contoh format JSON:

```json
{
    "id": 1,
    "name": "John Doe",
    "email": "john.doe@example.com"
}
```

## HTTP Response Code
Kode respons HTTP memberikan informasi tentang hasil dari permintaan API. Berikut beberapa kode yang sering digunakan:
- **200 OK**: Permintaan berhasil dan server mengembalikan data yang diminta.
- **201 Created**: Resource baru berhasil dibuat.
- **400 Bad Request**: Permintaan tidak valid.
- **401 Unauthorized**: Permintaan memerlukan otentikasi.
- **404 Not Found**: Resource yang diminta tidak ditemukan.
- **500 Internal Server Error**: Terjadi kesalahan pada server.

## Tools
Untuk bekerja dengan RESTful API, terdapat beberapa alat yang dapat digunakan, di antaranya:
- **Postman**: Alat populer untuk menguji dan mengelola API.
- **cURL**: Alat command-line untuk mengirim permintaan HTTP dan mengelola API.
- **Insomnia**: Alternatif dari Postman untuk menguji API.

## Format Respons JSON
Respons API biasanya diberikan dalam format JSON. Contoh respons JSON dari API:

```json
{
    "page": 1,
    "total": 12,
    "data": [
        {
            "id": 1,
            "name": "John Doe",
            "email": "john.doe@example.com"
        },
        {
            "id": 2,
            "name": "Jane Doe",
            "email": "jane.doe@example.com"
        }
    ]
}
```

Format ini sering kali digunakan karena mudah diuraikan oleh aplikasi dan mendukung berbagai tipe data.

## Filtering, Sorting, Paging, dan Field Selection
RESTful API sering menyediakan fitur-fitur berikut untuk mengelola data dalam jumlah besar:

- **Filtering (Penyaringan)**: Menyaring hasil berdasarkan kriteria tertentu, misalnya:
  ```
  GET /users?name=John
  ```
  Akan mengembalikan data pengguna dengan nama "John".

- **Sorting (Pengurutan)**: Mengurutkan hasil berdasarkan parameter, misalnya:
  ```
  GET /users?sort=created_at
  ```
  Mengurutkan data berdasarkan tanggal pembuatan.

- **Paging (Pembagian Halaman)**: Memisahkan hasil menjadi beberapa halaman untuk mengurangi beban transfer data, misalnya:
  ```
  GET /users?page=2&per_page=10
  ```
  Mengambil halaman kedua dengan 10 pengguna per halaman.

- **Field Selection (Pemilihan Kolom)**: Memilih hanya kolom tertentu dari data yang diambil, misalnya:
  ```
  GET /users?fields=id,name,email
  ```
  Mengembalikan hanya `id`, `name`, dan `email` dari resource pengguna.

Dengan fitur-fitur ini, RESTful API dapat memberikan hasil yang lebih relevan dan efisien berdasarkan kebutuhan aplikasi pengguna.
