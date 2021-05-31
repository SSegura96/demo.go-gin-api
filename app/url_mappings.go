package app

import "KPA/controllers/status"

func mapUrls() {
	router.GET("/status", status.Status)

}
