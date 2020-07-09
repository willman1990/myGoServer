package main

import (
	"database/sql"
	"fmt"
	"github.com/astaxie/beego"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"myGoWeb/goApiServer/models"
	"myGoWeb/goApiServer/providers"
	_ "myGoWeb/goApiServer/routers"
	"os"
	"time"
	"go.uber.org/fx"
	"github.com/labstack/echo"
)

func main() {
	// init to connect DB
	db, _ := connectDB()

	// init provider
	providers := initProvider(db)
	// init model
	initModels(providers)

	startBeegoWebServer()

	app := fx.New(

		// 一系列构造函数
		fx.Provide(
			// NewMyConstruct,
			// NewHandler,
			// NewMux,
			// NewLogger,
		),

		// 构造函数执行完后，执行初始化函数
		fx.Invoke(
			// invokeNothingUse, invokeRegister, invokeAnotherFunc, invokeUseMyconstruct
		),
	)

	startCtx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()
	if err := app.Start(startCtx); err != nil {
		log.Fatal(err)
	}

	stopCtx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()
	if err := app.Stop(stopCtx); err != nil {
		log.Fatal(err)
	}

}

func connectDB() (*sql.DB, error) {
	var db *sql.DB
	var err error
	connectString := "willjiang:willjiang@tcp(localhost:3306)/angular?charset=utf8"
	for {
		db, err = sql.Open("mysql", connectString)
		if err != nil {
			fmt.Println("Connect DB mysql failed , trying after 5 second. errMsg ==> " + err.Error())
			time.Sleep( 5 * time.Second)
			continue
		}
		break
	}
	return db, nil
}

func initProvider(db *sql.DB) *providers.Providers {
	return providers.NewProviders(db)
}

func initModels(providers *providers.Providers) {
	models.NewModels(providers)
}

func startBeegoWebServer() {
	beego.Run()
}

// *log.Logger 类型对象的构造函数 （注意：这里指针与非指针类型是严格区分的）
func NewLogger(lc fx.Lifecycle) *log.Logger {
	logger := log.New(os.Stdout, "" /* prefix */, 0 /* flags */)
	logger.Print("Executing NewLogger.")

	lc.Append(fx.Hook{
		OnStart: func(i context.Context) error {
			logger.Println("logger onstart..")
			return nil
		},
		OnStop: func(i context.Context) error {
			logger.Println("logger onstop..")
			return nil
		},
	})
	return logger
}