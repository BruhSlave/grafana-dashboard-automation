package main

import (
	"net/http"
	"sync"

	"github.com/gin-gonic/gin"
)	

type album struct {
	ID     string  `json:"id"`
	Title  string  `json:"title"`
	Artist string  `json:"artist"`
	Price  float64 `json:"price"`
}

type Store struct {
	mu sync.RWMutex
	items map[string]album
}

func NewStore() *Store {
	return &Store{
		items: make(map[string]album),
	}
}

func main() {
	store := NewStore()
	router := gin.Default()
	router.POST("/items", store.createItem)
	router.GET("/items/:id", getAlbumByID)
	router.PUT("/items/:id", putAlbumByID)

	if err := router.Run("localhost:8080"); err != nil {
		panic(err)
	}
}

func health(c *gin.Context) {
	c.Status(http.StatusOK)
}


func (s *Store) createItem(c *gin.Context) {
	var item album

	if err := c.ShouldBindJSON(&item); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid JSON or missing ID",
		})
		return
	}

	s.mu.Lock()
	defer s.mu.Unlock()

	if _, exists := s.items[item.ID]; exists {
		c.JSON(http.StatusConflict, gin.H{
			"error": "Item already exists",
		})
		return
	}

	s.items[item.ID] = item

	c.JSON(http.StatusCreated, item)
}

func getAlbumByID(c *gin.Context) {
	id := c.Param("id")

	for _, a := range albums {
		if a.ID == id {
			c.IndentedJSON(http.StatusOK, a)
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{
		"message:": "Album not found",
		"code:":    http.StatusNotFound,
	})
}

func putAlbumByID(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error:": "Invalid ID",
			"code:":  http.StatusBadRequest,
		})
		return
	}

	var updateAlbum album
	if err := c.BindJSON(&updateAlbum); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{
			"error:": "Invalid JSON",
			"code:": http.StatusBadRequest,
		})
	}

	for 
}
