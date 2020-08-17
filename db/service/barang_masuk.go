package service

import (
	"Inventory_Project/db"
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"

	"net/http"
)

type BarangMasuk struct {
	IdBarangMasuk string
	IdBarang      string
	IdSupplier    string
	JumlahMasuk   string
	WaktuMasuk    string
}


func InsertBarangMasuk(data BarangMasuk)Message{
	db:= db.Connect()
	defer db.Close()
	fmt.Println("Tambah Barang Masuk")

	var(
		id_barang string
		id_supplier string
		jumlah_masuk string
		waktu_masuk string
	)

	id_barang = data.IdBarang
	id_supplier = data.IdSupplier
	jumlah_masuk = data.JumlahMasuk
	waktu_masuk = data.WaktuMasuk

	insert,err := db.Prepare("INSERT INTO barang_masuk(id_barang, id_supplier, jumlah_masuk, waktu_masuk) VALUES(?,?,?,?)")
	if err != nil {
		message:=Message{}
		message.Status=http.StatusBadRequest
		message.Message = "Ada error terkait Insert barang masuk karena " + err.Error()

		return message
	}

	insert.Exec(id_barang, id_supplier, jumlah_masuk, waktu_masuk)
	
	message:=Message{}
	message.Status=http.StatusOK
	message.Message = "Sukses insert barang masuk"


	return message
}

func EditBarangMasuk(data BarangMasuk)Message{
	db:= db.Connect()
	defer db.Close()
	fmt.Println("Edit Barang Masuk")

	var (
		idBarangMasuk string
		idBarang      string
		idSupplier    string
		jumlahMasuk   string
		waktuMasuk    string
	)

	idBarangMasuk = data.IdBarangMasuk

	if FindBarangMasuk(idBarangMasuk) {

		idBarang = data.IdBarang
		idSupplier = data.IdSupplier
		jumlahMasuk = data.JumlahMasuk
		waktuMasuk = data.WaktuMasuk


		update,err := db.Prepare("UPDATE barang_masuk SET id_barang=?, id_supplier=?, jumlah_masuk =?, waktu_masuk = ? WHERE id_barang_masuk = ?")
		if err != nil {
		message:=Message{}
		message.Status=http.StatusBadRequest
		message.Message = "Ada error terkait Edit Barang masuk" + err.Error()

		return message
		}

		message:=Message{}
		message.Status=http.StatusOK
		message.Message = "Sukses input barang masuk"
		fmt.Println(err)
		update.Exec(idBarang, idSupplier, jumlahMasuk, waktuMasuk, idBarangMasuk)

		return message
	
		
	}
	message:=Message{}
	message.Status=http.StatusNotFound
	message.Message = "Ada error karena id tidak ada didatabase untuk edit barng masuk"


	return message

}

func ShowAllBarangMasuk()[]BarangMasuk{
	db := db.Connect()
	defer db.Close()

	barangMasuk:= BarangMasuk{}
	barangMasukList:= []BarangMasuk{}


	selDB,err := db.Query("SELECT * FROM barang_masuk ORDER BY id_barang_masuk DESC")
	if err != nil {
		fmt.Println("List barang masuk nill karena error saat query semua barang masuk" + err.Error())
		return barangMasukList
	}



	for selDB.Next(){
		var idBarangMasuk string
		var idBarang string
		var idSupplier string
		var jumlah_masuk string
		var waktuMasuk string

		err = selDB.Scan(&idBarangMasuk, &idBarang, &idSupplier, &jumlah_masuk, &waktuMasuk)
		if err != nil {
			log.Fatalf("Terjadi error terkait scan Show All Barang Masuk karena: ",err)
		}

		barangMasuk.IdBarangMasuk = idBarangMasuk
		barangMasuk.IdBarang = idBarang
		barangMasuk.IdSupplier = idSupplier
		barangMasuk.JumlahMasuk = jumlah_masuk
		barangMasuk.WaktuMasuk = waktuMasuk

		barangMasukList = append(barangMasukList, barangMasuk)
	}

	
	return barangMasukList



}

func ShowPerBarangMasuk(data BarangMasuk) BarangMasuk{
	db := db.Connect()
	defer db.Close()

	var (
		id_barang_masuk string
	)
	barMas:= BarangMasuk{}

	id_barang_masuk = data.IdBarangMasuk

	if FindBarangMasuk(id_barang_masuk) {

	selDB,err := db.Query("SELECT * FROM barang_masuk WHERE id_barang_masuk = ?",id_barang_masuk)
	if err != nil {
		fmt.Println("Barang masuknya nil karena ada error ", err.Error())
		return barMas
	}

	

	for selDB.Next(){
		var id_barang_masuk1 string
		var id_barang string
		var id_supplier string
		var jumlah_masuk string
		var waktu_masuk string

		err = selDB.Scan(&id_barang_masuk1, &id_barang, &id_supplier, &jumlah_masuk, &waktu_masuk)
		if err != nil {
			fmt.Println("barang masuk nil karena ada error saat di scan barang masuk")
			return barMas
		}

		barMas.IdBarangMasuk = id_barang_masuk1
		barMas.IdBarang = id_barang
		barMas.IdSupplier = id_supplier
		barMas.JumlahMasuk = jumlah_masuk
		barMas.WaktuMasuk = waktu_masuk
	}

	return barMas
	} else {
		return barMas
		fmt.Println("barang masuk nil karena id tidak ada di database")
	}

	fmt.Println("barang masuk nil karena tidak input apa apa")
	return barMas

	//hhahahaha
}

func DeleteBarangMasuk(data BarangMasuk)Message{
	db :=db.Connect()
	defer db.Close()

	var id_barang_masuk string
	fmt.Println("Masukkan id")

	id_barang_masuk =data.IdBarangMasuk


	if FindBarangMasuk(id_barang_masuk){
		result,err := db.Exec("DELETE FROM barang_masuk WHERE id_barang_masuk=?", id_barang_masuk)

	if err!= nil {
		message:=Message{}
		message.Status=http.StatusBadRequest
		message.Message = "Ada error terkait delete barang masuk: " + err.Error()
		fmt.Println(err)

		return message
	}
	fmt.Println(result)
	message:=Message{}
		message.Status=http.StatusOK
		message.Message = "Sukses delete barang masuk"
		fmt.Println(err)

		return message
	}
	message:=Message{}
	message.Status=http.StatusNotFound
	message.Message = "Id tidak ada di database saat mau menghapus barang masuk"
	return message
}


func FindBarangMasuk(id_barang_masuk string) bool{
	db:=db.Connect()
	defer db.Close()

	row := db.QueryRow("SELECT * FROM barang_masuk WHERE id_barang_masuk = ?", id_barang_masuk)

	bar := BarangMasuk{}
	err:= row.Scan(&bar.IdBarangMasuk, &bar.IdBarang, &bar.IdSupplier, &bar.JumlahMasuk, &bar.WaktuMasuk)

	switch {
	case err == sql.ErrNoRows:
		// log.Fatalf("Barang dari id yang dimasukkan tidak ada di database")
		return false
	case err != nil:
		log.Fatalf("Error saat mencari barang karena : ", err)
		return false
	}

	fmt.Println(bar)
	return true
}
