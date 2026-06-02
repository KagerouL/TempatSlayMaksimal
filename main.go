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

type ProductList [NMAX]Product
type ProductArray [NMAX]Product

var productsArr ProductArray
var countData int
var historyCountData int

// ======================================================
// UTILITY
// ======================================================

func logo() {
	var pink, logo string
	pink = "\033[38;2;245;187;212m"
	logo = fmt.Sprintf(`
%s
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
func title(text string) {
	line(40)
	fmt.Printf("|%26s%-12s|\n", text, "")
	line(40)
}

func line(n int) {
	var i int
	for i = 0; i < n; i++ {
		fmt.Print("=")
	}
	fmt.Println()
}

func inputInt(prompt string) int {
	var value int
	var err error

	for {
		fmt.Print(prompt)
		_, err = fmt.Scan(&value)
		if err != nil {
			warningMessage("Input harus berupa angka!")
			continue
		}
		if value < 0 {
			warningMessage("Input tidak boleh negatif!")
			continue
		}
		return value
	}
}

func inputString(text string) string {
	var x string
	fmt.Print(text)
	fmt.Scan(&x)
	return x
}

func clearScreen() {
	fmt.Print("\033[H\033[2J")
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

	fmt.Println("0. Kembali")
	line(40)
}

func crudMenu() {
	logo()
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
			createProduct(&productsArr)

		case 2:
			fmt.Println("UPDATE")

		case 3:
			fmt.Println("DELETE")

		case 4:
			fmt.Println("VIEW")

		case 5:
			clearScreen()

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
			viewProduct(productsArr, countData)
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

func createProduct(data *ProductArray) {
	var i int

	logo()
	title("TAMBAH PRODUK")

	if countData >= NMAX {
		fmt.Println("Data produk penuh!")
		return
	}

	data[countData].ID = generateID(*data)

	data[countData].Name = inputString("Nama Produk        : ")
	data[countData].Category = inputString("Kategori           : ")
	data[countData].Price = inputInt("Harga              : ")

	data[countData].BrandInfo.Name = inputString("Nama Brand         : ")
	data[countData].BrandInfo.Country = inputString("Negara datasal Brand  : ")

	data[countData].DetailInfo.Description = inputString("Deskripsi Produk   : ")
	data[countData].DetailInfo.SkinType = inputString("Jenis Kulit        : ")
	data[countData].DetailInfo.ExpiredYear = inputInt("Tahun Kedaluwarsa  : ")

	data[countData].VariantCount = inputInt("Jumlah Varian (1-5): ")

	if data[countData].VariantCount < 1 {
		data[countData].VariantCount = 1
	}

	if data[countData].VariantCount > MAX_VARIANT {
		data[countData].VariantCount = MAX_VARIANT
	}

	for i = 0; i < data[countData].VariantCount; i++ {
		fmt.Println()
		fmt.Println("Varian", i+1)

		data[countData].Variants[i].Color = inputString("Warna  : ")
		data[countData].Variants[i].Size = inputString("Ukuran : ")
		data[countData].Variants[i].Stock = inputInt("Stok   : ")
	}

	data[countData].Sold = 0
	data[countData].RateInfo.Score = 0
	data[countData].RateInfo.TotalReview = 0
	data[countData].ReviewCount = 0

	countData++

	fmt.Println()
	fmt.Println("Produk berhasil ditambahkan!")
}

func viewProduct(A ProductArray, n int) {
	var menus [5]string
	var c, i int
	var s string

	menus[0] = "Semua Data"
	menus[1] = "Data Spesifik"
	menus[2] = "Data Detail"

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
				s = inputString("Masukkan ID data yang ingin dilihat: ")
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
			case 0:
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
	if price > 0 {
		return true
	}

	return false
}

func validateStock(stock int) bool {
	if stock >= 0 {
		return true
	}

	return false
}

// ======================================================
// SEARCHING
// ======================================================

// Sequential Search
func sequentialSearchName(A ProductArray, n int, name string) int {
	for i := 0; i < n; i++ {
		if A[i].Name == name {
			return i
		}
	}

	return -1
}

func sequentialSearchCategory(A ProductArray, n int, category string) int {
	for i := 0; i < n; i++ {
		if A[i].Category == category {
			return i
		}
	}

	return -1
}

// Binary Search
func binarySearchID(A ProductArray, n int, id string) int {
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
	var i, j int
	var temp Product

	i = 0

	for i <= n-1 {
		j = i
		temp = A[j]

		for j > 0 && temp.Price < A[j-1].Price {
			A[j] = A[j-1]
			j = j - 1
		}
		A[j] = temp
		i = i + 1
	}
}

func insertionSortPriceDesc(A *ProductArray, n int) {
	var i, j int
	var temp Product

	i = 0

	for i <= n-1 {
		j = i
		temp = A[j]

		for j > 0 && temp.Price > A[j-1].Price {
			A[j] = A[j-1]
			j = j - 1
		}
		A[j] = temp
		i = i + 1
	}
}

// Selection Sort
func selectionSortNameAsc(A *ProductArray, n int) {
	var idxMin, j, i int
	var temp Product

	for i <= n {
		idxMin = i
		j = i

		for j > n {
			if A[idxMin].Name > A[j].Name {
				idxMin = j
			}
			j++
		}
		temp = A[idxMin]
		A[idxMin] = A[i]
		A[i] = temp
		i++
	}
}

func selectionSortNameDesc(A *ProductArray, n int) {
	var idxMax, j, i int
	var temp Product

	for i <= n {
		idxMax = n
		j = i

		for j < n {
			if A[idxMax].Name < A[j].Name {
				idxMax = j
			}
			j++
		}
		temp = A[idxMax]
		A[idxMax] = A[i]
		A[i] = temp
		i++
	}
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

func errorMessage(message string) {
	title("ERROR >_<")
	fmt.Println(message)
	line(40)
}

func successMessage(message string) {
	title("SUCCESS <3")
	fmt.Println(message)
	line(40)
}

func warningMessage(message string) {
	title("WARNING -_-")
	fmt.Println(message)
	line(40)
}

func findIndexByID(A ProductArray, n int, id string) int {
	var mid, left, right, isFound int
	left = 0
	right = n
	isFound = -1
	for left < right && isFound == -1 {
		mid = (left + right) / 2
		if id == A[mid].ID {
			isFound = mid
		}
		if id > A[mid].ID {
			right = mid - 1
		} else if id < A[mid].ID {
			left = mid + 1
		}
	}
	return -1
}

func generateID(A ProductArray) string {
	return "TSM" + fmt.Sprintf("%03d", historyCountData+1)
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
	//clearScreen()
	logo()
	mainMenu()
}
