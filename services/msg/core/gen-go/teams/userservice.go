// Autogenerated by Thrift Compiler (0.9.2)
// DO NOT EDIT UNLESS YOU ARE SURE THAT YOU KNOW WHAT YOU ARE DOING

package teams

import (
	"bytes"
	"fmt"
	"git.apache.org/thrift.git/lib/go/thrift"
)

// (needed to ensure safety because of naive import list construction.)
var _ = thrift.ZERO
var _ = fmt.Printf
var _ = bytes.Equal

type UserService interface { //Base service operations.  These dont care about authorization for now and
	//assume the user is authorized.  Authn (and possible Authz) have to be taken
	//care of seperately.

	// Removes all entries.
	RemoveAllUsers() (err error)
	// Get user info by ID or username/team
	//
	// Parameters:
	//  - User
	GetUser(user User) (r User, err error)
	// Saves a user.
	// 	If the ID param is empty:
	// 		If username/team does not already exist a new one is created.
	// 		otherwise, it is updated and returned if override=true otherwise
	// 		false is returned.
	// 	Otherwise:
	// 		If username/team does not exist then it is written as is (Create or Update)
	// 		otherwise if IDs of curr and existing are different errow is thrown,
	// 		otherwise object is updated.
	//
	// Parameters:
	//  - Request
	SaveUser(request SaveUserRequest) (err error)
}

//Base service operations.  These dont care about authorization for now and
//assume the user is authorized.  Authn (and possible Authz) have to be taken
//care of seperately.
type UserServiceClient struct {
	Transport       thrift.TTransport
	ProtocolFactory thrift.TProtocolFactory
	InputProtocol   thrift.TProtocol
	OutputProtocol  thrift.TProtocol
	SeqId           int32
}

func NewUserServiceClientFactory(t thrift.TTransport, f thrift.TProtocolFactory) *UserServiceClient {
	return &UserServiceClient{Transport: t,
		ProtocolFactory: f,
		InputProtocol:   f.GetProtocol(t),
		OutputProtocol:  f.GetProtocol(t),
		SeqId:           0,
	}
}

func NewUserServiceClientProtocol(t thrift.TTransport, iprot thrift.TProtocol, oprot thrift.TProtocol) *UserServiceClient {
	return &UserServiceClient{Transport: t,
		ProtocolFactory: nil,
		InputProtocol:   iprot,
		OutputProtocol:  oprot,
		SeqId:           0,
	}
}

// Removes all entries.
func (p *UserServiceClient) RemoveAllUsers() (err error) {
	if err = p.sendRemoveAllUsers(); err != nil {
		return
	}
	return p.recvRemoveAllUsers()
}

func (p *UserServiceClient) sendRemoveAllUsers() (err error) {
	oprot := p.OutputProtocol
	if oprot == nil {
		oprot = p.ProtocolFactory.GetProtocol(p.Transport)
		p.OutputProtocol = oprot
	}
	p.SeqId++
	if err = oprot.WriteMessageBegin("RemoveAllUsers", thrift.CALL, p.SeqId); err != nil {
		return
	}
	args := RemoveAllUsersArgs{}
	if err = args.Write(oprot); err != nil {
		return
	}
	if err = oprot.WriteMessageEnd(); err != nil {
		return
	}
	return oprot.Flush()
}

func (p *UserServiceClient) recvRemoveAllUsers() (err error) {
	iprot := p.InputProtocol
	if iprot == nil {
		iprot = p.ProtocolFactory.GetProtocol(p.Transport)
		p.InputProtocol = iprot
	}
	_, mTypeId, seqId, err := iprot.ReadMessageBegin()
	if err != nil {
		return
	}
	if mTypeId == thrift.EXCEPTION {
		error0 := thrift.NewTApplicationException(thrift.UNKNOWN_APPLICATION_EXCEPTION, "Unknown Exception")
		var error1 error
		error1, err = error0.Read(iprot)
		if err != nil {
			return
		}
		if err = iprot.ReadMessageEnd(); err != nil {
			return
		}
		err = error1
		return
	}
	if p.SeqId != seqId {
		err = thrift.NewTApplicationException(thrift.BAD_SEQUENCE_ID, "RemoveAllUsers failed: out of sequence response")
		return
	}
	result := RemoveAllUsersResult{}
	if err = result.Read(iprot); err != nil {
		return
	}
	if err = iprot.ReadMessageEnd(); err != nil {
		return
	}
	return
}

// Get user info by ID or username/team
//
// Parameters:
//  - User
func (p *UserServiceClient) GetUser(user User) (r User, err error) {
	if err = p.sendGetUser(user); err != nil {
		return
	}
	return p.recvGetUser()
}

func (p *UserServiceClient) sendGetUser(user User) (err error) {
	oprot := p.OutputProtocol
	if oprot == nil {
		oprot = p.ProtocolFactory.GetProtocol(p.Transport)
		p.OutputProtocol = oprot
	}
	p.SeqId++
	if err = oprot.WriteMessageBegin("GetUser", thrift.CALL, p.SeqId); err != nil {
		return
	}
	args := GetUserArgs{
		User: user,
	}
	if err = args.Write(oprot); err != nil {
		return
	}
	if err = oprot.WriteMessageEnd(); err != nil {
		return
	}
	return oprot.Flush()
}

func (p *UserServiceClient) recvGetUser() (value User, err error) {
	iprot := p.InputProtocol
	if iprot == nil {
		iprot = p.ProtocolFactory.GetProtocol(p.Transport)
		p.InputProtocol = iprot
	}
	_, mTypeId, seqId, err := iprot.ReadMessageBegin()
	if err != nil {
		return
	}
	if mTypeId == thrift.EXCEPTION {
		error2 := thrift.NewTApplicationException(thrift.UNKNOWN_APPLICATION_EXCEPTION, "Unknown Exception")
		var error3 error
		error3, err = error2.Read(iprot)
		if err != nil {
			return
		}
		if err = iprot.ReadMessageEnd(); err != nil {
			return
		}
		err = error3
		return
	}
	if p.SeqId != seqId {
		err = thrift.NewTApplicationException(thrift.BAD_SEQUENCE_ID, "GetUser failed: out of sequence response")
		return
	}
	result := GetUserResult{}
	if err = result.Read(iprot); err != nil {
		return
	}
	if err = iprot.ReadMessageEnd(); err != nil {
		return
	}
	if result.ErrorA1 != nil {
		err = result.ErrorA1
		return
	}
	value = result.GetSuccess()
	return
}

// Saves a user.
// 	If the ID param is empty:
// 		If username/team does not already exist a new one is created.
// 		otherwise, it is updated and returned if override=true otherwise
// 		false is returned.
// 	Otherwise:
// 		If username/team does not exist then it is written as is (Create or Update)
// 		otherwise if IDs of curr and existing are different errow is thrown,
// 		otherwise object is updated.
//
// Parameters:
//  - Request
func (p *UserServiceClient) SaveUser(request SaveUserRequest) (err error) {
	if err = p.sendSaveUser(request); err != nil {
		return
	}
	return p.recvSaveUser()
}

func (p *UserServiceClient) sendSaveUser(request SaveUserRequest) (err error) {
	oprot := p.OutputProtocol
	if oprot == nil {
		oprot = p.ProtocolFactory.GetProtocol(p.Transport)
		p.OutputProtocol = oprot
	}
	p.SeqId++
	if err = oprot.WriteMessageBegin("SaveUser", thrift.CALL, p.SeqId); err != nil {
		return
	}
	args := SaveUserArgs{
		Request: request,
	}
	if err = args.Write(oprot); err != nil {
		return
	}
	if err = oprot.WriteMessageEnd(); err != nil {
		return
	}
	return oprot.Flush()
}

func (p *UserServiceClient) recvSaveUser() (err error) {
	iprot := p.InputProtocol
	if iprot == nil {
		iprot = p.ProtocolFactory.GetProtocol(p.Transport)
		p.InputProtocol = iprot
	}
	_, mTypeId, seqId, err := iprot.ReadMessageBegin()
	if err != nil {
		return
	}
	if mTypeId == thrift.EXCEPTION {
		error4 := thrift.NewTApplicationException(thrift.UNKNOWN_APPLICATION_EXCEPTION, "Unknown Exception")
		var error5 error
		error5, err = error4.Read(iprot)
		if err != nil {
			return
		}
		if err = iprot.ReadMessageEnd(); err != nil {
			return
		}
		err = error5
		return
	}
	if p.SeqId != seqId {
		err = thrift.NewTApplicationException(thrift.BAD_SEQUENCE_ID, "SaveUser failed: out of sequence response")
		return
	}
	result := SaveUserResult{}
	if err = result.Read(iprot); err != nil {
		return
	}
	if err = iprot.ReadMessageEnd(); err != nil {
		return
	}
	if result.ErrorA1 != nil {
		err = result.ErrorA1
		return
	}
	return
}

type UserServiceProcessor struct {
	processorMap map[string]thrift.TProcessorFunction
	handler      UserService
}

func (p *UserServiceProcessor) AddToProcessorMap(key string, processor thrift.TProcessorFunction) {
	p.processorMap[key] = processor
}

func (p *UserServiceProcessor) GetProcessorFunction(key string) (processor thrift.TProcessorFunction, ok bool) {
	processor, ok = p.processorMap[key]
	return processor, ok
}

func (p *UserServiceProcessor) ProcessorMap() map[string]thrift.TProcessorFunction {
	return p.processorMap
}

func NewUserServiceProcessor(handler UserService) *UserServiceProcessor {

	self6 := &UserServiceProcessor{handler: handler, processorMap: make(map[string]thrift.TProcessorFunction)}
	self6.processorMap["RemoveAllUsers"] = &userServiceProcessorRemoveAllUsers{handler: handler}
	self6.processorMap["GetUser"] = &userServiceProcessorGetUser{handler: handler}
	self6.processorMap["SaveUser"] = &userServiceProcessorSaveUser{handler: handler}
	return self6
}

func (p *UserServiceProcessor) Process(iprot, oprot thrift.TProtocol) (success bool, err thrift.TException) {
	name, _, seqId, err := iprot.ReadMessageBegin()
	if err != nil {
		return false, err
	}
	if processor, ok := p.GetProcessorFunction(name); ok {
		return processor.Process(seqId, iprot, oprot)
	}
	iprot.Skip(thrift.STRUCT)
	iprot.ReadMessageEnd()
	x7 := thrift.NewTApplicationException(thrift.UNKNOWN_METHOD, "Unknown function "+name)
	oprot.WriteMessageBegin(name, thrift.EXCEPTION, seqId)
	x7.Write(oprot)
	oprot.WriteMessageEnd()
	oprot.Flush()
	return false, x7

}

type userServiceProcessorRemoveAllUsers struct {
	handler UserService
}

func (p *userServiceProcessorRemoveAllUsers) Process(seqId int32, iprot, oprot thrift.TProtocol) (success bool, err thrift.TException) {
	args := RemoveAllUsersArgs{}
	if err = args.Read(iprot); err != nil {
		iprot.ReadMessageEnd()
		x := thrift.NewTApplicationException(thrift.PROTOCOL_ERROR, err.Error())
		oprot.WriteMessageBegin("RemoveAllUsers", thrift.EXCEPTION, seqId)
		x.Write(oprot)
		oprot.WriteMessageEnd()
		oprot.Flush()
		return false, err
	}

	iprot.ReadMessageEnd()
	result := RemoveAllUsersResult{}
	var err2 error
	if err2 = p.handler.RemoveAllUsers(); err2 != nil {
		x := thrift.NewTApplicationException(thrift.INTERNAL_ERROR, "Internal error processing RemoveAllUsers: "+err2.Error())
		oprot.WriteMessageBegin("RemoveAllUsers", thrift.EXCEPTION, seqId)
		x.Write(oprot)
		oprot.WriteMessageEnd()
		oprot.Flush()
		return true, err2
	}
	if err2 = oprot.WriteMessageBegin("RemoveAllUsers", thrift.REPLY, seqId); err2 != nil {
		err = err2
	}
	if err2 = result.Write(oprot); err == nil && err2 != nil {
		err = err2
	}
	if err2 = oprot.WriteMessageEnd(); err == nil && err2 != nil {
		err = err2
	}
	if err2 = oprot.Flush(); err == nil && err2 != nil {
		err = err2
	}
	if err != nil {
		return
	}
	return true, err
}

type userServiceProcessorGetUser struct {
	handler UserService
}

func (p *userServiceProcessorGetUser) Process(seqId int32, iprot, oprot thrift.TProtocol) (success bool, err thrift.TException) {
	args := GetUserArgs{}
	if err = args.Read(iprot); err != nil {
		iprot.ReadMessageEnd()
		x := thrift.NewTApplicationException(thrift.PROTOCOL_ERROR, err.Error())
		oprot.WriteMessageBegin("GetUser", thrift.EXCEPTION, seqId)
		x.Write(oprot)
		oprot.WriteMessageEnd()
		oprot.Flush()
		return false, err
	}

	iprot.ReadMessageEnd()
	result := GetUserResult{}
	var retval User
	var err2 error
	if retval, err2 = p.handler.GetUser(args.User); err2 != nil {
		switch v := err2.(type) {
		case *Error:
			result.ErrorA1 = v
		default:
			x := thrift.NewTApplicationException(thrift.INTERNAL_ERROR, "Internal error processing GetUser: "+err2.Error())
			oprot.WriteMessageBegin("GetUser", thrift.EXCEPTION, seqId)
			x.Write(oprot)
			oprot.WriteMessageEnd()
			oprot.Flush()
			return true, err2
		}
	} else {
		result.Success = retval
	}
	if err2 = oprot.WriteMessageBegin("GetUser", thrift.REPLY, seqId); err2 != nil {
		err = err2
	}
	if err2 = result.Write(oprot); err == nil && err2 != nil {
		err = err2
	}
	if err2 = oprot.WriteMessageEnd(); err == nil && err2 != nil {
		err = err2
	}
	if err2 = oprot.Flush(); err == nil && err2 != nil {
		err = err2
	}
	if err != nil {
		return
	}
	return true, err
}

type userServiceProcessorSaveUser struct {
	handler UserService
}

func (p *userServiceProcessorSaveUser) Process(seqId int32, iprot, oprot thrift.TProtocol) (success bool, err thrift.TException) {
	args := SaveUserArgs{}
	if err = args.Read(iprot); err != nil {
		iprot.ReadMessageEnd()
		x := thrift.NewTApplicationException(thrift.PROTOCOL_ERROR, err.Error())
		oprot.WriteMessageBegin("SaveUser", thrift.EXCEPTION, seqId)
		x.Write(oprot)
		oprot.WriteMessageEnd()
		oprot.Flush()
		return false, err
	}

	iprot.ReadMessageEnd()
	result := SaveUserResult{}
	var err2 error
	if err2 = p.handler.SaveUser(args.Request); err2 != nil {
		switch v := err2.(type) {
		case *Error:
			result.ErrorA1 = v
		default:
			x := thrift.NewTApplicationException(thrift.INTERNAL_ERROR, "Internal error processing SaveUser: "+err2.Error())
			oprot.WriteMessageBegin("SaveUser", thrift.EXCEPTION, seqId)
			x.Write(oprot)
			oprot.WriteMessageEnd()
			oprot.Flush()
			return true, err2
		}
	}
	if err2 = oprot.WriteMessageBegin("SaveUser", thrift.REPLY, seqId); err2 != nil {
		err = err2
	}
	if err2 = result.Write(oprot); err == nil && err2 != nil {
		err = err2
	}
	if err2 = oprot.WriteMessageEnd(); err == nil && err2 != nil {
		err = err2
	}
	if err2 = oprot.Flush(); err == nil && err2 != nil {
		err = err2
	}
	if err != nil {
		return
	}
	return true, err
}

// HELPER FUNCTIONS AND STRUCTURES

type RemoveAllUsersArgs struct {
}

func NewRemoveAllUsersArgs() *RemoveAllUsersArgs {
	return &RemoveAllUsersArgs{}
}

func (p *RemoveAllUsersArgs) Read(iprot thrift.TProtocol) error {
	if _, err := iprot.ReadStructBegin(); err != nil {
		return fmt.Errorf("%T read error: %s", p, err)
	}
	for {
		_, fieldTypeId, fieldId, err := iprot.ReadFieldBegin()
		if err != nil {
			return fmt.Errorf("%T field %d read error: %s", p, fieldId, err)
		}
		if fieldTypeId == thrift.STOP {
			break
		}
		if err := iprot.Skip(fieldTypeId); err != nil {
			return err
		}
		if err := iprot.ReadFieldEnd(); err != nil {
			return err
		}
	}
	if err := iprot.ReadStructEnd(); err != nil {
		return fmt.Errorf("%T read struct end error: %s", p, err)
	}
	return nil
}

func (p *RemoveAllUsersArgs) Write(oprot thrift.TProtocol) error {
	if err := oprot.WriteStructBegin("RemoveAllUsers_args"); err != nil {
		return fmt.Errorf("%T write struct begin error: %s", p, err)
	}
	if err := oprot.WriteFieldStop(); err != nil {
		return fmt.Errorf("write field stop error: %s", err)
	}
	if err := oprot.WriteStructEnd(); err != nil {
		return fmt.Errorf("write struct stop error: %s", err)
	}
	return nil
}

func (p *RemoveAllUsersArgs) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("RemoveAllUsersArgs(%+v)", *p)
}

type RemoveAllUsersResult struct {
}

func NewRemoveAllUsersResult() *RemoveAllUsersResult {
	return &RemoveAllUsersResult{}
}

func (p *RemoveAllUsersResult) Read(iprot thrift.TProtocol) error {
	if _, err := iprot.ReadStructBegin(); err != nil {
		return fmt.Errorf("%T read error: %s", p, err)
	}
	for {
		_, fieldTypeId, fieldId, err := iprot.ReadFieldBegin()
		if err != nil {
			return fmt.Errorf("%T field %d read error: %s", p, fieldId, err)
		}
		if fieldTypeId == thrift.STOP {
			break
		}
		if err := iprot.Skip(fieldTypeId); err != nil {
			return err
		}
		if err := iprot.ReadFieldEnd(); err != nil {
			return err
		}
	}
	if err := iprot.ReadStructEnd(); err != nil {
		return fmt.Errorf("%T read struct end error: %s", p, err)
	}
	return nil
}

func (p *RemoveAllUsersResult) Write(oprot thrift.TProtocol) error {
	if err := oprot.WriteStructBegin("RemoveAllUsers_result"); err != nil {
		return fmt.Errorf("%T write struct begin error: %s", p, err)
	}
	if err := oprot.WriteFieldStop(); err != nil {
		return fmt.Errorf("write field stop error: %s", err)
	}
	if err := oprot.WriteStructEnd(); err != nil {
		return fmt.Errorf("write struct stop error: %s", err)
	}
	return nil
}

func (p *RemoveAllUsersResult) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("RemoveAllUsersResult(%+v)", *p)
}

type GetUserArgs struct {
	User *User `thrift:"user,1" json:"user"`
}

func NewGetUserArgs() *GetUserArgs {
	return &GetUserArgs{}
}

var GetUserArgs_User_DEFAULT User

func (p *GetUserArgs) GetUser() User {
	if !p.IsSetUser() {
		return GetUserArgs_User_DEFAULT
	}
	return *p.User
}
func (p *GetUserArgs) IsSetUser() bool {
	return p.User != nil
}

func (p *GetUserArgs) Read(iprot thrift.TProtocol) error {
	if _, err := iprot.ReadStructBegin(); err != nil {
		return fmt.Errorf("%T read error: %s", p, err)
	}
	for {
		_, fieldTypeId, fieldId, err := iprot.ReadFieldBegin()
		if err != nil {
			return fmt.Errorf("%T field %d read error: %s", p, fieldId, err)
		}
		if fieldTypeId == thrift.STOP {
			break
		}
		switch fieldId {
		case 1:
			if err := p.ReadField1(iprot); err != nil {
				return err
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
		return fmt.Errorf("%T read struct end error: %s", p, err)
	}
	return nil
}

func (p *GetUserArgs) ReadField1(iprot thrift.TProtocol) error {
	p.User = &User{}
	if err := p.User.Read(iprot); err != nil {
		return fmt.Errorf("%T error reading struct: %s", p.User, err)
	}
	return nil
}

func (p *GetUserArgs) Write(oprot thrift.TProtocol) error {
	if err := oprot.WriteStructBegin("GetUser_args"); err != nil {
		return fmt.Errorf("%T write struct begin error: %s", p, err)
	}
	if err := p.writeField1(oprot); err != nil {
		return err
	}
	if err := oprot.WriteFieldStop(); err != nil {
		return fmt.Errorf("write field stop error: %s", err)
	}
	if err := oprot.WriteStructEnd(); err != nil {
		return fmt.Errorf("write struct stop error: %s", err)
	}
	return nil
}

func (p *GetUserArgs) writeField1(oprot thrift.TProtocol) (err error) {
	if err := oprot.WriteFieldBegin("user", thrift.STRUCT, 1); err != nil {
		return fmt.Errorf("%T write field begin error 1:user: %s", p, err)
	}
	if err := p.User.Write(oprot); err != nil {
		return fmt.Errorf("%T error writing struct: %s", p.User, err)
	}
	if err := oprot.WriteFieldEnd(); err != nil {
		return fmt.Errorf("%T write field end error 1:user: %s", p, err)
	}
	return err
}

func (p *GetUserArgs) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("GetUserArgs(%+v)", *p)
}

type GetUserResult struct {
	Success *User  `thrift:"success,0" json:"success"`
	ErrorA1 *Error `thrift:"error,1" json:"error"`
}

func NewGetUserResult() *GetUserResult {
	return &GetUserResult{}
}

var GetUserResult_Success_DEFAULT User

func (p *GetUserResult) GetSuccess() User {
	if !p.IsSetSuccess() {
		return GetUserResult_Success_DEFAULT
	}
	return *p.Success
}

var GetUserResult_ErrorA1_DEFAULT *Error

func (p *GetUserResult) GetErrorA1() *Error {
	if !p.IsSetErrorA1() {
		return GetUserResult_ErrorA1_DEFAULT
	}
	return p.ErrorA1
}
func (p *GetUserResult) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *GetUserResult) IsSetErrorA1() bool {
	return p.ErrorA1 != nil
}

func (p *GetUserResult) Read(iprot thrift.TProtocol) error {
	if _, err := iprot.ReadStructBegin(); err != nil {
		return fmt.Errorf("%T read error: %s", p, err)
	}
	for {
		_, fieldTypeId, fieldId, err := iprot.ReadFieldBegin()
		if err != nil {
			return fmt.Errorf("%T field %d read error: %s", p, fieldId, err)
		}
		if fieldTypeId == thrift.STOP {
			break
		}
		switch fieldId {
		case 0:
			if err := p.ReadField0(iprot); err != nil {
				return err
			}
		case 1:
			if err := p.ReadField1(iprot); err != nil {
				return err
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
		return fmt.Errorf("%T read struct end error: %s", p, err)
	}
	return nil
}

func (p *GetUserResult) ReadField0(iprot thrift.TProtocol) error {
	p.Success = &User{}
	if err := p.Success.Read(iprot); err != nil {
		return fmt.Errorf("%T error reading struct: %s", p.Success, err)
	}
	return nil
}

func (p *GetUserResult) ReadField1(iprot thrift.TProtocol) error {
	p.ErrorA1 = &Error{}
	if err := p.ErrorA1.Read(iprot); err != nil {
		return fmt.Errorf("%T error reading struct: %s", p.ErrorA1, err)
	}
	return nil
}

func (p *GetUserResult) Write(oprot thrift.TProtocol) error {
	if err := oprot.WriteStructBegin("GetUser_result"); err != nil {
		return fmt.Errorf("%T write struct begin error: %s", p, err)
	}
	if err := p.writeField0(oprot); err != nil {
		return err
	}
	if err := p.writeField1(oprot); err != nil {
		return err
	}
	if err := oprot.WriteFieldStop(); err != nil {
		return fmt.Errorf("write field stop error: %s", err)
	}
	if err := oprot.WriteStructEnd(); err != nil {
		return fmt.Errorf("write struct stop error: %s", err)
	}
	return nil
}

func (p *GetUserResult) writeField0(oprot thrift.TProtocol) (err error) {
	if p.IsSetSuccess() {
		if err := oprot.WriteFieldBegin("success", thrift.STRUCT, 0); err != nil {
			return fmt.Errorf("%T write field begin error 0:success: %s", p, err)
		}
		if err := p.Success.Write(oprot); err != nil {
			return fmt.Errorf("%T error writing struct: %s", p.Success, err)
		}
		if err := oprot.WriteFieldEnd(); err != nil {
			return fmt.Errorf("%T write field end error 0:success: %s", p, err)
		}
	}
	return err
}

func (p *GetUserResult) writeField1(oprot thrift.TProtocol) (err error) {
	if p.IsSetErrorA1() {
		if err := oprot.WriteFieldBegin("error", thrift.STRUCT, 1); err != nil {
			return fmt.Errorf("%T write field begin error 1:error: %s", p, err)
		}
		if err := p.ErrorA1.Write(oprot); err != nil {
			return fmt.Errorf("%T error writing struct: %s", p.ErrorA1, err)
		}
		if err := oprot.WriteFieldEnd(); err != nil {
			return fmt.Errorf("%T write field end error 1:error: %s", p, err)
		}
	}
	return err
}

func (p *GetUserResult) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("GetUserResult(%+v)", *p)
}

type SaveUserArgs struct {
	Request *SaveUserRequest `thrift:"request,1" json:"request"`
}

func NewSaveUserArgs() *SaveUserArgs {
	return &SaveUserArgs{}
}

var SaveUserArgs_Request_DEFAULT SaveUserRequest

func (p *SaveUserArgs) GetRequest() SaveUserRequest {
	if !p.IsSetRequest() {
		return SaveUserArgs_Request_DEFAULT
	}
	return *p.Request
}
func (p *SaveUserArgs) IsSetRequest() bool {
	return p.Request != nil
}

func (p *SaveUserArgs) Read(iprot thrift.TProtocol) error {
	if _, err := iprot.ReadStructBegin(); err != nil {
		return fmt.Errorf("%T read error: %s", p, err)
	}
	for {
		_, fieldTypeId, fieldId, err := iprot.ReadFieldBegin()
		if err != nil {
			return fmt.Errorf("%T field %d read error: %s", p, fieldId, err)
		}
		if fieldTypeId == thrift.STOP {
			break
		}
		switch fieldId {
		case 1:
			if err := p.ReadField1(iprot); err != nil {
				return err
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
		return fmt.Errorf("%T read struct end error: %s", p, err)
	}
	return nil
}

func (p *SaveUserArgs) ReadField1(iprot thrift.TProtocol) error {
	p.Request = &SaveUserRequest{}
	if err := p.Request.Read(iprot); err != nil {
		return fmt.Errorf("%T error reading struct: %s", p.Request, err)
	}
	return nil
}

func (p *SaveUserArgs) Write(oprot thrift.TProtocol) error {
	if err := oprot.WriteStructBegin("SaveUser_args"); err != nil {
		return fmt.Errorf("%T write struct begin error: %s", p, err)
	}
	if err := p.writeField1(oprot); err != nil {
		return err
	}
	if err := oprot.WriteFieldStop(); err != nil {
		return fmt.Errorf("write field stop error: %s", err)
	}
	if err := oprot.WriteStructEnd(); err != nil {
		return fmt.Errorf("write struct stop error: %s", err)
	}
	return nil
}

func (p *SaveUserArgs) writeField1(oprot thrift.TProtocol) (err error) {
	if err := oprot.WriteFieldBegin("request", thrift.STRUCT, 1); err != nil {
		return fmt.Errorf("%T write field begin error 1:request: %s", p, err)
	}
	if err := p.Request.Write(oprot); err != nil {
		return fmt.Errorf("%T error writing struct: %s", p.Request, err)
	}
	if err := oprot.WriteFieldEnd(); err != nil {
		return fmt.Errorf("%T write field end error 1:request: %s", p, err)
	}
	return err
}

func (p *SaveUserArgs) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("SaveUserArgs(%+v)", *p)
}

type SaveUserResult struct {
	ErrorA1 *Error `thrift:"error,1" json:"error"`
}

func NewSaveUserResult() *SaveUserResult {
	return &SaveUserResult{}
}

var SaveUserResult_ErrorA1_DEFAULT *Error

func (p *SaveUserResult) GetErrorA1() *Error {
	if !p.IsSetErrorA1() {
		return SaveUserResult_ErrorA1_DEFAULT
	}
	return p.ErrorA1
}
func (p *SaveUserResult) IsSetErrorA1() bool {
	return p.ErrorA1 != nil
}

func (p *SaveUserResult) Read(iprot thrift.TProtocol) error {
	if _, err := iprot.ReadStructBegin(); err != nil {
		return fmt.Errorf("%T read error: %s", p, err)
	}
	for {
		_, fieldTypeId, fieldId, err := iprot.ReadFieldBegin()
		if err != nil {
			return fmt.Errorf("%T field %d read error: %s", p, fieldId, err)
		}
		if fieldTypeId == thrift.STOP {
			break
		}
		switch fieldId {
		case 1:
			if err := p.ReadField1(iprot); err != nil {
				return err
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
		return fmt.Errorf("%T read struct end error: %s", p, err)
	}
	return nil
}

func (p *SaveUserResult) ReadField1(iprot thrift.TProtocol) error {
	p.ErrorA1 = &Error{}
	if err := p.ErrorA1.Read(iprot); err != nil {
		return fmt.Errorf("%T error reading struct: %s", p.ErrorA1, err)
	}
	return nil
}

func (p *SaveUserResult) Write(oprot thrift.TProtocol) error {
	if err := oprot.WriteStructBegin("SaveUser_result"); err != nil {
		return fmt.Errorf("%T write struct begin error: %s", p, err)
	}
	if err := p.writeField1(oprot); err != nil {
		return err
	}
	if err := oprot.WriteFieldStop(); err != nil {
		return fmt.Errorf("write field stop error: %s", err)
	}
	if err := oprot.WriteStructEnd(); err != nil {
		return fmt.Errorf("write struct stop error: %s", err)
	}
	return nil
}

func (p *SaveUserResult) writeField1(oprot thrift.TProtocol) (err error) {
	if p.IsSetErrorA1() {
		if err := oprot.WriteFieldBegin("error", thrift.STRUCT, 1); err != nil {
			return fmt.Errorf("%T write field begin error 1:error: %s", p, err)
		}
		if err := p.ErrorA1.Write(oprot); err != nil {
			return fmt.Errorf("%T error writing struct: %s", p.ErrorA1, err)
		}
		if err := oprot.WriteFieldEnd(); err != nil {
			return fmt.Errorf("%T write field end error 1:error: %s", p, err)
		}
	}
	return err
}

func (p *SaveUserResult) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("SaveUserResult(%+v)", *p)
}