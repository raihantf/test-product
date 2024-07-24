package database

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"time"

	validation "github.com/go-ozzo/ozzo-validation"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type MySqlDb struct {
	Conn *gorm.DB
}

type DBConfig struct {
	Host     string
	Port     string
	User     string
	Password string
	Database string
}

func (DBValidate *DBConfig) Validate() error {
	if err := validation.ValidateStruct(
		validation.Field(&DBValidate.Host, validation.Required),
		validation.Field(&DBValidate.Port, validation.Required),
		validation.Field(&DBValidate.User, validation.Required),
		validation.Field(&DBValidate.Password, validation.Required),
		validation.Field(&DBValidate.Database, validation.Required),
	); err != nil {
		return err
	}
	return nil
}

func DBInit() (MySqlDb, error) {
	var setSql DBConfig

	db := MySqlDb{}
	setSql.Host = os.Getenv("DB_HOST")
	setSql.Port = os.Getenv("DB_PORT")
	setSql.User = os.Getenv("DB_USER")
	setSql.Password = os.Getenv("DB_PASSWORD")
	setSql.Database = os.Getenv("DB_DATABASE")

	dsn := fmt.Sprintf(`%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True`,
		setSql.User, setSql.Password, setSql.Host, setSql.Port, setSql.Database)

	sqlDB, err := sql.Open("mysql", dsn)
	if err != nil {
		panic("DSN is invalid")
	}
	gormDB, err := gorm.Open(mysql.New(mysql.Config{
		Conn: sqlDB,
	}), &gorm.Config{
		Logger:                 NewLogger(),
		SkipDefaultTransaction: true,
	})

	if err != nil {
		panic("Can't connect to database")
	}

	db.Conn = gormDB

	fmt.Println("Database Connecetion Success!")
	return db, nil
}

func NewLogger() logger.Interface {
	return logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
		logger.Config{
			SlowThreshold:             time.Second, // Slow SQL threshold
			LogLevel:                  logger.Info, // Log level
			IgnoreRecordNotFoundError: false,       // Ignore ErrRecordNotFound error for logger
			ParameterizedQueries:      false,       // Don't include params in the SQL log
			Colorful:                  true,        // Disable color
		},
	)
}
