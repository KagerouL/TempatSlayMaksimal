package main

import (
	"fmt"
)

const (
	NMAX            int = 100
	MAX_VARIANT     int = 5
	MAX_REVIEW      int = 10
	MAX_ITEM        int = 10
)

type Brand struct {
	Name    string
	Country string
}

type Rating struct {
	Score       float64
	TotalReview int
}

type Review struct {
	Username string
	Comment  string
	Value    int
}

type ProductDetail struct {
	Description string
	SkinType    string
	ExpiredYear int
}

type Variant struct {
	Color string
	Size  string
	Stock int
}

type Product struct {
	ID       string
	Name     string
	Category string
	Price    int
	Sold     int

	BrandInfo  Brand
	RateInfo   Rating
	DetailInfo ProductDetail

	Variants     [MAX_VARIANT]Variant
	VariantCount int

	Reviews     [MAX_REVIEW]Review
	ReviewCount int
}

type ProductArray [NMAX]Product

var productsArr ProductArray
var countData int
var historyCountData int


func init() {
	countData = 30
	historyCountData = 30

	productsArr = ProductArray{
		{
			ID: "TSM001", Name: "Lip Cream Matte", Category: "makeup", Price: 89000, Sold: 234,
			BrandInfo:  Brand{"Wardah", "Indonesia"},
			RateInfo:   Rating{4.7, 128},
			DetailInfo: ProductDetail{"Lip cream matte finish dengan formula ringan dan tahan lama", "all", 2026},
			Variants: [MAX_VARIANT]Variant{
				{"Red Velvet", "5ml", 45},
				{"Rose Pink", "5ml", 32},
				{"Mauve", "5ml", 28},
			},
			VariantCount: 3,
			Reviews: [MAX_REVIEW]Review{
				{"@beautygirl", "Warnanya pigmented banget!", 5},
				{"@makeupaddict", "Matte tapi tidak kering", 4},
			},
			ReviewCount: 2,
		},
		// 2
		{
			ID: "TSM002", Name: "BB Cream SPF 30", Category: "makeup", Price: 75000, Sold: 412,
			BrandInfo:  Brand{"Wardah", "Indonesia"},
			RateInfo:   Rating{4.5, 210},
			DetailInfo: ProductDetail{"BB cream dengan perlindungan SPF 30 untuk kulit cerah merata", "normal", 2027},
			Variants: [MAX_VARIANT]Variant{
				{"Natural Beige", "30ml", 60},
				{"Ivory", "30ml", 40},
			},
			VariantCount: 2,
			Reviews: [MAX_REVIEW]Review{
				{"@skincareenthusiast", "Coverage-nya bagus dan ringan!", 5},
				{"@dailymakeup", "Cocok untuk kulit normal", 4},
				{"@beautyreview", "SPF-nya terasa", 5},
			},
			ReviewCount: 3,
		},
		// 3
		{
			ID: "TSM003", Name: "Cushion Foundation", Category: "makeup", Price: 135000, Sold: 189,
			BrandInfo:  Brand{"Emina", "Indonesia"},
			RateInfo:   Rating{4.3, 95},
			DetailInfo: ProductDetail{"Cushion foundation dengan hasil akhir glowing dan tahan hingga 12 jam", "dry", 2026},
			Variants: [MAX_VARIANT]Variant{
				{"Porcelain", "15g", 25},
				{"Beige", "15g", 30},
				{"Sand", "15g", 20},
			},
			VariantCount: 3,
			Reviews: [MAX_REVIEW]Review{
				{"@glowingskin", "Hasilnya dewy banget!", 5},
				{"@makeupjunkie", "Tahan lama seharian", 4},
			},
			ReviewCount: 2,
		},
		// 4
		{
			ID: "TSM004", Name: "Eyeshadow Palette", Category: "makeup", Price: 165000, Sold: 302,
			BrandInfo:  Brand{"Implora", "Indonesia"},
			RateInfo:   Rating{4.6, 175},
			DetailInfo: ProductDetail{"Palette eyeshadow 12 warna dengan pigmentasi tinggi dan blendable", "all", 2027},
			Variants: [MAX_VARIANT]Variant{
				{"Warm Nude", "12g", 50},
				{"Cool Smoke", "12g", 35},
			},
			VariantCount: 2,
			Reviews: [MAX_REVIEW]Review{
				{"@eyemakeup", "Pigmentasi sangat bagus!", 5},
				{"@beautyblend", "Mudah di-blend", 5},
				{"@makeupid", "Worth the price!", 4},
			},
			ReviewCount: 3,
		},
		// 5
		{
			ID: "TSM005", Name: "Loose Powder Translucent", Category: "makeup", Price: 55000, Sold: 520,
			BrandInfo:  Brand{"Viva", "Indonesia"},
			RateInfo:   Rating{4.4, 300},
			DetailInfo: ProductDetail{"Bedak tabur translucent untuk mengunci makeup sepanjang hari", "oily", 2026},
			Variants: [MAX_VARIANT]Variant{
				{"Translucent", "20g", 80},
			},
			VariantCount: 1,
			Reviews: [MAX_REVIEW]Review{
				{"@oilyskincare", "Ampuh banget nahan minyak!", 5},
				{"@powderlover", "Ringan di wajah", 4},
			},
			ReviewCount: 2,
		},
		// 6
		{
			ID: "TSM006", Name: "Facial Wash Salicylic Acid", Category: "skincare", Price: 48000, Sold: 675,
			BrandInfo:  Brand{"Somethinc", "Indonesia"},
			RateInfo:   Rating{4.8, 420},
			DetailInfo: ProductDetail{"Sabun cuci muka dengan salicylic acid untuk kulit berjerawat", "oily", 2027},
			Variants: [MAX_VARIANT]Variant{
				{"Original", "100ml", 100},
			},
			VariantCount: 1,
			Reviews: [MAX_REVIEW]Review{
				{"@acneskin", "Jerawat cepat kempes!", 5},
				{"@skincaredaily", "Recommended banget!", 5},
				{"@beautycare", "Kulitku jadi lebih bersih", 4},
			},
			ReviewCount: 3,
		},
		// 7
		{
			ID: "TSM007", Name: "Moisturizer Hyaluronic Acid", Category: "skincare", Price: 120000, Sold: 388,
			BrandInfo:  Brand{"Somethinc", "Indonesia"},
			RateInfo:   Rating{4.9, 250},
			DetailInfo: ProductDetail{"Pelembap dengan hyaluronic acid untuk kulit terhidrasi maksimal", "dry", 2027},
			Variants: [MAX_VARIANT]Variant{
				{"Original", "50ml", 70},
			},
			VariantCount: 1,
			Reviews: [MAX_REVIEW]Review{
				{"@dryskin101", "Kulitku jadi lembap banget!", 5},
				{"@hydratedskin", "Teksturnya ringan dan cepat meresap", 5},
			},
			ReviewCount: 2,
		},
		// 8
		{
			ID: "TSM008", Name: "Sunscreen SPF 50 PA++++", Category: "skincare", Price: 95000, Sold: 810,
			BrandInfo:  Brand{"Azarine", "Indonesia"},
			RateInfo:   Rating{4.8, 600},
			DetailInfo: ProductDetail{"Sunscreen SPF 50 PA++++ dengan formula ringan dan tidak lengket", "all", 2026},
			Variants: [MAX_VARIANT]Variant{
				{"Original", "50ml", 120},
			},
			VariantCount: 1,
			Reviews: [MAX_REVIEW]Review{
				{"@sunscreenlovers", "Tidak putih atau lengket sama sekali!", 5},
				{"@spfeveryday", "Jadi andalan sehari-hari", 5},
				{"@skinprotect", "No white cast!", 5},
			},
			ReviewCount: 3,
		},
		// 9
		{
			ID: "TSM009", Name: "Serum Niacinamide 10%", Category: "skincare", Price: 85000, Sold: 543,
			BrandInfo:  Brand{"Skintific", "Indonesia"},
			RateInfo:   Rating{4.7, 380},
			DetailInfo: ProductDetail{"Serum niacinamide 10% untuk mencerahkan dan mengecilkan pori", "all", 2027},
			Variants: [MAX_VARIANT]Variant{
				{"Original", "20ml", 90},
			},
			VariantCount: 1,
			Reviews: [MAX_REVIEW]Review{
				{"@brightenskin", "Pori-pori terlihat mengecil!", 5},
				{"@niacinamideuser", "Wajah jadi lebih cerah dalam seminggu", 4},
			},
			ReviewCount: 2,
		},
		// 10
		{
			ID: "TSM010", Name: "Toner Centella Asiatica", Category: "skincare", Price: 68000, Sold: 297,
			BrandInfo:  Brand{"Avoskin", "Indonesia"},
			RateInfo:   Rating{4.5, 180},
			DetailInfo: ProductDetail{"Toner dengan centella asiatica untuk menenangkan kulit sensitif", "sensitive", 2027},
			Variants: [MAX_VARIANT]Variant{
				{"Original", "100ml", 55},
			},
			VariantCount: 1,
			Reviews: [MAX_REVIEW]Review{
				{"@sensitiveskin", "Kulit jadi lebih tenang dan tidak merah", 5},
				{"@cicalover", "Toner favorit!", 5},
				{"@skinroutine", "Cocok banget buat kulit sensitif", 4},
			},
			ReviewCount: 3,
		},
		// 11
		{
			ID: "TSM011", Name: "Blush On Cream", Category: "makeup", Price: 72000, Sold: 260,
			BrandInfo:  Brand{"Emina", "Indonesia"},
			RateInfo:   Rating{4.4, 140},
			DetailInfo: ProductDetail{"Blush on cream dengan hasil akhir natural dan tahan lama", "all", 2026},
			Variants: [MAX_VARIANT]Variant{
				{"Peach", "5g", 40},
				{"Coral", "5g", 35},
				{"Rose", "5g", 28},
			},
			VariantCount: 3,
			Reviews: [MAX_REVIEW]Review{
				{"@blushonlover", "Warnanya cantik banget!", 5},
				{"@creamblush", "Mudah diaplikasikan", 4},
			},
			ReviewCount: 2,
		},
		// 12
		{
			ID: "TSM012", Name: "Mascara Waterproof", Category: "makeup", Price: 58000, Sold: 315,
			BrandInfo:  Brand{"Viva", "Indonesia"},
			RateInfo:   Rating{4.3, 190},
			DetailInfo: ProductDetail{"Maskara waterproof untuk bulu mata lebih panjang dan bervolume", "all", 2025},
			Variants: [MAX_VARIANT]Variant{
				{"Black", "8ml", 60},
			},
			VariantCount: 1,
			Reviews: [MAX_REVIEW]Review{
				{"@mascaralovers", "Tidak luntur meski hujan!", 5},
				{"@eyelashqueen", "Bulu mata jadi lebih tebal", 4},
			},
			ReviewCount: 2,
		},
		// 13
		{
			ID: "TSM013", Name: "Eyeliner Pen Waterproof", Category: "makeup", Price: 45000, Sold: 490,
			BrandInfo:  Brand{"Implora", "Indonesia"},
			RateInfo:   Rating{4.6, 320},
			DetailInfo: ProductDetail{"Eyeliner pen waterproof dengan ujung yang presisi dan tahan lama", "all", 2026},
			Variants: [MAX_VARIANT]Variant{
				{"Black", "0.5ml", 90},
				{"Brown", "0.5ml", 45},
			},
			VariantCount: 2,
			Reviews: [MAX_REVIEW]Review{
				{"@eyelinerfan", "Presisi banget ujungnya!", 5},
				{"@makeupnerd", "Tahan lama seharian", 5},
				{"@beautyid", "Gampang dipake", 4},
			},
			ReviewCount: 3,
		},
		// 14
		{
			ID: "TSM014", Name: "Lip Balm SPF 15", Category: "lip care", Price: 35000, Sold: 620,
			BrandInfo:  Brand{"Wardah", "Indonesia"},
			RateInfo:   Rating{4.5, 270},
			DetailInfo: ProductDetail{"Lip balm dengan SPF 15 untuk menjaga kelembapan bibir", "all", 2026},
			Variants: [MAX_VARIANT]Variant{
				{"Original", "4g", 100},
				{"Strawberry", "4g", 85},
				{"Vanilla", "4g", 70},
			},
			VariantCount: 3,
			Reviews: [MAX_REVIEW]Review{
				{"@liplover", "Bibir jadi lembap seharian!", 5},
				{"@lipcare101", "Aroma strawberry-nya enak", 5},
			},
			ReviewCount: 2,
		},
		// 15
		{
			ID: "TSM015", Name: "Face Mask Sheet Aloe Vera", Category: "skincare", Price: 15000, Sold: 980,
			BrandInfo:  Brand{"Esqa", "Indonesia"},
			RateInfo:   Rating{4.4, 540},
			DetailInfo: ProductDetail{"Sheet mask aloe vera untuk menenangkan dan menghidrasi kulit", "sensitive", 2026},
			Variants: [MAX_VARIANT]Variant{
				{"Aloe Vera", "25ml", 200},
			},
			VariantCount: 1,
			Reviews: [MAX_REVIEW]Review{
				{"@sheetmasklover", "Kulit terasa segar setelahnya!", 5},
				{"@skincareholic", "Cocok banget buat kulit sensitif", 4},
				{"@beautycare", "Harganya terjangkau banget", 5},
			},
			ReviewCount: 3,
		},
		// 16
		{
			ID: "TSM016", Name: "Concealer Stick", Category: "makeup", Price: 62000, Sold: 275,
			BrandInfo:  Brand{"Emina", "Indonesia"},
			RateInfo:   Rating{4.3, 155},
			DetailInfo: ProductDetail{"Concealer stick untuk menyamarkan noda dan mata panda", "all", 2026},
			Variants: [MAX_VARIANT]Variant{
				{"Porcelain", "3g", 45},
				{"Beige", "3g", 50},
				{"Honey", "3g", 30},
			},
			VariantCount: 3,
			Reviews: [MAX_REVIEW]Review{
				{"@concealerfan", "Nutupin mata panda dengan baik!", 5},
				{"@makeupbase", "Mudah di-blend", 4},
			},
			ReviewCount: 2,
		},
		// 17
		{
			ID: "TSM017", Name: "Exfoliating Toner AHA BHA", Category: "skincare", Price: 145000, Sold: 330,
			BrandInfo:  Brand{"Skintific", "Indonesia"},
			RateInfo:   Rating{4.7, 210},
			DetailInfo: ProductDetail{"Toner eksfoliasi AHA BHA untuk mengangkat sel kulit mati", "oily", 2027},
			Variants: [MAX_VARIANT]Variant{
				{"Original", "100ml", 65},
			},
			VariantCount: 1,
			Reviews: [MAX_REVIEW]Review{
				{"@exfoliatefan", "Kulit terasa halus setelah pakai!", 5},
				{"@ahabhalovers", "Tekstur kulit jadi lebih merata", 5},
				{"@skinroutineid", "Jangan lupa sunscreen ya!", 4},
			},
			ReviewCount: 3,
		},
		// 18
		{
			ID: "TSM018", Name: "Micellar Water", Category: "skincare", Price: 42000, Sold: 710,
			BrandInfo:  Brand{"Wardah", "Indonesia"},
			RateInfo:   Rating{4.6, 480},
			DetailInfo: ProductDetail{"Micellar water untuk membersihkan makeup dan kotoran secara lembut", "all", 2027},
			Variants: [MAX_VARIANT]Variant{
				{"Original", "200ml", 130},
			},
			VariantCount: 1,
			Reviews: [MAX_REVIEW]Review{
				{"@makeupremover", "Bersih banget tanpa perlu digosok!", 5},
				{"@cleansinglover", "Cocok untuk semua jenis kulit", 4},
			},
			ReviewCount: 2,
		},
		// 19
		{
			ID: "TSM019", Name: "Highlighter Powder", Category: "makeup", Price: 88000, Sold: 198,
			BrandInfo:  Brand{"Esqa", "Indonesia"},
			RateInfo:   Rating{4.8, 120},
			DetailInfo: ProductDetail{"Highlighter powder dengan shimmer halus untuk tampilan glowing", "all", 2027},
			Variants: [MAX_VARIANT]Variant{
				{"Champagne Gold", "8g", 35},
				{"Rose Gold", "8g", 40},
				{"Pearl White", "8g", 25},
			},
			VariantCount: 3,
			Reviews: [MAX_REVIEW]Review{
				{"@glowingskin", "Highlighter terbaik!", 5},
				{"@shimmerqueen", "Shimmer-nya halus dan tidak glittery", 5},
				{"@makeupglow", "Pigmentasinya luar biasa", 5},
			},
			ReviewCount: 3,
		},
		// 20
		{
			ID: "TSM020", Name: "Eye Cream Retinol", Category: "skincare", Price: 185000, Sold: 145,
			BrandInfo:  Brand{"Avoskin", "Indonesia"},
			RateInfo:   Rating{4.6, 85},
			DetailInfo: ProductDetail{"Krim mata dengan retinol untuk mengurangi kerutan dan lingkar hitam", "normal", 2027},
			Variants: [MAX_VARIANT]Variant{
				{"Original", "15ml", 30},
			},
			VariantCount: 1,
			Reviews: [MAX_REVIEW]Review{
				{"@antiaging", "Mata terasa lebih segar!", 5},
				{"@eyecare101", "Lingkar hitam berkurang signifikan", 4},
			},
			ReviewCount: 2,
		},
		// 21
		{
			ID: "TSM021", Name: "Contour Stick Duo", Category: "makeup", Price: 78000, Sold: 230,
			BrandInfo:  Brand{"Implora", "Indonesia"},
			RateInfo:   Rating{4.4, 130},
			DetailInfo: ProductDetail{"Contour stick duo untuk shading dan highlighting wajah", "all", 2026},
			Variants: [MAX_VARIANT]Variant{
				{"Light", "5g", 40},
				{"Medium", "5g", 38},
				{"Dark", "5g", 22},
			},
			VariantCount: 3,
			Reviews: [MAX_REVIEW]Review{
				{"@contourlover", "Mudah diaplikasikan!", 5},
				{"@facesculpt", "Hasilnya natural banget", 4},
			},
			ReviewCount: 2,
		},
		// 22
		{
			ID: "TSM022", Name: "Setting Spray", Category: "makeup", Price: 55000, Sold: 365,
			BrandInfo:  Brand{"Emina", "Indonesia"},
			RateInfo:   Rating{4.5, 200},
			DetailInfo: ProductDetail{"Setting spray untuk mengunci makeup agar tahan lama seharian", "all", 2026},
			Variants: [MAX_VARIANT]Variant{
				{"Matte Finish", "60ml", 75},
				{"Dewy Finish", "60ml", 60},
			},
			VariantCount: 2,
			Reviews: [MAX_REVIEW]Review{
				{"@makeuplast", "Makeup bertahan seharian!", 5},
				{"@settingspray", "Packaging-nya mungil tapi isinya banyak", 4},
				{"@beautyfix", "Worth it banget!", 5},
			},
			ReviewCount: 3,
		},
		// 23
		{
			ID: "TSM023", Name: "Vitamin C Serum 20%", Category: "skincare", Price: 160000, Sold: 420,
			BrandInfo:  Brand{"Somethinc", "Indonesia"},
			RateInfo:   Rating{4.8, 310},
			DetailInfo: ProductDetail{"Serum vitamin C 20% untuk mencerahkan kulit dan memudarkan bekas jerawat", "all", 2027},
			Variants: [MAX_VARIANT]Variant{
				{"Original", "20ml", 80},
			},
			VariantCount: 1,
			Reviews: [MAX_REVIEW]Review{
				{"@vitaminclover", "Wajah cerah dalam 2 minggu!", 5},
				{"@brightenskin", "Bekas jerawat mulai memudar", 5},
			},
			ReviewCount: 2,
		},
		// 24
		{
			ID: "TSM024", Name: "Lip Tint Water", Category: "lip care", Price: 40000, Sold: 534,
			BrandInfo:  Brand{"Wardah", "Indonesia"},
			RateInfo:   Rating{4.6, 290},
			DetailInfo: ProductDetail{"Lip tint berbahan dasar air dengan warna natural dan tahan lama", "all", 2026},
			Variants: [MAX_VARIANT]Variant{
				{"Cherry Red", "5ml", 55},
				{"Rosy Pink", "5ml", 62},
				{"Coral Orange", "5ml", 48},
			},
			VariantCount: 3,
			Reviews: [MAX_REVIEW]Review{
				{"@liptintlover", "Warnanya cantik dan natural!", 5},
				{"@kbeautyfan", "K-beauty vibes banget!", 4},
				{"@liplover", "Ringan di bibir", 5},
			},
			ReviewCount: 3,
		},
		// 25
		{
			ID: "TSM025", Name: "Clay Mask Kaolin", Category: "skincare", Price: 52000, Sold: 395,
			BrandInfo:  Brand{"Skintific", "Indonesia"},
			RateInfo:   Rating{4.5, 220},
			DetailInfo: ProductDetail{"Clay mask kaolin untuk membersihkan pori dan mengontrol minyak berlebih", "oily", 2026},
			Variants: [MAX_VARIANT]Variant{
				{"Original", "60ml", 85},
			},
			VariantCount: 1,
			Reviews: [MAX_REVIEW]Review{
				{"@claymaskfan", "Pori-pori bersih banget setelahnya!", 5},
				{"@oilcontrol", "Cocok banget buat kulit berminyak", 5},
			},
			ReviewCount: 2,
		},
		// 26
		{
			ID: "TSM026", Name: "Brow Pencil Micro", Category: "makeup", Price: 48000, Sold: 445,
			BrandInfo:  Brand{"Implora", "Indonesia"},
			RateInfo:   Rating{4.7, 260},
			DetailInfo: ProductDetail{"Pensil alis micro dengan ujung sangat tipis untuk hasil natural", "all", 2026},
			Variants: [MAX_VARIANT]Variant{
				{"Dark Brown", "0.09g", 90},
				{"Medium Brown", "0.09g", 75},
				{"Soft Black", "0.09g", 60},
			},
			VariantCount: 3,
			Reviews: [MAX_REVIEW]Review{
				{"@browsonfleek", "Sangat presisi dan natural!", 5},
				{"@browgame", "Mirip banget sama bulu alis asli", 5},
				{"@makeupartist", "Highly recommended!", 5},
			},
			ReviewCount: 3,
		},
		// 27
		{
			ID: "TSM027", Name: "Body Lotion Glowing", Category: "body care", Price: 65000, Sold: 602,
			BrandInfo:  Brand{"Viva", "Indonesia"},
			RateInfo:   Rating{4.4, 370},
			DetailInfo: ProductDetail{"Body lotion dengan formula glowing untuk kulit tubuh cerah dan lembap", "all", 2027},
			Variants: [MAX_VARIANT]Variant{
				{"Original", "200ml", 110},
			},
			VariantCount: 1,
			Reviews: [MAX_REVIEW]Review{
				{"@glowingbody", "Kulit badan jadi lebih cerah!", 5},
				{"@bodycarelovers", "Aroma wanginya tahan lama", 4},
			},
			ReviewCount: 2,
		},
		// 28
		{
			ID: "TSM028", Name: "Primer Face Pore Minimizer", Category: "makeup", Price: 92000, Sold: 278,
			BrandInfo:  Brand{"Azarine", "Indonesia"},
			RateInfo:   Rating{4.6, 165},
			DetailInfo: ProductDetail{"Primer wajah untuk menyamarkan pori dan membuat makeup lebih tahan lama", "all", 2027},
			Variants: [MAX_VARIANT]Variant{
				{"Original", "30ml", 55},
			},
			VariantCount: 1,
			Reviews: [MAX_REVIEW]Review{
				{"@primerlover", "Makeup jadi tahan lama dan pori samar!", 5},
				{"@makeupbase", "Teksturnya ringan dan cepat meresap", 4},
				{"@beautyprep", "Wajib punya nih!", 5},
			},
			ReviewCount: 3,
		},
		// 29
		{
			ID: "TSM029", Name: "Lip Scrub Sugar", Category: "lip care", Price: 28000, Sold: 480,
			BrandInfo:  Brand{"Emina", "Indonesia"},
			RateInfo:   Rating{4.5, 245},
			DetailInfo: ProductDetail{"Scrub bibir dengan gula alami untuk mengangkat sel kulit mati pada bibir", "all", 2026},
			Variants: [MAX_VARIANT]Variant{
				{"Strawberry", "10g", 80},
				{"Vanilla", "10g", 65},
			},
			VariantCount: 2,
			Reviews: [MAX_REVIEW]Review{
				{"@lipscrub", "Bibir jadi halus seketika!", 5},
				{"@skincareroutine", "Rasanya enak juga wkwk", 5},
			},
			ReviewCount: 2,
		},
		// 30
		{
			ID: "TSM030", Name: "Retinol Night Cream", Category: "skincare", Price: 210000, Sold: 165,
			BrandInfo:  Brand{"Avoskin", "Indonesia"},
			RateInfo:   Rating{4.8, 100},
			DetailInfo: ProductDetail{"Night cream dengan retinol dan peptide untuk regenerasi kulit saat tidur", "normal", 2027},
			Variants: [MAX_VARIANT]Variant{
				{"Original", "30ml", 40},
			},
			VariantCount: 1,
			Reviews: [MAX_REVIEW]Review{
				{"@nightskincare", "Kulit terasa lebih kencang setelah 2 minggu!", 5},
				{"@retinolfan", "Wajib buat anti-aging routine", 5},
				{"@glowingskin", "Hasilnya nyata banget!", 4},
			},
			ReviewCount: 3,
		},
	}
}

// ======================================================
// UI
// ======================================================

func hLine(n int) {
	var i int
	for i = 0; i < n; i++ {
		fmt.Print("=")
	}
	fmt.Println()
}

func hLineThin(n int) {
	var i int
	for i = 0; i < n; i++ {
		fmt.Print("-")
	}
	fmt.Println()
}

func hLineSep(n int) {
	fmt.Print("+")
	var i int
	for i = 0; i < n-2; i++ {
		fmt.Print("-")
	}
	fmt.Print("+")
	fmt.Println()
}

func boxTop(n int) {
	fmt.Print("+")
	var i int
	for i = 0; i < n-2; i++ {
		fmt.Print("=")
	}
	fmt.Print("+")
	fmt.Println()
}

func boxBottom(n int) {
	fmt.Print("+")
	var i int
	for i = 0; i < n-2; i++ {
		fmt.Print("=")
	}
	fmt.Print("+")
	fmt.Println()
}

func clearScreen() {
	fmt.Print("\033[H\033[2J")
}

func next() {
	fmt.Println()
	fmt.Print("  Tekan Enter untuk lanjut...")
	fmt.Scanln()
	fmt.Scanln()
	clearScreen()
}

// ======================================================
// LOGO & TITLE
// ======================================================

func logo() {
	clearScreen()
	var pink, reset string
	pink = "\033[38;2;245;187;212m"
	reset = "\033[0m"
	fmt.Printf(`%s
0110101010110101010101010 1 0 10101 011 0101 1 1 01 1 10 1 1 1 101 1 01 01 1 01 1 01
10010      10101 10 1 1 0     10 10 01  10   0 0 10 1 00 0 1 0 010 0 10 0  0 0  0
  001       011  0  0  0      0  1  1    1       1  0  1 0 1    1  1  1 0  1 0  1
   1         0   1     0         0       1          0      1       0    1    0  1
             1                           0                                      1
             0                           1
      _____                    _____                    _____
     /\    \                  /\    \                  /\    \
    /::\    \                /::\    \                /::\____\
    \:::\    \              /::::\    \              /::::|   |
     \:::\    \            /::::::\    \            /:::::|   |
      \:::\    \          /:::/\:::\    \          /::::::|   |
       \:::\    \        /:::/__\:::\    \        /:::/|::|   |
       /::::\    \       \:::\   \:::\    \      /:::/ |::|   |
      /::::::\    \    ___\:::\   \:::\    \    /:::/  |::|___|______
     /:::/\:::\    \  /\   \:::\   \:::\    \  /:::/   |::::::::\    \
    /:::/  \:::\____\/::\   \:::\   \::/    /\::/    / ~~~~~/:::/    /
   /:::/    \::/    /\:::\   \:::\   \/____/  \/____/      /:::/    /
  /:::/    / \/____/  \:::\   \:::\    \                  /:::/    /
 /:::/    /            \:::\   \:::\____\                /:::/    /
/:::/    /              \:::\   \:::\    \               /:::/    /
\::/    /                \:::\  /:::/    /               /:::/    /
 \/____/                  \:::\/:::/    /               /:::/    /
                           \::::::/    /               /:::/    /
                            \::::/    /               /:::/    /
                             \::/    /                \::/    /
                              \/____/                  \/____/
%s`, pink, reset)
	fmt.Println()
	hLine(72)
	fmt.Println("                   Tetap Slay Maksimal")
	hLine(72)
	fmt.Println()
	next()
}

func pageTitle(text string) {
	fmt.Println()
	boxTop(54)

	fmt.Printf("| %-50s |\n", text)
	boxBottom(54)
	fmt.Println()
}

// ======================================================
// MESSAGE BOXES
// ======================================================

func msgSuccess(msg string) {
	fmt.Println()
	boxTop(54)
	fmt.Printf("| %-50s |\n", "SUCCESS <3")
	hLineSep(54)
	fmt.Printf("| %-50s |\n", msg)
	boxBottom(54)
	fmt.Println()
	fmt.Println("  Tekan Enter untuk melanjutkan!!")
	fmt.Scanln()
}

func msgError(msg string) {
	fmt.Println()
	boxTop(54)
	fmt.Printf("| %-50s |\n", "ERROR -_-")
	hLineSep(54)
	fmt.Printf("| %-50s |\n", msg)
	boxBottom(54)
	fmt.Println()
	fmt.Println("  Tekan Enter untuk melanjutkan!!")
	fmt.Scanln()
}

func msgWarning(msg string) {
	fmt.Println()
	boxTop(54)
	fmt.Printf("| %-50s |\n", " PERINGATAN >_<")
	hLineSep(54)
	fmt.Printf("| %-50s |\n", msg)
	boxBottom(54)
	fmt.Println()
	fmt.Println("  Tekan Enter untuk melanjutkan!!")
	fmt.Scanln()
}

func msgInfo(msg string) {
	fmt.Println()
	hLineSep(54)
	fmt.Printf("| %-50s |\n", msg)
	hLineSep(54)
	fmt.Println()
	fmt.Println("  Tekan Enter untuk melanjutkan!!")
	fmt.Scanln()
}

// ======================================================
// INPUT HELPERS
// ======================================================

func inputInt(prompt string) int {
	var err error
	var value int
	var valid bool
	valid = false
	for !valid {
		fmt.Print("  " + prompt)
		_, err = fmt.Scan(&value)
		if err != nil {
			msgWarning("Input harus berupa angka!")
			var dummy string
			fmt.Scanln(&dummy)
		} else if value < 0 {
			msgWarning("Input tidak boleh negatif!")
		} else {
			valid = true
		}
	}
	return value
}

func inputString(prompt string) string {
	var x string
	fmt.Print("  " + prompt)
	fmt.Scan(&x)
	return x
}

func inputStringConfirm(prompt string) string {
	var x string
	fmt.Print("  " + prompt + " [y/n]: ")
	fmt.Scan(&x)
	return x
}

// ======================================================
// MENU RENDERER
// ======================================================

func renderMenu(titleText string, items []string, n int) {
	pageTitle(titleText)
	var i int
	for i = 0; i < n; i++ {
		fmt.Printf("  %d.  %s\n", i+1, items[i])
	}
	fmt.Println("  0.  Kembali")
	fmt.Println()
	hLineThin(54)
}

// ======================================================
// MENU FLOWS
// ======================================================

func mainMenu() {
	clearScreen()
	for {
		var items = []string{
			"Kelola Produk",
			"Lihat Produk",
			"Rekomendasi",
			"Laporan Penjualan",
		}
		renderMenu("MENU UTAMA", items, 4)
		var c int
		c = inputInt("Pilih menu: ")
		switch c {
		case 1:
			crudMenu()
		case 2:
			viewMenu()
		case 3:
			recommendationMenu()
		case 4:
			salesMenu()
		case 0:
			var confirm string
			for {
				confirm = inputStringConfirm("Apakah Anda yakin ingin keluar dari program?")
				if confirm == "y" || confirm == "Y" {
					clearScreen()
					fmt.Println()
					msgInfo("Program selesai. Terima kasih!")
					fmt.Println()
					return
				}
				if confirm == "n" || confirm == "N" {
					clearScreen()
					mainMenu()
					return
				}
				msgWarning("Masukkan hanya 'y' atau 'n'.")
			}
		default:
			msgWarning("Menu tidak tersedia. Pilih 0-4.")
		}
		clearScreen()
	}
}

func crudMenu() {
	for {
		var items = []string{
			"Tambah Produk",
			"Perbarui Produk",
			"Hapus Produk",
			"Lihat Semua Produk",
		}
		renderMenu("KELOLA PRODUK", items, 4)
		var c int
		c = inputInt("Pilih menu: ")
		switch c {
		case 1:
			createProduct()
		case 2:
			updateProduct()
		case 3:
			deleteProduct()
		case 4:
			printAllProducts()
			next()
		case 0:
			return
		default:
			msgWarning("Menu tidak tersedia. Pilih 0-4.")
		}
		clearScreen()
	}
}

func viewMenu() {
	if isEmpty() {
		msgWarning("Belum ada data produk.")
		return
	}
	for {
		var items = []string{
			"Semua Produk",
			"Cari Produk by ID",
			"Cari Produk by Nama",
			"Detail Lengkap Produk",
			"Urutkan Produk",
		}
		renderMenu("LIHAT PRODUK", items, 5)
		var c int
		c = inputInt("Pilih menu: ")
		switch c {
		case 1:
			printAllProducts()
			next()
		case 2:
			clearScreen()
			pageTitle("CARI PRODUK BY ID")
			fmt.Println("  (Binary Search — ID harus urut)")
			fmt.Println()
			viewByID(false)
			next()
		case 3:
			clearScreen()
			pageTitle("CARI PRODUK BY NAMA")
			fmt.Println("  (Sequential Search)")
			fmt.Println()
			searchByName()
			next()
		case 4:
			clearScreen()
			pageTitle("DETAIL LENGKAP PRODUK")
			fmt.Println()
			viewByID(true)
			next()
		case 5:
			sortMenu()
		case 0:
			return
		default:
			msgWarning("Menu tidak tersedia. Pilih 0-5.")
		}
		clearScreen()
	}
}

func recommendationMenu() {
	for {
		var items = []string{
			"Produk Terlaris",
			"Stok Menipis",
		}
		renderMenu("REKOMENDASI", items, 2)
		var c int
		c = inputInt("Pilih menu: ")
		switch c {
		case 1:
			bestSellingRecommendation()
		case 2:
			lowStockRecommendation()
		case 0:
			return
		default:
			msgWarning("Menu tidak tersedia. Pilih 0-2.")
		}
		clearScreen()
	}
}

func salesMenu() {
	for {
		var items = []string{
			"Produk Paling Laku",
			"Produk Paling Sedikit Laku",
			"Total Penjualan",
		}
		renderMenu("LAPORAN PENJUALAN", items, 3)
		var c int
		c = inputInt("Pilih menu: ")
		switch c {
		case 1:
			mostSoldProduct()
			next()
		case 2:
			leastSoldProduct()
			next()
		case 3:
			fmt.Printf("\n  Total penjualan keseluruhan: %d unit\n\n", totalSales())
			next()
		case 0:
			return
		default:
			msgWarning("Menu tidak tersedia. Pilih 0-3.")
		}
		clearScreen()
	}
}

func sortMenu() {
	for {
		var items = []string{
			"Harga: Murah -> Mahal  (Insertion Sort)",
			"Harga: Mahal -> Murah  (Insertion Sort)",
			"Nama:  A -> Z          (Selection Sort)",
			"Nama:  Z -> A          (Selection Sort)",
		}
		renderMenu("URUTKAN PRODUK", items, 4)
		var c int
		c = inputInt("Pilih urutan: ")
		switch c {
		case 1:
			printSortedProducts("insertion_asc")
			next()
		case 2:
			printSortedProducts("insertion_desc")
			next()
		case 3:
			printSortedProducts("selection_asc")
			next()
		case 4:
			printSortedProducts("selection_desc")
			next()
		case 0:
			return
		default:
			msgWarning("Menu tidak tersedia. Pilih 0-4.")
		}
		clearScreen()
	}
}

// ======================================================
// CRUD
// ======================================================

func createProduct() {
	clearScreen()
	pageTitle("TAMBAH PRODUK BARU")

	if isFull(countData) {
		msgError("Data produk penuh! Kapasitas maksimal " + fmt.Sprintf("%d", NMAX) + " produk.")
		next()
		return
	}

	var p Product
	p.ID = generateID()

	fmt.Println()
	hLineSep(54)
	fmt.Println("| LANGKAH 1/4 - INFO PRODUK |")
	hLineSep(54)

	p.Name = inputValidString("Nama Produk        : ", 50)
	p.Category = inputValidString("Kategori           : ", 30)
	p.Price = inputValidPrice("Harga (Rp)         : ")

	fmt.Println()
	hLineSep(54)
	fmt.Println("| LANGKAH 2/4 - INFO BRAND |")
	hLineSep(54)

	p.BrandInfo.Name = inputValidString("Nama Brand         : ", 50)
	p.BrandInfo.Country = inputValidString("Negara Asal Brand  : ", 30)

	fmt.Println()
	hLineSep(54)
	fmt.Println("| LANGKAH 3/4 - DETAIL PRODUK |")
	hLineSep(54)

	p.DetailInfo.Description = inputValidString("Deskripsi          : ", 100)
	p.DetailInfo.SkinType = inputValidString("Jenis Kulit        : ", 30)
	p.DetailInfo.ExpiredYear = inputValidYear("Tahun Kedaluwarsa  : ")

	fmt.Println()
	hLineSep(54)
	fmt.Println("| LANGKAH 4/4 - VARIAN PRODUK |")
	hLineSep(54)

	var variantCount int
	variantCount = 0
	for variantCount < 1 || variantCount > MAX_VARIANT {
		variantCount = inputInt(fmt.Sprintf("Jumlah Varian (1-%d): ", MAX_VARIANT))
		if variantCount < 1 || variantCount > MAX_VARIANT {
			msgWarning(fmt.Sprintf("Jumlah varian harus antara 1 sampai %d.", MAX_VARIANT))
		}
	}
	p.VariantCount = variantCount

	var i int
	for i = 0; i < p.VariantCount; i++ {
		fmt.Println()
		fmt.Printf("  -- Varian %d dari %d --\n", i+1, p.VariantCount)
		p.Variants[i].Color = inputValidString("  Warna  : ", 20)
		p.Variants[i].Size = inputValidString("  Ukuran : ", 10)
		p.Variants[i].Stock = inputValidStock("  Stok   : ")
	}

	p.Sold = 0
	p.RateInfo.Score = 0
	p.RateInfo.TotalReview = 0
	p.ReviewCount = 0

	clearScreen()
	pageTitle("KONFIRMASI PRODUK BARU")
	hLineSep(54)
	fmt.Println("| RINGKASAN DATA |")
	hLineSep(54)
	fmt.Printf("  ID Baru            : %s\n", p.ID)
	fmt.Printf("  Nama               : %s\n", p.Name)
	fmt.Printf("  Kategori           : %s\n", p.Category)
	fmt.Printf("  Harga              : Rp %d\n", p.Price)
	fmt.Printf("  Brand              : %s (%s)\n", p.BrandInfo.Name, p.BrandInfo.Country)
	fmt.Printf("  Jumlah Varian      : %d\n", p.VariantCount)
	hLineSep(54)

	var confirm string
	confirm = inputStringConfirm("Simpan produk ini?")
	if confirm != "y" && confirm != "Y" {
		msgInfo("Penambahan produk dibatalkan.")
		next()
		return
	}

	productsArr[countData] = p
	countData++
	historyCountData++

	msgSuccess("Produk berhasil ditambahkan! ID: " + p.ID)
	next()
}

func updateProduct() {
	clearScreen()
	pageTitle("PERBARUI PRODUK")

	if isEmpty() {
		msgWarning("Belum ada data produk untuk diperbarui.")
		next()
		return
	}

	printAllProducts()

	fmt.Println()
	fmt.Println("  Ketik 'batal' untuk membatalkan.")

	var id string
	id = inputString("ID yang ingin diperbarui: ")

	if id == "batal" {
		msgInfo("Update dibatalkan.")
		next()
		return
	}

	var idx int
	idx = findIndexByID(id)

	if idx == -1 {
		msgError("ID tidak ditemukan.")
		next()
		return
	}

	var oldData Product
	oldData = productsArr[idx]

	clearScreen()
	pageTitle("DATA SAAT INI")

	hLineSep(60)
	fmt.Println("  ID              :", productsArr[idx].ID)
	fmt.Println("  Nama            :", productsArr[idx].Name)
	fmt.Println("  Kategori        :", productsArr[idx].Category)
	fmt.Println("  Harga           :", productsArr[idx].Price)
	fmt.Println("  Brand           :", productsArr[idx].BrandInfo.Name)
	fmt.Println("  Negara Brand    :", productsArr[idx].BrandInfo.Country)
	fmt.Println("  Deskripsi       :", productsArr[idx].DetailInfo.Description)
	fmt.Println("  Jenis Kulit     :", productsArr[idx].DetailInfo.SkinType)
	fmt.Println("  Tahun Expired   :", productsArr[idx].DetailInfo.ExpiredYear)
	hLineSep(60)

	fmt.Println()
	fmt.Println("  Pilih bagian yang ingin diperbarui:")
	fmt.Println()
	fmt.Println("  1. Nama Produk")
	fmt.Println("  2. Kategori")
	fmt.Println("  3. Harga")
	fmt.Println("  4. Nama Brand")
	fmt.Println("  5. Negara Brand")
	fmt.Println("  6. Deskripsi")
	fmt.Println("  7. Jenis Kulit")
	fmt.Println("  8. Tahun Expired")
	fmt.Println("  9. Update Semua Data")
	fmt.Println("  0. Batal")
	fmt.Println()
	hLineThin(60)

	var pilihan int
	pilihan = inputInt("Pilih Menu : ")

	switch pilihan {

	case 0:
		msgInfo("Update dibatalkan.")
		next()
		return

	case 1:
		clearScreen()
		pageTitle("UPDATE NAMA PRODUK")
		productsArr[idx].Name =
			inputValidString("Nama Produk Baru : ", 50)

	case 2:
		clearScreen()
		pageTitle("UPDATE KATEGORI")
		productsArr[idx].Category =
			inputValidString("Kategori Baru : ", 30)

	case 3:
		clearScreen()
		pageTitle("UPDATE HARGA")
		productsArr[idx].Price =
			inputValidPrice("Harga Baru : ")

	case 4:
		clearScreen()
		pageTitle("UPDATE NAMA BRAND")
		productsArr[idx].BrandInfo.Name =
			inputValidString("Nama Brand Baru : ", 50)

	case 5:
		clearScreen()
		pageTitle("UPDATE NEGARA BRAND")
		productsArr[idx].BrandInfo.Country =
			inputValidString("Negara Brand Baru : ", 30)

	case 6:
		clearScreen()
		pageTitle("UPDATE DESKRIPSI")
		productsArr[idx].DetailInfo.Description =
			inputValidString("Deskripsi Baru : ", 100)

	case 7:
		clearScreen()
		pageTitle("UPDATE JENIS KULIT")
		productsArr[idx].DetailInfo.SkinType =
			inputValidString("Jenis Kulit Baru : ", 30)

	case 8:
		clearScreen()
		pageTitle("UPDATE TAHUN EXPIRED")
		productsArr[idx].DetailInfo.ExpiredYear =
			inputValidYear("Tahun Expired Baru : ")

	case 9:
		clearScreen()
		pageTitle("UPDATE SEMUA DATA")

		hLineSep(54)
		fmt.Println("| INFO PRODUK |")
		hLineSep(54)

		productsArr[idx].Name =
			inputValidString("Nama Produk        : ", 50)

		productsArr[idx].Category =
			inputValidString("Kategori           : ", 30)

		productsArr[idx].Price =
			inputValidPrice("Harga (Rp)         : ")

		fmt.Println()
		hLineSep(54)
		fmt.Println("| INFO BRAND |")
		hLineSep(54)

		productsArr[idx].BrandInfo.Name =
			inputValidString("Nama Brand         : ", 50)

		productsArr[idx].BrandInfo.Country =
			inputValidString("Negara Asal Brand  : ", 30)

		fmt.Println()
		hLineSep(54)
		fmt.Println("| DETAIL PRODUK |")
		hLineSep(54)

		productsArr[idx].DetailInfo.Description =
			inputValidString("Deskripsi          : ", 100)

		productsArr[idx].DetailInfo.SkinType =
			inputValidString("Jenis Kulit        : ", 30)

		productsArr[idx].DetailInfo.ExpiredYear =
			inputValidYear("Tahun Kedaluwarsa  : ")

	default:
		msgError("Pilihan tidak valid.")
		next()
		return
	}

	clearScreen()
	pageTitle("HASIL UPDATE PRODUK")

	hLineSep(60)
	fmt.Println("| SEBELUM |")
	hLineSep(60)
	fmt.Println("  ID              :", oldData.ID)
	fmt.Println("  Nama            :", oldData.Name)
	fmt.Println("  Kategori        :", oldData.Category)
	fmt.Println("  Harga           :", oldData.Price)
	fmt.Println("  Brand           :", oldData.BrandInfo.Name)
	fmt.Println("  Negara Brand    :", oldData.BrandInfo.Country)
	fmt.Println("  Deskripsi       :", oldData.DetailInfo.Description)
	fmt.Println("  Jenis Kulit     :", oldData.DetailInfo.SkinType)
	fmt.Println("  Tahun Expired   :", oldData.DetailInfo.ExpiredYear)

	fmt.Println()

	hLineSep(60)
	fmt.Println("| SESUDAH |")
	hLineSep(60)
	fmt.Println("  ID              :", productsArr[idx].ID)
	fmt.Println("  Nama            :", productsArr[idx].Name)
	fmt.Println("  Kategori        :", productsArr[idx].Category)
	fmt.Println("  Harga           :", productsArr[idx].Price)
	fmt.Println("  Brand           :", productsArr[idx].BrandInfo.Name)
	fmt.Println("  Negara Brand    :", productsArr[idx].BrandInfo.Country)
	fmt.Println("  Deskripsi       :", productsArr[idx].DetailInfo.Description)
	fmt.Println("  Jenis Kulit     :", productsArr[idx].DetailInfo.SkinType)
	fmt.Println("  Tahun Expired   :", productsArr[idx].DetailInfo.ExpiredYear)

	fmt.Println()

	hLineSep(60)
	fmt.Println("| RINGKASAN PERUBAHAN |")
	hLineSep(60)

	var adaPerubahan bool
	adaPerubahan = false

	if oldData.Name != productsArr[idx].Name {
		fmt.Println("  Nama :", oldData.Name, "->", productsArr[idx].Name)
		adaPerubahan = true
	}

	if oldData.Category != productsArr[idx].Category {
		fmt.Println("  Kategori :", oldData.Category, "->", productsArr[idx].Category)
		adaPerubahan = true
	}

	if oldData.Price != productsArr[idx].Price {
		fmt.Println("  Harga :", oldData.Price, "->", productsArr[idx].Price)
		adaPerubahan = true
	}

	if oldData.BrandInfo.Name != productsArr[idx].BrandInfo.Name {
		fmt.Println("  Brand :", oldData.BrandInfo.Name, "->", productsArr[idx].BrandInfo.Name)
		adaPerubahan = true
	}

	if oldData.BrandInfo.Country != productsArr[idx].BrandInfo.Country {
		fmt.Println("  Negara Brand :", oldData.BrandInfo.Country, "->", productsArr[idx].BrandInfo.Country)
		adaPerubahan = true
	}

	if oldData.DetailInfo.Description != productsArr[idx].DetailInfo.Description {
		fmt.Println("  Deskripsi :", oldData.DetailInfo.Description, "->", productsArr[idx].DetailInfo.Description)
		adaPerubahan = true
	}

	if oldData.DetailInfo.SkinType != productsArr[idx].DetailInfo.SkinType {
		fmt.Println("  Jenis Kulit :", oldData.DetailInfo.SkinType, "->", productsArr[idx].DetailInfo.SkinType)
		adaPerubahan = true
	}

	if oldData.DetailInfo.ExpiredYear != productsArr[idx].DetailInfo.ExpiredYear {
		fmt.Println("  Tahun Expired :", oldData.DetailInfo.ExpiredYear, "->", productsArr[idx].DetailInfo.ExpiredYear)
		adaPerubahan = true
	}

	if !adaPerubahan {
		fmt.Println("  (Tidak ada perubahan data)")
	}

	fmt.Println()
	msgSuccess("Produk berhasil diperbarui.")
	next()
}

func deleteProduct() {
	clearScreen()
	pageTitle("HAPUS PRODUK")

	if isEmpty() {
		msgWarning("Belum ada data produk untuk dihapus.")
		next()
		return
	}

	printAllProducts()
	fmt.Println()
	fmt.Println("  Ketik 'batal' untuk membatalkan.")

	var id string
	id = inputString("ID yang ingin dihapus: ")
	if id == "batal" {
		msgInfo("Hapus dibatalkan.")
		next()
		return
	}

	var idx int
	idx = findIndexByID(id)
	if idx == -1 {
		msgError("ID '" + id + "' tidak ditemukan dalam data.")
		next()
		return
	}

	clearScreen()
	pageTitle("KONFIRMASI HAPUS")
	hLineSep(54)
	fmt.Println("| PRODUK YANG AKAN DIHAPUS |")
	hLineSep(54)
	fmt.Printf("  ID     : %s\n", productsArr[idx].ID)
	fmt.Printf("  Nama   : %s\n", productsArr[idx].Name)
	fmt.Printf("  Harga  : Rp %d\n", productsArr[idx].Price)
	fmt.Printf("  Stok terjual : %d unit\n", productsArr[idx].Sold)
	hLineSep(54)
	fmt.Println("  Perhatian: tindakan ini tidak dapat dibatalkan setelah disimpan.")

	var confirm string
	confirm = inputStringConfirm("Yakin ingin menghapus produk ini?")
	if confirm != "y" && confirm != "Y" {
		msgInfo("Hapus dibatalkan.")
		next()
		return
	}

	var i int
	for i = idx; i < countData-1; i++ {
		productsArr[i] = productsArr[i+1]
	}
	var empty Product
	productsArr[countData-1] = empty
	countData--

	msgSuccess("Produk berhasil dihapus.")
	next()
}

// ======================================================
// VIEW
// ======================================================

func viewByID(detail bool) {
	fmt.Println()
	var id string
	id = inputString("Masukkan ID produk: ")

	var idx int
	idx = findIndexByID(id)

	if idx == -1 {
		msgError("Produk dengan ID '" + id + "' tidak ditemukan.")
		return
	}

	if detail {
		printProductDetail(idx)
	} else {
		printProductRow(idx)
	}
}

func printAllProducts() {
	clearScreen()
	pageTitle("DAFTAR PRODUK")

	if isEmpty() {
		msgInfo("Belum ada data produk.")
		return
	}

	var w int
	w = 120
	hLine(w)
	fmt.Printf("  %-10s | %-40s | %-18s | %-14s | %-8s\n",
		"ID", "Nama Produk", "Kategori", "Harga (Rp)", "Terjual")
	hLineThin(w)
	var i int
	for i = 0; i < countData; i++ {
		fmt.Printf("  %-10s | %-40s | %-18s | %-14d | %-8d\n",
			productsArr[i].ID,
			productsArr[i].Name,
			productsArr[i].Category,
			productsArr[i].Price,
			productsArr[i].Sold,
		)
	}
	hLine(w)
	fmt.Printf("  Total produk: %d\n", countData)
}

func printProductRow(idx int) {
	var w int
	w = 120
	hLine(w)
	fmt.Printf("  %-10s | %-40s | %-18s | %-14s | %-8s\n",
		"ID", "Nama Produk", "Kategori", "Harga (Rp)", "Terjual")
	hLineThin(w)
	fmt.Printf("  %-10s | %-40s | %-18s | %-14d | %-8d\n",
		productsArr[idx].ID,
		productsArr[idx].Name,
		productsArr[idx].Category,
		productsArr[idx].Price,
		productsArr[idx].Sold,
	)
	hLine(w)
}

func printProductDetail(idx int) {
	var p Product
	p = productsArr[idx]

	fmt.Println()
	boxTop(60)
	fmt.Println("|                 DETAIL PRODUK                 |")
	boxBottom(60)
	fmt.Println()

	hLineSep(60)
	fmt.Println("| INFO UMUM |")
	hLineSep(60)
	fmt.Printf("  %-22s: %s\n", "ID", p.ID)
	fmt.Printf("  %-22s: %s\n", "Nama", p.Name)
	fmt.Printf("  %-22s: %s\n", "Kategori", p.Category)
	fmt.Printf("  %-22s: Rp %d\n", "Harga", p.Price)
	fmt.Printf("  %-22s: %d unit\n", "Terjual", p.Sold)

	hLineSep(60)
	fmt.Println("| BRAND |")
	hLineSep(60)
	fmt.Printf("  %-22s: %s\n", "Nama Brand", p.BrandInfo.Name)
	fmt.Printf("  %-22s: %s\n", "Negara Asal", p.BrandInfo.Country)

	hLineSep(60)
	fmt.Println("| DETAIL PRODUK |")
	hLineSep(60)
	fmt.Printf("  %-22s: %s\n", "Deskripsi", p.DetailInfo.Description)
	fmt.Printf("  %-22s: %s\n", "Jenis Kulit", p.DetailInfo.SkinType)
	fmt.Printf("  %-22s: %d\n", "Tahun Kedaluwarsa", p.DetailInfo.ExpiredYear)

	hLineSep(60)
	fmt.Println("| VARIAN |")
	hLineSep(60)
	var i int
	for i = 0; i < p.VariantCount; i++ {
		fmt.Printf("  Varian %-2d  Warna: %-15s  Ukuran: %-8s  Stok: %d\n",
			i+1,
			p.Variants[i].Color,
			p.Variants[i].Size,
			p.Variants[i].Stock,
		)
	}
	hLineSep(60)
}

func printSortedProducts(mode string) {
	clearScreen()

	var sorted ProductArray
	sorted = productsArr

	switch mode {
	case "insertion_asc":
		pageTitle("PRODUK: HARGA MURAH KE MAHAL")
		insertionSortPriceAsc(&sorted, countData)
	case "insertion_desc":
		pageTitle("PRODUK: HARGA MAHAL KE MURAH")
		insertionSortPriceDesc(&sorted, countData)
	case "selection_asc":
		pageTitle("PRODUK: NAMA A -> Z")
		selectionSortNameAsc(&sorted, countData)
	default:
		pageTitle("PRODUK: NAMA Z -> A")
		selectionSortNameDesc(&sorted, countData)
	}

	var w int
	w = 120
	hLine(w)
	fmt.Printf("  %-10s | %-40s | %-18s | %-14s | %-8s\n",
		"ID", "Nama Produk", "Kategori", "Harga (Rp)", "Terjual")
	hLineThin(w)

	var i int
	for i = 0; i < countData; i++ {
		fmt.Printf("  %-10s | %-40s | %-18s | %-14d | %-8d\n",
			sorted[i].ID,
			sorted[i].Name,
			sorted[i].Category,
			sorted[i].Price,
			sorted[i].Sold,
		)
	}
	hLine(w)
	fmt.Printf("  Total produk: %d\n", countData)
	fmt.Println()

	if mode == "insertion_asc" || mode == "insertion_desc" {
		fmt.Println("  [Algoritma: Insertion Sort]")
	} else {
		fmt.Println("  [Algoritma: Selection Sort]")
	}
}

// ======================================================
// RECOMMENDATION
// ======================================================

func bestSellingRecommendation() {
	clearScreen()
	pageTitle("PRODUK TERLARIS")

	if isEmpty() {
		msgWarning("Belum ada data produk.")
		next()
		return
	}

	var sorted ProductArray
	sorted = productsArr
	var i int
	insertionSortSoldDesc(&sorted, countData)

	fmt.Println()
	hLine(80)
	fmt.Printf("  %-10s | %-35s | %-15s | %-8s\n", "ID", "Nama Produk", "Harga (Rp)", "Terjual")
	hLineThin(80)

	var limit int
	limit = countData
	if limit > 5 {
		limit = 5
	}
	for i = 0; i < limit; i++ {
		fmt.Printf("  %-10s | %-35s | %-15d | %-8d\n",
			sorted[i].ID, sorted[i].Name, sorted[i].Price, sorted[i].Sold)
	}
	hLine(80)
	fmt.Printf("\n  Menampilkan %d produk terlaris dari total %d produk.\n", limit, countData)
	next()
}

func lowStockRecommendation() {
	clearScreen()
	pageTitle("PRODUK STOK MENIPIS")

	if isEmpty() {
		msgWarning("Belum ada data produk.")
		next()
		return
	}

	var threshold int
	threshold = 5
	var found bool
	found = false

	fmt.Println()
	hLine(80)
	fmt.Printf("  %-10s | %-30s | %-10s | %-8s | %-8s\n",
		"ID", "Nama Produk", "Warna", "Ukuran", "Stok")
	hLineThin(80)

	var i, j int
	for i = 0; i < countData; i++ {
		for j = 0; j < productsArr[i].VariantCount; j++ {
			if productsArr[i].Variants[j].Stock <= threshold {
				found = true
				fmt.Printf("  %-10s | %-30s | %-10s | %-8s | %-8d\n",
					productsArr[i].ID,
					productsArr[i].Name,
					productsArr[i].Variants[j].Color,
					productsArr[i].Variants[j].Size,
					productsArr[i].Variants[j].Stock,
				)
			}
		}
	}
	hLine(80)
	if !found {
		fmt.Println()
		msgInfo(fmt.Sprintf("Aman! Tidak ada produk dengan stok menipis (stok <= %d).", threshold))
	}
	next()
}

// ======================================================
// SALES / STATISTICS
// ======================================================

func mostSoldProduct() {
	clearScreen()
	pageTitle("PRODUK PALING LAKU")

	if isEmpty() {
		msgWarning("Belum ada data produk.")
		return
	}

	var most int
	most = 0
	var i int
	for i = 1; i < countData; i++ {
		if productsArr[i].Sold > productsArr[most].Sold {
			most = i
		}
	}

	fmt.Println()
	printProductRow(most)
	fmt.Printf("\n  Total terjual: %d unit\n", productsArr[most].Sold)
}

func leastSoldProduct() {
	clearScreen()
	pageTitle("PRODUK PALING SEDIKIT LAKU")

	if isEmpty() {
		msgWarning("Belum ada data produk.")
		return
	}

	var i int
	var least int
	least = 0
	
	for i = 1; i < countData; i++ {
		if productsArr[i].Sold < productsArr[least].Sold {
			least = i
		}
	}

	fmt.Println()
	printProductRow(least)
	fmt.Printf("\n  Total terjual: %d unit\n", productsArr[least].Sold)
}

func totalSales() int {
	var i int
	var total int
	total = 0

	for i = 0; i < countData; i++ {
		total = total + productsArr[i].Sold
	}
	return total
}

// ======================================================
// SEARCHING
// ======================================================
func searchByName() {
	var name string
	name = inputString("Masukkan nama produk: ")

	var idx int
	idx = sequentialSearchName(name)

	if idx == -1 {
		msgError("Produk dengan nama '" + name + "' tidak ditemukan.")
		return
	}

	fmt.Println()
	hLineSep(54)
	fmt.Println("| PRODUK DITEMUKAN |")
	hLineSep(54)
	printProductRow(idx)
	fmt.Println()

	var jawab string
	jawab = inputStringConfirm("Tampilkan detail lengkap?")
	if jawab == "y" || jawab == "Y" {
		printProductDetail(idx)
	}
}

func sequentialSearchName(name string) int {
	var i int
	for i = 0; i < countData; i++ {
		if productsArr[i].Name == name {
			return i
		}
	}
	return -1
}

func findIndexByID(id string) int {
	var left, right, mid int
	left = 0
	right = countData - 1
	for left <= right {
		mid = (left + right) / 2
		if id == productsArr[mid].ID {
			return mid
		}
		if id > productsArr[mid].ID {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	return -1
}

// ======================================================
// SORTING
// ======================================================

func insertionSortPriceAsc(A *ProductArray, n int) {
	var i, j int
	var temp Product
	for i = 1; i <= n-1; i++ {
		j = i
		temp = A[j]
		for j > 0 && temp.Price < A[j-1].Price {
			A[j] = A[j-1]
			j = j - 1
		}
		A[j] = temp
	}
}

func insertionSortPriceDesc(A *ProductArray, n int) {
	var i, j int
	var temp Product
	for i = 1; i <= n-1; i++ {
		j = i
		temp = A[j]
		for j > 0 && temp.Price > A[j-1].Price {
			A[j] = A[j-1]
			j = j - 1
		}
		A[j] = temp
	}
}

func insertionSortSoldDesc(A *ProductArray, n int) {
	var i, j int
	var temp Product
	for i = 1; i <= n-1; i++ {
		j = i
		temp = A[j]
		for j > 0 && temp.Sold > A[j-1].Sold {
			A[j] = A[j-1]
			j = j - 1
		}
		A[j] = temp
	}
}

func selectionSortNameAsc(A *ProductArray, n int) {
	var i, j, idxMin int
	var temp Product
	for i = 0; i < n-1; i++ {
		idxMin = i
		for j = i + 1; j < n; j++ {
			if A[j].Name < A[idxMin].Name {
				idxMin = j
			}
		}
		temp = A[idxMin]
		A[idxMin] = A[i]
		A[i] = temp
	}
}

func selectionSortNameDesc(A *ProductArray, n int) {
	var i, j, idxMax int
	var temp Product
	for i = 0; i < n-1; i++ {
		idxMax = i
		for j = i + 1; j < n; j++ {
			if A[j].Name > A[idxMax].Name {
				idxMax = j
			}
		}
		temp = A[idxMax]
		A[idxMax] = A[i]
		A[i] = temp
	}
}

// ======================================================
// VALIDATION
// ======================================================

func isEmpty() bool {
	return countData == 0
}

func isFull(n int) bool {
	return n >= NMAX
}

func inputValidString(prompt string, maxLen int) string {
	var val string
	var valid bool
	valid = false
	for !valid {
		val = inputString(prompt)
		if val == "" {
			msgWarning("Input tidak boleh kosong!")
		} else if len(val) > maxLen {
			msgWarning(fmt.Sprintf("Input terlalu panjang! Maksimal %d karakter.", maxLen))
		} else {
			valid = true
		}
	}
	return val
}

func inputValidPrice(prompt string) int {
	var val int
	var valid bool
	valid = false
	for !valid {
		val = inputInt(prompt)
		if val <= 0 {
			msgWarning("Harga harus lebih dari 0!")
		} else {
			valid = true
		}
	}
	return val
}

func inputValidStock(prompt string) int {
	var val int
	var valid bool
	valid = false
	for !valid {
		val = inputInt(prompt)
		if val < 0 {
			msgWarning("Stok tidak boleh negatif!")
		} else {
			valid = true
		}
	}
	return val
}

func inputValidYear(prompt string) int {
	var val int
	var valid bool
	valid = false
	for !valid {
		val = inputInt(prompt)
		if val < 2026 {
			msgWarning("Tahun kedaluwarsa tidak valid! Minimal 2024.")
		} else if val > 2100 {
			msgWarning("Tahun kedaluwarsa tidak masuk akal! Maksimal 2100.")
		} else {
			valid = true
		}
	}
	return val
}

// ======================================================
// HELPER
// ======================================================

func generateID() string {
	return "TSM" + fmt.Sprintf("%03d", historyCountData+1)
}

// ======================================================
// MAIN
// ======================================================

func main() {
	logo()
	mainMenu()
}