package controller

import (
	"net/http"
	"ticketsale/model"
	"time"

	"github.com/gin-gonic/gin"
)

// ticket CURD
func CreateTicket(ctx *gin.Context) {
	type form struct {
		Name      string `json:"name"`
		Introduce string `json:"introduce"`
		Count     int    `json:"count"`
		ExpiresAt string `json:"expires_at"`
	}
	var req form
	var model model.Tickets
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"msg":   "Invalid form",
			"error": err.Error(),
		})
		return
	}
	model.Name = req.Name
	model.Introduce = req.Introduce
	model.Count = req.Count
	timeLayout := "2006-01-02 15:04:05"
	expireAt, err := time.Parse(timeLayout, req.ExpiresAt)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"msg":   "Invalid time format",
			"error": err.Error(),
		})
		return
	}
	model.ExpiresAt = expireAt
	err = model.Create()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"msg":   "Create ticket failed",
			"error": err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"msg": "Create ticket success",
	})
}

func UpdateTicket(ctx *gin.Context) {
	id := ctx.Param("id")
	type form struct {
		Count     int    `json:"count"`
		Introduce string `json:"introduce"`
		ExpiresAt string `json:"expires_at"`
	}
	var req form
	var model model.Tickets
	_, err := model.Find(map[string]interface{}{"id": id})
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"msg":   "Ticket not found",
			"error": err.Error(),
		})
		return
	}

	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"msg":   "Invalid form",
			"error": err.Error(),
		})
		return
	}
	model.Count = req.Count
	model.Introduce = req.Introduce
	timeLayout := "2006-01-02 15:04:05"
	expireAt, err := time.Parse(timeLayout, req.ExpiresAt)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"msg":   "Invalid time format",
			"error": err.Error(),
		})
		return
	}
	model.ExpiresAt = expireAt
	err = model.Update()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"msg":   "Update ticket failed",
			"error": err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"msg": "Update ticket success",
	})

}

func DeleteTicket(ctx *gin.Context) {
	id := ctx.Param("id")
	var model model.Tickets
	_, err := model.Find(map[string]interface{}{"id": id})
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"msg":   "Ticket not found",
			"error": err.Error(),
		})
		return
	}
	err = model.Delete()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"msg":   "Delete ticket failed",
			"error": err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"msg": "Delete ticket success",
	})
}

// ticket type CURD
func CreateTicketType(ctx *gin.Context) {
	type form struct {
		Name      string `json:"name"`
		Introduce string `json:"introduce"`
	}
	var req form
	var model model.TicketsType
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"msg":   "Invalid form",
			"error": err.Error(),
		})
		return
	}
	model.Name = req.Name
	model.Introduce = req.Introduce
	err := model.Create()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"msg":   "Create ticket type failed",
			"error": err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"msg": "Create ticket type success",
	})
}

func UpdateTicketType(ctx *gin.Context) {
	id := ctx.Param("id")
	type form struct {
		Name      string `json:"name"`
		Introduce string `json:"introduce"`
	}
	var req form
	var model model.TicketsType
	_, err := model.Find(map[string]interface{}{"id": id})
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"msg":   "Ticket type not found",
			"error": err.Error(),
		})
		return
	}
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"msg":   "Invalid form",
			"error": err.Error(),
		})
		return
	}
	model.Name = req.Name
	model.Introduce = req.Introduce
	err = model.Update()
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"msg":   "Update ticket type failed",
			"error": err.Error(),
		})
		return
	}
}

func DeleteTicketType(ctx *gin.Context) {
	id := ctx.Param("id")
	var model model.TicketsType
	_, err := model.Find(map[string]interface{}{"id": id})
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"msg":   "Ticket type not found",
			"error": err.Error(),
		})
		return
	}
	err = model.Delete()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"msg":   "Delete ticket type failed",
			"error": err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"msg": "Delete ticket type success",
	})
}

// ticket sale price CURD
func CreateTicketSalePrice(ctx *gin.Context) {
	type form struct {
		TickeketsID   int     `json:"tickets_id"`
		TicketsTypeID int     `json:"tickets_type_id"`
		Price         float64 `json:"price"`
		ExpiresAt     string  `json:"expires_at"`
	}
	var req form
	var m model.TicketsSalePrice
	var ticketModel model.Tickets
	var ticketTypeModel model.TicketsType
	// bind the form to the struct
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"msg":   "Invalid form",
			"error": err.Error(),
		})
		return
	}
	// check the ticket and ticket type exists
	_, err := ticketModel.Find(map[string]interface{}{"id": req.TickeketsID})
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"msg":   "Ticket not found",
			"error": err.Error(),
		})
		return
	}
	_, err = ticketTypeModel.Find(map[string]interface{}{"id": req.TicketsTypeID})
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"msg":   "Ticket type not found",
			"error": err.Error(),
		})
		return
	}
	// create the ticket sale price
	m.TicketsID = req.TickeketsID
	m.TicketsTypeID = req.TicketsTypeID
	m.Price = req.Price
	timeLayout := "2006-01-02 15:04:05"
	expireAt, err := time.Parse(timeLayout, req.ExpiresAt)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"msg":   "Invalid time format",
			"error": err.Error(),
		})
		return
	}
	m.ExpiresAt = expireAt
	err = m.Create()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"msg":   "Create ticket sale price failed",
			"error": err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"msg": "Create ticket sale price success",
	})
}

func UpdateTicketSalePrice(ctx *gin.Context) {
	id := ctx.Param("id")
	type form struct {
		TickeketsID   int     `json:"tickets_id"`
		TicketsTypeID int     `json:"tickets_type_id"`
		Price         float64 `json:"price"`
		ExpiresAt     string  `json:"expires_at"`
	}
	var req form
	var m model.TicketsSalePrice

	_, err := m.Find(map[string]interface{}{"id": id})
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"msg":   "Ticket sale price not found",
			"error": err.Error(),
		})
		return
	}
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"msg":   "Invalid form",
			"error": err.Error(),
		})
		return
	}
	m.TicketsID = req.TickeketsID
	m.TicketsTypeID = req.TicketsTypeID
	m.Price = req.Price
	timeLayout := "2006-01-02 15:04:05"
	expireAt, err := time.Parse(timeLayout, req.ExpiresAt)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"msg":   "Invalid time format",
			"error": err.Error(),
		})
		return
	}
	m.ExpiresAt = expireAt
	err = m.Update()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"msg":   "Update ticket sale price failed",
			"error": err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"msg": "Update ticket sale price success",
	})
}

func DeleteTicketSalePrice(ctx *gin.Context) {
	id := ctx.Param("id")
	var model model.TicketsSalePrice
	_, err := model.Find(map[string]interface{}{"id": id})
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"msg":   "Ticket sale price not found",
			"error": err.Error(),
		})
		return
	}
	err = model.Delete()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"msg":   "Delete ticket sale price failed",
			"error": err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"msg": "Delete ticket sale price success",
	})
}
