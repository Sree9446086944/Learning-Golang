package controllers

import (
	"github.com/Sree9446086944/jsonCrudApi/initializers"
	"github.com/Sree9446086944/jsonCrudApi/models"
	"github.com/gin-gonic/gin"
)

func Test(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "pong",
	})
}

/*
//hard coded req body
func PostsCreate(c *gin.Context) {
	// https://gorm.io/docs/create.html - refer

	//get data off req body since post req
	//not getting req body now, hardcoding it
	//create a post
	post := models.Post{Title: "Hello", Body: "Post Body"}

	result := initializers.DB.Create(&post) //created post struct
	//result has - ID(primary key), Error, RowsAffected
	if result.Error != nil {
		c.Status(400)
		return
	}

	//return it
	c.JSON(200, gin.H{
		"post": post,
	})
}
*/

// with actual req body
func PostsCreate(c *gin.Context) {
	// https://gorm.io/docs/create.html - refer

	//get data off req body since post req
	var body struct {
		Body  string
		Title string
	}

	c.Bind(&body) //pass reference of the struct of req body, Bind() parses the req body json/xml, decodes the json payload into the struct specified as a pointer

	//create a post
	post := models.Post{Title: body.Title, Body: body.Body}

	result := initializers.DB.Create(&post) //created post struct
	//result has - ID(primary key), Error, RowsAffected
	if result.Error != nil {
		c.Status(400)
		return
	}

	//return it
	c.JSON(200, gin.H{
		"post": post,
	})
}

// http://localhost:3000/posts - POST
// {
//     "Title":"from req body",
//     "Body":"from req body"
// }

// fetch all
func PostsIndex(c *gin.Context) {
	//check gorm docs
	//Get the posts
	var posts []models.Post
	initializers.DB.Find(&posts)

	//respond with them
	c.JSON(200, gin.H{
		"posts": posts,
	})
}

// http://localhost:3000/posts - GET

// get post by id
func PostsShow(c *gin.Context) {
	//Get id from url param(dynamic route :id)
	id := c.Param("id") //get id from param and assign to id var
	//Get the post
	var post models.Post
	initializers.DB.First(&post, id) //get post by id and write to post struct

	//respond with them
	c.JSON(200, gin.H{
		"posts": post,
	})

}

// http://localhost:3000/posts/1

// post update
func PostsUpdate(c *gin.Context) {
	//get id off url
	id := c.Param("id")

	//get data off req body
	var body struct {
		Body  string
		Title string
	}

	c.Bind(&body) //binging req json to body struct using reference(original var change)

	//find post were updating  - https://gorm.io/docs/update.html  -> update multiple columns
	var post models.Post
	initializers.DB.First(&post, id) // find post by id from db and write to post struct

	//update it
	initializers.DB.Model(&post).Updates(models.Post{
		Title: body.Title,
		Body:  body.Body,
	})

	// respond with it
	c.JSON(200, gin.H{
		"post": post,
	})

}

/*http://localhost:3000/posts/1 - PUT
{
    "Title":"updated req body",     // need to send only  values need to be updated
    "Body":"updated req body"
}*/

func PostsDelete(c *gin.Context) {
	//get id from url
	id := c.Param("id")

	//Delete the posts
	initializers.DB.Delete(&models.Post{}, id) //Delete(type reference, id(condition))
	//Delete() will soft delete, i.e, the data will be in db but the deleted at column will be updated to current time if null,
	//can revert if needed, if fetch all values, deleted record wont show up

	//respond
	c.Status(200)
}

// http://localhost:3000/posts/1 - DELETE
