package handler

import (
	"net/http"

	"backend/internal/domain/repository"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

type FileHandler struct {
	fileRepository repository.FileRepository
}

func NewFileHandler(fileRepository repository.FileRepository) *FileHandler {
	return &FileHandler{
		fileRepository: fileRepository,
	}
}

func (h *FileHandler) GetImage(c echo.Context) error {
	idParam := c.Param("id")

	// UUIDの検証
	fileID, err := uuid.Parse(idParam)
	if err != nil {
		return errorResponse(c, http.StatusBadRequest, "Invalid image ID format")
	}

	// 画像を取得
	imageReader, contentType, err := h.fileRepository.GetImage(c.Request().Context(), fileID)
	if err != nil {
		return errorResponse(c, http.StatusNotFound, "Image not found")
	}
	defer imageReader.Close()

	// レスポンスヘッダーを設定
	c.Response().Header().Set("Content-Type", contentType)
	c.Response().Header().Set("Cache-Control", "public, max-age=31536000") // 1年間キャッシュ

	// 画像データをストリーミング
	return c.Stream(http.StatusOK, contentType, imageReader)
}
