package main 
import(
	"net/http"
	"github.com/labstack/echo"
	//"github.com/labstack/echo/middleware"
	"html/template"
	"io"
	"path/filepath"
	// "os"
	"strconv"
	"io/ioutil"
	"log"
	//"strings"
)

var directory string = "./file"
var url string = "/files"

type User struct{
	Username string `json:"username" form:"username" query:"username"`
	Password string `json:"password" form:"password" query:"password"`
}

type Template struct{
	templates *template.Template
}

type ConfigHtml struct{
	IndexTitle string
	ListFileTitle string
	UploadTitle string
}

type File struct{
	Name string `json:"name"`
	Size string `json:"size"`
	Type string `json:"type"`
	Path string `json:"path"`
}

var d = ConfigHtml{
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

// func listFile(c echo.Context) error{
// 	return c.Render(http.StatusOK, "listFile", d)
// }

func listAllFiles(c echo.Context) error{
	isAjax := c.Request().Header.Get("isAjax")
	if isAjax == "true"{
		url_t := c.Request().URL.EscapedPath()
		//log.Println(url_t)
		var localPath string
		listOfFiles := make([]File, 0)
		if len(url_t) > len(url){ 
			localPath = filepath.Join(directory,  url_t[len(url) : len(url_t)])
		}else{
			localPath = filepath.Join(directory, "")
		}
		dir_list, e := ioutil.ReadDir(localPath)
		if e != nil{
			log.Fatal("read dir error")
		}
		var t File
		for _, v := range dir_list{
			t.Name = v.Name()
			t.Size = strconv.FormatFloat(float64(v.Size())/1024, 'f', 4, 64)
			if v.IsDir(){
				t.Type = "dir"
			}else{
				t.Type = "file"
			}
			t.Path = localPath
			listOfFiles = append(listOfFiles, t)
		}
		//log.Println(listOfFiles)
		return c.JSON(http.StatusOK, listOfFiles)
	}else{
		return c.Render(http.StatusOK, "listFile", d)
	}
}

func main(){
	e := echo.New()
	//配置静态文件路径
	e.Static("/static", "assets")
	e.Static("/file", "file")
	t := &Template{
		templates : template.Must(template.ParseGlob("views/*.html")),
	}
	e.Renderer = t
	e.GET("/index", index)
	// e.GET("/listFile", listFile)
	e.GET("/files", listAllFiles)
	e.GET("/files/*", listAllFiles)
	e.Logger.Fatal(e.Start(":8000"))
}