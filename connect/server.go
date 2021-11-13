package connect

import (
	"bitbucket/models"
	"fmt"
	"github.com/spf13/viper"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	cfg *viper.Viper
	db  *gorm.DB
	err error
)

type Connect interface {
	SqlDb() *gorm.DB
	Config() *viper.Viper
	ApiServer(addr []string) string
}

type connect struct {}

func (c *connect) SqlDb() *gorm.DB {
	dbUser := c.Config().GetString("DB_USER")
	dbPassword := c.Config().GetString("DB_PASSWORD")
	dbHost := c.Config().GetString("DB_HOST")
	dbPort := c.Config().GetString("DB_PORT")
	dbName := c.Config().GetString("DB_NAME")

	db, err = gorm.Open(postgres.Open(fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=require", dbHost, dbUser, dbPassword, dbName, dbPort)))
	db.AutoMigrate(&models.Inventory{},&models.Merchant{},&models.Outlet{}, &models.Pelanggan{}, &models.Product{}, &models.Transaksi{}, &models.TransaksiDetail{}, &models.User{})

	if err != nil {
		panic(err)
	}

	return db
}

func (c *connect) Config() *viper.Viper {
	viper.AddConfigPath(".")
	viper.SetConfigFile(".env")
	viper.AutomaticEnv()
	viper.ReadInConfig()
	cfg = viper.GetViper()
	return cfg}

func (c *connect) ApiServer(addr []string) string {
	switch len(addr) {
	case 0:
		if port := c.Config().GetString("PORT"); port != "" {
			debugPrint("Environment variable PORT=" + port)
			return ":" + port
		}
		debugPrint("Environment variable PORT is undefined. Using port :8081 by default")
		return ":8081"
	case 1:
		return addr[0]
	default:
		panic("too many parameters")
	}}

func debugPrint(format string, values ...interface{}) {
	fmt.Println(format)
}

func NewConnect() Connect  {
	return &connect{}
}


