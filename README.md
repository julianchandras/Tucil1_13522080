# Tugas Kecil 1 IF2211 Strategi Algoritma
> Julian Chandra Sutadi (13522080)

## Daftar Konten
* [Informasi Umum](#informasi-umum)
* [Deskripsi Singkat](#deskripsi-singkat)
* [Setup dan Penggunaan](#setup-dan-penggunaan)
* [Struktur File](#struktur-file)

## Informasi Umum
Repository ini berisi program penyelesaian Cyberpunk 2077 Breach Protocol dengan Algoritma Brute Force.

## Deskripsi Singkat
Pemain menyediakan matriks berisi dua karakter alfanumerik dan beberapa sekuens beserta reward dari masing-masing sekuens. Program akan mencari solusi terbaik (reward maksimal) dari matriks dimulai dari salah satu kolom pada baris pertama, lalu bergerak vertikal dan horizontal secara bergantian dengan algoritma brute force.<br>Program dibuat dalam bahasa pemrograman Go.<br><br>Contoh terdapat pada laman https://cyberpunk-hacker.com/.

## Setup dan Penggunaan
1. Clone repository ini dengan perintah
```
git clone https://github.com/julianchandras/Tucil1_13522080.git
```
2. Jika ingin mengkompilasi program jalankan perintah
```
go build -o bin/main src/main.go src/functions.go src/txtprocessor.go src/type.go
```
3. Pada Windows, jalankan program dengan perintah:
```
.\bin\main.exe
```
Sedangkan pada Linux, jalankan program dengan perintah:
```
bin/main
```
Untuk Linux, juga langsung dapat menjalankan (tanpa perlu compile terlebih dahulu)
```
./run.sh
```
4. Ikuti arahan program dan berikan input.
5. Jika ingin menginput dari file, gunakan format berikut pada file txt dan tempatkan pada folder test.

```
buffer_size<br>
matrix_width matrix_height<br>
matrix<br>
number_of_sequences<br>
sequences_1<br>
sequences_1_reward<br>
sequences_2<br>
sequences_2_reward<br>
…<br>
sequences_n<br>
sequences_n_reward<br>
```

## Struktur File
Direktori tugas ini memiliki struktur berikut:

.
├── README.md
├── bin
│   ├── main
│   └── main.exe
├── doc
│   ├── Tucil1_K2_13522080_Julian Chandra Sutadi.docx
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
