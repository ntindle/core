package data

import (
	"fmt"
	"net/http"

	olc "github.com/google/open-location-code/go"
	"github.com/libp2p/go-libp2p-core/peer"
	"github.com/libp2p/go-libp2p-core/protocol"
	util "github.com/sonr-io/core/pkg/util"
	"google.golang.org/protobuf/proto"
)

// ** ─── ConnectionRequest MANAGEMENT ────────────────────────────────────────────────────────
// Method Returns Formatted Info for Address
func (ip *ConnectionRequest_IPAddress) Info() string {
	return fmt.Sprintf("%s \n \t(%s) \n \t\tAddress: %s \n \t\tMac: %s \n \t\tIs Internal: %t \n", ip.GetName(), ip.GetFamily().String(), ip.GetValue(), ip.GetMac(), ip.GetInternal())
}

// Method Checks if this Address is IPv4 and not a Loopback
func (ip *ConnectionRequest_IPAddress) IsIPv4() bool {
	return ip.GetFamily() == ConnectionRequest_IPAddress_IPV4 && !ip.GetInternal()
}

// Method Checks if this Address is IPv6 and not a Loopback
func (ip *ConnectionRequest_IPAddress) IsIPv6() bool {
	return ip.GetFamily() == ConnectionRequest_IPAddress_IPV6 && !ip.GetInternal()
}

// Method Converts this IPAddress into a MultiAddress String
func (ip *ConnectionRequest_IPAddress) MultiAddrStr(onlyIPv4 bool, port int) (string, error) {
	if ip.IsIPv4() {
		LogInfo(ip.Info())
		return fmt.Sprintf("/ip4/%s/tcp/%d", ip.GetValue(), port), nil
	} else if ip.IsIPv6() && !onlyIPv4 {
		LogInfo(ip.Info())
		return fmt.Sprintf("/ip6/%s/tcp/%d", ip.GetValue(), port), nil
	} else {
		return "", fmt.Errorf("Invalid IP Address")
	}

}

// ** ─── VerifyRequest MANAGEMENT ────────────────────────────────────────────────────────
// Checks if VerifyRequest is for String Value
func (vr *VerifyRequest) IsString() bool {
	switch vr.Data.(type) {
	case *VerifyRequest_TextValue:
		return true
	default:
		return false
	}
}

// Checks if VerifyRequest is for Buffer Value
func (vr *VerifyRequest) IsBuffer() bool {
	switch vr.Data.(type) {
	case *VerifyRequest_BufferValue:
		return true
	default:
		return false
	}
}

// ** ─── VerifyResponse MANAGEMENT ────────────────────────────────────────────────────────
// Create  VerifyResponse as GIVEN VALUE
func NewVerifyResponse(result bool) *VerifyResponse {
	return &VerifyResponse{Success: result}
}

// Create Marshalled VerifyResponse as GIVEN VALUE
func NewVerifyResponseBuf(result bool) []byte {
	if buf, err := proto.Marshal(&VerifyResponse{Success: result}); err != nil {
		return nil
	} else {
		return buf
	}
}

// Create Marshalled VerifyResponse as TRUE
func NewValidVerifyResponseBuf() []byte {
	if buf, err := proto.Marshal(&VerifyResponse{Success: true}); err != nil {
		return nil
	} else {
		return buf
	}
}

// Create Marshalled VerifyResponse as FALSE
func NewInvalidVerifyResponseBuf() []byte {
	if buf, err := proto.Marshal(&VerifyResponse{Success: false}); err != nil {
		return nil
	} else {
		return buf
	}
}

// ** ─── InitializeRequest MANAGEMENT ────────────────────────────────────────────────────────
// Check for Logging Enabled
func (i *InitializeRequest) IsLoggingEnabled() bool {
	if i.GetLogLevel() != InitializeRequest_NONE {
		return true
	}
	return false
}

// Check if Info Log Enabled
func (i *InitializeRequest) HasInfoLog() bool {
	if i.GetLogLevel() >= InitializeRequest_INFO {
		return true
	}
	return false
}

// Check if Debug Log Enabled
func (i *InitializeRequest) HasDebugLog() bool {
	if i.GetLogLevel() >= InitializeRequest_DEBUG {
		return true
	}
	return false
}

// Check if Debug Log Enabled
func (i *InitializeRequest) HasWarningLog() bool {
	if i.GetLogLevel() >= InitializeRequest_WARNING {
		return true
	}
	return false
}

// Check if Debug Log Enabled
func (i *InitializeRequest) HasCriticalLog() bool {
	if i.GetLogLevel() >= InitializeRequest_CRITICAL {
		return true
	}
	return false
}

// Check if Debug Log Enabled
func (i *InitializeRequest) HasFatalLog() bool {
	if i.GetLogLevel() >= InitializeRequest_FATAL {
		return true
	}
	return false
}

// Check if Request requires temporary keys
func (i *InitializeRequest) ShouldCreateTempKeys() bool {
	return i.GetClient() == InitializeRequest_RPC && i.GetOptions().GetIsPairing() && i.GetDevice().GetPlatform().IsMobile()
}

// Check if Request is Existing User and Verified
func (i *InitializeRequest) ShouldLoadKeychain() bool {
	return i.GetOptions().GetIsVerified()
}

// Check if Request is to Reset all keys
func (i *InitializeRequest) ShouldResetKeys() bool {
	return i.GetOptions().GetResetKeys()
}

// Get Assumed User State
func (i *InitializeRequest) UserState() User_State {
	if i.ShouldCreateTempKeys() {
		return User_UNVERIFIED
	} else if i.ShouldLoadKeychain() {
		return User_VERIFIED
	} else {
		return User_UNVERIFIED
	}
}

// ** ─── URLLink MANAGEMENT ────────────────────────────────────────────────────────
// Creates New Link
func NewURLLink(url string) *URLLink {
	link := &URLLink{
		Url:         url,
		Initialized: false,
	}
	link.SetData()
	return link
}

// Sets URLLink Data
func (u *URLLink) SetData() {
	if !u.Initialized {
		// Create Request
		resp, err := http.Get(u.Url)
		if err != nil {
			return
		}

		// Get Info
		info, err := util.GetPageData(resp)
		if err != nil {
			return
		}

		// Set Link
		u.Initialized = true
		u.Title = info.Title
		u.Type = info.Type
		u.Site = info.Site
		u.SiteName = info.SiteName
		u.Description = info.Description
		u.Locale = info.Locale

		// Get Images
		if info.Images != nil {
			for _, v := range info.Images {
				u.Images = append(u.Images, &URLLink_OpenGraphImage{
					Url:       v.Url,
					SecureUrl: v.SecureUrl,
					Width:     int32(v.Width),
					Height:    int32(v.Height),
					Type:      v.Type,
				})
			}
		}

		// Get Videos
		if info.Videos != nil {
			for _, v := range info.Videos {
				u.Videos = append(u.Videos, &URLLink_OpenGraphVideo{
					Url:       v.Url,
					SecureUrl: v.SecureUrl,
					Width:     int32(v.Width),
					Height:    int32(v.Height),
					Type:      v.Type,
				})
			}
		}

		// Get Audios
		if info.Audios != nil {
			for _, v := range info.Videos {
				u.Audios = append(u.Audios, &URLLink_OpenGraphAudio{
					Url:       v.Url,
					SecureUrl: v.SecureUrl,
					Type:      v.Type,
				})
			}
		}

		// Get Twitter
		if info.Twitter != nil {
			u.Twitter = &URLLink_TwitterCard{
				Card:        info.Twitter.Card,
				Site:        info.Twitter.Site,
				SiteId:      info.Twitter.SiteId,
				Creator:     info.Twitter.Creator,
				CreatorId:   info.Twitter.CreatorId,
				Description: info.Twitter.Description,
				Title:       info.Twitter.Title,
				Image:       info.Twitter.Image,
				ImageAlt:    info.Twitter.ImageAlt,
				Url:         info.Twitter.Url,
				Player: &URLLink_TwitterCard_Player{
					Url:    info.Twitter.Player.Url,
					Width:  int32(info.Twitter.Player.Width),
					Height: int32(info.Twitter.Player.Height),
					Stream: info.Twitter.Player.Stream,
				},
				Iphone: &URLLink_TwitterCard_IPhone{
					Name: info.Twitter.IPhone.Name,
					Id:   info.Twitter.IPhone.Id,
					Url:  info.Twitter.IPhone.Url,
				},
				Ipad: &URLLink_TwitterCard_IPad{
					Name: info.Twitter.IPad.Name,
					Id:   info.Twitter.IPad.Id,
					Url:  info.Twitter.IPad.Url,
				},
				GooglePlay: &URLLink_TwitterCard_GooglePlay{
					Name: info.Twitter.Googleplay.Name,
					Id:   info.Twitter.Googleplay.Id,
					Url:  info.Twitter.Googleplay.Url,
				},
			}
		}
	}
}

// ** ─── DecisionRequest MANAGEMENT ────────────────────────────────────────────────────────
// The Response of DecisionRequest is ACCEPT
func (d Decision) Accepted() bool {
	return d == Decision_ACCEPT
}

// The Response of DecisionRequest is DECLINE
func (d Decision) Declined() bool {
	return d == Decision_DECLINE
}

// The Response of DecisionRequest is IGNORE
func (d Decision) Ignored() bool {
	return d == Decision_IGNORE
}

// The Response of DecisionRequest is CANCEL
func (d Decision) Cancelled() bool {
	return d == Decision_CANCEL
}

// Convert this DecisionRequest InviteResponse
func (dr *DecisionRequest) ToResponse() *InviteResponse {
	return &InviteResponse{
		Decision: dr.Decision,
		To:       dr.GetTo(),
		From:     dr.GetFrom(),
		Transfer: dr.GetTransfer(),
		Type:     InviteResponse_Type(dr.GetType()),
		Payload:  dr.GetPayload(),
		Protocol: dr.GetProtocol(),
	}
}

// ** ─── LinkRequest MANAGEMENT ────────────────────────────────────────────────────────
// Checks if LinkRequest is SEND type
func (r *LinkRequest) IsSend() bool {
	return r.Type == LinkRequest_SEND
}

// Checks if LinkRequest is RECEIVE type
func (r *LinkRequest) IsReceive() bool {
	return r.Type == LinkRequest_RECEIVE
}

// Checks if LinkRequest is CANCEL type
func (r *LinkRequest) IsCancel() bool {
	return r.Type == LinkRequest_CANCEL
}

// ** ─── InviteResponse MANAGEMENT ────────────────────────────────────────────────────────
// Checks if Peer Accepted Transfer
func (r *InviteResponse) HasAcceptedTransfer() bool {
	return r.GetDecision().Accepted()
}

// Returns Protocol ID Set by Peer
func (r *InviteResponse) ProtocolID() protocol.ID {
	return protocol.ID(r.GetProtocol())
}

// ** ─── InviteRequest MANAGEMENT ────────────────────────────────────────────────────────
// Returns Invite Contact
func (i *InviteRequest) GetContact() *Contact {
	return i.GetTransfer().GetContact()
}

// Returns Invite File
func (i *InviteRequest) GetFile() *SFile {
	return i.GetTransfer().GetFile()
}

// Returns Invite URL
func (i *InviteRequest) GetUrl() *URLLink {
	return i.GetTransfer().GetUrl()
}

// Checks if Payload is Contact
func (i *InviteRequest) IsPayloadContact() bool {
	return i.Payload == Payload_CONTACT
}

// Checks if Payload is File Transfer
func (i *InviteRequest) IsPayloadTransfer() bool {
	return i.Payload == Payload_FILE || i.Payload == Payload_FILES || i.Payload == Payload_MEDIA || i.Payload == Payload_ALBUM
}

// Checks if Payload is Url
func (i *InviteRequest) IsPayloadUrl() bool {
	return i.Payload == Payload_URL
}

// Checks for Flat Invite
func (i *InviteRequest) IsDirectInvite() bool {
	return i.GetType() == InviteRequest_DIRECT
}

// Returns Protocol ID Set by Peer
func (r *InviteRequest) ProtocolID() protocol.ID {
	return protocol.ID(r.GetProtocol())
}

// Set Protocol for Invite and Return ID
func (i *InviteRequest) SetProtocol(p SonrProtocol, id peer.ID) protocol.ID {
	// Set Name and Return ID
	if i.Protocol != nil {
		p := string(p.NewIDProtocol(id))
		i.Protocol = &p
	}
	return p.NewIDProtocol(id)
}

// Validates LinkRequest has From Parameter
func (u *Device) SignLink(i *LinkRequest) *LinkRequest {
	// Set From
	if i.From == nil {
		i.From = u.GetPeer()
	}

	// Set Contact
	if i.Contact == nil {
		i.Contact = u.GetContact()
	}
	return i
}

// ** ─── Location MANAGEMENT ────────────────────────────────────────────────────────
func (l *Location) OLC(scope int) string {
	return olc.Encode(float64(l.GetLatitude()), float64(l.GetLongitude()), scope)
}

// ** ─── Router MANAGEMENT ────────────────────────────────────────────────────────
// Construct New Protocol ID given Method Name String and id Peer.ID
func (p SonrProtocol) NewIDProtocol(id peer.ID) protocol.ID {
	return protocol.ID(fmt.Sprintf("/sonr/%s/%s", p.Method(), id.String()))
}

// Construct New Protocol ID given Method Name String and Value String
func (p SonrProtocol) NewValueProtocol(value string) protocol.ID {
	return protocol.ID(fmt.Sprintf("/sonr/%s/%s", p.Method(), value))
}

// Returns Method Name for this Protocol
func (p SonrProtocol) Method() string {
	switch p {
	case SonrProtocol_Authorize:
		return "auth-service"
	case SonrProtocol_Devices:
		return "user-devices"
	case SonrProtocol_Linker:
		return "user-linker"
	case SonrProtocol_LocalTransfer:
		return "local-transfer"
	case SonrProtocol_RemoteTransfer:
		return "remote-transfer"
	default:
		return ""
	}
}

// ** ─── Status MANAGEMENT ────────────────────────────────────────────────────────
var (
	defaultStatus = Status_DEFAULT
	available     = Status_AVAILABLE
	connected     = Status_CONNECTED
	failed        = Status_FAILED
)

// Update Connected Connection Status
func (u *Device) SetConnected(value bool) *StatusEvent {
	// Update Status
	if value {
		u.Status = &available
	} else {
		u.Status = &failed
	}

	// Returns Status Update
	return &StatusEvent{Value: *u.Status}
}

// Update Bootstrap Connection Status
func (u *Device) SetAvailable(value bool) *StatusEvent {
	// Update Status
	if value {
		u.Status = &available
	} else {
		u.Status = &failed
	}

	// Returns Status Update
	return &StatusEvent{Value: *u.Status}
}

// Update Node Status
func (u *Device) SetStatus(ns Status) *StatusEvent {
	// Set Value
	u.Status = &ns

	// Returns Status Update
	return &StatusEvent{Value: *u.Status}
}

// Checks if Status is Given Value
func (u *Device) IsStatus(gs Status) bool {
	return u.Status == &gs
}

// Checks if Status is Not Given Value
func (u *Device) IsNotStatus(gs Status) bool {
	return u.Status != &gs
}

// ** ─── Generic Callback MANAGEMENT ───────────────────────────────────────────

// Create New Generic CONNECTION Response Message
func (o *ConnectionResponse) ToGeneric() ([]byte, error) {
	// Marshal Original
	ogBuf, err := proto.Marshal(o)
	if err != nil {
		return nil, err
	}

	// Create Generic
	generic := &GenericResponse{
		Type: GenericResponse_CONNECTION,
		Data: ogBuf,
	}

	// Marshal Generic
	genBuf, err := proto.Marshal(generic)
	if err != nil {
		return nil, err
	}
	return genBuf, nil
}

// Create New Generic INVITE Request Message
func (o *InviteRequest) ToGeneric() ([]byte, error) {
	// Marshal Original
	ogBuf, err := proto.Marshal(o)
	if err != nil {
		return nil, err
	}

	// Create Generic
	generic := &GenericRequest{
		Type: GenericRequest_INVITE,
		Data: ogBuf,
	}

	// Marshal Generic
	genBuf, err := proto.Marshal(generic)
	if err != nil {
		return nil, err
	}
	return genBuf, nil
}

// Create New Generic REPLY Response Message
func (o *InviteResponse) ToGeneric() ([]byte, error) {
	// Marshal Original
	ogBuf, err := proto.Marshal(o)
	if err != nil {
		return nil, err
	}

	// Create Generic
	generic := &GenericResponse{
		Type: GenericResponse_REPLY,
		Data: ogBuf,
	}

	// Marshal Generic
	genBuf, err := proto.Marshal(generic)
	if err != nil {
		return nil, err
	}
	return genBuf, nil
}

// Create New Generic COMPLETE Event Message
func (o *CompleteEvent) ToGeneric() ([]byte, error) {
	// Marshal Original
	ogBuf, err := proto.Marshal(o)
	if err != nil {
		return nil, err
	}

	// Create Generic
	generic := &GenericEvent{
		Type: GenericEvent_COMPLETE,
		Data: ogBuf,
	}

	// Marshal Generic
	genBuf, err := proto.Marshal(generic)
	if err != nil {
		return nil, err
	}
	return genBuf, nil
}

// Create New Generic MAIL Event Message
func (o *MailEvent) ToGeneric() ([]byte, error) {
	// Marshal Original
	ogBuf, err := proto.Marshal(o)
	if err != nil {
		return nil, err
	}

	// Create Generic
	generic := &GenericEvent{
		Type: GenericEvent_MAIL,
		Data: ogBuf,
	}

	// Marshal Generic
	genBuf, err := proto.Marshal(generic)
	if err != nil {
		return nil, err
	}
	return genBuf, nil
}

// Create New Generic PROGRESS Event Message
func (o *ProgressEvent) ToGeneric() ([]byte, error) {
	// Marshal Original
	ogBuf, err := proto.Marshal(o)
	if err != nil {
		return nil, err
	}

	// Create Generic
	generic := &GenericEvent{
		Type: GenericEvent_PROGRESS,
		Data: ogBuf,
	}

	// Marshal Generic
	genBuf, err := proto.Marshal(generic)
	if err != nil {
		return nil, err
	}
	return genBuf, nil
}

// Create New Generic LINK Event Message
func (o *LinkEvent) ToGeneric() ([]byte, error) {
	// Marshal Original
	ogBuf, err := proto.Marshal(o)
	if err != nil {
		return nil, err
	}

	// Create Generic
	generic := &GenericEvent{
		Type: GenericEvent_LINK,
		Data: ogBuf,
	}

	// Marshal Generic
	genBuf, err := proto.Marshal(generic)
	if err != nil {
		return nil, err
	}
	return genBuf, nil
}

// Create New Generic Room Event Message
func (o *RoomEvent) ToGeneric() ([]byte, error) {
	// Marshal Original
	ogBuf, err := proto.Marshal(o)
	if err != nil {
		return nil, err
	}

	// Create Generic
	generic := &GenericEvent{
		Type: GenericEvent_ROOM,
		Data: ogBuf,
	}

	// Marshal Generic
	genBuf, err := proto.Marshal(generic)
	if err != nil {
		return nil, err
	}
	return genBuf, nil
}

// Create New Generic LINK Event Message
func (o *SyncEvent) ToGeneric() ([]byte, error) {
	// Marshal Original
	ogBuf, err := proto.Marshal(o)
	if err != nil {
		return nil, err
	}

	// Create Generic
	generic := &GenericEvent{
		Type: GenericEvent_SYNC,
		Data: ogBuf,
	}

	// Marshal Generic
	genBuf, err := proto.Marshal(generic)
	if err != nil {
		return nil, err
	}
	return genBuf, nil
}
