# Aplikasi Manajemen Produk & Penjualan - Spesifikasi Lengkap

##  Daftar Isi
1. [Deskripsi Aplikasi](#deskripsi-aplikasi)
2. [Struktur Data](#struktur-data)
3. [Fitur Aplikasi](#fitur-aplikasi)
4. [Contoh JSON](#contoh-json)
5. [Endpoint / Fungsi Utama](#endpoint--fungsi-utama)
6. [Flow Aplikasi](#flow-aplikasi)

---

## Deskripsi Aplikasi

Aplikasi ini adalah sistem manajemen inventori dan penjualan produk fashion/kosmetik dengan fitur lengkap meliputi:

- **CRUD Produk** (Create, Read, Update, Delete)
- **Manajemen Varian** (warna, ukuran, stok)
- **Sistem Review & Rating**
- **Transaksi Penjualan**
- **Rekomendasi Produk** (terlaris, stok menipis, kategori tertentu)
- **Statistik Penjualan**
- **Searching & Sorting** (berbagai kriteria)

### Batasan Sistem

| Komponen | Maksimal |
|----------|----------|
| Total Produk | 999 |
| Varian per produk | 5 |
| Review per produk | 10 |
| Item per transaksi | 10 |

---

## Struktur Data

### 1. Brand

```go
type Brand struct {
    Name    string  // Nama brand
    Country string  // Negara asal
}
```

### 2. Rating

```go
type Rating struct {
    Score       float64  // Rata-rata rating (1-5)
    TotalReview int      // Total review
}
```

### 3. Review

```go
type Review struct {
    Username string  // Nama pengulas
    Comment  string  // Isi komentar
    Value    int     // Rating (1-5)
}
```

### 4. ProductDetail

```go
type ProductDetail struct {
    Description string  // Deskripsi produk
    SkinType    string  // Jenis kulit (dry/oily/normal/combination)
    ExpiredYear int     // Tahun kadaluarsa
}
```

### 5. Variant

```go
type Variant struct {
    Color string  // Warna
    Size  string  // Ukuran (S/M/L/XL)
    Stock int     // Jumlah stok
}
```

### 6. Product (Inti)

```go
type Product struct {
    ID           int            // ID unik produk
    Name         string         // Nama produk
    Category     string         // Kategori (fashion_pria, fashion_wanita, skincare, makeup)
    Price        int            // Harga (Rupiah)
    Sold         int            // Jumlah terjual
    
    BrandInfo    Brand          // Informasi brand
    RateInfo     Rating         // Informasi rating
    DetailInfo   ProductDetail  // Detail produk
    
    Variants     [MAX_VARIANT]Variant  // Array varian
    VariantCount int                   // Jumlah varian aktif
    
    Reviews      [MAX_REVIEW]Review    // Array review
    ReviewCount  int                   // Jumlah review
}
```

### 7. Customer

```go
type Customer struct {
    ID    int    // ID pelanggan
    Name  string // Nama pelanggan
    Phone string // No telepon
}
```

### 8. SalesItem

```go
type SalesItem struct {
    ProductID   int    // ID produk
    ProductName string // Nama produk
    Price       int    // Harga satuan
    Quantity    int    // Jumlah beli
    Subtotal    int    // Subtotal (Price * Quantity)
}
```

### 9. Transaction

```go
type Transaction struct {
    TransactionID int           // ID transaksi
    Buyer         Customer      // Data pembeli
    Items         [MAX_ITEM]SalesItem  // Item yang dibeli
    ItemCount     int           // Jumlah item unik
    TotalPayment  int           // Total pembayaran
}
```

### Array Container

```go
type ProductList [NMAX]Product          // Array untuk semua produk
type TransactionList [NMAX]Transaction  // Array untuk semua transaksi
type arrProduct [NMAX]Product           // Alias untuk kemudahan
```

---

## Fitur Aplikasi

### 🔧 Menu CRUD (Pengelola)

| Fungsi | Deskripsi |
|--------|-----------|
| `createBarang()` | Menambah produk baru dengan semua detailnya |
| `updateBarang()` | Mengedit data produk (nama, harga, stok, dll) |
| `deleteBarang()` | Menghapus produk berdasarkan ID |
| `viewBarang()` | Menampilkan semua produk |
| `viewDetailBarang()` | Menampilkan detail lengkap 1 produk |

### 🔍 Menu View Barang

| Fungsi | Deskripsi |
|--------|-----------|
| `sequentialSearchNama()` | Cari produk berdasarkan nama (sequential) |
| `sequentialSearchKategori()` | Cari produk berdasarkan kategori (sequential) |
| `binarySearchID()` | Cari produk berdasarkan ID (binary) |
| `binarySearchHarga()` | Cari produk berdasarkan harga (binary) |

### 📊 Menu Sorting

| Fungsi | Deskripsi |
|--------|-----------|
| `insertionSortHargaAsc()` | Urutkan harga dari termurah (insertion sort) |
| `insertionSortHargaDesc()` | Urutkan harga dari termahal (insertion sort) |
| `selectionSortNamaAsc()` | Urutkan nama A-Z (selection sort) |
| `selectionSortNamaDesc()` | Urutkan nama Z-A (selection sort) |

### 💰 Menu Penjualan

| Fungsi | Deskripsi |
|--------|-----------|
| `penjualanBarang()` | Proses transaksi pembelian |
| `checkoutBarang()` | Menyelesaikan transaksi dan update stok |
| `hitungTotalBelanja()` | Menghitung total pembayaran |

### ⭐ Menu Rekomendasi

| Fungsi | Deskripsi |
|--------|-----------|
| `rekomendasiTerlaris()` | 5 produk dengan sold terbanyak |
| `rekomendasiStokSedikit()` | Produk dengan stok terendah |
| `rekomendasiFashionPria()` | Semua produk fashion pria |
| `rekomendasiFashionWanita()` | Semua produk fashion wanita |

### 📈 Menu Statistik

| Fungsi | Deskripsi |
|--------|-----------|
| `totalBarang()` | Total jumlah produk |
| `totalStok()` | Total stok semua produk |
| `totalPenjualan()` | Total unit terjual |
| `barangPalingLaris()` | Produk dengan sold tertinggi |
| `barangPalingSedikitTerjual()` | Produk dengan sold terendah |

### 💾 File Handling (Optional)

| Fungsi | Deskripsi |
|--------|-----------|
| `saveToFile()` | Simpan data produk ke file JSON |
| `loadFromFile()` | Load data produk dari file JSON |

---

## Contoh JSON

### 1. Contoh Single Product (JSON)

```json
{
  "ID": 1001,
  "Name": "Lip Cream Matte",
  "Category": "makeup",
  "Price": 89000,
  "Sold": 234,
  "BrandInfo": {
    "Name": "Wardah",
    "Country": "Indonesia"
  },
  "RateInfo": {
    "Score": 4.7,
    "TotalReview": 128
  },
  "DetailInfo": {
    "Description": "Lip cream matte finish dengan formula ringan dan tahan lama",
    "SkinType": "all",
    "ExpiredYear": 2026
  },
  "Variants": [
    {
      "Color": "Red Velvet",
      "Size": "5ml",
      "Stock": 45
    },
    {
      "Color": "Rose Pink",
      "Size": "5ml",
      "Stock": 32
    },
    {
      "Color": "Mauve",
      "Size": "5ml",
      "Stock": 28
    }
  ],
  "VariantCount": 3,
  "Reviews": [
    {
      "Username": "@beautygirl",
      "Comment": "Warnanya pigmented banget!",
      "Value": 5
    },
    {
      "Username": "@makeupaddict",
      "Comment": "Matte tapi tidak kering",
      "Value": 4
    }
  ],
  "ReviewCount": 2
}
```

### 2. Contoh Multiple Products (JSON)

```json
{
  "products": [
    {
      "ID": 1001,
      "Name": "Lip Cream Matte",
      "Category": "makeup",
      "Price": 89000,
      "Sold": 234,
      "BrandInfo": {
        "Name": "Wardah",
        "Country": "Indonesia"
      },
      "RateInfo": {
        "Score": 4.7,
        "TotalReview": 128
      },
      "DetailInfo": {
        "Description": "Lip cream matte finish dengan formula ringan",
        "SkinType": "all",
        "ExpiredYear": 2026
      },
      "Variants": [
        {
          "Color": "Red Velvet",
          "Size": "5ml",
          "Stock": 45
        }
      ],
      "VariantCount": 1,
      "Reviews": [],
      "ReviewCount": 0
    },
    {
      "ID": 1002,
      "Name": "Kemeja Flanel Pria",
      "Category": "fashion_pria",
      "Price": 149000,
      "Sold": 89,
      "BrandInfo": {
        "Name": "Uniqlo",
        "Country": "Japan"
      },
      "RateInfo": {
        "Score": 4.5,
        "TotalReview": 56
      },
      "DetailInfo": {
        "Description": "Kemeja flanel bahan nyaman dan adem",
        "SkinType": "",
        "ExpiredYear": 0
      },
      "Variants": [
        {
          "Color": "Red Plaid",
          "Size": "S",
          "Stock": 12
        },
        {
          "Color": "Red Plaid",
          "Size": "M",
          "Stock": 8
        },
        {
          "Color": "Red Plaid",
          "Size": "L",
          "Stock": 5
        },
        {
          "Color": "Blue Plaid",
          "Size": "M",
          "Stock": 10
        }
      ],
      "VariantCount": 4,
      "Reviews": [
        {
          "Username": "@fashionpria",
          "Comment": "Bagus dan nyaman dipakai",
          "Value": 5
        }
      ],
      "ReviewCount": 1
    }
  ],
  "totalProducts": 2,
  "lastUpdated": "2026-01-15T10:30:00Z"
}
```

### 3. Contoh Transaction (JSON)

```json
{
  "TransactionID": 5001,
  "Buyer": {
    "ID": 8001,
    "Name": "Andi Wijaya",
    "Phone": "081234567890"
  },
  "Items": [
    {
      "ProductID": 1002,
      "ProductName": "Kemeja Flanel Pria",
      "Price": 149000,
      "Quantity": 2,
      "Subtotal": 298000
    },
    {
      "ProductID": 1005,
      "ProductName": "Jeans Pria",
      "Price": 199000,
      "Quantity": 1,
      "Subtotal": 199000
    }
  ],
  "ItemCount": 2,
  "TotalPayment": 497000
}
```

### 4. Contoh File Konfigurasi (config.json)

```json
{
  "appName": "BeautyFashion Shop Manager",
  "version": "1.0.0",
  "maxProducts": 999,
  "maxVariants": 5,
  "maxReviews": 10,
  "categories": [
    "fashion_pria",
    "fashion_wanita",
    "skincare",
    "makeup"
  ],
  "skinTypes": [
    "dry",
    "oily",
    "normal",
    "combination",
    "all"
  ],
  "filePaths": {
    "products": "./data/products.json",
    "transactions": "./data/transactions.json",
    "backup": "./backup/"
  }
}
```

---

## Endpoint / Fungsi Utama

### Fungsi Validasi

```go
func isEmpty(n int) bool {
    return n == 0
}

func isFull(n int) bool {
    return n == NMAX
}

func validasiHarga(harga int) bool {
    return harga > 0
}

func validasiStok(stok int) bool {
    return stok >= 0
}
```

### Fungsi Helper

```go
func cariIndexByID(A arrProduct, n int, id int) int {
    for i := 0; i < n; i++ {
        if A[i].ID == id {
            return i
        }
    }
    return -1
}

func generateID(A arrProduct, n int) int {
    if n == 0 {
        return 1
    }
    maxID := A[0].ID
    for i := 1; i < n; i++ {
        if A[i].ID > maxID {
            maxID = A[i].ID
        }
    }
    return maxID + 1
}

func swap(A *Product, B *Product) {
    temp := *A
    *A = *B
    *B = temp
}
```

### Algoritma Sorting (Implementasi)

**Insertion Sort (Harga Ascending)**

```go
func insertionSortHargaAsc(A *arrProduct, n int) {
    for i := 1; i < n; i++ {
        key := A[i]
        j := i - 1
        for j >= 0 && A[j].Price > key.Price {
            A[j+1] = A[j]
            j--
        }
        A[j+1] = key
    }
}
```

**Insertion Sort (Harga Descending)**

```go
func insertionSortHargaDesc(A *arrProduct, n int) {
    for i := 1; i < n; i++ {
        key := A[i]
        j := i - 1
        for j >= 0 && A[j].Price < key.Price {
            A[j+1] = A[j]
            j--
        }
        A[j+1] = key
    }
}
```

**Selection Sort (Nama Ascending)**

```go
func selectionSortNamaAsc(A *arrProduct, n int) {
    for i := 0; i < n-1; i++ {
        minIdx := i
        for j := i + 1; j < n; j++ {
            if A[j].Name < A[minIdx].Name {
                minIdx = j
            }
        }
        if minIdx != i {
            swap(&A[i], &A[minIdx])
        }
    }
}
```

**Selection Sort (Nama Descending)**

```go
func selectionSortNamaDesc(A *arrProduct, n int) {
    for i := 0; i < n-1; i++ {
        maxIdx := i
        for j := i + 1; j < n; j++ {
            if A[j].Name > A[maxIdx].Name {
                maxIdx = j
            }
        }
        if maxIdx != i {
            swap(&A[i], &A[maxIdx])
        }
    }
}
```

### Algoritma Searching (Implementasi)

**Sequential Search by Name**

```go
func sequentialSearchNama(A arrProduct, n int, nama string) int {
    for i := 0; i < n; i++ {
        if A[i].Name == nama {
            return i
        }
    }
    return -1
}
```

**Sequential Search by Category**

```go
func sequentialSearchKategori(A arrProduct, n int, kategori string) int {
    for i := 0; i < n; i++ {
        if A[i].Category == kategori {
            return i
        }
    }
    return -1
}
```

**Binary Search by ID** (asumsi array sudah terurut ID)

```go
func binarySearchID(A arrProduct, n int, id int) int {
    left, right := 0, n-1
    for left <= right {
        mid := (left + right) / 2
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
```

**Binary Search by Price** (asumsi array sudah terurut harga)

```go
func binarySearchHarga(A arrProduct, n int, harga int) int {
    left, right := 0, n-1
    for left <= right {
        mid := (left + right) / 2
        if A[mid].Price == harga {
            return mid
        } else if A[mid].Price < harga {
            left = mid + 1
        } else {
            right = mid - 1
        }
    }
    return -1
}
```

---

## Flow Aplikasi

```
START
  │
  ▼
┌─────────────────┐
│   tampilkanLogo()│
└────────┬────────┘
         │
         ▼
┌─────────────────────────────────────┐
│           MENU UTAMA                │
│  ┌─────────────────────────────┐    │
│  │ 1. Pengelola (CRUD)         │    │
│  │ 2. View Barang (Search/Sort)│    │
│  │ 3. Rekomendasi              │    │
│  │ 4. Penjualan                │    │
│  │ 0. Exit                     │    │
│  └─────────────────────────────┘    │
└────────┬────────────────────────────┘
         │
    ┌────┴────────────────────────────────────┐
    │                                         │
    ▼                                         ▼
┌───────────────────┐                 ┌───────────────────┐
│    MENU CRUD      │                 │    VIEW MODE      │
│ ┌───────────────┐ │                 │ ┌───────────────┐ │
│ │ 1. Create     │ │                 │ │ - Search      │ │
│ │ 2. Update     │ │                 │ │ - Sort        │ │
│ │ 3. Delete     │ │                 │ │ - Detail      │ │
│ │ 4. View All   │ │                 │ └───────────────┘ │
│ │ 5. Back       │ │                 └───────────────────┘
│ └───────────────┘ │
└─────────┬─────────┘
          │
    ┌─────┴─────────────────┐
    │                       │
    ▼                       ▼
┌──────────────┐     ┌──────────────┐
│ REKOMENDASI  │     │  PENJUALAN   │
│ ┌──────────┐ │     │ ┌──────────┐ │
│ │Terlaris  │ │     │ │Pilih Item│ │
│ │Stok Tipis│ │     │ │Checkout  │ │
│ │Pria      │ │     │ │Update Stok│ │
│ │Wanita    │ │     │ └──────────┘ │
│ └──────────┘ │     └──────┬───────┘
└──────┬───────┘            │
       │                    │
       └────────┬───────────┘
                ▼
       ┌────────────────┐
       │   STATISTIK    │
       │ ┌────────────┐ │
       │ │Total Barang│ │
       │ │Total Stok  │ │
       │ │Top Sales   │ │
       │ │Low Sales   │ │
       │ └────────────┘ │
       └────────┬───────┘
                │
                ▼
               END
```

---

## Contoh Penggunaan Fungsi CRUD

### Create Product

```go
func createBarang(A *arrProduct, n *int) {
    if isFull(*n) {
        fmt.Println("Database produk penuh!")
        return
    }
    
    var p Product
    p.ID = generateID(*A, *n)
    
    fmt.Print("Nama produk: ")
    fmt.Scan(&p.Name)
    fmt.Print("Kategori (fashion_pria/fashion_wanita/skincare/makeup): ")
    fmt.Scan(&p.Category)
    fmt.Print("Harga: Rp ")
    fmt.Scan(&p.Price)
    
    if !validasiHarga(p.Price) {
        fmt.Println("Harga tidak valid!")
        return
    }
    
    // Input brand
    fmt.Print("Brand: ")
    fmt.Scan(&p.BrandInfo.Name)
    fmt.Print("Negara asal: ")
    fmt.Scan(&p.BrandInfo.Country)
    
    // Input detail
    fmt.Print("Deskripsi: ")
    fmt.Scan(&p.DetailInfo.Description)
    fmt.Print("Jenis kulit (dry/oily/normal/combination/all): ")
    fmt.Scan(&p.DetailInfo.SkinType)
    fmt.Print("Tahun kadaluarsa: ")
    fmt.Scan(&p.DetailInfo.ExpiredYear)
    
    // Input varian
    fmt.Print("Jumlah varian (max 5): ")
    fmt.Scan(&p.VariantCount)
    if p.VariantCount > MAX_VARIANT {
        p.VariantCount = MAX_VARIANT
    }
    
    for i := 0; i < p.VariantCount; i++ {
        fmt.Printf("\n=== Varian %d ===\n", i+1)
        fmt.Print("Warna: ")
        fmt.Scan(&p.Variants[i].Color)
        fmt.Print("Ukuran (S/M/L/XL): ")
        fmt.Scan(&p.Variants[i].Size)
        fmt.Print("Stok: ")
        fmt.Scan(&p.Variants[i].Stock)
    }
    
    // Initialize rating
    p.RateInfo.Score = 0
    p.RateInfo.TotalReview = 0
    
    A[*n] = p
    *n++
    fmt.Println("\n✅ Produk berhasil ditambahkan!")
}
```

### Update Product

```go
func updateBarang(A *arrProduct, n int) {
    if isEmpty(n) {
        fmt.Println("Belum ada produk!")
        return
    }
    
    id := inputInt("Masukkan ID produk yang akan diupdate: ")
    idx := cariIndexByID(*A, n, id)
    
    if idx == -1 {
        fmt.Println("❌ Produk tidak ditemukan!")
        return
    }
    
    fmt.Println("\n📝 Data Lama:")
    fmt.Printf("Nama: %s\n", A[idx].Name)
    fmt.Printf("Harga: Rp %d\n", A[idx].Price)
    fmt.Printf("Kategori: %s\n", A[idx].Category)
    
    fmt.Println("\n📝 Masukkan Data Baru:")
    fmt.Print("Nama baru: ")
    fmt.Scan(&A[idx].Name)
    fmt.Print("Harga baru: Rp ")
    fmt.Scan(&A[idx].Price)
    fmt.Print("Kategori baru: ")
    fmt.Scan(&A[idx].Category)
    
    fmt.Println("✅ Produk berhasil diupdate!")
}
```

### Delete Product

```go
func deleteBarang(A *arrProduct, n *int) {
    if isEmpty(*n) {
        fmt.Println("Belum ada produk!")
        return
    }
    
    id := inputInt("Masukkan ID produk yang akan dihapus: ")
    idx := cariIndexByID(*A, *n, id)
    
    if idx == -1 {
        fmt.Println("❌ Produk tidak ditemukan!")
        return
    }
    
    fmt.Printf("Apakah anda yakin ingin menghapus %s? (y/n): ", A[idx].Name)
    var confirm string
    fmt.Scan(&confirm)
    
    if confirm != "y" && confirm != "Y" {
        fmt.Println("Penghapusan dibatalkan.")
        return
    }
    
    // Shift elemen ke kiri
    for i := idx; i < *n-1; i++ {
        A[i] = A[i+1]
    }
    *n--
    fmt.Println("✅ Produk berhasil dihapus!")
}
```

### View All Products

```go
func viewBarang(A arrProduct, n int) {
    if isEmpty(n) {
        fmt.Println("📭 Belum ada produk dalam database.")
        return
    }
    
    fmt.Println("\n" + strings.Repeat("=", 80))
    fmt.Printf("%-6s | %-25s | %-15s | %-10s | %-8s | %-6s\n", 
        "ID", "Nama Produk", "Kategori", "Harga", "Terjual", "Rating")
    fmt.Println(strings.Repeat("-", 80))
    
    for i := 0; i < n; i++ {
        fmt.Printf("%-6d | %-25s | %-15s | Rp %-7d | %-8d | %.1f\n",
            A[i].ID,
            truncateString(A[i].Name, 25),
            A[i].Category,
            A[i].Price,
            A[i].Sold,
            A[i].RateInfo.Score)
    }
    fmt.Println(strings.Repeat("=", 80))
    fmt.Printf("Total produk: %d\n", n)
}
```

---

## Contoh Skenario Transaksi

```go
func penjualanBarang(A *arrProduct, n int) {
    if isEmpty(n) {
        fmt.Println("Belum ada produk untuk dijual!")
        return
    }
    
    var trans Transaction
    trans.TransactionID = generateTransactionID()
    
    fmt.Println("\n💰 FORM PEMBELIAN")
    fmt.Println(strings.Repeat("-", 40))
    fmt.Print("Nama pembeli: ")
    fmt.Scan(&trans.Buyer.Name)
    fmt.Print("No telepon: ")
    fmt.Scan(&trans.Buyer.Phone)
    
    for {
        fmt.Println("\n🛍️ TAMBAH ITEM")
        fmt.Println(strings.Repeat("-", 40))
        viewBarang(*A, n)
        
        id := inputInt("\nMasukkan ID produk (0 selesai): ")
        if id == 0 {
            break
        }
        
        idx := cariIndexByID(*A, n, id)
        if idx == -1 {
            fmt.Println("❌ Produk tidak ditemukan!")
            continue
        }
        
        // Tampilkan varian produk
        fmt.Printf("\n📦 Produk: %s\n", A[idx].Name)
        fmt.Printf("Harga: Rp %d\n", A[idx].Price)
        fmt.Println("\nVarian yang tersedia:")
        
        for i := 0; i < A[idx].VariantCount; i++ {
            fmt.Printf("%d. %s - %s (Stok: %d)\n", 
                i+1, 
                A[idx].Variants[i].Color, 
                A[idx].Variants[i].Size, 
                A[idx].Variants[i].Stock)
        }
        
        var variantIdx int
        fmt.Print("Pilih varian (1-5): ")
        fmt.Scan(&variantIdx)
        variantIdx--
        
        if variantIdx < 0 || variantIdx >= A[idx].VariantCount {
            fmt.Println("❌ Varian tidak valid!")
            continue
        }
        
        qty := inputInt("Jumlah beli: ")
        if qty <= 0 {
            fmt.Println("❌ Jumlah tidak valid!")
            continue
        }
        
        if qty > A[idx].Variants[variantIdx].Stock {
            fmt.Printf("❌ Stok tidak mencukupi! Tersisa: %d\n", A[idx].Variants[variantIdx].Stock)
            continue
        }
        
        subtotal := A[idx].Price * qty
        
        // Kurangi stok
        A[idx].Variants[variantIdx].Stock -= qty
        A[idx].Sold += qty
        
        // Tambah ke transaksi
        trans.Items[trans.ItemCount] = SalesItem{
            ProductID:   A[idx].ID,
            ProductName: A[idx].Name,
            Price:       A[idx].Price,
            Quantity:    qty,
            Subtotal:    subtotal,
        }
        trans.ItemCount++
        trans.TotalPayment += subtotal
        
        fmt.Printf("✅ Item ditambahkan! Subtotal: Rp %d\n", subtotal)
    }
    
    if trans.ItemCount == 0 {
        fmt.Println("Tidak ada item yang dibeli.")
        return
    }
    
    // Tampilkan struk
    fmt.Println("\n" + strings.Repeat("=", 50))
    fmt.Println("           STRUK PEMBELIAN")
    fmt.Println(strings.Repeat("=", 50))
    fmt.Printf("ID Transaksi : %d\n", trans.TransactionID)
    fmt.Printf("Tanggal      : %s\n", time.Now().Format("02/01/2006 15:04:05"))
    fmt.Printf("Pembeli      : %s (%s)\n", trans.Buyer.Name, trans.Buyer.Phone)
    fmt.Println(strings.Repeat("-", 50))
    fmt.Println("Detail Pembelian:")
    
    for i := 0; i < trans.ItemCount; i++ {
        fmt.Printf("  %d. %s\n", i+1, trans.Items[i].ProductName)
        fmt.Printf("     %d x Rp %d = Rp %d\n", 
            trans.Items[i].Quantity,
            trans.Items[i].Price,
            trans.Items[i].Subtotal)
    }
    
    fmt.Println(strings.Repeat("-", 50))
    fmt.Printf("TOTAL BAYAR : Rp %d\n", trans.TotalPayment)
    fmt.Println(strings.Repeat("=", 50))
    fmt.Println("Terima kasih sudah berbelanja! 💖")
    fmt.Println(strings.Repeat("=", 50))
}
```

---

## Catatan Implementasi

1. **Array vs Slice**: Gunakan array statis dengan konstanta NMAX untuk memenuhi batasan

2. **Case Sensitive**: Search menggunakan pencocokan eksak (case-sensitive)

3. **Binary Search**: Pastikan array sudah terurut sebelum menggunakan binary search

4. **Rating Score**: Hitung rata-rata dari semua review setiap kali ada review baru
   ```go
   func updateRating(p *Product) {
       total := 0
       for i := 0; i < p.ReviewCount; i++ {
           total += p.Reviews[i].Value
       }
       if p.ReviewCount > 0 {
           p.RateInfo.Score = float64(total) / float64(p.ReviewCount)
       }
       p.RateInfo.TotalReview = p.ReviewCount
   }
   ```

5. **Stok Varian**: Saat checkout, kurangi stok varian yang dipilih

6. **ID Generation**: ID baru = ID maksimum + 1

7. **Validasi**: Selalu validasi input sebelum memproses data

---

## Contoh Output Program

```
=======================================================================
|                        Tetap Slay Maksimal                          |
=======================================================================

MENU UTAMA
========================================
1. Pengelola (CRUD)
2. View Barang (Search/Sort)
3. Rekomendasi
4. Penjualan
0. Exit
========================================
Pilih menu: 1

MENU CRUD
========================================
1. Create Barang
2. Update Barang
3. Delete Barang
4. View Barang
5. Back
========================================
Pilih menu: 4

VIEW BARANG
================================================================================
ID    | Nama Produk               | Kategori        | Harga      | Terjual | Rating
--------------------------------------------------------------------------------
1001  | Lip Cream Matte           | makeup          | Rp 89000   | 234     | 4.7
1002  | Kemeja Flanel Pria        | fashion_pria    | Rp 149000  | 89      | 4.5
1003  | Moisturizer               | skincare        | Rp 125000  | 567     | 4.9
================================================================================
Total produk: 3

Menu:
1. Detail produk
2. Cari produk
3. Urutkan
4. Kembali
Pilih: 
```

---

## Kesimpulan

Aplikasi ini menyediakan sistem manajemen toko fashion dan kosmetik yang lengkap dengan fitur:

- ✅ Manajemen produk dengan varian
- ✅ Sistem review dan rating
- ✅ Transaksi penjualan
- ✅ Rekomendasi produk
- ✅ Statistik penjualan
- ✅ Searching dan sorting
- ✅ Export/Import JSON

Aplikasi siap diimplementasikan!