package youtube_service

import (
	"github.com/gin-gonic/gin"

	ports "youtube_service_arq_hex/youtube-service/ports/drivens"
	models "youtube_service_arq_hex/youtube-service/models"

)

type YouTubeService struct {
	searchYouTube ports.ForSearchYouTube
}

// functions
func (yts *YouTubeService) GetSearch(gc *gin.Context) (*[]models.DataYT, error) {
	dataYt, err := yts.searchYouTube.GetSearch(gc)
	if err != nil {
		gc.JSON(500, gin.H{
			"error": err.Error(),
		})
		return nil, err
	}

	gc.JSON(200, dataYt)
	return dataYt, nil
}

// New Youtube Service
func NewYouTubeService(searchYouTube ports.ForSearchYouTube) *YouTubeService {
	return &YouTubeService{
		searchYouTube: searchYouTube,
	}
}
