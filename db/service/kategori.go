package service

import (
	"Inventory_Project/db"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"database/sql"
	"strconv"

	"net/http"
)



type Kategori struct {
	IdKategori string
	NamaKategori string
	Satuan string
}
type Satuan struct {
	IdSatuan string
	NamaSatuan string
}


func InsertKategori(data Kategori) Message{
	db:= db.Connect()
	defer db.Close()
	fmt.Println("Tambah Kategori")

	var(
		id_satuan string
		nama_kategori string
	)

	// fmt.Println("Masukkan nama kategori")
	// fmt.Scan(&nama_kategori)

	nama_kategori = data.NamaKategori
	id_satuan = data.Satuan

	insert,err := db.Prepare("INSERT INTO kategori(id_satuan, nama_kategori) VALUES(?,?)")
	if err != nil {

		message:=Message{}
		message.Status=http.StatusBadRequest
		message.Message = "Ada error terkait Insert Kategori"
		fmt.Println(err)

		return message
	}

	insert.Exec(id_satuan, nama_kategori)


	message:=Message{}
	message.Status= http.StatusOK
	message.Message = "Sukses Insert Kategori"
	fmt.Println(err)

	return message

}

func EditKategori(data Kategori)Message{
	db:= db.Connect()
	defer db.Close()
	fmt.Println("Edit Kategori")

	var (
		id_kategori int
		// nama_kategori string
	)

	// fmt.Println("Masukkan id kategori")
	// fmt.Scan(&id_kategori)

	id_kategori,_ = strconv.Atoi(data.IdKategori)

	if FindKategori(id_kategori) {
		nama_kategori := data.NamaKategori
		id_satuan := data.Satuan
		fmt.Println("idsatuan :",id_satuan)
		update,_ := db.Prepare("UPDATE kategori SET nama_kategori=?,id_satuan=? WHERE id_kategori = ?")
		_, err:=update.Exec(nama_kategori,id_satuan,id_kategori)
		if err != nil {
			message:=Message{}
			message.Status=http.StatusBadRequest
			message.Message = "Ada error terkait Edit Kategori"
			fmt.Println(err)
			return message
		}
		message:=Message{}
		message.Status=http.StatusOK
		message.Message = "Sukses Edit Kategori"

		return message
	} else {
		fmt.Println("Id tidak ada")
	}
		message:=Message{}
			message.Status=http.StatusBadRequest
			message.Message = "Tidak input apa-apa"
	
			return message
}

func ShowAllKategori(data Kategori) []Kategori{
	db:= db.Connect()
	defer db.Close()

	selDB,err := db.Query("SELECT * FROM kategori ORDER BY id_kategori DESC")
	if err != nil {
		log.Fatalf("Terjadi error terkait Show All Kategori karena: ", err)
	}

	kate:= Kategori{}
	kateList:= []Kategori{}

	for selDB.Next(){
		var id_kategori int
		var nama_kategori string
		var id_satuan int

		err = selDB.Scan(&id_kategori, &nama_kategori, &id_satuan)
		if err != nil {
			return kateList
		}

		kate.IdKategori = string(id_kategori)
		kate.NamaKategori = nama_kategori
		kate.Satuan = string(id_satuan)

		kateList = append(kateList, kate)
		return kateList
	}

	fmt.Println(kateList)

	return kateList

}

func ShowPerKategori(data Kategori)Kategori{
	db:= db.Connect()
	defer db.Close()

	var (
		id_kategori int
	)

	id_kategori,_ = strconv.Atoi(data.IdKategori)

	is_exist := FindKategori(id_kategori)

	kate:= Kategori{}

	if is_exist {

	selDB,err := db.Query("SELECT * FROM kategori WHERE id_kategori = ?",id_kategori)
	if err != nil {
		return kate
	}

	

	for selDB.Next(){
		var id_kategori1 int
		var nama_kategori string

		err = selDB.Scan(&id_kategori1, &nama_kategori)
		if err != nil {
			log.Fatalf("Terjadi error dikarenakan scan kategori by id", err)
		}

		kate.IdKategori = string(id_kategori1)
		kate.NamaKategori = nama_kategori
		
	}
	return kate

	fmt.Println(kate)
	} else {
		fmt.Println("Id tidak ada")
	}

	return kate

}

func DeleteKategori(data Kategori) Message{
	db :=db.Connect()
	defer db.Close()

	var id_kategori int

	id_kategori,_ = strconv.Atoi(data.IdKategori)

	is_exist := FindKategori(id_kategori)

	if is_exist == true {
			
	result,err := db.Exec("DELETE FROM kategori WHERE id_kategori=?", id_kategori)

	if err!= nil {
		message:=Message{}
		message.Status=http.StatusBadRequest
		message.Message = "Ada error terkait delete kategori berdasarkan id"
		fmt.Println(err)

		return message
		
	}
	message:=Message{}
	message.Status=http.StatusOK
	message.Message = "sukses menghapus kategori berdasarkan id"

	return message

	//biar "result" nya "kepake"
	fmt.Println(result)
	} else {
		message:=Message{}
	message.Status=http.StatusBadRequest
	message.Message = "id tidak ada di database"
		fmt.Println("Id tidak ada")
		return message
	}

	message:=Message{}
	message.Status=http.StatusBadRequest
	message.Message = "tidak input apa apa"
		fmt.Println("Id tidak ada")
		return message
}

func FindKategori(id_kategori int) bool{
	db:=db.Connect()
	defer db.Close()

	row := db.QueryRow("SELECT * FROM kategori WHERE id_kategori = ?", id_kategori)

	kate := Kategori{}
	err:= row.Scan(&kate.IdKategori, &kate.NamaKategori,&kate.Satuan)

	switch {
	case err == sql.ErrNoRows:
		// log.Fatalf("Barang dari id yang dimasukkan tidak ada di database")
		return false
	case err != nil:
		log.Fatalf("Error saat mencari barang karena : ", err)
		return false
	}

	fmt.Println(kate)
	return true
}