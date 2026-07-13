package shared

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

func Err_handler_middleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()

		if len(c.Errors) == 0 {
			return
		}

		err := c.Errors.Last().Err

		var validation_errs validator.ValidationErrors
		var not_found *Not_fount_error

		switch {
		case errors.As(err, &validation_errs):
			fields := make([]gin.H, 0, len(validation_errs))
			for _, fe := range validation_errs {
				fields = append(fields, gin.H{
					"field":   fe.Field(),
					"message": fe.Error(),
				})
			}

			c.JSON(http.StatusBadRequest, gin.H{
				"title":             "validation error",
				"details":           "validation request error",
				"status":            http.StatusBadRequest,
				"validation_errors": fields,
			})
		case errors.As(err, &not_found):
			c.JSON(http.StatusNotFound, gin.H{
				"title":   "resurce not found",
				"details": not_found.Error(),
				"status":  http.StatusNotFound,
			})

		default:
			c.JSON(http.StatusInternalServerError, gin.H{
				"title":   "internal server error",
				"details": "unexpected error occured",
				"status":  http.StatusInternalServerError,
			})

		}

	}
}
