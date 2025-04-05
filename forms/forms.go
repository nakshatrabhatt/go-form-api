package forms

import (
	"encoding/json"
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/mux"

	"github.com/nakshatrabhatt/go-form-api/database"
)

// Form represents a form with all its fields
type Form struct {
	ID          int    `json:"id"`
	UserID      int    `json:"user_id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Field1      string `json:"field1"`
	Field2      string `json:"field2"`
	Field3      string `json:"field3"`
	Field4      string `json:"field4"`
	Field5      string `json:"field5"`
	Field6      string `json:"field6"`
	Field7      string `json:"field7"`
	Field8      string `json:"field8"`
	Field9      string `json:"field9"`
	Field10     string `json:"field10"`
	Field11     string `json:"field11"`
	Field12     string `json:"field12"`
	Field13     string `json:"field13"`
	Field14     string `json:"field14"`
	Field15     string `json:"field15"`
	Field16     string `json:"field16"`
	Field17     string `json:"field17"`
	Field18     string `json:"field18"`
	Field19     string `json:"field19"`
	Field20     string `json:"field20"`
	Field21     string `json:"field21"`
	Field22     string `json:"field22"`
	Field23     string `json:"field23"`
	Field24     string `json:"field24"`
	Field25     string `json:"field25"`
	Field26     string `json:"field26"`
	Field27     string `json:"field27"`
	Field28     string `json:"field28"`
	Field29     string `json:"field29"`
	Field30     string `json:"field30"`
	CreatedAt   string `json:"created_at"`
}

// FormField represents a structured field in a form
type FormField struct {
	ID          int    `json:"id"`
	FormID      int    `json:"form_id"`
	Label       string `json:"label"`
	Type        string `json:"type"`
	Required    bool   `json:"required"`
	Placeholder string `json:"placeholder"`
	Options     string `json:"options"`
}

// CreateFormRequest represents the request structure for creating a form
type CreateFormRequest struct {
	Title       string      `json:"title"`
	Description string      `json:"description"`
	Fields      []FormField `json:"fields"`
	// Original fields
	Field1  string `json:"field1"`
	Field2  string `json:"field2"`
	Field3  string `json:"field3"`
	Field4  string `json:"field4"`
	Field5  string `json:"field5"`
	Field6  string `json:"field6"`
	Field7  string `json:"field7"`
	Field8  string `json:"field8"`
	Field9  string `json:"field9"`
	Field10 string `json:"field10"`
	Field11 string `json:"field11"`
	Field12 string `json:"field12"`
	Field13 string `json:"field13"`
	Field14 string `json:"field14"`
	Field15 string `json:"field15"`
	Field16 string `json:"field16"`
	Field17 string `json:"field17"`
	Field18 string `json:"field18"`
	Field19 string `json:"field19"`
	Field20 string `json:"field20"`
	Field21 string `json:"field21"`
	Field22 string `json:"field22"`
	Field23 string `json:"field23"`
	Field24 string `json:"field24"`
	Field25 string `json:"field25"`
	Field26 string `json:"field26"`
	Field27 string `json:"field27"`
	Field28 string `json:"field28"`
	Field29 string `json:"field29"`
	Field30 string `json:"field30"`
}

// CreateForm handles new form submissions
func CreateForm(w http.ResponseWriter, r *http.Request) {
	// Set Content-Type header
	w.Header().Set("Content-Type", "application/json")

	// Parse the request body
	var request CreateFormRequest
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		http.Error(w, "Error decoding request", http.StatusBadRequest)
		return
	}

	// Extract user ID from context (set by auth middleware)
	userID := r.Context().Value("user_id").(int)

	// Create form instance
	form := Form{
		UserID:      userID,
		Title:       request.Title,
		Description: request.Description,
		Field1:      request.Field1,
		Field2:      request.Field2,
		Field3:      request.Field3,
		Field4:      request.Field4,
		Field5:      request.Field5,
		Field6:      request.Field6,
		Field7:      request.Field7,
		Field8:      request.Field8,
		Field9:      request.Field9,
		Field10:     request.Field10,
		Field11:     request.Field11,
		Field12:     request.Field12,
		Field13:     request.Field13,
		Field14:     request.Field14,
		Field15:     request.Field15,
		Field16:     request.Field16,
		Field17:     request.Field17,
		Field18:     request.Field18,
		Field19:     request.Field19,
		Field20:     request.Field20,
		Field21:     request.Field21,
		Field22:     request.Field22,
		Field23:     request.Field23,
		Field24:     request.Field24,
		Field25:     request.Field25,
		Field26:     request.Field26,
		Field27:     request.Field27,
		Field28:     request.Field28,
		Field29:     request.Field29,
		Field30:     request.Field30,
		CreatedAt:   time.Now().Format(time.RFC3339),
	}

	// Begin a database transaction
	tx := database.DB.Begin()
	if tx.Error != nil {
		http.Error(w, "Failed to begin transaction", http.StatusInternalServerError)
		return
	}

	// Insert the form into the database
	result := tx.Exec(`
		INSERT INTO forms (
			user_id, title, description, 
			field1, field2, field3, field4, field5, field6, field7, field8, field9, field10,
			field11, field12, field13, field14, field15, field16, field17, field18, field19, field20,
			field21, field22, field23, field24, field25, field26, field27, field28, field29, field30,
			created_at
		) VALUES (
			?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?,
			?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?
		)`,
		form.UserID, form.Title, form.Description,
		form.Field1, form.Field2, form.Field3, form.Field4, form.Field5,
		form.Field6, form.Field7, form.Field8, form.Field9, form.Field10,
		form.Field11, form.Field12, form.Field13, form.Field14, form.Field15,
		form.Field16, form.Field17, form.Field18, form.Field19, form.Field20,
		form.Field21, form.Field22, form.Field23, form.Field24, form.Field25,
		form.Field26, form.Field27, form.Field28, form.Field29, form.Field30,
		form.CreatedAt,
	)

	if result.Error != nil {
		tx.Rollback()
		http.Error(w, "Failed to create form", http.StatusInternalServerError)
		return
	}

	// Get the ID of the inserted form
	formID := result.RowsAffected
	if formID == 0 {
		tx.Rollback()
		http.Error(w, "Failed to get form ID", http.StatusInternalServerError)
		return
	}

	// For GORM, we need to get the last inserted ID through a separate query if needed
	var lastInsertID int
	if err := tx.Raw("SELECT LAST_INSERT_ID()").Scan(&lastInsertID).Error; err != nil {
		tx.Rollback()
		http.Error(w, "Failed to get form ID", http.StatusInternalServerError)
		return
	}

	form.ID = lastInsertID

	// Insert form fields if provided
	if len(request.Fields) > 0 {
		for _, field := range request.Fields {
			result := tx.Exec(`
				INSERT INTO form_fields (form_id, label, type, required, placeholder, options)
				VALUES (?, ?, ?, ?, ?, ?)
			`, lastInsertID, field.Label, field.Type, field.Required, field.Placeholder, field.Options)

			if result.Error != nil {
				tx.Rollback()
				http.Error(w, "Failed to create form fields", http.StatusInternalServerError)
				return
			}
		}
	}

	// Commit the transaction
	if err := tx.Commit().Error; err != nil {
		http.Error(w, "Transaction failed", http.StatusInternalServerError)
		return
	}

	// Load the form fields to include in the response
	var fields []FormField
	rows, err := database.DB.Raw("SELECT id, form_id, label, type, required, placeholder, options FROM form_fields WHERE form_id = ?", lastInsertID).Rows()
	if err != nil {
		http.Error(w, "Failed to retrieve form fields", http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	for rows.Next() {
		var field FormField
		if err := database.DB.ScanRows(rows, &field); err != nil {
			http.Error(w, "Failed to scan form fields", http.StatusInternalServerError)
			return
		}
		fields = append(fields, field)
	}

	// Create response with both form and fields
	response := struct {
		Form   Form        `json:"form"`
		Fields []FormField `json:"fields"`
	}{
		Form:   form,
		Fields: fields,
	}

	// Return the created form with its fields
	json.NewEncoder(w).Encode(response)
}

// GetForms returns all forms for the authenticated user
func GetForms(w http.ResponseWriter, r *http.Request) {
	// Set Content-Type header
	w.Header().Set("Content-Type", "application/json")

	// Extract user ID from context (set by auth middleware)
	userID := r.Context().Value("user_id").(int)

	// Query all forms owned by this user
	var forms []Form
	rows, err := database.DB.Raw(`
		SELECT id, user_id, title, description,
		field1, field2, field3, field4, field5, field6, field7, field8, field9, field10,
		field11, field12, field13, field14, field15, field16, field17, field18, field19, field20,
		field21, field22, field23, field24, field25, field26, field27, field28, field29, field30,
		created_at
		FROM forms WHERE user_id = ? ORDER BY created_at DESC
	`, userID).Rows()

	if err != nil {
		http.Error(w, "Failed to retrieve forms", http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	// Parse query results into form objects
	for rows.Next() {
		var form Form
		if err := database.DB.ScanRows(rows, &form); err != nil {
			http.Error(w, "Error scanning form data", http.StatusInternalServerError)
			return
		}
		forms = append(forms, form)
	}

	// Return the forms as JSON
	json.NewEncoder(w).Encode(forms)
}

// GetFormByID returns a specific form and its fields
func GetFormByID(w http.ResponseWriter, r *http.Request) {
	// Set Content-Type header
	w.Header().Set("Content-Type", "application/json")

	// Extract form ID from URL
	vars := mux.Vars(r)
	formID, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, "Invalid form ID", http.StatusBadRequest)
		return
	}

	// Extract user ID from context (set by auth middleware)
	userID := r.Context().Value("user_id").(int)

	// Query the form and verify ownership
	var form Form
	result := database.DB.Raw(`
		SELECT id, user_id, title, description,
		field1, field2, field3, field4, field5, field6, field7, field8, field9, field10,
		field11, field12, field13, field14, field15, field16, field17, field18, field19, field20,
		field21, field22, field23, field24, field25, field26, field27, field28, field29, field30,
		created_at
		FROM forms WHERE id = ?
	`, formID).Scan(&form)

	if result.Error != nil {
		http.Error(w, "Form not found", http.StatusNotFound)
		return
	}

	// Verify form belongs to the authenticated user
	if form.UserID != userID {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	// Query form fields
	var fields []FormField
	rows, err := database.DB.Raw("SELECT id, form_id, label, type, required, placeholder, options FROM form_fields WHERE form_id = ?", formID).Rows()
	if err != nil {
		http.Error(w, "Failed to retrieve form fields", http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	for rows.Next() {
		var field FormField
		if err := database.DB.ScanRows(rows, &field); err != nil {
			http.Error(w, "Failed to scan form fields", http.StatusInternalServerError)
			return
		}
		fields = append(fields, field)
	}

	// Create response with both form and fields
	response := struct {
		Form   Form        `json:"form"`
		Fields []FormField `json:"fields"`
	}{
		Form:   form,
		Fields: fields,
	}

	// Return the form with its fields
	json.NewEncoder(w).Encode(response)
}

// UpdateForm handles form updates
func UpdateForm(w http.ResponseWriter, r *http.Request) {
	// Set Content-Type header
	w.Header().Set("Content-Type", "application/json")

	// Extract form ID from URL
	vars := mux.Vars(r)
	formID, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, "Invalid form ID", http.StatusBadRequest)
		return
	}

	// Extract user ID from context (set by auth middleware)
	userID := r.Context().Value("user_id").(int)

	// Verify form ownership
	var ownerID int
	result := database.DB.Raw("SELECT user_id FROM forms WHERE id = ?", formID).Scan(&ownerID)
	if result.Error != nil {
		http.Error(w, "Form not found", http.StatusNotFound)
		return
	}

	if ownerID != userID {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	// Parse the request body
	var request CreateFormRequest
	err = json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		http.Error(w, "Error decoding request", http.StatusBadRequest)
		return
	}

	// Begin a database transaction
	tx := database.DB.Begin()
	if tx.Error != nil {
		http.Error(w, "Failed to begin transaction", http.StatusInternalServerError)
		return
	}

	// Update the form
	result = tx.Exec(`
		UPDATE forms SET
		title = ?, description = ?,
		field1 = ?, field2 = ?, field3 = ?, field4 = ?, field5 = ?,
		field6 = ?, field7 = ?, field8 = ?, field9 = ?, field10 = ?,
		field11 = ?, field12 = ?, field13 = ?, field14 = ?, field15 = ?,
		field16 = ?, field17 = ?, field18 = ?, field19 = ?, field20 = ?,
		field21 = ?, field22 = ?, field23 = ?, field24 = ?, field25 = ?,
		field26 = ?, field27 = ?, field28 = ?, field29 = ?, field30 = ?
		WHERE id = ? AND user_id = ?
	`,
		request.Title, request.Description,
		request.Field1, request.Field2, request.Field3, request.Field4, request.Field5,
		request.Field6, request.Field7, request.Field8, request.Field9, request.Field10,
		request.Field11, request.Field12, request.Field13, request.Field14, request.Field15,
		request.Field16, request.Field17, request.Field18, request.Field19, request.Field20,
		request.Field21, request.Field22, request.Field23, request.Field24, request.Field25,
		request.Field26, request.Field27, request.Field28, request.Field29, request.Field30,
		formID, userID,
	)

	if result.Error != nil {
		tx.Rollback()
		http.Error(w, "Failed to update form", http.StatusInternalServerError)
		return
	}

	// Delete existing fields
	result = tx.Exec("DELETE FROM form_fields WHERE form_id = ?", formID)
	if result.Error != nil {
		tx.Rollback()
		http.Error(w, "Failed to update form fields", http.StatusInternalServerError)
		return
	}

	// Insert new fields
	if len(request.Fields) > 0 {
		for _, field := range request.Fields {
			result := tx.Exec(`
				INSERT INTO form_fields (form_id, label, type, required, placeholder, options)
				VALUES (?, ?, ?, ?, ?, ?)
			`, formID, field.Label, field.Type, field.Required, field.Placeholder, field.Options)

			if result.Error != nil {
				tx.Rollback()
				http.Error(w, "Failed to create form fields", http.StatusInternalServerError)
				return
			}
		}
	}

	// Commit the transaction
	if err := tx.Commit().Error; err != nil {
		http.Error(w, "Transaction failed", http.StatusInternalServerError)
		return
	}

	// Query the updated form
	var form Form
	result = database.DB.Raw(`
		SELECT id, user_id, title, description,
		field1, field2, field3, field4, field5, field6, field7, field8, field9, field10,
		field11, field12, field13, field14, field15, field16, field17, field18, field19, field20,
		field21, field22, field23, field24, field25, field26, field27, field28, field29, field30,
		created_at
		FROM forms WHERE id = ?
	`, formID).Scan(&form)

	if result.Error != nil {
		http.Error(w, "Failed to retrieve updated form", http.StatusInternalServerError)
		return
	}

	// Query updated form fields
	var fields []FormField
	rows, err := database.DB.Raw("SELECT id, form_id, label, type, required, placeholder, options FROM form_fields WHERE form_id = ?", formID).Rows()
	if err != nil {
		http.Error(w, "Failed to retrieve form fields", http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	for rows.Next() {
		var field FormField
		if err := database.DB.ScanRows(rows, &field); err != nil {
			http.Error(w, "Failed to scan form fields", http.StatusInternalServerError)
			return
		}
		fields = append(fields, field)
	}

	// Create response with both form and fields
	response := struct {
		Form   Form        `json:"form"`
		Fields []FormField `json:"fields"`
	}{
		Form:   form,
		Fields: fields,
	}

	// Return the updated form with its fields
	json.NewEncoder(w).Encode(response)
}

// DeleteForm removes a form and its fields from the database
func DeleteForm(w http.ResponseWriter, r *http.Request) {
	// Set Content-Type header
	w.Header().Set("Content-Type", "application/json")

	// Extract form ID from URL
	vars := mux.Vars(r)
	formID, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, "Invalid form ID", http.StatusBadRequest)
		return
	}

	// Extract user ID from context (set by auth middleware)
	userID := r.Context().Value("user_id").(int)

	// Verify form ownership
	var ownerID int
	result := database.DB.Raw("SELECT user_id FROM forms WHERE id = ?", formID).Scan(&ownerID)
	if result.Error != nil {
		http.Error(w, "Form not found", http.StatusNotFound)
		return
	}

	if ownerID != userID {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	// Begin a database transaction
	tx := database.DB.Begin()
	if tx.Error != nil {
		http.Error(w, "Failed to begin transaction", http.StatusInternalServerError)
		return
	}

	// Delete form fields first to maintain referential integrity
	result = tx.Exec("DELETE FROM form_fields WHERE form_id = ?", formID)
	if result.Error != nil {
		tx.Rollback()
		http.Error(w, "Failed to delete form fields", http.StatusInternalServerError)
		return
	}

	// Delete form responses and answers if they exist
	result = tx.Exec("DELETE FROM form_response_answers WHERE response_id IN (SELECT id FROM form_responses WHERE form_id = ?)", formID)
	if result.Error != nil {
		tx.Rollback()
		http.Error(w, "Failed to delete form responses", http.StatusInternalServerError)
		return
	}

	result = tx.Exec("DELETE FROM form_responses WHERE form_id = ?", formID)
	if result.Error != nil {
		tx.Rollback()
		http.Error(w, "Failed to delete form responses", http.StatusInternalServerError)
		return
	}

	// Finally delete the form
	result = tx.Exec("DELETE FROM forms WHERE id = ? AND user_id = ?", formID, userID)
	if result.Error != nil {
		tx.Rollback()
		http.Error(w, "Failed to delete form", http.StatusInternalServerError)
		return
	}

	// Commit the transaction
	if err := tx.Commit().Error; err != nil {
		http.Error(w, "Transaction failed", http.StatusInternalServerError)
		return
	}

	// Return success message
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"message": "Form deleted successfully"})
}
