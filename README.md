
# **Project Mini Golang AirQuality**

API ini adalah proyek mini yang dibangun dengan Golang untuk memantau dan mengelola data kualitas udara. Fitur utama meliputi:
- Registrasi pengguna.
- Login/Logout.
- Pengelolaan data kualitas udara.
- Analisis berbasis AI untuk alert dan rekomendasi tindakan.

---

## **Daftar Isi**
1. [Entity Relationship Diagram (ERD)](#entity-relationship-diagram-erd)
2. [High-Level Architecture Diagram (HLA)](#high-level-architecture-diagram-hla)
3. [Dokumentasi API](#dokumentasi-api)
4. [Cara Instalasi](#cara-instalasi)
5. [Teknologi yang Digunakan](#teknologi-yang-digunakan)
6. [Contoh Hasil Postman](#contoh-hasil-postman)

---

## **Entity Relationship Diagram (ERD)**

Berikut adalah Entity Relationship Diagram (ERD) untuk API AirQuality yang mengelola data pengguna, kualitas udara, dan hasil analisis berbasis AI.

### **ERD Diagram**
![ERD](image.png)

### **Penjelasan ERD**
#### **Tabel dan Relasi**
1. **Users**
   - **Tabel ini menyimpan informasi pengguna aplikasi.**
     - `id` (integer): Primary key.
     - `email` (varchar): Alamat email pengguna.
     - `password` (varchar): Hash password pengguna.
   - **Relasi:**
     - Satu pengguna dapat mengakses banyak data kualitas udara.

2. **AirQuality**
   - **Tabel ini menyimpan data kualitas udara.**
     - `id` (integer): Primary key.
     - `user_id` (integer): Foreign key ke tabel Users.
     - `location` (varchar): Lokasi pengukuran.
     - `aqi` (integer): Indeks kualitas udara.
     - `pollutants` (varchar): Detail polutan yang terdeteksi.
     - `recorded_at` (datetime): Waktu pencatatan data kualitas udara.
   - **Relasi:**
     - Satu pengguna memiliki banyak data kualitas udara.

3. **AI_Response**
   - **Tabel ini menyimpan hasil analisis dan rekomendasi dari AI.**
     - `id` (integer): Primary key.
     - `air_quality_id` (integer): Foreign key ke tabel AirQuality.
     - `alert_level` (varchar): Tingkat kewaspadaan berdasarkan data AI.
     - `recommendations` (text): Rekomendasi tindakan dari AI.
   - **Relasi:**
     - Satu data kualitas udara memiliki satu hasil analisis AI.

#### **Relasi Antar Tabel**
- **Users → AirQuality:** Relasi one-to-many.
- **AirQuality → AI_Response:** Relasi one-to-one.

---

## **High-Level Architecture Diagram (HLA)**

### **Deskripsi Umum**
High-Level Architecture Diagram (HLA) memberikan gambaran menyeluruh tentang alur kerja dan komponen utama dalam sistem backend API AirQuality.

### **Diagram Visual**
![HLA](hla_project.png)

### **Komponen Utama**

1. **User**
   - Representasi pengguna aplikasi (melalui web atau Postman).
   - Mengirimkan permintaan ke API untuk mengakses fitur seperti autentikasi, pengelolaan data kualitas udara, dan rekomendasi.

2. **API Gateway**
   - Berfungsi sebagai pintu masuk utama untuk semua permintaan API.
   - Mengatur routing, autentikasi, dan load balancing.

3. **Backend (Echo Framework)**
   - Berfungsi sebagai inti sistem untuk mengelola logika bisnis dan komunikasi dengan komponen lain.
   - **Modul Utama:**
     - **Autentikasi:** Mengelola login dan registrasi pengguna menggunakan JWT.
     - **Manajemen Kualitas Udara:** CRUD data kualitas udara.
     - **Analisis AI:** Mengintegrasikan AI untuk alert dan rekomendasi.

4. **Database (MySQL)**
   - Menyimpan data pengguna, kualitas udara, dan hasil analisis AI.

5. **AI Module**
   - Menganalisis data kualitas udara untuk memberikan alert dan rekomendasi berbasis logika AI.

---

## **Dokumentasi API**

### **Fitur Autentikasi**
| No | Method | Endpoint    | Request Body                              | Deskripsi                               |
|----|--------|-------------|------------------------------------------|-----------------------------------------|
| 1  | POST   | `/register` | `{ "email": "test@example.com", "password": "12345" }` | Registrasi pengguna baru               |
| 2  | POST   | `/login`    | `{ "email": "test@example.com", "password": "12345" }` | Login pengguna dan mendapatkan token   |

### **Fitur Manajemen Kualitas Udara**
| No | Method  | Endpoint                | Request Body                                                                                     | Deskripsi                                         |
|----|---------|-------------------------|-------------------------------------------------------------------------------------------------|-------------------------------------------------|
| 1  | GET     | `/airquality`          | -                                                                                               | Mendapatkan seluruh data kualitas udara          |
| 2  | GET     | `/airquality/{id}`     | -                                                                                               | Mendapatkan data kualitas udara berdasarkan ID   |
| 3  | POST    | `/airquality`          | `{ "location": "Jakarta", "aqi": 150, "pollutants": "PM2.5", "recorded_at": "2024-11-17" }`    | Menambahkan data kualitas udara baru            |
| 4  | PUT     | `/airquality/{id}`     | `{ "location": "Bandung", "aqi": 100, "pollutants": "PM10", "recorded_at": "2024-11-17" }`     | Memperbarui data kualitas udara berdasarkan ID  |
| 5  | DELETE  | `/airquality/{id}`     | -                                                                                               | Menghapus data kualitas udara berdasarkan ID     |

### **Fitur AI untuk Alert dan Rekomendasi**
| No | Method | Endpoint                | Request Body | Deskripsi                                    |
|----|--------|-------------------------|--------------|----------------------------------------------|
| 1  | GET    | `/airquality/alert`    | -            | Mendapatkan alert tingkat kewaspadaan dari AI |
| 2  | GET    | `/airquality/rekomendasi` | -            | Mendapatkan rekomendasi berbasis AI          |

---

## **Cara Instalasi**

1. **Clone Repository**
   ```bash
   git clone https://github.com/Barmono-Berutu/Barmono_Berutu_Golang_Mini_Project.git
   cd Barmono_Berutu_Golang_Mini_Project
   ```

2. **Setup File Konfigurasi**
   ```bash
   cp .env.example .env
   ```
   Sesuaikan parameter di `.env` seperti `DB_HOST`, `DB_USER`, `DB_PASSWORD`, dan `JWT_SECRET`.

3. **Install Dependencies**
   ```bash
   go mod tidy
   ```

4. **Run Server**
   ```bash
   go run main.go
   ```

---

## **Teknologi yang Digunakan**
- **Bahasa Pemrograman**: Golang
- **Framework**: Echo Framework
- **Database**: MySQL
- **Autentikasi**: JSON Web Token (JWT)
- **AI**: Google Generative AI

---

## **Contoh Hasil Postman**

### **1. User Autentikasi**
- **Registrasi**  
  ![register](image-2.png)
- **Login**  
  ![login](image-3.png)

### **2. CRUD AirQuality**
- **Menambahkan Data**  
  ![add](image-4.png)
- **Melihat Semua Data**  
  ![view-all](image-5.png)
- **Melihat Data Berdasarkan ID**  
  ![view-id](image-6.png)
- **Mengedit Data**  
  ![edit](image-7.png)
- **Menghapus Data**  
  ![delete](image-8.png)

### **3. Alert dan Rekomendasi**
- **Alert dari AI**  
  ![alert](image-9.png)
- **Rekomendasi dari AI**  
  ![recommendation](image-10.png)

---

Dokumen ini memberikan panduan lengkap untuk memahami, menginstal, dan menggunakan API **Project Mini Golang AirQuality**. Terimakasih
