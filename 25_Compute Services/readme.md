Berikut adalah versi teks Anda yang telah diformat ke dalam Markdown:

```markdown
# Compute Service

## Docker
Docker adalah platform yang memungkinkan developer untuk mengemas aplikasi beserta seluruh dependensinya ke dalam container, sehingga aplikasi dapat berjalan konsisten di berbagai lingkungan.

### How Container Works
Container bekerja dengan cara mengisolasi aplikasi beserta seluruh dependensinya dalam satu lingkungan yang terpisah dari sistem operasi host. Setiap container berjalan di atas kernel host, tetapi tidak memerlukan sistem operasi lengkap seperti pada virtual machine (VM). Hal ini membuat container lebih ringan dan efisien.

### To Launch Container Using Docker
Untuk menjalankan container menggunakan Docker, Anda dapat menggunakan perintah berikut:
```bash
docker run -d -p 80:80 nginx
```
Contohnya, menjalankan container dari image `nginx`.

### To Build and Launch Your Own Container Images
Docker memungkinkan Anda membuat image kustom menggunakan Dockerfile. Dockerfile berisi instruksi untuk membangun image secara otomatis. Berikut langkah-langkahnya:
1. Buat file bernama `Dockerfile` di root direktori aplikasi Anda.
2. Tambahkan instruksi di Dockerfile sesuai kebutuhan.
3. Gunakan perintah berikut untuk membangun image:
   ```bash
   docker build -t [image-name] .
   ```
4. Setelah image selesai dibangun, jalankan dengan perintah:
   ```bash
   docker run [options] [image-name]
   ```

### To Deploy Your Application as Container
Untuk mendistribusikan aplikasi sebagai container, pastikan aplikasi dikemas sebagai image Docker, lalu dorong image ke Docker registry (misalnya, Docker Hub). Selanjutnya, tarik dan jalankan image tersebut di server atau environment produksi.

## Container vs VM
Container dan VM keduanya memungkinkan isolasi, tetapi container lebih ringan karena berbagi kernel host, sedangkan VM membutuhkan sistem operasi lengkap untuk setiap instance. VM memberikan isolasi lengkap pada tingkat sistem operasi, sementara container memberikan isolasi pada tingkat aplikasi.

## Deployment
Deployment adalah proses menyebarkan aplikasi dari lingkungan development ke production. Berikut beberapa strategi dan siklus dalam deployment.

### Definisi dan Strategi Deploy
- **Blue-Green Deployment**: Dua lingkungan identik digunakan (Blue sebagai current, Green sebagai new version), sehingga peralihan bisa dilakukan dengan lancar.
- **Canary Deployment**: Sebagian kecil pengguna dihadapkan pada versi baru sebelum peluncuran penuh.
- **Rolling Deployment**: Perbaruan dilakukan secara bertahap untuk meminimalkan downtime.

### Deployment Cycle
Deployment cycle meliputi langkah-langkah:
1. **Build**: Membuat aplikasi yang siap diproduksi.
2. **Test**: Menguji build untuk memastikan stabilitas.
3. **Deploy**: Mendistribusikan aplikasi ke production.

### Proses Deployment
1. Pastikan aplikasi berada di versi terbaru dan siap di-deploy.
2. Build dan push image Docker ke registry jika menggunakan container.
3. Terapkan strategi deployment yang sesuai.
4. Monitor aplikasi selama dan setelah deployment.

## Git sebagai Repository
Git adalah sistem version control terdistribusi yang membantu developer melacak perubahan kode. Sebagai repo, Git menyimpan histori dan memungkinkan kolaborasi dalam pengembangan aplikasi.

## AWS
Amazon Web Services (AWS) menyediakan infrastruktur cloud untuk menjalankan aplikasi. Compute service seperti EC2 dan layanan manajemen database seperti RDS adalah layanan AWS yang sering digunakan untuk deployment aplikasi.

### Linux Server
Server Linux sering digunakan sebagai host dalam deployment aplikasi karena kestabilan dan kompatibilitasnya dengan banyak tool.

### Setup dan Instalasi Webserver
Instalasi webserver di server Linux biasanya melibatkan tool seperti Nginx atau Apache. Setelah webserver terinstal, konfigurasi dapat dilakukan untuk mengatur virtual hosts atau SSL.

### Setup Domain
Setup domain melibatkan penghubungan domain ke IP publik server agar aplikasi dapat diakses melalui URL. Hal ini dilakukan di pengaturan DNS provider domain, dengan membuat A record yang menunjuk ke alamat IP server.