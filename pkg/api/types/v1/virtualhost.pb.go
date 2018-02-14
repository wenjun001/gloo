// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: virtualhost.proto

package v1

import proto "github.com/gogo/protobuf/proto"
import fmt "fmt"
import math "math"
import google_protobuf "github.com/gogo/protobuf/types"
import _ "github.com/gogo/protobuf/gogoproto"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type VirtualHost struct {
	Name      string     `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Domains   []string   `protobuf:"bytes,2,rep,name=domains" json:"domains,omitempty"`
	Routes    []*Route   `protobuf:"bytes,3,rep,name=routes" json:"routes,omitempty"`
	SslConfig *SSLConfig `protobuf:"bytes,4,opt,name=ssl_config,json=sslConfig" json:"ssl_config,omitempty"`
	// read only
	Status    *Status `protobuf:"bytes,5,opt,name=status" json:"status,omitempty"`
	*Metadata `protobuf:"bytes,6,opt,name=metadata,embedded=metadata" json:"metadata,omitempty"`
}

func (m *VirtualHost) Reset()                    { *m = VirtualHost{} }
func (m *VirtualHost) String() string            { return proto.CompactTextString(m) }
func (*VirtualHost) ProtoMessage()               {}
func (*VirtualHost) Descriptor() ([]byte, []int) { return fileDescriptorVirtualhost, []int{0} }

func (m *VirtualHost) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *VirtualHost) GetDomains() []string {
	if m != nil {
		return m.Domains
	}
	return nil
}

func (m *VirtualHost) GetRoutes() []*Route {
	if m != nil {
		return m.Routes
	}
	return nil
}

func (m *VirtualHost) GetSslConfig() *SSLConfig {
	if m != nil {
		return m.SslConfig
	}
	return nil
}

func (m *VirtualHost) GetStatus() *Status {
	if m != nil {
		return m.Status
	}
	return nil
}

type Route struct {
	Matcher              *Matcher                `protobuf:"bytes,1,opt,name=matcher" json:"matcher,omitempty"`
	MultipleDestinations []*WeightedDestination  `protobuf:"bytes,2,rep,name=multiple_destinations,json=multipleDestinations" json:"multiple_destinations,omitempty"`
	SingleDestination    *Destination            `protobuf:"bytes,3,opt,name=single_destination,json=singleDestination" json:"single_destination,omitempty"`
	PrefixRewrite        string                  `protobuf:"bytes,4,opt,name=prefix_rewrite,json=prefixRewrite,proto3" json:"prefix_rewrite,omitempty"`
	Extensions           *google_protobuf.Struct `protobuf:"bytes,5,opt,name=extensions" json:"extensions,omitempty"`
}

func (m *Route) Reset()                    { *m = Route{} }
func (m *Route) String() string            { return proto.CompactTextString(m) }
func (*Route) ProtoMessage()               {}
func (*Route) Descriptor() ([]byte, []int) { return fileDescriptorVirtualhost, []int{1} }

func (m *Route) GetMatcher() *Matcher {
	if m != nil {
		return m.Matcher
	}
	return nil
}

func (m *Route) GetMultipleDestinations() []*WeightedDestination {
	if m != nil {
		return m.MultipleDestinations
	}
	return nil
}

func (m *Route) GetSingleDestination() *Destination {
	if m != nil {
		return m.SingleDestination
	}
	return nil
}

func (m *Route) GetPrefixRewrite() string {
	if m != nil {
		return m.PrefixRewrite
	}
	return ""
}

func (m *Route) GetExtensions() *google_protobuf.Struct {
	if m != nil {
		return m.Extensions
	}
	return nil
}

type Matcher struct {
	// Types that are valid to be assigned to Path:
	//	*Matcher_PathPrefix
	//	*Matcher_PathRegex
	//	*Matcher_PathExact
	Path        isMatcher_Path    `protobuf_oneof:"path"`
	Headers     map[string]string `protobuf:"bytes,4,rep,name=headers" json:"headers,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	QueryParams map[string]string `protobuf:"bytes,5,rep,name=query_params,json=queryParams" json:"query_params,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	Verbs       []string          `protobuf:"bytes,6,rep,name=verbs" json:"verbs,omitempty"`
}

func (m *Matcher) Reset()                    { *m = Matcher{} }
func (m *Matcher) String() string            { return proto.CompactTextString(m) }
func (*Matcher) ProtoMessage()               {}
func (*Matcher) Descriptor() ([]byte, []int) { return fileDescriptorVirtualhost, []int{2} }

type isMatcher_Path interface {
	isMatcher_Path()
	Equal(interface{}) bool
}

type Matcher_PathPrefix struct {
	PathPrefix string `protobuf:"bytes,1,opt,name=path_prefix,json=pathPrefix,proto3,oneof"`
}
type Matcher_PathRegex struct {
	PathRegex string `protobuf:"bytes,2,opt,name=path_regex,json=pathRegex,proto3,oneof"`
}
type Matcher_PathExact struct {
	PathExact string `protobuf:"bytes,3,opt,name=path_exact,json=pathExact,proto3,oneof"`
}

func (*Matcher_PathPrefix) isMatcher_Path() {}
func (*Matcher_PathRegex) isMatcher_Path()  {}
func (*Matcher_PathExact) isMatcher_Path()  {}

func (m *Matcher) GetPath() isMatcher_Path {
	if m != nil {
		return m.Path
	}
	return nil
}

func (m *Matcher) GetPathPrefix() string {
	if x, ok := m.GetPath().(*Matcher_PathPrefix); ok {
		return x.PathPrefix
	}
	return ""
}

func (m *Matcher) GetPathRegex() string {
	if x, ok := m.GetPath().(*Matcher_PathRegex); ok {
		return x.PathRegex
	}
	return ""
}

func (m *Matcher) GetPathExact() string {
	if x, ok := m.GetPath().(*Matcher_PathExact); ok {
		return x.PathExact
	}
	return ""
}

func (m *Matcher) GetHeaders() map[string]string {
	if m != nil {
		return m.Headers
	}
	return nil
}

func (m *Matcher) GetQueryParams() map[string]string {
	if m != nil {
		return m.QueryParams
	}
	return nil
}

func (m *Matcher) GetVerbs() []string {
	if m != nil {
		return m.Verbs
	}
	return nil
}

// XXX_OneofFuncs is for the internal use of the proto package.
func (*Matcher) XXX_OneofFuncs() (func(msg proto.Message, b *proto.Buffer) error, func(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error), func(msg proto.Message) (n int), []interface{}) {
	return _Matcher_OneofMarshaler, _Matcher_OneofUnmarshaler, _Matcher_OneofSizer, []interface{}{
		(*Matcher_PathPrefix)(nil),
		(*Matcher_PathRegex)(nil),
		(*Matcher_PathExact)(nil),
	}
}

func _Matcher_OneofMarshaler(msg proto.Message, b *proto.Buffer) error {
	m := msg.(*Matcher)
	// path
	switch x := m.Path.(type) {
	case *Matcher_PathPrefix:
		_ = b.EncodeVarint(1<<3 | proto.WireBytes)
		_ = b.EncodeStringBytes(x.PathPrefix)
	case *Matcher_PathRegex:
		_ = b.EncodeVarint(2<<3 | proto.WireBytes)
		_ = b.EncodeStringBytes(x.PathRegex)
	case *Matcher_PathExact:
		_ = b.EncodeVarint(3<<3 | proto.WireBytes)
		_ = b.EncodeStringBytes(x.PathExact)
	case nil:
	default:
		return fmt.Errorf("Matcher.Path has unexpected type %T", x)
	}
	return nil
}

func _Matcher_OneofUnmarshaler(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error) {
	m := msg.(*Matcher)
	switch tag {
	case 1: // path.path_prefix
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		x, err := b.DecodeStringBytes()
		m.Path = &Matcher_PathPrefix{x}
		return true, err
	case 2: // path.path_regex
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		x, err := b.DecodeStringBytes()
		m.Path = &Matcher_PathRegex{x}
		return true, err
	case 3: // path.path_exact
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		x, err := b.DecodeStringBytes()
		m.Path = &Matcher_PathExact{x}
		return true, err
	default:
		return false, nil
	}
}

func _Matcher_OneofSizer(msg proto.Message) (n int) {
	m := msg.(*Matcher)
	// path
	switch x := m.Path.(type) {
	case *Matcher_PathPrefix:
		n += proto.SizeVarint(1<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(len(x.PathPrefix)))
		n += len(x.PathPrefix)
	case *Matcher_PathRegex:
		n += proto.SizeVarint(2<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(len(x.PathRegex)))
		n += len(x.PathRegex)
	case *Matcher_PathExact:
		n += proto.SizeVarint(3<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(len(x.PathExact)))
		n += len(x.PathExact)
	case nil:
	default:
		panic(fmt.Sprintf("proto: unexpected type %T in oneof", x))
	}
	return n
}

type WeightedDestination struct {
	*Destination `protobuf:"bytes,1,opt,name=destination,embedded=destination" json:"destination,omitempty"`
	Weight       uint32 `protobuf:"varint,2,opt,name=weight,proto3" json:"weight,omitempty"`
}

func (m *WeightedDestination) Reset()                    { *m = WeightedDestination{} }
func (m *WeightedDestination) String() string            { return proto.CompactTextString(m) }
func (*WeightedDestination) ProtoMessage()               {}
func (*WeightedDestination) Descriptor() ([]byte, []int) { return fileDescriptorVirtualhost, []int{3} }

func (m *WeightedDestination) GetWeight() uint32 {
	if m != nil {
		return m.Weight
	}
	return 0
}

type Destination struct {
	// Types that are valid to be assigned to DestinationType:
	//	*Destination_Function
	//	*Destination_Upstream
	DestinationType isDestination_DestinationType `protobuf_oneof:"destination_type"`
}

func (m *Destination) Reset()                    { *m = Destination{} }
func (m *Destination) String() string            { return proto.CompactTextString(m) }
func (*Destination) ProtoMessage()               {}
func (*Destination) Descriptor() ([]byte, []int) { return fileDescriptorVirtualhost, []int{4} }

type isDestination_DestinationType interface {
	isDestination_DestinationType()
	Equal(interface{}) bool
}

type Destination_Function struct {
	Function *FunctionDestination `protobuf:"bytes,1,opt,name=function,oneof"`
}
type Destination_Upstream struct {
	Upstream *UpstreamDestination `protobuf:"bytes,2,opt,name=upstream,oneof"`
}

func (*Destination_Function) isDestination_DestinationType() {}
func (*Destination_Upstream) isDestination_DestinationType() {}

func (m *Destination) GetDestinationType() isDestination_DestinationType {
	if m != nil {
		return m.DestinationType
	}
	return nil
}

func (m *Destination) GetFunction() *FunctionDestination {
	if x, ok := m.GetDestinationType().(*Destination_Function); ok {
		return x.Function
	}
	return nil
}

func (m *Destination) GetUpstream() *UpstreamDestination {
	if x, ok := m.GetDestinationType().(*Destination_Upstream); ok {
		return x.Upstream
	}
	return nil
}

// XXX_OneofFuncs is for the internal use of the proto package.
func (*Destination) XXX_OneofFuncs() (func(msg proto.Message, b *proto.Buffer) error, func(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error), func(msg proto.Message) (n int), []interface{}) {
	return _Destination_OneofMarshaler, _Destination_OneofUnmarshaler, _Destination_OneofSizer, []interface{}{
		(*Destination_Function)(nil),
		(*Destination_Upstream)(nil),
	}
}

func _Destination_OneofMarshaler(msg proto.Message, b *proto.Buffer) error {
	m := msg.(*Destination)
	// destination_type
	switch x := m.DestinationType.(type) {
	case *Destination_Function:
		_ = b.EncodeVarint(1<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.Function); err != nil {
			return err
		}
	case *Destination_Upstream:
		_ = b.EncodeVarint(2<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.Upstream); err != nil {
			return err
		}
	case nil:
	default:
		return fmt.Errorf("Destination.DestinationType has unexpected type %T", x)
	}
	return nil
}

func _Destination_OneofUnmarshaler(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error) {
	m := msg.(*Destination)
	switch tag {
	case 1: // destination_type.function
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(FunctionDestination)
		err := b.DecodeMessage(msg)
		m.DestinationType = &Destination_Function{msg}
		return true, err
	case 2: // destination_type.upstream
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(UpstreamDestination)
		err := b.DecodeMessage(msg)
		m.DestinationType = &Destination_Upstream{msg}
		return true, err
	default:
		return false, nil
	}
}

func _Destination_OneofSizer(msg proto.Message) (n int) {
	m := msg.(*Destination)
	// destination_type
	switch x := m.DestinationType.(type) {
	case *Destination_Function:
		s := proto.Size(x.Function)
		n += proto.SizeVarint(1<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(s))
		n += s
	case *Destination_Upstream:
		s := proto.Size(x.Upstream)
		n += proto.SizeVarint(2<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(s))
		n += s
	case nil:
	default:
		panic(fmt.Sprintf("proto: unexpected type %T in oneof", x))
	}
	return n
}

type FunctionDestination struct {
	UpstreamName string `protobuf:"bytes,1,opt,name=upstream_name,json=upstreamName,proto3" json:"upstream_name,omitempty"`
	FunctionName string `protobuf:"bytes,2,opt,name=function_name,json=functionName,proto3" json:"function_name,omitempty"`
}

func (m *FunctionDestination) Reset()                    { *m = FunctionDestination{} }
func (m *FunctionDestination) String() string            { return proto.CompactTextString(m) }
func (*FunctionDestination) ProtoMessage()               {}
func (*FunctionDestination) Descriptor() ([]byte, []int) { return fileDescriptorVirtualhost, []int{5} }

func (m *FunctionDestination) GetUpstreamName() string {
	if m != nil {
		return m.UpstreamName
	}
	return ""
}

func (m *FunctionDestination) GetFunctionName() string {
	if m != nil {
		return m.FunctionName
	}
	return ""
}

type UpstreamDestination struct {
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
}

func (m *UpstreamDestination) Reset()                    { *m = UpstreamDestination{} }
func (m *UpstreamDestination) String() string            { return proto.CompactTextString(m) }
func (*UpstreamDestination) ProtoMessage()               {}
func (*UpstreamDestination) Descriptor() ([]byte, []int) { return fileDescriptorVirtualhost, []int{6} }

func (m *UpstreamDestination) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

type SSLConfig struct {
	SecretRef string `protobuf:"bytes,1,opt,name=secret_ref,json=secretRef,proto3" json:"secret_ref,omitempty"`
}

func (m *SSLConfig) Reset()                    { *m = SSLConfig{} }
func (m *SSLConfig) String() string            { return proto.CompactTextString(m) }
func (*SSLConfig) ProtoMessage()               {}
func (*SSLConfig) Descriptor() ([]byte, []int) { return fileDescriptorVirtualhost, []int{7} }

func (m *SSLConfig) GetSecretRef() string {
	if m != nil {
		return m.SecretRef
	}
	return ""
}

func init() {
	proto.RegisterType((*VirtualHost)(nil), "v1.VirtualHost")
	proto.RegisterType((*Route)(nil), "v1.Route")
	proto.RegisterType((*Matcher)(nil), "v1.Matcher")
	proto.RegisterType((*WeightedDestination)(nil), "v1.WeightedDestination")
	proto.RegisterType((*Destination)(nil), "v1.Destination")
	proto.RegisterType((*FunctionDestination)(nil), "v1.FunctionDestination")
	proto.RegisterType((*UpstreamDestination)(nil), "v1.UpstreamDestination")
	proto.RegisterType((*SSLConfig)(nil), "v1.SSLConfig")
}
func (this *VirtualHost) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*VirtualHost)
	if !ok {
		that2, ok := that.(VirtualHost)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if this.Name != that1.Name {
		return false
	}
	if len(this.Domains) != len(that1.Domains) {
		return false
	}
	for i := range this.Domains {
		if this.Domains[i] != that1.Domains[i] {
			return false
		}
	}
	if len(this.Routes) != len(that1.Routes) {
		return false
	}
	for i := range this.Routes {
		if !this.Routes[i].Equal(that1.Routes[i]) {
			return false
		}
	}
	if !this.SslConfig.Equal(that1.SslConfig) {
		return false
	}
	if !this.Status.Equal(that1.Status) {
		return false
	}
	if !this.Metadata.Equal(that1.Metadata) {
		return false
	}
	return true
}
func (this *Route) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*Route)
	if !ok {
		that2, ok := that.(Route)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if !this.Matcher.Equal(that1.Matcher) {
		return false
	}
	if len(this.MultipleDestinations) != len(that1.MultipleDestinations) {
		return false
	}
	for i := range this.MultipleDestinations {
		if !this.MultipleDestinations[i].Equal(that1.MultipleDestinations[i]) {
			return false
		}
	}
	if !this.SingleDestination.Equal(that1.SingleDestination) {
		return false
	}
	if this.PrefixRewrite != that1.PrefixRewrite {
		return false
	}
	if !this.Extensions.Equal(that1.Extensions) {
		return false
	}
	return true
}
func (this *Matcher) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*Matcher)
	if !ok {
		that2, ok := that.(Matcher)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if that1.Path == nil {
		if this.Path != nil {
			return false
		}
	} else if this.Path == nil {
		return false
	} else if !this.Path.Equal(that1.Path) {
		return false
	}
	if len(this.Headers) != len(that1.Headers) {
		return false
	}
	for i := range this.Headers {
		if this.Headers[i] != that1.Headers[i] {
			return false
		}
	}
	if len(this.QueryParams) != len(that1.QueryParams) {
		return false
	}
	for i := range this.QueryParams {
		if this.QueryParams[i] != that1.QueryParams[i] {
			return false
		}
	}
	if len(this.Verbs) != len(that1.Verbs) {
		return false
	}
	for i := range this.Verbs {
		if this.Verbs[i] != that1.Verbs[i] {
			return false
		}
	}
	return true
}
func (this *Matcher_PathPrefix) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*Matcher_PathPrefix)
	if !ok {
		that2, ok := that.(Matcher_PathPrefix)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if this.PathPrefix != that1.PathPrefix {
		return false
	}
	return true
}
func (this *Matcher_PathRegex) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*Matcher_PathRegex)
	if !ok {
		that2, ok := that.(Matcher_PathRegex)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if this.PathRegex != that1.PathRegex {
		return false
	}
	return true
}
func (this *Matcher_PathExact) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*Matcher_PathExact)
	if !ok {
		that2, ok := that.(Matcher_PathExact)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if this.PathExact != that1.PathExact {
		return false
	}
	return true
}
func (this *WeightedDestination) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*WeightedDestination)
	if !ok {
		that2, ok := that.(WeightedDestination)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if !this.Destination.Equal(that1.Destination) {
		return false
	}
	if this.Weight != that1.Weight {
		return false
	}
	return true
}
func (this *Destination) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*Destination)
	if !ok {
		that2, ok := that.(Destination)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if that1.DestinationType == nil {
		if this.DestinationType != nil {
			return false
		}
	} else if this.DestinationType == nil {
		return false
	} else if !this.DestinationType.Equal(that1.DestinationType) {
		return false
	}
	return true
}
func (this *Destination_Function) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*Destination_Function)
	if !ok {
		that2, ok := that.(Destination_Function)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if !this.Function.Equal(that1.Function) {
		return false
	}
	return true
}
func (this *Destination_Upstream) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*Destination_Upstream)
	if !ok {
		that2, ok := that.(Destination_Upstream)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if !this.Upstream.Equal(that1.Upstream) {
		return false
	}
	return true
}
func (this *FunctionDestination) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*FunctionDestination)
	if !ok {
		that2, ok := that.(FunctionDestination)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if this.UpstreamName != that1.UpstreamName {
		return false
	}
	if this.FunctionName != that1.FunctionName {
		return false
	}
	return true
}
func (this *UpstreamDestination) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*UpstreamDestination)
	if !ok {
		that2, ok := that.(UpstreamDestination)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if this.Name != that1.Name {
		return false
	}
	return true
}
func (this *SSLConfig) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*SSLConfig)
	if !ok {
		that2, ok := that.(SSLConfig)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if this.SecretRef != that1.SecretRef {
		return false
	}
	return true
}

func init() { proto.RegisterFile("virtualhost.proto", fileDescriptorVirtualhost) }

var fileDescriptorVirtualhost = []byte{
	// 710 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x54, 0x5d, 0x6e, 0xd3, 0x4c,
	0x14, 0xad, 0x93, 0x34, 0xa9, 0xaf, 0x93, 0x7e, 0xed, 0xb4, 0x1f, 0xb5, 0xa2, 0x42, 0x53, 0x57,
	0x95, 0x02, 0x42, 0xae, 0x1a, 0x84, 0x8a, 0xfa, 0x50, 0xa4, 0x40, 0x51, 0x1e, 0x0a, 0x2a, 0x53,
	0x01, 0x8f, 0xd6, 0x34, 0x99, 0x38, 0x16, 0x89, 0xed, 0xce, 0x8c, 0xd3, 0x64, 0x15, 0x6c, 0x80,
	0x05, 0xb0, 0x1c, 0x56, 0x80, 0x10, 0xfb, 0x40, 0x42, 0xf3, 0xe3, 0xd4, 0x54, 0x79, 0xe1, 0x6d,
	0xee, 0xb9, 0xe7, 0xdc, 0xbf, 0x9c, 0x18, 0x36, 0xa7, 0x11, 0x13, 0x19, 0x19, 0x8f, 0x12, 0x2e,
	0xfc, 0x94, 0x25, 0x22, 0x41, 0xa5, 0xe9, 0x71, 0x73, 0x37, 0x4c, 0x92, 0x70, 0x4c, 0x8f, 0x14,
	0x72, 0x9d, 0x0d, 0x8f, 0xb8, 0x60, 0x59, 0xdf, 0x30, 0x9a, 0xdb, 0x61, 0x12, 0x26, 0xea, 0x79,
	0x24, 0x5f, 0x06, 0xad, 0x73, 0x41, 0x44, 0xc6, 0x4d, 0xb4, 0x3e, 0xa1, 0x82, 0x0c, 0x88, 0x20,
	0x3a, 0xf6, 0x7e, 0x5a, 0xe0, 0x7c, 0xd4, 0xbd, 0x7a, 0x09, 0x17, 0x08, 0x41, 0x25, 0x26, 0x13,
	0xea, 0x5a, 0x2d, 0xab, 0x6d, 0x63, 0xf5, 0x46, 0x2e, 0xd4, 0x06, 0xc9, 0x84, 0x44, 0x31, 0x77,
	0x4b, 0xad, 0x72, 0xdb, 0xc6, 0x79, 0x88, 0xf6, 0xa1, 0xca, 0x92, 0x4c, 0x50, 0xee, 0x96, 0x5b,
	0xe5, 0xb6, 0xd3, 0xb1, 0xfd, 0xe9, 0xb1, 0x8f, 0x25, 0x82, 0x4d, 0x02, 0x3d, 0x05, 0xe0, 0x7c,
	0x1c, 0xf4, 0x93, 0x78, 0x18, 0x85, 0x6e, 0xa5, 0x65, 0xb5, 0x9d, 0x4e, 0x43, 0xd2, 0xae, 0xae,
	0x2e, 0x5e, 0x29, 0x10, 0xdb, 0x9c, 0x8f, 0xf5, 0x13, 0x79, 0x50, 0xd5, 0xe3, 0xba, 0xab, 0x8a,
	0x09, 0x8a, 0xa9, 0x10, 0x6c, 0x32, 0xc8, 0x87, 0xb5, 0x7c, 0x09, 0xb7, 0xaa, 0x58, 0x75, 0xc9,
	0x7a, 0x6b, 0xb0, 0x6e, 0xe5, 0xfb, 0x8f, 0x3d, 0x0b, 0x2f, 0x38, 0xde, 0xd7, 0x12, 0xac, 0xaa,
	0x99, 0xd0, 0x21, 0xd4, 0x26, 0x44, 0xf4, 0x47, 0x94, 0xa9, 0xfd, 0x9c, 0x8e, 0xa3, 0x84, 0x1a,
	0xc2, 0x79, 0x0e, 0x5d, 0xc0, 0xff, 0x93, 0x6c, 0x2c, 0xa2, 0x74, 0x4c, 0x83, 0x01, 0xe5, 0x22,
	0x8a, 0x89, 0x88, 0x12, 0xb3, 0xbd, 0xd3, 0xd9, 0x91, 0xa2, 0x4f, 0x34, 0x0a, 0x47, 0x82, 0x0e,
	0x5e, 0xdf, 0xe5, 0xf1, 0x76, 0xae, 0x2a, 0x80, 0x1c, 0x9d, 0x01, 0xe2, 0x51, 0x1c, 0xfe, 0x5d,
	0xcb, 0x2d, 0xab, 0xfe, 0xff, 0xc9, 0x52, 0xc5, 0x12, 0x9b, 0x9a, 0x5a, 0x80, 0xd0, 0x21, 0xac,
	0xa7, 0x8c, 0x0e, 0xa3, 0x59, 0xc0, 0xe8, 0x2d, 0x8b, 0x04, 0x55, 0x47, 0xb4, 0x71, 0x43, 0xa3,
	0x58, 0x83, 0xe8, 0x04, 0x80, 0xce, 0x04, 0x8d, 0xb9, 0x9a, 0x54, 0x5f, 0x6f, 0xc7, 0xd7, 0x7e,
	0xf1, 0x73, 0xbf, 0xf8, 0x57, 0xca, 0x2f, 0xb8, 0x40, 0xf5, 0x7e, 0x97, 0xa0, 0x66, 0x4e, 0x80,
	0xf6, 0xc1, 0x49, 0x89, 0x18, 0x05, 0xba, 0xb4, 0x36, 0x41, 0x6f, 0x05, 0x83, 0x04, 0x2f, 0x15,
	0x86, 0xf6, 0x40, 0x45, 0x01, 0xa3, 0x21, 0x9d, 0xb9, 0x25, 0xc3, 0xb0, 0x25, 0x86, 0x25, 0xb4,
	0x20, 0xd0, 0x19, 0xe9, 0x0b, 0xb5, 0xe7, 0x82, 0x70, 0x2e, 0x21, 0xd4, 0x81, 0xda, 0x88, 0x92,
	0x01, 0x65, 0xdc, 0xad, 0xa8, 0x83, 0xba, 0x85, 0x5f, 0xc1, 0xef, 0xe9, 0xd4, 0x79, 0x2c, 0xd8,
	0x1c, 0xe7, 0x44, 0xf4, 0x12, 0xea, 0x37, 0x19, 0x65, 0xf3, 0x20, 0x25, 0x8c, 0x4c, 0xe4, 0x7e,
	0x52, 0xb8, 0x5b, 0x14, 0xbe, 0x97, 0xf9, 0x4b, 0x95, 0xd6, 0x62, 0xe7, 0xe6, 0x0e, 0x41, 0xdb,
	0xb0, 0x3a, 0xa5, 0xec, 0x9a, 0xbb, 0x55, 0xe5, 0x60, 0x1d, 0x34, 0x4f, 0xa1, 0x5e, 0xec, 0x87,
	0x36, 0xa0, 0xfc, 0x99, 0xce, 0x8d, 0xf9, 0xe5, 0x53, 0xe9, 0xc8, 0x38, 0xa3, 0x7a, 0x53, 0xac,
	0x83, 0xd3, 0xd2, 0x0b, 0xab, 0x79, 0x06, 0x1b, 0xf7, 0x5b, 0xfe, 0x8b, 0xbe, 0x5b, 0x85, 0x8a,
	0xbc, 0x89, 0x37, 0x84, 0xad, 0x25, 0x66, 0x42, 0x27, 0xe0, 0x14, 0xfd, 0x62, 0x2d, 0xf5, 0x8b,
	0xf1, 0x7a, 0x91, 0x89, 0x1e, 0x40, 0xf5, 0x56, 0xd5, 0x53, 0x2d, 0x1b, 0xd8, 0x44, 0xde, 0x17,
	0x0b, 0x9c, 0x62, 0x83, 0xe7, 0xb0, 0x36, 0xcc, 0xe2, 0x7e, 0xa1, 0xba, 0x32, 0xf6, 0x1b, 0x83,
	0x15, 0xa8, 0xbd, 0x15, 0xbc, 0xa0, 0x4a, 0x59, 0x96, 0x72, 0xc1, 0x28, 0x99, 0xa8, 0x06, 0x46,
	0xf6, 0xc1, 0x60, 0xf7, 0x64, 0x39, 0xb5, 0x8b, 0x60, 0xa3, 0x30, 0x64, 0x20, 0xe6, 0x29, 0xf5,
	0x02, 0xd8, 0x5a, 0xd2, 0x0d, 0x1d, 0x40, 0x23, 0x97, 0x05, 0x85, 0x6f, 0x51, 0x3d, 0x07, 0xdf,
	0xc9, 0x6f, 0xd2, 0x01, 0x34, 0xf2, 0x91, 0x34, 0x49, 0xdf, 0xb7, 0x9e, 0x83, 0x92, 0xe4, 0x3d,
	0x86, 0xad, 0x25, 0x73, 0x2d, 0xfb, 0xc6, 0x79, 0x4f, 0xc0, 0x5e, 0x7c, 0x90, 0xd0, 0x43, 0x00,
	0x4e, 0xfb, 0x8c, 0x8a, 0x80, 0xd1, 0xa1, 0xa1, 0xd9, 0x1a, 0xc1, 0x74, 0xd8, 0xad, 0x7c, 0xfb,
	0xf5, 0xc8, 0xba, 0xae, 0xaa, 0x3f, 0xd5, 0xb3, 0x3f, 0x01, 0x00, 0x00, 0xff, 0xff, 0x39, 0xa4,
	0x52, 0xd5, 0xab, 0x05, 0x00, 0x00,
}