package vouchers

import (
	"github.com/gin-gonic/gin"
	voucherApi "github.com/h-varmazyar/voucher/api/proto"
	"github.com/h-varmazyar/voucher/pkg/grpcext"
	log "github.com/sirupsen/logrus"
	"net/http"
)

type Handler struct {
	voucherService voucherApi.VoucherServiceClient
	logger         *log.Logger
}

func NewHandler(configs *Configs, logger *log.Logger) *Handler {
	voucherConn := grpcext.NewConnection(configs.VoucherServiceAddress)
	handler := &Handler{
		logger:         logger,
		voucherService: voucherApi.NewVoucherServiceClient(voucherConn),
	}

	return handler
}

func (h *Handler) RegisterRoutes(router *gin.Engine) {
	voucherRoutes := router.Group("/vouchers")

	voucherRoutes.POST("/", h.createVoucher)
	voucherRoutes.GET("/:voucher_code/usages", h.voucherUsages)
	voucherRoutes.POST("/:voucher_code/apply", h.applyVoucher)

}

func (h *Handler) createVoucher(c *gin.Context) {
	req := new(voucherApi.VoucherCreateReq)
	if err := c.BindJSON(req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}

	h.logger.Infof("vt: %v", req.Type)
	if voucher, err := h.voucherService.Create(c, req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	} else {
		c.JSON(http.StatusOK, voucher)
	}
}

func (h *Handler) voucherUsages(c *gin.Context) {
	req := new(voucherApi.VoucherUsageReq)
	req.Identifier = &voucherApi.VoucherUsageReq_Code{Code: c.Param("voucher_code")}
	if usages, err := h.voucherService.Usage(c, req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	} else {
		c.JSON(http.StatusOK, usages)
	}
}

func (h *Handler) applyVoucher(c *gin.Context) {
	req := new(voucherApi.VoucherApplyReq)
	req.Code = c.Param("voucher_code")
	if err := c.BindJSON(req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}
	if _, err := h.voucherService.Apply(c, req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	} else {
		c.Status(http.StatusOK)
	}
}
