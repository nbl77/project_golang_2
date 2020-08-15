package service

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"time"
)

type BarangKeluar struct {
	Id_barang_keluar int
	Id_barang int
	Alamat string
	Jumlah_keluar int
	Waktu_keluar string
}


func Insert_barang_keluar(){
	db:= dbConn()
	defer db.Close()
	fmt.Println("Tambah Barang Keluar")

	var(
		id_barang int
		alamat string
		jumlah_keluar int
	)

	fmt.Println("Masukkan id barang")
	fmt.Scan(&id_barang)

	fmt.Println("Masukkan alamat pengiriman")
	fmt.Scan(&alamat)

	fmt.Println("Masukkan jumlah keluar")
	fmt.Scan(&jumlah_keluar)

	currentTime := time.Now()

	insert,err := db.Prepare("INSERT INTO barang_keluar(id_barang, alamat, jumlah_keluar, waktu_keluar) VALUES(?,?,?,?)")
	if err != nil {
		log.Fatalf("Terjadi error terkait input barang keluar di database karena: ", err)
	}

	insert.Exec(id_barang, alamat, jumlah_keluar, currentTime.String())
	fmt.Println("Berhasil input barang keluar")
}

func Edit_barang_keluar(){
	db:= dbConn()
	defer db.Close()
	fmt.Println("Edit Barang Masuk")

	var (
		id_barang_keluar int
		id_barang int
		alamat string
		jumlah_keluar int
	)



	fmt.Println("Masukkan id barang keluar")
	fmt.Scan(&id_barang_keluar)

	fmt.Println("Masukkan id barang")
	fmt.Scan(&id_barang)

	fmt.Println("Masukkan alamat")
	fmt.Scan(&alamat)

	fmt.Println("Masukkan jumlah_keluar")
	fmt.Scan(&jumlah_keluar)


	update,err := db.Prepare("UPDATE barang_keluar SET id_barang=?, alamat=?, jumlah_keluar =? WHERE id_barang_keluar = ?")
	if err != nil {
		log.Fatalf("Terjadi error terkait edit barang keluar karena: ", err)
	}

	update.Exec(id_barang,alamat,jumlah_keluar,id_barang_keluar)
}

func ShowAllBarangKeluar(){
	db := dbConn()
	defer db.Close()

	selDB,err := db.Query("SELECT * FROM barang_keluar ORDER BY id_barang_keluar DESC")
	if err != nil {
		log.Fatalf("Terjadi error terkait Show All Barang Keluar karena: ", err)
	}

	barangKeluar:= BarangKeluar{}
	barangKeluarList:= []BarangKeluar{}

	for selDB.Next(){
		var id_barang_keluar int
		var id_barang int
		var alamat string
		var jumlah_keluar int
		var waktu_keluar string

		err = selDB.Scan(&id_barang_keluar, &id_barang, &alamat, &jumlah_keluar, &waktu_keluar)
		if err != nil {
			log.Fatalf("Terjadi error terkait scan Show All Barang Keluar karena: ",err)
		}

		barangKeluar.Id_barang_keluar = id_barang_keluar
		barangKeluar.Id_barang = id_barang
		barangKeluar.Alamat = alamat
		barangKeluar.Jumlah_keluar = jumlah_keluar
		barangKeluar.Waktu_keluar = waktu_keluar

		barangKeluarList = append(barangKeluarList, barangKeluar)
	}

	fmt.Println(barangKeluarList)


}

func ShowPerBarangKeluar(){
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
}

func Delete_barang_keluar(){
	db :=dbConn()
	defer db.Close()

	var id_barang_keluar int
	fmt.Println("Masukkan id")

	fmt.Scan(&id_barang_keluar)

	
	result,err := db.Exec("DELETE FROM barang_keluar WHERE id_barang_keluar=?", id_barang_keluar)

	if err!= nil {
		log.Fatalln("Terjadi error terkait result", err)
	}
	fmt.Println(result)


}