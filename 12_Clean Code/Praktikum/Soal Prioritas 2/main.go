package main

import "fmt"

// Mengubah 'set' menjadi 'Set' untuk mengikuti konvensi penamaan struct dengan huruf kapital
type Set struct {
    elements map[int]struct{} // Menggunakan struct{} untuk menghemat memori (karena hanya perlu menyimpan keberadaan elemen, bukan nilainya)
}

// Gunakan pointer receiver karena kita memodifikasi struct
func (s *Set) Add(element int) {
    s.elements[element] = struct{}{}
}

func (s *Set) GetAll() []int {
    var result []int

    for element := range s.elements {
        result = append(result, element)
    }

    return result
}

// Gunakan pointer receiver untuk memodifikasi struct
func (s *Set) Remove(element int) {
    delete(s.elements, element)
}

func main() {
    // Inisialisasi set dengan nama lebih deskriptif dan jelas
    set := &Set{
        elements: make(map[int]struct{}),
    }

    // Tambahkan elemen ke set
    set.Add(1)
    set.Add(2)
    set.Add(1) // Duplikasi elemen tidak akan berpengaruh pada set
    set.Add(3)
    set.Add(4)
    set.Add(5)

    // Coba hapus elemen yang tidak ada, tidak ada efek negatif
    set.Remove(100)

    // Tampilkan semua elemen yang ada di set
    fmt.Println(set.GetAll())
}
