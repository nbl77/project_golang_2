package Barang

import (
	"Inventory_Project/config"
	"Inventory_Project/cookie_conf"
	"Inventory_Project/db"
	"Inventory_Project/db/service"
	"github.com/labstack/echo/v4"
	"log"
	"net/http"
	"strconv"
)

type BarangKeluar struct {
	IdBarangKeluar string
	NamaBarang string
	Alamat string
	JumlahKeluar string
	WaktuKeluar string
}
func Keluar(ctx echo.Context) error {
	rowsBarangKeluar := db.SelectAll("SELECT barang_keluar.id_barang_keluar, barang.nama_barang,barang_keluar.alamat,barang_keluar.jumlah_keluar, barang_keluar.waktu_keluar FROM barang_keluar LEFT JOIN barang ON barang_keluar.id_barang=barang.id_barang")
	rowsBarang := db.SelectAll("SELECT id_barang,nama_barang FROM barang")
	var (
		barangKeluarList []BarangKeluar
		barangList []Barang
	)
	for rowsBarangKeluar.Next() {
		var barangKeluar = new(BarangKeluar)
		rowsBarangKeluar.Scan(&barangKeluar.IdBarangKeluar,&barangKeluar.NamaBarang,&barangKeluar.Alamat,&barangKeluar.JumlahKeluar,&barangKeluar.WaktuKeluar)
		barangKeluarList = append(barangKeluarList,*barangKeluar)
	}
	for rowsBarang.Next(){
		var barang = new(Barang)
		rowsBarang.Scan(&barang.IdBarang,&barang.NamaBarang)
		barangList = append(barangList,*barang)
	}
	rowsBarangKeluar.Close()
	rowsBarang.Close()
	data := &config.M{
		"title": "Barang Keluar",
		"path": "barang-keluar",
		"alert":      cookie_conf.CookieExist(ctx, "alert"),
		"alert_data": cookie_conf.GetCookieAlert(ctx),
		"data":barangKeluarList,
		"barang":barangList,
	}
	ctx.Render(http.StatusOK,"header",data)
	ctx.Render(http.StatusOK,"sidenav",data)
	return ctx.Render(http.StatusOK,"barang_keluar.html",data)
}
func PostbarangKeluar(ctx echo.Context) error {
	data := service.BarangKeluar{
		IdBarang:     ctx.FormValue("barang"),
		Alamat:       ctx.FormValue("tujuan"),
		JumlahKeluar: ctx.FormValue("jumlah"),
		WaktuKeluar:  ctx.FormValue("tanggal_keluar"),
	}
	var currStok int
	jmlOut,_ := strconv.Atoi(data.JumlahKeluar)
	rowsBarangKeluar := db.SelectParam("SELECT stok FROM barang WHERE id_barang=?",data.IdBarang)
	rowsBarangKeluar.Scan(&currStok)
	if jmlOut > currStok {
		cookie_conf.SetCookieAlert(ctx,"danger","Jumlah Keluar Melebihi total!")
		return ctx.Redirect(http.StatusFound,"/barang-keluar")
	}
	mess := service.InsertBarangKeluar(data)
	if mess.Status == http.StatusOK {
		db.Execute("UPDATE barang SET stok= stok - ? WHERE id_barang=?",data.JumlahKeluar,data.IdBarang)
		cookie_conf.SetCookieAlert(ctx,"success","Berhasil Menambahkan barang keluar!")
		log.Println("Menambah data barang keluar")
		return ctx.Redirect(http.StatusFound,"/barang-keluar")
	}
	cookie_conf.SetCookieAlert(ctx,"danger","Gagal Menambahkan barang keluar!")
	return ctx.Redirect(http.StatusFound,"/barang-keluar")
}
func DeleteBarangKeluar(ctx echo.Context) error {
	var currJmlKeluar int
	var idBarang int
	data := service.BarangKeluar{IdBarangKeluar: ctx.Param("id")}
	rowBarangKeluar := db.SelectParam("SELECT jumlah_keluar,id_barang FROM barang_keluar WHERE id_barang_keluar=?",data.IdBarangKeluar)
	rowBarangKeluar.Scan(&currJmlKeluar,&idBarang)
	mess := service.DeleteBarangKeluar(data)
	if mess.Status == http.StatusOK {
		db.Execute("UPDATE barang SET stok=stok + ? WHERE id_barang=?",currJmlKeluar,idBarang)
		cookie_conf.SetCookieAlert(ctx,"success","Berhasil Menghapus barang keluar!")
		log.Println("Menghapus data barang keluar")
		return ctx.Redirect(http.StatusFound,"/barang-keluar")
	}
	cookie_conf.SetCookieAlert(ctx,"danger","Gagal Menghapus barang keluar!")
	return ctx.Redirect(http.StatusFound,"/barang-keluar")
}
func ShowEditBarangKeluar(ctx echo.Context) error {
	rowBarangKeluar := db.SelectParam("SELECT * FROM barang_keluar WHERE id_barang_keluar=?",ctx.Param("id"))
	rowsBarang := db.SelectAll("SELECT id_barang,nama_barang FROM barang")
	var (
		barangKeluar service.BarangKeluar
		barangList []Barang
	)
	rowBarangKeluar.Scan(&barangKeluar.IdBarangKeluar,&barangKeluar.IdBarang,&barangKeluar.Alamat,&barangKeluar.JumlahKeluar,&barangKeluar.WaktuKeluar)
	for rowsBarang.Next(){
		var barang = new(Barang)
		rowsBarang.Scan(&barang.IdBarang,&barang.NamaBarang)
		barangList = append(barangList,*barang)
	}
	data := &config.M{
		"title": "Edit Barang Keluar",
		"path": "barang-keluar",
		"alert":      cookie_conf.CookieExist(ctx, "alert"),
		"alert_data": cookie_conf.GetCookieAlert(ctx),
		"data":barangKeluar,
		"barang":barangList,
	}
	ctx.Render(http.StatusOK,"header",data)
	ctx.Render(http.StatusOK,"sidenav",data)
	return ctx.Render(http.StatusOK,"action_barang_keluar.html",data)
}
func EditBarangKeluar(ctx echo.Context) error {
	data := service.BarangKeluar{
		IdBarangKeluar: ctx.Param("id"),
		IdBarang:       ctx.FormValue("barang"),
		JumlahKeluar:   ctx.FormValue("jumlah"),
		Alamat:         ctx.FormValue("tujuan"),
		WaktuKeluar: ctx.FormValue("tanggal_keluar"),
	}
	var JmlSkrg int
	var stok int
	rowBarangMasuk := db.SelectParam("SELECT jumlah_keluar FROM barang_keluar WHERE id_barang_keluar=?",data.IdBarangKeluar)
	rowBarangMasuk.Scan(&JmlSkrg)
	newJml,_ := strconv.Atoi(data.JumlahKeluar)
	rowBarang := db.SelectParam("SELECT stok FROM barang WHERE id_barang=?",data.IdBarang)
	rowBarang.Scan(&stok)
	selisih := newJml - JmlSkrg
	if stok - selisih < 0  {
		cookie_conf.SetCookieAlert(ctx,"danger","Jumlah Barang Kurang!")
		return ctx.Redirect(http.StatusFound,"/barang-keluar")
	}
	mess := service.EditBarangKeluar(data)
	if mess.Status == http.StatusOK {
		db.Execute("UPDATE barang SET stok=stok - ? WHERE id_barang=?",selisih,data.IdBarang)
		cookie_conf.SetCookieAlert(ctx,"success","Berhasil Mengubah barang keluar!")
		log.Println("Mengubah data barang keluar")
		return ctx.Redirect(http.StatusFound,"/barang-keluar")
	}
	cookie_conf.SetCookieAlert(ctx,"danger","Data tidak ditemukan!")
	return ctx.Redirect(http.StatusFound,"/barang-keluar")
}