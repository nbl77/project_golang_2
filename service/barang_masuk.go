package service

import (
	"fmt"
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"time"
	
	"net/http"
)

type BarangMasuk struct {
	Id_barang_masuk int
	Id_barang int
	Id_supplier int
	Jumlah_masuk int
	Waktu_masuk time.Time
}


func Insert_barang_masuk(data BarangMasuk)Message{
	db:= dbConn()
	defer db.Close()
	fmt.Println("Tambah Barang Masuk")

	var(
		id_barang int
		id_supplier int
		jumlah_masuk int
		waktu_masuk time.Time
	)

	id_barang = data.Id_barang
	id_supplier = data.Id_supplier
	jumlah_masuk = data.Jumlah_masuk
	waktu_masuk = data.Waktu_masuk

	insert,err := db.Prepare("INSERT INTO barang_masuk(id_barang, id_supplier, jumlah_masuk, waktu_masuk) VALUES(?,?,?,?)")
	if err != nil {
		message:=Message{}
		message.Status=string(http.StatusBadRequest)
		message.Message = "Ada error terkait Insert barang masuk karena " + err.Error()

		return message
	}

	insert.Exec(id_barang, id_supplier, jumlah_masuk, waktu_masuk)
	
	message:=Message{}
	message.Status=string(http.StatusNotFound)
	message.Message = "Sukses insert barang masuk"


	return message
}

func Edit_barang_masuk(data BarangMasuk)Message{
	db:= dbConn()
	defer db.Close()
	fmt.Println("Edit Barang Masuk")

	var (
		id_barang_masuk int
		id_barang int
		id_supplier int
		jumlah_masuk int
		waktu_masuk time.Time
	)

	id_barang_masuk = data.Id_barang_masuk

	if FindBarangMasuk(id_barang_masuk) {

		id_barang = data.Id_barang
		id_supplier = data.Id_supplier
		jumlah_masuk = data.Jumlah_masuk
		waktu_masuk = data.Waktu_masuk


		update,err := db.Prepare("UPDATE barang_masuk SET id_barang=?, id_supplier=?, jumlah_masuk =?, waktu_masuk = ? WHERE id_barang_masuk = ?")
		if err != nil {
		message:=Message{}
		message.Status=string(http.StatusBadRequest)
		message.Message = "Ada error terkait Edit Barang masuk" + err.Error()

		return message
		}

		message:=Message{}
		message.Status=string(http.StatusOK)
		message.Message = "Sukses input barang masuk"
		fmt.Println(err)
		update.Exec(id_barang,id_supplier,jumlah_masuk,waktu_masuk,id_barang_masuk)

		return message
	
		
	} else {
		message:=Message{}
		message.Status=string(http.StatusNotFound)
		message.Message = "Ada error karena id tidak ada didatabase untuk edit barng masuk"
		

		return message
	}

	message:=Message{}
	message.Status=string(http.StatusNotFound)
	message.Message = "tidak input apa apa untuk edit barang masuk"

	return message

}

func ShowAllBarangMasuk()[]BarangMasuk{
	db := dbConn()
	defer db.Close()

	barangMasuk:= BarangMasuk{}
	barangMasukList:= []BarangMasuk{}


	selDB,err := db.Query("SELECT * FROM barang_masuk ORDER BY id_barang_masuk DESC")
	if err != nil {
		fmt.Println("List barang masuk nill karena error saat query semua barang masuk" + err.Error())
		return barangMasukList
	}



	for selDB.Next(){
		var id_barang_masuk int
		var id_barang int
		var id_supplier int
		var jumlah_masuk int
		var waktu_masuk time.Time

		err = selDB.Scan(&id_barang_masuk, &id_barang, &id_supplier, &jumlah_masuk, &waktu_masuk)
		if err != nil {
			log.Fatalf("Terjadi error terkait scan Show All Barang Masuk karena: ",err)
		}

		barangMasuk.Id_barang_masuk = id_barang_masuk
		barangMasuk.Id_barang = id_barang
		barangMasuk.Id_supplier = id_supplier
		barangMasuk.Jumlah_masuk = jumlah_masuk
		barangMasuk.Waktu_masuk = waktu_masuk

		barangMasukList = append(barangMasukList, barangMasuk)
	}

	
	return barangMasukList



}

func ShowPerBarangMasuk(data BarangMasuk) BarangMasuk{
	db := dbConn()
	defer db.Close()

	var (
		id_barang_masuk int
	)
	barMas:= BarangMasuk{}

	id_barang_masuk = data.Id_barang_masuk

	if FindBarangMasuk(id_barang_masuk) {

	selDB,err := db.Query("SELECT * FROM barang_masuk WHERE id_barang_masuk = ?",id_barang_masuk)
	if err != nil {
		fmt.Println("Barang masuknya nil karena ada error ", err.Error())
		return barMas
	}

	

	for selDB.Next(){
		var id_barang_masuk1 int
		var id_barang int
		var id_supplier int
		var jumlah_masuk int
		var waktu_masuk time.Time

		err = selDB.Scan(&id_barang_masuk1, &id_barang, &id_supplier, &jumlah_masuk, &waktu_masuk)
		if err != nil {
			fmt.Println("barang masuk nil karena ada error saat di scan barang masuk")
			return barMas
		}

		barMas.Id_barang_masuk = id_barang_masuk1
		barMas.Id_barang = id_barang
		barMas.Id_supplier = id_supplier
		barMas.Jumlah_masuk = jumlah_masuk
		barMas.Waktu_masuk = waktu_masuk
	}

	return barMas
	} else {
		return barMas
		fmt.Println("barang masuk nil karena id tidak ada di database")
	}

	fmt.Println("barang masuk nil karena tidak input apa apa")
	return barMas

	//hhahahaha
}

func Delete_barang_masuk(data BarangMasuk)Message{
	db :=dbConn()
	defer db.Close()

	var id_barang_masuk int
	fmt.Println("Masukkan id")

	id_barang_masuk =data.Id_barang_masuk


	if FindBarangMasuk(id_barang_masuk){
		result,err := db.Exec("DELETE FROM barang_masuk WHERE id_barang_masuk=?", id_barang_masuk)

	if err!= nil {
		message:=Message{}
		message.Status=string(http.StatusBadRequest)
		message.Message = "Ada error terkait delete barang masuk: " + err.Error()
		fmt.Println(err)

		return message
	}
	fmt.Println(result)
	message:=Message{}
		message.Status=string(http.StatusOK)
		message.Message = "Sukses delete barang masuk"
		fmt.Println(err)

		return message
	} else {
		message:=Message{}
		message.Status=string(http.StatusNotFound)
		message.Message = "Id tidak ada di database saat mau menghapus barang masuk"
		return message
	}

	message:=Message{}
		message.Status=string(http.StatusNotFound)
		message.Message = "Tidak input apa apa terkait delete barang masuk"
		

		return message
	
	


}


func FindBarangMasuk(id_barang_masuk int) bool{
	db:=dbConn()
	defer db.Close()

	row := db.QueryRow("SELECT * FROM barang_masuk WHERE id_barang_masuk = ?", id_barang_masuk)

	bar := BarangMasuk{}
	err:= row.Scan(&bar.Id_barang_masuk, &bar.Id_barang, &bar.Id_supplier, &bar.Jumlah_masuk, &bar.Waktu_masuk)

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
