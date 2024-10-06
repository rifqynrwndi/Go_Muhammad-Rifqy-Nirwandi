package main

import "fmt"

// Kekurangan 1: Nama struct 'user' sebaiknya menggunakan huruf kapital 'User' sesuai kebiasaan Go untuk public types.
type user struct {  
  Email    string
  Password string
}

// Kekurangan 2: Nama struct 'userRepo' sebaiknya menggunakan huruf kapital 'UserRepo' jika ingin bisa diakses di luar package.
type userRepo struct {  
  DB []user
}

func (r userRepo) Register(u user) {
  // Kekurangan 3: Kondisi pengecekan 'if' tidak diakhiri dengan return setelah menampilkan pesan.
  // Jika Email atau Password kosong, fungsi seharusnya berhenti (dengan return) agar tidak melanjutkan eksekusi dan mengubah DB.
  if u.Email == "" || u.Password == "" {
      fmt.Println("register failed")
      return // Tambahkan return untuk menghentikan eksekusi setelah kegagalan
  }

  // Kekurangan 4: Metode ini seharusnya menerima pointer receiver (*userRepo) karena akan memodifikasi DB.
  // Tanpa pointer, perubahan yang dilakukan tidak akan berpengaruh pada 'r.DB' yang asli.
  r.DB = append(r.DB, u) 
}

func (r userRepo) Login(u user) string {
  // Kekurangan 5: Sama seperti di fungsi Register, pengecekan gagal login seharusnya diikuti dengan return untuk mencegah kode berikutnya berjalan.
  if u.Email == "" || u.Password == "" {
      fmt.Println("login failed")
      return "" // Tambahkan return untuk menghentikan eksekusi setelah kegagalan
  }

  for _, us := range r.DB {
      // Kekurangan 6: Seharusnya gunakan perbandingan string case-insensitive jika sistem mengizinkan login dengan variasi huruf besar/kecil.
      if us.Email == u.Email && us.Password == u.Password {
          return "auth token"
      }
  }

  // Kekurangan 7: Sebaiknya tambahkan log atau pesan error jika login gagal untuk memberitahukan alasan kegagalan.
  return "" 
}
