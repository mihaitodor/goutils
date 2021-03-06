// Autogenerated by Thrift Compiler (1.0.0-upfluence)
// DO NOT EDIT UNLESS YOU ARE SURE THAT YOU KNOW WHAT YOU ARE DOING

package monitoring

import (
	"bytes"
	"fmt"
	"github.com/upfluence/goutils/Godeps/_workspace/src/github.com/upfluence/thrift/lib/go/thrift"
	"time"
)

// (needed to ensure safety because of naive import list construction.)
var _ = thrift.ZERO
var _ = fmt.Printf
var _ = bytes.Equal
var _ = time.Now()

var GoUnusedProtection__ int

type MetricID string

func MetricIDPtr(v MetricID) *MetricID { return &v }

type Metrics map[MetricID]float64

func MetricsPtr(v Metrics) *Metrics { return &v }

// Attributes:
//  - Key
type UnknownMetric struct {
	Key MetricID `thrift:"key,1,required" json:"key"`
}

func NewUnknownMetric() *UnknownMetric {
	return &UnknownMetric{}
}

func (p *UnknownMetric) GetKey() MetricID {
	return p.Key
}
func (p *UnknownMetric) Read(iprot thrift.TProtocol) error {
	if _, err := iprot.ReadStructBegin(); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T read error: ", p), err)
	}

	var issetKey bool = false

	for {
		_, fieldTypeId, fieldId, err := iprot.ReadFieldBegin()
		if err != nil {
			return thrift.PrependError(fmt.Sprintf("%T field %d read error: ", p, fieldId), err)
		}
		if fieldTypeId == thrift.STOP {
			break
		}
		switch fieldId {
		case 1:
			if err := p.ReadField1(iprot); err != nil {
				return err
			}
			issetKey = true
		default:
			if err := iprot.Skip(fieldTypeId); err != nil {
				return err
			}
		}
		if err := iprot.ReadFieldEnd(); err != nil {
			return err
		}
	}
	if err := iprot.ReadStructEnd(); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T read struct end error: ", p), err)
	}
	if !issetKey {
		return thrift.NewTProtocolExceptionWithType(thrift.INVALID_DATA, fmt.Errorf("Required field Key is not set"))
	}
	return nil
}

func (p *UnknownMetric) ReadField1(iprot thrift.TProtocol) error {
	if v, err := iprot.ReadString(); err != nil {
		return thrift.PrependError("error reading field 1: ", err)
	} else {
		temp := MetricID(v)
		p.Key = temp
	}
	return nil
}

func (p *UnknownMetric) Write(oprot thrift.TProtocol) error {
	if err := oprot.WriteStructBegin("UnknownMetric"); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write struct begin error: ", p), err)
	}
	if err := p.writeField1(oprot); err != nil {
		return err
	}
	if err := oprot.WriteFieldStop(); err != nil {
		return thrift.PrependError("write field stop error: ", err)
	}
	if err := oprot.WriteStructEnd(); err != nil {
		return thrift.PrependError("write struct stop error: ", err)
	}
	return nil
}

func (p *UnknownMetric) writeField1(oprot thrift.TProtocol) (err error) {
	if err := oprot.WriteFieldBegin("key", thrift.STRING, 1); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field begin error 1:key: ", p), err)
	}
	if err := oprot.WriteString(string(p.Key)); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T.key (1) field write error: ", p), err)
	}
	if err := oprot.WriteFieldEnd(); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field end error 1:key: ", p), err)
	}
	return err
}

func (p *UnknownMetric) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("UnknownMetric(%+v)", *p)
}

func (p *UnknownMetric) Error() string {
	return p.String()
}

// Attributes:
//  - Reason
type ServiceNotAvailable struct {
	Reason string `thrift:"reason,1,required" json:"reason"`
}

func NewServiceNotAvailable() *ServiceNotAvailable {
	return &ServiceNotAvailable{}
}

func (p *ServiceNotAvailable) GetReason() string {
	return p.Reason
}
func (p *ServiceNotAvailable) Read(iprot thrift.TProtocol) error {
	if _, err := iprot.ReadStructBegin(); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T read error: ", p), err)
	}

	var issetReason bool = false

	for {
		_, fieldTypeId, fieldId, err := iprot.ReadFieldBegin()
		if err != nil {
			return thrift.PrependError(fmt.Sprintf("%T field %d read error: ", p, fieldId), err)
		}
		if fieldTypeId == thrift.STOP {
			break
		}
		switch fieldId {
		case 1:
			if err := p.ReadField1(iprot); err != nil {
				return err
			}
			issetReason = true
		default:
			if err := iprot.Skip(fieldTypeId); err != nil {
				return err
			}
		}
		if err := iprot.ReadFieldEnd(); err != nil {
			return err
		}
	}
	if err := iprot.ReadStructEnd(); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T read struct end error: ", p), err)
	}
	if !issetReason {
		return thrift.NewTProtocolExceptionWithType(thrift.INVALID_DATA, fmt.Errorf("Required field Reason is not set"))
	}
	return nil
}

func (p *ServiceNotAvailable) ReadField1(iprot thrift.TProtocol) error {
	if v, err := iprot.ReadString(); err != nil {
		return thrift.PrependError("error reading field 1: ", err)
	} else {
		p.Reason = v
	}
	return nil
}

func (p *ServiceNotAvailable) Write(oprot thrift.TProtocol) error {
	if err := oprot.WriteStructBegin("ServiceNotAvailable"); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write struct begin error: ", p), err)
	}
	if err := p.writeField1(oprot); err != nil {
		return err
	}
	if err := oprot.WriteFieldStop(); err != nil {
		return thrift.PrependError("write field stop error: ", err)
	}
	if err := oprot.WriteStructEnd(); err != nil {
		return thrift.PrependError("write struct stop error: ", err)
	}
	return nil
}

func (p *ServiceNotAvailable) writeField1(oprot thrift.TProtocol) (err error) {
	if err := oprot.WriteFieldBegin("reason", thrift.STRING, 1); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field begin error 1:reason: ", p), err)
	}
	if err := oprot.WriteString(string(p.Reason)); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T.reason (1) field write error: ", p), err)
	}
	if err := oprot.WriteFieldEnd(); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field end error 1:reason: ", p), err)
	}
	return err
}

func (p *ServiceNotAvailable) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("ServiceNotAvailable(%+v)", *p)
}

func (p *ServiceNotAvailable) Error() string {
	return p.String()
}
