// Autogenerated by Thrift Compiler (0.13.0)
// DO NOT EDIT UNLESS YOU ARE SURE THAT YOU KNOW WHAT YOU ARE DOING

package iron_man

import(
	"bytes"
	"context"
	"reflect"
	"fmt"
	"github.com/apache/thrift/lib/go/thrift"
	"github.com/fizx/iron-man/thrift/reddit/baseplate"

)

// (needed to ensure safety because of naive import list construction.)
var _ = thrift.ZERO
var _ = fmt.Printf
var _ = context.Background
var _ = reflect.DeepEqual
var _ = bytes.Equal

var _ = baseplate.GoUnusedProtection__
// Attributes:
//  - Content
type EchoRequest struct {
  Content string `thrift:"content,1" db:"content" json:"content"`
}

func NewEchoRequest() *EchoRequest {
  return &EchoRequest{}
}


func (p *EchoRequest) GetContent() string {
  return p.Content
}
func (p *EchoRequest) Read(iprot thrift.TProtocol) error {
  if _, err := iprot.ReadStructBegin(); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T read error: ", p), err)
  }


  for {
    _, fieldTypeId, fieldId, err := iprot.ReadFieldBegin()
    if err != nil {
      return thrift.PrependError(fmt.Sprintf("%T field %d read error: ", p, fieldId), err)
    }
    if fieldTypeId == thrift.STOP { break; }
    switch fieldId {
    case 1:
      if fieldTypeId == thrift.STRING {
        if err := p.ReadField1(iprot); err != nil {
          return err
        }
      } else {
        if err := iprot.Skip(fieldTypeId); err != nil {
          return err
        }
      }
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
  return nil
}

func (p *EchoRequest)  ReadField1(iprot thrift.TProtocol) error {
  if v, err := iprot.ReadString(); err != nil {
  return thrift.PrependError("error reading field 1: ", err)
} else {
  p.Content = v
}
  return nil
}

func (p *EchoRequest) Write(oprot thrift.TProtocol) error {
  if err := oprot.WriteStructBegin("EchoRequest"); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T write struct begin error: ", p), err) }
  if p != nil {
    if err := p.writeField1(oprot); err != nil { return err }
  }
  if err := oprot.WriteFieldStop(); err != nil {
    return thrift.PrependError("write field stop error: ", err) }
  if err := oprot.WriteStructEnd(); err != nil {
    return thrift.PrependError("write struct stop error: ", err) }
  return nil
}

func (p *EchoRequest) writeField1(oprot thrift.TProtocol) (err error) {
  if err := oprot.WriteFieldBegin("content", thrift.STRING, 1); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T write field begin error 1:content: ", p), err) }
  if err := oprot.WriteString(string(p.Content)); err != nil {
  return thrift.PrependError(fmt.Sprintf("%T.content (1) field write error: ", p), err) }
  if err := oprot.WriteFieldEnd(); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T write field end error 1:content: ", p), err) }
  return err
}

func (p *EchoRequest) String() string {
  if p == nil {
    return "<nil>"
  }
  return fmt.Sprintf("EchoRequest(%+v)", *p)
}

// Attributes:
//  - Content
type EchoResponse struct {
  Content string `thrift:"content,1" db:"content" json:"content"`
}

func NewEchoResponse() *EchoResponse {
  return &EchoResponse{}
}


func (p *EchoResponse) GetContent() string {
  return p.Content
}
func (p *EchoResponse) Read(iprot thrift.TProtocol) error {
  if _, err := iprot.ReadStructBegin(); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T read error: ", p), err)
  }


  for {
    _, fieldTypeId, fieldId, err := iprot.ReadFieldBegin()
    if err != nil {
      return thrift.PrependError(fmt.Sprintf("%T field %d read error: ", p, fieldId), err)
    }
    if fieldTypeId == thrift.STOP { break; }
    switch fieldId {
    case 1:
      if fieldTypeId == thrift.STRING {
        if err := p.ReadField1(iprot); err != nil {
          return err
        }
      } else {
        if err := iprot.Skip(fieldTypeId); err != nil {
          return err
        }
      }
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
  return nil
}

func (p *EchoResponse)  ReadField1(iprot thrift.TProtocol) error {
  if v, err := iprot.ReadString(); err != nil {
  return thrift.PrependError("error reading field 1: ", err)
} else {
  p.Content = v
}
  return nil
}

func (p *EchoResponse) Write(oprot thrift.TProtocol) error {
  if err := oprot.WriteStructBegin("EchoResponse"); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T write struct begin error: ", p), err) }
  if p != nil {
    if err := p.writeField1(oprot); err != nil { return err }
  }
  if err := oprot.WriteFieldStop(); err != nil {
    return thrift.PrependError("write field stop error: ", err) }
  if err := oprot.WriteStructEnd(); err != nil {
    return thrift.PrependError("write struct stop error: ", err) }
  return nil
}

func (p *EchoResponse) writeField1(oprot thrift.TProtocol) (err error) {
  if err := oprot.WriteFieldBegin("content", thrift.STRING, 1); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T write field begin error 1:content: ", p), err) }
  if err := oprot.WriteString(string(p.Content)); err != nil {
  return thrift.PrependError(fmt.Sprintf("%T.content (1) field write error: ", p), err) }
  if err := oprot.WriteFieldEnd(); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T write field end error 1:content: ", p), err) }
  return err
}

func (p *EchoResponse) String() string {
  if p == nil {
    return "<nil>"
  }
  return fmt.Sprintf("EchoResponse(%+v)", *p)
}

type IronManService interface {
  baseplate.BaseplateService

  // Parameters:
  //  - Request
  Echo(ctx context.Context, request *EchoRequest) (r *EchoResponse, err error)
}

type IronManServiceClient struct {
  *baseplate.BaseplateServiceClient
}

func NewIronManServiceClientFactory(t thrift.TTransport, f thrift.TProtocolFactory) *IronManServiceClient {
  return &IronManServiceClient{BaseplateServiceClient: baseplate.NewBaseplateServiceClientFactory(t, f)}}

func NewIronManServiceClientProtocol(t thrift.TTransport, iprot thrift.TProtocol, oprot thrift.TProtocol) *IronManServiceClient {
  return &IronManServiceClient{BaseplateServiceClient: baseplate.NewBaseplateServiceClientProtocol(t, iprot, oprot)}
}

func NewIronManServiceClient(c thrift.TClient) *IronManServiceClient {
  return &IronManServiceClient{
    BaseplateServiceClient: baseplate.NewBaseplateServiceClient(c),
  }
}

// Parameters:
//  - Request
func (p *IronManServiceClient) Echo(ctx context.Context, request *EchoRequest) (r *EchoResponse, err error) {
  var _args0 IronManServiceEchoArgs
  _args0.Request = request
  var _result1 IronManServiceEchoResult
  if err = p.Client_().Call(ctx, "echo", &_args0, &_result1); err != nil {
    return
  }
  return _result1.GetSuccess(), nil
}

type IronManServiceProcessor struct {
  *baseplate.BaseplateServiceProcessor
}

func NewIronManServiceProcessor(handler IronManService) *IronManServiceProcessor {
  self2 := &IronManServiceProcessor{baseplate.NewBaseplateServiceProcessor(handler)}
  self2.AddToProcessorMap("echo", &ironManServiceProcessorEcho{handler:handler})
  return self2
}

type ironManServiceProcessorEcho struct {
  handler IronManService
}

func (p *ironManServiceProcessorEcho) Process(ctx context.Context, seqId int32, iprot, oprot thrift.TProtocol) (success bool, err thrift.TException) {
  args := IronManServiceEchoArgs{}
  if err = args.Read(iprot); err != nil {
    iprot.ReadMessageEnd()
    x := thrift.NewTApplicationException(thrift.PROTOCOL_ERROR, err.Error())
    oprot.WriteMessageBegin("echo", thrift.EXCEPTION, seqId)
    x.Write(oprot)
    oprot.WriteMessageEnd()
    oprot.Flush(ctx)
    return false, err
  }

  iprot.ReadMessageEnd()
  result := IronManServiceEchoResult{}
var retval *EchoResponse
  var err2 error
  if retval, err2 = p.handler.Echo(ctx, args.Request); err2 != nil {
    x := thrift.NewTApplicationException(thrift.INTERNAL_ERROR, "Internal error processing echo: " + err2.Error())
    oprot.WriteMessageBegin("echo", thrift.EXCEPTION, seqId)
    x.Write(oprot)
    oprot.WriteMessageEnd()
    oprot.Flush(ctx)
    return true, err2
  } else {
    result.Success = retval
}
  if err2 = oprot.WriteMessageBegin("echo", thrift.REPLY, seqId); err2 != nil {
    err = err2
  }
  if err2 = result.Write(oprot); err == nil && err2 != nil {
    err = err2
  }
  if err2 = oprot.WriteMessageEnd(); err == nil && err2 != nil {
    err = err2
  }
  if err2 = oprot.Flush(ctx); err == nil && err2 != nil {
    err = err2
  }
  if err != nil {
    return
  }
  return true, err
}


// HELPER FUNCTIONS AND STRUCTURES

// Attributes:
//  - Request
type IronManServiceEchoArgs struct {
  Request *EchoRequest `thrift:"request,1" db:"request" json:"request"`
}

func NewIronManServiceEchoArgs() *IronManServiceEchoArgs {
  return &IronManServiceEchoArgs{}
}

var IronManServiceEchoArgs_Request_DEFAULT *EchoRequest
func (p *IronManServiceEchoArgs) GetRequest() *EchoRequest {
  if !p.IsSetRequest() {
    return IronManServiceEchoArgs_Request_DEFAULT
  }
return p.Request
}
func (p *IronManServiceEchoArgs) IsSetRequest() bool {
  return p.Request != nil
}

func (p *IronManServiceEchoArgs) Read(iprot thrift.TProtocol) error {
  if _, err := iprot.ReadStructBegin(); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T read error: ", p), err)
  }


  for {
    _, fieldTypeId, fieldId, err := iprot.ReadFieldBegin()
    if err != nil {
      return thrift.PrependError(fmt.Sprintf("%T field %d read error: ", p, fieldId), err)
    }
    if fieldTypeId == thrift.STOP { break; }
    switch fieldId {
    case 1:
      if fieldTypeId == thrift.STRUCT {
        if err := p.ReadField1(iprot); err != nil {
          return err
        }
      } else {
        if err := iprot.Skip(fieldTypeId); err != nil {
          return err
        }
      }
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
  return nil
}

func (p *IronManServiceEchoArgs)  ReadField1(iprot thrift.TProtocol) error {
  p.Request = &EchoRequest{}
  if err := p.Request.Read(iprot); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T error reading struct: ", p.Request), err)
  }
  return nil
}

func (p *IronManServiceEchoArgs) Write(oprot thrift.TProtocol) error {
  if err := oprot.WriteStructBegin("echo_args"); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T write struct begin error: ", p), err) }
  if p != nil {
    if err := p.writeField1(oprot); err != nil { return err }
  }
  if err := oprot.WriteFieldStop(); err != nil {
    return thrift.PrependError("write field stop error: ", err) }
  if err := oprot.WriteStructEnd(); err != nil {
    return thrift.PrependError("write struct stop error: ", err) }
  return nil
}

func (p *IronManServiceEchoArgs) writeField1(oprot thrift.TProtocol) (err error) {
  if err := oprot.WriteFieldBegin("request", thrift.STRUCT, 1); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T write field begin error 1:request: ", p), err) }
  if err := p.Request.Write(oprot); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T error writing struct: ", p.Request), err)
  }
  if err := oprot.WriteFieldEnd(); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T write field end error 1:request: ", p), err) }
  return err
}

func (p *IronManServiceEchoArgs) String() string {
  if p == nil {
    return "<nil>"
  }
  return fmt.Sprintf("IronManServiceEchoArgs(%+v)", *p)
}

// Attributes:
//  - Success
type IronManServiceEchoResult struct {
  Success *EchoResponse `thrift:"success,0" db:"success" json:"success,omitempty"`
}

func NewIronManServiceEchoResult() *IronManServiceEchoResult {
  return &IronManServiceEchoResult{}
}

var IronManServiceEchoResult_Success_DEFAULT *EchoResponse
func (p *IronManServiceEchoResult) GetSuccess() *EchoResponse {
  if !p.IsSetSuccess() {
    return IronManServiceEchoResult_Success_DEFAULT
  }
return p.Success
}
func (p *IronManServiceEchoResult) IsSetSuccess() bool {
  return p.Success != nil
}

func (p *IronManServiceEchoResult) Read(iprot thrift.TProtocol) error {
  if _, err := iprot.ReadStructBegin(); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T read error: ", p), err)
  }


  for {
    _, fieldTypeId, fieldId, err := iprot.ReadFieldBegin()
    if err != nil {
      return thrift.PrependError(fmt.Sprintf("%T field %d read error: ", p, fieldId), err)
    }
    if fieldTypeId == thrift.STOP { break; }
    switch fieldId {
    case 0:
      if fieldTypeId == thrift.STRUCT {
        if err := p.ReadField0(iprot); err != nil {
          return err
        }
      } else {
        if err := iprot.Skip(fieldTypeId); err != nil {
          return err
        }
      }
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
  return nil
}

func (p *IronManServiceEchoResult)  ReadField0(iprot thrift.TProtocol) error {
  p.Success = &EchoResponse{}
  if err := p.Success.Read(iprot); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T error reading struct: ", p.Success), err)
  }
  return nil
}

func (p *IronManServiceEchoResult) Write(oprot thrift.TProtocol) error {
  if err := oprot.WriteStructBegin("echo_result"); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T write struct begin error: ", p), err) }
  if p != nil {
    if err := p.writeField0(oprot); err != nil { return err }
  }
  if err := oprot.WriteFieldStop(); err != nil {
    return thrift.PrependError("write field stop error: ", err) }
  if err := oprot.WriteStructEnd(); err != nil {
    return thrift.PrependError("write struct stop error: ", err) }
  return nil
}

func (p *IronManServiceEchoResult) writeField0(oprot thrift.TProtocol) (err error) {
  if p.IsSetSuccess() {
    if err := oprot.WriteFieldBegin("success", thrift.STRUCT, 0); err != nil {
      return thrift.PrependError(fmt.Sprintf("%T write field begin error 0:success: ", p), err) }
    if err := p.Success.Write(oprot); err != nil {
      return thrift.PrependError(fmt.Sprintf("%T error writing struct: ", p.Success), err)
    }
    if err := oprot.WriteFieldEnd(); err != nil {
      return thrift.PrependError(fmt.Sprintf("%T write field end error 0:success: ", p), err) }
  }
  return err
}

func (p *IronManServiceEchoResult) String() string {
  if p == nil {
    return "<nil>"
  }
  return fmt.Sprintf("IronManServiceEchoResult(%+v)", *p)
}


