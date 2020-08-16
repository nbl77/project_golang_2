package Barang

import (
	"Inventory_Project/config"
	"Inventory_Project/cookie_conf"
	"Inventory_Project/db/service"
	"github.com/labstack/echo/v4"
	"log"
	"net/http"
	"strconv"
)
type Supplier struct {
	Id_supplier int
	Nama_supplier string
	Alamat string
	NoTelp string
}

func SupplierMaster(ctx echo.Context) error {
	data_supplier := service.ShowAllSupplier()
	data := &config.M{
		"title": "Supplier",
		"path": "supplier",
		"data":data_supplier,
		"alert":      cookie_conf.CookieExist(ctx, "alert"),
		"alert_data": cookie_conf.GetCookieAlert(ctx),
	}
	ctx.Render(http.StatusOK,"header",data)
	ctx.Render(http.StatusOK,"sidenav",data)
	return ctx.Render(http.StatusOK,"supplier.html",data)
}
func PostSupplier(ctx echo.Context) error {

	data := service.Supplier{
		NamaSupplier: ctx.FormValue("supplier"),
		Alamat:       ctx.FormValue("alamat"),
		NoTelp:       ctx.FormValue("no_telp"),
	}
	_,err := strconv.Atoi(data.NoTelp)
	if err != nil {
		cookie_conf.SetCookieAlert(ctx,"danger","Tidak dapat memproses data!")
		return ctx.Redirect(http.StatusFound,"/supplier")
	}
	if data.NamaSupplier == "" || data.Alamat == "" || data.NoTelp == ""{
		cookie_conf.SetCookieAlert(ctx,"danger","Field Tidak Boleh Kosong!")
		return ctx.Redirect(http.StatusFound,"/supplier")
	}
	mess := service.InsertSupplier(data)
	if mess.Status == http.StatusOK {
		cookie_conf.SetCookieAlert(ctx,"success","Berhasil Menambah Supplier!")
		return ctx.Redirect(http.StatusFound,"/supplier")
	}
	cookie_conf.SetCookieAlert(ctx,"danger","Gagal Menambah Supplier!")
	return ctx.Redirect(http.StatusFound,"/supplier")
}
func DeleteSupplier(ctx echo.Context) error {
	data := service.Supplier{IdSupplier: ctx.Param("id")}
	mess := service.DeleteSupplier(data)
	if mess.Status == http.StatusOK {
		cookie_conf.SetCookieAlert(ctx,"success","Berhasil Menghapus Supplier!")
		log.Println("Menghapus Data Supplier")
		return ctx.Redirect(http.StatusFound,"/supplier")
	}
	cookie_conf.SetCookieAlert(ctx,"danger","Gagal Menghapus Supplier!")
	return ctx.Redirect(http.StatusFound,"/supplier")
}
func ShowEditSupplier(ctx echo.Context) error {
	dataSup := service.Supplier{
		IdSupplier: ctx.Param("id"),
	}
	mess := service.ShowPerSupplier(dataSup)
	if mess.IdSupplier == "" {
		cookie_conf.SetCookieAlert(ctx,"danger","Supplier Tidak ditemukan!")
		return ctx.Redirect(http.StatusFound,"/supplier")
	}
	data := &config.M{
		"title":      "Edit Data Supplier",
		"path":       "barang",
		"alert":      cookie_conf.CookieExist(ctx, "alert"),
		"alert_data": cookie_conf.GetCookieAlert(ctx),
		"supplier":mess,
	}
	ctx.Render(http.StatusOK,"header",data)
	ctx.Render(http.StatusOK,"sidenav",data)
	return ctx.Render(http.StatusOK,"action_supplier.html",data)
}
func EditSupplier(ctx echo.Context) error {
	IdSupplier := ctx.Param("id")
	NamaSupplier := ctx.FormValue("supplier")
	NoTelp := ctx.FormValue("no_telp")
	Alamat := ctx.FormValue("alamat")

	if NamaSupplier == "" || NoTelp == "" || Alamat == ""{
		cookie_conf.SetCookieAlert(ctx,"danger","Field Tidak Boleh Kosong!")
		return ctx.Redirect(http.StatusFound,"/supplier")
	}
	data := service.Supplier{
		IdSupplier: IdSupplier,
		NamaSupplier: NamaSupplier,
		NoTelp: NoTelp,
		Alamat: Alamat,
	}
	mess := service.EditSupplier(data)
	if mess.Status == http.StatusOK{
		cookie_conf.SetCookieAlert(ctx,"success","Berhasil Mengubah Supplier!")
		log.Println("Mengubah Data Supplier")
		return ctx.Redirect(http.StatusFound,"/supplier")
	}
	cookie_conf.SetCookieAlert(ctx,"danger","Gagal Mengubah Supplier!")
	return ctx.Redirect(http.StatusFound,"/supplier")
}