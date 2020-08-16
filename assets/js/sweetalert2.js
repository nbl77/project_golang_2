
(function($) {
    'use strict'

    $(function() {
        $('[data-toggle="sweet-alert"]').on('click', function(){
            event.preventDefault()
            var type = $(this).data('sweet-alert');
            switch (type) {
                case 'basic':
                    swal({
                        title: $(this).data("title"),
                        text: $(this).data("message"),
                        buttonsStyling: false,
                        confirmButtonClass: 'btn btn-primary'
                    })
                break;

                case 'info':
                    swal({
                        title: $(this).data("title"),
                        text: $(this).data("message"),
                        type: 'info',
                        buttonsStyling: false,
                        confirmButtonClass: 'btn btn-info'
                    })
                break;

                case 'info':
                    swal({
                        title: $(this).data("title"),
                        text: $(this).data("message"),
                        type: 'info',
                        buttonsStyling: false,
                        confirmButtonClass: 'btn btn-info'
                    })
                break;

                case 'success':
                    swal({
                        title: $(this).data("title"),
                        text: $(this).data("message"),
                        type: 'success',
                        buttonsStyling: false,
                        confirmButtonClass: 'btn btn-success'
                    })
                break;

                case 'warning':
                    swal({
                        title: $(this).data("title"),
                        text: $(this).data("message"),
                        type: 'warning',
                        buttonsStyling: false,
                        confirmButtonClass: 'btn btn-warning'
                    })
                break;

                case 'question':
                    swal({
                        title: $(this).data("title"),
                        text: $(this).data("message"),
                        type: 'question',
                        buttonsStyling: false,
                        confirmButtonClass: 'btn btn-default'
                    })
                break;

                case 'confirm':

                    swal({
                        title: $(this).data("title"),
                        text: $(this).data("message"),
                        type: 'warning',
                        showCancelButton: true,
                        buttonsStyling: false,
                        confirmButtonClass: 'btn btn-danger',
                        confirmButtonText: 'Yes, delete it!',
                        cancelButtonClass: 'btn btn-secondary'
                    }).then((result) => {
                      if (result.value) {
                        window.location = $(this).attr("href")
                      }
                    })
                break;

                case 'image':
                    swal({
                        title: $(this).data("title"),
                        text: $(this).data("message"),
                        imageUrl: '/assets/img/ill/ill-1.svg',
                        buttonsStyling: false,
                        confirmButtonClass: 'btn btn-primary',
                        confirmButtonText: 'Super!'
                });
                break;

                case 'timer':
                    swal({
                        title: $(this).data("title"),
                        text: $(this).data("message"),
                        timer: 2000,
                        showConfirmButton: false
                    });
                break;
            }
        });

    });
}(jQuery))
