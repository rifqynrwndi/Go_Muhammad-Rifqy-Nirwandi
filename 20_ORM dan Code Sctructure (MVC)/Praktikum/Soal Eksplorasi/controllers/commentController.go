package controllers

import (
	_ "database/sql"
	"strconv"

	"net/http"
	"post-comment/models"

	"github.com/labstack/echo/v4"
)

func CreateComment(c echo.Context) error {
    
    postID := c.Param("post_id")
    var comment models.Comment
    if err := c.Bind(&comment); err != nil {
        return c.JSON(http.StatusBadRequest, echo.Map{"error": "Invalid request"})
    }

    postIDInt, err := strconv.Atoi(postID)
    if err != nil {
        return c.JSON(http.StatusBadRequest, echo.Map{"error": "Invalid post ID"})
    }

    comment.PostID = postIDInt
    result, err := db.Exec("INSERT INTO comments (post_id, content) VALUES (?, ?)", comment.PostID, comment.Content)
    if err != nil {
        return c.JSON(http.StatusInternalServerError, echo.Map{"error": "Failed to create comment"})
    }
    id, _ := result.LastInsertId()
    comment.ID = int(id)
    return c.JSON(http.StatusCreated, comment)
}

func UpdateComment(c echo.Context) error {
    id := c.Param("id")
    var comment models.Comment
    if err := c.Bind(&comment); err != nil {
        return c.JSON(http.StatusBadRequest, echo.Map{"error": "Invalid request"})
    }
    _, err := db.Exec("UPDATE comments SET content = ? WHERE id = ?", comment.Content, id)
    if err != nil {
        return c.JSON(http.StatusInternalServerError, echo.Map{"error": "Failed to update comment"})
    }
    return c.JSON(http.StatusOK, comment)
}

func DeleteComment(c echo.Context) error {
    id := c.Param("id")
    _, err := db.Exec("DELETE FROM comments WHERE id = ?", id)
    if err != nil {
        return c.JSON(http.StatusNotFound, echo.Map{"error": "Comment not found"})
    }
    return c.JSON(http.StatusOK, echo.Map{"message": "Deleted successfully"})
}
