package controllers

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/nakshatrabhatt/go-form-api/database"
	"github.com/nakshatrabhatt/go-form-api/models"
	"gorm.io/gorm"
)

// CreateFormRequest represents the request to create a form
type CreateFormRequest struct {
	Title       string             `json:"title" binding:"required"`
	Description string             `json:"description"`
	Fields      []CreateFieldInput `json:"fields" binding:"required,dive"`
	// Original fields from forms.go
	Field1    string `json:"field1"`
	Field2    string `json:"field2"`
	Field3    string `json:"field3"`
	Field4    string `json:"field4"`
	Field5    string `json:"field5"`
	Field6    string `json:"field6"`
	Field7    string `json:"field7"`
	Field8    string `json:"field8"`
	Field9    string `json:"field9"`
	Field10   string `json:"field10"`
	Field11   string `json:"field11"`
	Field12   string `json:"field12"`
	Field13   string `json:"field13"`
	Field14   string `json:"field14"`
	Field15   string `json:"field15"`
	Field16   string `json:"field16"`
	Field17   string `json:"field17"`
	Field18   string `json:"field18"`
	Field19   string `json:"field19"`
	Field20   string `json:"field20"`
	Field21   string `json:"field21"`
	Field22   string `json:"field22"`
	Field23   string `json:"field23"`
	Field24   string `json:"field24"`
	Field25   string `json:"field25"`
	Field26   string `json:"field26"`
	Field27   string `json:"field27"`
	Field28   string `json:"field28"`
	Field29   string `json:"field29"`
	Field30   string `json:"field30"`
}

// CreateFieldInput represents a field in create form request
type CreateFieldInput struct {
	Label       string `json:"label" binding:"required"`
	Type        string `json:"type" binding:"required"`
	Required    bool   `json:"required"`
	Placeholder string `json:"placeholder"`
	Options     string `json:"options"`
}

// UpdateFormRequest represents the request to update a form
type UpdateFormRequest struct {
	Title       string             `json:"title"`
	Description string             `json:"description"`
	Fields      []CreateFieldInput `json:"fields"`
	// Original fields from forms.go
	Field1    string `json:"field1"`
	Field2    string `json:"field2"`
	Field3    string `json:"field3"`
	Field4    string `json:"field4"`
	Field5    string `json:"field5"`
	Field6    string `json:"field6"`
	Field7    string `json:"field7"`
	Field8    string `json:"field8"`
	Field9    string `json:"field9"`
	Field10   string `json:"field10"`
	Field11   string `json:"field11"`
	Field12   string `json:"field12"`
	Field13   string `json:"field13"`
	Field14   string `json:"field14"`
	Field15   string `json:"field15"`
	Field16   string `json:"field16"`
	Field17   string `json:"field17"`
	Field18   string `json:"field18"`
	Field19   string `json:"field19"`
	Field20   string `json:"field20"`
	Field21   string `json:"field21"`
	Field22   string `json:"field22"`
	Field23   string `json:"field23"`
	Field24   string `json:"field24"`
	Field25   string `json:"field25"`
	Field26   string `json:"field26"`
	Field27   string `json:"field27"`
	Field28   string `json:"field28"`
	Field29   string `json:"field29"`
	Field30   string `json:"field30"`
}

// CreateForm handles form creation
func CreateForm(c *gin.Context) {
	var request CreateFormRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	
	// Get user ID from context (set by auth middleware)
	userID, exists := c.Get("user_id")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}
	
	// Create form
	form := models.Form{
		Title:       request.Title,
		Description: request.Description,
		UserID:      userID.(uint),
		// Original fields
		Field1:  request.Field1,
		Field2:  request.Field2,
		Field3:  request.Field3,
		Field4:  request.Field4,
		Field5:  request.Field5,
		Field6:  request.Field6,
		Field7:  request.Field7,
		Field8:  request.Field8,
		Field9:  request.Field9,
		Field10: request.Field10,
		Field11: request.Field11,
		Field12: request.Field12,
		Field13: request.Field13,
		Field14: request.Field14,
		Field15: request.Field15,
		Field16: request.Field16,
		Field17: request.Field17,
		Field18: request.Field18,
		Field19: request.Field19,
		Field20: request.Field20,
		Field21: request.Field21,
		Field22: request.Field22,
		Field23: request.Field23,
		Field24: request.Field24,
		Field25: request.Field25,
		Field26: request.Field26,
		Field27: request.Field27,
		Field28: request.Field28,
		Field29: request.Field29,
		Field30: request.Field30,
	}
	
	// Begin a transaction
	tx := database.DB.Begin()
	
	if err := tx.Create(&form).Error; err != nil {
		tx.Rollback()
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create form"})
		return
	}
	
	// Add fields to the form
	for _, fieldInput := range request.Fields {
		field := models.FormField{
			FormID:      form.ID,
			Label:       fieldInput.Label,
			Type:        fieldInput.Type,
			Required:    fieldInput.Required,
			Placeholder: fieldInput.Placeholder,
			Options:     fieldInput.Options,
		}
		
		if err := tx.Create(&field).Error; err != nil {
			tx.Rollback()
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create form fields"})
			return
		}
	}
	
	// Commit the transaction
	if err := tx.Commit().Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Transaction failed"})
		return
	}
	
	// Load the form with its fields
	database.DB.Preload("Fields").First(&form, form.ID)
	
	c.JSON(http.StatusCreated, form)
}

// GetForms returns all forms for the authenticated user
func GetForms(c *gin.Context) {
	// Get user ID from context
	userID, exists := c.Get("user_id")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}
	
	var forms []models.Form
	
	// Retrieve forms for the user
	if err := database.DB.Where("user_id = ?", userID).Preload("Fields").Find(&forms).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve forms"})
		return
	}
	
	c.JSON(http.StatusOK, forms)
}

// GetFormByID returns a specific form by ID
func GetFormByID(c *gin.Context) {
	// Get user ID from context
	userID, exists := c.Get("user_id")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}
	
	// Get form ID from URL parameter
	formID, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid form ID"})
		return
	}
	
	var form models.Form
	
	// Find form by ID and make sure it belongs to the user
	if err := database.DB.Where("id = ? AND user_id = ?", formID, userID).Preload("Fields").First(&form).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.JSON(http.StatusNotFound, gin.H{"error": "Form not found"})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve form"})
		}
		return
	}
	
	c.JSON(http.StatusOK, form)
}

// UpdateForm updates an existing form
func UpdateForm(c *gin.Context) {
	// Get user ID from context
	userID, exists := c.Get("user_id")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}
	
	// Get form ID from URL parameter
	formID, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid form ID"})
		return
	}
	
	var form models.Form
	
	// Find form by ID and make sure it belongs to the user
	if err := database.DB.Where("id = ? AND user_id = ?", formID, userID).First(&form).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.JSON(http.StatusNotFound, gin.H{"error": "Form not found"})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve form"})
		}
		return
	}
	
	// Parse request body
	var request UpdateFormRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	
	// Begin a transaction
	tx := database.DB.Begin()
	
	// Update form attributes
	if request.Title != "" {
		form.Title = request.Title
	}
	form.Description = request.Description
	
	// Update original fields
	form.Field1 = request.Field1
	form.Field2 = request.Field2
	form.Field3 = request.Field3
	form.Field4 = request.Field4
	form.Field5 = request.Field5
	form.Field6 = request.Field6
	form.Field7 = request.Field7
	form.Field8 = request.Field8
	form.Field9 = request.Field9
	form.Field10 = request.Field10
	form.Field11 = request.Field11
	form.Field12 = request.Field12
	form.Field13 = request.Field13
	form.Field14 = request.Field14
	form.Field15 = request.Field15
	form.Field16 = request.Field16
	form.Field17 = request.Field17
	form.Field18 = request.Field18
	form.Field19 = request.Field19
	form.Field20 = request.Field20
	form.Field21 = request.Field21
	form.Field22 = request.Field22
	form.Field23 = request.Field23
	form.Field24 = request.Field24
	form.Field25 = request.Field25
	form.Field26 = request.Field26
	form.Field27 = request.Field27
	form.Field28 = request.Field28
	form.Field29 = request.Field29
	form.Field30 = request.Field30
	
	if err := tx.Save(&form).Error; err != nil {
		tx.Rollback()
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update form"})
		return
	}
	
	// Update fields if provided
	if request.Fields != nil {
		// Delete existing fields
		if err := tx.Where("form_id = ?", form.ID).Delete(&models.FormField{}).Error; err != nil {
			tx.Rollback()
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update form fields"})
			return
		}
		
		// Create new fields
		for _, fieldInput := range request.Fields {
			field := models.FormField{
				FormID:      form.ID,
				Label:       fieldInput.Label,
				Type:        fieldInput.Type,
				Required:    fieldInput.Required,
				Placeholder: fieldInput.Placeholder,
				Options:     fieldInput.Options,
			}
			
			if err := tx.Create(&field).Error; err != nil {
				tx.Rollback()
				c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create form fields"})
				return
			}
		}
	}
	
	// Commit the transaction
	if err := tx.Commit().Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Transaction failed"})
		return
	}
	
	// Load the updated form with its fields
	database.DB.Preload("Fields").First(&form, form.ID)
	
	c.JSON(http.StatusOK, form)
}

// DeleteForm deletes a form by ID
func DeleteForm(c *gin.Context) {
	// Get user ID from context
	userID, exists := c.Get("user_id")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}
	
	// Get form ID from URL parameter
	formID, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid form ID"})
		return
	}
	
	var form models.Form
	
	// Find form by ID and make sure it belongs to the user
	if err := database.DB.Where("id = ? AND user_id = ?", formID, userID).First(&form).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.JSON(http.StatusNotFound, gin.H{"error": "Form not found"})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve form"})
		}
		return
	}
	
	// Begin a transaction
	tx := database.DB.Begin()
	
	// Delete form fields first to maintain referential integrity
	if err := tx.Where("form_id = ?", formID).Delete(&models.FormField{}).Error; err != nil {
		tx.Rollback()
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete form fields"})
		return
	}
	
	// Delete the form
	if err := tx.Delete(&form).Error; err != nil {
		tx.Rollback()
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete form"})
		return
	}
	
	// Commit the transaction
	if err := tx.Commit().Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Transaction failed"})
		return
	}
	
	c.JSON(http.StatusOK, gin.H{"message": "Form deleted successfully"})
}

// SubmitFormResponse handles form submission
func SubmitFormResponse(c *gin.Context) {
	// Get form ID from URL parameter
	formID, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid form ID"})
		return
	}
	
	// Check if form exists
	var form models.Form
	if err := database.DB.Where("id = ?", formID).Preload("Fields").First(&form).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.JSON(http.StatusNotFound, gin.H{"error": "Form not found"})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve form"})
		}
		return
	}
	
	// Map of field IDs to field objects for validation
	fieldMap := make(map[uint]models.FormField)
	for _, field := range form.Fields {
		fieldMap[field.ID] = field
	}
	
	// Parse request body
	var answers map[string]string
	if err := c.ShouldBindJSON(&answers); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	
	// Begin a transaction
	tx := database.DB.Begin()
	
	// Create form response
	response := models.FormResponse{
		FormID: uint(formID),
	}
	
	if err := tx.Create(&response).Error; err != nil {
		tx.Rollback()
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create form response"})
		return
	}
	
	// Add answers to the response
	for fieldIDStr, value := range answers {
		fieldID, err := strconv.ParseUint(fieldIDStr, 10, 32)
		if err != nil {
			continue // Skip invalid field IDs
		}
		
		// Check if field exists in this form
		if _, exists := fieldMap[uint(fieldID)]; !exists {
			continue // Skip fields that don't belong to this form
		}
		
		answer := models.FormResponseAnswer{
			ResponseID: response.ID,
			FieldID:    uint(fieldID),
			Value:      value,
		}
		
		if err := tx.Create(&answer).Error; err != nil {
			tx.Rollback()
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to save form answers"})
			return
		}
	}
	
	// Commit the transaction
	if err := tx.Commit().Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Transaction failed"})
		return
	}
	
	c.JSON(http.StatusCreated, gin.H{
		"message":    "Form submitted successfully",
		"responseId": response.ID,
	})
}

// GetFormResponses returns all responses for a specific form
func GetFormResponses(c *gin.Context) {
	// Get user ID from context
	userID, exists := c.Get("user_id")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}
	
	// Get form ID from URL parameter
	formID, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid form ID"})
		return
	}
	
	// Check if form exists and belongs to the user
	var form models.Form
	if err := database.DB.Where("id = ? AND user_id = ?", formID, userID).First(&form).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.JSON(http.StatusNotFound, gin.H{"error": "Form not found"})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve form"})
		}
		return
	}
	
	// Retrieve all responses for this form
	var responses []models.FormResponse
	if err := database.DB.Where("form_id = ?", formID).Preload("Answers").Find(&responses).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve form responses"})
		return
	}
	
	c.JSON(http.StatusOK, responses)
}