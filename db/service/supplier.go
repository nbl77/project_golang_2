package service

import (
	"Inventory_Project/db"
	"fmt"
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"log"


	"net/http"
)



type Supplier struct {
	IdSupplier   string
	NamaSupplier string
	Alamat       string
	NoTelp       string
	TotalMasuk string
	LastInsert string
}



func InsertSupplier(data Supplier)Message{
	db:= db.Connect()
	defer db.Close()
	fmt.Println("Tambah Supplier")

	var(
		nama_supplier string
		alamat string
		no_telp string
	)

	nama_supplier = data.NamaSupplier
	alamat = data.Alamat
	no_telp = data.NoTelp

	insert,err := db.Prepare("INSERT INTO supplier(nama_supplier,no_telp, alamat) VALUES(?,?,?)")
	if err != nil {
		message:=Message{}
		message.Status=http.StatusBadRequest
		message.Message = "Ada error terkait Insert supplier"
		fmt.Println(err)

		return message
	}

	insert.Exec(nama_supplier,no_telp, alamat)
	message:=Message{}
		message.Status=http.StatusOK
		message.Message = "Sukses insert supplier"
		fmt.Println(err)

		return message
}

func EditSupplier(data Supplier)Message{
	db:= db.Connect()
	defer db.Close()
	fmt.Println("Edit Supplier")

	var (
		id_supplier string
		nama_supplier string
		no_telp string
		alamat string
	)

	id_supplier = data.IdSupplier

	if FindSupplier(id_supplier) {

	nama_supplier = data.NamaSupplier
	no_telp = data.NoTelp
	alamat = data.Alamat

	update,err := db.Prepare("UPDATE supplier SET nama_supplier=?,no_telp=?, alamat=? WHERE id_supplier = ?")
	if err != nil {
		message:=Message{}
		message.Status=http.StatusOK
		message.Message = "Ada error terkait Edit supplier" + err.Error()

		return message
	}


	update.Exec(nama_supplier,no_telp,alamat,id_supplier)
	message:=Message{}
	message.Status=http.StatusOK
	message.Message = "Sukses edit supplier"
	fmt.Println(err)

	return message
	}

	message:=Message{}
	message.Status=http.StatusNotFound
	message.Message = "Id tidak ada di database untuk edit supplier"


	return message
}

func ShowAllSupplier()[]Supplier{
	db2 := db.Connect()
	defer db2.Close()

	selDB,err := db2.Query("SELECT supplier.* FROM supplier ORDER BY id_supplier DESC")
	if err != nil {
		log.Fatalf("Terjadi error terkait Show All Supplier karena: ", err)
	}

	supp:= Supplier{}
	suppList:= []Supplier{}

	for selDB.Next(){
		var idSupplier string
		var namaSupplier string
		var noTelp string
		var alamat string
		var totalBarang sql.NullString
		var lastInsert sql.NullString
		err = selDB.Scan(&idSupplier, &namaSupplier,&noTelp, &alamat)
		db.SelectParam("SELECT SUM(jumlah_masuk),MAX(waktu_masuk) FROM barang_masuk WHERE id_supplier=?", idSupplier).Scan(&totalBarang,&lastInsert)
		if err != nil {
			fmt.Println("Supplier List nil karena ada error saat scan " + err.Error())
		}

		supp.IdSupplier = idSupplier
		supp.NamaSupplier = namaSupplier
		supp.Alamat = alamat
		supp.NoTelp = noTelp
		if totalBarang.Valid {
			supp.TotalMasuk = totalBarang.String
		}
		if lastInsert.Valid {
			supp.LastInsert = lastInsert.String
		}

		suppList = append(suppList, supp)
	}

	return suppList
}

func ShowPerSupplier(data Supplier)Supplier{
	db := db.Connect()
	defer db.Close()

	var (
		id_supplier string
	)

	id_supplier = data.IdSupplier
	supp:= Supplier{}

	if FindSupplier(id_supplier) {
		selDB,err := db.Query("SELECT * FROM supplier WHERE id_supplier = ?",id_supplier)
	if err != nil {
		fmt.Println("Supplier nill (per supplier) karena ada error di query ", err.Error())
		return supp
	}

	

	for selDB.Next(){
		var id_supplier1 string
		var nama_supplier string
		var no_telp string
		var alamat string

		err = selDB.Scan(&id_supplier1, &nama_supplier,&no_telp, &alamat)
		if err != nil {
			log.Fatalf("Terjadi error dikarenakan scan supplier by id", err)
		}

		supp.IdSupplier = id_supplier1
		supp.NamaSupplier = nama_supplier
		supp.Alamat = alamat
		supp.NoTelp = no_telp
	}

	return supp
	}
	fmt.Println("Supplier nill karena id tidak ada di database")
	return supp
}

func DeleteSupplier(data Supplier)Message{
	db :=db.Connect()
	defer db.Close()

	var id_supplier string

	id_supplier = data.IdSupplier

	if FindSupplier(id_supplier) {
		result,err := db.Exec("DELETE FROM supplier WHERE id_supplier=?", id_supplier)

		if err!= nil {
			message:=Message{}
		message.Status=http.StatusBadRequest
		message.Message = "Ada error saat mengeksekusi delete untuk supplier" + err.Error()
		fmt.Println(err)

		return message
		}
		message:=Message{}
		message.Status=http.StatusOK
		message.Message = "Berhasil menghapus supplier"
		fmt.Println(result)
		return message
	}
	message:=Message{}
	message.Status=http.StatusNotFound
	message.Message = "error karena id tidak ada di database"
	return message
}

func FindSupplier(id_supplier string) bool{
	db:=db.Connect()
	defer db.Close()

	row := db.QueryRow("SELECT * FROM supplier WHERE id_supplier = ?", id_supplier)

	sup := Supplier{}
	err:= row.Scan(&sup.IdSupplier, &sup.NamaSupplier,&sup.NoTelp, &sup.Alamat)

	switch {
	case err == sql.ErrNoRows:
		// log.Fatalf("Barang dari id yang dimasukkan tidak ada di database")
		return false
	case err != nil:
		log.Fatalf("Error saat mencari barang karena : ", err)
		return false
	}

	fmt.Println(sup)
	return true
}