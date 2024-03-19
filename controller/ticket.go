package controller

import (
	"net/http"
	"ticketsale/model"
	"ticketsale/service"
	"time"

	"github.com/gin-gonic/gin"
)

func TicketList(ctx *gin.Context) {
	var model model.Tickets
	ret, err := model.Find(nil)
	if err != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"code":  400,
			"msg":   "Get ticket list failed",
			"error": err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"code": 200,
		"msg":  "Get ticket list success",
		"data": ret,
	})
}

// ticket CURD
func CreateTicket(ctx *gin.Context) {
	type form struct {
		Name      string `json:"name"`
		Introduce string `json:"introduce"`
		Count     int    `json:"count"`
		StartAt   string `json:"start_at"`
		ExpiresAt string `json:"expires_at"`
	}
	var req form
	var model model.Tickets
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"code":  400,
			"msg":   "Invalid form",
			"error": err.Error(),
		})
		return
	}
	uid := service.GetUidFromToken(ctx)
	model.Name = req.Name
	model.Introduce = req.Introduce
	model.Count = req.Count
	model.OwnerID = uid
	timeLayout := "2006-01-02 15:04:05"
	startAt, err := time.Parse(timeLayout, req.StartAt)
	if err != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"code":  400,
			"msg":   "Invalid time format",
			"error": err.Error(),
		})
		return
	}
	expireAt, err := time.Parse(timeLayout, req.ExpiresAt)
	if err != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"code":  400,
			"msg":   "Invalid time format",
			"error": err.Error(),
		})
		return
	}
	model.StartAt = startAt
	model.ExpiresAt = expireAt
	err = model.Create()
	if err != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"code":  400,
			"msg":   "Create ticket failed",
			"error": err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"code": 200,
		"msg":  "Create ticket success",
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
	target, err := model.First(map[string]interface{}{"id": id})
	if err != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"code":  400,
			"msg":   "Ticket not found",
			"error": err.Error(),
		})
		return
	}

	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"code":  400,
			"msg":   "Invalid form",
			"error": err.Error(),
		})
		return
	}
	target.Count = req.Count
	target.Introduce = req.Introduce
	timeLayout := "2006-01-02 15:04:05"
	expireAt, err := time.Parse(timeLayout, req.ExpiresAt)
	if err != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"code":  400,
			"msg":   "Invalid time format",
			"error": err.Error(),
		})
		return
	}
	target.ExpiresAt = expireAt
	err = target.Update()
	if err != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"code":  400,
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
	target, err := model.First(map[string]interface{}{"id": id})
	if err != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"code":  400,
			"msg":   "Ticket not found",
			"error": err.Error(),
		})
		return
	}
	err = target.Delete()
	if err != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"code":  400,
			"msg":   "Delete ticket failed",
			"error": err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"code": 200,
		"msg":  "Delete ticket success",
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
		ctx.JSON(http.StatusOK, gin.H{
			"code":  400,
			"msg":   "Invalid form",
			"error": err.Error(),
		})
		return
	}
	model.Name = req.Name
	model.Introduce = req.Introduce
	err := model.Create()
	if err != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"code":  400,
			"msg":   "Create ticket type failed",
			"error": err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"code": 200,
		"msg":  "Create ticket type success",
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
	target, err := model.First(map[string]interface{}{"id": id})
	if err != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"code":  400,
			"msg":   "Ticket type not found",
			"error": err.Error(),
		})
		return
	}
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"code":  400,
			"msg":   "Invalid form",
			"error": err.Error(),
		})
		return
	}
	target.Name = req.Name
	target.Introduce = req.Introduce
	err = target.Update()
	if err != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"code":  400,
			"msg":   "Update ticket type failed",
			"error": err.Error(),
		})
		return
	}
}

func DeleteTicketType(ctx *gin.Context) {
	id := ctx.Param("id")
	var model model.TicketsType
	target, err := model.First(map[string]interface{}{"id": id})
	if err != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"code":  400,
			"msg":   "Ticket type not found",
			"error": err.Error(),
		})
		return
	}
	err = target.Delete()
	if err != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"code":  400,
			"msg":   "Delete ticket type failed",
			"error": err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"code": 200,
		"msg":  "Delete ticket type success",
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
		ctx.JSON(http.StatusOK, gin.H{
			"code":  400,
			"msg":   "Invalid form",
			"error": err.Error(),
		})
		return
	}
	// check the ticket and ticket type exists
	_, err := ticketModel.First(map[string]interface{}{"id": req.TickeketsID})
	if err != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"code":  400,
			"msg":   "Ticket not found",
			"error": err.Error(),
		})
		return
	}
	_, err = ticketTypeModel.First(map[string]interface{}{"id": req.TicketsTypeID})
	if err != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"code":  400,
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
		ctx.JSON(http.StatusOK, gin.H{
			"code":  400,
			"msg":   "Invalid time format",
			"error": err.Error(),
		})
		return
	}
	m.ExpiresAt = expireAt
	err = m.Create()
	if err != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"code":  400,
			"msg":   "Create ticket sale price failed",
			"error": err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"code": 200,
		"msg":  "Create ticket sale price success",
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

	target, err := m.First(map[string]interface{}{"id": id})
	if err != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"code":  400,
			"msg":   "Ticket sale price not found",
			"error": err.Error(),
		})
		return
	}
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"code":  400,
			"msg":   "Invalid form",
			"error": err.Error(),
		})
		return
	}
	target.TicketsID = req.TickeketsID
	target.TicketsTypeID = req.TicketsTypeID
	target.Price = req.Price
	timeLayout := "2006-01-02 15:04:05"
	expireAt, err := time.Parse(timeLayout, req.ExpiresAt)
	if err != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"code":  400,
			"msg":   "Invalid time format",
			"error": err.Error(),
		})
		return
	}
	target.ExpiresAt = expireAt
	err = m.Update()
	if err != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"code":  400,
			"msg":   "Update ticket sale price failed",
			"error": err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"code": 200,
		"msg":  "Update ticket sale price success",
	})
}

func DeleteTicketSalePrice(ctx *gin.Context) {
	id := ctx.Param("id")
	var model model.TicketsSalePrice
	target, err := model.First(map[string]interface{}{"id": id})
	if err != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"code":  400,
			"msg":   "Ticket sale price not found",
			"error": err.Error(),
		})
		return
	}
	err = target.Delete()
	if err != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"code":  400,
			"msg":   "Delete ticket sale price failed",
			"error": err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"code": 200,
		"msg":  "Delete ticket sale price success",
	})
}
