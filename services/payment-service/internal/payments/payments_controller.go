package payments

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"os"
	"payment-service/pkg/models"

	"github.com/gin-gonic/gin"
	"github.com/stripe/stripe-go/v74"
	"github.com/stripe/stripe-go/v74/webhook"
)

// PaymentsController handles HTTP requests
type PaymentsController struct {
	service *PaymentsService
}

// GetPaymentsController creates a new controller
func GetPaymentsController(service *PaymentsService) *PaymentsController {
	return &PaymentsController{
		service: service,
	}
}

// GetPaymentsHandler godoc
// @Summary      Get all payments
// @Description  Returns a list of all payments
// @Tags         payments
// @Produce      json
// @Success      200  {array}   Payment
// @Failure      500  {object}  models.ErrorResponse
// @Router       /payments [get]
func (c *PaymentsController) GetPaymentsHandler(ctx *gin.Context) {
	payments, err := c.service.GetPayments()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, models.ErrorResponse{Error: err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, payments)
}

// CreatePaymentHandler handles creating a new payment
// @Summary      Create a new payment
// @Description  Creates a Stripe PaymentIntent with the given amount, currency, and description
// @Tags         payments
// @Accept       json
// @Produce      json
// @Param        payment  body      CreatePaymentRequest  true  "Payment info"
// @Success      200      {object}  CreatePaymentResponse
// @Failure      400      {object}  models.ErrorResponse
// @Failure      500      {object}  models.ErrorResponse
// @Router       /payments [post]
func (c *PaymentsController) CreatePaymentHandler(ctx *gin.Context) {
	var req CreatePaymentRequest

	// Bind JSON request
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, models.ErrorResponse{Error: err.Error()})
		return
	}

	// Call service
	id, err := c.service.CreatePayment(req)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, models.ErrorResponse{Error: err.Error()})
		return
	}

	// Success response
	ctx.JSON(http.StatusOK, CreatePaymentResponse{PaymentIntentId: id})
}

// WebhookController služi kot ločen vhod za Stripe dogodke.
// @Router /stripe/webhook [post]
func (c *PaymentsController) HandleWebhook(ctx *gin.Context) {
	// 1. Preberite surov Body zahtevka
	const MaxBodyBytes = int64(65536)
	ctx.Request.Body = http.MaxBytesReader(ctx.Writer, ctx.Request.Body, MaxBodyBytes)
	payload, _ := io.ReadAll(ctx.Request.Body)
	// ... obravnava napak ...

	// 2. Preverite podpis Webhooka (za varnost!)
	// To zahteva vaš Stripe Webhook Secret (os.Getenv("STRIPE_WEBHOOK_SECRET"))
	event, _ := webhook.ConstructEvent(
		payload,
		ctx.GetHeader("Stripe-Signature"),
		os.Getenv("STRIPE_WEBHOOK_SECRET"),
	)
	// ... obravnava napak ...

	// 3. Obdelava dogodka
	if event.Type == "payment_intent.succeeded" {
		var pi stripe.PaymentIntent
		json.Unmarshal(event.Data.Raw, &pi)
		// ... obravnava napak ...

		// TUKAJ je mesto, kjer POSODOBITE svojo bazo:
		// - Posodobite plačilo v svoji DB na status 'SUCCESS'
		// - Generirajte račun (Invoice)
		// - Pošljite potrditveno e-pošto
		log.Printf("PaymentIntent succeeded: %s", pi.ID)
	}

	// ... obravnava drugih dogodkov (failed, cancelled itd.) ...

	ctx.Status(http.StatusOK) // Vedno vrnite 200 OK, sicer Stripe ponovi pošiljanje
}
