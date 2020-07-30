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

// GetPrivateGetMarginsReader is a Reader for the GetPrivateGetMargins structure.
type GetPrivateGetMarginsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetPrivateGetMarginsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetPrivateGetMarginsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetPrivateGetMarginsOK creates a GetPrivateGetMarginsOK with default headers values
func NewGetPrivateGetMarginsOK() *GetPrivateGetMarginsOK {
	return &GetPrivateGetMarginsOK{}
}

/*GetPrivateGetMarginsOK handles this case with default header values.

GetPrivateGetMarginsOK get private get margins o k
*/
type GetPrivateGetMarginsOK struct {
	Payload *models.PrivateGetMarginsResponse
}

func (o *GetPrivateGetMarginsOK) Error() string {
	return fmt.Sprintf("[GET /private/get_margins][%d] getPrivateGetMarginsOK  %+v", 200, o.Payload)
}

func (o *GetPrivateGetMarginsOK) GetPayload() *models.PrivateGetMarginsResponse {
	return o.Payload
}

func (o *GetPrivateGetMarginsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.PrivateGetMarginsResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
