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
type ProductArray [NMAX]Product

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

func inputString(text string) string {
	var x string
	fmt.Print(text)
	fmt.Scan(&x)
	return x
}

func clearScreen() {

}

func pause() {

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

func crudMenu() {
	var menus [5]string
	var c int

	menus[0] = "Tambah Produk"
	menus[1] = "Perbarui Produk"
	menus[2] = "Hapus Produk"
	menus[3] = "Lihat Produk"
	menus[4] = "Kembali"

	for {
		showMenu("CRUD MENU", menus, 5)

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
			fmt.Println("Menu ini tidak tersedia")
		}
	}
}

func mainMenu() {
	var menus [5]string
	var c int

	menus[0] = "Kelola"
	menus[1] = "Lihat Produk"
	menus[2] = "Rekomendasi"
	menus[3] = "Sales"

	for {
		showMenu("MAIN MENU", menus, 4)

		c = inputInt("Pilih menu: ")

		switch c {
		case 1:
			crudMenu()

		case 2:
			fmt.Println("LIHAT PRODUCT")

		case 3:
			fmt.Println("RECOMMENDATION")

		case 4:
			fmt.Println("SALES")

		case 0:
			fmt.Println("Program finished")
			return

		default:
			fmt.Println("Menu not available")
		}
	}
}

func searchMenu(A *ProductArray, n *int) {

}

func sortMenu(A *ProductArray, n *int) {

}

func salesMenu(A *ProductArray, n *int) {

}

func recommendationMenu(A *ProductArray, n *int) {

}

func statisticMenu(A *ProductArray, n *int) {

}

// ======================================================
// CRUD
// ======================================================

func createProduct(A *ProductArray, n *int) {

}

func viewProduct(A ProductArray, n int) {
	var menus [5]string
	var c, s, i int

	menus[0] = "Semua Data"
	menus[1] = "Data Spesifik"
	menus[2] = "Data Detail"
	menus[3] = "Kembali"

	for { 
		showMenu("View Menu", menus, 3)
		if len(A) == 0 {
			fmt.Println("Belum ada data")
			return
		} else {
			c = inputInt("Pilih menu: ")
			
			switch c {
				case 1: 
					fmt.Printf("%-5s | %-20s | %-20s | %-20s | %-10s |\n", "ID", "Nama", "Kategori", "Harga", "Terjual")
					for i := 0; i < n; i++ {
						fmt.Printf("%-5d | %-20s | %-20s | %-20d | %-10d |\n", A[i].ID, A[i].Name, A[i].Category, A[i].Price, A[i].Sold)
					}
				
				case 2:
					s = inputInt("Masukkan ID data yang ingin dilihat: ")
					i = binarySearchID(A, n, s)

					if i == -1 {
						fmt.Println("Barang tidak ditemukan")
						return
					} else {
						fmt.Printf("%-5d | %-20s | %-20s | %-20d | %-10d |\n", A[i].ID, A[i].Name, A[i].Category, A[i].Price, A[i].Sold)
					}
				case 3: 
					viewProductDetail(A, n)
				case 4: 
					return
				default:
					fmt.Println("Menu tidak ditemukan")
			}

			
		}
	}

	
}

func viewProductDetail(A ProductArray, n int) {

}

func updateProduct(A *ProductArray, n int) {

}

func deleteProduct(A *ProductArray, n *int) {

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

func validatePrice(price int) bool {
	return false
}

func validateStock(stock int) bool {
	return false
}

// ======================================================
// SEARCHING
// ======================================================

// Sequential Search
func sequentialSearchName(A ProductArray, n int, name string) int {
	return -1
}

func sequentialSearchCategory(A ProductArray, n int, category string) int {
	return -1
}

// Binary Search
func binarySearchID(A ProductArray, n int, id int) int {
	var left, right, mid int

	left = 0
	right = n

	for left <= right {
		mid = (left + right) / 2 

		if A[mid].ID == id {
			return mid
		} else if A[mid].ID < id {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}

	return -1
}

func binarySearchPrice(A ProductArray, n int, price int) int {
	var left, right, mid int

	left = 0
	right = n

	for left <= right {
		mid = (left + right) / 2 

		if A[mid].Price == price {
			return mid
		} else if A[mid].Price < price {
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

// Insertion Sort
func insertionSortPriceAsc(A *ProductArray, n int) {
	
}

func insertionSortPriceDesc(A *ProductArray, n int) {

}

// Selection Sort
func selectionSortNameAsc(A *ProductArray, n int) {

}

func selectionSortNameDesc(A *ProductArray, n int) {

}

// ======================================================
// SALES
// ======================================================

func productSales(A *ProductArray, n int) {

}

func checkoutProduct(A *ProductArray, n int) {

}

func calculateTotalShopping(A ProductArray, n int) int {
	return 0
}

// ======================================================
// RECOMMENDATION
// ======================================================

func bestSellingRecommendation(A ProductArray, n int) {

}

func lowStockRecommendation(A ProductArray, n int) {

}

func menFashionRecommendation(A ProductArray, n int) {

}

func womenFashionRecommendation(A ProductArray, n int) {

}

// ======================================================
// STATISTIC
// ======================================================

func totalProduct(A ProductArray, n int) int {
	return 0
}

func totalStock(A ProductArray, n int) int {
	return 0
}

func totalSales(A ProductArray, n int) int {
	return 0
}

func mostSoldProduct(A ProductArray, n int) {

}

func leastSoldProduct(A ProductArray, n int) {

}

// ======================================================
// HELPER
// ======================================================

func findIndexByID(A ProductArray, n int, id int) int {
	return -1
}

func generateID(A ProductArray, n int) int {
	return 0
}

func swap(A *Product, B *Product) {

}

func printTableHeader() {

}

func printProduct(B Product) {

}

// ======================================================
// MAIN
// ======================================================

func main() {
	logo()
	mainMenu()
}