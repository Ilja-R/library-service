package http


func (s *Server) endpoints() {
	s.router.GET("/ping", s.Ping)

	

	apiG := s.router.Group("/api", s.checkUserAuthentication)
	{
		apiG.GET("/books", s.GetAllBooks)
		apiG.GET("/books/:id", s.GetBookByID)
		apiG.POST("/books", s.checkIsAdmin, s.CreateProduct)
		apiG.PUT("/books/:id", s.checkIsAdmin, s.UpdateProductByID)
		apiG.DELETE("/books/:id", s.checkIsAdmin, s.DeleteProductByID)
	}

}