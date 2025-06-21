package handler

import (
	"backend/internal/domain/model"
	"backend/internal/domain/repository"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"net/http"
	"strconv"
	"time"
)

const maxContentLength = 1024
const maxImages = 4

type ReviewHandler struct {
	reviewRepo repository.ReviewRepository
}

func NewReviewHandler(reviewRepo repository.ReviewRepository) *ReviewHandler {
	return &ReviewHandler{
		reviewRepo: reviewRepo,
	}
}

type Review struct {
	ID string `json:"id"`

	// レビュー投稿者のユーザーID
	Author string `json:"author"`

	// レビュー対象の店舗ID
	Shop string `json:"shop"`

	// 評価（0から3まで）
	Rating int32 `json:"rating,omitempty"`

	CreatedAt time.Time `json:"created_at,omitempty"`

	UpdatedAt time.Time `json:"updated_at,omitempty"`

	Content string `json:"content,omitempty"`

	Images []string `json:"images,omitempty"`
}

func (d *Review) FromModel(r *model.Review) {
	images := make([]string, len(r.Images))
	for i, img := range r.Images {
		images[i] = img.ID.String()
	}

	d.ID = r.ID.String()
	d.Author = string(r.Author)
	d.Shop = r.Shop.String()
	d.Rating = int32(r.Rating)
	d.CreatedAt = r.CreatedAt
	d.UpdatedAt = r.UpdatedAt
	d.Content = r.Content
	d.Images = images
}

type APIV1ReviewsPostRequest struct {
	Author string `json:"author"`

	Shop string `json:"shop"`

	Rating int32 `json:"rating"`

	Content string `json:"content,omitempty"`

	Images []string `json:"images,omitempty"`
}

func (h *ReviewHandler) GetReviews(c echo.Context) error {
	limit, err := strconv.Atoi(c.QueryParam("limit"))
	if err != nil || limit <= 0 {
		limit = 30
	}

	offset, err := strconv.Atoi(c.QueryParam("offset"))
	if err != nil || offset < 0 {
		offset = 0
	}

	after, err := time.Parse(time.RFC3339, c.QueryParam("after"))
	if err != nil {
		after = time.Time{}
	}

	before, err := time.Parse(time.RFC3339, c.QueryParam("before"))
	if err != nil {
		before = time.Now()
	}

	shopID, err := uuid.Parse(c.QueryParam("shop_id"))
	if err != nil {
		shopID = uuid.Nil
	}

	userID := c.QueryParam("user_id")

	reviews, err := h.reviewRepo.FindRecentReviews(
		c.Request().Context(),
		after,
		before,
		limit,
		offset,
		shopID,
		model.UserID(userID),
	)
	if err != nil {
		return errorResponse(c, http.StatusInternalServerError, "Failed to fetch reviews")
	}

	responses := make([]*Review, len(reviews))

	for i, review := range reviews {
		reviewDto := &Review{}
		reviewDto.FromModel(review)
		responses[i] = reviewDto
	}

	return c.JSON(http.StatusOK, responses)
}

func (h *ReviewHandler) CreateReview(c echo.Context) error {
	var req APIV1ReviewsPostRequest
	if err := c.Bind(&req); err != nil {
		return errorResponse(c, http.StatusBadRequest, "Invalid request payload")

	}

	userID, err := GetUserID(c)
	if err != nil || userID != req.Author {
		return errorResponse(c, http.StatusForbidden, "You are not allowed to post this review")
	}

	shopID, err := uuid.Parse(req.Shop)
	if err != nil {
		return errorResponse(c, http.StatusBadRequest, "Invalid shop ID")
	}

	rating, err := model.NewRating(int(req.Rating))
	if err != nil {
		return errorResponse(c, http.StatusBadRequest, "Invalid rating value")
	}

	content := req.Content
	if len(content) > maxContentLength {
		return errorResponse(c, http.StatusBadRequest, "Invalid content length")
	}

	if len(req.Images) > maxImages {
		return errorResponse(c, http.StatusBadRequest, "Too many images")
	}

	images := make([]model.ImageFile, 0, len(req.Images))

	for i, img := range req.Images {
		imgID, err := uuid.Parse(img)
		if err != nil {
			return errorResponse(c, http.StatusBadRequest, "Invalid image ID")
		}

		images[i] = *model.NewImageFile(imgID)
	}

	review, err := model.NewReview(
		model.UserID(userID),
		shopID,
		rating,
		content,
		images,
	)
	if err != nil {
		return errorResponse(c, http.StatusInternalServerError, "Failed to create review")
	}

	err = h.reviewRepo.Save(c.Request().Context(), review)

	reviewDto := &Review{}
	reviewDto.FromModel(review)

	return c.JSON(http.StatusCreated, reviewDto)
}
