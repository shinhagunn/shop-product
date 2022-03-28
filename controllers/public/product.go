package public

import (
	"context"
	"reflect"

	"github.com/gofiber/fiber/v2"
	"github.com/kamva/mgm/v3/operator"
	"github.com/shinhagunn/shop-product/config/collection"
	"github.com/shinhagunn/shop-product/controllers"
	"github.com/shinhagunn/shop-product/models"
	"github.com/shinhagunn/shop-product/utils"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func checkOrderbyInProductSort(orderby string) bool {
	x := &ProductSort{
		Name:     "0",
		Price:    2.3,
		Discount: 9.3,
	}

	for i := 0; i < reflect.ValueOf(*x).NumField(); i++ {
		if reflect.ValueOf(*x).Type().Field(i).Name == utils.UpperFirstLetter(orderby) {
			return true
		}
	}
	return false
}

func GetProductByID(c *fiber.Ctx) error {
	id := c.Params("id")
	product := new(models.Product)

	if err := collection.Product.FindByID(id, product); err != nil {
		return c.Status(422).JSON(controllers.FailedConnectDataInDatabase)
	}

	return c.JSON(product)
}

func GetCategories(c *fiber.Ctx) error {
	categories := []models.Category{}

	if err := collection.Category.SimpleFind(&categories, bson.M{}); err != nil {
		return c.Status(422).JSON(controllers.FailedConnectDataInDatabase)
	}

	return c.JSON(categories)
}

type ProductQuery struct {
	Name     string `query:"name"`
	Category string `query:"category"`
	Order    string `query:"order"`
	Orderby  string `query:"orderby"`
	Limit    int    `query:"limit"`
}

type ProductSort struct {
	Name     string
	Price    float64
	Discount float64
}

func GetProducts(c *fiber.Ctx) error {
	products := []models.Product{}
	queries := new(ProductQuery)

	if err := c.QueryParser(queries); err != nil {
		return c.Status(422).JSON(controllers.FailedToParseBody)
	}

	filter := bson.M{}

	if len(queries.Name) != 0 {
		filter["name"] = bson.M{operator.Regex: queries.Name}
	}

	if len(queries.Category) != 0 {
		r, err := primitive.ObjectIDFromHex(queries.Category)

		if err != nil {
			return c.Status(500).JSON(controllers.ServerInternalError)
		}

		filter["category_id"] = r
	}

	// SORT PRODUCTS
	if queries.Orderby != "" && queries.Order != "" && checkOrderbyInProductSort(queries.Orderby) {
		var orderNumber int64
		if queries.Order == "asc" {
			orderNumber = 1
		} else if queries.Order == "desc" {
			orderNumber = -1
		}

		findOptions := options.Find()
		findOptions.SetSort(bson.D{{queries.Orderby, orderNumber}})

		cursor, err := collection.Product.Find(context.TODO(), filter, findOptions)

		if err != nil {
			panic(err)
		}

		result := []bson.M{}
		if err := cursor.All(context.TODO(), &result); err != nil {
			return c.Status(422).JSON(controllers.FailedConnectDataInDatabase)
		}

		if queries.Limit != 0 {
			return c.JSON(result[0:queries.Limit])
		}

		return c.JSON(result)
	}

	if err := collection.Product.SimpleFind(&products, filter); err != nil {
		return c.Status(422).JSON(controllers.FailedConnectDataInDatabase)
	}

	if queries.Limit != 0 {
		return c.JSON(products[0:queries.Limit])
	}

	return c.JSON(products)
}
