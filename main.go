package main

import(
	// "fmt"
	"log"
	// "golang_echo/service"
	"golang_echo/service/config"
)
func main(){
	// flag := false

	// for !flag {
	// 	fmt.Println("Masukkan menu")

	// 	var angkaSwitch int
	// 	fmt.Scan(&angkaSwitch)

	// 	switch angkaSwitch{
	// 		case 1:
	// 			service.Insert_barang()
	// 			break
	// 		case 2:
	// 			service.Insert_kategori()
	// 			break
	// 		case 3:
	// 			service.Insert_barang_keluar()
	// 			break
	// 		case 4:
	// 			service.Insert_satuan()
	// 			break
	// 		case 5:
	// 			service.Insert_supplier()
	// 			break
	// 		case 6:
	// 			service.Insert_barang_masuk()
	// 			break
			
	// 		//edit
	// 		case 7:
	// 			service.Edit_barang()
	// 			break
	// 		case 8:
	// 			service.Edit_kategori()
	// 			break
	// 		case 9:
	// 			service.Edit_barang_keluar()
	// 			break
	// 		case 10:
	// 			service.Edit_satuan()
	// 			break
	// 		case 11:
	// 			service.Edit_supplier()
	// 			break
	// 		case 12:
	// 			service.Edit_barang_masuk()
	// 			break

	// 		//Shown all
	// 		case 13:
	// 			service.ShowAllBarang()
	// 			break
	// 		case 14:
	// 			service.ShowAllKategori()
	// 			break
	// 		case 15:
	// 			service.ShowAllBarangKeluar()
	// 			break
	// 		case 16:
	// 			service.ShowAllSatuan()
	// 			break
	// 		case 17:
	// 			service.ShowAllSupplier()
	// 			break
	// 		case 18:
	// 			service.ShowAllBarangMasuk()

	// 		//show per
	// 		case 19:
	// 			service.ShowPerBarang()
	// 			break
	// 		case 20:
	// 			service.ShowPerKategori()
	// 			break
	// 		case 21:
	// 			service.ShowPerBarangKeluar()
	// 			break
	// 		case 22:
	// 			service.ShowPerSatuan()
	// 			break
	// 		case 23:
	// 			service.ShowPerSupplier()
	// 			break
	// 		case 24:
	// 			service.ShowPerBarangMasuk()
	// 			break

	// 		//delete
	// 		case 25:
	// 			service.Delete_barang()
	// 			break
	// 		case 26:
	// 			service.Delete_kategori()
	// 			break
	// 		case 27:
	// 			service.Delete_barang_keluar()
	// 			break
	// 		case 28:
	// 			service.Delete_satuan()
	// 			break
	// 		case 29:
	// 			service.Delete_supplier()
	// 			break
	// 		case 30:
	// 			service.Delete_barang_masuk()
	// 			break

			log.Print("Server start at ", ":8080")

			r:= config.Server()
			
			r.Start(":8080")
	
}