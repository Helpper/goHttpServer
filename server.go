package main 
import(
	"net/http"
	"github.com/labstack/echo"
	//"github.com/labstack/echo/middleware"
	"html/template"
	"io"
	"path/filepath"
	"os"
	"strconv"
	//"fmt"
)

var directory string = "./file"

type User struct{
	Username string `json:"username" form:"username" query:"username"`
	Password string `json:"password" form:"password" query:"password"`
}

type Template struct{
	templates *template.Template
}

type Title struct{
	IndexTitle string
	ListFileTitle string
	UploadTitle string
}

type File struct{
	Name string `json:"name"`
	Size string `json:"size"`
}

var d = Title{
	IndexTitle: "fileServer",
	ListFileTitle: "listFile",
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

func listAllFiles(c echo.Context) error{
	listOfFiles := make([]File, 0)
	filepath.Walk(directory, func(path string, f os.FileInfo, err error) error{
		if f == nil{
			return err
		}
		if f.IsDir(){
			return nil
		}
		var t File
		t.Name = f.Name()
		t.Size = strconv.FormatInt(f.Size()/1024, 10)
		listOfFiles = append(listOfFiles, t)
		return nil
	})
	return c.JSON(http.StatusOK, listOfFiles)
}

func main(){
	e := echo.New()
	e.Static("/static", "assets")
	t := &Template{
		templates : template.Must(template.ParseGlob("views/*.html")),
	}
	e.Renderer = t
	e.GET("/index", index)
	e.GET("/listFile", listFile)
	e.GET("/files", listAllFiles)
	e.Logger.Fatal(e.Start(":8000"))
}