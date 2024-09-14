# Basic Programming

## 1. Introduction to Golang
- Pengenalan tentang **Go** sebagai bahasa pemrograman, sejarahnya, dan alasan mengapa Go populer dalam pengembangan perangkat lunak modern.
- Cara menginstal Go di berbagai platform dan menulis program pertama menggunakan Go.
- Struktur dasar program Go, seperti fungsi `main`, deklarasi paket (`package`), dan cara menjalankan program Go.

```go
package main

import "fmt"

func main() {
    fmt.Println("Hello, World!")
}
```

## 2. Variables and Data Types in Go

Variabel merupakan nama yang diberikan ke sebuah pemrograman untuk menyimpan suatu nilai data. Terdapat 3 tipe data utama di golang yaitu `boolean`, `numeric`, dan `string`. Variabel declaration di Golang dilakukan dengan `type value address` contohnya `var age int` atau kalau variabel sudah memiliki nilai bisa dideklarasi dengan `<variable_name> := <value>` contohnya `name := rifqy` atau `age := 18` dan golang akan otomatis menentukan tipe datanya.

## Operators in Go

Operators Go digunakan untuk melakukan operasi pada variabel dan nilai. Berikut adalah beberapa jenis operator di Go:

Arithmetic Operators: Digunakan untuk operasi matematika dasar.
- `+` (Penjumlahan)
- `-` (Pengurangan)
- `*` (Perkalian)
- `/` (Pembagian)
- `%` (Modulus atau Sisa bagi)

Logical Operators : Digunakan untuk menggabungkan beberapa kondisi logika
- `&&` (AND)
- `||` (OR)
- `!` (NOT)

Relational Operators: Digunakan untuk membandingkan dua nilai, menghasilkan true atau false.
- `==` (Sama dengan)
- `!=` (Tidak sama dengan)
- `>` (Lebih besar)
- `<` (Lebih kecil)
- `>=` (Lebih besar atau sama dengan)
- `<=` (Lebih kecil atau sama dengan)

Bitwise Operators: Mengoperasikan bit-level.
- `&` (AND)
- `|` (OR)
- `^` (XOR)
- `<<` (Left shift)
- `>>` (Right shift)