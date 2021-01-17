package controllers

import (
	"goapi/models"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

var DB *Database = InitDB(os.Getenv("DBUSER"), os.Getenv("DBPASS"), os.Getenv("DBHOST"), os.Getenv("DBPORT"), os.Getenv("DBNAME"), os.Getenv("DBTABLE"))

func GetUser(c *gin.Context) {
	var user []models.User

	// TODO: add checking/creating of DB handler
	_ = DB.CheckDB()
	err := DB.CheckTable("test")
	if err != nil {
		c.JSON(500, gin.H{"error": "Something wrong with DB"})
		return
	}
	err = DB.SelectAll(&user)
	if err == nil {
		if len(user) == 0 {
			c.JSON(200, gin.H{"error": "no users registered yet"})
			return
		}
		c.JSON(200, user)
	} else {
		c.JSON(404, gin.H{"error": "user not found"})
	}
}

func PostUser(c *gin.Context) {
	var user models.User
	c.Bind(&user)
	log.Println(user)
	if user.Username != "" && user.Password != "" && user.Firstname != "" && user.Lastname != "" {
		err := DB.InsertUSER(user)
		if err != nil {
			c.JSON(400, gin.H{"error": err})
		}
	} else {
		c.JSON(400, gin.H{"error": "Fields are empty"})
	}
}

func GetUserDetail(c *gin.Context) {
	id := c.Params.ByName("id")
	var user models.User

	err := DB.GetUser(&user, id)

	if err == nil {
		// user.Id, _ = strconv.ParseInt(id, 0, 64)
		c.JSON(200, user)
	} else {
		c.JSON(404, gin.H{"error": "user not found"})
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
