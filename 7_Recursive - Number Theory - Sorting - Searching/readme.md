# Recursive - Number Theory - Sorting - Searching

## Recursive
Rekursif adalah teknik pemrograman di mana sebuah fungsi memanggil dirinya sendiri untuk menyelesaikan masalah yang lebih kecil dari masalah aslinya. Teknik ini sering digunakan untuk menyelesaikan masalah yang dapat dipecah menjadi sub-masalah yang lebih kecil dan serupa. Rekursif biasanya digunakan untuk penyelesaian masalah yang lebih kompleks seperti pohon biner, graf, dan lainnya. Rekursif terdiri atas `base case`, `recurrence relation (relasi rekursif)`, dan `penyusunan kembali hasil`.

## Number Theory

- ### Factorial
Faktorial dari sebuah bilangan n (dilambangkan dengan n!) adalah hasil perkalian semua bilangan bulat positif kurang dari atau sama dengan n.
- ### Prime Number
Bilangan prima adalah bilangan bulat positif yang hanya memiliki dua faktor: 1 dan dirinya sendiri.
- ### Greatest Common Divisor (GCD)
Greatest Common Divisor atau GCD atau FPB dari dua bilangan adalah bilangan bulat positif terbesar yang membagi kedua bilangan tersebut tanpa sisa.
- ### Least Common Multiple (LCM)
Least Common Multiple atau LCM atau KPK dari dua bilangan adalah bilangan bulat positif terkecil yang merupakan kelipatan dari kedua bilangan tersebut.

## Searching

- ### Linear Search
Pencarian linear adalah metode pencarian di mana kita memeriksa setiap elemen dari awal hingga akhir hingga kita menemukan elemen yang dicari atau mencapai akhir dari daftar. Linear Search memiliki time complexity sebesar `O(n)`
Contoh: 
```go
func linearSearch(arr []int, x int) int {
    for i, v := range arr {
        if v == x {
            return i
        }
    }
    return -1
}
```
- ### Builtin Search
Golang juga memiliki package bawaan yang sudah disediakan di dalam library nya.
Contoh:
```go
elements := []int{12,31,23,26,13,15,20}
value := 26

findIndex := sort.SearchInts(elements, value)
if value == elements[findIndex] {
    fmt.Println("Value found in element")
} else {
    fmt.Println("Value not found in element")
}
```
## Sorting

- ### Selection Sort
Selection sort adalah algoritma pengurutan di mana kita berulang kali menemukan elemen terkecil (atau terbesar) dari bagian yang belum diurutkan dan menukarnya dengan elemen pertama dari bagian yang belum diurutkan. Selection Sort memiliki Time Complexity sebesar `O(n^2)`
Contoh: 
```go
func selectionSort(arr []int) []int {
    for i := 0; i < len(arr); i++ {
        minIdx := i
        for j := i + 1; j < len(arr); j++ {
            if arr[j] < arr[minIdx] {
                minIdx = j
            }
        }
        arr[i], arr[minIdx] = arr[minIdx], arr[i]
    }
    return arr
}
```
- ### Counting Sort
Counting sort adalah algoritma pengurutan yang bekerja dengan menghitung jumlah kemunculan setiap elemen yang berbeda dalam array input. Counting sort memiliki Time Complexity sebesar `O(n+k)`.
Contoh:
```go
func countingSort(arr []int) []int {
    maxVal := max(arr)
    count := make([]int, maxVal+1)
    for _, v := range arr {
        count[v]++
    }
    sortedArr := make([]int, 0)
    for i, c := range count {
        for j := 0; j < c; j++ {
            sortedArr = append(sortedArr, i)
        }
    }
    return sortedArr
}

func max(arr []int) int {
    maxVal := arr[0]
    for _, v := range arr {
        if v > maxVal {
            maxVal = v
        }
    }
    return maxVal
}
```
- ### Merge Sort
Merge sort adalah algoritma pengurutan berbasis pembagian dan penaklukan yang membagi array menjadi dua bagian, mengurutkan masing-masing bagian, dan kemudian menggabungkannya kembali. Merge sort memiliki Time Complexity sebesar `O(log n)`.
Contoh:
```go
func mergeSort(arr []int) []int {
    if len(arr) <= 1 {
        return arr
    }
    mid := len(arr) / 2
    leftHalf := arr[:mid]
    rightHalf := arr[mid:]

    leftHalf = mergeSort(leftHalf)
    rightHalf = mergeSort(rightHalf)

    return merge(leftHalf, rightHalf)
}

func merge(left, right []int) []int {
    merged := make([]int, 0)
    for len(left) > 0 && len(right) > 0 {
        if left[0] <= right[0] {
            merged = append(merged, left[0])
            left = left[1:]
        } else {
            merged = append(merged, right[0])
            right = right[1:]
        }
    }
    merged = append(merged, left...)
    merged = append(merged, right...)
    return merged
}
```
- ### Builtin Sort
Sama seperti search, golang juga memiliki fungsi bawaan untuk melakukan sorting di dalam library nya.
Contoh:
```go
strs := []string{"c", "a", "b"}
slices.Sort(strs)
fmt.Println("Strings:", strs)

ints := []int{7, 2, 4}
slices.Sort(ints)
fmt.Println("Ints:   ", ints)

s := slices.IsSorted(ints)
fmt.Println("Sorted: ", s)
```
