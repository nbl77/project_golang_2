package service

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"bufio"
	"os"
)



type Kategori struct {
	Id_kategori int
	Nama_kategori string
}


func Insert_kategori(){
	db:= dbConn()
	defer db.Close()
	fmt.Println("Tambah Kategori")

	var(
		nama_kategori string
	)

	fmt.Println("Masukkan nama kategori")
	fmt.Scan(&nama_kategori)

	insert,err := db.Prepare("INSERT INTO kategori(nama_kategori) VALUES(?)")
	if err != nil {
		log.Fatalf("Terjadi error terkait input kategori di database karena: ", err)
	}

	insert.Exec(nama_kategori)
	fmt.Println("Berhasil input kategori")
}

func Edit_kategori(){
	db:= dbConn()
	defer db.Close()
	fmt.Println("Edit Kategori")

	scanner := bufio.NewScanner(os.Stdin)

	var (
		id_kategori int
		// nama_kategori string
	)

	fmt.Println("Masukkan id kategori")
	fmt.Scan(&id_kategori)

	fmt.Println("Masukkan nama kategori")
	scanner.Scan()
	nama_kategori := scanner.Text()
	fmt.Println(nama_kategori)

	update,err := db.Prepare("UPDATE kategori SET nama_kategori=? WHERE id_kategori = ?")
	if err != nil {
		log.Fatalf("Terjadi error terkait edit kategori karena: ", err)
	}

	update.Exec("mobil balap",id_kategori)
}

func ShowAllKategori(){
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

		err = selDB.Scan(&id_kategori, &nama_kategori)
		if err != nil {
			log.Fatalf("Terjadi error terkait scan Show All kategori karena: ",err)
		}

		kate.Id_kategori = id_kategori
		kate.Nama_kategori = nama_kategori

		kateList = append(kateList, kate)
	}

	fmt.Println(kateList)


}

func ShowPerKategori(){
	db := dbConn()
	defer db.Close()

	var (
		id_kategori int
	)

	fmt.Println("Masukkan id kategori")
	fmt.Scan(&id_kategori)

	selDB,err := db.Query("SELECT * FROM kategori WHERE id_kategori = ?",id_kategori)
	if err != nil {
		log.Fatalf("Terjadi error terkait query kategori by id karena error: ", err)
	}

	kate:= Kategori{}

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

	fmt.Println(kate)
}

func Delete_kategori(){
	db :=dbConn()
	defer db.Close()

	var id_kategori int
	fmt.Println("Masukkan id")

	fmt.Scan(&id_kategori)

	
	result,err := db.Exec("DELETE FROM kategori WHERE id_kategori=?", id_kategori)

	if err!= nil {
		log.Fatalln("Terjadi error terkait delForm", err)
	}
	fmt.Println(result)


}