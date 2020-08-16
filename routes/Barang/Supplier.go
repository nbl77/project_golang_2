package Barang

import (
	"Inventory_Project/config"
	"Inventory_Project/cookie_conf"
	db2 "Inventory_Project/db"
	"fmt"
	"github.com/labstack/echo/v4"
	"net/http"
)
type Supplier struct {
	Id_supplier int
	Nama_supplier string
	Alamat string
	NoTelp string
}

func SupplierMaster(ctx echo.Context) error {
	data := &config.M{
		"title": "Supplier",
		"path": "supplier",
		"alert":      cookie_conf.CookieExist(ctx, "alert"),
		"alert_data": cookie_conf.GetCookieAlert(ctx),
	}
	ctx.Render(http.StatusOK,"header",data)
	ctx.Render(http.StatusOK,"sidenav",data)
	return ctx.Render(http.StatusOK,"supplier.html",data)
}
func PostSupplier(ctx echo.Context) error {

	data := Supplier{
		Nama_supplier: ctx.FormValue("supplier"),
		Alamat:        ctx.FormValue("alamat"),
		NoTelp: ctx.FormValue("no_telp"),
	}
	if data.Nama_supplier == "" || data.Alamat == "" || data.NoTelp == ""{
		cookie_conf.SetCookieAlert(ctx,"danger","Field Tidak Boleh Kosong!")
		return ctx.Redirect(http.StatusFound,"/supplier")
	}
	mess := InsertSupplier(data)
	if mess.Status == http.StatusOK {
		cookie_conf.SetCookieAlert(ctx,"success","Berhasil Menambah Supplier!")
		return ctx.Redirect(http.StatusFound,"/supplier")
	}
	cookie_conf.SetCookieAlert(ctx,"danger","Gagal Menambah Supplier!")
	return ctx.Redirect(http.StatusFound,"/supplier")
}
func InsertSupplier(data Supplier)Message{
	db:= db2.Connect()
	defer db.Close()
	fmt.Println("Tambah Supplier")

	var(
		nama_supplier string
		alamat string
		no_telp string
	)
	nama_supplier = data.Nama_supplier
	alamat = data.Alamat
	no_telp = data.NoTelp
	insert,err := db.Prepare("INSERT INTO supplier(nama_supplier, alamat, no_telp) VALUES(?,?,?)")
	if err != nil {
		message:=Message{}
		message.Status=string(http.StatusBadRequest)
		message.Message = "Ada error terkait Insert supplier"
		fmt.Println(err)
		return message
	}

	insert.Exec(nama_supplier, alamat, no_telp)

	message:=Message{}
	message.Status=http.StatusOK
	message.Message = "Sukses insert supplier"
	fmt.Println(err)

	return message

}