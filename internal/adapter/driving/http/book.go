package http

import (
	"net/http"
	"strconv"

	"github.com/Ilja-R/library-service/internal/domain"
	"github.com/Ilja-R/library-service/internal/errs"
	"github.com/gin-gonic/gin"
)


type Book struct{

}
func (b *Book) FromDomain(dbBook domain.Book) {
	
}



func (s*Server)Ping(c*gin.Context){
	c.JSON(http.StatusOK,gin.H{
		"ping":"pong",
	})
}



func (s*Server)GetAllBooks(c*gin.Context){
	userID := c.GetInt(userIDCtx)
	if userID == 0 {
		c.JSON(http.StatusBadRequest, CommonError{Error: "invalid userID in context"})
		return
	}

	dBooks, err := s.uc.BookGetter.GetAllBooks(c)
	if err != nil {
		s.handleError(c, err)
		return
	}

	var (
		books []Book
		book  Book
	)
	for _, dBook := range dBooks {
		book.FromDomain(dBook)
		books = append(books, book)
	}

	c.JSON(http.StatusOK, books)
}


func (s*Server)GetBookByID(c*gin.Context){
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil || id < 1 {
		s.handleError(c, errs.ErrInvalidBookID)
		return
	}
	dBook, err := s.uc.BookGetter.GetBookByID(c, id)
	if err != nil {
		s.handleError(c, err)
		return
	}

	var book Book
	book.FromDomain(dBook)

	c.JSON(http.StatusOK, book)
}