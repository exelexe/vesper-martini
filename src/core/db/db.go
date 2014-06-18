package db

import (
	"core"
	. "core/config"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"log"
)

var DB gorm.DB

func init() {
	var err error

	log.Println("DB DataSource: " + Conf.Dsn)

	DB, err = gorm.Open("mysql", Conf.Dsn)
	core.CheckFatal(err, "Database connect error")
	DB.LogMode(true)

	// DB, err = gorm.Open("postgres", "user=gorm dbname=gorm sslmode=disable")
	// db, err = gorm.Open("sqlite3", "/tmp/gorm.db")

	// Connection string parameters for Postgres - http://godoc.org/github.com/lib/pq, if you are using another
	// database refer to the relevant driver's documentation.

	// * dbname - The name of the database to connect to
	// * user - The user to sign in as
	// * password - The user's password
	// * host - The host to connect to. Values that start with / are for unix domain sockets.
	//   (default is localhost)
	// * port - The port to bind to. (default is 5432)
	// * sslmode - Whether or not to use SSL (default is require, this is not the default for libpq)
	//   Valid SSL modes:
	//    * disable - No SSL
	//    * require - Always SSL (skip verification)
	//    * verify-full - Always SSL (require verification)
}
