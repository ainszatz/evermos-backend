# Store Management API

Store Management API adalah sistem backend berbasis Go dengan framework **Fiber** dan menggunakan **Gorm** sebagai ORM untuk berinteraksi dengan database. Aplikasi ini dirancang untuk mengelola toko dan produk, serta mendukung otorisasi pengguna berbasis JWT.

## 🚀 Fitur

### 🔐 **Autentikasi & Otorisasi**
- Service **Login dan Register** menggunakan JWT
- **Toko otomatis terbuat** ketika user mendaftar
- Middleware untuk validasi token pengguna
- Role-based access control (**RBAC**) untuk memastikan hanya pemilik toko yang dapat mengelola tokonya
- **User tidak dapat mengakses atau mengubah data user lain**

### 🏪 **Manajemen Toko**
- **CRUD Toko**: Tambah, edit, hapus, dan lihat detail toko
- **Upload Logo**: Pemilik toko dapat mengunggah atau mengubah logo tokonya
- **Keamanan**:
  - Hanya pemilik toko yang dapat mengelola tokonya
  - User tidak dapat mengelola toko user lain

### 📦 **Manajemen Produk**
- **CRUD Produk**: Tambah, edit, hapus, dan lihat daftar produk dalam toko
- **User hanya bisa mengelola produk di tokonya sendiri**
- **Tabel log produk menyimpan data setiap transaksi**

### 💳 **Manajemen Transaksi**
- **User hanya bisa melakukan transaksi produk yang ada**
- **Alamat diperlukan untuk pengiriman produk**
- **Tabel log produk diisi otomatis setiap transaksi**
- **User tidak dapat mengelola transaksi user lain**

### 🏠 **Manajemen Alamat**
- **CRUD Alamat Pengiriman**
- **User hanya dapat mengelola alamatnya sendiri**

### 📂 **Manajemen Kategori**
- **CRUD Kategori** (Hanya dapat dikelola oleh Admin)
- **Hanya admin yang dapat menambah/menghapus kategori** (Perubahan admin dilakukan langsung di database)

### 📊 **Filtering & Pagination**
- **Menerapkan Pagination** seperti di Postman
- **Menerapkan Filtering Data** berdasarkan harga, tanggal, dll.

### 📤 **File Upload**
- **Upload Logo Toko**
- **Upload Gambar Produk**
- **Hanya pemilik toko yang dapat mengunggah gambar produknya sendiri**

### 📊 **Monitoring & Logging**
- Logging aktivitas API untuk debugging
- Error handling dengan pesan yang jelas

---

## 🛠 Teknologi yang Digunakan
- **Golang (Fiber)** - Web framework yang ringan dan cepat
- **GORM** - ORM untuk interaksi dengan database
- **MySQL** - Database utama
- **JWT (JSON Web Token)** - Untuk autentikasi dan otorisasi
- **Chat GPT** - Membantu dalam debugging, optimasi kode, pembuatan dokumentasi API, serta memberikan solusi dan best practices dalam pengembangan backend

---

## 🔧 Instalasi & Menjalankan Proyek

### 1️⃣ **Clone Repository**
```sh
 git clone https://github.com/ainszatz/evermos-backend
 cd evermos-backend
```

### 2️⃣ **Setup Environment Variables**
Buat file `.env` dan isi dengan konfigurasi berikut:
```env
DB_USER=DB_USERNAME
DB_PASS=DB_PASSWORD
DB_HOST=DB_HOST
DB_PORT=DB_PORT
DB_NAME=DB_NAME
APP_PORT=APP_PORT
JWT_SECRET=supersecretkey
```

### 3️⃣ **Jalankan Server**
```sh
go run main.go
```

---

## 📌 API Endpoints

### 🔐 **Auth Routes**
| Method | Endpoint        | Description          |
|--------|----------------|----------------------|
| POST   | `/login`        | Login user          |
| POST   | `/logout`       | Logout user         |

### 🏪 **Store Routes**
| Method | Endpoint                | Description                  |
|--------|-------------------------|------------------------------|
| POST   | `/stores`               | Tambah toko baru             |
| GET    | `/stores/:id`           | Lihat detail toko            |
| PUT    | `/stores/:id`           | Update toko (hanya pemilik)  |
| DELETE | `/stores/:id`           | Hapus toko (hanya pemilik)   |
| POST   | `/stores/:id/logo`      | Upload atau update logo toko |

### 📦 **Product Routes**
| Method | Endpoint                | Description                  |
|--------|-------------------------|------------------------------|
| POST   | `/stores/:id/products`  | Tambah produk ke toko        |
| GET    | `/stores/:id/products`  | Lihat semua produk toko      |
| PUT    | `/products/:id`         | Update produk                |
| DELETE | `/products/:id`         | Hapus produk                 |

### 🏠 **Address Routes**
| Method | Endpoint                | Description                  |
|--------|-------------------------|------------------------------|
| POST   | `/addresses`            | Tambah alamat pengguna       |
| GET    | `/addresses`            | Lihat alamat pengguna        |
| PUT    | `/addresses/:id`        | Update alamat (hanya pemilik)|
| DELETE | `/addresses/:id`        | Hapus alamat (hanya pemilik) |

### 📊 **Category Routes (Admin Only)**
| Method | Endpoint                | Description                  |
|--------|-------------------------|------------------------------|
| POST   | `/categories`           | Tambah kategori (Admin)      |
| GET    | `/categories`           | Lihat daftar kategori        |
| PUT    | `/categories/:id`       | Update kategori (Admin)      |
| DELETE | `/categories/:id`       | Hapus kategori (Admin)       |

---

## 🤝 Kontribusi
Jika ingin berkontribusi, silakan fork repository ini dan buat pull request!

---

## 📄 Lisensi
Proyek ini menggunakan lisensi MIT. Silakan gunakan dan modifikasi sesuai kebutuhan.

🚀 **Happy Coding!**

