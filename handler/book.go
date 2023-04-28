package handler

import (
	"challange-10/helper"
	"challange-10/model"
	"strconv"

	"github.com/gin-gonic/gin"
)

func (h HttpServer) AddBook(c *gin.Context) {
	in := model.Book{}

	err := c.BindJSON(&in)
	if err != nil {
		helper.BadRequest(c, err.Error())
		return
	}

	res, err := h.app.AddBook(in)
	if err != nil {
		helper.InternalServerError(c, err.Error())
		return
	}

	helper.Ok(c, res)
}

func (h HttpServer) GetAllBook(c *gin.Context) {

	res, err := h.app.GetAllBook()
	if err != nil {
		helper.InternalServerError(c, err.Error())
		return
	}

	helper.Ok(c, res)
}

func (h HttpServer) GetBookId(c *gin.Context) {
	id := c.Param("id")

	idInt, err := strconv.Atoi(id)
	if err != nil {
		helper.BadRequest(c, err.Error())
		return
	}

	res, err := h.app.GetBookId(int64(idInt))
	if err != nil {
		helper.InternalServerError(c, "Buku dengan ID tersebut tidak ditemukan")
		return
	}

	helper.Ok(c, res)
}

func (h HttpServer) UpdateBook(c *gin.Context) {
	id := c.Param("id")

	idInt, err := strconv.Atoi(id)
	if err != nil {
		helper.BadRequest(c, err.Error())
		return
	}

	in := model.Book{}

	err = c.BindJSON(&in)

	in.ID = idInt

	_, err = h.app.UpdateBook(in)
	if err != nil {
		helper.InternalServerError(c, "Buku dengan ID tersebut tidak ditemukan")
		return
	}

	helper.Ok(c, "Berhasil update Data Buku")
}

func (h HttpServer) DeleteBook(c *gin.Context) {
	id := c.Param("id")

	idInt, err := strconv.Atoi(id)
	if err != nil {
		helper.BadRequest(c, err.Error())
		return
	}

	err = h.app.DeleteBook(int64(idInt))
	if err != nil {
		helper.InternalServerError(c, err.Error())
		return
	}

	helper.Ok(c, "Data Buku berhasil dihapus")
}
