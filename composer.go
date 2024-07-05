package main

import (
	"github.com/gin-gonic/gin"

	youtube_service "youtube_service_arq_hex/youtube-service"
	youtube_adapt_drivens "youtube_service_arq_hex/youtube-service/adapters/drivens"
	youtube_adapt_drivers "youtube_service_arq_hex/youtube-service/adapters/drivers"
	youtube_ports_drivers "youtube_service_arq_hex/youtube-service/ports/drivers"

)

func Compose() youtube_ports_drivers.ForYouTube {
	// Create gin context
	gc := &gin.Context{}
	 // YouTube adapters Drivens
	youtubeAdaptDrivens:= youtube_adapt_drivens.NewYoutubeSearchAdapter(gc)

	// YouTube Service
 	youtubeService := youtube_service.NewYouTubeService(youtubeAdaptDrivens)

	// Search YouTube Adapter Drivers
	searchYouTubeAdapterDrivers := youtube_adapt_drivers.NewSearchYouTubeAdapter(gc, youtubeService)

	return searchYouTubeAdapterDrivers
}
