package controllers

import (
	"buy2play/config"
	"buy2play/models"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"net/http"
)

// GetAllProducts fetches all products REDUNDANT
func GetAllProducts(c *gin.Context) {
	db := config.DB

	var products []models.Product
	if err := db.Preload("ProductSubCategory.ProductCategory").Find(&products).Error; err != nil {
		log.Errorf("Error fetching products: %v", err) // Log error for debugging
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch products"})
		return
	}

	categoryMap := make(map[string]map[string][]models.Product)

	for _, product := range products {
		categoryName := product.ProductSubCategory.ProductCategory.Name
		subcategoryName := product.ProductSubCategory.Name

		if _, exists := categoryMap[categoryName]; !exists {
			categoryMap[categoryName] = make(map[string][]models.Product)
		}

		categoryMap[categoryName][subcategoryName] = append(categoryMap[categoryName][subcategoryName], product)
	}

	var response []struct {
		CategoryName  string `json:"category_name"`
		Subcategories []struct {
			SubcategoryName string           `json:"subcategory_name"`
			Products        []models.Product `json:"products"`
		} `json:"subcategories"`
	}

	for categoryName, subcategories := range categoryMap {
		var subcategoryList []struct {
			SubcategoryName string           `json:"subcategory_name"`
			Products        []models.Product `json:"products"`
		}

		for subcategoryName, products := range subcategories {
			subcategoryList = append(subcategoryList, struct {
				SubcategoryName string           `json:"subcategory_name"`
				Products        []models.Product `json:"products"`
			}{
				SubcategoryName: subcategoryName,
				Products:        products,
			})
		}

		response = append(response, struct {
			CategoryName  string `json:"category_name"`
			Subcategories []struct {
				SubcategoryName string           `json:"subcategory_name"`
				Products        []models.Product `json:"products"`
			} `json:"subcategories"`
		}{
			CategoryName:  categoryName,
			Subcategories: subcategoryList,
		})
	}

	c.JSON(http.StatusOK, response)
}

func GetCategoryProducts(c *gin.Context) {
	var input struct {
		SubCategoryID uint `json:"subcategory_id"`
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request format"})
		return
	}

	db := config.DB

	var products []models.Product
	if err := db.Model(&models.Product{}).Where("product_sub_category_id = ?", input.SubCategoryID).Find(&products).Error; err != nil {
		log.Errorf("Error fetching products: %v", err) // Log error for debugging
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch products"})
		return
	}

	var response []struct {
		ProductID       uint   `json:"product_id"`
		ProductName     string `json:"product_name"`
		ProductPrice    int    `json:"product_price"`
		ProductImageURL string `json:"product_image_url"`
	}

	for _, product := range products {
		response = append(response, struct {
			ProductID       uint   `json:"product_id"`
			ProductName     string `json:"product_name"`
			ProductPrice    int    `json:"product_price"`
			ProductImageURL string `json:"product_image_url"`
		}{
			ProductID:       product.ID,
			ProductName:     product.Name,
			ProductPrice:    product.Price,
			ProductImageURL: product.ImageURL,
		})
	}

	c.JSON(http.StatusOK, response)
}

func GetCategories(c *gin.Context) {
	db := config.DB

	var subcategories []models.ProductSubCategory
	if err := db.Preload("ProductCategory").Find(&subcategories).Error; err != nil {
		log.Errorf("Error fetching subcategories: %v", err) // Log error for debugging
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch subcategories"})
		return
	}

	var response []struct {
		CategoryID    uint   `json:"category_id"`
		CategoryName  string `json:"category_name"`
		Subcategories []struct {
			SubcategoryID   uint   `json:"subcategory_id"`
			SubcategoryName string `json:"subcategory_name"`
		} `json:"subcategories"`
	}

	categoryMap := make(map[uint]*struct {
		CategoryID    uint   `json:"category_id"`
		CategoryName  string `json:"category_name"`
		Subcategories []struct {
			SubcategoryID   uint   `json:"subcategory_id"`
			SubcategoryName string `json:"subcategory_name"`
		} `json:"subcategories"`
	})

	for _, subcategory := range subcategories {
		categoryID := subcategory.ProductCategory.ID

		if _, exists := categoryMap[categoryID]; !exists {
			categoryMap[categoryID] = &struct {
				CategoryID    uint   `json:"category_id"`
				CategoryName  string `json:"category_name"`
				Subcategories []struct {
					SubcategoryID   uint   `json:"subcategory_id"`
					SubcategoryName string `json:"subcategory_name"`
				} `json:"subcategories"`
			}{
				CategoryID:   categoryID,
				CategoryName: subcategory.ProductCategory.Name,
			}
		}

		category := categoryMap[categoryID]
		category.Subcategories = append(category.Subcategories, struct {
			SubcategoryID   uint   `json:"subcategory_id"`
			SubcategoryName string `json:"subcategory_name"`
		}{
			SubcategoryID:   subcategory.ID,
			SubcategoryName: subcategory.Name,
		})
	}

	for _, category := range categoryMap {
		response = append(response, *category)
	}

	// Return the JSON response
	c.JSON(http.StatusOK, response)
}

// GetProduct fetches a single product by ID
func GetProduct(c *gin.Context) {
	var input struct {
		ProductID uint `json:"product_id"`
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request format"})
		return
	}

	db := config.DB

	var product models.Product
	if err := db.Preload("ProductSubCategory.ProductCategory").First(&product, input.ProductID).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Product not found"})
		return
	}

	response := struct {
		ProductID          uint   `json:"product_id"`
		ProductName        string `json:"product_name"`
		ProductPrice       int    `json:"product_price"`
		ProductDescription string `json:"product_description"`
		ProductType        string `json:"product_type"`
		ProductImageURL    string `json:"product_image_url"`
	}{
		ProductID:          product.ID,
		ProductName:        product.Name,
		ProductPrice:       product.Price,
		ProductDescription: product.Description,
		ProductType:        string(product.Type),
		ProductImageURL:    product.ImageURL,
	}

	c.JSON(http.StatusOK, response)
}

// CreateProduct creates a new product
func CreateProduct(c *gin.Context) {
	var productInput struct {
		Name                 string `json:"product_name"`
		Price                int    `json:"product_price"`
		Description          string `json:"product_description"`
		Type                 string `json:"product_type"`
		ImageURL             string `json:"product_image_url"`
		ProductSubCategoryID uint   `json:"product_category_id"`
	}

	if err := c.ShouldBindJSON(&productInput); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request format"})
		return
	}

	validTypes := map[string]models.ProductType{
		string(models.Auto):   models.Auto,
		string(models.Manual): models.Manual,
	}

	productType, exists := validTypes[productInput.Type]
	if !exists {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid product type"})
		return
	}

	if productInput.Name == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Product name is required"})
		return
	}

	if productInput.Description == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Product description is required"})
		return
	}

	if productInput.Price == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Product price is required"})
		return
	}

	if productInput.ProductSubCategoryID == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Product Category ID is required"})
		return
	}

	if productInput.Price <= 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Price must be greater than zero"})
		return
	}

	product := models.Product{
		Name:                 productInput.Name,
		Description:          productInput.Description,
		Price:                productInput.Price,
		Type:                 productType,
		ImageURL:             productInput.ImageURL,
		ProductSubCategoryID: productInput.ProductSubCategoryID,
	}

	db := config.DB
	if err := db.Create(&product).Error; err != nil {
		log.Error("Error creating product: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create product"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"product_id": product.ID})
}

// UpdateProduct updates an existing product
func UpdateProduct(c *gin.Context) {
	var input struct {
		ProductID            uint   `json:"product_id"`
		Name                 string `json:"product_name"`
		Price                int    `json:"product_price"`
		Description          string `json:"product_description"`
		Type                 string `json:"product_type"`
		ImageURL             string `json:"product_image_url"`
		ProductSubCategoryID uint   `json:"product_category_id"`
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request format"})
		return
	}

	validTypes := map[string]models.ProductType{
		string(models.Auto):   models.Auto,
		string(models.Manual): models.Manual,
	}

	productType, exists := validTypes[input.Type]
	if !exists {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid product type"})
		return
	}

	if input.Name == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Product name is required"})
		return
	}

	if input.Description == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Product description is required"})
		return
	}

	if input.Price == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Product price is required"})
		return
	}

	if input.ProductSubCategoryID == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Product Category ID is required"})
		return
	}

	if input.Price <= 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Price must be greater than zero"})
		return
	}

	db := config.DB
	var existingProduct models.Product
	if err := db.Preload("ProductSubCategory.ProductCategory").First(&existingProduct, input.ProductID).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Product not found"})
		return
	}

	existingProduct.Name = input.Name
	existingProduct.Description = input.Description
	existingProduct.Price = input.Price
	existingProduct.Type = productType
	existingProduct.ImageURL = input.ImageURL
	existingProduct.ProductSubCategoryID = input.ProductSubCategoryID

	if err := db.Save(&existingProduct).Error; err != nil {
		log.Errorf("Error updating product: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update product"})
		return
	}

	c.JSON(http.StatusOK, existingProduct)
}

// DeleteProduct deletes a product
func DeleteProduct(c *gin.Context) {
	var input struct {
		ProductID uint `json:"product_id"`
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request format"})
		return
	}

	db := config.DB

	var product models.Product
	if err := db.First(&product, input.ProductID).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Product not found"})
		return
	}

	if err := db.Delete(&product).Error; err != nil {
		log.Errorf("Error deleting product: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete product"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Product deleted"})
}
