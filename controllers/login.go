package controllers

import (
	"goapi/models"
	"os"

	_ "goapi/docs"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

var erro models.Err

var DB *Database = InitDB(os.Getenv("DBUSER"), os.Getenv("DBPASS"), os.Getenv("DBHOST"), os.Getenv("DBPORT"), os.Getenv("DBNAME"), os.Getenv("DBTABLE"))

// GetUSER godoc
// @Summary Get all users
// @Description Fetches all available users from DB
// @Tags Users
// @Produce json
// @Success 200 {object} []models.User
// @Failure 400 {object} models.Err
// @Failure 404 {object} models.Err
// @Failure 500 {object} models.Err
// @Router /users [get]
func GetUser(c *gin.Context) {
	var user []models.User

	// TODO: add checking/creating of DB handler
	_ = DB.CheckDB()
	err := DB.CheckTable("test")
	if err != nil {
		erro.Severity, erro.Body = "error", "Something wrong with DB"
		c.JSON(500, erro)
		return
	}
	err = DB.SelectAll(&user)
	if err == nil {
		if len(user) == 0 {
			erro.Severity, erro.Body = "error", "no users registered yet"
			c.JSON(200, erro)
			return
		}
		c.JSON(200, user)
	} else {
		erro.Severity, erro.Body = "error", "user not found"
		c.JSON(404, erro)
	}
}

// PostUser godoc
// @Summary Create new user
// @Description Create/Add new user to DB via POST request
// @Tags Users
// @Accept json
// @Produce json
// @Param users body models.User true "Specification for user which should be added"
// @Success 200 {object} models.Err
// @Failure 400 {object} models.Err
// @Failure 404 {object} models.Err
// @Router /users [post]
func PostUser(c *gin.Context) {
	var user models.User
	c.Bind(&user)
	if user.Username != "" && user.Password != "" && user.Firstname != "" && user.Lastname != "" {
		err := DB.InsertUSER(user)
		if err != nil {
			erro.Severity, erro.Body = "error", err
			c.JSON(400, erro)
		} else {
			erro.Severity, erro.Body = "info", "User has been added."
			c.JSON(200, erro)
		}
	} else {
		erro.Severity, erro.Body = "error", "Fields are empty"
		c.JSON(400, erro)
	}
}

func GetUserDetail(c *gin.Context) {
	id := c.Params.ByName("id")
	var user models.User

	err := DB.GetUser(&user, id)

	if err == nil {
		c.JSON(200, user)
	} else {
		erro.Severity, erro.Body = "error", "user not found"
		c.JSON(404, erro)
	}
}

// func Login(c *gin.Context) {
// 	var user models.User
// 	c.Bind(&user)
// 	err := dbmap.SelectOne(&user, "select * from user where Username=? LIMIT 1", user.Username)
// 	if err == nil {
// 		user_id := user.Id

// 		content := &models.User{
// 			Id:        user_id,
// 			Username:  user.Username,
// 			Password:  user.Password,
// 			Firstname: user.Firstname,
// 			Lastname:  user.Lastname,
// 		}
// 		c.JSON(200, content)
// 	} else {
// 		c.JSON(404, gin.H{"error": "user not found"})
// 	}
// }

// func UpdateUser(c *gin.Context) {
// 	id := c.Params.ByName("id")
// 	var user models.User
// 	err := dbmap.SelectOne(&user, "SELECT * FROM user WHERE id=?", id)
// 	if err == nil {
// 		var json models.User
// 		c.Bind(&json)
// 		user_id, _ := strconv.ParseInt(id, 0, 64)
// 		user := models.User{
// 			Id:        user_id,
// 			Username:  user.Username,
// 			Password:  user.Password,
// 			Firstname: json.Firstname,
// 			Lastname:  json.Lastname,
// 		}

// 		if user.Firstname != "" && user.Lastname != "" {
// 			_, err = dbmap.Update(&user)
// 			if err == nil {
// 				c.JSON(200, user)
// 			} else {
// 				checkErr(err, "Updated failed")
// 			}
// 		} else {
// 			c.JSON(400, gin.H{"error": "fields are empty"})
// 		}
// 	} else {
// 		c.JSON(404, gin.H{"error": "user not found"})
// 	}
// }
