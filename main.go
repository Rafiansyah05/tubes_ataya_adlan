package main

import "fmt"

const NMAX int = 100

type Interaction struct {
	Mood    string
	Message string 
}

type SaranKesehatan struct {
	Topik    string 
	IsiSaran string 
}

type tabInteraksi [NMAX]Interaction
var jumlahInteraksi int = 0 

type tabSaran [NMAX]SaranKesehatan
var daftarSaran tabSaran    
var jumlahSaran int = 0    

func main() {
	var namaPengguna string
	var riwayatInteraksiPengguna tabInteraksi

	inisialisasiDaftarSaran() 

	// Judul Aplikasi

	fmt.Println("\n========== CHATBOT KESEHATAN MENTAL ==========")
	fmt.Println("Halo! Saya adalah teman bicara untuk kesehatan mental.")
	fmt.Println("\nSebelum kita mulai, boleh saya tahu namamu?")
	fmt.Print("> ")
	fmt.Scanln(&namaPengguna)

	fmt.Printf("\nHalo, %s! 😊\n", namaPengguna)

	var sesiAktif bool = true
	for sesiAktif {
		var pilihanMenuUtama string
		var pesanPengguna string
		var suasanaHatiPengguna string = "" 

		fmt.Println("\nBagaimana perasaanmu hari ini? (1-5, atau 6 untuk Opsi Lanjutan/Keluar)")
		fmt.Println("1. Bahagia 😊")
		fmt.Println("2. Sedih 😢")
		fmt.Println("3. Cemas 😰")
		fmt.Println("4. Marah 😠")
		fmt.Println("5. Lelah 😴")
		fmt.Println("6. Opsi Lanjutan / Keluar")
		fmt.Print("> ")
		fmt.Scanln(&pilihanMenuUtama)

		if pilihanMenuUtama == "6" {
			sesiAktif = tanganiOpsiLanjutan(&riwayatInteraksiPengguna)
			if !sesiAktif { 
				break
			}
			continue 
		}

		switch pilihanMenuUtama {
		case "1":
			suasanaHatiPengguna = "bahagia"
			fmt.Println("Wah, senang mendengarnya! Ada cerita menyenangkan apa hari ini?")
			fmt.Println("(Notes : Gunakan titik sebagai spasi. Contoh: Aku.habis.makan.enak)")
			fmt.Print("> ")
			fmt.Scanln(&pesanPengguna)
		case "2":
			suasanaHatiPengguna = "sedih"
			fmt.Println("Aku di sini untuk mendengarmu. Mau ceritakan apa yang membuatmu sedih?")
			fmt.Println("(Notes : Gunakan titik sebagai spasi. Contoh: Aku.habis.diputusin.pacar)")
			fmt.Print("> ")
			fmt.Scanln(&pesanPengguna)
		case "3":
			suasanaHatiPengguna = "cemas"
			fmt.Println("Kecemasan bisa terasa berat. Apa yang sedang mengganggu pikiranmu?")
			fmt.Println("(Notes : Gunakan titik sebagai spasi. Contoh: Besok.aku.akan.dikaderisasi)")
			fmt.Print("> ")
			fmt.Scanln(&pesanPengguna)
		case "4":
			suasanaHatiPengguna = "marah"
			fmt.Println("Marah adalah perasaan yang wajar. Apa yang memicunya?")
			fmt.Println("(Notes : Gunakan titik sebagai spasi. Contoh: Dia.menyebalkan)")
			fmt.Print("> ")
			fmt.Scanln(&pesanPengguna)
		case "5":
			suasanaHatiPengguna = "lelah"
			fmt.Println("Istirahat itu penting. Apakah kamu merasa terbebani akhir-akhir ini?")
			fmt.Println("(Notes : Gunakan titik sebagai spasi. Contoh: Baru.pulang.jogging)")
			fmt.Print("> ")
			fmt.Scanln(&pesanPengguna)
		default:
			fmt.Println("Maaf, aku tidak mengerti pilihanmu.")
			continue 
		}


		if suasanaHatiPengguna != "" && pesanPengguna != "" {
			tambahInteraksi(&riwayatInteraksiPengguna, suasanaHatiPengguna, pesanPengguna)
			dapatkanResponSuasanaHati(suasanaHatiPengguna)
			sarankanAktivitas(suasanaHatiPengguna)
			tawarkanRelaksasi(suasanaHatiPengguna)
		} else if suasanaHatiPengguna != "" && pesanPengguna == "" {
			fmt.Println("Kamu belum menceritakan apa-apa. Ceritakanlah!")
		}
	}

	fmt.Println("\nTerima kasih sudah berbicara denganku hari ini, ", namaPengguna, ".")
	fmt.Println("Ingatlah: Perasaanmu valid dan penting.")
	fmt.Println("Sampai jumpa lagi! ❤")
}

func tanganiOpsiLanjutan(A *tabInteraksi) bool {
	var pilihanMenuLanjutan string

	for {
		fmt.Println("\n--- Opsi Lanjutan & Riwayat ---")
		fmt.Println("1. Lihat Riwayat Interaksi")
		fmt.Println("2. Ubah Interaksi Tertentu")
		fmt.Println("3. Hapus Interaksi Tertentu")
		fmt.Println("4. Cari Saran Kesehatan Mental")
		fmt.Println("5. Urutkan Riwayat Interaksi")
		fmt.Println("6. Kembali ke Chat Utama")
		fmt.Println("7. Selesai Sesi & Keluar")
		fmt.Print("> ")
		fmt.Scanln(&pilihanMenuLanjutan)

		switch pilihanMenuLanjutan {
		case "1":
			cetakRiwayat(*A)
		case "2":
			ubahInteraksiTertentu(A)
		case "3":
			hapusInteraksiTertentu(A)
		case "4":
			cariSaranKesehatan()
		case "5":
			urutkanRiwayatInteraksi(A)
		case "6":
			return true 
		case "7":
			return false 
		default:
			fmt.Println("Pilihan tidak dikenali. Silakan coba lagi.")
		}
	}

}

func tambahInteraksi(A *tabInteraksi, suasanaHati string, pesan string) {
	if jumlahInteraksi < NMAX {
		A[jumlahInteraksi].Mood = suasanaHati
		A[jumlahInteraksi].Message = pesan
		jumlahInteraksi++
		fmt.Println("\n(Interaksi dicatat.)")
	} else {
		fmt.Println("Maaf, riwayat interaksi sudah penuh.")
	}
}

func dapatkanResponSuasanaHati(suasanaHati string) {
	switch suasanaHati {
	case "bahagia":
		fmt.Println("\nSenang sekali kamu merasa bahagia hari ini!")
		fmt.Println("Ingatlah momen-momen seperti ini ketika hari-hari terasa sulit.")
	case "sedih":
		fmt.Println("\nPerasaan sedih itu wajar dan manusiawi.")
		fmt.Println("Terkadang kita perlu merasakan kesedihan. Aku ada di sini untukmu.")
	case "cemas":
		fmt.Println("\nKecemasan bisa terasa seperti badai dalam pikiran.")
		fmt.Println("Coba tarik napas dalam-dalam beberapa kali. Kamu tidak sendirian.")
	case "marah":
		fmt.Println("\nMarah adalah emosi yang valid. Penting untuk menyalurkannya dengan sehat.")
		fmt.Println("Apa yang bisa kamu lakukan untuk meredakan perasaan ini sedikit?")
	case "lelah":
		fmt.Println("\nTubuh dan pikiranmu butuh istirahat. Jangan paksakan diri.")
		fmt.Println("Mungkin ini saatnya untuk beristirahat sejenak.")
	default:
		fmt.Println("\nTerima kasih sudah berbagi perasaanmu denganku.")
	}
}

func sarankanAktivitas(suasanaHati string) {
	if suasanaHati == "bahagia" {
		fmt.Println("Lanjutkan hal yang membuatmu bahagia!")
		return
	}

	var inputPengguna string
	fmt.Println("\nButuh saran aktivitas untuk membantumu merasa lebih baik? (ya/tidak)")
	fmt.Print("> ")
	fmt.Scanln(&inputPengguna)

	if inputPengguna == "ya" {
		fmt.Println("\nBerdasarkan perasaanmu, berikut beberapa aktivitas yang bisa membantu:")
		switch suasanaHati {
		case "sedih":
			fmt.Println("- Dengarkan musik yang menenangkan atau uplifting.")
			fmt.Println("- Tulis perasaanmu dalam jurnal, biarkan semua keluar.")
			fmt.Println("- Hubungi teman terdekat atau orang yang kamu percaya.")
		case "cemas":
			fmt.Println("- Coba latihan pernapasan ringan (misal: 4-7-8).")
			fmt.Println("- Lakukan peregangan atau jalan santai di alam terbuka.")
			fmt.Println("- Alihkan pikiran dengan aktivitas yang kamu nikmati (membaca, menggambar).")
		case "marah":
			fmt.Println("- Lakukan aktivitas fisik seperti lari atau olahraga untuk melepaskan energi.")
			fmt.Println("- Coba teknik grounding: sebutkan 5 hal yang kamu lihat, 4 yang disentuh, dst.")
			fmt.Println("- Tuliskan apa yang membuatmu marah, lalu robek kertasnya jika perlu.")
		case "lelah":
			fmt.Println("- Istirahat sejenak tanpa gangguan (jauhkan ponsel).")
			fmt.Println("- Minum air putih yang cukup dan pertimbangkan tidur lebih awal.")
			fmt.Println("- Dengarkan musik relaksasi atau meditasi singkat.")
		default:
			fmt.Println("- Ambil waktu sejenak untuk dirimu sendiri, lakukan sesuatu yang menenangkan.")
		}
	}
}

func tawarkanRelaksasi(suasanaHati string) {
	var inputPengguna string
	if suasanaHati == "bahagia" {
		return 
	}

	fmt.Println("\nMaukah kamu mencoba teknik relaksasi sederhana bersamaku? (ya/tidak)")
	fmt.Print("> ")
	fmt.Scanln(&inputPengguna)

	if inputPengguna == "ya" {
		fmt.Println("\nBaiklah. Sekarang, pejamkan matamu jika nyaman.")
		fmt.Println("Tarik napas dalam melalui hidung... hitung sampai empat...")
		fmt.Println("Tahan napasmu... hitung sampai tujuh...")
		fmt.Println("Dan hembuskan perlahan melalui mulut... hitung sampai delapan.")
		fmt.Println("Ulangi beberapa kali sampai kamu merasa lebih tenang. Aku tunggu di sini.")
	}
}

func cetakRiwayat(A tabInteraksi) {
	if jumlahInteraksi == 0 {
		fmt.Println("\nBelum ada riwayat interaksi.")
		return
	}

	fmt.Println("\n--- Riwayat Interaksi ---")
	var i int
	for i = 0; i < jumlahInteraksi; i++ {
		fmt.Printf("%d. Mood: %s\n   Pesan: %s\n", i+1, A[i].Mood, A[i].Message)
	}
	fmt.Println("-------------------------")
}

func ubahInteraksiTertentu(A *tabInteraksi) {
	if jumlahInteraksi == 0 {
		fmt.Println("\nBelum ada riwayat untuk diubah.")
		return
	}
	cetakRiwayat(*A) 
	var nomorUbah int
	var suasanaHatiBaru, pesanBaru string

	fmt.Print("Masukkan nomor interaksi yang ingin diubah: ")
	fmt.Scanln(&nomorUbah) 

	if nomorUbah < 1 || nomorUbah > jumlahInteraksi {
		fmt.Println("Nomor interaksi tidak valid.")
		return
	}
	var indeksUntukDiubah int = nomorUbah - 1 

	fmt.Printf("Mengubah interaksi ke-%d:\n", nomorUbah)
	fmt.Printf("Mood saat ini: %s. Masukkan mood baru (atau '-' untuk tidak mengubah): ", A[indeksUntukDiubah].Mood)
	fmt.Scanln(&suasanaHatiBaru)
	fmt.Printf("Pesan saat ini: %s. Masukkan pesan baru (atau '-' untuk tidak mengubah): ", A[indeksUntukDiubah].Message)
	fmt.Println("(Notes : Gunakan titik sebagai spasi. Contoh: Cerita.Baru)")
	fmt.Scanln(&pesanBaru)


	if suasanaHatiBaru != "-" && suasanaHatiBaru != "" {
		A[indeksUntukDiubah].Mood = suasanaHatiBaru
	}

	if pesanBaru != "-" && pesanBaru != "" {
		A[indeksUntukDiubah].Message = pesanBaru
	}

	fmt.Println("Interaksi berhasil diubah.")
}

func hapusInteraksiTertentu(A *tabInteraksi) {
	if jumlahInteraksi == 0 {
		fmt.Println("\nBelum ada riwayat untuk dihapus.")
		return
	}
	cetakRiwayat(*A) 
	var nomorHapus int

	fmt.Print("Masukkan nomor interaksi yang ingin dihapus: ")
	var inputStr string
	fmt.Scanln(&inputStr)

	var kesalahanKonversi bool = false
	nomorHapus = 0
	if len(inputStr) > 0 {
		var temp int = 0
		
		var mulai int = 0

		if inputStr[0] == '+' { mulai = 1 }

		for k := mulai; k < len(inputStr); k++ {
			if inputStr[k] >= '0' && inputStr[k] <= '9' {
				temp = temp*10 + int(inputStr[k]-'0')
			} else {
				kesalahanKonversi = true
				break
			}
		}
		if !kesalahanKonversi {
			nomorHapus = temp
		}
	} else {
		kesalahanKonversi = true
	}

	if kesalahanKonversi || nomorHapus < 1 || nomorHapus > jumlahInteraksi {
		fmt.Println("Nomor interaksi tidak valid.")
		return
	}

	var indeksUntukDihapus int = nomorHapus - 1 
	var i int
	for i = indeksUntukDihapus; i < jumlahInteraksi-1; i++ {
		(*A)[i] = (*A)[i+1]
	}
	jumlahInteraksi--
	fmt.Println("Interaksi berhasil dihapus.")
}

func inisialisasiDaftarSaran() {
	if jumlahSaran < NMAX {
		daftarSaran[jumlahSaran] = SaranKesehatan{"stres", "Cobalah teknik pernapasan 4-7-8: tarik napas 4 hitungan, tahan 7, hembuskan 8."}; jumlahSaran++
	}
	if jumlahSaran < NMAX {
		daftarSaran[jumlahSaran] = SaranKesehatan{"cemas", "Alihkan fokus pada hal yang kamu syukuri hari ini, tuliskan 3 hal."}; jumlahSaran++
	}
	if jumlahSaran < NMAX {
		daftarSaran[jumlahSaran] = SaranKesehatan{"sedih", "Ingatlah bahwa perasaan ini akan berlalu. Kamu kuat."}; jumlahSaran++
	}
	if jumlahSaran < NMAX {
		daftarSaran[jumlahSaran] = SaranKesehatan{"motivasi", "Setiap langkah kecil adalah kemajuan. Teruslah bergerak maju!"}; jumlahSaran++
	}
	if jumlahSaran < NMAX {
		daftarSaran[jumlahSaran] = SaranKesehatan{"tidur", "Ciptakan rutinitas tidur yang konsisten dan hindari kafein sebelum tidur."}; jumlahSaran++
	}
}

func cariSaranKesehatan() {
	if jumlahSaran == 0 {
		fmt.Println("Maaf, daftar saran belum tersedia.")
		return
	}
	var pilihanCari, kataKunci string
	fmt.Println("\n--- Cari Saran Kesehatan Mental ---")
	fmt.Println("Pilih metode pencarian:")
	fmt.Println("1. Pencarian Sekuensial")
	fmt.Println("2. Pencarian Biner (Data harus diurutkan berdasarkan Topik dulu)")
	fmt.Print("> ")
	fmt.Scanln(&pilihanCari)

	fmt.Print("Masukkan topik atau kata kunci saran yang ingin dicari: ")
	fmt.Scanln(&kataKunci)

	var ditemukan bool = false
	if pilihanCari == "1" {
		fmt.Println("\nHasil Pencarian Sekuensial untuk '", kataKunci, "':")
		var i int
		for i = 0; i < jumlahSaran; i++ {
			if stringMengandung(daftarSaran[i].Topik, kataKunci) || stringMengandung(daftarSaran[i].IsiSaran, kataKunci) {
				fmt.Printf("- Topik: %s\n  Saran: %s\n", daftarSaran[i].Topik, daftarSaran[i].IsiSaran)
				ditemukan = true
			}
		}
	} else if pilihanCari == "2" {

		urutkanDaftarSaranBerdasarkanTopik(&daftarSaran, jumlahSaran) 
		fmt.Println("\nHasil Pencarian Biner untuk Topik '", kataKunci, "':")

		var rendah int = 0
		var tinggi int = jumlahSaran - 1
		var tengah int
		for rendah <= tinggi {
			tengah = rendah + (tinggi-rendah)/2
			if daftarSaran[tengah].Topik == kataKunci {
				fmt.Printf("- Topik: %s\n  Saran: %s\n", daftarSaran[tengah].Topik, daftarSaran[tengah].IsiSaran)
				ditemukan = true
		
				var l int = tengah - 1
				for l >= 0 && daftarSaran[l].Topik == kataKunci {
					fmt.Printf("- Topik: %s\n  Saran: %s\n", daftarSaran[l].Topik, daftarSaran[l].IsiSaran)
					l--
				}

				var r int = tengah + 1
				for r < jumlahSaran && daftarSaran[r].Topik == kataKunci {
					fmt.Printf("- Topik: %s\n  Saran: %s\n", daftarSaran[r].Topik, daftarSaran[r].IsiSaran)
					r++
				}
				break 
			} else if daftarSaran[tengah].Topik < kataKunci {
				rendah = tengah + 1
			} else {
				tinggi = tengah - 1
			}
		}
	} else {
		fmt.Println("Pilihan metode pencarian tidak valid.")
		return
	}

	if !ditemukan {
		fmt.Println("Saran untuk '", kataKunci, "' tidak ditemukan.")
	}
}

func stringMengandung(teksUtama, teksCari string) bool {
	var n int = len(teksUtama)
	var m int = len(teksCari)
	if m == 0 { return true } 
	if m > n { return false }
	var i, j int
	for i = 0; i <= n-m; i++ {
		var cocok bool = true
		for j = 0; j < m; j++ {
			if teksUtama[i+j] != teksCari[j] {
				cocok = false
				break
			}
		}
		if cocok {
			return true
		}
	}
	return false
}

func urutkanDaftarSaranBerdasarkanTopik(A *tabSaran, n int) {
	var i, j int
	var kunci SaranKesehatan
	for i = 1; i < n; i++ {
		kunci = A[i]
		j = i - 1
	
		for j >= 0 && A[j].Topik > kunci.Topik {
			A[j+1] = A[j]
			j--
		}
		A[j+1] = kunci 
	}
}

func urutkanRiwayatInteraksi(A *tabInteraksi) {
	if jumlahInteraksi == 0 {
		fmt.Println("Tidak ada riwayat untuk diurutkan.")
		return
	}
	var pilihanUrut, kriteriaUrut string
	fmt.Println("\n--- Urutkan Riwayat Interaksi ---")
	fmt.Println("Pilih metode pengurutan:")
	fmt.Println("1. Urut Seleksi (Selection Sort)")
	fmt.Println("2. Urut Sisip (Insertion Sort)")
	fmt.Print("> ")
	fmt.Scanln(&pilihanUrut)

	fmt.Println("Urutkan berdasarkan:")
	fmt.Println("1. Mood (A-Z)") 
	fmt.Println("2. Tingkat Urgensi (Tinggi ke Rendah)") 
	fmt.Print("> ")
	fmt.Scanln(&kriteriaUrut)

	var menaik bool = true
	if kriteriaUrut == "2" {
		menaik = false
	}

	if pilihanUrut == "1" {
		urutSeleksiInteraksi(A, jumlahInteraksi, kriteriaUrut, menaik)
	} else if pilihanUrut == "2" {
		urutSisipInteraksi(A, jumlahInteraksi, kriteriaUrut, menaik)
	} else {
		fmt.Println("Pilihan metode tidak valid.")
		return
	}
	fmt.Println("\nRiwayat setelah diurutkan:")
	cetakRiwayat(*A)
}

func dapatkanSkorUrgensi(suasanaHati string) int {
	switch suasanaHati {
	case "marah", "cemas": return 3 
	case "sedih": return 2          
	case "lelah": return 1          
	case "bahagia": return 0       
	default: return -1              
	}
}

func urutSeleksiInteraksi(A *tabInteraksi, n int, kriteria string, menaik bool) {
	var i, j, indeksEkstremum int
	var temp Interaction
	for i = 0; i < n-1; i++ {
		indeksEkstremum = i
		for j = i + 1; j < n; j++ {
			var kondisiBanding bool = false
			if kriteria == "1" { 
				if menaik { 
					kondisiBanding = A[j].Mood < A[indeksEkstremum].Mood
				} else { 
					kondisiBanding = A[j].Mood > A[indeksEkstremum].Mood
				}
			} else if kriteria == "2" {
				skorJ := dapatkanSkorUrgensi(A[j].Mood)
				skorEkstremum := dapatkanSkorUrgensi(A[indeksEkstremum].Mood)
				if menaik { 
					kondisiBanding = skorJ < skorEkstremum
				} else { 
					kondisiBanding = skorJ > skorEkstremum
				}
			}

			if kondisiBanding {
				indeksEkstremum = j
			}
		}
		if indeksEkstremum != i { 
			temp = A[i]
			A[i] = A[indeksEkstremum]
			A[indeksEkstremum] = temp
		}
	}
}

func urutSisipInteraksi(A *tabInteraksi, n int, kriteria string, menaik bool) {
	var i, j int
	var kunci Interaction
	for i = 1; i < n; i++ {
		kunci = A[i]
		j = i - 1

		var kondisiGeser bool = false
		if j >= 0 {
			if kriteria == "1" { 
				if menaik { kondisiGeser = A[j].Mood > kunci.Mood
				} else { kondisiGeser = A[j].Mood < kunci.Mood }
			} else if kriteria == "2" { // Urgensi
				skorJ := dapatkanSkorUrgensi(A[j].Mood)
				skorKunci := dapatkanSkorUrgensi(kunci.Mood)
				if menaik { kondisiGeser = skorJ > skorKunci
				} else { kondisiGeser = skorJ < skorKunci }
			}
		}
		
		
		for j >= 0 && kondisiGeser {
			A[j+1] = A[j]
			j--
			kondisiGeser = false 
			if j >= 0 {
				if kriteria == "1" {
					if menaik { kondisiGeser = A[j].Mood > kunci.Mood 
					} else { kondisiGeser = A[j].Mood < kunci.Mood }
				} else if kriteria == "2" {
					skorJ := dapatkanSkorUrgensi(A[j].Mood)
					skorKunci := dapatkanSkorUrgensi(kunci.Mood)
					if menaik { kondisiGeser = skorJ > skorKunci
					} else { kondisiGeser = skorJ < skorKunci }
				}
			}
		}
		A[j+1] = kunci 
	}
}
