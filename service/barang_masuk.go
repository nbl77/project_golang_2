package service

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"time"
)

type BarangMasuk struct {
	Id_barang_masuk int
	Id_barang int
	Id_supplier int
	Jumlah_masuk int
	Waktu_masuk string
}


func Insert_barang_masuk(){
	db:= dbConn()
	defer db.Close()
	fmt.Println("Tambah Barang Masuk")

	var(
		id_barang int
		id_supplier int
		jumlah_masuk int
	)

	fmt.Println("Masukkan id barang")
	fmt.Scan(&id_barang)

	fmt.Println("Masukkan id supplier")
	fmt.Scan(&id_supplier)

	fmt.Println("Masukkan jumlah masuk")
	fmt.Scan(&jumlah_masuk)

	currentTime := time.Now()

	insert,err := db.Prepare("INSERT INTO barang_masuk(id_barang, id_supplier, jumlah_masuk, waktu_masuk) VALUES(?,?,?,?)")
	if err != nil {
		log.Fatalf("Terjadi error terkait input barang masuk di database karena: ", err)
	}

	insert.Exec(id_barang, id_supplier, jumlah_masuk, currentTime.String())
	fmt.Println("Berhasil input barang masuk")
}

func Edit_barang_masuk(){
	db:= dbConn()
	defer db.Close()
	fmt.Println("Edit Barang Masuk")

	var (
		id_barang_masuk int
		id_barang int
		id_supplier int
		jumlah_masuk int
	)



	fmt.Println("Masukkan id barang masuk")
	fmt.Scan(&id_barang_masuk)

	fmt.Println("Masukkan id barang")
	fmt.Scan(&id_barang)

	fmt.Println("Masukkan id supplier")
	fmt.Scan(&id_supplier)

	fmt.Println("Masukkan jumlah_masuk")
	fmt.Scan(&jumlah_masuk)


	update,err := db.Prepare("UPDATE barang_masuk SET id_barang=?, id_supplier=?, jumlah_masuk =? WHERE id_barang_masuk = ?")
	if err != nil {
		log.Fatalf("Terjadi error terkait edit barang masuk karena: ", err)
	}

	update.Exec(id_barang,id_supplier,jumlah_masuk,id_barang_masuk)
}

func ShowAllBarangMasuk(){
	db := dbConn()
	defer db.Close()

	selDB,err := db.Query("SELECT * FROM barang_masuk ORDER BY id_barang_masuk DESC")
	if err != nil {
		log.Fatalf("Terjadi error terkait Show All Barang masuk karena: ", err)
	}

	barangMasuk:= BarangMasuk{}
	barangMasukList:= []BarangMasuk{}

	for selDB.Next(){
		var id_barang_masuk int
		var id_barang int
		var id_supplier int
		var jumlah_masuk int
		var waktu_masuk string

		err = selDB.Scan(&id_barang_masuk, &id_barang, &id_supplier, &jumlah_masuk, &waktu_masuk)
		if err != nil {
			log.Fatalf("Terjadi error terkait scan Show All Barang Masuk karena: ",err)
		}

		barangMasuk.Id_barang_masuk = id_barang_masuk
		barangMasuk.Id_barang = id_barang
		barangMasuk.Id_supplier = id_supplier
		barangMasuk.Jumlah_masuk = jumlah_masuk
		barangMasuk.Waktu_masuk = waktu_masuk

		barangMasukList = append(barangMasukList, barangMasuk)
	}

	fmt.Println(barangMasukList)


}

func ShowPerBarangMasuk(){
	db := dbConn()
	defer db.Close()

	var (
		id_barang_masuk int
	)

	fmt.Println("Masukkan id barang masuk")
	fmt.Scan(&id_barang_masuk)

	selDB,err := db.Query("SELECT * FROM barang_masuk WHERE id_barang_masuk = ?",id_barang_masuk)
	if err != nil {
		log.Fatalf("Terjadi error terkait query barang masuk by id karena error: ", err)
	}

	barMas:= BarangMasuk{}

	for selDB.Next(){
		var id_barang_masuk1 int
		var id_barang int
		var id_supplier int
		var jumlah_masuk int
		var waktu_masuk string

		err = selDB.Scan(&id_barang_masuk1, &id_barang, &id_supplier, &jumlah_masuk, &waktu_masuk)
		if err != nil {
			log.Fatalf("Terjadi error dikarenakan scan barang masuk by id", err)
		}

		barMas.Id_barang_masuk = id_barang_masuk1
		barMas.Id_barang = id_barang
		barMas.Id_supplier = id_supplier
		barMas.Jumlah_masuk = jumlah_masuk
		barMas.Waktu_masuk = waktu_masuk
	}

	fmt.Println(barMas)
	//hhahahaha
}

func Delete_barang_masuk(){
	db :=dbConn()
	defer db.Close()

	var id_barang_masuk int
	fmt.Println("Masukkan id")

	fmt.Scan(&id_barang_masuk)

	
	result,err := db.Exec("DELETE FROM barang_masuk WHERE id_barang_masuk=?", id_barang_masuk)

	if err!= nil {
		log.Fatalln("Terjadi error terkait result hapus barang masuk", err)
	}
	fmt.Println(result)


}