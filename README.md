# Tugas Kecil 1 IF2211 Strategi Algoritma

## Daftar Konten
* [Informasi Umum](#informasi-umum)
* [Deskripsi Singkat](#deskripsi-singkat)
* [Kebutuhan Program](#kebutuhan-program)
* [Setup dan Penggunaan](#setup-dan-penggunaan)
* [Struktur File](#struktur-file)
* [Identitas](#identitas)

## Informasi Umum
Repository ini berisi program penyelesaian Cyberpunk 2077 Breach Protocol dengan Algoritma Brute Force.

## Deskripsi Singkat
Pemain menyediakan matriks berisi dua karakter alfanumerik dan beberapa sekuens beserta reward dari masing-masing sekuens. Program akan mencari solusi terbaik (reward maksimal) dari matriks dimulai dari salah satu kolom pada baris pertama, lalu bergerak vertikal dan horizontal secara bergantian dengan algoritma brute force.<br><br>Contoh terdapat pada laman https://cyberpunk-hacker.com/.

## Kebutuhan Program
Program dibuat dengan bahasa pemrograman Go. Untuk menjalankan program ini, pastikan bahwa Go sudah terinstall. Gunakan perintah di bawah ini untuk mengecek versi Go yang terinstall
```
$ go version
```
Jika belum terinstall, download dan ikuti petunjuk pada laman [ini](https://go.dev/doc/install)
## Setup dan Penggunaan
### 1. Clone repository ini dengan perintah:
```
$ git clone https://github.com/julianchandras/Tucil1_13522080.git
```
### 2. Kompilasi program
Untuk Windows, kompilasi program dengan perintah:
```
$ go build -o bin/main.exe src/main.go src/functions.go src/txtprocessor.go src/type.go
```
Sedangkan pada Linux, kompilasi program dengan perintah:
```
$ go build -o bin/main src/main.go src/functions.go src/txtprocessor.go src/type.go
```
### 3. Jalankan program<br>
Untuk Windows, jalankan program dengan perintah:
```
$ .\bin\main.exe
```
Sedangkan pada Linux, jalankan program dengan perintah:
```
$ bin/main
```
Pada Linux, juga langsung dapat jalankan perintah (tanpa perlu kompilasi program terlebih dahulu)
```
$ ./run.sh
```
### 4. Ikuti arahan program dan berikan input.
### 5. Jika ingin menginput dari file, gunakan format berikut pada file txt dan tempatkan pada folder test.

```
buffer_size
matrix_width matrix_height
matrix
number_of_sequences
sequences_1
sequences_1_reward
sequences_2
sequences_2_reward
…
sequences_n
sequences_n_reward
```

## Struktur File
Direktori tugas ini memiliki struktur berikut:
```
.
├── README.md
├── bin
│   ├── main
│   └── main.exe
├── doc
│   └── Tucil1_K2_13522080_Julian Chandra Sutadi.pdf
├── go.mod
├── run.sh
├── src
│   ├── functions.go
│   ├── main.go
│   ├── txtprocessor.go
│   └── type.go
└── test
    ├── 1-result.txt
    ├── 1.txt
    ├── 2-result.txt
    ├── 2.txt
    ├── 3-result.txt
    ├── 3.txt
    ├── 4-result.txt
    ├── 4.txt
    ├── 5-result.txt
    ├── 5.txt
    ├── 6-result.txt
    └── 6.txt
```
## Identitas
Program ini dibuat oleh Julian Chandra Sutadi (13522080)
