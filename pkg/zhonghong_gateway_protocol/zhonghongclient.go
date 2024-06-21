package zhonghonggatewayprotocol

// // ClientHandler is the interface that groups the Packager and Transporter methods.
// type ClientHandler interface {
// 	// Packager
// 	// Transporter
// }

// type client struct {
// 	// packager    Packager
// 	// transporter Transporter
// }

// NewClient creates a new Zhonghong client with given backend handler.
// func NewClient(handler ClientHandler) Client {
// 	return &client{packager: handler, transporter: handler}
// }

// func (mb *client) ReadGateway() (results []byte, err error) {
// 	request := ProtocolDataUnit{
// 		Header:       HeadCodeReadGateway,
// 		FunctionCode: FuncCodeReadGateway,
// 		Data:         []byte{0x00, 0x00, 0x00, 0x00},
// 	}
// resp, err := mb.send(&request)
// if err != nil {
// 	return nil, err
// }

// return resp.Data, nil
// }

// func (mb *client) send(request *ProtocolDataUnit) (response *ProtocolDataUnit, err error) {
// 	aduRequest, err := mb.packager.Encode(request)
// 	if err != nil {
// 		return
// 	}
// 	aduResponse, err := mb.transporter.Send(aduRequest)
// 	if err != nil {
// 		return
// 	}
// 	if err = mb.packager.Verify(aduRequest, aduResponse); err != nil {
// 		return
// 	}
// 	response, err = mb.packager.Decode(aduResponse)
// 	if err != nil {
// 		return
// 	}
// 	// Check correct function code returned (exception)
// 	if response.FunctionCode != request.FunctionCode {
// 		err = responseError(response)
// 		return
// 	}
// 	if response.Data == nil || len(response.Data) == 0 {
// 		// Empty response
// 		err = fmt.Errorf("Zhonghong: response data is empty")
// 		return
// 	}
// 	return
// }
