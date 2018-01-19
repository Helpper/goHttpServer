package main 
import(
	"net/http"
	"github.com/labstack/echo"
	//"github.com/labstack/echo/middleware"
	"html/template"
	"io"
	//"fmt"
)

type User struct{
	Username string `json:"username" form:"username" query:"username"`
	Password string `json:"password" form:"password" query:"password"`
}

type Template struct{
	templates *template.Template
}

type Title struct{
	IndexTitle string
	ListFileTiTle string
	UploadTitle string
}

var d = Title{
	IndexTitle: "fileServer",
	ListFileTiTle: "listFile",
	UploadTitle: "uploadfile",
}

func (t *Template) Render(w io.Writer, name string, data interface{}, c echo.Context) error{
	if viewContext, isMap := data.(map[string]interface{}); isMap {
		viewContext["reverse"] = c.Echo().Reverse
	}
	return t.templates.ExecuteTemplate(w, name, data)
}

func index(c echo.Context) error{
	return c.Render(http.StatusOK, "index", d)
}

func listFile(c echo.Context) error{
	return c.Render(http.StatusOK, "listFile", d)
}

func main(){
	e := echo.New()
	e.Static("/", "static")
	t := &Template{
		templates : template.Must(template.ParseGlob("views/*.html")),
	}
	e.Renderer = t
	e.GET("/index", index)
	e.GET("/files", listFile)
	e.Logger.Fatal(e.Start(":8000"))
}