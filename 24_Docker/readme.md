# Docker

## 1. Docker
Docker adalah platform untuk mengotomatisasi penyebaran aplikasi sebagai container portabel dan mandiri. Docker memungkinkan developer untuk membuat, menguji, dan menerapkan aplikasi di lingkungan yang konsisten.

### Keuntungan Docker:
- **Portabilitas**: Docker dapat dijalankan di berbagai lingkungan (lokal, server, cloud).
- **Isolasi**: Aplikasi berjalan dalam container yang terisolasi satu sama lain.
- **Efisiensi Sumber Daya**: Container menggunakan sumber daya yang lebih rendah dibandingkan VM karena berbagi kernel host.

## 2. Virtual Machine
Virtual Machine (VM) adalah lingkungan komputasi virtual yang meniru perangkat keras fisik. VM memungkinkan beberapa sistem operasi berjalan di atas hypervisor yang sama, tetapi lebih berat dibandingkan container.

### Keuntungan VM:
- **Isolasi Lengkap**: Setiap VM beroperasi secara independen dengan OS-nya sendiri.
- **Kompatibilitas Beragam**: VM mendukung berbagai OS.

### Kekurangan VM:
- **Penggunaan Sumber Daya Tinggi**: VM memerlukan sumber daya komputasi yang besar karena setiap VM memiliki OS dan kernel sendiri.

## 3. Container
Container adalah lingkungan yang ringan yang mengemas aplikasi dan semua dependensinya sehingga dapat dijalankan di mana saja. Container lebih efisien daripada VM karena berbagi kernel OS host.

### Keuntungan Container:
- **Ringan dan Cepat**: Tidak memerlukan OS terpisah seperti VM.
- **Penggunaan Sumber Daya Rendah**: Menggunakan kernel host sehingga lebih hemat sumber daya.
- **Cepat di-deploy dan di-scale**: Container dapat dibuat, dihentikan, dan digandakan dengan cepat.

## 4. Cheat Sheet Docker
### Basic Commands
- **docker run [options] image**: Menjalankan container dari image.
- **docker build -t [name] .**: Membangun image baru dari Dockerfile di direktori.
- **docker ps**: Melihat container yang sedang berjalan.
- **docker stop [container_id]**: Menghentikan container yang berjalan.

## 5. Basic Command
- **docker pull [image]**: Mengunduh image dari Docker Hub.
- **docker rm [container_id]**: Menghapus container yang sudah dihentikan.
- **docker rmi [image_id]**: Menghapus image.

## Introduction to Docker
Docker mempermudah proses membuat, menjalankan, dan mengelola aplikasi dalam lingkungan container yang konsisten, memungkinkan developer untuk menghindari masalah perbedaan lingkungan.

## Docker Basics
### Images dan Container
- **Image**: Template yang berisi aplikasi dan dependensinya.
- **Container**: Instance dari image yang sedang berjalan.

## Docker Compose
Docker Compose memungkinkan definisi dan pengelolaan beberapa container sekaligus menggunakan file YAML.

### Perintah Dasar Docker Compose
- **docker-compose up**: Menjalankan seluruh service yang didefinisikan dalam file YAML.
- **docker-compose down**: Menghentikan dan menghapus semua service.
