package repository

import (
	"go-api/entities"

	"gorm.io/gorm"
)

type DocumentRepository struct {
	connection *gorm.DB
}

func NewDocumentRepository(connection *gorm.DB) DocumentRepository {
	return DocumentRepository{
		connection: connection,
	}
}

func (doc *DocumentRepository) GetDocuments() ([]entities.Document, error){
	var document []entities.Document
	doc.connection.Find(&document)
	return document, nil
}

func (doc *DocumentRepository) CreateDocuments(document entities.Document) (entities.Document, error){
	if err := doc.connection.Create(&document).Error; err != nil {
		return entities.Document{}, err
	}
	return document, nil
}