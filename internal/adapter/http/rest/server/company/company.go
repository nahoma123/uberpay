package company

import (
	"fmt"
	"mime/multipart"
	"net/http"
	"os"
	"path/filepath"
	"ride_plus/internal/adapter/http/rest/server"
	"ride_plus/internal/adapter/http/rest/server/image"
	custErr "ride_plus/internal/constant/errors"
	model "ride_plus/internal/constant/model/dbmodel"
	"ride_plus/internal/constant/rest"
	"ride_plus/internal/module"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	uuid "github.com/satori/go.uuid"
)

// companyHandler defines all the things necessary for company handlers
type companyHandler struct {
	companyUsecase module.CompanyUsecase
	store          image.Storage
}

//CompanyInit initializes a company handler for the domain company
func CompanyInit(cmp module.CompanyUsecase, store image.Storage) server.CompanyHandler {
	return companyHandler{
		companyUsecase: cmp,
		store:          store,
	}
}

func (com companyHandler) StoreCompanyImage(ctx *gin.Context) {
	image := &model.Image{}
	thumbnail := &model.ImageFormat{}
	small := &model.ImageFormat{}
	err := ctx.Bind(image)
	//fmt.Println("form ",image)
	if err != nil {
		rest.ErrorResponseJson(ctx, custErr.ServiceError(custErr.ErrorUnableToBindJsonToStruct), http.StatusBadRequest)
		return
	}
	extension := filepath.Ext(image.ImageFile.Filename)
	hash, err := com.store.AverageHash(image.ImageFile)
	if err != nil {
		rest.ErrorResponseJson(ctx, err, http.StatusBadRequest)
		return
	}
	//puts types of hashing algorithm with image hash separating by _ (eg a: by a_ alg type=a)
	hash_rep := strings.Replace(hash, ":", "_", -1)
	imageId := hash_rep
	image.Name = imageId + extension
	image.Ext = extension
	//finds the original input form image dimension
	width, height := com.store.GetImageDimension(image.ImageFile)
	if width < 600 || height < 600 {
		nrr := custErr.ErrorModel{
			ErrorCode:        custErr.ErrCodes[custErr.ErrUnSupportedSize],
			ErrorMessage:     custErr.ErrUnSupportedSize.Error(),
			ErrorDescription: custErr.Descriptions[custErr.ErrUnSupportedSize],
		}
		rest.ErrorResponseJson(ctx, nrr, http.StatusBadRequest)
		return
	}

	m, err := com.store.Resize(image.ImageFile, 600, 600)
	rnimg, err := com.store.FileInfo(extension, m)
	if err != nil {
		rest.ErrorResponseJson(ctx, err, http.StatusBadRequest)
		return
	}
	imageInfo, err := rnimg.Stat()
	fmt.Println("line 97 error ", err)
	if err != nil {
		rest.ErrorResponseJson(ctx, err, http.StatusBadRequest)
		return
	}
	image.Size = imageInfo.Size()
	image.Width = 600
	image.Height = 600

	image.Hash = hash_rep
	image.Mime = "image/" + extension
	image.Url = "/uploads/" + image.Hash + extension
	image.PreviewUrl = "/uploads/" + image.Hash + extension
	image.CreatedAt = time.Now()
	image.UpdatedAt = time.Now() //this is might be changed later when updated put request comes from client
	//thumbnail image format
	thumbnail.Name = "thumbnail_" + imageId + extension
	thumbnail.Hash = "thumbnail_" + imageId
	thumbnail.Ext = extension
	thumbnail.Mime = image.Mime
	m, err = com.store.Resize(image.ImageFile, 156, 156)
	rtimg, err := com.store.FileInfo(extension, m)
	if err != nil {
		rest.ErrorResponseJson(ctx, err, http.StatusBadRequest)
		return
	}
	rtInfo, err := rtimg.Stat()
	fmt.Println("line 97 error ", err)
	if err != nil {
		rest.ErrorResponseJson(ctx, err, http.StatusBadRequest)
		return
	}
	thumbnail.Width = 156
	thumbnail.Height = 156
	thumbnail.Size = rtInfo.Size()
	thumbnail.FormatType = "thumbnail"
	thumbnail.Url = "/uploads/" + thumbnail.Hash + extension
	//small image format
	small.Name = "small_" + imageId + extension
	small.Hash = "small_" + imageId
	small.Ext = extension
	small.Mime = image.Mime
	m, err = com.store.Resize(image.ImageFile, 500, 500)
	rsimg, err := com.store.FileInfo(extension, m)
	if err != nil {
		rest.ErrorResponseJson(ctx, err, http.StatusBadRequest)
		return
	}
	rsInfo, err := rsimg.Stat()
	fmt.Println("line 117 error ", err)
	if err != nil {
		rest.ErrorResponseJson(ctx, err, http.StatusBadRequest)
		return
	}
	small.Width = 500
	small.Height = 500
	small.Size = rsInfo.Size()
	small.FormatType = "small"
	small.Url = "/uploads/" + small.Hash + extension
	parm := model.CompanyImage{}
	parm.Image = image
	parm.Formats.Thumbnail = thumbnail
	parm.Formats.Small = small
	companyImage, err := com.companyUsecase.StoreCompanyImage(ctx.Request.Context(), parm)
	fmt.Println("line 130 error ", err)
	if err != nil {
		nrr := custErr.ServiceError(err)
		rest.ErrorResponseJson(ctx, nrr, http.StatusBadRequest)
		return
	}
	//save normal image at server assets/images/normal directory
	err = com.SaveFile(image.ImageFile, "normal", image.Name, 600, 600)
	fmt.Println("error ", err)
	if err != nil {
		rest.ErrorResponseJson(ctx, err, http.StatusBadRequest)
		return
	}
	//save thumbnail image at server assets/images/thumbnail directory
	err = com.SaveFile(image.ImageFile, "thumbnail", thumbnail.Name, 156, 156)
	if err != nil {
		rest.ErrorResponseJson(ctx, err, http.StatusBadRequest)
		return
	}
	//save small image at server assets/images/small directory
	err = com.SaveFile(image.ImageFile, "small", small.Name, 500, 500)
	if err != nil {
		rest.ErrorResponseJson(ctx, err, http.StatusBadRequest)
		return
	}

	rest.ErrorResponseJson(ctx, companyImage, http.StatusOK)
}
func (com companyHandler) UpdateCompanyImage(ctx *gin.Context) {
	image := &model.Image{}
	thumbnail := &model.ImageFormat{}
	small := &model.ImageFormat{}
	err := ctx.Bind(image)
	//fmt.Println("form ",image)
	if err != nil {
		rest.ErrorResponseJson(ctx, custErr.ServiceError(custErr.ErrorUnableToBindJsonToStruct), http.StatusBadRequest)
		return
	}
	extension := filepath.Ext(image.ImageFile.Filename)
	hash, err := com.store.AverageHash(image.ImageFile)
	if err != nil {
		rest.ErrorResponseJson(ctx, err, http.StatusBadRequest)
		return
	}
	hash_rep := strings.Replace(hash, ":", "_", -1)
	imageId := hash_rep
	image.Name = imageId + extension
	image.Ext = extension
	width, height := com.store.GetImageDimension(image.ImageFile)
	if width < 600 || height < 600 {
		nrr := custErr.ErrorModel{
			ErrorCode:        custErr.ErrCodes[custErr.ErrUnSupportedSize],
			ErrorMessage:     custErr.ErrUnSupportedSize.Error(),
			ErrorDescription: custErr.Descriptions[custErr.ErrUnSupportedSize],
		}
		rest.ErrorResponseJson(ctx, nrr, http.StatusBadRequest)
		return
	}
	m, err := com.store.Resize(image.ImageFile, 600, 600)
	rnimg, err := com.store.FileInfo(extension, m)
	if err != nil {
		rest.ErrorResponseJson(ctx, err, http.StatusBadRequest)
		return
	}
	imageInfo, err := rnimg.Stat()
	fmt.Println("line 97 error ", err)
	if err != nil {
		rest.ErrorResponseJson(ctx, err, http.StatusBadRequest)
		return
	}
	image.Size = imageInfo.Size()
	image.Width = 600
	image.Height = 600

	image.Hash = hash_rep
	image.Mime = "image/" + extension
	image.Url = "/uploads/" + image.Hash + extension
	image.PreviewUrl = "/uploads/" + image.Hash + extension
	image.CreatedAt = time.Now()
	image.UpdatedAt = time.Now() //this is might be changed later when updated put request comes from client
	//thumbnail image format
	thumbnail.Name = "thumbnail_" + imageId + extension
	thumbnail.Hash = "thumbnail_" + image.Hash
	thumbnail.Ext = extension
	thumbnail.Mime = image.Mime
	m, err = com.store.Resize(image.ImageFile, 156, 156)
	rtimg, err := com.store.FileInfo(extension, m)
	if err != nil {
		rest.ErrorResponseJson(ctx, err, http.StatusBadRequest)
		return
	}
	rtInfo, err := rtimg.Stat()
	fmt.Println("line 97 error ", err)
	if err != nil {
		rest.ErrorResponseJson(ctx, err, http.StatusBadRequest)
		return
	}
	thumbnail.Width = 156
	thumbnail.Height = 156
	thumbnail.Size = rtInfo.Size()
	thumbnail.FormatType = "thumbnail"
	thumbnail.Url = "/uploads/" + thumbnail.Hash + extension
	//small image format
	small.Name = "small_" + imageId + extension
	small.Hash = "small_" + image.Hash
	small.Ext = extension
	small.Mime = image.Mime

	m, err = com.store.Resize(image.ImageFile, 300, 300)
	rsimg, err := com.store.FileInfo(extension, m)
	if err != nil {
		rest.ErrorResponseJson(ctx, err, http.StatusBadRequest)
		return
	}

	rsInfo, err := rsimg.Stat()
	fmt.Println("line 117 error ", err)
	fmt.Println("size ", rsInfo.Size())
	if err != nil {
		rest.ErrorResponseJson(ctx, err, http.StatusBadRequest)
		return
	}
	small.Width = 300
	small.Height = 300
	small.Size = rsInfo.Size()
	small.FormatType = "small"
	small.Url = "/uploads/" + small.Hash + extension
	parm := model.CompanyImage{}
	parm.Image = image
	parm.Formats.Thumbnail = thumbnail
	parm.Formats.Small = small
	companyImage, err := com.companyUsecase.UpdateCompanyImage(ctx.Request.Context(), parm)
	fmt.Println("line 130 error ", err)
	if err != nil {
		nrr := custErr.ServiceError(err)
		rest.ErrorResponseJson(ctx, nrr, http.StatusBadRequest)
		return
	}

	//save normal image at server assets/images/normal directory
	err = com.UpdateFile(image.ImageFile, "normal", image.Name, 600, 600)
	fmt.Println("line 291 error ", err)
	if err != nil {
		rest.ErrorResponseJson(ctx, err, http.StatusBadRequest)
		return
	}
	//save thumbnail image at server assets/images/thumbnail directory
	err = com.UpdateFile(image.ImageFile, "thumbnail", thumbnail.Name, 156, 156)
	if err != nil {
		rest.ErrorResponseJson(ctx, err, http.StatusBadRequest)
		return
	}
	//save small image at server assets/images/small directory
	err = com.UpdateFile(image.ImageFile, "small", small.Name, 300, 300)
	if err != nil {
		rest.ErrorResponseJson(ctx, err, http.StatusBadRequest)
		return
	}
	rest.ErrorResponseJson(ctx, companyImage, http.StatusOK)
}
func (com companyHandler) CompanyImages(c *gin.Context) {
	ctx := c.Request.Context()
	successData, err := com.companyUsecase.CompanyImages(ctx)
	if err != nil {
		nrr := custErr.ServiceError(err)
		rest.ErrorResponseJson(c, nrr, http.StatusBadRequest)
		return
	}
	rest.ErrorResponseJson(c, successData, http.StatusOK)
}
func (com companyHandler) CompanyByID(c *gin.Context) {
	id, err := uuid.FromString(c.Param("company-id"))
	if err != nil {
		rest.ErrorResponseJson(c, custErr.ConvertionError(), http.StatusBadRequest)
	}
	comp := model.Company{ID: id}
	ctx := c.Request.Context()
	successData, err := com.companyUsecase.CompanyByID(ctx, comp)
	fmt.Println("error ", err)
	if err != nil {
		nrr := custErr.ServiceError(err)
		rest.ErrorResponseJson(c, nrr, http.StatusBadRequest)
		return
	}
	rest.ErrorResponseJson(c, successData, http.StatusOK)
}
func (com companyHandler) Companies(c *gin.Context) {
	ctx := c.Request.Context()
	successData, err := com.companyUsecase.Companies(ctx)
	if err != nil {
		nrr := custErr.ServiceError(err)
		rest.ErrorResponseJson(c, nrr, http.StatusBadRequest)
		return
	}
	rest.ErrorResponseJson(c, successData, http.StatusOK)
}
func (com companyHandler) StoreCompany(c *gin.Context) {
	comp := &model.Company{}
	err := c.Bind(comp)
	if err != nil {
		rest.ErrorResponseJson(c, custErr.ServiceError(custErr.ErrorUnableToBindJsonToStruct), http.StatusBadRequest)
		return
	}
	ctx := c.Request.Context()
	successData, err := com.companyUsecase.StoreCompany(ctx, *comp)
	if err != nil {
		if strings.Contains(err.Error(), os.Getenv("ErrSecretKey")) {
			e := strings.Replace(err.Error(), os.Getenv("ErrSecretKey"), "", 1)
			errValue := custErr.ErrorModel{
				ErrorCode:        custErr.ErrCodes[custErr.ErrInvalidField],
				ErrorDescription: custErr.Descriptions[custErr.ErrInvalidField],
				ErrorMessage:     e,
			}
			rest.ErrorResponseJson(c, errValue, http.StatusBadRequest)
			return
		}
		rest.ErrorResponseJson(c, custErr.ServiceError(err), custErr.ErrCodes[err])
		return
	}
	rest.ErrorResponseJson(c, *successData, http.StatusOK)
}
func (com companyHandler) UpdateCompany(c *gin.Context) {
	id, err := uuid.FromString(c.Param("company-id"))
	if err != nil {
		rest.ErrorResponseJson(c, custErr.ConvertionError(), http.StatusBadRequest)
		return
	}
	comp := model.Company{ID: id}
	ctx := c.Request.Context()
	successData, err := com.companyUsecase.UpdateCompany(ctx, comp)
	if err != nil {
		if strings.Contains(err.Error(), os.Getenv("ErrSecretKey")) {
			e := strings.Replace(err.Error(), os.Getenv("ErrSecretKey"), "", 1)
			errValue := custErr.ErrorModel{
				ErrorCode:        custErr.ErrCodes[custErr.ErrInvalidField],
				ErrorDescription: custErr.Descriptions[custErr.ErrInvalidField],
				ErrorMessage:     e,
			}
			rest.ErrorResponseJson(c, errValue, http.StatusBadRequest)
			return
		}
		rest.ErrorResponseJson(c, custErr.ServiceError(err), custErr.ErrCodes[err])
		return
	}
	rest.ErrorResponseJson(c, *successData, http.StatusOK)
}
func (com companyHandler) DeleteCompany(c *gin.Context) {
	id, err := uuid.FromString(c.Param("company-id"))
	if err != nil {
		rest.ErrorResponseJson(c, custErr.ConvertionError(), http.StatusBadRequest)
		return
	}
	comp := model.Company{ID: id}
	ctx := c.Request.Context()
	err = com.companyUsecase.DeleteCompany(ctx, comp)

	if err != nil {
		nrr := custErr.ServiceError(err)
		rest.ErrorResponseJson(c, nrr, http.StatusBadRequest)
		return
	}
	rest.ErrorResponseJson(c, "company Deleted", http.StatusOK)
}
func (n companyHandler) SaveFile(f *multipart.FileHeader, format, path string, rwidth, rheiht uint) error {
	fp := filepath.Join("assets", "images", format, path)
	fmt.Println("path join ", fp)
	m, err := n.store.Resize(f, rwidth, rheiht)
	fmt.Println("line 411 error ", err)
	if err != nil {
		return err
	}
	_, err = n.store.Save(fp, m)
	fmt.Println("line 416 error ", err)
	if err != nil {
		return err
	}
	return nil
}
func (n companyHandler) UpdateFile(f *multipart.FileHeader, format, path string, rwidth, rheiht uint) error {
	fp := filepath.Join("assets", "images", format, path)
	fmt.Println("path join ", fp)
	m, err := n.store.Resize(f, rwidth, rheiht)
	fmt.Println("line 411 error ", err)
	if err != nil {
		return err
	}
	p := n.store.FullPath(fp)
	fmt.Println("fullpath ", p)
	err = os.Remove(p)
	fmt.Println("error remove ", err)
	if err != nil {
		return err
	}
	_, err = n.store.Save(fp, m)
	fmt.Println("line 416 error ", err)
	if err != nil {
		return err
	}
	return nil
}
