package main

import (
	"fmt"
	"os"
)

const NMAX int = 1000

type Kegiatan struct {
	ketua      string
	anggota    [4]string
	prodFak    string
	judul      string
	sumberDana string
	luaran     string
	tahun      int
}

type ArrKegiatan struct {
	tabKegiatan   [NMAX - 1]Kegiatan
	totalKegiatan int
	tipeKegiatan  string
}

type DataTahun struct {
	tahun                                                     int
	totalPublikasi, totalProduk, totalSeminar, totalPelatihan int
	sumKegiatan                                               int
}

type ArrDataTahun struct {
	data      [NMAX - 1]DataTahun
	tipe      string
	totalData int
}

func main() {
	var dataPenelitian, dataAbdimas ArrKegiatan
	dataPenelitian.tipeKegiatan = "penelitian"
	dataAbdimas.tipeKegiatan = "abdimas"

	var dataTahunPenelitian, dataTahunAbdimas ArrDataTahun
	dataTahunPenelitian.tipe = "penelitian"
	dataTahunAbdimas.tipe = "abdimas"

	for true {
		fmt.Println("=============\t MAIN MENU \t=============")
		fmt.Println("1. Kelola Data Penelitian")
		fmt.Println("2. Kelola Data Abdimas")
		fmt.Println("3. Sorting Data")
		fmt.Println("4. Tampilkan Data")
		fmt.Println("0. Keluar")

		var menu int
		fmt.Print("\nPilih Menu : ")
		fmt.Scan(&menu)
		fmt.Print("==================================================\n\n")

		switch menu {
		case 1:
			KelolaDataPenelitian(&dataPenelitian, &dataTahunPenelitian)
		case 2:
			KelolaDataAbdimas(&dataAbdimas, &dataTahunAbdimas)
		case 3:
			SortingData(&dataPenelitian, &dataAbdimas, &dataTahunPenelitian, &dataTahunAbdimas)
		case 4:
			TampilkanData(dataPenelitian, dataAbdimas, dataTahunPenelitian, dataTahunAbdimas)
		case 0:
			os.Exit(0)

		}
	}
}

func KelolaDataPenelitian(dataPenelitian *ArrKegiatan, dataTahunPenelitian *ArrDataTahun) {
	// var dataPenelitian ArrKegiatan
	dataPenelitian.tipeKegiatan = "penelitian"

	// var dataTahunPenelitian ArrDataTahun
	dataTahunPenelitian.tipe = "penelitian"

	for true {
		fmt.Println("========= Menu Pengolahan Data Penelitian =========")
		fmt.Println("1. Tambahkan Data Penelitian")
		fmt.Println("2. Edit Data Penelitian")
		fmt.Println("3. Hapus Data Penelitian")
		fmt.Println("0. Kembali")

		var pilihan int
		fmt.Print("\nPilih Menu : ")
		fmt.Scan(&pilihan)
		switch pilihan {
		case 1:
			InputDataKegiatan(dataPenelitian, dataTahunPenelitian)
		case 2:
			EditDataKegiatan(dataPenelitian)
		case 3:
			HapusDataKegiatan(dataPenelitian)
		case 0:
			return
		}
	}
}

func KelolaDataAbdimas(dataAbdimas *ArrKegiatan, dataTahunAbdimas *ArrDataTahun) {
	// var dataAbdimas ArrKegiatan
	dataAbdimas.tipeKegiatan = "abdimas"

	// var dataTahunAbdimas ArrDataTahun
	dataTahunAbdimas.tipe = "abdimas"

	for true {
		fmt.Println("===========\t Menu Pengolahan Data Abdimas \t===========")
		fmt.Println("1. Tambahkan Data Abdimas")
		fmt.Println("2. Edit Data Abdimas")
		fmt.Println("3. Hapus Data Abdimas")
		fmt.Println("0. Kembali")

		var pilihan int
		fmt.Print("\nPilih Menu : ")
		fmt.Scan(&pilihan)
		switch pilihan {
		case 1:
			InputDataKegiatan(dataAbdimas, dataTahunAbdimas)
		case 2:
			EditDataKegiatan(dataAbdimas)
		case 3:
			HapusDataKegiatan(dataAbdimas)
		case 0:
			return
		}
	}
}

func SortingData(dataPenelitian, dataAbdimas *ArrKegiatan, dataTahunPenelitian, dataTahunAbdimas *ArrDataTahun) {
	// var dataPenelitian, dataAbdimas ArrKegiatan
	dataPenelitian.tipeKegiatan = "penelitian"
	dataAbdimas.tipeKegiatan = "abdimas"

	// var dataTahunPenelitian, dataTahunAbdimas ArrDataTahun
	dataTahunPenelitian.tipe = "penelitian"
	dataTahunAbdimas.tipe = "abdimas"

	for true {
		fmt.Println("===========\t Menu Sorting Data \t===========")
		fmt.Println("1. Sorting Data Penelitian")
		fmt.Println("2. Sorting Data Abdimas")
		fmt.Println("0. Kembali")

		var pilihan int
		fmt.Print("\nPilih Menu : ")
		fmt.Scan(&pilihan)
		switch pilihan {
		case 1:
			fmt.Print("Sorting data penelitian berdasarkan: \n1. Tahun \n2. Jumlah Kegiatan Pertahun \n")
			var pilihan int
			fmt.Print("Pilih menu : ")
			fmt.Scan(&pilihan)

			switch pilihan {
			case 1:
				fmt.Print("Sorting data dari: \n1. Tahun Terlama \n2. Tahun Terbaru")
				var menu int
				fmt.Print("\nPilih Menu : ")
				fmt.Scan(&menu)

				switch menu {
				case 1:
					InSortData_Tahun(dataPenelitian)
					PrintData(*dataPenelitian)
				case 2:
					SelSortData_Tahun(dataPenelitian)
					PrintData(*dataPenelitian)
				}

			case 2:
				fmt.Print("Sorting data dari: \n1. Kegiatan paling sedikit \n2. Kegiatan paling banyak")
				var menu int
				fmt.Print("\nPilih Menu : ")
				fmt.Scan(&menu)

				switch menu {
				case 1:
					SortingJumlahKegiatan(dataTahunPenelitian, true)
					outputDataTahun(*dataTahunPenelitian)
				case 2:
					SortingJumlahKegiatan(dataTahunPenelitian, false)
					outputDataTahun(*dataTahunPenelitian)
				}
			}

		case 2:
			fmt.Print("Sorting data abdimas berdasarkan: \n1. Tahun \n2. Jumlah Kegiatan Pertahun \n")
			var pilihan int
			fmt.Print("Pilih menu : ")
			fmt.Scan(&pilihan)

			switch pilihan {
			case 1:
				fmt.Print("Sorting data dari: \n1. Tahun Terlama \n2. Tahun Terbaru")
				var menu int
				fmt.Print("\nPilih Menu : ")
				fmt.Scan(&menu)

				switch menu {
				case 1:
					InSortData_Tahun(dataAbdimas)
					PrintData(*dataAbdimas)
				case 2:
					SelSortData_Tahun(dataAbdimas)
					PrintData(*dataAbdimas)
				}

			case 2:
				fmt.Print("Sorting data dari: \n1. Kegiatan paling sedikit \n2. Kegiatan paling banyak")
				var menu int
				fmt.Print("\nPilih Menu : ")
				fmt.Scan(&menu)

				switch menu {
				case 1:
					SortingJumlahKegiatan(dataTahunAbdimas, true)
					outputDataTahun(*dataTahunAbdimas)
				case 2:
					SortingJumlahKegiatan(dataTahunAbdimas, false)
					outputDataTahun(*dataTahunAbdimas)
				}
			}

		case 0:
			return
		}
	}
}

func TampilkanData(dataPenelitian, dataAbdimas ArrKegiatan, dataTahunPenelitian, dataTahunAbdimas ArrDataTahun) {
	// var dataPenelitian, dataAbdimas ArrKegiatan
	dataPenelitian.tipeKegiatan = "penelitian"
	dataAbdimas.tipeKegiatan = "abdimas"

	// var dataTahunPenelitian, dataTahunAbdimas ArrDataTahun
	dataTahunPenelitian.tipe = "penelitian"
	dataTahunAbdimas.tipe = "abdimas"
	for true {
		fmt.Println("===========\t Menu Tampilkan Data \t===========")
		fmt.Println("1. Penelitian")
		fmt.Println("2. Abdimas")
		fmt.Println("0. Kembali")

		var menu int
		fmt.Print("\nPilih Menu : ")
		fmt.Scan(&menu)

		switch menu {
		case 1:
			fmt.Print("Tampilkan data berdasarkan: \n1. Data Keseluruhan \n2. Tahun \n3. Fakultas/Prodi \n4. Judul \n")

			var pilihan int
			fmt.Print("Pilih menu : ")
			fmt.Scan(&pilihan)

			switch pilihan {
			case 1:
				PrintData(dataPenelitian)
			case 2:
				SearchingData_Tahun(&dataPenelitian)
			case 3:
				SearchingData_FakProdi(&dataPenelitian)
			case 4:
				SortingJudul(&dataPenelitian)
				var index = SearchingData_Judul(dataPenelitian)
				var data Kegiatan = dataPenelitian.tabKegiatan[index]
				PrintDataKegiatan(data)
			}

		case 2:
			fmt.Print("Tampilkan data berdasarkan: \n1. Data Keseluruhan \n2. Tahun \n3. Fakultas/Prodi \n4. Judul \n")

			var pilihan int
			fmt.Print("Pilih menu : ")
			fmt.Scan(&pilihan)

			switch pilihan {
			case 1:
				PrintData(dataAbdimas)
			case 2:
				SearchingData_Tahun(&dataAbdimas)
			case 3:
				SearchingData_FakProdi(&dataAbdimas)
			case 4:
				SortingJudul(&dataAbdimas)
				var index = SearchingData_Judul(dataAbdimas)
				var data Kegiatan = dataAbdimas.tabKegiatan[index]
				PrintDataKegiatan(data)
			}

		case 0:
			return
		}
	}
}

// Kelola Data = Tambahkan, Edit, Hapus
func InputDataKegiatan(P *ArrKegiatan, T *ArrDataTahun) {
	var i int
	fmt.Print("Banyak data : ")
	fmt.Scan(&i)
	fmt.Println("---------------")
	for i > 0 {
		fmt.Print("Judul Kegiatan \t: ")
		fmt.Scan(&P.tabKegiatan[P.totalKegiatan].judul)
		fmt.Print("Ketua \t\t: ")
		fmt.Scan(&P.tabKegiatan[P.totalKegiatan].ketua)
		fmt.Println("Anggota \t: ")
		var ang int
		var anggota string
		fmt.Print("- ")
		fmt.Scan(&anggota)
		for ang < 4 && anggota != "xxx" {
			P.tabKegiatan[P.totalKegiatan].anggota[ang] = anggota
			ang++
			if ang != 4 {
				fmt.Print("- ")
				fmt.Scan(&anggota)
			}
		}
		fmt.Print("Prodi / Fakultas \t\t\t: ")
		fmt.Scan(&P.tabKegiatan[P.totalKegiatan].prodFak)
		fmt.Print("Sumber Dana (internal / eksternal) \t: ")
		fmt.Scan(&P.tabKegiatan[P.totalKegiatan].sumberDana)
		if P.tipeKegiatan == "penelitian" {
			fmt.Print("Luaran (publikasi / produk) \t\t: ")
			fmt.Scan(&P.tabKegiatan[P.totalKegiatan].luaran)
		} else {
			fmt.Print("Luaran (publikasi / produk / seminar / pelatihan) : ")
			fmt.Scan(&P.tabKegiatan[P.totalKegiatan].luaran)
		}
		fmt.Print("Tahun Kegiatan \t\t\t\t: ")
		fmt.Scan(&P.tabKegiatan[P.totalKegiatan].tahun)
		TambahDataTahun(T, P.tabKegiatan[P.totalKegiatan].tahun, P.tabKegiatan[P.totalKegiatan].luaran)
		P.totalKegiatan++
		fmt.Println("==================================================")
		i--
	}
}

func EditDataKegiatan(P *ArrKegiatan) {
	fmt.Print("Masukkan indeks data yang akan diubah : ")
	var index int
	fmt.Scan(&index)

	if index >= 0 && index < P.totalKegiatan {
		fmt.Print("Judul Kegiatan \t: ")
		fmt.Scan(&P.tabKegiatan[index].judul)
		fmt.Print("Ketua \t: ")
		fmt.Scan(&P.tabKegiatan[index].ketua)
		fmt.Println("Anggota \t: ")
		var ang int
		var anggota string
		fmt.Print("- ")
		fmt.Scan(&anggota)
		for ang < 4 && anggota != "xxx" {
			P.tabKegiatan[P.totalKegiatan].anggota[ang] = anggota
			ang++
			if ang != 4 {
				fmt.Print("- ")
				fmt.Scan(&anggota)
			}
		}
		fmt.Print("Prodi / Fakultas \t\t\t: ")
		fmt.Scan(&P.tabKegiatan[index].prodFak)
		fmt.Print("Sumber Dana (internal / eksternal) \t: ")
		fmt.Scan(&P.tabKegiatan[index].sumberDana)
		if P.tipeKegiatan == "penelitian" {
			fmt.Print("Luaran (publikasi / produk) \t\t: ")
			fmt.Scan(&P.tabKegiatan[index].luaran)
		} else {
			fmt.Print("Luaran (publikasi / produk / seminar / pelatihan) : ")
			fmt.Scan(&P.tabKegiatan[index].luaran)
		}
		fmt.Print("Tahun Kegiatan \t\t\t\t: ")
		fmt.Scan(&P.tabKegiatan[index].tahun)
		fmt.Println("Data berhasil diubah.")
		fmt.Println("==================================================")
	} else {
		fmt.Println("Indeks data tidak valid.")
		fmt.Println("==================================================")
	}
}

func HapusDataKegiatan(P *ArrKegiatan) {
	fmt.Print("Masukkan indeks data yang akan dihapus : ")
	var index int
	fmt.Scan(&index)

	if index >= 0 && index < P.totalKegiatan {
		for i := index; i < P.totalKegiatan-1; i++ {
			P.tabKegiatan[i] = P.tabKegiatan[i+1]
		}
		P.totalKegiatan--
		fmt.Println("Data berhasil dihapus.")
		fmt.Println("==================================================")
	} else {
		fmt.Println("Indeks data tidak valid.")
		fmt.Println("==================================================")
	}
}

// Sorting Data = Sorting berdasarkan tahun atau jumlah kegiatan pertahun (secara ascending/descending)
// Sorting berdasarkan tahun (Ascending)
func InSortData_Tahun(P *ArrKegiatan) {
	//Insertion Sort = Ascending
	var pass, i int
	pass = 1
	for pass < P.totalKegiatan {
		i = pass
		temp := P.tabKegiatan[pass]
		for i > 0 && temp.tahun < P.tabKegiatan[i-1].tahun {
			P.tabKegiatan[i].tahun = P.tabKegiatan[i-1].tahun
			i--
		}
		P.tabKegiatan[i] = temp
		pass++
	}
}

// Sorting berdasarkan tahun (Descending)
func SelSortData_Tahun(P *ArrKegiatan) {
	//Selection Sort = Descending
	var pass, idx, i int
	pass = 1
	for pass < P.totalKegiatan {
		idx = pass - 1
		i = pass
		for i < P.totalKegiatan {
			if P.tabKegiatan[idx].tahun < P.tabKegiatan[i].tahun {
				idx = i
			}
			i++
		}
		temp := P.tabKegiatan[pass-1].tahun
		P.tabKegiatan[pass-1].tahun = P.tabKegiatan[idx].tahun
		P.tabKegiatan[idx].tahun = temp
		pass++
	}
}

// Sorting berdasarkan jumlah kegiatan
func CariTahun(T ArrDataTahun, tahun int) int {
	for i := 0; i < T.totalData; i++ {
		if T.data[i].tahun == tahun {
			return i
		}
	}
	return -1
}

func TambahDataTahun(T *ArrDataTahun, tahun int, luaran string) {
	var index int = CariTahun(*T, tahun)
	var dataTahun DataTahun
	if index > -1 {
		dataTahun = T.data[index]
	}
	switch luaran {
	case "publikasi":
		dataTahun.totalPublikasi++
		dataTahun.sumKegiatan++
	case "produk":
		dataTahun.totalProduk++
		dataTahun.sumKegiatan++
	case "seminar":
		dataTahun.totalSeminar++
		dataTahun.sumKegiatan++
	case "pelatihan":
		dataTahun.totalPelatihan++
		dataTahun.sumKegiatan++
	}

	if index >= 0 {
		T.data[index] = dataTahun
	} else {
		dataTahun.tahun = tahun
		T.data[T.totalData] = dataTahun
		T.totalData++
	}
}

func outputDataTahun(T ArrDataTahun) {
	fmt.Println("--------------------------------------------------")
	for i := 0; i < T.totalData; i++ {
		fmt.Println("Data Tahun", T.data[i].tahun)
		fmt.Printf("Jumlah Kegiatan pada tahun %d adalah : %d\n", T.data[i].tahun, T.data[i].sumKegiatan)

		fmt.Println("Jumlah publikasi \t:", T.data[i].totalPublikasi)
		fmt.Println("Jumlah produk \t\t:", T.data[i].totalProduk)
		if T.tipe == "abdimas" {
			fmt.Println("Jumlah seminar \t\t:", T.data[i].totalSeminar)
			fmt.Println("Jumlah pelatihan \t:", T.data[i].totalPelatihan)
		}
		fmt.Println("--------------------------------------------------")
	}
}

func SortingJumlahKegiatan(T *ArrDataTahun, asc bool) {
	//Insertion Sort
	var i, pass int
	var temp DataTahun

	pass = 1
	for pass < T.totalData {
		i = pass
		temp = T.data[pass]
		if asc {
			//Ascending
			for i > 0 && temp.sumKegiatan < T.data[i-1].sumKegiatan {
				T.data[i] = T.data[i-1]
				i--
			}
		} else {
			//Descending
			for i > 0 && temp.sumKegiatan > T.data[i-1].sumKegiatan {
				T.data[i] = T.data[i-1]
				i--
			}
		}

		T.data[i] = temp
		pass++
	}
}

// Sorting berdasarkan judul
func SortingJudul(P *ArrKegiatan) {
	// Insertion Sort (Ascending)
	var pass, i int
	pass = 1
	for pass < P.totalKegiatan {
		i = pass
		temp := P.tabKegiatan[pass]
		for i > 0 && temp.judul < P.tabKegiatan[i-1].judul {
			P.tabKegiatan[i] = P.tabKegiatan[i-1]
			i--
		}
		P.tabKegiatan[i] = temp
		pass++
	}
}

// Searching Data
// Binary Search
func SearchingData_Judul(P ArrKegiatan) int {
	var cariJudul string
	fmt.Print("Masukkan judul : ")
	fmt.Scan(&cariJudul)

	var left, right, med int
	right = P.totalKegiatan - 1
	for left <= right {
		med = (left + right) / 2
		if P.tabKegiatan[med].judul == cariJudul {
			return med
		} else if cariJudul > P.tabKegiatan[med].judul {
			left = med + 1
		} else {
			right = med - 1
		}
	}
	return -1
}

// Sequential Search
func SearchingData_Tahun(P *ArrKegiatan) {
	var valid bool = false
	var thn int
	fmt.Print("Masukkan tahun : ")
	fmt.Scan(&thn)

	fmt.Println("========== Data Tahun", thn, "==========")
	for i := 0; i < P.totalKegiatan; i++ {
		if P.tabKegiatan[i].tahun == thn {
			fmt.Println("Judul Kegiatan \t: ", P.tabKegiatan[i].judul)
			fmt.Println("Ketua \t\t: ", P.tabKegiatan[i].ketua)
			fmt.Println("Anggota \t: ")
			for j := 0; j < 4; j++ {
				if P.tabKegiatan[i].anggota[j] != "" {
					fmt.Println("-", P.tabKegiatan[i].anggota[j])
				}
			}
			fmt.Println("Prodi / Fakultas : ", P.tabKegiatan[i].prodFak)
			fmt.Println("Sumber Dana \t:", P.tabKegiatan[i].sumberDana)
			fmt.Println("Luaran \t\t:", P.tabKegiatan[i].luaran)
			fmt.Println("Tahun Kegiatan \t:", P.tabKegiatan[i].tahun)
			fmt.Println("--------------------------------------------------")
			valid = true
		}
	}
	if valid == false {
		fmt.Println("Data tahun yang diminta tidak ditemukan")
		fmt.Println("--------------------------------------------------")
	}
}

// Sequential Search
func SearchingData_FakProdi(P *ArrKegiatan) {
	var valid bool = false
	var FakProdi string
	fmt.Print("Masukkan Fakultas / prodi : ")
	fmt.Scan(&FakProdi)

	fmt.Println("========== Data Fakultas / Prodi", FakProdi, "==========")
	for i := 0; i < P.totalKegiatan; i++ {
		if P.tabKegiatan[i].prodFak == FakProdi {
			fmt.Println("Judul \t: ", P.tabKegiatan[i].judul)
			fmt.Println("Ketua \t: ", P.tabKegiatan[i].ketua)
			fmt.Println("Anggota : ")
			for j := 0; j < 4; j++ {
				if P.tabKegiatan[i].anggota[j] != "" {
					fmt.Println("-", P.tabKegiatan[i].anggota[j])
				}
			}
			fmt.Println("Prodi / Fakultas : ", P.tabKegiatan[i].prodFak)
			fmt.Println("Sumber Dana \t:", P.tabKegiatan[i].sumberDana)
			fmt.Println("Luaran \t\t:", P.tabKegiatan[i].luaran)
			fmt.Println("Tahun \t\t:", P.tabKegiatan[i].tahun)
			fmt.Println("--------------------------------------------------")
			valid = true
		}
	}
	if valid == false {
		fmt.Println("Data fakultas / prodi yang diminta tidak ditemukan")
		fmt.Println("--------------------------------------------------")
	}
}

func PrintData(P ArrKegiatan) {
	fmt.Println("--------------------------------------------------")
	for i := 0; i < P.totalKegiatan; i++ {
		fmt.Print("Judul \t: ", P.tabKegiatan[i].judul, "\nKetua \t: ", P.tabKegiatan[i].ketua, "\nAnggota \t: ", P.tabKegiatan[i].anggota,
			"\nFakultas/Prodi \t: ", P.tabKegiatan[i].prodFak, "\nSumber Dana \t: ", P.tabKegiatan[i].sumberDana,
			"\nLuaran \t\t: ", P.tabKegiatan[i].luaran, "\nTahun \t\t: ", P.tabKegiatan[i].tahun, "\n---------------")
		fmt.Println()
	}
	fmt.Println("--------------------------------------------------")
}

func PrintDataKegiatan(data Kegiatan) {
	fmt.Println("--------------------------------------------------")
	fmt.Print("Judul \t: ", data.judul, "\nKetua \t: ", data.ketua, "\nAnggota \t: ", data.anggota,
		"\nFakultas/Prodi \t: ", data.prodFak, "\nSumber Dana \t: ", data.sumberDana,
		"\nLuaran \t\t: ", data.luaran, "\nTahun \t\t: ", data.tahun, "\n")
	fmt.Println("--------------------------------------------------")
}
