package documents

import (
	"database/sql"
	"encoding/json"
	"net/http"
	"strconv"
	"time"
	"upload-service/configs"
	"upload-service/pkg/common"

	"github.com/google/uuid"
)

type Controller struct {
	repo    *Repository
	storage *common.FileStorage
}

func DocumentController(db *sql.DB) *Controller {
	return &Controller{repo: DocumentRepository(db), storage: common.NewFileStorage(
		configs.GetStorageConfig(),
	)}
}

func (c *Controller) Index(w http.ResponseWriter, r *http.Request) {
	page, err := strconv.Atoi(r.URL.Query().Get("page"))
	if err != nil || page < 1 {
		page = 1
	}

	pageSize := 10
	offset := (page - 1) * pageSize

	total, err := c.repo.GetTotalDocuments()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(common.APIResponse{Error: "Failed to fetch documents"})
		return
	}

	docs, err := c.repo.GetDocumentsPaginated(pageSize, offset)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(common.APIResponse{Error: "Failed to fetch documents"})
		return
	}

	documents := []DocumentResource{}

	for _, doc := range docs {
		resource := DocumentResource{
			ID:         doc.ID,
			Name:       doc.Name,
			URL:        c.storage.GetFileURL(doc.Path),
			UploadedAt: doc.UploadedAt,
		}
		documents = append(documents, resource)
	}

	response := common.APIResponse{
		Data: common.PaginatedResponse{
			Results: documents,
			Total:   total,
		},
	}

	json.NewEncoder(w).Encode(response)
}

func (c *Controller) Store(w http.ResponseWriter, r *http.Request) {
	// Limit the size of the request body to prevent large file uploads
	r.Body = http.MaxBytesReader(w, r.Body, 10<<20)
	if err := r.ParseMultipartForm(10 << 20); err != nil {
		w.WriteHeader(http.StatusUnprocessableEntity)
		json.NewEncoder(w).Encode(common.APIResponse{Error: "No file provided or file too large"})
		return
	}

	file, handler, err := r.FormFile("file")
	if err != nil {
		w.WriteHeader(http.StatusUnprocessableEntity)
		json.NewEncoder(w).Encode(common.APIResponse{Error: "No file provided"})
		return
	}
	defer file.Close()

	docID := uuid.New().String()

	docPath, err := c.storage.StoreFile(handler, docID)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(common.APIResponse{Error: "Failed to store file"})
		return
	}

	doc := Document{
		ID:         docID,
		Name:       handler.Filename,
		Path:       docPath,
		UploadedAt: time.Now(),
	}

	err = c.repo.Store(doc)
	if err != nil {
		c.storage.DeleteFile(doc.Path)
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(common.APIResponse{Error: "Failed to save document record"})
		return
	}

	response := common.APIResponse{
		Data: DocumentResource{
			ID:         doc.ID,
			Name:       doc.Name,
			URL:        c.storage.GetFileURL(doc.Path),
			UploadedAt: doc.UploadedAt,
		},
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(response)
}

func (c *Controller) Delete(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	if id == "" {
		response := common.APIResponse{Error: "Document ID is required"}
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(response)
		return
	}

	doc, err := c.repo.GetByID(id)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(common.APIResponse{Error: "Document not found"})
		return
	}

	err = c.repo.Delete(id)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(common.APIResponse{Error: "Failed to delete document"})
		return
	}

	c.storage.DeleteFile(doc.Path)

	response := common.APIResponse{
		Data: DocumentResource{
			ID:         doc.ID,
			Name:       doc.Name,
			URL:        c.storage.GetFileURL(doc.Path),
			UploadedAt: doc.UploadedAt,
		},
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
}
