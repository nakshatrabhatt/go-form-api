package models

import (
	"gorm.io/gorm"
)

// Form represents a form in the database
type Form struct {
	gorm.Model
	Title       string      `json:"title"`
	Description string      `json:"description"`
	UserID      uint        `json:"userId"`
	User        User        `json:"-" gorm:"foreignKey:UserID"` // Relationship with User
	Fields      []FormField `json:"fields" gorm:"foreignKey:FormID"`
	// Original fields from your forms.go
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

// FormField represents a single field in a form
type FormField struct {
	gorm.Model
	FormID      uint   `json:"formId" gorm:"index"`
	Label       string `json:"label"`
	Type        string `json:"type"` // text, number, email, checkbox, radio, etc.
	Required    bool   `json:"required"`
	Placeholder string `json:"placeholder"`
	Options     string `json:"options"` // JSON string for select/radio/checkbox options
}

// FormResponse is used when submitting responses to a form
type FormResponse struct {
	gorm.Model
	FormID  uint                 `json:"formId" gorm:"index"`
	Form    Form                 `json:"-" gorm:"foreignKey:FormID"`
	Answers []FormResponseAnswer `json:"answers" gorm:"foreignKey:ResponseID"`
}

// FormResponseAnswer stores individual answers for each field
type FormResponseAnswer struct {
	gorm.Model
	ResponseID uint   `json:"responseId" gorm:"index"`
	FieldID    uint   `json:"fieldId" gorm:"index"`
	Value      string `json:"value"`
}

// User represents a user in the database
type User struct {
	gorm.Model
	Username string `json:"username" gorm:"unique"`
	Email    string `json:"email" gorm:"unique"`
	Password string `json:"-"` // Password is not included in JSON responses
	Forms    []Form `json:"-" gorm:"foreignKey:UserID"`
}
