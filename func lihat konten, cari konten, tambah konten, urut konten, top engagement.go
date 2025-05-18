func tambahKonten() {
	var p, t, i, h, c string
	var e int

	fmt.Println("\n=== Tambah Ide Konten ===")
	fmt.Println("Pilih Platform (Instagram / Tiktok / Facebook / Lainnya): ")
	fmt.Scan(&p)

	if p == "Lainnya" {
		fmt.Print("Masukkan platform custom: ")
		fmt.Scan(&p)
	}

	fmt.Print("Tanggal (YYYY-MM-DD): ")
	fmt.Scan(&t)
	fmt.Print("Ide Konten: ")
	fmt.Scan(&i)
	fmt.Print("Hashtag: ")
	fmt.Scan(&h)
	fmt.Print("Caption: ")
	fmt.Scan(&c)
	fmt.Print("Engagement: ")
	fmt.Scan(&e)

	KL[jumlahKonten] = Konten{p, t, e, i, h, c}
	jumlahKonten++

	fmt.Println("Konten berhasil ditambahkan.")
}

func lihatKonten() {
	if jumlahKonten == 0 {
		fmt.Println("\nBelum ada konten.")
		return
	}

	fmt.Println("\n=== Daftar Ide Konten ===")
	var i int
	for i = 0; i < jumlahKonten; i++ {
		fmt.Printf("%d. %s | %s | %d | %s | %s | %s\n",
			i+1, KL[i].Platform, KL[i].Tanggal, KL[i].Engagement, KL[i].Ide_Konten, KL[i].Hashtag, KL[i].Caption)
	}
}

func cariKonten() {
	if jumlahKonten == 0 {
		fmt.Println("\nBelum ada konten.")
		return
	}

	var keyword string
	var ketemu bool
	ketemu = false

	fmt.Print("\nMasukkan keyword ide konten: ")
	fmt.Scan(&keyword)

	var i int
	for i = 0; i < jumlahKonten; i++ {
		if KL[i].Ide_Konten == keyword {
			fmt.Printf("Ditemukan di data ke-%d: %s | %s | %d | %s | %s | %s\n",
				i+1, KL[i].Platform, KL[i].Tanggal, KL[i].Engagement, KL[i].Ide_Konten, KL[i].Hashtag, KL[i].Caption)
			ketemu = true
		}
	}

	if !ketemu {
		fmt.Println("Ide konten tidak ditemukan.")
	}
}

func urutkanKonten() {
	var i, j, maxIdx int
	var temp Konten

	for i = 0; i < jumlahKonten-1; i++ {
		maxIdx = i
		for j = i + 1; j < jumlahKonten; j++ {
			if KL[j].Engagement > KL[maxIdx].Engagement {
				maxIdx = j
			}
		}
		temp = KL[i]
		KL[i] = KL[maxIdx]
		KL[maxIdx] = temp
	}

	fmt.Println("Konten berhasil diurutkan berdasarkan engagement (tertinggi ke terendah).")
}

func topEngagement() {
	if jumlahKonten == 0 {
		fmt.Println("\nBelum ada konten.")
		return
	}

	var max Konten
	var i int
	max = KL[0]

	for i = 1; i < jumlahKonten; i++ {
		if KL[i].Engagement > max.Engagement {
			max = KL[i]
		}
	}

	fmt.Println("\nKonten dengan engagement tertinggi:")
	fmt.Printf("%s | %s | %d | %s | %s | %s\n",
		max.Platform, max.Tanggal, max.Engagement, max.Ide_Konten, max.Hashtag, max.Caption)
}