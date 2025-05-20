package main

import "fmt"

type Konten struct {
	Platform   string
	Tanggal    string
	Engagement int
	Ide_Konten string
	Hashtag string
	Caption string
}

var KL [100]Konten
var Platform [100]string

func main() {
    var pilih int
   
    KL[0].Platform = "Instagram"
    KL[1].Platform = "Tiktok"
    KL[2].Platform = "Facebook"
    
    

	for pilih != 6 {
		fmt.Println("\n=== MENU UTAMA ===")
		fmt.Println("1. Tambah Ide Konten")
		fmt.Println("2. Lihat Semua Konten")
		fmt.Println("3. Cari Konten (Sequential)")
		fmt.Println("4. Urutkan Konten")
		fmt.Println("5. Engagement Tertinggi")
		fmt.Println("6. Keluar")
		fmt.Print("Pilih menu: ")

		var pilih int
		fmt.Scan(&pilih)

		switch pilih {
		case 1:
			tambahKonten()
		case 2:
			lihatKonten()
		case 3:
			cariKonten()
		case 4:
			urutkanKonten()
		case 5:
			topEngagement()
		case 6:
			fmt.Println("Terima kasih.")
		default:
			fmt.Println("Pilihan tidak valid.")
		}
	}
}