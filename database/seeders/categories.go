package seeders

import (
	"fmt"
	"real-time-application/database"
	"real-time-application/services"
)

func CategorySeeder() {
	categories := []services.Category{
		{
			Name:        "Jersey Tim",
			Description: "Koleksi jersey berkualitas tinggi untuk menunjang performa di lapangan dan menonjolkan gaya tim Anda.",
			Slug:        "jersey",
		},
		{
			Name:        "Sepatu Futsal",
			Description: "Pilihan sepatu futsal yang dirancang untuk kenyamanan, daya cengkeram, dan performa maksimal di setiap pertandingan.",
			Slug:        "sepatu-futsal",
		},
		{
			Name:        "Bola Futsal",
			Description: "Bola futsal dengan bahan premium yang sesuai standar resmi, ideal untuk latihan maupun kompetisi.",
			Slug:        "bola-futsal",
		},
		{
			Name:        "Aksesoris Futsal",
			Description: "Beragam aksesoris futsal seperti sarung tangan kiper, pelindung kaki, dan tas olahraga yang mendukung kebutuhan pemain.",
			Slug:        "aksesoris",
		},
		{
			Name:        "Pakaian Kasual",
			Description: "Koleksi pakaian santai untuk pemain dan penggemar futsal yang ingin tampil sporty di luar lapangan.",
			Slug:        "kasual",
		},
		{
			Name:        "Perlengkapan Latihan",
			Description: "Alat bantu latihan seperti cone, agility ladder, dan peluit untuk meningkatkan skill dan kebugaran tim.",
			Slug:        "perlengkapan-latihan",
		},
		{
			Name:        "Peralatan Wasit",
			Description: "Produk penting untuk wasit seperti kartu pertandingan, peluit berkualitas, dan stopwatch.",
			Slug:        "peralatan-wasit",
		},
		{
			Name:        "Merchandise Futsal",
			Description: "Produk merchandise eksklusif untuk pecinta futsal, termasuk kaos, topi, dan gantungan kunci bertema futsal.",
			Slug:        "merchendise",
		},
		{
			Name:        "Minuman & Energi",
			Description: "Pilihan minuman isotonic dan makanan ringan untuk menjaga energi pemain selama pertandingan.",
			Slug:        "foods",
		},
		{
			Name:        "Dekorasi Lapangan",
			Description: "Produk untuk mempercantik dan melengkapi lapangan futsal, seperti gawang portable dan penanda garis.",
			Slug:        "decorations",
		},
	}

	for _, category := range categories {
		if err := database.DB.Create(&category).Error; err != nil {
			fmt.Println("Create Categories")
		}
	}

	fmt.Println("Success Create Categories")
}
