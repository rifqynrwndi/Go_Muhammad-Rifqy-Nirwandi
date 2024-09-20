# Data Structure

## Functions in Go
Fungsi dalam Go merupakan blok kode yang dapat dipanggil untuk menjalankan tugas tertentu. Setiap fungsi memiliki nama, parameter, dan dapat mengembalikan nilai. Fungsi bisa memiliki return values.
```go
package main

import "fmt"

func add(x, y int) int {
    return x + y
}

func main() {
    fmt.Println(add(3, 5)) // Output: 8
}
```

## Array, Slices, and Map in Go
- Array adalah kumpulan data bertipe sama, yang disimpan dalam sebuah variabel. Array memiliki kapasitas yang nilainya ditentukan pada saat pembuatan, menjadikan elemen/data yang disimpan di array tersebut jumlahnya tidak boleh melebihi yang sudah dialokasikan.
- Slices adalah abstraksi array dengan panjang yang bisa berubah. Slices lebih fleksibel dan sering digunakan dibanding array.
- Maps menyimpan pasangan key-value, mirip dengan dictionary di bahasa lain. Maps memiliki performa lookup yang cepat.
```go
// Array
var arr [3]int = [3]int{1, 2, 3}

// Slice
slice := []int{4, 5, 6}

// Map
m := map[string]int{"one": 1, "two": 2}
```

## Slice and Map Package
Go memiliki package bawaan dengan berbagai fungsi untuk memanipulasi slice dan maps. Beberapa package adalah sebagai berikut :
- Slice Package
1. `append(slice, elem)` : Menambah elemen ke slice.
2. `copy(dstSlice, srcSlice)` : Menyalin data dari satu slice ke slice lainnya.
3. `len(slice)` : Mengembalikan panjang slice.

- Map Package
1. `delete(map, key)` : Menghapus key dari map.
2. `value, exists := map[key]` : Mengecek apakah key ada dan mengambil nilainya.