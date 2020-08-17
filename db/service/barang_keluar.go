package service

import (
	"Inventory_Project/db"
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"

	"net/http"
)

type BarangKeluar struct {
	IdBarangKeluar string
	IdBarang       string
	Alamat         string
	JumlahKeluar   string
	WaktuKeluar    string
}


func InsertBarangKeluar(data BarangKeluar)Message{
	db:= db.Connect()
	defer db.Close()
	fmt.Println("Tambah Barang Keluar")

	var(
		id_barang string
		alamat string
		jumlah_keluar string
		waktu_keluar string
	)

	// fmt.Println("Masukkan id barang")
	// fmt.Scan(&id_barang)

	// fmt.Println("Masukkan alamat pengiriman")
	// fmt.Scan(&alamat)

	// fmt.Println("Masukkan jumlah keluar")
	// fmt.Scan(&jumlah_keluar)

	// currentTime := time.Now()

	id_barang = data.IdBarang
	alamat = data.Alamat
	jumlah_keluar = data.JumlahKeluar
	waktu_keluar = data.WaktuKeluar

	insert,err := db.Prepare("INSERT INTO barang_keluar(id_barang, alamat, jumlah_keluar, waktu_keluar) VALUES(?,?,?,?)")
	if err != nil {
		message:=Message{}
		message.Status=http.StatusBadRequest
		message.Message = "Ada error terkait Insert Barang Keluar"
		fmt.Println(err)

		return message
	}

	message:=Message{}
	message.Status=http.StatusOK
	message.Message = "Sukses Insert Kategori"

	insert.Exec(id_barang, alamat, jumlah_keluar, waktu_keluar)
	fmt.Println("Berhasil input barang keluar")

	return message

	

}

func EditBarangKeluar(data BarangKeluar)Message{
	db:= db.Connect()
	defer db.Close()
	fmt.Println("Edit Barang Masuk")

	var (
		id_barang_keluar string
		id_barang string
		alamat string
		jumlah_keluar string
		waktu_keluar string
	)
	id_barang_keluar = data.IdBarangKeluar

	if FindBarangKeluar(id_barang_keluar) {
		id_barang = data.IdBarang
		alamat = data.Alamat
		jumlah_keluar = data.JumlahKeluar
		waktu_keluar = data.WaktuKeluar
	
	
		update,err := db.Prepare("UPDATE barang_keluar SET id_barang=?, alamat=?, jumlah_keluar =? , waktu_keluar = ? WHERE id_barang_keluar = ?")
		if err != nil {
			message:=Message{}
			message.Status=http.StatusBadRequest
			message.Message = "Ada error terkait Edit Barang Keluar"
			fmt.Println(err)
	
			return message
		}

		message:=Message{}
		message.Status=http.StatusOK
		message.Message = "Sukses Edit Barang Keluar"
		update.Exec(id_barang,alamat,jumlah_keluar, waktu_keluar,id_barang_keluar)

		return message
	
	}
	message:=Message{}
	message.Status=http.StatusBadRequest
	message.Message = "Id tidak ada"

	return message

}

func ShowAllBarangKeluar()[]BarangKeluar{
	db := db.Connect()
	defer db.Close()

	selDB,err := db.Query("SELECT * FROM barang_keluar ORDER BY id_barang_keluar DESC")
	if err != nil {
		log.Fatalf("Terjadi error terkait Show All Barang Keluar karena: ", err)
	}

	barangKeluar:= BarangKeluar{}
	barangKeluarList:= []BarangKeluar{}

	for selDB.Next(){
		var idBarangKeluar string
		var idBarang string
		var alamat string
		var jumlahKeluar string
		var waktuKeluar string

		err = selDB.Scan(&idBarangKeluar, &idBarang, &alamat, &jumlahKeluar, &waktuKeluar)
		if err != nil {
			fmt.Println("Ada error saat scan untuk show all barang keluar karena: ", err)
			return barangKeluarList
		}

		barangKeluar.IdBarangKeluar = idBarangKeluar
		barangKeluar.IdBarang = idBarang
		barangKeluar.Alamat = alamat
		barangKeluar.JumlahKeluar = jumlahKeluar
		barangKeluar.WaktuKeluar = waktuKeluar

		barangKeluarList = append(barangKeluarList, barangKeluar)
		return barangKeluarList
	}

	fmt.Println("barang keluar list null karena tidak input apa apa")
	return barangKeluarList


}

func ShowPerBarangKeluar(data BarangKeluar)BarangKeluar{
	db := db.Connect()
	defer db.Close()

	var (
		id_barang_keluar string
	)

	id_barang_keluar = data.IdBarangKeluar

	is_exist := FindBarangKeluar(id_barang_keluar)

	barKel:= BarangKeluar{}

	if is_exist {
		selDB,err := db.Query("SELECT * FROM barang_keluar WHERE id_barang_keluar = ?",id_barang_keluar)
		if err != nil {
			fmt.Println("Terjadi error saat select barang keluar berdasarkan id karena: ", err)
			return barKel
		}
	
	
		for selDB.Next(){
			var id_barang_keluar1 string
			var id_barang string
			var alamat string
			var jumlah_keluar string
			var waktu_keluar string
	
			err = selDB.Scan(&id_barang_keluar1, &id_barang, &alamat, &jumlah_keluar, &waktu_keluar)
			if err != nil {
				fmt.Println("Terjadi error saat scan barang keluar")
				return barKel
			}

			barKel.IdBarangKeluar = id_barang_keluar1
			barKel.IdBarang = id_barang
			barKel.Alamat = alamat
			barKel.JumlahKeluar = jumlah_keluar
			barKel.WaktuKeluar = waktu_keluar
			
		}
		return barKel
	
		
	} else {
		fmt.Println("Id tidak ada")
		return barKel
		
	}

	fmt.Println("Tidak input apa apa")
	return barKel


}

func DeleteBarangKeluar(data BarangKeluar)Message{
	db :=db.Connect()
	defer db.Close()

	var id_barang_keluar string

	id_barang_keluar = data.IdBarangKeluar

	is_exist := FindBarangKeluar(id_barang_keluar)

	if is_exist {
		result,err := db.Exec("DELETE FROM barang_keluar WHERE id_barang_keluar=?", id_barang_keluar)

		if err!= nil {
			
		message:=Message{}
		message.Status=http.StatusBadRequest
		message.Message = "Ada error saat mengeksekusi delete di barang keluar karena: " + err.Error()
		fmt.Println(err)

		return message
		}

		message:=Message{}
		message.Status=http.StatusOK
		message.Message = "Berhasil delete barang keluar"

		return message

		//ini resultnya biar kepake
		fmt.Println(result)
	} else {
		message:=Message{}
		message.Status=http.StatusNotFound
		message.Message = "id tidak ditemukan"

		return message
	}


	message:=Message{}
	message.Status=http.StatusNotFound
	message.Message = "tidak input apa apa"

	return message
}

func FindBarangKeluar(id_barang_keluar string) bool{
	db:=db.Connect()
	defer db.Close()

	row := db.QueryRow("SELECT * FROM barang_keluar WHERE id_barang_keluar = ?", id_barang_keluar)

	bar := BarangKeluar{}
	err:= row.Scan(&bar.IdBarangKeluar, &bar.IdBarang, &bar.Alamat, &bar.JumlahKeluar, &bar.WaktuKeluar)

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
