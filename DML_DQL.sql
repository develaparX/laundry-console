SELECT * FROM pelanggan;

INSERT INTO pelanggan(id_pelanggan, nama_pelanggan, no_handphone) VALUES (1, 'Jessica', '08133549999');

SELECT * FROM layanan;

INSERT INTO layanan(id_layanan, layanan, satuan, harga)
VALUES (1, 'Cuci + Setrika', 'KG',7000),
(2, 'Laundry Bedcover', 'Buah',50000), 
(3, 'Laundry Boneka', 'Buah', 25000);

SELECT * FROM transaksi;

INSERT INTO transaksi(id_transaksi, id_pelanggan, tanggal_masuk, tanggal_selesai, nama_penerima)
VALUES (1, 1, '2022-08-18', '2022-08-20', 'Mirna');

SELECT * FROM detail_transaksi;

INSERT INTO detail_transaksi(id_detail, id_transaksi, id_layanan, jumlah_layanan)
VALUES (1, 1, 1, 5 ),(2, 1, 2, 1 ),(3, 1, 3, 2 );


--saya melakukan semua input data ini, bagaimana melihat transaksi dengan id 1 dengan format (id_transaksi, nama_pelanggan, layanan, satuan, harga,jumlah, tanggal mulai, tanggal selesai, penerima, total(jumlah * harga) )

SELECT 
    t.id_transaksi,
    p.nama_pelanggan,
	p.no_handphone,
    l.layanan,
    l.satuan,
    l.harga,
    dt.jumlah_layanan AS jumlah,
    t.tanggal_masuk,
    t.tanggal_selesai,
    t.nama_penerima,
    (l.harga * dt.jumlah_layanan) AS total
FROM 
    transaksi t
JOIN 
    pelanggan p ON t.id_pelanggan = p.id_pelanggan
JOIN 
    detail_transaksi dt ON t.id_transaksi = dt.id_transaksi
JOIN 
    layanan l ON dt.id_layanan = l.id_layanan;
	
	




