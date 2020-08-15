package service

import (
	"fmt"
	// "database/sql"
	_ "github.com/go-sql-driver/mysql"
	"log"
)



type Supplier struct {
	Id_supplier int
	Nama_supplier string
	Alamat string
}



func Insert_supplier(){
	db:= dbConn()
	defer db.Close()
	fmt.Println("Tambah Kategori")

	var(
		nama_supplier string
		alamat string
	)

	fmt.Println("Masukkan nama supplier")
	fmt.Scan(&nama_supplier)

	fmt.Println("Masukkan alamat supplier")
	fmt.Scan(&alamat)

	insert,err := db.Prepare("INSERT INTO supplier(nama_supplier, alamat) VALUES(?,?)")
	if err != nil {
		log.Fatalf("Terjadi error terkait input supplier di database karena: ", err)
	}

	insert.Exec(nama_supplier, alamat)
	fmt.Println("Berhasil input kategori")
}

func Edit_supplier(){
	db:= dbConn()
	defer db.Close()
	fmt.Println("Edit Supplier")

	var (
		id_supplier int
		nama_supplier string
		alamat string
	)

	fmt.Println("Masukkan id supplier")
	fmt.Scan(&id_supplier)

	fmt.Println("Masukkan nama supplier")
	fmt.Scan(&nama_supplier)

	fmt.Println("Masukkan alamat")
	fmt.Scan(&alamat)

	update,err := db.Prepare("UPDATE supplier SET nama_supplier=?, alamat=? WHERE id_supplier = ?")
	if err != nil {
		log.Fatalf("Terjadi error terkait edit supplier karena: ", err)
	}

	update.Exec(nama_supplier,alamat,id_supplier)
}

func ShowAllSupplier(){
	db := dbConn()
	defer db.Close()

	selDB,err := db.Query("SELECT * FROM supplier ORDER BY id_supplier DESC")
	if err != nil {
		log.Fatalf("Terjadi error terkait Show All Supplier karena: ", err)
	}

	supp:= Supplier{}
	suppList:= []Supplier{}

	for selDB.Next(){
		var id_supplier int
		var nama_supplier string
		var alamat string

		err = selDB.Scan(&id_supplier, &nama_supplier, &alamat)
		if err != nil {
			log.Fatalf("Terjadi error terkait scan Show All Supplier karena: ",err)
		}

		supp.Id_supplier = id_supplier
		supp.Nama_supplier = nama_supplier
		supp.Alamat = alamat

		suppList = append(suppList, supp)
	}

	fmt.Println(suppList)


}

func ShowPerSupplier(){
	db := dbConn()
	defer db.Close()

	var (
		id_supplier int
	)

	fmt.Println("Masukkan id supplier")
	fmt.Scan(&id_supplier)

	selDB,err := db.Query("SELECT * FROM supplier WHERE id_supplier = ?",id_supplier)
	if err != nil {
		log.Fatalf("Terjadi error terkait query supplier by id karena error: ", err)
	}

	supp:= Supplier{}

	for selDB.Next(){
		var id_supplier1 int
		var nama_supplier string
		var alamat string

		err = selDB.Scan(&id_supplier1, &nama_supplier, &alamat)
		if err != nil {
			log.Fatalf("Terjadi error dikarenakan scan supplier by id", err)
		}

		supp.Id_supplier = id_supplier1
		supp.Nama_supplier = nama_supplier
		supp.Alamat = alamat
	}

	fmt.Println(supp)
}

func Delete_supplier(){
	db :=dbConn()
	defer db.Close()

	var id_supplier int
	fmt.Println("Masukkan id")

	fmt.Scan(&id_supplier)

	
	result,err := db.Exec("DELETE FROM supplier WHERE id_supplier=?", id_supplier)

	if err!= nil {
		log.Fatalln("Terjadi error terkait delForm", err)
	}
	fmt.Println(result)


}