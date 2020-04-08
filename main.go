package main

import (
	"flag"
	"fmt"
	"github.com/go-xorm/xorm"
	"github.com/hyecheonlee/echosample/controller"
	"github.com/hyecheonlee/echosample/models"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	_ "github.com/mattn/go-sqlite3"
	"github.com/pangpanglabs/echoswagger"
	configutil "github.com/pangpanglabs/goutils/config"
	"github.com/pangpanglabs/goutils/echomiddleware"
	"github.com/pangpanglabs/goutils/echotpl"
	"log"
	"os"
	"runtime"
)

func main() {
	appEnv := flag.String("app-env", os.Getenv("APP_ENV"), "app env")
	flag.Parse()
	var c Config
	if err := configutil.Read(*appEnv, &c); err != nil {
		panic(err)
	}
	fmt.Println(c)
	db, err := initDB(c.Database.Driver, c.Database.Connection)
	if err != nil {
		panic(err)
	}
	defer db.Close()
	e := echo.New()
	echoswagger.New(e, "/doc", &echoswagger.Info{
		Title:       "Echo sample",
		Description: "This is API doc for Echo Sample",
		Version:     "1.0",
	})

	controller.HomeController{}.Init(e.Group("/"))

	//정적 파일 경로
	e.Static("/static", "static")
	e.Pre(middleware.RemoveTrailingSlash())

	e.Pre(echomiddleware.ContextBase())

	//샘플코드 만든 착한분의 유틸 나중에 보자.
	e.Renderer = echotpl.New()

	if err := e.Start(":" + c.HttpPort); err != nil {
		log.Println(err)
	}

}

func initDB(driver, connection string) (*xorm.Engine, error) {
	db, err := xorm.NewEngine(driver, connection)
	if err != nil {
		return nil, err
	}

	if driver == "sqlite3" {
		runtime.GOMAXPROCS(1)
	}

	db.Sync(new(models.Discount))
	return db, nil
}

type Config struct {
	Database struct {
		Driver     string
		Connection string
		Logger     struct {
			Kafka echomiddleware.KafkaConfig
		}
	}
	BehaviorLog struct {
		Kafka echomiddleware.KafkaConfig
	}
	Trace struct {
		Zipkin echomiddleware.ZipkinConfig
	}

	Debug    bool
	Service  string
	HttpPort string
}
