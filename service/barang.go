package service

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

type Barang struct {
	Id_barang int
	Nama_barang string
	Stok int
	Id_kategori int
	Id_satuan int
}


func Insert_barang(){
	db:= dbConn()
	defer db.Close()
	fmt.Println("Tambah Barang")

	var(
		nama_barang string
		stok int
		id_kategori int
		id_satuan int
	)

	fmt.Println("Masukkan nama barang")
	fmt.Scan(&nama_barang)

	fmt.Println("Masukkan jumlah stok")
	fmt.Scan(&stok)

	fmt.Println("Masukkan id_kategori")
	fmt.Scan(&id_kategori)

	fmt.Println("Masukkan id satuan")
	fmt.Scan(&id_satuan)

	insert,err := db.Prepare("INSERT INTO barang(nama_barang, stok, id_kategori, id_satuan) VALUES(?,?,?,?)")
	if err != nil {
		log.Fatalf("Terjadi error terkait input barang di database karena: ", err)
	}

	insert.Exec(nama_barang, stok, id_kategori, id_satuan)
	fmt.Println("Berhasil input barang")
}

func Edit_barang(){
	db:= dbConn()
	defer db.Close()
	fmt.Println("Edit Barang")

	var (
		id_barang int
		nama_barang string
		stok int
		id_kategori int
		id_satuan int
	)

	fmt.Println("Masukkan id barang")
	fmt.Scan(&id_barang)

	fmt.Println("Masukkan nama barang")
	fmt.Scan(&nama_barang)

	fmt.Println("Masukkan stok")
	fmt.Scan(&stok)

	fmt.Println("Masukkan id_kategori")
	fmt.Scan(&id_kategori)

	fmt.Println("Masukkan id_satuan")
	fmt.Scan(&id_satuan)

	update,err := db.Prepare("UPDATE barang SET nama_barang=?, stok=?, id_kategori =?, id_satuan =? WHERE id_barang = ?")
	if err != nil {
		log.Fatalf("Terjadi error terkait edit supplier karena: ", err)
	}

	update.Exec(nama_barang,stok,id_kategori, id_satuan, id_barang)
}

func ShowAllBarang(){
	db := dbConn()
	defer db.Close()

	selDB,err := db.Query("SELECT * FROM barang ORDER BY id_barang DESC")
	if err != nil {
		log.Fatalf("Terjadi error terkait Show All Supplier karena: ", err)
	}

	bar:= Barang{}
	barList:= []Barang{}

	for selDB.Next(){
		var id_barang int
		var nama_barang string
		var stok int
		var id_kategori int
		var id_satuan int

		err = selDB.Scan(&id_barang, &nama_barang, &stok, &id_kategori, &id_satuan)
		if err != nil {
			log.Fatalf("Terjadi error terkait scan Show All Supplier karena: ",err)
		}

		bar.Id_barang = id_barang
		bar.Nama_barang = nama_barang
		bar.Stok = stok
		bar.Id_kategori = id_kategori
		bar.Id_satuan = id_satuan

		barList = append(barList, bar)
	}

	fmt.Println(barList)


}

func ShowPerBarang(){
	db := dbConn()
	defer db.Close()

	var (
		id_barang int
	)

	fmt.Println("Masukkan id barang")
	fmt.Scan(&id_barang)

	selDB,err := db.Query("SELECT * FROM barang WHERE id_barang = ?",id_barang)
	if err != nil {
		log.Fatalf("Terjadi error terkait query barang by id karena error: ", err)
	}

	bar:= Barang{}

	for selDB.Next(){
		var id_barang1 int
		var nama_barang string
		var stok int
		var id_kategori int
		var id_satuan int

		err = selDB.Scan(&id_barang1, &nama_barang, &stok, &id_kategori, &id_satuan)
		if err != nil {
			log.Fatalf("Terjadi error dikarenakan scan barang by id", err)
		}

		bar.Id_barang = id_barang1
		bar.Nama_barang = nama_barang
		bar.Stok = stok
		bar.Id_kategori = id_kategori
		bar.Id_satuan = id_satuan
	}

	fmt.Println(bar)
}

func Delete_barang(){
	db :=dbConn()
	defer db.Close()

	var id_barang int
	fmt.Println("Masukkan id")

	fmt.Scan(&id_barang)

	
	result,err := db.Exec("DELETE FROM barang WHERE id_barang=?", id_barang)

	if err!= nil {
		log.Fatalln("Terjadi error terkait delForm", err)
	}
	fmt.Println(result)


}
