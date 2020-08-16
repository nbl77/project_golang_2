package service

// func dbConn() (db *sql.DB) {
//     dbDriver := "mysql"
//     dbUser := "root"
//     dbPass := "@Sekolahpagi23175"
//     dbName := "db_inventory"
//     db, err := sql.Open(dbDriver, dbUser+":"+dbPass+"@/"+dbName)
//     if err != nil {
//         panic(err.Error())
//     }
//     return db
// }

type Model map[string]interface{}