# System Design

## 1. Diagram

Di dalam proses pengembangan sistem, **diagram** adalah salah satu cara terbaik untuk memahami keseluruhan alur kerja dari sistem. Diagram visual membantu dalam:
- Menggambarkan entitas dan relasinya (seperti **Entity Relationship Diagram (ERD)**)
- Menunjukkan **use cases** untuk user dan admin
- Menggambarkan **High Level Architecture** yang menjelaskan hubungan antara frontend, backend, dan database

Beberapa contoh diagram dalam sistem:
- **ERD**: Menunjukkan entitas seperti `Pasien`, `Dokter`, `Rekam Medis`, dan `Feedback` beserta relasi antar entitas.
- **Use Case Diagram**: Memperlihatkan interaksi antara pasien, admin, dan sistem.
- **High-Level Architecture Diagram**: Menggambarkan arsitektur sistem termasuk komponen frontend, backend (API), dan database.

## 2. Distributed System

Distributed system adalah sistem yang komponennya terdistribusi di beberapa lokasi atau server namun saling berkoordinasi untuk bekerja sebagai satu kesatuan. Sistem terdistribusi memungkinkan:
- **Scalability**: Dengan menambahkan lebih banyak server untuk menangani beban kerja yang meningkat.
- **Reliability**: Jika satu komponen gagal, sistem tetap berfungsi dengan komponen yang lain.
- **Fault Tolerance**: Dengan adanya redundansi dan mekanisme pemulihan kesalahan.

Contoh penggunaan dalam sistem rumah sakit:
- Server yang terdistribusi di beberapa lokasi untuk menangani data pasien, dokter, dan rekam medis.

## 3. Job Queue dan Microservices

**Job Queue** digunakan untuk mengelola pekerjaan yang dieksekusi secara asynchronous, di mana tugas-tugas besar dapat didelegasikan ke antrian dan diproses satu per satu di latar belakang tanpa memblokir proses utama.

**Microservices** adalah pendekatan desain arsitektur di mana aplikasi dipecah menjadi layanan-layanan kecil yang mandiri. Setiap layanan bertanggung jawab atas satu fungsi tertentu dan berkomunikasi dengan layanan lain menggunakan API. Keunggulan microservices:
- **Scalability**: Setiap layanan dapat diskalakan secara independen.
- **Maintainability**: Layanan yang lebih kecil lebih mudah dikelola dan diupdate.
- **Fault Isolation**: Jika salah satu layanan gagal, sistem lain tetap berfungsi.

Contoh implementasi:
- **Pasien Service**, **Dokter Service**, dan **Rekam Medis Service** yang diimplementasikan sebagai microservices dan saling berkomunikasi melalui API.

## 4. SQL dan NoSQL

### SQL (Structured Query Language):
- **Relational Databases** seperti MySQL, PostgreSQL menggunakan SQL sebagai bahasa utama untuk mengelola data yang terstruktur.
- SQL memiliki **schema-based** desain yang ketat dan lebih cocok untuk data yang membutuhkan integritas tinggi dan relasi antar tabel.

Contoh: 
- **Rekam Medis** dan **Data Pasien** biasanya menggunakan SQL karena membutuhkan relasi antar entitas.

### NoSQL:
- **Non-relational databases** seperti MongoDB, Cassandra lebih fleksibel karena tidak bergantung pada schema yang ketat.
- NoSQL cocok untuk aplikasi yang membutuhkan skalabilitas tinggi dan struktur data yang tidak teratur atau berubah-ubah.

Contoh:
- **Feedback** dari pasien dapat disimpan dalam NoSQL karena struktur data bisa bervariasi.

## 5. Cache dan Indexing

**Cache** adalah penyimpanan sementara yang digunakan untuk menyimpan data yang sering diakses, sehingga waktu akses menjadi lebih cepat. Cache biasanya ditempatkan di memori dengan kecepatan tinggi, seperti **Redis** atau **Memcached**.

**Indexing** adalah teknik yang digunakan di database untuk mempercepat proses pencarian data. Dengan membuat index pada kolom-kolom tertentu, database bisa lebih cepat menemukan data yang diminta tanpa harus melakukan pemindaian seluruh tabel.

Penggunaan di sistem:
- **Cache** dapat digunakan untuk menyimpan data pasien yang sering diakses, seperti informasi login.
- **Indexing** dapat diterapkan pada kolom seperti `pasien_id` atau `dokter_id` di tabel rekam medis untuk mempercepat pencarian.
