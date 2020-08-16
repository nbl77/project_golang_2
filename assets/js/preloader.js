document.addEventListener("DOMContentLoaded",function (){
    setTimeout(function (){
        $(".loader").fadeOut("slow")
    },1000)
    let index = $(".index");
    index.each(function (){
        let num = parseInt($(this).text());
        $(this).html(num + 1)
    })
})