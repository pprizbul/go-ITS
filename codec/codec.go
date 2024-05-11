package codec

import (
	"encoding/asn1"
	"fmt"

	"github.com/pprizbul/go-ITS/protocols/CAM"
	"github.com/pprizbul/go-ITS/protocols/CDD"
	"github.com/pprizbul/go-ITS/protocols/CVM"
)

type Header struct {
	Header CDD.ItsPduHeader
}

type AsnMessage interface{}

func Encode(message AsnMessage) ([]byte, error) {
	var header CDD.ItsPduHeader
	header.ProtocolVersion = 1
	header.StationId = 0

	switch message := message.(type) {
	case CAM.CamPayload:
		camMessage := CAM.CAM{}
		header.MessageId = CDD.Cam
		camMessage.Header = header
		camMessage.Cam = message
		return asn1.Marshal(camMessage)
	case CVM.CvmPayload:
		cvmMessage := CVM.CVM{}
		header.MessageId = CVM.Cvm
		cvmMessage.Header = header
		cvmMessage.Cvm = message
		return asn1.Marshal(cvmMessage)

	default:
		return nil, fmt.Errorf("unsupported protocol")
	}
}

func Decode(data []byte) (AsnMessage, error) {
	var h Header
	_, err := asn1.Unmarshal(data, &h)
	if err != nil {
		return nil, fmt.Errorf("decoding header: %w", err)
	}

	switch h.Header.MessageId {
	case CDD.Cam:
		var cam CAM.CAM
		_, err := asn1.Unmarshal(data, &cam)
		if err != nil {
			return nil, fmt.Errorf("decoding CAM: %w", err)
		}
		return cam, nil
	case CVM.Cvm:
		var cvm CVM.CVM
		_, err := asn1.Unmarshal(data, &cvm)
		if err != nil {
			return nil, fmt.Errorf("decoding CVM: %w", err)
		}
		return cvm, nil
	default:
		return nil, fmt.Errorf("unsupported protocol")
	}

}
