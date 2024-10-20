# ORM dan Struktur Kode (MVC)

## Apa itu ORM

ORM (Object-Relational Mapping) adalah teknik yang digunakan untuk menghubungkan database relasional dengan pemrograman berorientasi objek. Dengan ORM, objek dalam kode dapat dipetakan ke tabel dalam database, sehingga mempermudah manipulasi data tanpa perlu menulis query SQL secara manual.

## Apa itu GORM

GORM adalah sebuah ORM untuk bahasa pemrograman Go yang memungkinkan pengembang untuk bekerja dengan database menggunakan cara yang lebih intuitif. GORM mendukung berbagai database, termasuk MySQL, PostgreSQL, SQLite, dan SQL Server. Dengan GORM, operasi CRUD (Create, Read, Update, Delete) dapat dilakukan dengan mudah menggunakan method di dalam model.

## Keuntungan Menggunakan ORM

- **Abstraksi Database**: ORM menyembunyikan detail implementasi database, sehingga pengembang dapat lebih fokus pada logika bisnis daripada query SQL.
- **Pengurangan Waktu Pengembangan**: Dengan ORM, proses pengembangan aplikasi menjadi lebih cepat karena mengurangi jumlah kode yang harus ditulis.
- **Keterbacaan Kode**: Kode menjadi lebih mudah dibaca dan dipahami, karena menggunakan objek daripada query SQL yang kompleks.
- **Pengelolaan Relasi**: ORM menyediakan cara yang mudah untuk mengelola relasi antar tabel dalam database.

## Kekurangan ORM

- **Kinerja**: Penggunaan ORM dapat mengurangi kinerja dibandingkan dengan penggunaan query SQL langsung, terutama untuk query yang kompleks.
- **Kurangnya Kontrol**: Terkadang, ORM dapat membatasi kontrol yang dimiliki pengembang terhadap query yang dijalankan, sehingga tidak dapat mengoptimalkan performa.
- **Learning Curve**: Meskipun ORM dapat menyederhanakan interaksi dengan database, mempelajari cara kerja ORM dan cara terbaik untuk menggunakannya juga membutuhkan waktu.

## Migrasi Database

Migrasi database adalah proses memindahkan skema database dari satu versi ke versi lain. GORM menyediakan fitur migrasi yang memungkinkan pengembang untuk memperbarui struktur database secara otomatis sesuai dengan perubahan pada model.

## Relasi Database di GORM

GORM mendukung berbagai jenis relasi database, seperti relasi satu-ke-satu, satu-ke-banyak, dan banyak-ke-banyak. Relasi ini dapat didefinisikan dalam model menggunakan tag dan method yang disediakan oleh GORM.

## Instalasi GORM dan Driver MySQL

Untuk menginstal GORM dan driver MySQL, gunakan perintah berikut:

```bash
go get -u gorm.io/gorm
go get -u gorm.io/driver/mysql
```

## Menghubungkan Aplikasi ke Database
Untuk menghubungkan aplikasi ke database menggunakan GORM, buatlah koneksi ke database dengan kode berikut:

```go
import (
    "gorm.io/driver/mysql"
    "gorm.io/gorm"
)

func connectDatabase() (*gorm.DB, error) {
    dsn := "username:password@tcp(localhost:3306)/dbname?charset=utf8mb4&parseTime=True&loc=Local"
    db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
    return db, err
}
```

## Struktur Kode MVC
MVC (Model-View-Controller) adalah pola arsitektur perangkat lunak yang membagi aplikasi menjadi tiga komponen utama:

- **Model**: Berfungsi untuk mengelola data dan logika bisnis aplikasi. Model berinteraksi dengan database dan menyediakan data yang dibutuhkan oleh controller.
- **View**: Menangani presentasi data kepada pengguna. View berfungsi untuk menampilkan antarmuka pengguna dan menerima input.
- **Controller**: Mengatur alur aplikasi dan menghubungkan model dan view. Controller menerima input dari view, memprosesnya, dan memberikan respons yang sesuai.

## Keuntungan MVC
- **Pemisahan Tanggung Jawab**: MVC memisahkan logika bisnis, presentasi, dan pengendalian alur, yang memudahkan pengembangan dan pemeliharaan aplikasi.
- **Kemudahan Pengujian**: Dengan pemisahan ini, setiap komponen dapat diuji secara terpisah, yang meningkatkan kualitas kode dan mengurangi bug.
- **Kolaborasi Tim**: Pemisahan tanggung jawab memungkinkan tim pengembang untuk bekerja secara paralel pada bagian yang berbeda dari aplikasi.
