package ports

import (
	"github.com/gin-gonic/gin"

	models "youtube_service_arq_hex/youtube-service/models"

)

type ForYouTube interface {
	GetSearch(gc *gin.Context) (*[]models.DataYT, error)
}
