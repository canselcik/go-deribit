// Code generated by go-swagger; DO NOT EDIT.

package wallet

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/canselcik/go-deribit/models"
)

// GetPrivateGetTransfersReader is a Reader for the GetPrivateGetTransfers structure.
type GetPrivateGetTransfersReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetPrivateGetTransfersReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetPrivateGetTransfersOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetPrivateGetTransfersOK creates a GetPrivateGetTransfersOK with default headers values
func NewGetPrivateGetTransfersOK() *GetPrivateGetTransfersOK {
	return &GetPrivateGetTransfersOK{}
}

/*GetPrivateGetTransfersOK handles this case with default header values.

GetPrivateGetTransfersOK get private get transfers o k
*/
type GetPrivateGetTransfersOK struct {
	Payload *models.PrivateGetTransfersResponse
}

func (o *GetPrivateGetTransfersOK) Error() string {
	return fmt.Sprintf("[GET /private/get_transfers][%d] getPrivateGetTransfersOK  %+v", 200, o.Payload)
}

func (o *GetPrivateGetTransfersOK) GetPayload() *models.PrivateGetTransfersResponse {
	return o.Payload
}

func (o *GetPrivateGetTransfersOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.PrivateGetTransfersResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
