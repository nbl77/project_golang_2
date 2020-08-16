package web

import (
  "fmt"
  jwtGo "github.com/dgrijalva/jwt-go"
  "github.com/labstack/echo/v4"
  myMiddleware "github.com/labstack/echo/v4/middleware"
  "log"
  "net/http"
  "strings"
  "time"
)
type NameUser struct {
  Name string
  Alamat string
}
type JwtClaims struct {
    Name string `json:"name"`
    jwtGo.StandardClaims
}
var (
    CookieSessionLogin = "SessionLogin"
    LoginSuccessCookieVal = "LoginSuccess"
    UserIdExample = "20022012"
    SecretKeyExample = "mLmHu8f1IxFo4dWurBG3jEf1Ex0wDZvvwND6eFmcaX"
    SigningMethodExample = "HS512"
)
func main()  {
  e := echo.New()
  e.GET("/", Welcome)
  e.GET("/user/:id", GetUserID)
  e.POST("/name", GetUserName)
  g := e.Group("/admin")
  g.Use(myMiddleware.Logger())
  g.Use(myMiddleware.BasicAuth(validateUser))
  g.Use(myMiddleware.LoggerWithConfig(myMiddleware.LoggerConfig {
   Format: `[${time_rfc3339} ${status} ${method} ${host}${path} ${latency_human}]` + "\n",
  }))
  e.Use(CustomMiddleWare)
  g.GET("/main", mainAdmin)
  // INIT API Group
  jwtGroup := e.Group("/jwt")
  jwtGroup.Use(myMiddleware.JWTWithConfig(myMiddleware.JWTConfig{
    SigningMethod : SigningMethodExample,
    SigningKey :    []byte(SecretKeyExample),
    }))
  jwtGroup.GET("/main", mainJWT)
  // INIT API COOKIE
  cookieGroup := e.Group("/cookie")
  // COOKIE MIDDLEWARE
  cookieGroup.Use(checkCookie)

  // Routing `cookieGroup`
  cookieGroup.GET("/main", mainCookie)

  // Routing login
  e.GET("/login", login)


  e.Logger.Fatal(e.Start(":8001"))
}
func Welcome(e echo.Context) error {
  return e.String(http.StatusOK, "Hello world")
}
func GetUserID(c echo.Context) error {
  id:=c.Param("id")
  return c.String(http.StatusOK, "ID Yang anda masukan :" + id)
}
func GetUserName(c echo.Context) error {
  q := new(NameUser)
  c.Bind(q)
  nama := q.Name
  alamat := q.Alamat
  fmt.Println("Nama Saya :",nama)
  fmt.Println("Alamat Saya :",alamat)
  return c.JSON(http.StatusOK, q)
}
func validateUser(username, password string, c echo.Context) (bool, error) {
  if username == "nabil" && password == "amdnabil" {
    return true,nil
  }
  return false, nil
}
func mainAdmin(c echo.Context) error {
  return c.String(http.StatusOK,"Anda adalah admin")
}
func CustomMiddleWare(next echo.HandlerFunc) echo.HandlerFunc {
  return func (c echo.Context) error {
    c.Response().Header().Set(echo.HeaderServer, "ServerNBL/1.0")
    c.Response().Header().Set("Custom-header", "Hae KAKA")
    return next(c)
  }
}
func login(c echo.Context) error {
  username := c.QueryParam("username")
  password := c.QueryParam("password")

  if username == "nabil" && password == "amdnabil" {
    cookie := &http.Cookie{} //Ini sama seperti cookie := new(http.Cookie)
    cookie.Name = CookieSessionLogin
    cookie.Value = LoginSuccessCookieVal
    cookie.Expires = time.Now().Add(48 * time.Hour)
    c.SetCookie(cookie)

    // TODO: create JWT token
    token,err := createJwtToken()
    if err != nil {
      log.Println("Error saat membuat jwt token", err)
      return c.String(http.StatusInternalServerError, "ERROR : Terjadi kesalahan saat membuat JWT token")
    }
    return c.JSON(http.StatusOK, map[string] string{
      "message" : "You are logged in",
      "token" : token,
    })
  }
  return c.String(http.StatusUnauthorized,"Warning : Username atau password yang anda masukan salah")
}
func createJwtToken() (string, error) {
  claims := JwtClaims{
    "Nabil",
    jwtGo.StandardClaims{
      Id :        UserIdExample,
      ExpiresAt : time.Now().Add(24 * time.Hour).Unix(),
    },
  }
  // Hash the jwt claims
  rawToken := jwtGo.NewWithClaims(jwtGo.SigningMethodHS512, claims)
  token, err := rawToken.SignedString([]byte(SecretKeyExample))
  if err != nil {
    return "",nil
  }
  return token, nil
}
func checkCookie(next echo.HandlerFunc) echo.HandlerFunc {
  return func (c echo.Context) error {
    cookie, err := c.Cookie(CookieSessionLogin)
    if err != nil {
      if strings.Contains(err.Error(), "named cookie not present") {
        return c.String(http.StatusUnauthorized,"Anda tidak memiliki cookie")
      }
      log.Println(err)
      return err
    }
    if cookie.Value == LoginSuccessCookieVal {
      return next(c)
    }
    return c.String(http.StatusUnauthorized, "WARNING: You don't have the right cookie")
  }
}

func mainCookie(c echo.Context) error {
  return c.String(http.StatusOK, "Anda berada di halaman cookie")
}
func mainJWT(c echo.Context) error {
  user := c.Get("user")
  token := user.(*jwtGo.Token)
  claims := token.Claims.(jwtGo.MapClaims)
  log.Println("Username : ", claims["name"],", User ID : ",claims["jti"])

  return c.JSON(http.StatusOK, map[string] interface{}{
      "claims" : claims,
      "message" : "SUCCESS : Anda berada di halaman jwt",
    })
}
