package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	postgres "github.com/jackc/pgx"
	"mime/multipart"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
)

type settings struct {
	host     string
	port     string
	user     string
	password string
	dbname   string
}

var s *settings = settingsLoad()

var db *postgres.Conn = createConnectionDB(s)

func main() {

	r := gin.Default()
	r.POST("/", loadFile)
	r.Run(":5000")
}

func settingsLoad() *settings {
	return &settings{
		host:     os.Getenv("DB-HOST"),
		port:     os.Getenv("DB_PORT"),
		user:     os.Getenv("DB_USER"),
		password: os.Getenv("DB_PASSWORD"),
		dbname:   os.Getenv("DB_NAME"),
	}
}

func loadFile(c *gin.Context) {

	form, err := c.MultipartForm()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	files := form.File["file"]
	if len(files) == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "No files uploaded"})
		return
	}

	for _, fileHeader := range files {
		file, err := fileHeader.Open()
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Unable to open file"})
			return
		}
		defer file.Close()

		filename, extension := extractFileMetadata(fileHeader)
		fmt.Sprintf("%s, %s", filename, extension)
		//createMetadata(filename, extension)

	}
	c.JSON(http.StatusOK, gin.H{"message": "Files uploaded successfully"})

}

func extractFileMetadata(fileHeader *multipart.FileHeader) (string, string) {
	filename := fileHeader.Filename
	extension := filepath.Ext(filename)
	return filename, extension
}

func createConnectionDB(s *settings) *postgres.Conn {

	portInt, _ := strconv.Atoi(s.port)

	confDB := postgres.ConnConfig{
		Host:     s.host,
		Port:     uint16(portInt),
		Database: s.dbname,
		User:     s.user,
		Password: s.password,
	}

	cfg, err := postgres.Connect(confDB)
	if err != nil {
		panic("Unable to connect to database")
	}

	return cfg
}

//func createMetadata(filename string, extension string) {
//	sqlStatement := `
//        INSERT INTO documents_metadata (filename, extension)
//        VALUES ($1, $2)`
//
//	_, err := db.Exec(sqlStatement, filename, extension)
//	if err != nil {
//		panic("Error inserting entry")
//	}
//}
