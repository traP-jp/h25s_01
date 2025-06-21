package handler

import (
	"backend/internal/domain/model"
	"backend/internal/domain/repository"
	"log"
	"mime/multipart"
	"net/http"
	"strconv"
	"time"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

const maxContentLength = 1024
const maxImages = 4

type ReviewHandler struct {
	reviewRepo repository.ReviewRepository
	fileRepo   repository.FileRepository
}

func NewReviewHandler(reviewRepo repository.ReviewRepository, fileRepo repository.FileRepository) *ReviewHandler {
	return &ReviewHandler{
		reviewRepo: reviewRepo,
		fileRepo:   fileRepo,
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
	if err != nil {
		return errorResponse(c, http.StatusUnauthorized, "Unauthorized: Failed to get user ID")
	}

	if err := validateAuthor(userID, req.Author); err != nil {
		return errorResponse(c, http.StatusForbidden, err.Error())
	}

	shopID, err := parseShopID(req.Shop)
	if err != nil {
		return errorResponse(c, http.StatusBadRequest, err.Error())
	}

	rating, err := parseRating(req.Rating)
	if err != nil {
		return errorResponse(c, http.StatusBadRequest, err.Error())
	}

	if err := validateContent(req.Content); err != nil {
		return errorResponse(c, http.StatusBadRequest, err.Error())
	}

	images, err := parseImages(req.Images)
	if err != nil {
		return errorResponse(c, http.StatusBadRequest, err.Error())
	}

	review, err := model.NewReview(
		model.UserID(userID),
		shopID,
		rating,
		req.Content,
		images,
	)
	if err != nil {
		return errorResponse(c, http.StatusInternalServerError, "Failed to create review")
	}

	err = h.reviewRepo.Save(c.Request().Context(), review)
	if err != nil {
		return errorResponse(c, http.StatusInternalServerError, "Failed to save review")
	}

	reviewDto := &Review{}
	reviewDto.FromModel(review)

	return c.JSON(http.StatusCreated, reviewDto)
}

func (h *ReviewHandler) GetReview(c echo.Context) error {
	id, err := uuid.Parse(c.Param("id"))
	if err != nil {
		return errorResponse(c, http.StatusBadRequest, "Invalid review ID")
	}

	review, err := h.reviewRepo.FindByID(c.Request().Context(), id)

	if err != nil {
		return errorResponse(c, http.StatusInternalServerError, "Failed to fetch review")
	}

	reviewDto := &Review{}
	reviewDto.FromModel(review)

	return c.JSON(http.StatusOK, reviewDto)
}

func (h *ReviewHandler) UpdateReview(c echo.Context) error {
	id, err := uuid.Parse(c.Param("id"))
	if err != nil {
		return errorResponse(c, http.StatusBadRequest, "Invalid review ID")
	}

	var req APIV1ReviewsPostRequest
	if err := c.Bind(&req); err != nil {
		return errorResponse(c, http.StatusBadRequest, "Invalid request payload")
	}

	userID, err := GetUserID(c)
	if err != nil {
		return errorResponse(c, http.StatusInternalServerError, "Failed to get user ID")
	}

	if err := validateAuthor(userID, req.Author); err != nil {
		return errorResponse(c, http.StatusForbidden, err.Error())
	}

	shopID, err := parseShopID(req.Shop)
	if err != nil {
		return errorResponse(c, http.StatusBadRequest, err.Error())
	}

	rating, err := parseRating(req.Rating)
	if err != nil {
		return errorResponse(c, http.StatusBadRequest, err.Error())
	}

	if err := validateContent(req.Content); err != nil {
		return errorResponse(c, http.StatusBadRequest, err.Error())
	}

	images, err := parseImages(req.Images)
	if err != nil {
		return errorResponse(c, http.StatusBadRequest, err.Error())
	}

	review, err := h.reviewRepo.FindByID(c.Request().Context(), id)
	if err != nil {
		return err
	}

	review.Author = model.UserID(userID)
	review.Shop = shopID
	review.Rating = rating
	review.Content = req.Content
	review.Images = images
	review.UpdatedAt = time.Now()

	err = h.reviewRepo.Save(c.Request().Context(), review)
	if err != nil {
		return errorResponse(c, http.StatusInternalServerError, "Failed to save review")
	}

	reviewDto := &Review{}
	reviewDto.FromModel(review)

	return c.JSON(http.StatusCreated, reviewDto)
}

func (h *ReviewHandler) DeleteReview(c echo.Context) error {
	id, err := uuid.Parse(c.Param("id"))
	if err != nil {
		return errorResponse(c, http.StatusBadRequest, "Invalid review ID")
	}

	userID, err := GetUserID(c)
	if err != nil {
		return errorResponse(c, http.StatusInternalServerError, "Failed to get user ID")
	}

	review, err := h.reviewRepo.FindByID(c.Request().Context(), id)
	if err != nil {
		return errorResponse(c, http.StatusNotFound, "Review not found")
	}

	if err := validateAuthor(userID, string(review.Author)); err != nil {
		return errorResponse(c, http.StatusForbidden, err.Error())
	}

	for _, image := range review.Images {
		_ = h.fileRepo.DeleteImage(c.Request().Context(), image.ID)
	}

	err = h.reviewRepo.Delete(c.Request().Context(), id)
	if err != nil {
		return errorResponse(c, http.StatusInternalServerError, "Failed to delete review")
	}

	return c.JSON(http.StatusOK, map[string]string{
		"message": "Review deleted successfully",
	})
}

func (h *ReviewHandler) UploadImage(c echo.Context) error {
	reviewID, err := uuid.Parse(c.Param("id"))
	if err != nil {
		return errorResponse(c, http.StatusBadRequest, "Invalid review ID")
	}

	review, err := h.reviewRepo.FindByID(c.Request().Context(), reviewID)
	if err != nil {
		return errorResponse(c, http.StatusNotFound, "Review not found")
	}

	fileHeader, err := c.FormFile("image")
	if err != nil {
		return errorResponse(c, http.StatusBadRequest, "Invalid image file")
	}

	contentType := fileHeader.Header.Get("Content-Type")

	file, err := fileHeader.Open()
	if err != nil {
		return errorResponse(c, http.StatusInternalServerError, "Failed to open image file")
	}
	defer func(file multipart.File) {
		err := file.Close()
		if err != nil {
			log.Println(err)
		}
	}(file)

	userID, err := GetUserID(c)
	if err != nil {
		return errorResponse(c, http.StatusInternalServerError, "Failed to get user ID")
	}

	if err := validateAuthor(userID, string(review.Author)); err != nil {
		return errorResponse(c, http.StatusForbidden, err.Error())
	}

	FileID, err := h.fileRepo.UploadImage(c.Request().Context(), contentType, file)
	if err != nil {
		return errorResponse(c, http.StatusInternalServerError, "Failed to save image file")
	}

	review.Images = append(review.Images, *model.NewImageFile(FileID))

	if err := h.reviewRepo.Save(c.Request().Context(), review); err != nil {
		return errorResponse(c, http.StatusInternalServerError, "Failed to save review with new image")
	}

	return c.JSON(http.StatusCreated, map[string]string{
		"id": FileID.String(),
	})
}

func validateAuthor(userID, author string) error {
	if userID != author {
		return echo.NewHTTPError(http.StatusForbidden, "You are not allowed to post this review")
	}

	return nil
}

func parseShopID(shop string) (uuid.UUID, error) {
	id, err := uuid.Parse(shop)
	if err != nil {
		return uuid.Nil, echo.NewHTTPError(http.StatusBadRequest, "Invalid shop ID")
	}

	return id, nil
}

func parseRating(rating int32) (model.Rating, error) {
	r, err := model.NewRating(int(rating))
	if err != nil {
		return 0, echo.NewHTTPError(http.StatusBadRequest, "Invalid rating value")
	}

	return r, nil
}

func validateContent(content string) error {
	if len(content) > maxContentLength {
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid content length")
	}

	return nil
}

func parseImages(imagesReq []string) ([]model.ImageFile, error) {
	if len(imagesReq) > maxImages {
		return nil, echo.NewHTTPError(http.StatusBadRequest, "Too many images")
	}

	images := make([]model.ImageFile, 0, len(imagesReq))

	for _, img := range imagesReq {
		imgID, err := uuid.Parse(img)
		if err != nil {
			return nil, echo.NewHTTPError(http.StatusBadRequest, "Invalid image ID")
		}

		images = append(images, *model.NewImageFile(imgID))
	}

	return images, nil
}
