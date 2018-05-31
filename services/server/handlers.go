package server

import (
	"log"
	"net/http"
	"sync"

	"github.com/asaskevich/govalidator"
	"github.com/gin-gonic/gin"
	"github.com/kirill-a-belov/test-sg/models"
)

type ProductMap struct {
	sync.Mutex
	m map[string]*models.Product
}

func (s *Server) NewByURLHandler() func(c *gin.Context) {
	return func(c *gin.Context) {

		req := &models.ByURLRequest{}
		if err := c.BindJSON(req); err != nil {
			log.Print("error in NewByURLHandler", err)
			// 400 error was sended by c.BindJSON itself
			return
		}

		// Concurrency result processing
		res := &ProductMap{
			m: make(map[string]*models.Product),
		}

		wg := sync.WaitGroup{}
		wg.Add(len(req.URLs))

		for _, url := range req.URLs {
			go func(url string) {
				product, err := ScrapURL(url)
				if err != nil {
					log.Print("error in NewByURLHandler", err)
					return
				}

				if err := s.storage.SaveProduct(product); err != nil {
					log.Print("error in NewByURLHandler", err)
					return // Return or continue - TBD
				}

				res.Lock()
				res.m[url] = product
				res.Unlock()
				wg.Done()

			}(url)
		}

		wg.Wait()
		resp := &models.ByURLResponse{}
		for _, url := range req.URLs {
			resp.Products = append(resp.Products, res.m[url])
		}

		c.JSON(http.StatusOK, resp)
		return
	}
}

func (s *Server) NewByPIDHandler() func(c *gin.Context) {
	return func(c *gin.Context) {
		req := c.Param("pid")
		if !govalidator.IsUUID(req) {
			err := &models.Error{
				Code:    models.ErrUnknownRequestCode,
				Message: "Invalid request",
			}

			c.JSON(http.StatusBadRequest, err)
			return
		}

		res, err := s.storage.GetProduct(req)
		if err != nil {
			log.Print("error in NewByPIDHandler", err)

			err2 := &models.Error{
				Code:    models.ErrUnknownErrorCode,
				Message: "Server error",
			}

			c.JSON(http.StatusInternalServerError, err2)
			return
		}

		c.JSON(http.StatusOK, &models.ByIDResponse{Product: res})
		return
	}
}
