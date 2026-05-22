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

	fmt.Println("0. Exit")
	line(40)
}

func crudMenu() {
	logo()
	var menus [5]string
	var c int

	menus[0] = "Create Product"
	menus[1] = "Update Product"
	menus[2] = "Delete Product"
	menus[3] = "View Product"
	menus[4] = "Back"

	for {
		showMenu("CRUD MENU", menus, 5)

		c = inputInt("Choose menu: ")

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
			clearScreen()

		case 0:
			return

		default:
			fmt.Println("Menu not available")
		}
	}
}

func mainMenu() {
	var menus [5]string
	var c int

	menus[0] = "Manager"
	menus[1] = "View Product"
	menus[2] = "Recommendation"
	menus[3] = "Sales"

	for {
		showMenu("MAIN MENU", menus, 4)

		c = inputInt("Choose menu: ")
		d := 0
		fmt.Scan(&d)

		switch c {
		case 1:
			crudMenu()

		case 2:
			fmt.Println("VIEW PRODUCT")

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
	return -1
}

func binarySearchPrice(A ProductArray, n int, price int) int {
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
	//clearScreen()
	logo()
	mainMenu()
}
