# String - Advance Function - Pointer - Method - Struct and Interface

## String
Berikut beberapa fungsi dasar yang dapat digunakan untuk memanipulasi string.

- **Len**: Mengembalikan panjang string.
- **Compare**: Membandingkan dua string dan mengembalikan hasil perbandingan (-1, 0, 1).
- **Contains**: Memeriksa apakah substring terdapat di dalam string.

## Advance Funtion

Function memiliki beberapa jenis sebagai berikut

- **Closure**
Digunakan untuk mengakses variabel di luar cakupannya dan "mengikat" variabel tersebut ke dalam dirinya sendiri. Contoh penggunaannya seperti ini.
```go
package main

import "fmt"

func main() {
    counter := func() func() int {
        count := 0
        return func() int {
            count++
            return count
        }
    }

    increment := counter()

    fmt.Println(increment())
    fmt.Println(increment()) 
    fmt.Println(increment()) 
}

```
- **Defer**
Digunakan untuk menunda eksekusi sebuah statement hingga fungsi tersebut berakhir. Contoh penggunaannya seperti ini.
```go
package main

import "fmt"

func main() {
    defer fmt.Println("Print Terakhir")
    fmt.Println("Print")
}
```

## Pointer
Pointer adalah variabel yang meniympan alamat memori dari variabel lain. Pointer biasanya dideklarasi dengan `*` dan dipanggil dengan `&`.

## Struct
Struct adalah tipe data yang menyimpan nilai - nilai dengan tipe yang berbeda. 
``` go
type Person struct {
    Name string
    Age  int
}
``` 
## Method
Method adalah fungsi yang terkait dengan tipe tertentu contohnya pada Struct. Method berfungsi untuk membantu agar tidak ada konflik penamaan variabel, dan Method sangat mudah dipahami dan dipanggil daripada Function biasa.
```go
type Person struct {
    Name string
    Age  int
}

func (user Person) TambahNama() {
    fmt.Println("berhasil menambah nama")
}
```

## Interface
Kumpulan dari metode yang dapat diimplementasikan oleh berbagai tipe data. Biasanya Interface bisa juga disebut Behavior atau perlakuan dari suatu Object.
```go
type AddInfo interface {
    TambahNama() string
}

type Person struct {
    Name string
    Age  int
}

func (user Person) TambahNama() string {
    return user.Name
}
```

## Error Handling
Go menyediakan dua fungsi utama untuk menangani error secara eksplisit: panic dan recover.

`Panic`: Menghentikan eksekusi normal dari program dan memicu proses "panik", biasanya digunakan untuk kondisi error yang tidak dapat ditangani.
`Recover`: Menghentikan proses panik dan memungkinkan program untuk kembali ke jalur eksekusi normal.
```go
func main() {
    defer func() {
        if r := recover(); r != nil {
            fmt.Println("Recovered from panic:", r)
        }
    }()

    fmt.Println("Starting the program...")
    panic("Something went wrong!")
    fmt.Println("This will not be printed.")
}
```