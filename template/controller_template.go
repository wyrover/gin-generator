package controllers

import (
    "fmt"

    "github.com/gin-gonic/gin"
    "XXXXX/models"
)

// @Router /X@@@@s/ [get]
func Get@@@@@s(c *gin.Context) {
    if X@@@@s, err := models.FindAll@@@@@(); err != nil {
        c.AbortWithStatus(404)
        fmt.Println(err)
    } else {
        c.JSON(200, X@@@@s)
    }
    
}

// @Router /X@@@@s/{id} [get]
func Get@@@@@(c *gin.Context) {
    id := c.Params.ByName("id")
    if X@@@@, err := models.Find@@@@@ById(id); err != nil {
        c.AbortWithStatus(404)
        fmt.Println(err)
    } else {
        c.JSON(200, X@@@@)
    }

}

// @Router /X@@@@s [post]
func Post@@@@@(c *gin.Context) {
    var X@@@@ models.@@@@@
    c.BindJSON(&X@@@@)
    if X@@@@, err := models.Create@@@@@(&X@@@@); err != nil {
        c.AbortWithStatus(404)
        fmt.Println(err)
    } else {
        c.JSON(200, X@@@@)
    }
}

// @Router /X@@@@s/{id} [put]
func Put@@@@@(c *gin.Context) {
    id := c.Params.ByName("id")
    X@@@@, err := models.Find@@@@@ById(id)
    fmt.Println(X@@@@)
    if X@@@@.ID == 0 {
        c.AbortWithStatus(405)
        fmt.Println(err)
        return
    }

    c.BindJSON(&X@@@@)
    X@@@@Updated, err := models.Update@@@@@(&X@@@@)
    c.JSON(200, &X@@@@Updated)
}

// @Router /X@@@@s/{id} [delete]
func Delete@@@@@(c *gin.Context) {
    id := c.Params.ByName("id")
    if err := models.Delete@@@@@(id); err != nil {
        c.AbortWithStatus(405)
        fmt.Println(err)
        return
    }
    c.JSON(200, gin.H{"id #" + id: "deleted"})
}