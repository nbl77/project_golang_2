package service

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"
)



type Satuan struct {
	Id_satuan int
	Nama_satuan string
}


func Insert_satuan(){
	db:= dbConn()
	defer db.Close()
	fmt.Println("Tambah Satuan")

	var(
		nama_satuan string
	)

	fmt.Println("Masukkan nama satuan")
	fmt.Scan(&nama_satuan)

	insert,err := db.Prepare("INSERT INTO satuan(nama_satuan) VALUES(?)")
	if err != nil {
		log.Fatalf("Terjadi error terkait input satuan di database karena: ", err)
	}

	insert.Exec(nama_satuan)
	fmt.Println("Berhasil input satuan")
}

func Edit_satuan(){
	db:= dbConn()
	defer db.Close()
	fmt.Println("Edit Satuan")

	var (
		id_satuan int
		nama_satuan string
	)

	fmt.Println("Masukkan id satuan")
	fmt.Scan(&id_satuan)

	fmt.Println("Masukkan nama satuan")
	fmt.Scan(&nama_satuan)

	update,err := db.Prepare("UPDATE satuan SET nama_satuan=? WHERE id_satuan = ?")
	if err != nil {
		log.Fatalf("Terjadi error terkait edit satuan karena: ", err)
	}

	update.Exec(nama_satuan,id_satuan)
}

func ShowAllSatuan(){
	db := dbConn()
	defer db.Close()

	selDB,err := db.Query("SELECT * FROM satuan ORDER BY id_satuan DESC")
	if err != nil {
		log.Fatalf("Terjadi error terkait Show All Satuan karena: ", err)
	}

	satu:= Satuan{}
	satuList:= []Satuan{}

	for selDB.Next(){
		var id_satuan int
		var nama_satuan string

		err = selDB.Scan(&id_satuan, &nama_satuan)
		if err != nil {
			log.Fatalf("Terjadi error terkait scan Show All satuan karena: ",err)
		}

		satu.Id_satuan = id_satuan
		satu.Nama_satuan = nama_satuan

		satuList = append(satuList, satu)
	}

	fmt.Println(satuList)


}

func ShowPerSatuan(){
	db := dbConn()
	defer db.Close()

	var (
		id_satuan int
	)

	fmt.Println("Masukkan id satuan")
	fmt.Scan(&id_satuan)

	selDB,err := db.Query("SELECT * FROM satuan WHERE id_satuan = ?",id_satuan)
	if err != nil {
		log.Fatalf("Terjadi error terkait query satuan by id karena error: ", err)
	}

	satu:= Satuan{}

	for selDB.Next(){
		var id_satuan1 int
		var nama_satuan string

		err = selDB.Scan(&id_satuan1, &nama_satuan)
		if err != nil {
			log.Fatalf("Terjadi error dikarenakan scan satuan by id", err)
		}

		satu.Id_satuan = id_satuan1
		satu.Nama_satuan = nama_satuan
	}

	fmt.Println(satu)
}

func Delete_satuan(){
	db :=dbConn()
	defer db.Close()

	var id_satuan int
	fmt.Println("Masukkan id")

	fmt.Scan(&id_satuan)

	
	result,err := db.Exec("DELETE FROM satuan WHERE id_satuan=?", id_satuan)

	if err!= nil {
		log.Fatalln("Terjadi error terkait result hapus satuan", err)
	}
	fmt.Println(result)


}