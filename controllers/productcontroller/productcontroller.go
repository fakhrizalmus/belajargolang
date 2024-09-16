package productcontroller

import (
	"net/http"

	"github.com/fakhrizalmus/belajargolang/models"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func Index(c *gin.Context) {
	var products []models.Product
	models.DB.Find(&products)
	c.JSON(http.StatusOK, gin.H{"message": "Success", "products": products})
}

func Show(c *gin.Context) {
	var product models.Product
	id := c.Param("id")
	if err := models.DB.First(&product, id).Error; err != nil {
		switch err {
		case gorm.ErrRecordNotFound:
			c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"message": "Data Tidak Ditemukan"})
			return
		default:
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
			return
		}
	}

	c.JSON(http.StatusOK, gin.H{"produk": product})
}

func Create(c *gin.Context) {
	var product models.Product
	if err := c.ShouldBindJSON(&product); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	models.DB.Create(&product)
	c.JSON(http.StatusOK, gin.H{"produk": product})
}

func Edit(c *gin.Context) {
	var product models.Product
	id := c.Param("id")
	if err := c.ShouldBindJSON(&product); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	if models.DB.Model(&product).Where("id = ?", id).Updates(&product).RowsAffected == 0 {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "Edit Gagal"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Berhasil", "data": product})
}

func Delete(c *gin.Context) {
	var product models.Product
	id := c.Param("id")
	//cek id
	if id == "" {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "ID tidak valid"})
		return
	}
	//query cari id produk
	if models.DB.First(&product, id).RowsAffected == 0 {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"message": "Produk tidak ditemukan"})
		return
	}
	//hapus
	if models.DB.Delete(&product).RowsAffected == 0 {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "Gagal Hapus"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Berhasil", "data": product})
}
