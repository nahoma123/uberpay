package company

import (
	"context"
	"errors"
	"fmt"
	storage "ride_plus/internal/adapter/storage/persistence"
	custErr "ride_plus/internal/constant/errors"
	model "ride_plus/internal/constant/model/dbmodel"
	utils "ride_plus/internal/constant/model/init"

	"gorm.io/gorm"
)

type companyPersistence struct {
	conn *gorm.DB
}

func CompanyInit(utils utils.Utils) storage.CompanyPersistence {
	return &companyPersistence{
		conn: utils.Conn,
	}
}

func (r companyPersistence) CompanyByID(ctx context.Context, param model.Company) (*model.Company, error) {
	conn := r.conn.WithContext(ctx)
	company := &model.Company{}
	err := conn.Model(&model.Company{}).Where(&param).First(company).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, custErr.ErrRecordNotFound
		}
		return nil, custErr.ErrorUnableToFetch
	}
	return company, nil
}

func (r companyPersistence) Companies(ctx context.Context) ([]model.Company, error) {
	conn := r.conn.WithContext(ctx)
	companies := []model.Company{}
	err := conn.Model(&model.Company{}).Find(&companies).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, gorm.ErrRecordNotFound
		}
		return nil, custErr.ErrorUnableToFetch
	}
	return companies, err
}

func (r companyPersistence) StoreCompany(ctx context.Context, company model.Company) (*model.Company, error) {
	conn := r.conn.WithContext(ctx)
	err := conn.Model(&model.Company{}).Create(&company).Error
	if err != nil {
		if errors.Is(err, gorm.ErrInvalidTransaction) {
			return nil, gorm.ErrInvalidTransaction
		}
		return nil, custErr.ErrUnableToSave
	}
	return &company, nil
}
func (r companyPersistence) StoreCompanyImage(ctx context.Context, images model.CompanyImage) (*model.CompanyImage, error) {
	conn := r.conn.WithContext(ctx)
	normal_image_format := images.Image
	thumbnail_image_format := images.Formats.Thumbnail
	small_image_format := images.Formats.Small

	err := conn.Transaction(func(tx *gorm.DB) error {
		err := conn.Model(&model.Image{}).Create(normal_image_format).Error
		if err != nil {
			if errors.Is(err, gorm.ErrInvalidTransaction) {
				return gorm.ErrInvalidTransaction
			}
			return custErr.ErrUnableToSave
		}

		img := &model.Image{}
		err = tx.Model(&model.Image{}).Where(&normal_image_format).First(img).Error
		if err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				return custErr.ErrRecordNotFound
			}
			return custErr.ErrorUnableToFetch
		}
		thumbnail_image_format.ImageID = img.ID
		small_image_format.ImageID = img.ID

		err = tx.Model(&model.ImageFormat{}).Create(thumbnail_image_format).Error
		if err != nil {
			if errors.Is(err, gorm.ErrInvalidTransaction) {
				return gorm.ErrInvalidTransaction
			}
			return custErr.ErrUnableToSave
		}
		err = tx.Model(&model.ImageFormat{}).Create(small_image_format).Error
		if err != nil {
			if errors.Is(err, gorm.ErrInvalidTransaction) {
				return gorm.ErrInvalidTransaction
			}
			return custErr.ErrUnableToSave
		}
		return nil
	})

	if err != nil {
		return nil, custErr.ErrInvalidTransaction
	}
	return &images, nil
}
func (r companyPersistence) UpdateCompanyImage(ctx context.Context, images model.CompanyImage) (*model.CompanyImage, error) {
	conn := r.conn.WithContext(ctx)
	normal_image_format := images.Image
	thumbnail_image_format := images.Formats.Thumbnail
	small_image_format := images.Formats.Small
	thumbnail_image_format.ImageID = normal_image_format.ID
	small_image_format.ImageID = normal_image_format.ID
	err := conn.Transaction(func(tx *gorm.DB) error {
		err := conn.Model(&model.Image{}).Where(&model.Image{Hash: normal_image_format.Hash}).Updates(&normal_image_format).Error
		fmt.Println("error ", err)
		if err != nil {
			if errors.Is(err, gorm.ErrInvalidTransaction) {
				return gorm.ErrInvalidTransaction
			}
			return custErr.ErrUnableToSave
		}

		err = tx.Model(&model.ImageFormat{}).Where(&model.ImageFormat{Hash: thumbnail_image_format.Hash, FormatType: thumbnail_image_format.FormatType}).Updates(&thumbnail_image_format).Error
		fmt.Println("error ", err)
		if err != nil {
			if errors.Is(err, gorm.ErrInvalidTransaction) {
				return gorm.ErrInvalidTransaction
			}
			return custErr.ErrUnableToSave
		}
		err = tx.Model(&model.ImageFormat{}).Where(&model.ImageFormat{Hash: small_image_format.Hash, FormatType: small_image_format.FormatType}).Updates(&small_image_format).Error
		fmt.Println("error ", err)
		if err != nil {
			if errors.Is(err, gorm.ErrInvalidTransaction) {
				return gorm.ErrInvalidTransaction
			}
			return custErr.ErrUnableToSave
		}
		return nil
	})
	fmt.Println("error ", err)
	if err != nil {
		return nil, custErr.ErrInvalidTransaction
	}
	return &images, nil
}

func (r companyPersistence) CompanyImages(ctx context.Context) ([]model.CompanyImage, error) {
	conn := r.conn.WithContext(ctx)
	images := []model.CompanyImage{}
	img := []model.Image{}

	err := conn.Transaction(func(tx *gorm.DB) error {
		err := tx.Model(&model.Image{}).Find(&img).Error
		fmt.Println("error ", err)
		if err != nil {
			if errors.Is(err, gorm.ErrInvalidTransaction) {
				return gorm.ErrInvalidTransaction
			}
			return custErr.ErrUnableToSave
		}
		for _, image := range img {
			companyImage := model.CompanyImage{}
			thumbnail_image_format := &model.ImageFormat{}
			small_image_format := &model.ImageFormat{}

			err = tx.Model(&model.ImageFormat{}).Where(&model.ImageFormat{ImageID: image.ID, FormatType: "thumbnail"}).First(thumbnail_image_format).Error
			if err != nil {
				if errors.Is(err, gorm.ErrRecordNotFound) {
					return custErr.ErrRecordNotFound
				}
				return custErr.ErrorUnableToFetch
			}
			err = tx.Model(&model.ImageFormat{}).Where(&model.ImageFormat{ImageID: image.ID, FormatType: "small"}).First(small_image_format).Error
			if err != nil {
				if errors.Is(err, gorm.ErrRecordNotFound) {
					return custErr.ErrRecordNotFound
				}
				return custErr.ErrorUnableToFetch
			}
			companyImage.Image = &image
			companyImage.Formats.Thumbnail = thumbnail_image_format
			companyImage.Formats.Small = small_image_format
			images = append(images, companyImage)
		}
		return nil
	})
	if err != nil {
		return nil, custErr.ErrInvalidTransaction
	}
	return images, nil
}

func (r companyPersistence) UpdateCompany(ctx context.Context, company model.Company) (*model.Company, error) {
	conn := r.conn.WithContext(ctx)
	err := conn.Model(&model.Company{}).Where(&model.Company{ID: company.ID}).Updates(&company).Error
	if err != nil {
		if errors.Is(err, gorm.ErrInvalidTransaction) {
			return nil, gorm.ErrInvalidTransaction
		}
		return nil, custErr.ErrUnableToSave
	}
	return &company, nil
}

func (r companyPersistence) DeleteCompany(ctx context.Context, param model.Company) error {
	conn := r.conn.WithContext(ctx)
	err := conn.Model(&model.Company{}).Where(&param).Delete(&param).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return gorm.ErrRecordNotFound
		}
		return custErr.ErrorUnableToFetch
	}
	return nil
}
func (r companyPersistence) CompanyExists(ctx context.Context, param model.Company) (bool, error) {
	conn := r.conn.WithContext(ctx)
	var count int64
	err := conn.Model(&model.Company{}).Where(&param).Count(&count).Error
	if err != nil {
		return false, custErr.ErrRecordNotFound
	}
	return count > 0, nil
}
func (r companyPersistence) MigrateCompany(ctx context.Context) error {
	conn := r.conn.WithContext(ctx)
	err := conn.Migrator().AutoMigrate(&model.Company{})
	if err != nil {
		return err
	}

	err = conn.Migrator().AutoMigrate(&model.CompanyUser{})
	if err != nil {
		return err
	}
	return nil
}
func (r companyPersistence) ImageExists(param model.Image) (bool, error) {
	var count int64 = 0
	err := r.conn.Model(&model.Image{}).Where(&model.Image{Hash: param.Hash}).Count(&count).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return false, gorm.ErrRecordNotFound
		}
		return false, custErr.ErrorUnableToFetch
	}
	return count > 0, nil
}
