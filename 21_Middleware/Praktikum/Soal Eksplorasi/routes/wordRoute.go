package routes

import (
    "database/sql"
    "github.com/labstack/echo/v4"
    "post-comment/controllers"
)

func SetupRoutes(e *echo.Echo, db *sql.DB) {
    controllers.InitializeDB(db)

    // Routes untuk postingan
    e.GET("/api/v1/posts", controllers.GetAllPosts)
    e.GET("/api/v1/posts/:id", controllers.GetPostByID)
    e.POST("/api/v1/posts", controllers.CreatePost)
    e.PUT("/api/v1/posts/:id", controllers.UpdatePost)
    e.DELETE("/api/v1/posts/:id", controllers.DeletePost)

    // Routes untuk komentar
    e.POST("/api/v1/comments/:post_id", controllers.CreateComment)
    e.PUT("/api/v1/comments/:id", controllers.UpdateComment)
    e.DELETE("/api/v1/comments/:id", controllers.DeleteComment)
}
