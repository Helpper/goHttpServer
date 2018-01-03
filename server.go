package main 
import(
	"net/http"
	"github.com/labstack/echo"
	//"github.com/labstack/echo/middleware"
	"html/template"
	"io"
	//"github.com/garyburd/redigo/redis"
	//"fmt"
)

type User struct{
	Username string `json:"username" form:"username" query:"username"`
	Password string `json:"password" form:"password" query:"password"`
}

type Template struct{
	templates *template.Template
}

//var c_redis redis.Conn

func (t *Template) Render(w io.Writer, name string, data interface{}, c echo.Context) error{
	if viewContext, isMap := data.(map[string]interface{}); isMap {
		viewContext["reverse"] = c.Echo().Reverse
	}
	return t.templates.ExecuteTemplate(w, name, data)
}

func Hello(c echo.Context) error{
	return c.Render(http.StatusOK, "hello", nil)
}

func login(c echo.Context) error{
	u := new(User)
	if err := c.Bind(u); err != nil{
		return err
	}
	// username, err := redis.String(c_redis.Do("get", "username"))
	// if err != nil{
	// 	return c.String(http.StatusOK, "login failed")
	// }else{
	// 	fmt.Printf("username: %s", username)
	// }
	// password, err := redis.String(c_redis.Do("get", "password"))
	// if err != nil{
	// 	return c.String(http.StatusOK, "login failed")
	// }else{
	// 	fmt.Printf("password: %s", password)
	// }
	// if u.Username == username && u.Password == password{
	// 	return c.String(http.StatusOK, "login successfully!")
	// }else{
	// 	return c.String(http.StatusOK, "login failed!")
	// }
	return c.String(http.StatusOK, "login successfully!")
}

func main(){
	// var err error
	// c_redis, err = redis.Dial("tcp", "localhost:6379")
	// if err != nil{
	// 	fmt.Println("Connect to redis error.", err)
	// 	return
	// }
	// defer c_redis.Close()
	e := echo.New()
	////BasicAuth
	// e.Use(middleware.BasicAuth(func(username, password string, c echo.Context) (bool, error) {
	// 	if username == "joe" && password == "secret" {
	// 		return true, nil
	// 	}
	// 	return false, nil
	// }))
	//e.GET("/", Hello)
	e.Static("/", "static")
	t := &Template{
		templates : template.Must(template.ParseGlob("views/*.html")),
	}
	e.Renderer = t
	e.GET("/hello", Hello)
	e.POST("/login", login)
	e.Logger.Fatal(e.Start(":8000"))
}