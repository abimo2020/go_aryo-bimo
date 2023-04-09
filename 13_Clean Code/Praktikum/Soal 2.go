package main

// Penggunaan dua kata untuk suatu property harus menggunakan camelCase atau pascalCase agar mudah dibaca
type kendaraan struct {
	totalRoda       int
	kecepatanPerJam int
}

type mobil struct {
	kendaraan
}

// Penggunaan dua kata untuk suatu property harus menggunakan camelCase atau pascalCase agar mudah dibaca
func (m *mobil) berjalan() {
	m.tambahKecepatan(10)
}

// Penggunaan dua kata untuk suatu property harus menggunakan camelCase atau pascalCase agar mudah dibaca
func (m *mobil) tambahKecepatan(kecepatanBaru int) {
	m.kecepatanPerJam = m.kecepatanPerJam + kecepatanBaru
}

func main() {

	mobilCepat := mobil{}

	// Membuat sebuah perulangan untuk menghindari perulangan penulisan pemanggilan fungsi
	totalJalan := 3
	for i := 0; i < totalJalan; i++ {
		mobilCepat.berjalan()
	}

	mobilLamban := mobil{}
	mobilLamban.berjalan()
}
