package drivers
import (
	"github.com/go-sql-driver/mysql"
)
// TODO: DBの設定をgormにする?
GetDB() (mysql.Conn, error) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_HOST"), os.Getenv("DB_PORT"), os.Getenv("DATABASE"))
	conn, err := sql.Open("mysql", dsn)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	return  conn, nil 
}