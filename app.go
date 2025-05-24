package main

import (
    "RESERVA-DE-SALAS-API/sala"
    "github.com/gin-gonic/gin"
)

func main() {
    config.ConnectDatabase()

    r := gin.Default()

    // Ativar rotas de sala
    routes.SalaRoutes(r)

    r.Run(":6000")
}
