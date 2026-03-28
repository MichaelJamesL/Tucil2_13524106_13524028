# Tucil 3 IF2211 Strategi Algoritma — Voxelisasi 3D dengan Divide and Conquer

## Identitas Pembuat

| NIM       | Nama                    |
|-----------|---------------------|
| 13524028 | Muhammad Nur Majiid |
| 13524106 | Michael James Liman |

## Penjelasan Singkat Program

Program ini mengimplementasikan **Voxelisasi 3D** menggunakan pendekatan algoritma **Divide and Conquer** dengan struktur data **Octree**. Voxelisasi adalah proses mengubah representasi model 3D berbasis permukaan (mesh segitiga dari file OBJ) menjadi representasi berbasis volume (kumpulan kubus-kubus kecil / voxel).

Alur kerja program:
1. Membaca file model 3D berformat `.obj` untuk mengekstrak vertex dan segitiga
2. Menghitung bounding box yang melingkupi seluruh model
3. Membangun pohon Octree secara rekursif dengan **concurrency** (goroutine pada level dangkal) untuk membagi ruang 3D dan mendeteksi area yang terisi model
4. Mengumpulkan titik-titik pusat voxel dari daun-daun Octree
5. Mengekspor hasil voxelisasi ke file `.obj` baru

## Dependency

Program ini ditulis dalam bahasa **Go (Golang)** dan hanya menggunakan **standard library**, tanpa dependency eksternal.

| Dependency         | Versi Minimum | Keterangan                    |
|--------------------|---------------|-------------------------------|
| Go (Golang)        | 1.18          | Bahasa pemrograman utama      |

Package standard library yang digunakan:
- `bufio` — Buffered I/O untuk membaca dan menulis file
- `fmt` — Formatted I/O (parsing dan output)
- `math` — Fungsi matematika (Min, Max, Pow)
- `os` — Operasi file (Open, Create)
- `strconv` — Konversi string ke integer
- `strings` — Manipulasi string (TrimSpace, Fields, Index)
- `sync` — Sinkronisasi concurrency (WaitGroup)

## Cara Instalasi dan Menjalankan Program

### Prasyarat
Pastikan **Go** sudah terinstal di sistem Anda:
```bash
go version
```
Jika belum terinstal, unduh dari [https://go.dev/dl/](https://go.dev/dl/)

### Langkah Instalasi
1. Clone repository ini:
```bash
git clone https://github.com/MichaelJamesL/Tucil3_Stima.git
cd Tucil3_Stima
```

2. (Opsional) Verifikasi modul:
```bash
cat go.mod
```
sesuaikan dengan versi go yang anda gunakan

### Menjalankan Program
```bash
go run main.go voxelization.go
```

Program akan membaca file `28-pumpkin/pumpkin/pumpkin_.obj` dan menghasilkan file output `voxelized.obj` di direktori yang sama.

### Build (Opsional)
```bash
go build -o voxelizer main.go voxelization.go
./voxelizer
```

## Directory Tree

```
Tucil3_Stima/
├── go.mod                         # Go module definition
├── main.go                        # Program utama (entry point)
├── voxelization.go                # Implementasi algoritma voxelisasi & Octree
├── voxelized.obj                  # Output hasil voxelisasi (generated)
└── 28-pumpkin/                    # Data model 3D input
    └── pumpkin/
        ├── pumpkin_.obj           # File model 3D format OBJ (input utama)
        ├── pumpkin_.mtl           # Material definition file
        ├── pumpkin_.c4d           # File Cinema 4D (sumber model)
        ├── Sphere2.jpg            # Texture model
        └── N_map.jpg              # Normal map texture
```
