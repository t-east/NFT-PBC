package presenter

func Count(artLogs entities.ArtLogs) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Header("Access-Control-Allow-Origin", "*")
        c.JSON(http.StatusOK, logs)
	}
}