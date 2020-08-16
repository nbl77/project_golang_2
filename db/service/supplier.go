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
	Id_supplier int
	Nama_supplier string
	Alamat string
}



func Insert_supplier(data Supplier)Message{
	db:= db.Connect()
	defer db.Close()
	fmt.Println("Tambah Supplier")

	var(
		nama_supplier string
		alamat string
	)

	nama_supplier = data.Nama_supplier
	alamat = data.Alamat

	insert,err := db.Prepare("INSERT INTO supplier(nama_supplier, alamat) VALUES(?,?)")
	if err != nil {
		message:=Message{}
		message.Status=http.StatusBadRequest
		message.Message = "Ada error terkait Insert supplier"
		fmt.Println(err)

		return message
	}

	insert.Exec(nama_supplier, alamat)
	
	message:=Message{}
		message.Status=http.StatusOK
		message.Message = "Sukses insert supplier"
		fmt.Println(err)

		return message

}

func Edit_supplier(data Supplier)Message{
	db:= db.Connect()
	defer db.Close()
	fmt.Println("Edit Supplier")

	var (
		id_supplier int
		nama_supplier string
		alamat string
	)

	id_supplier = data.Id_supplier

	if FindSupplier(id_supplier) {

	nama_supplier = data.Nama_supplier

	alamat = data.Alamat

	update,err := db.Prepare("UPDATE supplier SET nama_supplier=?, alamat=? WHERE id_supplier = ?")
	if err != nil {
		message:=Message{}
		message.Status=http.StatusOK
		message.Message = "Ada error terkait Edit supplier" + err.Error()

		return message
	}


	update.Exec(nama_supplier,alamat,id_supplier)
	message:=Message{}
	message.Status=http.StatusOK
	message.Message = "Sukses edit supplier"
	fmt.Println(err)

	return message
	} else {
		message:=Message{}
		message.Status=http.StatusNotFound
		message.Message = "Id tidak ada di database untuk edit supplier"
	

		return message
	}

	message:=Message{}
	message.Status=http.StatusNotFound
	message.Message = "Tidak input apa apa untuk edit supplier"


	return message
}

func ShowAllSupplier()[]Supplier{
	db := db.Connect()
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
			fmt.Println("Supplier List nil karena ada error saat scan " + err.Error())
		}

		supp.Id_supplier = id_supplier
		supp.Nama_supplier = nama_supplier
		supp.Alamat = alamat

		suppList = append(suppList, supp)
	}

	return suppList
}

func ShowPerSupplier(data Supplier)Supplier{
	db := db.Connect()
	defer db.Close()

	var (
		id_supplier int
	)

	id_supplier = data.Id_supplier
	supp:= Supplier{}

	if FindSupplier(id_supplier) {
		selDB,err := db.Query("SELECT * FROM supplier WHERE id_supplier = ?",id_supplier)
	if err != nil {
		fmt.Println("Supplier nill (per supplier) karena ada error di query ", err.Error())
		return supp
	}

	

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

	return supp
	} else {
		fmt.Println("Supplier nill karena id tidak ada di database")
		return supp
		
	}

	fmt.Println("Supplier nill karena tidak input apa apa")
	return supp

	
}

func Delete_supplier(data Supplier)Message{
	db :=db.Connect()
	defer db.Close()

	var id_supplier int

	id_supplier = data.Id_supplier

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
	} else {
		message:=Message{}
		message.Status=http.StatusNotFound
		message.Message = "error karena id tidak ada di database"
		
		return message

	}

	message:=Message{}
	message.Status=http.StatusNotFound
	message.Message = "Ada error karena tidak input apa apa untuk delete supplier"

	return message



}

func FindSupplier(id_supplier int) bool{
	db:=db.Connect()
	defer db.Close()

	row := db.QueryRow("SELECT * FROM supplier WHERE id_supplier = ?", id_supplier)

	sup := Supplier{}
	err:= row.Scan(&sup.Id_supplier, &sup.Nama_supplier, &sup.Alamat)

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