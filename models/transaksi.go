package models

type Transaksi struct {
	ID             int
	Pelanggan      Pelanggan
	Layanan        Layanan
	Jumlah         int
	TanggalMulai   string
	TanggalSelesai string
	Penerima       string
	Total          int
}
