# Problem Solving Paradigm – Brute Force, Greedy and Dynamic Programming

## Problem Solving Paradigm
Problem Solving Paradigm adalah pendekatan yang digunakan untuk menyelesaikan masalah dalam ilmu komputer. Berbagai paradigma dapat digunakan untuk menyelesaikan masalah yang sama, namun masing-masing memiliki kelebihan dan kekurangan.

## Brute Force
Pendekatan brute force adalah metode paling sederhana dan langsung dalam pemecahan masalah. Dengan pendekatan ini, kita mencoba semua kemungkinan solusi hingga menemukan solusi yang benar. Teknik ini umumnya tidak efisien, tetapi menjamin bahwa solusi terbaik akan ditemukan, meskipun memerlukan waktu komputasi yang lama. Brute force digunakan biasa untuk complete search.

### Complete Search
Dalam complete search, semua solusi potensial dieksplorasi secara menyeluruh. Ini biasanya dilakukan menggunakan teknik rekursi atau perulangan. Contoh penggunaan:
```go
func permute(nums []int, l int, r int) {
	if l == r {
		fmt.Println(nums)
		return
	}
	for i := l; i <= r; i++ {
		nums[l], nums[i] = nums[i], nums[l]
		permute(nums, l+1, r)
		nums[l], nums[i] = nums[i], nums[l] // Backtrack
	}
}
```
## Divide and Conquer
Divide and Conquer adalah pendekatan di mana masalah besar dipecah menjadi submasalah yang lebih kecil, lalu solusi dari submasalah digabungkan untuk membentuk solusi masalah asli. Pendekatan ini sering digunakan untuk mengurangi kompleksitas waktu komputasi.

### -  Binary Search
Binary search adalah cara untuk mencari elemen di dalam array yang sudah terurut dengan cara membagi ruang pencarian menjadi dua bagian secara berulang hingga elemen ditemukan.
```go
func binarySearch(arr []int, target int) int {
	left, right := 0, len(arr)-1
	for left <= right {
		mid := left + (right-left)/2
		if arr[mid] == target {
			return mid
		} else if arr[mid] < target {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	return -1 // Jika elemen tidak ditemukan
}
```
### -  Power atau Pangkat
Power adalah cara untuk menghitung pangkat suatu bilangan dengan menggunakan teknik pembagian dan penggabungan, seperti exponentiation by squaring.
```go
func power(base int, exp int) int {
	if exp == 0 {
		return 1 // Basis: setiap bilangan pangkat 0 adalah 1
	}
	if exp%2 == 0 {
		half := power(base, exp/2)
		return half * half // Jika eksponen genap, bagi dua eksponen
	}
	return base * power(base, exp-1) // Jika eksponen ganjil, kalikan dengan base
}
```
## Greedy
Pendekatan greedy memilih solusi terbaik saat ini atau lokal dengan harapan solusi ini akan membawa pada solusi global yang optimal. Pendekatan ini tidak selalu memberikan solusi optimal, tetapi bekerja baik pada masalah-masalah tertentu di mana pilihan lokal terbaik akan menghasilkan solusi optimal global.
### - Coin Change Problen
Coin Change adalah contoh implementasi Greedy yang meminimalkan jumlah koin yang digunakan untuk mencapai jumlah uang tertentu dengan memilih denominasi terbesar terlebih dahulu.
```go
func coinChange(coins []int, amount int) int {
	count := 0
	for i := len(coins) - 1; i >= 0; i-- {
		if coins[i] <= amount {
			count += amount / coins[i]
			amount %= coins[i]
		}
	}
	if amount > 0 {
		return -1 // Tidak bisa mencapai jumlah yang diinginkan
	}
	return count
}
```
### - Dijkstra Algoritma
Dijkstra Algoritma adalah salah satu contoh implementasi Greedy yang digunakan untuk menemukan jalur terpendek dari satu simpul ke simpul lainnya dalam graf dengan memilih simpul terdekat yang belum dikunjungi.

## Dynamic Programming
Dynamic Programming (DP) adalah teknik yang digunakan untuk menyelesaikan masalah yang memiliki submasalah tumpang tindih dan struktur optimal subproblems. Dalam DP, masalah besar dipecah menjadi submasalah yang lebih kecil, dan hasil dari submasalah tersebut disimpan agar tidak dihitung kembali (menyimpan hasil intermediate). Metode DP terbagi menjadi dua jenis:

### Top - Down (memoization)
Pendekatan ini memulai dari masalah terbesar dan memecahnya menjadi submasalah yang lebih kecil. Hasil dari setiap submasalah disimpan (memoization) agar tidak dihitung kembali jika diperlukan lagi.

### Bottom - Up (tabulation)
Pada pendekatan ini, kita menyelesaikan submasalah terkecil terlebih dahulu dan menyimpan hasilnya dalam tabel. Kemudian, hasil submasalah ini digunakan untuk menyelesaikan masalah yang lebih besar, hingga mencapai masalah terbesar yang ingin dipecahkan.