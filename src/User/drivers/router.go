package drivers

import (
	// "database/sql"
	// "fmt"
	"log"
	"net/http"
	// "os"

	// blank import for MySQL driver
	"pairing_test/src/user/interfaces/controllers"
	eth "pairing_test/src/user/drivers/ethereum"

	_ "github.com/go-sql-driver/mysql"
)

// Serve はserverを起動させます．
func Serve(addr string) {
	// データベース情報を取得
	db, err := GetDB()
	if err != nil {
		log.Fatalf("Can't get DB. %+v", err)
	}

	// パラメータを取得
	param, err := eth.GetParam()
	if err != nil {
		log.Fatalf("Can't get Param from BC. %+v", err)
	}

	// コントローラの準備
	user := controllers.LoadUserController(db, param)
	content := controllers.LoadContentController(db, param)

	
	err = http.ListenAndServe(addr, nil)
	if err != nil {
		log.Fatalf("Listen and serve failed. %+v", err)
	}
}
