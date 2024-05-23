# Laundry Management System

## Deskripsi
Aplikasi manajemen laundry sederhana yang ditulis dalam bahasa Go. Aplikasi ini memungkinkan pengguna untuk melakukan operasi dasar pada tabel `pelanggan`, `layanan`, dan `transaksi`.

## Fitur
- Melihat daftar pelanggan, layanan, dan transaksi
- Melihat transaksi berdasarkan ID
- Menambahkan pelanggan baru
- Menambahkan layanan baru
- Menambahkan transaksi baru (dengan beberapa layanan)

## Prasyarat
- Go 1.15 atau yang lebih baru
- PostgreSQL

## Instalasi
1. Clone repository ini:
    ```sh
    git clone https://git.enigmacamp.com/enigma-20/arfian-saiful-rifai/challenge-godb.git
    ```

2. Buat database di PostgreSQL dan sesuaikan string koneksi di `database/database.go`:

3. Jalankan script DDL untuk membuat tabel dan Jalankan script DML untuk mengisi data, script bernama `DDL.sql dan DML_DQL.sql`
   
4. Jalankan aplikasi:
    ```sh
    go run main.go
    ```

## Penggunaan
1. Jalankan aplikasi.
2. Pilih menu yang diinginkan dengan memasukkan angka yang sesuai.
3. Ikuti petunjuk yang diberikan untuk memasukkan data.

### Menu Aplikasi
1. **Menu Pelanggan**: Menampilkan daftar pelanggan yang didalamnya berisi Lihat Pelanggan, Tambah Pelanggan, Ubah Pelanggan dan Hapus Pelanggan.
2. **Menu Layanan**: Menampilkan daftar layanan, yang didalamnya berisi Lihat Semua Layanan, Tambahkan Layanan, Ubah Layanan, Hapus Layanan.
3. **Menu Transaksi**: Menampilkan semua transaksi, yang didalamnya berisi Lihat Semua Transaksi, Lihat Transaksi Sesuai ID, Masukkan Transaksi Baru, Lihat Layanan, Lihat Pelanggan.

### Contoh Penggunaan
1. **Lihat Pelanggan**:
    - Pilih menu 1.
    - Aplikasi akan menampilkan daftar pelanggan yang ada.

2. **Tambah Pelanggan**:
    - Pilih menu 6.
    - Masukkan ID pelanggan, nama pelanggan, dan no handphone saat diminta.
    - Aplikasi akan menambahkan pelanggan baru jika data valid dan ID belum ada.

3. **Tambah Transaksi**:
    - Pilih menu 5.
    - Masukkan ID transaksi, ID pelanggan, tanggal masuk, tanggal selesai, dan nama penerima.
    - Tambahkan layanan dengan memasukkan ID layanan dan jumlah layanan (input 0 untuk berhenti menambahkan layanan).
    - Aplikasi akan menyimpan transaksi baru jika semua data valid dan ID transaksi belum ada.

4. **Lihat Transaksi Berdasarkan ID**:
    - Pilih menu 4.
    - Masukkan ID transaksi yang ingin dilihat.
    - Aplikasi akan menampilkan detail transaksi termasuk total harga keseluruhan.

