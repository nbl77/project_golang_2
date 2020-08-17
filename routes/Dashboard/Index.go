package Dashboard

import (
	"Inventory_Project/config"
	"Inventory_Project/db"
	"github.com/labstack/echo/v4"
	"net/http"
)

func Home(ctx echo.Context) error {
	var totalBarang int
	var totalBarangMasuk int
	var totalBarangKeluar int
	rowBarang := db.Select("SELECT SUM(stok) FROM barang")
	rowBarangMasuk := db.Select("SELECT SUM(jumlah_masuk) FROM barang_masuk")
	rowBarangKeluar := db.Select("SELECT SUM(jumlah_keluar) FROM barang_keluar")
	rowBarang.Scan(&totalBarang)
	rowBarangMasuk.Scan(&totalBarangMasuk)
	rowBarangKeluar.Scan(&totalBarangKeluar)
	data := &config.M{
		"title": "Dashboard",
		"path":  "dashboard",
		"total_barang": totalBarang,
		"total_barang_masuk": totalBarangMasuk,
		"total_barang_keluar": totalBarangKeluar,
	}
	ctx.Render(http.StatusOK,"header",data)
	ctx.Render(http.StatusOK,"sidenav",data)
	return ctx.Render(http.StatusOK,"index.html",data)
}