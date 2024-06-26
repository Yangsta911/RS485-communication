package clienthandler

import (
	"fmt"

	"github.com/Yangsta911/zhonghonghvac-go/pkg/client"
	"github.com/Yangsta911/zhonghonghvac-go/pkg/protocol"
)

const (
	AddressBytesLength = 2
)

type B27Packager struct {
}

func (p *B27Packager) Encode(pdu *protocol.ProtocolDataUnit) (adu []byte, err error) {
	length := 5 + len(pdu.Commands) + 1

	if length > rtuMaxSize {
		err = fmt.Errorf("b27-packager: length of data '%v' must not be bigger than '%v'", length, rtuMaxSize)
		return
	}
	adu = make([]byte, length)

	adu[0] = pdu.Header
	adu[1] = byte(length)
	copy(adu[2:4], pdu.Address)
	adu[4] = pdu.FunctionCode
	copy(adu[5:5+len(pdu.Commands)], pdu.Commands)
	adu[length-1] = client.CalculateByteSum(adu[0 : length-1])
	return
}

func (p *B27Packager) Decode(adu []byte) (pdu *protocol.ProtocolDataUnit, err error) {
	length := len(adu)
	receivedChecksum := uint8(adu[length-1])
	computedChecksum := client.CalculateByteSum(adu[0 : length-1])

	if computedChecksum != receivedChecksum {
		err = fmt.Errorf("b27-pack: response checksum '%v' does not match expected '%v'", receivedChecksum, computedChecksum)
		return
	}

	// Function code
	pdu = &protocol.ProtocolDataUnit{}
	pdu.Header = adu[0]
	copy(pdu.Address, adu[2:4])
	pdu.FunctionCode = adu[4]
	pdu.Data = adu[5 : length-1]
	return
}

func (p *B27Packager) Verify(aduRequest []byte, aduResponse []byte) error {
	if aduRequest[0] != 0xDD {
		return fmt.Errorf("b27-packager: protocol error")
	}

	if aduResponse[0] != 0xCC {
		return fmt.Errorf("b27-packager: protocol error")
	}

	// get last byte
	reqSum := aduRequest[len(aduRequest)-1]
	respSum := aduResponse[len(aduResponse)-1]

	if reqSum != respSum {
		return fmt.Errorf("b27-packager: checksum error")
	}

	return nil
}
