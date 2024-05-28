package services

import (
	"submission_promotion_api/internal/app/models"
	"submission_promotion_api/internal/app/repositories"
)

// PromotionService provides promotion-related services
type PromotionService interface {
	CreatePromotion(promo models.Promotion) (models.Promotion, error)
	GetAllPromotions() ([]models.Promotion, error)
	GetPromotionbyPromotionID(promotionID string) (models.Promotion, error)
	UpdatePromotionbyPromotionID(promo models.Promotion) (models.Promotion, error)
	DeletePromotionbyPromotionID(promotionID string) error
}

type PromotionServiceImpl struct {
	PromotionRepo repositories.PromotionRepository
}

// NewPromotionService creates a new instance of PromotionService
func NewPromotionService(PromotionRepo repositories.PromotionRepository) *PromotionServiceImpl {
	return &PromotionServiceImpl{
		PromotionRepo: PromotionRepo,
	}
}

// CreatePromotion creates a new promotion
func (s *PromotionServiceImpl) CreatePromotion(promo models.Promotion) (models.Promotion, error) {
	// Implementasi kamu taruh disini
	return s.PromotionRepo.CreatePromotion(promo)
}

// GetAllPromotions that already recorded on database
func (s *PromotionServiceImpl) GetAllPromotions() ([]models.Promotion, error) {
	return s.PromotionRepo.GetAllPromotions()
}

// GetPromotionByPromotionID will throw data based on promotionID request
func (s *PromotionServiceImpl) GetPromotionbyPromotionID(promotionID string) (models.Promotion, error) {
	// Implementasi kamu taruh disini

	// Check promotionID exist
	return s.PromotionRepo.GetPromotionbyPromotionID(promotionID)

}

// UpdatePromotion will update data based on promotionID request
func (s *PromotionServiceImpl) UpdatePromotionbyPromotionID(promo models.Promotion) (models.Promotion, error) {
	// Implementasi kamu taruh disini

	// Check Duplicate promotion / promotion id exist
	return s.PromotionRepo.UpdatePromotionbyPromotionID(promo)
}

// DeletePromotionByPromotionID will delete data based on promotionID request
func (s *PromotionServiceImpl) DeletePromotionbyPromotionID(promotionID string) error {
	// Implementasi kamu taruh disini

	// del using params promotionID checking in repositories
	return s.PromotionRepo.DeletePromotionbyPromotionID(promotionID)
}
