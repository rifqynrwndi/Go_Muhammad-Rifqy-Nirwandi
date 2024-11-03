# CI/CD

## Continuous Integration
Continuous Integration (CI) adalah praktik di mana developer secara rutin menggabungkan perubahan kode ke dalam branch utama sebuah repository. Setiap perubahan kode secara otomatis di-build dan diuji untuk mendeteksi masalah sedini mungkin.

### Manfaat Continuous Integration:
- Mengurangi risiko konflik kode saat penggabungan.
- Deteksi bug lebih awal melalui build dan testing otomatis.
- Meningkatkan kolaborasi antar tim.

## Continuous Delivery
Continuous Delivery (CD) adalah praktik di mana perubahan kode yang telah diintegrasikan ke dalam branch utama dapat dirilis ke lingkungan produksi kapan saja. Proses ini memastikan bahwa build yang telah lulus pengujian dapat di-deploy dengan aman.

### Manfaat Continuous Delivery:
- Memastikan kode selalu siap untuk di-deploy.
- Mengurangi waktu yang dibutuhkan untuk merilis fitur baru.
- Meminimalkan risiko bug di produksi melalui pengujian yang konsisten.

## Continuous Deployment
Continuous Deployment adalah langkah berikutnya setelah Continuous Delivery, di mana setiap perubahan kode yang telah lulus pengujian otomatis langsung di-deploy ke produksi tanpa memerlukan persetujuan manual.

### Manfaat Continuous Deployment:
- Mempercepat proses release.
- Memberikan feedback lebih cepat ke tim development.
- Memastikan pengguna mendapatkan fitur baru lebih cepat.

## How to CI & CD
Implementasi CI/CD melibatkan beberapa langkah:
1. **Setup Environment**: Persiapkan environment yang sesuai (pipeline CI/CD).
2. **Automated Testing**: Tambahkan pengujian otomatis untuk memastikan build stabil.
3. **Build Pipeline**: Buat pipeline yang secara otomatis membangun dan menguji kode.
4. **Deployment Pipeline**: Tambahkan tahap deploy setelah build berhasil untuk Continuous Delivery atau Continuous Deployment.

## CI/CD with GitHub Actions
GitHub Actions adalah platform integrasi bawaan dari GitHub yang memungkinkan pembuatan workflow otomatis untuk build, test, dan deploy aplikasi.

### Benefit of GitHub Actions:
- Integrasi yang mudah dengan repository GitHub.
- Tersedia banyak action dan template yang mempermudah setup pipeline.
- Mendukung berbagai bahasa pemrograman dan platform.

### Step by Step CI/CD with GitHub Actions:
1. **Buat File Workflow**: Tambahkan file workflow di folder `.github/workflows` dengan ekstensi `.yml` (contoh: `ci-cd.yml`).
2. **Definisikan Job**: Tentukan langkah-langkah build, test, dan deploy di dalam job.
3. **Tentukan Trigger**: Atur workflow agar berjalan otomatis saat ada push, pull request, atau schedule tertentu.
4. **Tambahkan Steps**: Setiap langkah dapat mencakup checkout kode, setup environment, build, test, dan deploy.
5. **Commit dan Push Workflow**: Push file workflow ke repository untuk mengaktifkan CI/CD.

### Contoh Workflow GitHub Actions:
```yaml
name: CI/CD Pipeline

on:
  push:
    branches:
      - main

jobs:
  build:
    runs-on: ubuntu-latest
    steps:
      - name: Checkout Code
        uses: actions/checkout@v2

      - name: Setup Node.js
        uses: actions/setup-node@v2
        with:
          node-version: '14'

      - name: Install Dependencies
        run: npm install

      - name: Run Tests
        run: npm test

      - name: Build Application
        run: npm run build

      - name: Deploy
        run: echo "Deploy step here"
```

GitHub Actions memungkinkan tim untuk mengatur pipeline CI/CD yang fleksibel dan terintegrasi langsung dengan repository GitHub.