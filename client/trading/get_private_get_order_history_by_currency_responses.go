// Code generated by go-swagger; DO NOT EDIT.

package trading

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/canselcik/go-deribit/models"
)

// GetPrivateGetOrderHistoryByCurrencyReader is a Reader for the GetPrivateGetOrderHistoryByCurrency structure.
type GetPrivateGetOrderHistoryByCurrencyReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetPrivateGetOrderHistoryByCurrencyReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetPrivateGetOrderHistoryByCurrencyOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetPrivateGetOrderHistoryByCurrencyOK creates a GetPrivateGetOrderHistoryByCurrencyOK with default headers values
func NewGetPrivateGetOrderHistoryByCurrencyOK() *GetPrivateGetOrderHistoryByCurrencyOK {
	return &GetPrivateGetOrderHistoryByCurrencyOK{}
}

/*GetPrivateGetOrderHistoryByCurrencyOK handles this case with default header values.

GetPrivateGetOrderHistoryByCurrencyOK get private get order history by currency o k
*/
type GetPrivateGetOrderHistoryByCurrencyOK struct {
	Payload *models.PrivateGetOrderHistoryResponse
}

func (o *GetPrivateGetOrderHistoryByCurrencyOK) Error() string {
	return fmt.Sprintf("[GET /private/get_order_history_by_currency][%d] getPrivateGetOrderHistoryByCurrencyOK  %+v", 200, o.Payload)
}

func (o *GetPrivateGetOrderHistoryByCurrencyOK) GetPayload() *models.PrivateGetOrderHistoryResponse {
	return o.Payload
}

func (o *GetPrivateGetOrderHistoryByCurrencyOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.PrivateGetOrderHistoryResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
