package main

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/labstack/echo"
	"net/http"
)

type User struct{
	gorm.Model
	Name string `json:"Name" xml:"Name" form:"Name" query:"Name"`
	Identify string
	Pw string `json:"Pw" xml:"Pw" form:"Pw" query:"Pw"`
}

func main(){
	db, err := gorm.Open("mysql", "root:qordls7410@tcp(localhost:3306)/WIKI")
	if err != nil {
		panic("failed to connect to database")
	}
	defer db.Close()
	db.AutoMigrate(&User{})
	e := echo.New()
	e.Static("/static","assets")
	e.File("/join","assets/join.html")
	e.POST("/join",func(c echo.Context)error{
		u := new(User)
		if err := c.Bind(u) ;err != nil{
			return err
		}
		if len(u.Name) == 0{
			return c.HTML(http.StatusOK, "<script>window.location.href='http://localhost:1324/join';alert('이름을 입력하세요')</sript>")
		}
		if len(u.Identify) < 5{
			return c.HTML(http.StatusOK, "<script>window.location.href='http://localhost:1324/join';alert('아이디를 입력하세요')</sript>")
		}
		if len(u.Name) < 8{
			return c.HTML(http.StatusOK, "<script>window.location.href='http://localhost:1324/join';alert('패스워드를 입력하세요')</sript>")
		}

		ur := &User {
			Identify: u.Identify,
			Name : u.Name,
			Pw : u.Pw,
		}

		db.Create(&ur)

		return c.String(http.StatusOK,"회원 가입 성공")
	})
	e.Logger.Fatal(e.Start(":1324"))
}
