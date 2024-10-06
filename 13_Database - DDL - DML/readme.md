# Database - DDL- DML

# Ringkasan Materi Database

## 1. Database Introduction
**Database** adalah kumpulan data yang terorganisir dan disimpan secara elektronik sehingga memudahkan pengelolaan, pengambilan, dan manipulasi data. Basis data digunakan dalam berbagai aplikasi untuk menyimpan informasi agar mudah diakses dan diperbarui.

### Fitur Utama:
- **Data Organizing**: Menyimpan data dalam format terstruktur (misalnya, tabel).
- **Data Management**: Memfasilitasi penambahan, modifikasi, penghapusan, dan pengambilan data.
- **Data Security**: Menyediakan mekanisme untuk melindungi dan mengontrol akses ke informasi sensitif.

---

## 2. Database Relation
**Database relation** adalah tipe basis data yang menyimpan data secara terstruktur menggunakan tabel (relasi) yang terdiri dari baris (rekor) dan kolom (atribut). Data di tabel yang berbeda dapat saling berhubungan melalui kunci (keys).

### Jenis-jenis Relation:
- **One-to-One**: Satu rekaman di satu tabel berhubungan dengan satu rekaman di tabel lain.
- **One-to-Many**: Satu rekaman di satu tabel berhubungan dengan banyak rekaman di tabel lain.
- **Many-to-Many**: Beberapa rekaman di satu tabel berhubungan dengan beberapa rekaman di tabel lain, biasanya melalui tabel penghubung.

---

## 3. Entity Relationship Diagram (ERD)
**Entity Relationship Diagram (ERD)** adalah representasi grafis dari entitas (objek) dan hubungan antar entitas di dalam basis data. ERD membantu dalam merancang struktur basis data sebelum membangun basis data secara nyata.

### Komponen Utama:
- **Entitas**: Objek atau konsep yang menyimpan data (misalnya, Pengguna, Produk, Pesanan).
- **Atribut**: Karakteristik atau properti dari entitas (misalnya, Nama Pengguna, Harga Produk).
- **Relasi**: Hubungan antar entitas (misalnya, Pengguna melakukan pesanan).

### Contoh Simbol ERD:
- **Persegi panjang** mewakili entitas.
- **Elips** mewakili atribut.
- **Berlian** mewakili relasi.

---

## 4. Pernyataan SQL
**SQL (Structured Query Language)** adalah bahasa yang digunakan untuk berinteraksi dengan basis data relasional. SQL memungkinkan pengguna melakukan berbagai operasi seperti membuat tabel, memasukkan data, melakukan query (pengambilan data), dan mengelola struktur basis data.

### Perintah SQL Umum:
- **CREATE**: Membuat basis data atau tabel baru.
  ```sql
  CREATE TABLE pengguna (
      id INT PRIMARY KEY,
      nama VARCHAR(255)
  );
  ```
  - **INSERT**: Memasukkan data baru ke dalam tabel.
  ```sql
  INSERT INTO pengguna (id, nama) VALUES (1, 'John Doe');
  ```
  - **SELECT**: Mengambil data dari tabel.
  ```sql
  SELECT * FROM pengguna;
  ```
  - **UPDATE**: Memodifikasi data yang ada dalam tabel.
  ```sql
  UPDATE pengguna SET nama = 'Jane Doe' WHERE id = 1;
  ```
  - **DELETE**: Menghapus data dari tabel.
  ```sql
  DELETE FROM pengguna WHERE id = 1;
  ```