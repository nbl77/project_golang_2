$(document).ready(function (){
    const stok = $("#stok_barang");
    const jumlah = $("#jumlah");
    const barang = $("#barang");
    const alertStok = $("#alert-stok");
    const form = $("form")
    alertStok.hide()
    barang.on("change",function (){
        let idBarang = $(this).val();
        $.ajax({
            url:`http://localhost:9000/get-single-barang/${idBarang}`,
            method:"get",
            dataType : "json",
            success:function (data){
                let jumStok = data.Stok
                stok.val(jumStok)
                jumlahFunc(data)
            }
        })
    })
    const jumlahFunc = function (data) {
        jumlah.on("keyup", function () {
            alertStok.hide()
            let curr = $(this).val();
            stok.val(data.Stok - curr);
            if (data.Stok - curr < 0) {
                alertStok.show()
                stok.val(0);
            }
        });
    };


})