package main

//https://patorjk.com/software/taag/#p=display&f=Graffiti&t=Type+Something+&x=none&v=4&h=4&w=80&we=false
import "fmt"

const (
	NMAX        int = 999
	MAX_VARIANT int = 5
	MAX_REVIEW  int = 10
	MAX_ITEM    int = 10
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
	ID       int
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

type Customer struct {
	ID    int
	Name  string
	Phone string
}

type SalesItem struct {
	ProductID   int
	ProductName string
	Price       int
	Quantity    int
	Subtotal    int
}

type Transaction struct {
	TransactionID int
	Buyer         Customer
	Items         [MAX_ITEM]SalesItem
	ItemCount     int
	TotalPayment  int
}

type ProductList [NMAX]Product
type TransactionList [NMAX]Transaction
type arrProduct [NMAX]Product

// ======================================================
// UTILITY
// ======================================================
func logo() {
	pink := "\033[38;2;245;187;212m"
	logo := fmt.Sprintf(`
%s
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
    /:::/  \:::\____\/::\   \:::\   \:::\____\/:::/    |:::::::::\____\
   /:::/    \::/    /\:::\   \:::\   \::/    /\::/    / ~~~~~/:::/    /
  /:::/    / \/____/  \:::\   \:::\   \/____/  \/____/      /:::/    / 
 /:::/    /            \:::\   \:::\    \                  /:::/    /  
/:::/    /              \:::\   \:::\____\                /:::/    /   
\::/    /                \:::\  /:::/    /               /:::/    /    
 \/____/                  \:::\/:::/    /               /:::/    /     
                           \::::::/    /               /:::/    /      
                            \::::/    /               /:::/    /       
                             \::/    /                \::/    /        
                              \/____/                  \/____/         
                                                                       
=======================================================================
|                        Tetap Slay Maksimal                          |
=======================================================================
`, pink)
	fmt.Println(logo)
}

func line(n int) {
	var i int
	for i = 0; i < n; i++ {
		fmt.Print("=")
	}
	fmt.Println()
}

func inputInt(text string) int {
	var x int
	fmt.Print(text)
	fmt.Scan(&x)
	return x
}

func clearScreen() {

}

func pause() {

}

func inputString(text string) string {
	return ""
}

// ======================================================
// MENU
// ======================================================

func showMenu(title string, menus [5]string, total int) {
	var i int
	fmt.Println(title)
	line(40)
	for i = 0; i < total; i++ {
		fmt.Printf("%d. %s\n", i+1, menus[i])
	}
	fmt.Println("0. Exit")
	line(40)
}

func menuCRUD() {
	var menus [5]string
	var c int

	menus[0] = "Create Barang"
	menus[1] = "Update Barang"
	menus[2] = "Delete Barang"
	menus[3] = "View Barang"
	menus[4] = "Back"

	for {
		showMenu("MENU CRUD", menus, 5)

		c = inputInt("Pilih menu: ")

		switch c {
		case 1:
			fmt.Println("CREATE")

		case 2:
			fmt.Println("UPDATE")

		case 3:
			fmt.Println("DELETE")

		case 4:
			fmt.Println("VIEW")

		case 5:
			return

		case 0:
			return

		default:
			fmt.Println("Menu tidak tersedia")
		}
	}
}

func menuUtama() {
	var menus [5]string
	var c int

	menus[0] = "Pengelola"
	menus[1] = "View Barang"
	menus[2] = "Rekomendasi"
	menus[3] = "Penjualan"

	for {
		showMenu("MENU UTAMA", menus, 4)

		c = inputInt("Pilih menu: ")

		switch c {
		case 1:
			menuCRUD()

		case 2:
			fmt.Println("VIEW BARANG")

		case 3:
			fmt.Println("REKOMENDASI")

		case 4:
			fmt.Println("PENJUALAN")

		case 0:
			fmt.Println("Program selesai")
			return

		default:
			fmt.Println("Menu tidak tersedia")
		}
	}
}

func menuSearch(A *arrProduct, n *int) {

}

func menuSort(A *arrProduct, n *int) {

}

func menuPenjualan(A *arrProduct, n *int) {

}

func menuRekomendasi(A *arrProduct, n *int) {

}

func menuStatistik(A *arrProduct, n *int) {

}

// ======================================================
// CRUD
// ======================================================

func createBarang(A *arrProduct, n *int) {

}

func viewBarang(A arrProduct, n int) {

}

func viewDetailBarang(A arrProduct, n int) {

}

func updateBarang(A *arrProduct, n int) {

}

func deleteBarang(A *arrProduct, n *int) {

}

// ======================================================
// VALIDATION
// ======================================================

func isEmpty(n int) bool {
	return false
}

func isFull(n int) bool {
	return false
}

func validasiHarga(harga int) bool {
	return false
}

func validasiStok(stok int) bool {
	return false
}

// ======================================================
// SEARCHING
// ======================================================

// Sequential Search
func sequentialSearchNama(A arrProduct, n int, nama string) int {
	return -1
}

func sequentialSearchKategori(A arrProduct, n int, kategori string) int {
	return -1
}

// Binary Search
func binarySearchID(A arrProduct, n int, id int) int {
	return -1
}

func binarySearchHarga(A arrProduct, n int, harga int) int {
	return -1
}

// ======================================================
// SORTING
// ======================================================

// Insertion Sort
func insertionSortHargaAsc(A *arrProduct, n int) {

}

func insertionSortHargaDesc(A *arrProduct, n int) {

}

// Selection Sort
func selectionSortNamaAsc(A *arrProduct, n int) {

}

func selectionSortNamaDesc(A *arrProduct, n int) {

}

// ======================================================
// PENJUALAN
// ======================================================

func penjualanBarang(A *arrProduct, n int) {

}

func checkoutBarang(A *arrProduct, n int) {

}

func hitungTotalBelanja(A arrProduct, n int) int {
	return 0
}

// ======================================================
// REKOMENDASI
// ======================================================

func rekomendasiTerlaris(A arrProduct, n int) {

}

func rekomendasiStokSedikit(A arrProduct, n int) {

}

func rekomendasiFashionPria(A arrProduct, n int) {

}

func rekomendasiFashionWanita(A arrProduct, n int) {

}

// ======================================================
// STATISTIK
// ======================================================

func totalBarang(A arrProduct, n int) int {
	return 0
}

func totalStok(A arrProduct, n int) int {
	return 0
}

func totalPenjualan(A arrProduct, n int) int {
	return 0
}

func barangPalingLaris(A arrProduct, n int) {

}

func barangPalingSedikitTerjual(A arrProduct, n int) {

}

// ======================================================
// HELPER
// ======================================================

func cariIndexByID(A arrProduct, n int, id int) int {
	return -1
}

func generateID(A arrProduct, n int) int {
	return 0
}

func swap(A *Product, B *Product) {

}

func printHeaderTable() {

}

func printBarang(B Product) {

}

// ======================================================
// MAIN
// ======================================================
func main() {
	logo()
	menuUtama()
}
