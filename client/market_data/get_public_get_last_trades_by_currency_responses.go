// Code generated by go-swagger; DO NOT EDIT.

package market_data

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/canselcik/go-deribit/models"
)

// GetPublicGetLastTradesByCurrencyReader is a Reader for the GetPublicGetLastTradesByCurrency structure.
type GetPublicGetLastTradesByCurrencyReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetPublicGetLastTradesByCurrencyReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetPublicGetLastTradesByCurrencyOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetPublicGetLastTradesByCurrencyOK creates a GetPublicGetLastTradesByCurrencyOK with default headers values
func NewGetPublicGetLastTradesByCurrencyOK() *GetPublicGetLastTradesByCurrencyOK {
	return &GetPublicGetLastTradesByCurrencyOK{}
}

/*GetPublicGetLastTradesByCurrencyOK handles this case with default header values.

GetPublicGetLastTradesByCurrencyOK get public get last trades by currency o k
*/
type GetPublicGetLastTradesByCurrencyOK struct {
	Payload *models.PublicTradesHistoryResponse
}

func (o *GetPublicGetLastTradesByCurrencyOK) Error() string {
	return fmt.Sprintf("[GET /public/get_last_trades_by_currency][%d] getPublicGetLastTradesByCurrencyOK  %+v", 200, o.Payload)
}

func (o *GetPublicGetLastTradesByCurrencyOK) GetPayload() *models.PublicTradesHistoryResponse {
	return o.Payload
}

func (o *GetPublicGetLastTradesByCurrencyOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.PublicTradesHistoryResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
