package sala

import (
    "github.com/gin-gonic/gin"
    "RESERVA-DE-SALAS-API/sala"
)

func SalaRoutes(router *gin.Engine) {
    salaGroup := router.Group("/salas")
    {
        salaGroup.GET("/", sala.GetSalas)
        salaGroup.POST("/", sala.CreateSala)
        salaGroup.GET("/:id", sala.GetSalaByID)
        salaGroup.PUT("/:id", sala.UpdateSala)
        salaGroup.DELETE("/:id", sala.DeleteSala)
    }
}