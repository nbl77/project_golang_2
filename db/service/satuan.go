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




func InsertSatuan(data Satuan)Message{
	db:= db.Connect()
	defer db.Close()
	log.Println("Tambah Satuan")

	var(
		nama_satuan string
	)

	nama_satuan = data.NamaSatuan

	insert,err := db.Prepare("INSERT INTO satuan(nama_satuan) VALUES(?)")
	if err != nil {
		message:=Message{}
		message.Status=string(http.StatusBadRequest)
		message.Message = "Ada error terkait Insert Satuan karena " + err.Error()
		log.Println(err)

		return message
	}
	insert.Exec(nama_satuan)
	message:=Message{}
	message.Status =http.StatusOK
	message.Message = "Sukses insert satuan"
	return message
}

func EditSatuan(data Satuan)Message{
	db:= db.Connect()
	defer db.Close()
	fmt.Println("Edit Satuan")

	var (
		id_satuan int
		nama_satuan string
	)

	id_satuan,_ = strconv.Atoi(data.IdSatuan)

	if FindSatuan(id_satuan) {

	nama_satuan = data.NamaSatuan

	update,err := db.Prepare("UPDATE satuan SET nama_satuan=? WHERE id_satuan = ?")
	if err != nil {
		message:=Message{}
		message.Status=http.StatusBadRequest
		message.Message = "Ada error terkait Edit Satuan karena: " + err.Error()
		fmt.Println(err)

		return message
	}

	update.Exec(nama_satuan,id_satuan)
	message:=Message{}
	message.Status=http.StatusOK
	message.Message = "Sukses Edit Satuan"

	return message

	}
	fmt.Println("Id tidak ada")
	message:=Message{}
	message.Status=http.StatusNotFound
	message.Message = "Id Tidak ada"

	return message
}

func ShowAllSatuan()[]Satuan{
	db:= db.Connect()
	defer db.Close()

	satu:= Satuan{}
	satuList:= []Satuan{}

	selDB,err := db.Query("SELECT * FROM satuan ORDER BY id_satuan DESC")
	if err != nil {
		fmt.Println("ada error saat query untuk satuan karena ", err.Error())
		return satuList
	}

	

	for selDB.Next(){
		var id_satuan int
		var nama_satuan string

		err = selDB.Scan(&id_satuan, &nama_satuan)
		if err != nil {
			fmt.Println("List untuk satuan nil karena error saat scan: ", err.Error())
			return satuList
		}

		satu.IdSatuan = string(id_satuan)
		satu.NamaSatuan = nama_satuan

		satuList = append(satuList, satu)
	}

	fmt.Println(satuList)
	return satuList
	

}

func ShowPerSatuan(data Satuan)Satuan{
	db:= db.Connect()
	defer db.Close()

	var (
		id_satuan int
	)

	id_satuan,_ = strconv.Atoi(data.IdSatuan)

	satu:= Satuan{}

	is_exist := FindSatuan(id_satuan)

	if is_exist {
		selDB,err := db.Query("SELECT * FROM satuan WHERE id_satuan = ?",id_satuan)
		if err != nil {
			fmt.Println("Error saat query untuk satuan (bukan yang all satuan) karena ", err.Error())
			return satu
		}
		for selDB.Next(){
			var id_satuan1 int
			var nama_satuan string
	
			err = selDB.Scan(&id_satuan1, &nama_satuan)
			if err != nil {
				fmt.Println("Error saat meng-scan satuan karena ", err.Error())
				return satu
			}
			satu.IdSatuan = string(id_satuan1)
			satu.NamaSatuan = nama_satuan
		}
		return satu
	} else {
		fmt.Println("Id tidak ada")
		return satu
	}
	fmt.Println("Satuan nill karena tidak input id apa apa")

	return satu
}

func DeleteSatuan(data Satuan)Message{
	db :=db.Connect()
	defer db.Close()

	var id_satuan int
	
	id_satuan,_ = strconv.Atoi(data.IdSatuan)
	if FindSatuan(id_satuan) {
		_,err := db.Exec("DELETE FROM satuan WHERE id_satuan=?", id_satuan)

	if err!= nil {
		message:=Message{}
		message.Status=http.StatusBadRequest
		message.Message = "Ada error terkait Delete satuan karena" + err.Error()
		fmt.Println(err)

		return message	
	}

	//result biar kepake doang soalnya eror ntar
	//fmt.Println(result)

	} else {
		message:=Message{}
		message.Status=http.StatusNotFound
		message.Message = "Ada error terkait Delete satuan karena id tidak ada"

		return message	
	}

	message:=Message{}
		message.Status=http.StatusOK
		message.Message = "tidak input apa apa untuk delete satuan"
		return message
	
}


func FindSatuan(id_satuan int) bool{
	db:=db.Connect()
	defer db.Close()

	row := db.QueryRow("SELECT * FROM satuan WHERE id_satuan = ?", id_satuan)

	sat := Satuan{}
	err:= row.Scan(&sat.IdSatuan, &sat.NamaSatuan)

	switch {
	case err == sql.ErrNoRows:
		// log.Fatalf("Barang dari id yang dimasukkan tidak ada di database")
		return false
	case err != nil:
		log.Fatalf("Error saat mencari barang karena : ", err)
		return false
	}

	fmt.Println(sat)
	return true
}
