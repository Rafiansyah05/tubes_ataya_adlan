package main

import "fmt"

const NMAX int = 100 

type Interaction struct {
	Mood    string
	Message string
}

type SaranKesehatan struct {
	Topik   string
	IsiSaran string
}

type tabInteraksi [NMAX]Interaction 
var jumlahInteraksi int = 0

type tabSaran [NMAX]SaranKesehatan
var daftarSaran tabSaran
var jumlahSaran int = 0

func main() {
	var name string
	var tbInter tabInteraksi 

	initDaftarSaran()
	fmt.Println("\n========== CHATBOT KESEHATAN MENTAL ==========")
	fmt.Println("Halo! Saya adalah teman bicara untuk kesehatan mental.")
	fmt.Println("\nSebelum kita mulai, boleh saya tahu namamu?")
	fmt.Print("> ")
	fmt.Scanln(&name)

	fmt.Printf("\nHalo, %s! 😊\n", name)

	var sessionActive bool = true
	for sessionActive {
		var pilihanMenuUtama string
		var userMessage string
		var mood string = "" 

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
			sessionActive = handleOpsiLanjutan(&tbInter) 
			if !sessionActive { 
				break 
			}
			continue
		}

		switch pilihanMenuUtama {
		case "1":
			mood = "bahagia"
			fmt.Println("Wah, senang mendengarnya! Ada cerita menyenangkan apa hari ini?")
			fmt.Println("(Notes : Gunakan titik sebagai spasi. Contoh: Aku.habis.makan.enak)")
			fmt.Print("> ")
			fmt.Scanln(&userMessage)
		case "2":
			mood = "sedih"
			fmt.Println("Aku di sini untuk mendengarmu. Mau ceritakan apa yang membuatmu sedih?")
			fmt.Println("(Notes : Gunakan titik sebagai spasi. Contoh: Aku.habis.diputusin.pacar)")
			fmt.Print("> ")
			fmt.Scanln(&userMessage)
		case "3":
			mood = "cemas"
			fmt.Println("Kecemasan bisa terasa berat. Apa yang sedang mengganggu pikiranmu?")
			fmt.Println("(Notes : Gunakan titik sebagai spasi. Contoh: Besok.aku.akan.dikaderisasi)")
			fmt.Print("> ")
			fmt.Scanln(&userMessage)
		case "4":
			mood = "marah"
			fmt.Println("Marah adalah perasaan yang wajar. Apa yang memicunya?")
			fmt.Println("(Notes : Gunakan titik sebagai spasi. Contoh: Dia.menyebalkan)")
			fmt.Print("> ")
			fmt.Scanln(&userMessage)
		case "5":
			mood = "lelah"
			fmt.Println("Istirahat itu penting. Apakah kamu merasa terbebani akhir-akhir ini?")
			fmt.Println("(Notes : Gunakan titik sebagai spasi. Contoh: Baru.pulang.jogging)")
			fmt.Print("> ")
			fmt.Scanln(&userMessage)
		default:
			fmt.Println("Maaf, aku tidak mengerti pilihanmu.")
			continue
		}

		if mood != "" && userMessage != "" {
			addInteraction(&tbInter, mood, userMessage)
			getMoodResponse(mood)
			suggestActivity(mood)
			offerRelaxation(mood)
		} else if mood != "" && userMessage == "" {
			fmt.Println("Kamu belum menceritakan apa-apa. Ceritakanlah!")
		}
	}

	fmt.Println("\nTerima kasih sudah berbicara denganku hari ini, ", name, ".")
	fmt.Println("Ingatlah: Perasaanmu valid dan penting.")
	fmt.Println("Sampai jumpa lagi! ❤")
}

func handleOpsiLanjutan(A *tabInteraksi) bool {
	var pilihanMenuLanjutan string
	var continueSession bool = true

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
			printHistory(*A)
		case "2":
			ubahInteraksiSpesifik(A)
		case "3":
			hapusInteraksiSpesifik(A)
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
	return continueSession 
}

func addInteraction(A *tabInteraksi, mood string, message string) {
	if jumlahInteraksi < NMAX {
		A[jumlahInteraksi].Mood = mood
		A[jumlahInteraksi].Message = message
		jumlahInteraksi++
		fmt.Println("\n(Interaksi dicatat.)")
	} else {
		fmt.Println("Maaf, riwayat interaksi sudah penuh.")
	}
}

func getMoodResponse(mood string) {
	switch mood {
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

func suggestActivity(mood string) {
	if mood == "bahagia" {
		fmt.Println("Lanjutkan hal yang membuatmu bahagia!")
		return
	}

	var input string
	fmt.Println("\nButuh saran aktivitas untuk membantumu merasa lebih baik? (ya/tidak)")
	fmt.Print("> ")
	fmt.Scanln(&input)

	if input == "ya" {
		fmt.Println("\nBerdasarkan perasaanmu, berikut beberapa aktivitas yang bisa membantu:")
		switch mood {
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

func offerRelaxation(mood string) {
	var input string
	if mood == "bahagia" {
		return
	}

	fmt.Println("\nMaukah kamu mencoba teknik relaksasi sederhana bersamaku? (ya/tidak)")
	fmt.Print("> ")
	fmt.Scanln(&input)

	if input == "ya" {
		fmt.Println("\nBaiklah. Sekarang, pejamkan matamu jika nyaman.")
		fmt.Println("Tarik napas dalam melalui hidung... hitung sampai empat...")
		fmt.Println("Tahan napasmu... hitung sampai tujuh...")
		fmt.Println("Dan hembuskan perlahan melalui mulut... hitung sampai delapan.")
		fmt.Println("Ulangi beberapa kali sampai kamu merasa lebih tenang. Aku tunggu di sini.")
	}
}

func printHistory(A tabInteraksi) {
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

func ubahInteraksiSpesifik(A *tabInteraksi) {
	if jumlahInteraksi == 0 {
		fmt.Println("\nBelum ada riwayat untuk diubah.")
		return
	}
	printHistory(*A)
	var nomorUbah int
	var newMood, newMessage string

	fmt.Print("Masukkan nomor interaksi yang ingin diubah: ")
	fmt.Scanln(&nomorUbah) 

	if nomorUbah < 1 || nomorUbah > jumlahInteraksi {
		fmt.Println("Nomor interaksi tidak valid.")
		return
	}
	var indexToChange int = nomorUbah - 1

	fmt.Printf("Mengubah interaksi ke-%d:\n", nomorUbah)
	fmt.Printf("Mood saat ini: %s. Masukkan mood baru (atau '-' untuk tidak mengubah): ", A[indexToChange].Mood)
	fmt.Scanln(&newMood)
	fmt.Printf("Pesan saat ini: %s. Masukkan pesan baru (atau '-' untuk tidak mengubah): ", A[indexToChange].Message)
	fmt.Println("(Notes : Gunakan titik sebagai spasi. Contoh: Cerita.Baru)")
	fmt.Scanln(&newMessage)

	if newMood != "-" && newMood != "" {
		A[indexToChange].Mood = newMood
	}
	if newMessage != "-" && newMessage != "" {
		A[indexToChange].Message = newMessage
	}
	fmt.Println("Interaksi berhasil diubah.")
}

func hapusInteraksiSpesifik(A *tabInteraksi) {
	if jumlahInteraksi == 0 {
		fmt.Println("\nBelum ada riwayat untuk dihapus.")
		return
	}
	printHistory(*A) 
	var nomorHapus int

	fmt.Print("Masukkan nomor interaksi yang ingin dihapus: ")
	var inputStr string
	fmt.Scanln(&inputStr)
	
	var errConv bool = false
	nomorHapus = 0
	if len(inputStr) > 0 {
		var temp int = 0
		var sign int = 1
		var start int = 0

		if inputStr[0] == '-' {
			errConv = true
		} else if inputStr[0] == '+' {
			start = 1
		}
		for k := start; k < len(inputStr); k++ {
			if inputStr[k] >= '0' && inputStr[k] <= '9' {
				temp = temp*10 + int(inputStr[k]-'0')
			} else {
				errConv = true
				break
			}
		}
		if !errConv {
			nomorHapus = temp * sign
		}
		
	} else {
		errConv = true
	}


	if errConv || nomorHapus < 1 || nomorHapus > jumlahInteraksi {
		fmt.Println("Nomor interaksi tidak valid.")
		return
	}

	var indexToDelete int = nomorHapus - 1
	
	var i int
	for i = indexToDelete; i < jumlahInteraksi-1; i++ {
		(*A)[i] = (*A)[i+1]
	}

	jumlahInteraksi--
	fmt.Println("Interaksi berhasil dihapus.")
}



func initDaftarSaran() {

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
	fmt.Println("1. Sequential Search")
	fmt.Println("2. Binary Search (Data harus diurutkan berdasarkan Topik dulu)")
	fmt.Print("> ")
	fmt.Scanln(&pilihanCari)

	fmt.Print("Masukkan topik atau kata kunci saran yang ingin dicari: ")
	fmt.Scanln(&kataKunci)

	var ditemukan bool = false
	if pilihanCari == "1" {
		fmt.Println("\nHasil Sequential Search untuk '", kataKunci, "':")
		var i int
		for i = 0; i < jumlahSaran; i++ {

			if stringContains(daftarSaran[i].Topik, kataKunci) || stringContains(daftarSaran[i].IsiSaran, kataKunci) {
				fmt.Printf("- Topik: %s\n  Saran: %s\n", daftarSaran[i].Topik, daftarSaran[i].IsiSaran)
				ditemukan = true
			}
		}
	} else if pilihanCari == "2" {
		urutkanDaftarSaranByTopik(&daftarSaran, jumlahSaran) 
		fmt.Println("\nHasil Binary Search untuk Topik '", kataKunci, "':")
		var low int = 0
		var high int = jumlahSaran - 1
		var mid int
		for low <= high {
			mid = low + (high-low)/2
			if daftarSaran[mid].Topik == kataKunci {
				fmt.Printf("- Topik: %s\n  Saran: %s\n", daftarSaran[mid].Topik, daftarSaran[mid].IsiSaran)
				ditemukan = true

				var l int = mid -1
				for l >=0 && daftarSaran[l].Topik == kataKunci {
					fmt.Printf("- Topik: %s\n  Saran: %s\n", daftarSaran[l].Topik, daftarSaran[l].IsiSaran)
					l--
				}
		
				var r int = mid + 1
				for r < jumlahSaran && daftarSaran[r].Topik == kataKunci {
					fmt.Printf("- Topik: %s\n  Saran: %s\n", daftarSaran[r].Topik, daftarSaran[r].IsiSaran)
					r++
				}
				break
			} else if daftarSaran[mid].Topik < kataKunci {
				low = mid + 1
			} else {
				high = mid - 1
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

func stringContains(haystack, needle string) bool {
	var n int = len(haystack)
	var m int = len(needle)
	if m == 0 { return true }
	if m > n { return false }
	var i, j int
	for i = 0; i <= n-m; i++ {
		var match bool = true
		for j = 0; j < m; j++ {
			if haystack[i+j] != needle[j] {
				match = false
				break
			}
		}
		if match {
			return true
		}
	}
	return false
}

func urutkanDaftarSaranByTopik(A *tabSaran, n int) {
	var i, j int
	var key SaranKesehatan
	for i = 1; i < n; i++ {
		key = A[i]
		j = i - 1
		for j >= 0 && A[j].Topik > key.Topik {
			A[j+1] = A[j]
			j--
		}
		A[j+1] = key
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
	fmt.Println("1. Selection Sort")
	fmt.Println("2. Insertion Sort")
	fmt.Print("> ")
	fmt.Scanln(&pilihanUrut)

	fmt.Println("Urutkan berdasarkan:")
	fmt.Println("1. Mood (A-Z)")
	fmt.Println("2. Tingkat Urgensi (Tinggi ke Rendah)")

	fmt.Print("> ")
	fmt.Scanln(&kriteriaUrut)

	var ascending bool = true 

	if kriteriaUrut == "2" { 
		ascending = false
	}


	if pilihanUrut == "1" { 
		selectionSortInteraksi(A, jumlahInteraksi, kriteriaUrut, ascending)
	} else if pilihanUrut == "2" { 
		insertionSortInteraksi(A, jumlahInteraksi, kriteriaUrut, ascending)
	} else {
		fmt.Println("Pilihan metode tidak valid.")
		return
	}
	fmt.Println("\nRiwayat setelah diurutkan:")
	printHistory(*A)
}

func getUrgensiScore(mood string) int {
	switch mood {
	case "marah", "cemas": return 3
	case "sedih": return 2      
	case "lelah": return 1         
	case "bahagia": return 0        
	default: return -1             
	}
}

func selectionSortInteraksi(A *tabInteraksi, n int, kriteria string, ascending bool) {
	var i, j, extremumIdx int
	var temp Interaction
	for i = 0; i < n-1; i++ {
		extremumIdx = i
		for j = i + 1; j < n; j++ {
			var A_j_lebihKecilAtauBesarDariExtremum bool = false
			if kriteria == "1" { 
				if ascending { 
					A_j_lebihKecilAtauBesarDariExtremum = A[j].Mood < A[extremumIdx].Mood
				} else { 
					A_j_lebihKecilAtauBesarDariExtremum = A[j].Mood > A[extremumIdx].Mood
				}
			} else if kriteria == "2" { 
				scoreJ := getUrgensiScore(A[j].Mood)
				scoreExtremum := getUrgensiScore(A[extremumIdx].Mood)
				if ascending {
					A_j_lebihKecilAtauBesarDariExtremum = scoreJ < scoreExtremum
				} else {
					A_j_lebihKecilAtauBesarDariExtremum = scoreJ > scoreExtremum
				}
			}

			if A_j_lebihKecilAtauBesarDariExtremum {
				extremumIdx = j
			}
		}
		if extremumIdx != i {
			temp = A[i]
			A[i] = A[extremumIdx]
			A[extremumIdx] = temp
		}
	}
}

func insertionSortInteraksi(A *tabInteraksi, n int, kriteria string, ascending bool) {
	var i, j int
	var key Interaction
	for i = 1; i < n; i++ {
		key = A[i]
		j = i - 1
		
		var kondisiGeser bool = false
		if j >= 0 {
			if kriteria == "1" {
				if ascending {
					kondisiGeser = A[j].Mood > key.Mood
				} else {
					kondisiGeser = A[j].Mood < key.Mood
				}
			} else if kriteria == "2" {
				scoreJ := getUrgensiScore(A[j].Mood)
				scoreKey := getUrgensiScore(key.Mood)
				if ascending {
					kondisiGeser = scoreJ > scoreKey
				} else {
					kondisiGeser = scoreJ < scoreKey
				}
			}
		}

		for j >= 0 && kondisiGeser {
			A[j+1] = A[j]
			j--

			kondisiGeser = false 

			if j >= 0 {
				if kriteria == "1" {
					if ascending { kondisiGeser = A[j].Mood > key.Mood 
					} else { kondisiGeser = A[j].Mood < key.Mood }
				} else if kriteria == "2" {
					scoreJ := getUrgensiScore(A[j].Mood)
					scoreKey := getUrgensiScore(key.Mood)
					if ascending { kondisiGeser = scoreJ > scoreKey
					} else { kondisiGeser = scoreJ < scoreKey }
				}
			}
		}
		A[j+1] = key
	}
}