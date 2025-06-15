// api/handler.go
package api

import "gorm.io/gorm"

type Handler struct {
	DB *gorm.DB
}
