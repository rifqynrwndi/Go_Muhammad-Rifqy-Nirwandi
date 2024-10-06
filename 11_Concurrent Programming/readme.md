# Concurrent Programming

1. **Goroutine**  
   Goroutine adalah unit eksekusi ringan yang memungkinkan eksekusi fungsi secara bersamaan dalam Go. Setiap goroutine dijalankan secara asinkron dan independen, memungkinkan Go untuk memanfaatkan fitur *multi-threading* tanpa harus secara eksplisit mengelola thread.
   ```go
   package main

   import "fmt"

   func printMessage() {
       fmt.Println("Hello from goroutine!")
   }

   func main() {
       go printMessage()
       fmt.Println("Hello from main function")
   }
   ```

2. **Channel and Select** 
Channel digunakan untuk komunikasi antar goroutine. Dengan channel, data dapat dikirim dan diterima antara goroutine dengan aman. Select memungkinkan memilih kanal yang siap menerima atau mengirim data, memberikan fleksibilitas untuk mengelola beberapa channel.
Contoh implementasi Channel dan Select:
```go
package main

import (
    "fmt"
    "time"
)

func sendMessage(ch chan string) {
    time.Sleep(2 * time.Second)
    ch <- "Hello from goroutine"
}

func main() {
    ch := make(chan string)
    go sendMessage(ch)

    select {
    case msg := <-ch:
        fmt.Println(msg)
    case <-time.After(3 * time.Second):
        fmt.Println("Timeout!")
    }
}

```

3. Race Condition
*Race condition* terjadi ketika beberapa goroutine mengakses atau memodifikasi variabel yang sama secara bersamaan, tanpa sinkronisasi yang tepat. Berikut ini adalah beberapa cara memperbaiki race condition:

   - **WaitGroups**  
     *WaitGroup* digunakan untuk memastikan bahwa semua goroutine telah selesai sebelum melanjutkan. Ini membantu menunggu seluruh goroutine menyelesaikan pekerjaannya.

     ```go
     package main

     import (
         "fmt"
         "sync"
     )

     var counter int

     func increment(wg *sync.WaitGroup) {
         counter++
         wg.Done()  // Menandakan bahwa goroutine ini selesai
     }

     func main() {
         var wg sync.WaitGroup
         for i := 0; i < 1000; i++ {
             wg.Add(1)
             go increment(&wg)
         }
         wg.Wait()
         fmt.Println("Final Counter:", counter)
     }
     ```

   - **Channel Blocking**  
     Channel blocking adalah cara untuk memastikan sinkronisasi antar goroutine, memastikan satu goroutine menunggu sampai channel menerima data dari goroutine lain.

     ```go
     package main

     import "fmt"

     func sendMessage(ch chan string) {
         ch <- "Message from goroutine"
     }

     func main() {
         ch := make(chan string)
         go sendMessage(ch)

         // Menunggu sampai ada data yang diterima dari channel
         msg := <-ch
         fmt.Println(msg)
     }
     ```

   - **Mutex**  
     Mutex digunakan untuk mengunci akses ke variabel tertentu sehingga hanya satu goroutine yang dapat memodifikasinya pada satu waktu, menghindari race condition.

     ```go
     package main

     import (
         "fmt"
         "sync"
     )

     var counter int
     var mu sync.Mutex

     func increment(wg *sync.WaitGroup) {
         mu.Lock()  // Mengunci sebelum akses
         counter++
         mu.Unlock()  // Membuka kunci setelah selesai
         wg.Done()
     }

     func main() {
         var wg sync.WaitGroup
         for i := 0; i < 1000; i++ {
             wg.Add(1)
             go increment(&wg)
         }
         wg.Wait()
         fmt.Println("Final Counter:", counter)
     }
     ```
