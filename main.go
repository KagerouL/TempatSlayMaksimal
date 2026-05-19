package main

//https://patorjk.com/software/taag/#p=display&f=Graffiti&t=Type+Something+&x=none&v=4&h=4&w=80&we=false
import "fmt"

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

func main() {
	logo()
	menuUtama()
}
