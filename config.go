package main

import (
    "gorm.io/driver/sqlite"
    "gorm.io/gorm"
    "meuprojeto/models"
)

var DB *gorm.DB

func ConnectDatabase() {
    database, err := gorm.Open(sqlite.Open("banco.db"), &gorm.Config{})
    if err != nil {
        panic("Erro ao conectar ao banco")
    }

    database.AutoMigrate(&models.Sala{})
    DB = database
}