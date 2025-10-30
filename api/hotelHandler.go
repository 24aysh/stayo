package api

import (
	"hotel-reservation/db"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/v2/bson"
)

type HotelHandler struct {
	store *db.Store
}
type HotelQueryParams struct {
	Rooms  bool
	Rating int
}

func NewHotelHandler(store *db.Store) *HotelHandler {
	return &HotelHandler{
		store: store,
	}
}

func (h *HotelHandler) HandleGetRooms(c *fiber.Ctx) error {
	id := c.Params("id")
	oid, err := bson.ObjectIDFromHex(id)
	if err != nil {
		return err
	}
	filter := bson.M{
		"hotelId": oid,
	}
	rooms, err := h.store.Room.GetRooms(c.Context(), filter)
	if err != nil {
		return nil
	}
	return c.JSON(rooms)
}
func (h *HotelHandler) HandleGetHotels(c *fiber.Ctx) error {
	var qParams HotelQueryParams
	if err := c.QueryParser(&qParams); err != nil {
		return err
	}
	hotels, err := h.store.Hotel.GetHotels(c.Context(), nil)
	if err != nil {
		return err
	}
	return c.JSON(hotels)
}
func (h *HotelHandler) HandleGetHotel(c *fiber.Ctx) error {
	id := c.Params("id")
	oid, err := bson.ObjectIDFromHex(id)
	if err != nil {
		return err
	}
	filter := bson.M{
		"_id": oid,
	}
	hotel, err := h.store.Hotel.GetHotel(c.Context(), filter)
	if err != nil {
		return err
	}
	return c.JSON(hotel)
}
