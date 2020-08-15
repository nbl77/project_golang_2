package service

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"time"
	"database/sql"
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


	if FindBarangKeluar(id_barang_keluar) {
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
	} else {
		fmt.Println("Id tidak ada")
	}


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
		id_barang_keluar int
	)

	fmt.Println("Masukkan id barang keluar")
	fmt.Scan(&id_barang_keluar)

	is_exist := FindBarangKeluar(id_barang_keluar)

	if is_exist {
		selDB,err := db.Query("SELECT * FROM barang_keluar WHERE id_barang_keluar = ?",id_barang_keluar)
		if err != nil {
			log.Fatalf("Terjadi error terkait query barang keluar by id karena error: ", err)
		}
	
		barKel:= BarangKeluar{}
	
		for selDB.Next(){
			var id_barang_keluar1 int
			var id_barang int
			var alamat string
			var jumlah_keluar int
			var waktu_keluar string
	
			err = selDB.Scan(&id_barang_keluar1, &id_barang, &alamat, &jumlah_keluar, &waktu_keluar)
			if err != nil {
				log.Fatalf("Terjadi error dikarenakan scan barang masuk by id", err)
			}

			barKel.Id_barang_keluar = id_barang_keluar1
			barKel.Id_barang = id_barang
			barKel.Alamat = alamat
			barKel.Jumlah_keluar = jumlah_keluar
			barKel.Waktu_keluar = waktu_keluar
			
		}
	
		fmt.Println(barKel)
	} else {
		fmt.Println("id tidak ada")
	}


}

func Delete_barang_keluar(){
	db :=dbConn()
	defer db.Close()

	var id_barang_keluar int
	fmt.Println("Masukkan id")

	fmt.Scan(&id_barang_keluar)

	is_exist := FindBarangKeluar(id_barang_keluar)

	if is_exist {
		result,err := db.Exec("DELETE FROM barang_keluar WHERE id_barang_keluar=?", id_barang_keluar)

		if err!= nil {
			log.Fatalln("Terjadi error terkait result", err)
		}
		fmt.Println(result)
	} else {
		fmt.Println("Id tidak ada")
	}

	



}

func FindBarangKeluar(id_barang_keluar int) bool{
	db:=dbConn()
	defer db.Close()

	row := db.QueryRow("SELECT * FROM barang_keluar WHERE id_barang_keluar = ?", id_barang_keluar)

	bar := BarangKeluar{}
	err:= row.Scan(&bar.Id_barang_keluar, &bar.Id_barang, &bar.Alamat, &bar.Jumlah_keluar, &bar.Waktu_keluar)

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
