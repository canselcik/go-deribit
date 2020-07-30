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

// GetPrivateEditReader is a Reader for the GetPrivateEdit structure.
type GetPrivateEditReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetPrivateEditReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetPrivateEditOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetPrivateEditOK creates a GetPrivateEditOK with default headers values
func NewGetPrivateEditOK() *GetPrivateEditOK {
	return &GetPrivateEditOK{}
}

/*GetPrivateEditOK handles this case with default header values.

GetPrivateEditOK get private edit o k
*/
type GetPrivateEditOK struct {
	Payload *models.PrivateEditResponse
}

func (o *GetPrivateEditOK) Error() string {
	return fmt.Sprintf("[GET /private/edit][%d] getPrivateEditOK  %+v", 200, o.Payload)
}

func (o *GetPrivateEditOK) GetPayload() *models.PrivateEditResponse {
	return o.Payload
}

func (o *GetPrivateEditOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.PrivateEditResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
