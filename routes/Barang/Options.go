package Barang

import (
	"Inventory_Project/config"
	"Inventory_Project/cookie_conf"
	db2 "Inventory_Project/db"
	"Inventory_Project/db/service"
	"github.com/labstack/echo/v4"
	"log"
	"net/http"
)
type Message struct{
	Status interface{}
	Message string
}
func Options(ctx echo.Context) error {
	rowsSatuan := db2.SelectAll("SELECT*FROM satuan")
	rowskategori := db2.SelectAll("SELECT kategori.id_kategori,kategori.nama_kategori,satuan.nama_satuan FROM kategori LEFT JOIN satuan ON kategori.id_satuan=satuan.id_satuan")
	var resultSatuan []Satuan
	var resultKategori []Kategori
	for rowsSatuan.Next(){
		var sat = new(Satuan)
		rowsSatuan.Scan(&sat.IdSatuan,&sat.NamaSatuan)
		resultSatuan = append(resultSatuan,*sat)
	}
	for rowskategori.Next(){
		var kat = new(Kategori)
		rowskategori.Scan(&kat.IdKategori,&kat.NamaKategori,&kat.Satuan)
		resultKategori = append(resultKategori,*kat)
	}
	rowsSatuan.Close()
	rowskategori.Close()
	data := &config.M{
		"title":"Options",
		"path": "barang",
		"satuan": resultSatuan,
		"kategori": resultKategori,
		"alert":      cookie_conf.CookieExist(ctx, "alert"),
		"alert_data": cookie_conf.GetCookieAlert(ctx),
	}
	ctx.Render(http.StatusOK,"header",data)
	ctx.Render(http.StatusOK,"sidenav",data)
	return ctx.Render(http.StatusOK,"options.html",data)
}
func PostSatuan(ctx echo.Context) error {
	data := service.Satuan{
		NamaSatuan: ctx.FormValue("satuan"),
	}
	if data.NamaSatuan == "" {
		cookie_conf.SetCookieAlert(ctx,"danger","Field Tidak Boleh Kosong!")
		return ctx.Redirect(http.StatusFound,"/options")
	}
	mess := service.InsertSatuan(data)
	if mess.Status == http.StatusOK {
		cookie_conf.SetCookieAlert(ctx,"success","Berhasil Menyimpan Satuan!")
		return ctx.Redirect(http.StatusFound,"/options")
	}
	cookie_conf.SetCookieAlert(ctx,"danger","Gagal Menyimpan Satuan!")
	return ctx.Redirect(http.StatusFound,"/options")
}
func PostKategori(ctx echo.Context) error {
	data := service.Kategori{
		NamaKategori: ctx.FormValue("kategori"),
		Satuan: ctx.FormValue("satuan"),
	}
	if data.NamaKategori == "" || data.Satuan == "" {
		cookie_conf.SetCookieAlert(ctx,"danger","Field Tidak Boleh Kosong!")
		return ctx.Redirect(http.StatusFound,"/options")
	}
	mess := service.InsertKategori(data)
	if mess.Status == http.StatusOK {
		cookie_conf.SetCookieAlert(ctx,"success","Berhasil Menyimpan Kategori!")
		return ctx.Redirect(http.StatusFound,"/options")
	}
	cookie_conf.SetCookieAlert(ctx,"danger","Gagal Menyimpan Kategori!")
	return ctx.Redirect(http.StatusFound,"/options")
}


func DeleteKategori(ctx echo.Context) error {
	data := service.Kategori{
		IdKategori: ctx.Param("id"),
	}
	mess := service.DeleteKategori(data)
	if mess.Status == http.StatusOK {
		cookie_conf.SetCookieAlert(ctx,"success", "Berhasil Menghapus Kategori")
		return ctx.Redirect(http.StatusFound,"/options")
	}
	cookie_conf.SetCookieAlert(ctx,"danger", "Gagal Menghapus Kategori")
	return ctx.Redirect(http.StatusFound,"/options")
}
func DeleteSatuan(ctx echo.Context) error {
	data := service.Satuan{
		IdSatuan: ctx.Param("id"),
	}
	mess := service.DeleteSatuan(data)
	if mess.Status == http.StatusOK {
		cookie_conf.SetCookieAlert(ctx,"success", "Berhasil Menghapus Satuan")
		return ctx.Redirect(http.StatusFound,"/options")
	}
	cookie_conf.SetCookieAlert(ctx,"danger", "Gagal Menghapus Satuan")
	return ctx.Redirect(http.StatusFound,"/options")
}
func ShowEditKategori(ctx echo.Context) error {
	IdKategori := ctx.Param("id")
	rowKategori := db2.SelectParam("SELECT*FROM kategori WHERE id_kategori=?",IdKategori)
	rowsSatuan := db2.SelectAll("SELECT*FROM satuan")
	var (
		resultKategori Kategori
		resultSatuan []Satuan
	)
	for rowsSatuan.Next(){
		var sat = new(Satuan)
		rowsSatuan.Scan(&sat.IdSatuan,&sat.NamaSatuan)
		resultSatuan = append(resultSatuan,*sat)
	}
	rowKategori.Scan(&resultKategori.IdKategori,&resultKategori.Satuan,&resultKategori.NamaKategori)
	defer rowsSatuan.Close()
	if resultKategori.IdKategori == "" {
		cookie_conf.SetCookieAlert(ctx,"danger","Kategori Tidak Di Temukan!")
		return ctx.Redirect(http.StatusFound,"/options")
	}
	data := &config.M{
		"title":      "Edit Data Barang",
		"path":       "barang",
		"alert":      cookie_conf.CookieExist(ctx, "alert"),
		"alert_data": cookie_conf.GetCookieAlert(ctx),
		"kategori":resultKategori,
		"satuan":resultSatuan,
	}
	ctx.Render(http.StatusOK,"header",data)
	ctx.Render(http.StatusOK,"sidenav",data)
	return ctx.Render(http.StatusOK,"action_kategori.html",data)
}
func EditKategori(ctx echo.Context) error {
	idKategori := ctx.Param("id")
	kategori := ctx.FormValue("kategori")
	satuan := ctx.FormValue("satuan")
	if kategori == "" || satuan == "" {
		cookie_conf.SetCookieAlert(ctx,"danger","Field Tidak Boleh Kosong!")
		return ctx.Redirect(http.StatusFound,"/options")
	}
	data := service.Kategori{
		IdKategori: idKategori,
		NamaKategori: kategori,
		Satuan: satuan,
	}
	mess := service.EditKategori(data)
	if mess.Status == http.StatusOK {
		cookie_conf.SetCookieAlert(ctx,"success","Berhasil Mengubah Kategori!")
		log.Println("Mengubah Data Kategori")
		return ctx.Redirect(http.StatusFound,"/options")
	}
	cookie_conf.SetCookieAlert(ctx,"danger","Berhasil Gagal Kategori!")
	return ctx.Redirect(http.StatusFound,"/options")
}
func ShowEditSatuan(ctx echo.Context) error {
	IdSatuan := ctx.Param("id")
	rowSatuan := db2.SelectParam("SELECT*FROM satuan WHERE id_satuan=?",IdSatuan)
	var resultSatuan Satuan
	rowSatuan.Scan(&resultSatuan.IdSatuan,&resultSatuan.NamaSatuan)
	if resultSatuan.IdSatuan == "" {
		cookie_conf.SetCookieAlert(ctx,"danger","Satuan Tidak Di Temukan!")
		return ctx.Redirect(http.StatusFound,"/options")
	}
	data := &config.M{
		"title":      "Edit Data Barang",
		"path":       "barang",
		"alert":      cookie_conf.CookieExist(ctx, "alert"),
		"alert_data": cookie_conf.GetCookieAlert(ctx),
		"satuan":resultSatuan,
	}
	ctx.Render(http.StatusOK,"header",data)
	ctx.Render(http.StatusOK,"sidenav",data)
	return ctx.Render(http.StatusOK,"action_satuan.html",data)
}
func EditSatuan(ctx echo.Context) error {
	IdSatuan := ctx.Param("id")
	Satuan := ctx.FormValue("satuan")
	if Satuan == ""{
		cookie_conf.SetCookieAlert(ctx,"danger","Field Tidak Boleh Kosong!")
		return ctx.Redirect(http.StatusFound,"/options")
	}
	data := service.Satuan{
		IdSatuan: IdSatuan,
		NamaSatuan: Satuan,
	}
	mess := service.EditSatuan(data)
	if mess.Status == http.StatusOK {
		cookie_conf.SetCookieAlert(ctx,"success","Berhasil Mengubah Satuan!")
		log.Println("Mengubah Data Satuan")
		return ctx.Redirect(http.StatusFound,"/options")
	}
	cookie_conf.SetCookieAlert(ctx,"danger","Gagal Mengubah Kategori!")
	return ctx.Redirect(http.StatusFound,"/options")
}