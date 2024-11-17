package handlers

import (
	"net/http"

	"github.com/ardiost/golang-clean-web-api/api/helper"
	"github.com/gin-gonic/gin"
)

type personInformation struct {
	FirstName    string `json:"first_name" binding:"required,alpha,min=4,max=10"`
	LastName     string `json:"last_name" binding:"required,alpha,min=5,max=20"`
	MobileNumber string `json:"mobile_number" binding:"required,mobile"`
}

type Header struct {
	UserId   string
	Platform string
}
type TestHandler struct{}

func NewTestHandler() *TestHandler {
	return &TestHandler{}
}

// Test godoc
// @Summary UserById
// @Description UserById
// @Tags Test
// @Accept json
// @Produce json
// @Success 200 {object} helper.BaseHttpResponse "Success"
// @Failure 400 {object} helper.BaseHttpResponse "Failed"
// @Router /v1/tester/test1/ [get]
func (h *TestHandler) TestHa(c *gin.Context) {
	c.JSON(http.StatusOK, helper.GenerateBaseResponse(gin.H{
		"result": "this is for this",
	}, true, 0))
}

func (h *TestHandler) GetUsers(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"result": "[{alireza 12 isf},{reza 13 shahreza}]",
	})
}

// UserbyId godoc
// @Summary UserById
// @Description UserById
// @Tags Test2
// @Accept json
// @Produce json
// @param id path int true "user id"
// @Success 200 {object} helper.BaseHttpResponse "Success"
// @Failure 400 {object} helper.BaseHttpResponse "Failed"
// @Router /v1/tester/test1/user/{id} [get]
func (h *TestHandler) GetUserById(c *gin.Context) {
	id := c.Param("id")
	c.JSON(http.StatusOK, helper.GenerateBaseResponse(gin.H{
		"result": "UserById",
		"id":     id,
	}, true, 0))
}

func (h *TestHandler) GetUserByUserName(c *gin.Context) {
	userName := c.Param("username")
	c.JSON(http.StatusOK, gin.H{
		"result":  "ok",
		"user-id": userName,
	})
}

func (h *TestHandler) GetAccounts(c *gin.Context) {
	id := c.Param("id")
	c.JSON(http.StatusOK, gin.H{
		"result":  "Accounts",
		"user-id": id,
	})
}

func (h *TestHandler) AddUser(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"result": "user added",
	})
}

func (h *TestHandler) Binder1(c *gin.Context) {
	id := c.GetHeader("userId")
	c.JSON(http.StatusOK, gin.H{
		"result": "userId",
		"id":     id,
	})
}

func (h *TestHandler) Binder2(c *gin.Context) {
	header := Header{}
	c.BindHeader(&header)
	c.JSON(http.StatusOK, gin.H{
		"result": header,
	})
}

func (h *TestHandler) Query1(c *gin.Context) {
	id := c.Query("id")
	name := c.Query("name")
	c.JSON(http.StatusOK, gin.H{
		"result": "Query1",
		"id":     id,
		"name":   name,
	})
}
func (h *TestHandler) Query2(c *gin.Context) {
	ids := c.QueryArray("id")
	name := c.Query("name")
	c.JSON(http.StatusOK, gin.H{
		"result": "Query1",
		"id":     ids,
		"name":   name,
	})
}

// BodyBind godoc
// @Summary BodyBind
// @Description BodyBind
// @Tags Test3
// @Accept json
// @Produce json
// @param person body personInformation true "person data"
// @Success 200 {object} helper.BaseHttpResponse "Success"
// @Failure 400 {object} helper.BaseHttpResponse "Failed"
// @Router /v1/tester/test1/body [post]
// @Security AuthBearer
func (h *TestHandler) BodyBind(c *gin.Context) {
	p := personInformation{}
	err := c.ShouldBindJSON(&p)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"validation error": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"result": "Query1",
		"data":   p,
	})
}
