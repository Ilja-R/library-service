package http

import (
	"net/http"

	"github.com/Ilja-R/library-service/internal/domain"
	"github.com/Ilja-R/library-service/pkg"
	"github.com/gin-gonic/gin"
)

const (
	authorizationHeader = "Authorization"
	userIDCtx           = "userID"
	userRoleCtx         = "userRole"
)

func (s *Server) checkUserAuthentication(c *gin.Context) {
	token, err := s.extractTokenFromHeader(c, authorizationHeader)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusUnauthorized, CommonError{Error: err.Error()})
		return
	}

	userID, isRefresh, userRole, err := pkg.ParseToken(token)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusUnauthorized, CommonError{Error: err.Error()})
		return
	}

	if isRefresh {
		c.AbortWithStatusJSON(http.StatusUnauthorized, CommonError{Error: "inappropriate token"})
		return
	}

	c.Set(userIDCtx, userID)
	c.Set(userRoleCtx, string(userRole))
}

func (s *Server) checkIsAdmin(c *gin.Context) {
	role := c.GetString(userRoleCtx)
	if role == "" {
		c.AbortWithStatusJSON(http.StatusUnauthorized, CommonError{Error: "role is not in context"})
		return
	}

	if role != domain.RoleAdmin {
		c.AbortWithStatusJSON(http.StatusForbidden, CommonError{Error: "permission denied"})
		return
	}

	c.Next()
}