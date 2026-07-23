package handler

import (
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"contact-book-api/internal/model"
	"contact-book-api/internal/service"
	"contact-book-api/pkg/response"
)

// ContactHandler holds the service it delegates to.
// The handler depends on the ContactServiceInterface, not the concrete type.
// This allows swapping implementations (e.g., mock for tests) without changing the handler.
type ContactHandler struct {
	contacts service.ContactServiceInterface
}

// NewContactHandler injects the service dependency via constructor.
// This is called in main.go when wiring layers together.
func NewContactHandler(contacts service.ContactServiceInterface) *ContactHandler {
	return &ContactHandler{contacts: contacts}
}

// CreateContactRequest is the JSON body shape for creating a contact.
// Gin's binding tag validates the input automatically.
type CreateContactRequest struct {
	FirstName string `json:"first_name" binding:"required"`
	LastName  string `json:"last_name" binding:"required"`
	Email     string `json:"email" binding:"required,email"`
	Phone     string `json:"phone" binding:"required"`
}

// UpdateContactRequest is the JSON body shape for updating a contact.
type UpdateContactRequest struct {
	FirstName string `json:"first_name" binding:"required"`
	LastName  string `json:"last_name" binding:"required"`
	Email     string `json:"email" binding:"required,email"`
	Phone     string `json:"phone" binding:"required"`
}

// ContactResponse is the JSON shape sent back to the client.
// It mirrors model.Contact but is a separate type so the API can evolve independently.
type ContactResponse struct {
	ID        int       `json:"id"`
	FirstName string    `json:"first_name"`
	LastName  string    `json:"last_name"`
	Email     string    `json:"email"`
	Phone     string    `json:"phone"`
	CreatedAt time.Time `json:"created_at"`
}

// toContactResponse converts a model.Contact to a ContactResponse DTO.
// We use a separate DTO (Data Transfer Object) instead of returning the model directly
// so the API can evolve independently of the database schema.
// For example, we might add "full_name" to the API response without changing the model.
func toContactResponse(c model.Contact) ContactResponse {
	return ContactResponse{
		ID:        c.ID,
		FirstName: c.FirstName,
		LastName:  c.LastName,
		Email:     c.Email,
		Phone:     c.Phone,
		CreatedAt: c.CreatedAt,
	}
}

// Create handles POST /contacts — receives JSON, validates via service, returns created contact.
func (h *ContactHandler) Create(c *gin.Context) {
	// c is the Gin context for this request.
	// c.ShouldBindJSON parses the request body into the struct.
	var req CreateContactRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.Fail(c.Writer, http.StatusBadRequest, err.Error())
		return
	}

	// c.Request.Context() carries the request's cancellation signal.
	// If the client disconnects mid-request, the context is cancelled.
	contact, err := h.contacts.Create(c.Request.Context(), req.FirstName, req.LastName, req.Email, req.Phone)
	if err != nil {
		if err == service.ErrEmailExists {
			response.Fail(c.Writer, http.StatusConflict, "email already exists")
			return
		}
		response.Fail(c.Writer, http.StatusInternalServerError, "internal error")
		return
	}

	response.OK(c.Writer, http.StatusCreated, toContactResponse(contact))
}

// GetAll handles GET /contacts — returns all contacts.
func (h *ContactHandler) GetAll(c *gin.Context) {
	contacts, err := h.contacts.GetAll(c.Request.Context())
	if err != nil {
		response.Fail(c.Writer, http.StatusInternalServerError, "internal error")
		return
	}

	// Convert each model.Contact to ContactResponse for the API.
	out := make([]ContactResponse, len(contacts))
	for i, c := range contacts {
		out[i] = toContactResponse(c)
	}

	response.OK(c.Writer, http.StatusOK, out)
}

// GetByID handles GET /contacts/:id — returns a single contact by ID.
func (h *ContactHandler) GetByID(c *gin.Context) {
	// c.Param("id") reads the :id segment from the URL path.
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		response.Fail(c.Writer, http.StatusBadRequest, "invalid contact id")
		return
	}

	contact, err := h.contacts.GetByID(c.Request.Context(), id)
	if err != nil {
		if err == service.ErrContactNotFound {
			response.Fail(c.Writer, http.StatusNotFound, "contact not found")
			return
		}
		response.Fail(c.Writer, http.StatusInternalServerError, "internal error")
		return
	}

	response.OK(c.Writer, http.StatusOK, toContactResponse(*contact))
}

// Update handles PUT /contacts/:id — updates an existing contact.
func (h *ContactHandler) Update(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		response.Fail(c.Writer, http.StatusBadRequest, "invalid contact id")
		return
	}

	var req UpdateContactRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.Fail(c.Writer, http.StatusBadRequest, err.Error())
		return
	}

	contact, err := h.contacts.Update(c.Request.Context(), id, req.FirstName, req.LastName, req.Email, req.Phone)
	if err != nil {
		if err == service.ErrContactNotFound {
			response.Fail(c.Writer, http.StatusNotFound, "contact not found")
			return
		}
		if err == service.ErrEmailExists {
			response.Fail(c.Writer, http.StatusConflict, "email already exists")
			return
		}
		response.Fail(c.Writer, http.StatusInternalServerError, "internal error")
		return
	}

	response.OK(c.Writer, http.StatusOK, toContactResponse(*contact))
}

// Delete handles DELETE /contacts/:id — removes a contact by ID.
func (h *ContactHandler) Delete(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		response.Fail(c.Writer, http.StatusBadRequest, "invalid contact id")
		return
	}

	err = h.contacts.Delete(c.Request.Context(), id)
	if err != nil {
		if err == service.ErrContactNotFound {
			response.Fail(c.Writer, http.StatusNotFound, "contact not found")
			return
		}
		response.Fail(c.Writer, http.StatusInternalServerError, "internal error")
		return
	}

	response.OK(c.Writer, http.StatusOK, gin.H{"message": "contact deleted"})
}