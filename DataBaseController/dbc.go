package DataBaseController

import (
	"database/sql"
	"fmt"
	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
	"net/http"
	"sina_project/Models"
	"sina_project/Token"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "*********"
	dbname   = "*********"
)

var (
	psqlInfo = fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)
	db, err = sql.Open("postgres", psqlInfo)
)

func Start()  {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)
	db, err := sql.Open("postgres", psqlInfo)
	defer db.Close()

	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("knihiughniefganiurvgiufdv")
	err = db.Ping()
	if err != nil {
		fmt.Println(err)
	}

}
func Login(c *gin.Context,user Models.Users){
	sqlStatement := `SELECT username , password FROM users WHERE username=$1;`
	var username string
	var password string
	row := db.QueryRow(sqlStatement, user.Username)
	switch err := row.Scan(&username,&password); err {
	case sql.ErrNoRows:
		c.JSON(200, gin.H{"username": "there is no such username" })
	case nil:
		if password==user.Password{
			c.JSON(http.StatusOK, gin.H{"message": "logged in successfully","token":Token.GetToken(user.Username),"status": http.StatusOK})
		}else{
			c.JSON(http.StatusOK, gin.H{"message": "wrong password", "status": http.StatusOK})
		}
	default:
		fmt.Println(err)
	}
}
func Register(c *gin.Context,user Models.Users)  {
	sqlStatement := `SELECT username FROM users WHERE username=$1;`
	var username string
	row := db.QueryRow(sqlStatement, user.Username)
	switch err := row.Scan(&username); err {
	case sql.ErrNoRows:
		sqlStatement = "INSERT INTO users (username, password) VALUES ($1, $2)"
		stmt,err:=db.Prepare(sqlStatement)
		if err != nil {fmt.Println("3")
			panic(err)
		}
		stmt.QueryRow(user.Username, user.Password)
		c.JSON(200, gin.H{"username": user.Username , "password":user.Password})
	case nil:
		c.JSON(http.StatusOK, gin.H{"message": "this username has been taken before", "status": http.StatusOK})
	default:
		panic(err)
	}
}
func IsUnique(name string,c *gin.Context) bool {
	sqlStatement := `SELECT file_id FROM permissions WHERE file_id=$1;`
	var username string
	row := db.QueryRow(sqlStatement, name)
	switch err := row.Scan(&username); err {
	case sql.ErrNoRows:
    return true
	case nil:
		c.JSON(http.StatusOK, gin.H{"message": "this file_id has been taken before", "status": http.StatusOK})
		return false
	default:
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"message": err,
		})
		return false
	}
}
func AddPermisssion(user string,file string)  {
	sqlStatement := "INSERT INTO permissions (file_id, username) VALUES ($1, $2)"
	stmt,err:=db.Prepare(sqlStatement)
	if err != nil {fmt.Println("3")
		panic(err)
	}
	stmt.QueryRow(file, user)
}
func IsPermited(username string,file_id string,c *gin.Context) bool{
	sqlStatement := `SELECT file_id FROM permissions WHERE file_id=$1 and username=$2;`
	var x string
	row := db.QueryRow(sqlStatement, file_id,username)
	err := row.Scan(&x)
	fmt.Println(x)
	switch  err {
	case sql.ErrNoRows:
		c.JSON(http.StatusOK, gin.H{"message": "you are not permitted to download this file", "status": http.StatusOK})
		return false

	case nil:

		return true
	default:
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"message": err,
		})
		return false
	}
	return true
}