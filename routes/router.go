package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/shinhagunn/shop-product/controllers/admin"
	"github.com/shinhagunn/shop-product/controllers/public"
	"github.com/shinhagunn/shop-product/controllers/resource"
	"github.com/shinhagunn/shop-product/routes/middlewares"
)

func InitRouter() {
	app := fiber.New()

	app.Use(cors.New(cors.Config{
		AllowOrigins: "*",
		AllowHeaders: "Origin, Content-Type, Accept",
	}))

	app.Use(logger.New(logger.Config{
		Format: "[${ip}]:${port} ${status} - ${method} ${path}\n",
	}))

	api_public := app.Group("/api/v2/public")
	{
		// Get all slides
		api_public.Get("/slides", public.GetSlides)
		// Get products
		api_public.Get("/products", public.GetProducts)
		// Get product by id
		api_public.Get("/product/:id", public.GetProductByID)
		// Get categories
		api_public.Get("/categories", public.GetCategories)
	}

	api_resource := app.Group("/api/v2/resource", middlewares.CheckRequest)
	{
		// Add product to cart
		api_resource.Post("/user/cart", resource.AddProductToCart)
		// Get all product in cart
		api_resource.Get("/user/cart", resource.GetCartProducts)
		// Remove product in cart
		api_resource.Delete("/user/cart/:id", resource.RemoveProductInCart)
		// Get comments by product
		api_public.Get("/product/:id/comments", resource.GetCommentsInProduct)
		// Add comment
		api_resource.Post("/product/:id/comment", resource.CreateComment)
		// Like comment
		api_resource.Get("/comment/:id/like", resource.LikeComment)
		// Dislike comment
		api_resource.Get("/comment/:id/dislike", resource.DislikeComment)
		// Add order
		api_resource.Get("/user/order", resource.HandleOrder)
	}

	api_admin := app.Group("/api/v2/admin", middlewares.CheckRequest)
	{
		// Add category
		api_admin.Post("/category", middlewares.MustAdmin, admin.CreateCategory)
		// Delete category
		api_admin.Delete("/category/:id", middlewares.MustAdmin, admin.DeleteCategory)
		// Add product
		api_admin.Post("/product", middlewares.MustAdmin, admin.CreateProduct)
		// Update product
		api_admin.Post("/product/:id", middlewares.MustAdmin, admin.UpdateProduct)
		// Delete product
		api_admin.Delete("/product/:id", middlewares.MustAdmin, admin.DeleteProduct)
		// Add slide
		api_admin.Post("/slide", middlewares.MustAdmin, admin.CreateSlide)
		// Update slide
		api_admin.Post("/slide/:id", middlewares.MustAdmin, admin.UpdateSlide)
		// Delete slide
		api_admin.Delete("/slide/:id", middlewares.MustAdmin, admin.DeleteSlide)
	}

	app.Listen(":3003")
}
