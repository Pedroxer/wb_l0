package api

import (
	"encoding/json"
	"github.com/Pedroxer/wbL0/db"
	"github.com/gin-gonic/gin"
	"net/http"
)

type uid struct {
	Order_uid string `json:"order_uid"`
}

func (serv *Server) getJson(ctx *gin.Context) {
	var id uid
	if err := ctx.ShouldBindJSON(&id); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	obj, b := serv.Cache.Get(id.Order_uid)
	if !b {
		ctx.JSON(http.StatusInternalServerError, "Non existing record")
		return
	}
	ctx.JSON(http.StatusOK, obj)
}

func (serv *Server) postJson(ctx *gin.Context) {
	var poston db.Order
	if err := ctx.ShouldBindJSON(&poston); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if (poston.Del == db.Delivery{
		Name:    "",
		Phone:   "",
		Zip:     "",
		City:    "",
		Address: "",
		Region:  "",
		Email:   "",
	}) ||
		(poston.Item == nil) ||
		(poston.Payment == db.Payments{
			Transaction:   "",
			Request_id:    "",
			Currency:      "",
			Provider:      "",
			Amount:        0,
			Payment_dt:    0,
			Bank:          "",
			Delivery_cost: 0,
			Goods_total:   0,
			Custom_fee:    0,
		}) {
		ctx.JSON(http.StatusBadRequest, "Invalid input")
		return
	}
	jposton, _ := json.Marshal(poston)
	serv.Con.Publish("foo", jposton)
	var p db.Order
	json.Unmarshal(jposton, &p)
	ctx.JSON(http.StatusOK, p)
}
