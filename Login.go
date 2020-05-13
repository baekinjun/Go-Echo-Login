package main

import (
	"net/http"

	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/labstack/echo"
)

type Login struct{
	Id  string `json:"id" xml:"id" form:"id" query:"id"`
	Pwd string `json:"pwd" xml:"pwd" form:"pwd" query:"pwd"`
}

type LoginResult struct{
	Id  string `json:"id" xml:"id" form:"id" query:"id"`
	Pwd string `json:"pwd" xml:"pwd" form:"pwd" query:"pwd"`
	Result string `json:"result" xml:"result" form:"result" query:"result"`
}

func main(){
	e := echo.New()
	e.Static("/static","assets")
	e.File("/login","assets/login.html")
	e.POST("/login",func (c echo.Context)error{
		u := new(Login)
		if err := c.Bind(u); err != nil {
			return err
		}

		oriID := "schema"
		oriPwd := "1004"

		r := &LoginResult{
			Id : u.Id,
			Pwd : u.Pwd,
		}
		if (u.Id == oriID && u.Pwd == oriPwd) {
			r.Result = "성공"
		}else{
			r.Result = "너 누구냐?"
		}
		return c.JSON(http.StatusCreated , r)
	})
	e.POST("/join",func(c echo.Context)error{


	})
	e.Logger.Fatal(e.Start(":1324"))

}
