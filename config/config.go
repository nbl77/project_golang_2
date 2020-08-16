package config
const(
  PORT = ":9000"
  DB_USER = "root"
  DB_PASS = "123"
  DB_HOST = "localhost"
  DB_NAME = "db_inventory"
  COOKIE_LOGIN_KEY = "SessionSignin"
  COOKIE_LOGIN_VAL = "LoginSuccess"
  COOKIE_ALERT = "alert"
  COOKIE_LOGIN_EXPIRES = 72 //jam
)

var (
  ALERT_STATUS = "basic"
  ALERT_MESSAGE = "Berhasil Inisialisasi"
  ALERT_TITLE = "Berhasil!"
  ALERT_VAL = `
                    swal({
                        title: '`+ALERT_TITLE+ `',
                        text: '`+ALERT_MESSAGE+`',
                        type: '`+ALERT_STATUS+`',
                        buttonsStyling: false,
                        confirmButtonClass: 'btn btn-success'
                    })`
)

type M map[string]interface{}

