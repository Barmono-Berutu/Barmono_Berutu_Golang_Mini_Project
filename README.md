# Project Mini Golang Airquality

API ini adalah proyek mini yang dibuat dengan Golang untuk memantau dan mengelola data kualitas udara. API ini memungkinkan pengguna untuk melakukan registrasi, login, dan mengakses data kualitas udara berdasarkan lokasi atau rekomendasi.

---

## Daftar Isi
1. [ERD (Entity-Relationship Diagram)](#erd-entity-relationship-diagram)
2. [HLA (High-Level Architecture)](#hla-high-level-architecture)
3. [Documentation List API](#documentation-list-api)
4. [Contoh Hasil](#contoh-hasil)
5. [Cara Instalasi](#cara-instalasi)
6. [Cara Penggunaan](#cara-penggunaan)

---

## ERD (Entity-Relationship Diagram)
Berikut adalah ERD yang menjelaskan tabel dan hubungan antar entitas pada database proyek ini:

![ERD](image.png)

---

## HLA (High-Level Architecture)
Berikut adalah arsitektur dari API ini:

![alt text](image-1.png)

---

## Documentation List API

Berikut adalah daftar endpoint API yang tersedia dalam sistem ini:

| Method | Endpoint                    | Description                                    |
|--------|-----------------------------|------------------------------------------------|
| POST   | `/login`                    | Login pengguna dan mendapatkan token           |
| POST   | `/register`                 | Registrasi pengguna baru                       |
| GET    | `/logout`                   | Logout pengguna dari sistem                    |
| POST   | `/airquality`               | Menambahkan data kualitas udara baru           |
| GET    | `/airquality`               | Mendapatkan semua data kualitas udara          |
| GET    | `/airquality/{id}`          | Mendapatkan data kualitas udara berdasarkan ID |
| PUT    | `/airquality/{id}`          | Memperbarui data kualitas udara berdasarkan ID |
| DELETE | `/airquality/{id}`          | Menghapus data kualitas udara berdasarkan ID   |
| GET    | `/airquality/alert`         | Mendapatkan data alert kualitas udara          |
| GET    | `/airquality/rekomendasi`   | Mendapatkan rekomendasi kualitas udara         |

---

## Contoh Hasil API
Berikut adalah contoh hasil dari masing-masing endpoint API:

### 1. Registrasi
![alt text](image-2.png)

### 2. Login
![alt text](image-3.png)

### 3. Tambah Data Kualitas Udara
![alt text](image-4.png)

### 4. Mendapatkan Semua Data Kualitas Udara
![alt text](image-5.png)

### 5. Mendapatkan Data Kualitas Udara Berdasarkan ID
![alt text](image-6.png)

### 6. Memperbarui Data Kualitas Udara
![alt text](image-7.png)

### 7. Menghapus Data Kualitas Udara
![alt text](image-8.png)

### 8. Mendapatkan Data Alert
![alt text](image-9.png)

### 9. Mendapatkan Rekomendasi Kualitas Udara
![alt text](image-10.png)

### 10. Logout
![alt text](image-11.png)

---

## Cara Instalasi

Ikuti langkah-langkah berikut untuk menginstal dan menjalankan API ini di mesin lokal Anda:

1. **Clone repository:**
   ```bash
   git clone https://github.com/Barmono-Berutu/Barmono_Berutu_Golang_Mini_Project.git
2. **Sebelum menjalankan kode, jalankan perintah ini terlebih dahulu:**
   ```bash
   cp .env.example .env
3. **Setelah mengubah semua data yang ada di .env, silakan jalankan kode:**
   ```bash
   go run main.go
