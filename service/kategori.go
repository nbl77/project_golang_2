package service

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"database/sql"


	"net/http"
)



type Kategori struct {
	Id_kategori int
	Nama_kategori string
	Id_satuan int
}


func Insert_kategori(data Kategori) Message{
	db:= dbConn()
	defer db.Close()
	fmt.Println("Tambah Kategori")

	var(
		nama_kategori string
	)

	// fmt.Println("Masukkan nama kategori")
	// fmt.Scan(&nama_kategori)

	nama_kategori = data.Nama_kategori

	insert,err := db.Prepare("INSERT INTO kategori(nama_kategori) VALUES(?)")
	if err != nil {

		message:=Message{}
		message.Status=string(http.StatusBadRequest)
		message.Message = "Ada error terkait Insert Kategori"
		fmt.Println(err)

		return message
	}

	insert.Exec(nama_kategori)


	message:=Message{}
	message.Status=string(http.StatusOK)
	message.Message = "Sukses Insert Kategori"
	fmt.Println(err)

	return message

}

func Edit_kategori(data Kategori)Message{
	db:= dbConn()
	defer db.Close()
	fmt.Println("Edit Kategori")

	var (
		id_kategori int
		// nama_kategori string
	)

	// fmt.Println("Masukkan id kategori")
	// fmt.Scan(&id_kategori)

	id_kategori = data.Id_kategori

	if FindKategori(id_kategori) {
	
		nama_kategori := data.Nama_kategori
		fmt.Println(nama_kategori)
	
		update,err := db.Prepare("UPDATE kategori SET nama_kategori=? WHERE id_kategori = ?")
		if err != nil {
			message:=Message{}
			message.Status=string(http.StatusBadRequest)
			message.Message = "Ada error terkait Edit Kategori"
			fmt.Println(err)
	
			return message
		}
	
		message:=Message{}
			message.Status=string(http.StatusOK)
			message.Message = "Sukses Edit Kategori"
			fmt.Println(err)
	
			return message
		update.Exec(nama_kategori,id_kategori)
	} else {
		fmt.Println("Id tidak ada")
	}


		message:=Message{}
			message.Status=string(http.StatusBadRequest)
			message.Message = "Tidak input apa-apa"
	
			return message
}

func ShowAllKategori(data Kategori) []Kategori{
	db := dbConn()
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

		kate.Id_kategori = id_kategori
		kate.Nama_kategori = nama_kategori
		kate.Id_satuan = id_satuan

		kateList = append(kateList, kate)
		return kateList
	}

	fmt.Println(kateList)

	return kateList

}

func ShowPerKategori(data Kategori)Kategori{
	db := dbConn()
	defer db.Close()

	var (
		id_kategori int
	)

	id_kategori = data.Id_kategori

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

		kate.Id_kategori = id_kategori1
		kate.Nama_kategori = nama_kategori
		
	}
	return kate

	fmt.Println(kate)
	} else {
		fmt.Println("Id tidak ada")
	}

	return kate

}

func Delete_kategori(data Kategori) Message{
	db :=dbConn()
	defer db.Close()

	var id_kategori int

	id_kategori = data.Id_kategori

	is_exist := FindKategori(id_kategori)

	if is_exist == true {
			
	result,err := db.Exec("DELETE FROM kategori WHERE id_kategori=?", id_kategori)

	if err!= nil {
		message:=Message{}
		message.Status=string(http.StatusBadRequest)
		message.Message = "Ada error terkait delete kategori berdasarkan id"
		fmt.Println(err)

		return message
		
	}
	message:=Message{}
	message.Status=string(http.StatusOK)
	message.Message = "sukses menghapus kategori berdasarkan id"

	return message

	//biar "result" nya "kepake"
	fmt.Println(result)
	} else {
		message:=Message{}
	message.Status=string(http.StatusBadRequest)
	message.Message = "id tidak ada di database"
		fmt.Println("Id tidak ada")
		return message
	}

	message:=Message{}
	message.Status=string(http.StatusBadRequest)
	message.Message = "tidak input apa apa"
		fmt.Println("Id tidak ada")
		return message


	


}

func FindKategori(id_kategori int) bool{
	db:=dbConn()
	defer db.Close()

	row := db.QueryRow("SELECT * FROM kategori WHERE id_kategori = ?", id_kategori)

	kate := Kategori{}
	err:= row.Scan(&kate.Id_kategori, &kate.Nama_kategori)

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