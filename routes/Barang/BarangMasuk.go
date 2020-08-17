package Barang

import (
	"Inventory_Project/config"
	"Inventory_Project/cookie_conf"
	"Inventory_Project/db"
	"Inventory_Project/db/service"
	"database/sql"
	"github.com/labstack/echo/v4"
	"log"
	"net/http"
	"strconv"
)

type BarangMasuk struct {
	IdBarangMasuk string
	NamaBarang      string
	NamaSupplier    sql.NullString
	JumlahMasuk   string
	WaktuMasuk    string
}
func Masuk(ctx echo.Context) error {
	rowsBarangMasuk := db.SelectAll("SELECT barang_masuk.id_barang_masuk,barang.nama_barang,supplier.nama_supplier,barang_masuk.jumlah_masuk,barang_masuk.waktu_masuk FROM barang_masuk LEFT JOIN barang ON barang_masuk.id_barang=barang.id_barang LEFT JOIN supplier ON barang_masuk.id_supplier=supplier.id_supplier ORDER BY barang_masuk.id_barang_masuk DESC ")
	rowsBarang := db.SelectAll("SELECT id_barang,nama_barang FROM barang")
	rowsSupplier := db.SelectAll("SELECT id_supplier,nama_supplier FROM supplier")
	var (
		barangMasuk []BarangMasuk
		barangList []Barang
		supplierList []Supplier
	)
	for rowsBarangMasuk.Next(){
		var BarangMasuk = new(BarangMasuk)
		rowsBarangMasuk.Scan(&BarangMasuk.IdBarangMasuk,&BarangMasuk.NamaBarang,&BarangMasuk.NamaSupplier,&BarangMasuk.JumlahMasuk,&BarangMasuk.WaktuMasuk)
		barangMasuk = append(barangMasuk,*BarangMasuk)
	}
	for rowsBarang.Next(){
		var barang = new(Barang)
		rowsBarang.Scan(&barang.IdBarang,&barang.NamaBarang)
		barangList = append(barangList,*barang)
	}
	for rowsSupplier.Next(){
		var supplier = new(Supplier)
		rowsSupplier.Scan(&supplier.IdSupplier,&supplier.NamaSupplier)
		supplierList = append(supplierList,*supplier)
	}
	rowsSupplier.Close()
	rowsBarang.Close()
	rowsBarangMasuk.Close()

	data := &config.M{
		"title":      "Barang Masuk",
		"path":       "barang-masuk",
		"alert":      cookie_conf.CookieExist(ctx, "alert"),
		"alert_data": cookie_conf.GetCookieAlert(ctx),
		"data":barangMasuk,
		"barang":barangList,
		"supplier":supplierList,
	}
	ctx.Render(http.StatusOK,"header",data)
	ctx.Render(http.StatusOK,"sidenav",data)
	return ctx.Render(http.StatusOK,"barang_masuk.html",data)
}
func PostbarangMasuk(ctx echo.Context) error {
	data := service.BarangMasuk{
		IdBarang:    ctx.FormValue("barang"),
		WaktuMasuk:  ctx.FormValue("tanggal_masuk"),
		IdSupplier:  ctx.FormValue("supplier"),
		JumlahMasuk: ctx.FormValue("jumlah"),
	}
	mess := service.InsertBarangMasuk(data)
	if mess.Status == http.StatusOK {
		db.Execute("UPDATE barang SET stok= ? + stok WHERE id_barang=?",data.JumlahMasuk,data.IdBarang)
		cookie_conf.SetCookieAlert(ctx,"success","Berhasil Menambahkan barang!")
		log.Println("Menambah data barang masuk")
		return ctx.Redirect(http.StatusFound,"/barang-masuk")
	}
	cookie_conf.SetCookieAlert(ctx,"danger","Gagal Menambahkan barang!")
	return ctx.Redirect(http.StatusFound,"/barang-masuk")
}
func DeleteBarangMasuk(ctx echo.Context) error {
	var currJmlMasuk int
	var idBarang int
	var stok int
	IdBarangMasuk := service.BarangMasuk{IdBarangMasuk: ctx.Param("id")}
	rowBarangKeluar := db.SelectParam("SELECT jumlah_masuk,id_barang FROM barang_masuk WHERE id_barang_masuk=?",IdBarangMasuk.IdBarangMasuk)
	rowBarangKeluar.Scan(&currJmlMasuk,&idBarang)
	rowBarang := db.SelectParam("SELECT stok FROM barang WHERE id_barang=?", idBarang)
	rowBarang.Scan(&stok)
	if stok < currJmlMasuk {
		cookie_conf.SetCookieAlert(ctx,"danger","Jumlah Barang Kurang!")
		return ctx.Redirect(http.StatusFound,"/barang-masuk")
	}
	mess := service.DeleteBarangMasuk(IdBarangMasuk)
	if mess.Status == http.StatusOK {
		db.Execute("UPDATE barang SET stok=stok - ? WHERE id_barang=?",currJmlMasuk,idBarang)
		cookie_conf.SetCookieAlert(ctx,"success","Berhasil Menghapus barang!")
		log.Println("Menghapus data barang masuk")
		return ctx.Redirect(http.StatusFound,"/barang-masuk")
	}
	cookie_conf.SetCookieAlert(ctx,"danger","Data Barang Masuk tidak ditemukan!")
	return ctx.Redirect(http.StatusFound,"/barang-masuk")
}
func ShowEditBarangMasuk(ctx echo.Context) error {
	rowsBarangMasuk := db.SelectParam("SELECT*FROM barang_masuk WHERE id_barang_masuk=?",ctx.Param("id"))
	rowsBarang := db.SelectAll("SELECT id_barang,nama_barang FROM barang")
	rowsSupplier := db.SelectAll("SELECT id_supplier,nama_supplier FROM supplier")
	var (
		barangMasuk service.BarangMasuk
		barangList []Barang
		supplierList []Supplier
	)
	rowsBarangMasuk.Scan(&barangMasuk.IdBarangMasuk,&barangMasuk.IdBarang,&barangMasuk.IdSupplier,&barangMasuk.JumlahMasuk,&barangMasuk.WaktuMasuk)
	for rowsBarang.Next(){
		var barang = new(Barang)
		rowsBarang.Scan(&barang.IdBarang,&barang.NamaBarang)
		barangList = append(barangList,*barang)
	}
	for rowsSupplier.Next(){
		var supplier = new(Supplier)
		rowsSupplier.Scan(&supplier.IdSupplier,&supplier.NamaSupplier)
		supplierList = append(supplierList,*supplier)
	}
	rowsSupplier.Close()
	rowsBarang.Close()
	if barangMasuk.IdBarangMasuk == ""{
		cookie_conf.SetCookieAlert(ctx,"danger","Data Tidak Di Temukan!")
		return ctx.Redirect(http.StatusFound,"/barang-masuk")
	}
	data := &config.M{
		"title":      "Barang Masuk",
		"path":       "barang-masuk",
		"alert":      cookie_conf.CookieExist(ctx, "alert"),
		"alert_data": cookie_conf.GetCookieAlert(ctx),
		"data":barangMasuk,
		"barang":barangList,
		"supplier":supplierList,
	}
	ctx.Render(http.StatusOK,"header",data)
	ctx.Render(http.StatusOK,"sidenav",data)
	return ctx.Render(http.StatusOK,"action_barang_masuk.html",data)
}
func EditBarangMasuk(ctx echo.Context) error {
	data := service.BarangMasuk{
		IdBarangMasuk: ctx.Param("id"),
		IdSupplier:    ctx.FormValue("supplier"),
		IdBarang:      ctx.FormValue("barang"),
		JumlahMasuk:   ctx.FormValue("jumlah"),
		WaktuMasuk: ctx.FormValue("tanggal_masuk"),
	}
	var JmlSkrg int
	var stok int
	rowBarangMasuk := db.SelectParam("SELECT jumlah_masuk FROM barang_masuk WHERE id_barang_masuk=?",data.IdBarangMasuk)
	rowBarangMasuk.Scan(&JmlSkrg)
	newJml,_ := strconv.Atoi(data.JumlahMasuk)
	rowBarang := db.SelectParam("SELECT stok FROM barang WHERE id_barang=?",data.IdBarang)
	rowBarang.Scan(&stok)
	selisih := newJml - JmlSkrg
	if stok + selisih < 0  {
		cookie_conf.SetCookieAlert(ctx,"danger","Jumlah Barang Kurang!")
		return ctx.Redirect(http.StatusFound,"/barang-masuk")
	}
	mess := service.EditBarangMasuk(data)
	if mess.Status == http.StatusOK {

		db.Execute("UPDATE barang SET stok=stok + ? WHERE id_barang=?",selisih,data.IdBarang)
		cookie_conf.SetCookieAlert(ctx,"success","Berhasil Mengubah barang!")
		log.Println("Mengubah data barang masuk")
		return ctx.Redirect(http.StatusFound,"/barang-masuk")
	}
	cookie_conf.SetCookieAlert(ctx,"danger","Data Barang Tidak Di temukan!")
	return ctx.Redirect(http.StatusFound,"/barang-masuk")

}