package DbConnection

import(
	"fmt"
	"log"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"github.com/ravinthiranpartheepan1407/go-ecommerce"
)

funct GetDB() *gorm.DB{
	dbConfig := connection.DbConfig()
	status:= fmt.Sprintf("host=%s, port=%s, password=%s, user=%s, dbname=%s, sslmode=%s", dbConfig.host, dbConfig.port, dbConfig.password, dbConfig.user, dbConfig.dbname, dbConfig.sslmode)
	db, err := gorm.open(postgres.Open(status), &gorm.Config{})
	if err != nil{
		fmt.Println("Falied to connect with postgres")
	}

	return db
}