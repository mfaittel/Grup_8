package main

import (
	"fmt"
	"strings"
)

// Struct untuk soal
type Soal struct {
	ID         string
	Pertanyaan string
	Pilihan    [4]string
	Jawaban    string
	Benar      int
	Salah      int
}

// Struct untuk peserta
type Peserta struct {
	Nama string
	Skor int
}

var (
	bankSoal      = make([]Soal, 0)
	dataPeserta   [100]Peserta
	passwordAdmin = "admin123"
	jumlahPeserta = 0
)

// Fungsi Admin - Tambah Soal
func tambahSoal() {
	var id, pertanyaan, jawaban string
	var pilihan [4]string
	opsi := [4]string{"A", "B", "C", "D"}

	fmt.Print("Masukkan ID Soal: ")
	fmt.Scanln(&id)
	id = strings.TrimSpace(id)

	// Cek apakah ID sudah ada
	for _, soal := range bankSoal {
		if soal.ID == id {
			fmt.Println("Soal dengan ID ini sudah ada!")
			return
		}
	}

	fmt.Print("Masukkan Pertanyaan: ")
	fmt.Scanln(&pertanyaan)
	pertanyaan = strings.TrimSpace(pertanyaan)

	// Input pilihan
	for i := 0; i < 4; i++ {
		fmt.Printf("Opsi %s: ", opsi[i])
		fmt.Scanln(&pilihan[i])
		pilihan[i] = strings.TrimSpace(pilihan[i])
	}

	// Input kunci jawaban
	fmt.Print("Masukkan Kunci Jawaban (A-D): ")
	fmt.Scanln(&jawaban)
	jawaban = strings.ToUpper(strings.TrimSpace(jawaban))

	// Validasi kunci jawaban
	if jawaban != "A" && jawaban != "B" && jawaban != "C" && jawaban != "D" {
		fmt.Println("Kunci jawaban tidak valid! Harus antara A-D.")
		return
	}

	// Tambahkan soal ke bank soal
	bankSoal = append(bankSoal, Soal{ID: id, Pertanyaan: pertanyaan, Pilihan: pilihan, Jawaban: jawaban})
	fmt.Println("Soal berhasil ditambahkan!")
}

// Fungsi Admin - Hapus Soal
func hapusSoal() {
	var id string
	fmt.Print("Masukkan ID Soal yang akan dihapus: ")
	fmt.Scanln(&id)
	index := -1
	for i, soal := range bankSoal {
		if soal.ID == id {
			index = i
			break
		}
	}

	if index == -1 {
		fmt.Println("Soal dengan ID tersebut tidak ditemukan!")
		return
	}
	bankSoal = append(bankSoal[:index], bankSoal[index+1:]...)
	fmt.Println("Soal berhasil dihapus!")
}

// Fungsi Admin - Update Soal
func updateSoal() {
	var id, pertanyaan, jawaban string
	var pilihan [4]string
	fmt.Print("Masukkan ID Soal yang akan diperbarui: ")
	fmt.Scanln(&id)

	index := -1
	for i, soal := range bankSoal {
		if soal.ID == id {
			index = i
			break
		}
	}

	if index == -1 {
		fmt.Println("Soal dengan ID tersebut tidak ditemukan!")
		return
	}
	fmt.Print("Masukkan Pertanyaan Baru: ")
	fmt.Scanln(&pertanyaan)

	for i := 0; i < 4; i++ {
		fmt.Printf("Opsi %d Baru: ", i+1)
		fmt.Scanln(&pilihan[i])
	}
	fmt.Print("Masukkan Kunci Jawaban Baru (1-4): ")
	fmt.Scanln(&jawaban)

	bankSoal[index] = Soal{
		ID:         id,
		Pertanyaan: pertanyaan,
		Pilihan:    pilihan,
		Jawaban:    jawaban,
		Benar:      bankSoal[index].Benar,
		Salah:      bankSoal[index].Salah,
	}

	fmt.Println("Soal berhasil diperbarui!")
}

// Fungsi Admin - Login
func loginAdmin() bool {
	var password string
	fmt.Print("Masukkan Password Admin: ")
	fmt.Scanln(&password)

	if password == passwordAdmin {
		fmt.Println("Login berhasil!")
		return true
	}
	fmt.Println("Password salah!")
	return false
}

// Fungsi Peserta - Mendaftar
func daftarPeserta() {
	var nama string
	fmt.Print("Masukkan Nama Anda: ")
	fmt.Scanln(&nama)

	for i := 0; i < jumlahPeserta; i++ {
		if dataPeserta[i].Nama == nama {
			fmt.Println("Nama sudah terdaftar!")
			return
		}
	}

	if jumlahPeserta < len(dataPeserta) {
		dataPeserta[jumlahPeserta] = Peserta{Nama: nama, Skor: 0}
		jumlahPeserta++
		fmt.Println("Berhasil mendaftar!")
	} else {
		fmt.Println("Kapasitas peserta penuh!")
	}
}

// Fungsi Peserta - Ikut Kuis
func ikutKuis() {
	var nama, jawaban string
	opsi := [4]string{"A", "B", "C", "D"}

	fmt.Print("Masukkan Nama Anda: ")
	fmt.Scanln(&nama)
	nama = strings.TrimSpace(nama)

	// Cari peserta
	index := -1
	for i := 0; i < jumlahPeserta; i++ {
		if dataPeserta[i].Nama == nama {
			index = i
			break
		}
	}

	if index == -1 {
		fmt.Println("Anda belum terdaftar!")
		return
	}

	skor := 0
	fmt.Println("\nMulai Kuis!")
	for i, soal := range bankSoal {
		fmt.Printf("\n%d. %s\n", i+1, soal.Pertanyaan)
		for j, pilihan := range soal.Pilihan {
			fmt.Printf("%s. %s\n", opsi[j], pilihan)
		}
		fmt.Print("Jawaban Anda (A-D): ")
		fmt.Scanln(&jawaban)
		jawaban = strings.ToUpper(strings.TrimSpace(jawaban))

		if jawaban == soal.Jawaban {
			fmt.Println("Benar!")
			bankSoal[i].Benar++
			skor += 10
		} else {
			fmt.Printf("Salah! Jawaban benar adalah %s\n", soal.Jawaban)
			bankSoal[i].Salah++
		}
	}

	fmt.Printf("Skor Anda: %d\n", skor)
	if skor > dataPeserta[index].Skor {
		dataPeserta[index].Skor = skor
	}
}

// Fungsi Peserta - Leaderboard
func papanPeringkat() {
	fmt.Println("\nPapan Peringkat:")
	for i := 0; i < jumlahPeserta-1; i++ {
		for j := 0; j < jumlahPeserta-i-1; j++ {
			if dataPeserta[j].Skor < dataPeserta[j+1].Skor {
				dataPeserta[j], dataPeserta[j+1] = dataPeserta[j+1], dataPeserta[j]
			}
		}
	}
	for i := 0; i < jumlahPeserta; i++ {
		fmt.Printf("%d. %s - Skor: %d\n", i+1, dataPeserta[i].Nama, dataPeserta[i].Skor)
	}
}

// Menu Utama
func main() {
	for {
		fmt.Println("\nAplikasi Kuis - Siapa yang Ingin Jadi Jutawan")
		fmt.Println("1. Admin")
		fmt.Println("2. Peserta")
		fmt.Println("3. Keluar")
		fmt.Print("Pilih peran: ")

		var choice string
		fmt.Scanln(&choice)

		switch choice {
		case "1":
			if loginAdmin() {
				menuAdmin()
			}
		case "2":
			menuPeserta()
		case "3":
			fmt.Println("Terima kasih!")
			return
		default:
			fmt.Println("Pilihan tidak valid!")
		}
	}
}

func menuAdmin() {
	for {
		fmt.Println("\nMenu Admin:")
		fmt.Println("1. Tambah Soal")
		fmt.Println("2. Hapus Soal")
		fmt.Println("3. Update Soal")
		fmt.Println("4. Keluar")
		fmt.Print("Pilih menu: ")

		var choice string
		fmt.Scanln(&choice)

		switch choice {
		case "1":
			tambahSoal()
		case "2":
			hapusSoal()
		case "3":
			updateSoal()
		case "4":
			return
		default:
			fmt.Println("Pilihan tidak valid!")
		}
	}
}

func menuPeserta() {
	for {
		fmt.Println("\nMenu Peserta:")
		fmt.Println("1. Daftar")
		fmt.Println("2. Ikut Kuis")
		fmt.Println("3. Papan Peringkat")
		fmt.Println("4. Keluar")
		fmt.Print("Pilih menu: ")

		var choice string
		fmt.Scanln(&choice)

		switch choice {
		case "1":
			daftarPeserta()
		case "2":
			ikutKuis()
		case "3":
			papanPeringkat()
		case "4":
			return
		default:
			fmt.Println("Pilihan tidak valid!")
		}
	}
}
