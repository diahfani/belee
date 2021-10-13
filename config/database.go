package config

import (
	// "github.com/jinzhu/gorm"
	"belee/models/buyer"
	"belee/models/owner"
	"belee/models/paymentMethod"
	"belee/models/products"
	"belee/models/productsType"
	"belee/models/transactions"
	"belee/models/warung"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

// **Contoh mas Rizky**

func DbConfig() {
	dsn := "admin0110:admin0110@tcp(db-belee.cxinhvzfptwa.us-east-2.rds.amazonaws.com:3306)/belee_test?charset=utf8mb4&parseTime=True&loc=Local"
	var err error
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("DB failed to connect")
	}
	Migration()
}

func Migration() {
	DB.AutoMigrate(&buyer.Buyers{}, &owner.Owners{}, &warung.Warungs{}, &paymentMethod.PaymentMethods{}, &products.Products{}, &productsType.ProductsType{})

	DB.AutoMigrate(&transactions.Transactions{})
	// DB.Migrator().RenameColumn(&transactions.Transactions{}, "products_name", "productsname")
	// DB.Migrator().RenameColumn(&transactions.Transactions{}, "total_qty", "totalqty")
	// DB.Migrator().RenameColumn(&transactions.Transactions{}, "total_price", "totalprice")
	// DB.Migrator().DropColumn(&warung.Warungs{}, "owners_id")
	// DB.Migrator().DropColumn(&products.Products{}, "warung_id")
	// DB.Migrator().DropColumn(&products.Products{}, "barang_type_id")
	// DB.Migrator().DropColumn(&transactions.Transactions{}, "barang_id")
	// DB.AutoMigrate()
}

func InitDBTest() {
	dsn := "admin0110:admin0110@tcp(db-belee.cxinhvzfptwa.us-east-2.rds.amazonaws.com:3306)/belee_testing?charset=utf8mb4&parseTime=True&loc=Local"
	var err error
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("DB failed to connect")
	}
	MigrationTest()

}

func MigrationTest() {
	DB.Migrator().DropTable(&buyer.Buyers{})
	DB.AutoMigrate(&buyer.Buyers{})
	DB.Migrator().DropTable(&owner.Owners{})
	DB.AutoMigrate(&owner.Owners{})
	DB.Migrator().DropTable(&warung.Warungs{})
	DB.AutoMigrate(&warung.Warungs{})
	DB.Migrator().DropTable(&paymentMethod.PaymentMethods{})
	DB.AutoMigrate(&paymentMethod.PaymentMethods{})
	DB.Migrator().DropTable(&products.Products{})
	DB.AutoMigrate(&products.Products{})
	DB.Migrator().DropTable(&productsType.ProductsType{})
	DB.AutoMigrate(&productsType.ProductsType{})
	DB.Migrator().DropTable(&transactions.Transactions{})
	DB.AutoMigrate(&transactions.Transactions{})
}
