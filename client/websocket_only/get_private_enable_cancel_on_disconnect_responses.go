// Code generated by go-swagger; DO NOT EDIT.

package websocket_only

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/canselcik/go-deribit/models"
)

// GetPrivateEnableCancelOnDisconnectReader is a Reader for the GetPrivateEnableCancelOnDisconnect structure.
type GetPrivateEnableCancelOnDisconnectReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetPrivateEnableCancelOnDisconnectReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetPrivateEnableCancelOnDisconnectOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetPrivateEnableCancelOnDisconnectBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetPrivateEnableCancelOnDisconnectOK creates a GetPrivateEnableCancelOnDisconnectOK with default headers values
func NewGetPrivateEnableCancelOnDisconnectOK() *GetPrivateEnableCancelOnDisconnectOK {
	return &GetPrivateEnableCancelOnDisconnectOK{}
}

/*GetPrivateEnableCancelOnDisconnectOK handles this case with default header values.

ok response
*/
type GetPrivateEnableCancelOnDisconnectOK struct {
	Payload *models.OkResponse
}

func (o *GetPrivateEnableCancelOnDisconnectOK) Error() string {
	return fmt.Sprintf("[GET /private/enable_cancel_on_disconnect][%d] getPrivateEnableCancelOnDisconnectOK  %+v", 200, o.Payload)
}

func (o *GetPrivateEnableCancelOnDisconnectOK) GetPayload() *models.OkResponse {
	return o.Payload
}

func (o *GetPrivateEnableCancelOnDisconnectOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.OkResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetPrivateEnableCancelOnDisconnectBadRequest creates a GetPrivateEnableCancelOnDisconnectBadRequest with default headers values
func NewGetPrivateEnableCancelOnDisconnectBadRequest() *GetPrivateEnableCancelOnDisconnectBadRequest {
	return &GetPrivateEnableCancelOnDisconnectBadRequest{}
}

/*GetPrivateEnableCancelOnDisconnectBadRequest handles this case with default header values.

result when used via rest/HTTP
*/
type GetPrivateEnableCancelOnDisconnectBadRequest struct {
	Payload *models.ErrorMessage
}

func (o *GetPrivateEnableCancelOnDisconnectBadRequest) Error() string {
	return fmt.Sprintf("[GET /private/enable_cancel_on_disconnect][%d] getPrivateEnableCancelOnDisconnectBadRequest  %+v", 400, o.Payload)
}

func (o *GetPrivateEnableCancelOnDisconnectBadRequest) GetPayload() *models.ErrorMessage {
	return o.Payload
}

func (o *GetPrivateEnableCancelOnDisconnectBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorMessage)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
