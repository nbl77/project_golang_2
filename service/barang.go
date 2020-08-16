package service

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"
	
	"net/http"
	
	
)

type Barang struct {
	Id_barang int		
	Nama_barang string
	Stok int
	Id_kategori int
	Id_satuan int
}

type Message struct{
	Status string
	Message string
}


func Insert_barang(data Barang) Message{
	db:= dbConn()
	defer db.Close()
	fmt.Println("Tambah Barang")



	// var(
	// 	nama_barang string
	// 	stok int
	// 	id_kategori int
	// 	id_satuan int
	// )

	// fmt.Println("Masukkan nama barang")
	// fmt.Scan(&nama_barang)

	// fmt.Println("Masukkan jumlah stok")
	// fmt.Scan(&stok)

	// fmt.Println("Masukkan id_kategori")
	// fmt.Scan(&id_kategori)

	// fmt.Println("Masukkan id satuan")
	// fmt.Scan(&id_satuan)


	nama_barang := data.Nama_barang
	stok:= 0
	id_kategori := data.Id_kategori
	id_satuan := data.Id_satuan


	insert,err := db.Prepare("INSERT INTO barang(nama_barang, stok, id_kategori, id_satuan) VALUES(?,?,?,?)")
	if err != nil {
		message:=Message{}
		message.Status=string(http.StatusBadRequest)
		message.Message = "Ada error terkait Insert Barang"
		fmt.Println(err)

		return message
	

	}
	insert.Exec(nama_barang, stok, id_kategori, id_satuan)

	return Message{
		Status: string(http.StatusOK),
		Message : "Sukses menambahkan barang",

	}
	
}

func Edit_barang(data Barang)Message{
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

	id_barang = data.Id_barang

	if FindBarang(id_barang) {
	
	nama_barang = data.Nama_barang
	stok = data.Stok
	id_kategori = data.Id_kategori
	id_satuan = data.Id_satuan


		update,err := db.Prepare("UPDATE barang SET nama_barang=?, stok=?, id_kategori =?, id_satuan =? WHERE id_barang = ?")
		if err != nil {
			message:=Message{}
			message.Status=string(http.StatusBadRequest)
			message.Message = "Ada error terkait Update Barang"
			fmt.Println(err)
	
			return message
		}

		// ini ga terlalu penting, buat jaga jaga aja
		if err == sql.ErrNoRows{
			message:=Message{}
			message.Status=string(http.StatusBadRequest)
			message.Message = "Ada error karena row tidak ditemukan"
			fmt.Println(err)
	
			return message
		
		}
	
		update.Exec(nama_barang,stok,id_kategori, id_satuan, id_barang)
	} else {
		fmt.Println("Id tidak ada")
	
		message:=Message{}
		message.Status=string(http.StatusBadRequest)
		message.Message = "Id tidak ada"

		return message

	}

	message:=Message{}
	message.Status=string(http.StatusOK)
	message.Message = "Sukses Edit barang"

	return message



}

func ShowAllBarang()([]Barang,Message){
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
			message:=Message{}
			message.Status=string(http.StatusBadRequest)
			message.Message = "ada error terkait scan semua barang"
	
			return barList,message
	
		}

		bar.Id_barang = id_barang
		bar.Nama_barang = nama_barang
		bar.Stok = stok
		bar.Id_kategori = id_kategori
		bar.Id_satuan = id_satuan

		barList = append(barList, bar)
	}

	fmt.Println(barList)
	message:=Message{}
	message.Status=string(http.StatusOK)
	message.Message = "Sukses menampilkan semua barang"

	return barList,message

}

func ShowPerBarang(data Barang)(Barang,Message){
	db := dbConn()
	defer db.Close()

	var (
		id_barang int
	)

	id_barang = data.Id_barang

	if FindBarang(id_barang) {
		selDB,err := db.Query("SELECT * FROM barang WHERE id_barang = ?",id_barang)
		if err != nil {

			message:=Message{}
			message.Status=string(http.StatusBadRequest)
			message.Message = "Terjadi error saat mau melakukan select semuanya di barang"
			
			return Barang{}, message
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

			message:=Message{}
			message.Status=string(http.StatusOK)
			message.Message = "Sukses Mengambil barang"
		
			return bar, message
	} else {

		message:=Message{}
		message.Status=string(http.StatusBadRequest)
		message.Message = "Id tidak ada"
	
		return Barang{}, message
		
	}


	return Barang{},Message{
		Status:string(http.StatusBadRequest),
		Message:"Tidak mengembalikan apa apa karena tidak menginput apa apa",
	}


	

}

func Delete_barang(data Barang) Message{
	db :=dbConn()
	defer db.Close()

	var id_barang int
	// fmt.Println("Masukkan id")

	// fmt.Scan(&id_barang)

	id_barang = data.Id_barang
	is_exist:= FindBarang(id_barang)

	if is_exist == true {
		result,err := db.Exec("DELETE FROM barang WHERE id_barang=?", id_barang)

		if err!= nil {
			message:=Message{}
			message.Status=string(http.StatusBadRequest)
			message.Message = "Terjadi error saat mengeksekusi delete barang"

			return message
		}

		// ini tidak penting, buat jaga jaga aja karena udah ada handlernya
		// di atas "is_exist"
		if err == sql.ErrNoRows{
			message:=Message{}
			message.Status=string(http.StatusBadRequest)
			message.Message = "Row tidak ditemukan"	

			return message
		}
		fmt.Println(result)
	} else {

		message:=Message{}
		message.Status=string(http.StatusBadRequest)
		message.Message = "Id tidak ada"	

		return message
		
	}

		message:=Message{}
		message.Status=string(http.StatusOK)
		message.Message = "Berhasil dihapus"	

		return message
	



}

func FindBarang(id_barang int) bool{
	db:=dbConn()
	defer db.Close()

	row := db.QueryRow("SELECT * FROM barang WHERE id_barang = ?", id_barang)

	bar := Barang{}
	err:= row.Scan(&bar.Id_barang, &bar.Nama_barang, &bar.Stok, &bar.Id_kategori, &bar.Id_satuan)

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
