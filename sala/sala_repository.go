package sala

import (
    "net/http"
    "strconv"

    "github.com/gin-gonic/gin"
    "RESERVA-DE-SALAS-API/sala"
)

// Buscar todas as salas
func GetSalas(c *gin.Context) {
    var salas []models.Sala
    config.DB.Find(&salas)
    c.JSON(http.StatusOK, salas)
}

// Criar nova sala
func CreateSala(c *gin.Context) {
    var sala models.Sala
    if err := c.ShouldBindJSON(&sala); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }
    config.DB.Create(&sala)
    c.JSON(http.StatusCreated, sala)
}

// Buscar sala por ID
func GetSalaByID(c *gin.Context) {
    id, _ := strconv.Atoi(c.Param("id"))
    var sala models.Sala
    result := config.DB.First(&sala, id)
    if result.Error != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "Sala não encontrada"})
        return
    }
    c.JSON(http.StatusOK, sala)
}

// Atualizar sala
func UpdateSala(c *gin.Context) {
    id, _ := strconv.Atoi(c.Param("id"))
    var sala models.Sala
    if err := config.DB.First(&sala, id).Error; err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "Sala não encontrada"})
        return
    }

    if err := c.ShouldBindJSON(&sala); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    config.DB.Save(&sala)
    c.JSON(http.StatusOK, sala)
}

// Deletar sala
func DeleteSala(c *gin.Context) {
    id, _ := strconv.Atoi(c.Param("id"))
    if err := config.DB.Delete(&models.Sala{}, id).Error; err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }
    c.JSON(http.StatusOK, gin.H{"message": "Sala deletada com sucesso"})
}