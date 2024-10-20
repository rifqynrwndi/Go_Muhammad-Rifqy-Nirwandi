package controllers

import (
    "database/sql"
    "net/http"
    "github.com/labstack/echo/v4"
    "post-comment/models"
)

var db *sql.DB

func InitializeDB(database *sql.DB) {
    db = database
}

func GetAllPosts(c echo.Context) error {
    rows, err := db.Query("SELECT * FROM posts")
    if err != nil {
        return c.JSON(http.StatusInternalServerError, echo.Map{"error": "Unable to fetch posts"})
    }
    defer rows.Close()

    var posts []models.Post
    for rows.Next() {
        var post models.Post
        if err := rows.Scan(&post.ID, &post.Title, &post.Content); err != nil {
            return err
        }
        posts = append(posts, post)
    }
    return c.JSON(http.StatusOK, posts)
}

func GetPostByID(c echo.Context) error {
    id := c.Param("id")
    var post models.Post
    err := db.QueryRow("SELECT * FROM posts WHERE id = ?", id).Scan(&post.ID, &post.Title, &post.Content)
    if err != nil {
        return c.JSON(http.StatusNotFound, echo.Map{"error": "Post not found"})
    }

    rows, err := db.Query("SELECT * FROM comments WHERE post_id = ?", post.ID)
    if err != nil {
        return c.JSON(http.StatusInternalServerError, echo.Map{"error": "Unable to fetch comments"})
    }
    defer rows.Close()

    var comments []models.Comment
    for rows.Next() {
        var comment models.Comment
        if err := rows.Scan(&comment.ID, &comment.PostID, &comment.Content); err != nil {
            return err
        }
        comments = append(comments, comment)
    }
    post.Comments = comments
    return c.JSON(http.StatusOK, post)
}

func CreatePost(c echo.Context) error {
    var post models.Post
    if err := c.Bind(&post); err != nil {
        return c.JSON(http.StatusBadRequest, echo.Map{"error": "Invalid request"})
    }
    result, err := db.Exec("INSERT INTO posts (title, content) VALUES (?, ?)", post.Title, post.Content)
    if err != nil {
        return c.JSON(http.StatusInternalServerError, echo.Map{"error": "Failed to create post"})
    }
    id, _ := result.LastInsertId()
    post.ID = int(id)
    return c.JSON(http.StatusCreated, post)
}

func UpdatePost(c echo.Context) error {
    id := c.Param("id")
    var post models.Post
    if err := c.Bind(&post); err != nil {
        return c.JSON(http.StatusBadRequest, echo.Map{"error": "Invalid request"})
    }
    _, err := db.Exec("UPDATE posts SET title = ?, content = ? WHERE id = ?", post.Title, post.Content, id)
    if err != nil {
        return c.JSON(http.StatusInternalServerError, echo.Map{"error": "Failed to update post"})
    }
    return c.JSON(http.StatusOK, post)
}

func DeletePost(c echo.Context) error {
    id := c.Param("id")
    _, err := db.Exec("DELETE FROM posts WHERE id = ?", id)
    if err != nil {
        return c.JSON(http.StatusNotFound, echo.Map{"error": "Post not found"})
    }
    return c.JSON(http.StatusOK, echo.Map{"message": "Deleted successfully"})
}
