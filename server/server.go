package server

import (
	"fmt"
	"github.com/gin-gonic/gin"
	response "gofile/common"
	"net/http"
	"os"
	"path/filepath"
	"sort"
	. "strconv"

	"syscall"
)

type Server struct {
}

type Attribute struct {
	Path  string `json:"path" form:"path"`
	Name  string `json:"name" form:"name"`
	IsDir bool   `json:"isDir" form:"isDir"`
}

type Attributes []Attribute

func (p Attributes) Len() int {
	return len(p)
}

func (p Attributes) Less(i, j int) bool {
	for k := 0; k < len(p[i].Name) && k < len(p[j].Name); k++ {
		if p[i].Name[k] != p[j].Name[k] {
			return p[i].Name[k] < p[j].Name[k]
		}
	}
	return len(p[i].Name) < len(p[j].Name)
}

func (p Attributes) Swap(i, j int) {
	p[i], p[j] = p[j], p[i]
}

func (s *Server) ReadFile(c *gin.Context) {
	pathname := c.Query("path")
	file, err := os.Open(pathname)
	defer file.Close()
	if err != nil {
		c.Status(http.StatusNotFound)
		return
	}
	c.Header("Content-Type", "application/octet-stream")
	c.Header("Content-Disposition", "attachment; filename="+file.Name())
	c.File(pathname)
}

func (s *Server) GetAllFile(c *gin.Context) {
	pathname := c.Query("path")
	rd, err := os.ReadDir(pathname)
	if err != nil {
		fmt.Println("read dir fail:", err)
		response.Fail(c)
		return
	}
	fullName, _ := filepath.Abs(pathname)
	attributesD := make(Attributes, 0)
	attributesF := make(Attributes, 0)
	for _, fi := range rd {
		var path string
		if len(pathname) > 3 {
			path = fullName + `\` + fi.Name()
		} else {
			path = fullName + fi.Name()
		}

		attribute := Attribute{
			Path:  path,
			Name:  fi.Name(),
			IsDir: fi.IsDir(),
		}

		if fi.IsDir() {
			attributesD = append(attributesD, attribute)
		} else {
			attributesF = append(attributesF, attribute)
		}

	}

	sort.Sort(attributesD)
	sort.Sort(attributesF)
	response.OkWithData(append(attributesD, attributesF...), c)
}

func (s *Server) GetLogicalDrives(c *gin.Context) {
	kernel32 := syscall.MustLoadDLL("kernel32.dll")

	GetLogicalDrives := kernel32.MustFindProc("GetLogicalDrives")

	n, _, _ := GetLogicalDrives.Call()

	ns := FormatInt(int64(n), 2)

	var drives_all = []string{"A:", "B:", "C:", "D:", "E:", "F:", "G:", "H:", "I:", "J:", "K:", "L:", "M:", "N:", "O:", "P:", "Q:", "R:", "S:", "T:", "U:", "V:", "W:", "X:", "Y:", "Z:"}

	temp := drives_all[0:len(ns)]

	var d []string

	for i, v := range ns {
		if v == 49 {
			l := len(ns) - i - 1
			d = append(d, temp[l])
		}

	}

	attributes := make([]Attribute, 0, len(d))
	for _, v := range d {
		attributes = append(attributes, Attribute{
			Path:  v + `\`,
			Name:  v + "盘",
			IsDir: true,
		})
	}
	response.OkWithData(attributes, c)
}
