package controllers

import (
	"assignment_2/dto"
	"assignment_2/models"
	"encoding/json"
	"fmt"
	"github.com/gomodule/redigo/redis"
	"github.com/labstack/echo/v4"
)

type OrderController struct {
	Base Controller
}

// Index list of resources godoc
// @Summary Index list of resources
// @Description Index list of resources
// @ID order-index
// @Produce json
// @Success 200 {object} Response
// @Failure 500 {object} Response
// @BasePath /api/v1
// @Router /api/v1/order/index [get]
func (c *OrderController) Index(ctx echo.Context) error {
	var result []models.Order
	err := c.Base.Db.Debug().Preload("Items").Find(&result).Error
	if err != nil {
		return c.Base.InternalServerError(ctx, err)
	}
	return c.Base.Data(ctx, result)
}

// Show detail of resource godoc
// @Summary Show detail of resource
// @Description Show detail of resource
// @ID order-show
// @Produce json
// @Success 200 {object} Response
// @Failure 500 {object} Response
// @BasePath /api/v1
// @Router /api/v1/order/show/{id} [get]
// @Param id path string true "Resource primary key"
func (c *OrderController) Show(ctx echo.Context) error {
	var result models.Order
	id := ctx.Param("id")

	redisKey := fmt.Sprintf(c.Base.RedisDb+"orders-%v", id)
	redisVal, err := redis.String(c.Base.Redis.Do("GET", redisKey))
	_ = json.Unmarshal([]byte(redisVal), &result)
	if redisVal != "" {
		return c.Base.Data(ctx, result)
	}

	err = c.Base.Db.Debug().Preload("Items").Find(&result, id).Error
	if err != nil {
		return c.Base.BadRequest(ctx, err.Error())
	}
	go func() {
		_, _ = c.Base.Redis.Do("SET", redisKey, result)
	}()
	return c.Base.Data(ctx, result)
}

// Create new resource godoc
// @Summary Create new resource to be saved to database
// @Description Create new resource to be saved to database
// @ID order-create
// @Accept json,x-www-form-urlencoded
// @Produce json
// @Param order body dto.OrderDTO true "request body json"
// @Success 200 {object} Response
// @Failure 500 {object} Response
// @BasePath /api/v1
// @Router /api/v1/order/create [post]
func (c *OrderController) Create(ctx echo.Context) error {
	orderDTO := new(dto.OrderDTO)
	err := ctx.Bind(orderDTO)
	if err != nil {
		return c.Base.InternalServerError(ctx, err)
	}
	var items []models.Item
	for _, each := range orderDTO.Items {
		item := models.Item{
			Code:      each.Code,
			Desc:      each.Desc,
			Qty:       each.Qty,
		}
		items = append(items, item)
	}
	order := models.Order{
		CustomerName: orderDTO.CustomerName,
		Items: items,
	}
	err = c.Base.Db.Debug().Create(&order).Error
	if err != nil {
		return c.Base.InternalServerError(ctx, err)
	}

	redisKey := fmt.Sprintf(c.Base.RedisDb+"orders-%v", order.ID)
	redisVal, _ := json.Marshal(order)
	go func() {
		_, _ = c.Base.Redis.Do("SET", redisKey, redisVal)
	}()
	return c.Base.Data(ctx, order)
}

// Edit existing resource godoc
// @Summary Edit existing resource
// @Description Edit existing resource
// @ID order-edit
// @Accept json,x-www-form-urlencoded
// @Produce json
// @Param order body dto.OrderDTO true "request body json"
// @Success 200 {object} Response
// @Failure 500 {object} Response
// @BasePath /api/v1
// @Router /api/v1/order/edit/{id} [patch]
// @Param id path string true "Resource primary key"
func (c *OrderController) Edit(ctx echo.Context) error {
	id := ctx.Param("id")
	orderDTO := new(dto.OrderDTO)
	err := ctx.Bind(orderDTO)
	if err != nil {
		return c.Base.InternalServerError(ctx, err)
	}

	var result models.Order
	err = c.Base.Db.Debug().Preload("Items").Where("id = ?", id).First(&result).Error
	if err != nil {
		return c.Base.InternalServerError(ctx, err)
	}

	c.Base.Db.Debug().Unscoped().Where("order_id = ?", result.ID).Delete(&[]models.Item{})

	var items []models.Item
	for _, each := range orderDTO.Items {
		item := models.Item{
			Code:      each.Code,
			Desc:      each.Desc,
			Qty:       each.Qty,
		}
		items = append(items, item)
	}
	result.CustomerName = orderDTO.CustomerName

	err = c.Base.Db.Debug().Where("id = ?", id).Updates(&result).Error
	if err != nil {
		return c.Base.InternalServerError(ctx, err)
	}
	err = c.Base.Db.Debug().Model(&result).Association("Items").Replace(&items)
	if err != nil {
		return c.Base.InternalServerError(ctx, err)
	}

	redisKey := fmt.Sprintf(c.Base.RedisDb+"orders-%v", result.ID)
	redisVal, _ := json.Marshal(result)
	go func() {
		_, _ = c.Base.Redis.Do("SET", redisKey, redisVal)
	}()
	return c.Base.Data(ctx, result)
}

// Delete existing resource godoc
// @Summary Delete existing resource
// @Description Delete existing resource
// @ID order-delete
// @Produce  json
// @Success 200 {object} Response
// @Failure 500 {object} Response
// @BasePath /api/v1
// @Router /api/v1/order/delete/{id} [delete]
// @Param id path string true "Resource primary key"
func (c *OrderController) Delete(ctx echo.Context) error {
	var result models.Order
	id := ctx.Param("id")
	err := c.Base.Db.Debug().Model(&result).Association("Items").Clear()
	if err != nil {
		return c.Base.InternalServerError(ctx, err)
	}
	err = c.Base.Db.Debug().Where("id = ?", id).Delete(&result).Error
	if err != nil {
		return c.Base.BadRequest(ctx, err.Error())
	}
	redisKey := fmt.Sprintf(c.Base.RedisDb+"orders-%v", result.ID)
	go func() {
		_, _ = c.Base.Redis.Do("DEL", redisKey)
	}()
	return c.Base.Ok(ctx)
}
