CREATE TABLE pelanggan (
    id_pelanggan INT PRIMARY KEY NOT NULL,
    nama_pelanggan VARCHAR(100) NOT NULL,
    no_handphone VARCHAR(15) NOT NULL
);

select * FROM pelanggan;

CREATE TABLE layanan(
	id_layanan INT PRIMARY KEY NOT NULL,
	layanan VARCHAR(50) NOT NULL,
	satuan VARCHAR(10) NOT NULL,
	HARGA INT 
);

SELECT * FROM layanan;


CREATE TABLE transaksi (
    id_transaksi INT PRIMARY KEY NOT NULL,
    id_pelanggan INT NOT NULL,
    id_layanan INT NOT NULL,
    jumlah_layanan INT NOT NULL,
    tanggal_masuk DATE NOT NULL,
    tanggal_selesai DATE NOT NULL,
    nama_penerima VARCHAR(100) NOT NULL,
    FOREIGN KEY (id_pelanggan) REFERENCES pelanggan(id_pelanggan),
    FOREIGN KEY (id_layanan) REFERENCES layanan(id_layanan)
);

CREATE TABLE detail_transaksi (
    id_detail SERIAL PRIMARY KEY,
    id_transaksi INT NOT NULL,
    id_layanan INT NOT NULL,
    jumlah_layanan INT NOT NULL,
    FOREIGN KEY (id_transaksi) REFERENCES transaksi(id_transaksi),
    FOREIGN KEY (id_layanan) REFERENCES layanan(id_layanan)
);