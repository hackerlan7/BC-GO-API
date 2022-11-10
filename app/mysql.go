package app

// import (
// 	"database/sql"
// 	"log"

// 	_ "github.com/go-sql-driver/mysql"
// )

// //定義全域 以後不用MySql直接換driver即可
// var db *sql.DB

// // 與DB連線。 init() 初始化，時間點比 main() 更早。
// func InitDb() {
// 	dbConnect, err := sql.Open(
// 		"mysql",
// 		//user+pass+@tcp(host)
// 		"root:linkroot@tcp(127.0.0.1:3306)/",
// 	)

// 	if err != nil {
// 		log.Fatalln(err)
// 	}

// 	err = dbConnect.Ping()
// 	if err != nil {
// 		log.Fatalln(err)
// 	}

// }
