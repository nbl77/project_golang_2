package Barang

import (
	"Inventory_Project/config"
	"Inventory_Project/cookie_conf"
	"Inventory_Project/db"
	"github.com/labstack/echo/v4"
	"log"
	"net/http"
)
type Barang struct {
	IdBarang string
	NamaBarang string
	Stok int
	Kategori string
	Satuan string
}
type Kategori struct {
	IdKategori string
	NamaKategori string
	Satuan string
}
type Satuan struct {
	IdSatuan string
	NamaSatuan string
}
func ShowMaster(ctx echo.Context) error {
	rowsBarang := db.SelectAll("SELECT barang.id_barang,barang.nama_barang,barang.stok,kategori.nama_kategori,satuan.nama_satuan FROM barang LEFT JOIN kategori ON barang.id_kategori=kategori.id_kategori LEFT JOIN satuan ON kategori.id_satuan=satuan.id_satuan")
	rowsKategori := db.SelectAll("SELECT id_kategori,nama_kategori FROM kategori")
	rowsSatuan := db.SelectAll("SELECT*FROM satuan")
	var (
		resultBarang []Barang
		resultKategori []Kategori
		resultSatuan []Satuan
	)

	for rowsBarang.Next(){
		var barang = new(Barang)
		rowsBarang.Scan(&barang.IdBarang,&barang.NamaBarang,&barang.Stok,&barang.Kategori,&barang.Satuan)
		resultBarang = append(resultBarang,*barang)
	}
	for rowsKategori.Next(){
		var kat = new(Kategori)
		rowsKategori.Scan(&kat.IdKategori,&kat.NamaKategori)
		resultKategori = append(resultKategori,*kat)
	}
	for rowsSatuan.Next(){
		var sat = new(Satuan)
		rowsSatuan.Scan(&sat.IdSatuan,&sat.NamaSatuan)
		resultSatuan = append(resultSatuan,*sat)
	}
	rowsBarang.Close()
	rowsKategori.Close()
	rowsSatuan.Close()
	data := &config.M{
		"title":      "Data Barang",
		"path":       "barang",
		"data":       resultBarang,
		"alert":      cookie_conf.CookieExist(ctx, "alert"),
		"alert_data": cookie_conf.GetCookieAlert(ctx),
		"kategori":resultKategori,
		"satuan":resultSatuan,
	}

	ctx.Render(http.StatusOK,"header",data)
	ctx.Render(http.StatusOK,"sidenav",data)
	return ctx.Render(http.StatusOK,"barang.html",data)
}

func ShowEditBarang(ctx echo.Context) error {
	IdBarang := ctx.Param("id")
	rowBarang := db.SelectParam("SELECT * FROM barang WHERE id_barang=?",IdBarang)
	rowsKategori := db.SelectAll("SELECT id_kategori,nama_kategori FROM kategori")
	rowsSatuan := db.SelectAll("SELECT*FROM satuan")
	var (
		resultKategori []Kategori
		resultSatuan []Satuan
		resultBarang Barang
	)
	for rowsKategori.Next(){
		var kat = new(Kategori)
		rowsKategori.Scan(&kat.IdKategori,&kat.NamaKategori)
		resultKategori = append(resultKategori,*kat)
	}
	for rowsSatuan.Next(){
		var sat = new(Satuan)
		rowsSatuan.Scan(&sat.IdSatuan,&sat.NamaSatuan)
		resultSatuan = append(resultSatuan,*sat)
	}
	rowBarang.Scan(&resultBarang.IdBarang,&resultBarang.NamaBarang,&resultBarang.Stok,&resultBarang.Kategori)
	if resultBarang.IdBarang == "" {
		cookie_conf.SetCookieAlert(ctx,"danger","Barang Tidak Di Temukan!")
		return ctx.Redirect(http.StatusFound,"/barang")
	}
	rowsKategori.Close()
	rowsSatuan.Close()
	data := &config.M{
		"title":      "Edit Data Barang",
		"path":       "barang",
		"data":       resultBarang,
		"alert":      cookie_conf.CookieExist(ctx, "alert"),
		"alert_data": cookie_conf.GetCookieAlert(ctx),
		"kategori":resultKategori,
		"satuan":resultSatuan,
	}
	ctx.Render(http.StatusOK,"header",data)
	ctx.Render(http.StatusOK,"sidenav",data)
	return ctx.Render(http.StatusOK,"action_barang.html",data)
}
func EditBarang(ctx echo.Context) error {
	IdBarang := ctx.Param("id")
	NamaBarang := ctx.FormValue("barang")
	Kategori := ctx.FormValue("kategori")
	data := []interface{}{
		NamaBarang,
		Kategori,
		IdBarang,
	}
	res,err := db.Execute("UPDATE barang SET nama_barang=?,id_kategori=? WHERE id_barang=?",data)
	if err != nil {
		log.Println(err)
	}
	if res > 0 {
		cookie_conf.SetCookieAlert(ctx,"success","Berhasil Mengubah data!")
		log.Println("Mengubah Data Barang")
		return ctx.Redirect(http.StatusFound,"/barang")
	}
	cookie_conf.SetCookieAlert(ctx,"danger","Gagal Mengubah data!")
	return ctx.Redirect(http.StatusFound,"/barang")
 return nil
}

func PostBarang(ctx echo.Context) error {
	barang := ctx.FormValue("barang")
	kategori := ctx.FormValue("kategori")
	if barang == "" || kategori == "" {
		cookie_conf.SetCookieAlert(ctx,"danger","Field Tidak Boleh Kosong!")
		return ctx.Redirect(http.StatusFound,"/barang")
	}
	var id_barang int
	row := db.Select("SELECT count(*) FROM barang")
	row.Scan(&id_barang)
	var data = []interface{} {
		id_barang + 1,
		barang,
		0,
		kategori,
	}
	res,err := db.Execute("INSERT INTO barang(id_barang,nama_barang,stok,id_kategori) VALUES(?,?,?,?)",data)
	if err != nil {
		log.Println(err)
	}
	if res > 0 {
		cookie_conf.SetCookieAlert(ctx,"success","Berhasil Menyimpan data!")
		log.Println("Menambahkan Data Barang")
		return ctx.Redirect(http.StatusFound,"/barang")
	}
	cookie_conf.SetCookieAlert(ctx,"danger","Gagal Menyimpan data!")
	return ctx.Redirect(http.StatusFound,"/barang")
}
func DeleteBarang(ctx echo.Context) error {
	res,err := db.Execute("DELETE FROM barang WHERE id_barang=?",ctx.Param("id"))
	if err != nil {
		log.Println(err)
	}
	if res > 0 {
		res, err = db.Execute("UPDATE barang SET id_barang = id_barang - 1 WHERE id_barang > ?",ctx.Param("id"))
		cookie_conf.SetCookieAlert(ctx,"success","Berhasil Menghapus data!")
		log.Println("Menghapus Data Barang")
		return ctx.Redirect(http.StatusFound,"/barang")
	}
	cookie_conf.SetCookieAlert(ctx,"danger","Gagal Menghapus data!")
	return ctx.Redirect(http.StatusFound,"/barang")
}