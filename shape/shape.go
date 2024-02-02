// Code generated by xgbgen from shape.xml. DO NOT EDIT.

// Package shape is the X client API for the SHAPE extension.
package shape

import (
	"github.com/jezek/xgb"

	"github.com/jezek/xgb/xproto"
)

// Init must be called before using the SHAPE extension.
func Init(c *xgb.Conn) error {
	reply, err := xproto.QueryExtension(c, 5, "SHAPE").Reply()
	switch {
	case err != nil:
		return err
	case !reply.Present:
		return xgb.Errorf("No extension named SHAPE could be found on on the server.")
	}

	c.ExtLock.Lock()
	c.Extensions["SHAPE"] = reply.MajorOpcode
	c.ExtLock.Unlock()
	for evNum, fun := range xgb.NewExtEventFuncs["SHAPE"] {
		xgb.NewEventFuncs[int(reply.FirstEvent)+evNum] = fun
	}
	for errNum, fun := range xgb.NewExtErrorFuncs["SHAPE"] {
		xgb.NewErrorFuncs[int(reply.FirstError)+errNum] = fun
	}
	return nil
}

func init() {
	xgb.NewExtEventFuncs["SHAPE"] = make(map[int]xgb.NewEventFun)
	xgb.NewExtErrorFuncs["SHAPE"] = make(map[int]xgb.NewErrorFun)
}

type Kind byte

// Notify is the event number for a NotifyEvent.
const Notify = 0

type NotifyEvent struct {
	Sequence       uint16
	ShapeKind      Kind
	AffectedWindow xproto.Window
	ExtentsX       int16
	ExtentsY       int16
	ExtentsWidth   uint16
	ExtentsHeight  uint16
	ServerTime     xproto.Timestamp
	Shaped         bool
	// padding: 11 bytes
}

// NotifyEventNew constructs a NotifyEvent value that implements xgb.Event from a byte slice.
func NotifyEventNew(buf []byte) xgb.Event {
	v := NotifyEvent{}
	b := 1 // don't read event number

	v.ShapeKind = Kind(buf[b])
	b += 1

	v.Sequence = xgb.Get16(buf[b:])
	b += 2

	v.AffectedWindow = xproto.Window(xgb.Get32(buf[b:]))
	b += 4

	v.ExtentsX = int16(xgb.Get16(buf[b:]))
	b += 2

	v.ExtentsY = int16(xgb.Get16(buf[b:]))
	b += 2

	v.ExtentsWidth = xgb.Get16(buf[b:])
	b += 2

	v.ExtentsHeight = xgb.Get16(buf[b:])
	b += 2

	v.ServerTime = xproto.Timestamp(xgb.Get32(buf[b:]))
	b += 4

	if buf[b] == 1 {
		v.Shaped = true
	} else {
		v.Shaped = false
	}
	b += 1

	b += 11 // padding

	return v
}

// Bytes writes a NotifyEvent value to a byte slice.
func (v NotifyEvent) Bytes() []byte {
	buf := make([]byte, 32)
	b := 0

	// write event number
	buf[b] = 0
	b += 1

	buf[b] = byte(v.ShapeKind)
	b += 1

	b += 2 // skip sequence number

	xgb.Put32(buf[b:], uint32(v.AffectedWindow))
	b += 4

	xgb.Put16(buf[b:], uint16(v.ExtentsX))
	b += 2

	xgb.Put16(buf[b:], uint16(v.ExtentsY))
	b += 2

	xgb.Put16(buf[b:], v.ExtentsWidth)
	b += 2

	xgb.Put16(buf[b:], v.ExtentsHeight)
	b += 2

	xgb.Put32(buf[b:], uint32(v.ServerTime))
	b += 4

	if v.Shaped {
		buf[b] = 1
	} else {
		buf[b] = 0
	}
	b += 1

	b += 11 // padding

	return buf
}

// SequenceId returns the sequence id attached to the Notify event.
// Events without a sequence number (KeymapNotify) return 0.
// This is mostly used internally.
func (v NotifyEvent) SequenceId() uint16 {
	return v.Sequence
}

// String is a rudimentary string representation of NotifyEvent.
func (v NotifyEvent) String() string {
	fieldVals := make([]string, 0, 9)
	fieldVals = append(fieldVals, xgb.Sprintf("Sequence: %d", v.Sequence))
	fieldVals = append(fieldVals, xgb.Sprintf("ShapeKind: %d", v.ShapeKind))
	fieldVals = append(fieldVals, xgb.Sprintf("AffectedWindow: %d", v.AffectedWindow))
	fieldVals = append(fieldVals, xgb.Sprintf("ExtentsX: %d", v.ExtentsX))
	fieldVals = append(fieldVals, xgb.Sprintf("ExtentsY: %d", v.ExtentsY))
	fieldVals = append(fieldVals, xgb.Sprintf("ExtentsWidth: %d", v.ExtentsWidth))
	fieldVals = append(fieldVals, xgb.Sprintf("ExtentsHeight: %d", v.ExtentsHeight))
	fieldVals = append(fieldVals, xgb.Sprintf("ServerTime: %d", v.ServerTime))
	fieldVals = append(fieldVals, xgb.Sprintf("Shaped: %t", v.Shaped))
	return "Notify {" + xgb.StringsJoin(fieldVals, ", ") + "}"
}

func init() {
	xgb.NewExtEventFuncs["SHAPE"][0] = NotifyEventNew
}

type Op byte

const (
	SkBounding = 0
	SkClip     = 1
	SkInput    = 2
)

const (
	SoSet       = 0
	SoUnion     = 1
	SoIntersect = 2
	SoSubtract  = 3
	SoInvert    = 4
)

// Skipping definition for base type 'Bool'

// Skipping definition for base type 'Byte'

// Skipping definition for base type 'Card8'

// Skipping definition for base type 'Char'

// Skipping definition for base type 'Void'

// Skipping definition for base type 'Double'

// Skipping definition for base type 'Float'

// Skipping definition for base type 'Int16'

// Skipping definition for base type 'Int32'

// Skipping definition for base type 'Int8'

// Skipping definition for base type 'Card16'

// Skipping definition for base type 'Card32'

// CombineCookie is a cookie used only for Combine requests.
type CombineCookie struct {
	*xgb.Cookie
}

// Combine sends an unchecked request.
// If an error occurs, it can only be retrieved using xgb.WaitForEvent or xgb.PollForEvent.
func Combine(c *xgb.Conn, Operation Op, DestinationKind Kind, SourceKind Kind, DestinationWindow xproto.Window, XOffset int16, YOffset int16, SourceWindow xproto.Window) CombineCookie {
	c.ExtLock.RLock()
	defer c.ExtLock.RUnlock()
	if _, ok := c.Extensions["SHAPE"]; !ok {
		panic("Cannot issue request 'Combine' using the uninitialized extension 'SHAPE'. shape.Init(connObj) must be called first.")
	}
	cookie := c.NewCookie(false, false)
	c.NewRequest(combineRequest(c, Operation, DestinationKind, SourceKind, DestinationWindow, XOffset, YOffset, SourceWindow), cookie)
	return CombineCookie{cookie}
}

// CombineChecked sends a checked request.
// If an error occurs, it can be retrieved using CombineCookie.Check()
func CombineChecked(c *xgb.Conn, Operation Op, DestinationKind Kind, SourceKind Kind, DestinationWindow xproto.Window, XOffset int16, YOffset int16, SourceWindow xproto.Window) CombineCookie {
	c.ExtLock.RLock()
	defer c.ExtLock.RUnlock()
	if _, ok := c.Extensions["SHAPE"]; !ok {
		panic("Cannot issue request 'Combine' using the uninitialized extension 'SHAPE'. shape.Init(connObj) must be called first.")
	}
	cookie := c.NewCookie(true, false)
	c.NewRequest(combineRequest(c, Operation, DestinationKind, SourceKind, DestinationWindow, XOffset, YOffset, SourceWindow), cookie)
	return CombineCookie{cookie}
}

// Check returns an error if one occurred for checked requests that are not expecting a reply.
// This cannot be called for requests expecting a reply, nor for unchecked requests.
func (cook CombineCookie) Check() error {
	return cook.Cookie.Check()
}

// Write request to wire for Combine
// combineRequest writes a Combine request to a byte slice.
func combineRequest(c *xgb.Conn, Operation Op, DestinationKind Kind, SourceKind Kind, DestinationWindow xproto.Window, XOffset int16, YOffset int16, SourceWindow xproto.Window) []byte {
	size := 20
	b := 0
	buf := make([]byte, size)

	c.ExtLock.RLock()
	buf[b] = c.Extensions["SHAPE"]
	c.ExtLock.RUnlock()
	b += 1

	buf[b] = 3 // request opcode
	b += 1

	xgb.Put16(buf[b:], uint16(size/4)) // write request size in 4-byte units
	b += 2

	buf[b] = byte(Operation)
	b += 1

	buf[b] = byte(DestinationKind)
	b += 1

	buf[b] = byte(SourceKind)
	b += 1

	b += 1 // padding

	xgb.Put32(buf[b:], uint32(DestinationWindow))
	b += 4

	xgb.Put16(buf[b:], uint16(XOffset))
	b += 2

	xgb.Put16(buf[b:], uint16(YOffset))
	b += 2

	xgb.Put32(buf[b:], uint32(SourceWindow))
	b += 4

	return buf
}

// GetRectanglesCookie is a cookie used only for GetRectangles requests.
type GetRectanglesCookie struct {
	*xgb.Cookie
}

// GetRectangles sends a checked request.
// If an error occurs, it will be returned with the reply by calling GetRectanglesCookie.Reply()
func GetRectangles(c *xgb.Conn, Window xproto.Window, SourceKind Kind) GetRectanglesCookie {
	c.ExtLock.RLock()
	defer c.ExtLock.RUnlock()
	if _, ok := c.Extensions["SHAPE"]; !ok {
		panic("Cannot issue request 'GetRectangles' using the uninitialized extension 'SHAPE'. shape.Init(connObj) must be called first.")
	}
	cookie := c.NewCookie(true, true)
	c.NewRequest(getRectanglesRequest(c, Window, SourceKind), cookie)
	return GetRectanglesCookie{cookie}
}

// GetRectanglesUnchecked sends an unchecked request.
// If an error occurs, it can only be retrieved using xgb.WaitForEvent or xgb.PollForEvent.
func GetRectanglesUnchecked(c *xgb.Conn, Window xproto.Window, SourceKind Kind) GetRectanglesCookie {
	c.ExtLock.RLock()
	defer c.ExtLock.RUnlock()
	if _, ok := c.Extensions["SHAPE"]; !ok {
		panic("Cannot issue request 'GetRectangles' using the uninitialized extension 'SHAPE'. shape.Init(connObj) must be called first.")
	}
	cookie := c.NewCookie(false, true)
	c.NewRequest(getRectanglesRequest(c, Window, SourceKind), cookie)
	return GetRectanglesCookie{cookie}
}

// GetRectanglesReply represents the data returned from a GetRectangles request.
type GetRectanglesReply struct {
	Sequence      uint16 // sequence number of the request for this reply
	Length        uint32 // number of bytes in this reply
	Ordering      byte
	RectanglesLen uint32
	// padding: 20 bytes
	Rectangles []xproto.Rectangle // size: xgb.Pad((int(RectanglesLen) * 8))
}

// Reply blocks and returns the reply data for a GetRectangles request.
func (cook GetRectanglesCookie) Reply() (*GetRectanglesReply, error) {
	buf, err := cook.Cookie.Reply()
	if err != nil {
		return nil, err
	}
	if buf == nil {
		return nil, nil
	}
	return getRectanglesReply(buf), nil
}

// getRectanglesReply reads a byte slice into a GetRectanglesReply value.
func getRectanglesReply(buf []byte) *GetRectanglesReply {
	v := new(GetRectanglesReply)
	b := 1 // skip reply determinant

	v.Ordering = buf[b]
	b += 1

	v.Sequence = xgb.Get16(buf[b:])
	b += 2

	v.Length = xgb.Get32(buf[b:]) // 4-byte units
	b += 4

	v.RectanglesLen = xgb.Get32(buf[b:])
	b += 4

	b += 20 // padding

	v.Rectangles = make([]xproto.Rectangle, v.RectanglesLen)
	b += xproto.RectangleReadList(buf[b:], v.Rectangles)

	return v
}

// Write request to wire for GetRectangles
// getRectanglesRequest writes a GetRectangles request to a byte slice.
func getRectanglesRequest(c *xgb.Conn, Window xproto.Window, SourceKind Kind) []byte {
	size := 12
	b := 0
	buf := make([]byte, size)

	c.ExtLock.RLock()
	buf[b] = c.Extensions["SHAPE"]
	c.ExtLock.RUnlock()
	b += 1

	buf[b] = 8 // request opcode
	b += 1

	xgb.Put16(buf[b:], uint16(size/4)) // write request size in 4-byte units
	b += 2

	xgb.Put32(buf[b:], uint32(Window))
	b += 4

	buf[b] = byte(SourceKind)
	b += 1

	b += 3 // padding

	return buf
}

// InputSelectedCookie is a cookie used only for InputSelected requests.
type InputSelectedCookie struct {
	*xgb.Cookie
}

// InputSelected sends a checked request.
// If an error occurs, it will be returned with the reply by calling InputSelectedCookie.Reply()
func InputSelected(c *xgb.Conn, DestinationWindow xproto.Window) InputSelectedCookie {
	c.ExtLock.RLock()
	defer c.ExtLock.RUnlock()
	if _, ok := c.Extensions["SHAPE"]; !ok {
		panic("Cannot issue request 'InputSelected' using the uninitialized extension 'SHAPE'. shape.Init(connObj) must be called first.")
	}
	cookie := c.NewCookie(true, true)
	c.NewRequest(inputSelectedRequest(c, DestinationWindow), cookie)
	return InputSelectedCookie{cookie}
}

// InputSelectedUnchecked sends an unchecked request.
// If an error occurs, it can only be retrieved using xgb.WaitForEvent or xgb.PollForEvent.
func InputSelectedUnchecked(c *xgb.Conn, DestinationWindow xproto.Window) InputSelectedCookie {
	c.ExtLock.RLock()
	defer c.ExtLock.RUnlock()
	if _, ok := c.Extensions["SHAPE"]; !ok {
		panic("Cannot issue request 'InputSelected' using the uninitialized extension 'SHAPE'. shape.Init(connObj) must be called first.")
	}
	cookie := c.NewCookie(false, true)
	c.NewRequest(inputSelectedRequest(c, DestinationWindow), cookie)
	return InputSelectedCookie{cookie}
}

// InputSelectedReply represents the data returned from a InputSelected request.
type InputSelectedReply struct {
	Sequence uint16 // sequence number of the request for this reply
	Length   uint32 // number of bytes in this reply
	Enabled  bool
}

// Reply blocks and returns the reply data for a InputSelected request.
func (cook InputSelectedCookie) Reply() (*InputSelectedReply, error) {
	buf, err := cook.Cookie.Reply()
	if err != nil {
		return nil, err
	}
	if buf == nil {
		return nil, nil
	}
	return inputSelectedReply(buf), nil
}

// inputSelectedReply reads a byte slice into a InputSelectedReply value.
func inputSelectedReply(buf []byte) *InputSelectedReply {
	v := new(InputSelectedReply)
	b := 1 // skip reply determinant

	if buf[b] == 1 {
		v.Enabled = true
	} else {
		v.Enabled = false
	}
	b += 1

	v.Sequence = xgb.Get16(buf[b:])
	b += 2

	v.Length = xgb.Get32(buf[b:]) // 4-byte units
	b += 4

	return v
}

// Write request to wire for InputSelected
// inputSelectedRequest writes a InputSelected request to a byte slice.
func inputSelectedRequest(c *xgb.Conn, DestinationWindow xproto.Window) []byte {
	size := 8
	b := 0
	buf := make([]byte, size)

	c.ExtLock.RLock()
	buf[b] = c.Extensions["SHAPE"]
	c.ExtLock.RUnlock()
	b += 1

	buf[b] = 7 // request opcode
	b += 1

	xgb.Put16(buf[b:], uint16(size/4)) // write request size in 4-byte units
	b += 2

	xgb.Put32(buf[b:], uint32(DestinationWindow))
	b += 4

	return buf
}

// MaskCookie is a cookie used only for Mask requests.
type MaskCookie struct {
	*xgb.Cookie
}

// Mask sends an unchecked request.
// If an error occurs, it can only be retrieved using xgb.WaitForEvent or xgb.PollForEvent.
func Mask(c *xgb.Conn, Operation Op, DestinationKind Kind, DestinationWindow xproto.Window, XOffset int16, YOffset int16, SourceBitmap xproto.Pixmap) MaskCookie {
	c.ExtLock.RLock()
	defer c.ExtLock.RUnlock()
	if _, ok := c.Extensions["SHAPE"]; !ok {
		panic("Cannot issue request 'Mask' using the uninitialized extension 'SHAPE'. shape.Init(connObj) must be called first.")
	}
	cookie := c.NewCookie(false, false)
	c.NewRequest(maskRequest(c, Operation, DestinationKind, DestinationWindow, XOffset, YOffset, SourceBitmap), cookie)
	return MaskCookie{cookie}
}

// MaskChecked sends a checked request.
// If an error occurs, it can be retrieved using MaskCookie.Check()
func MaskChecked(c *xgb.Conn, Operation Op, DestinationKind Kind, DestinationWindow xproto.Window, XOffset int16, YOffset int16, SourceBitmap xproto.Pixmap) MaskCookie {
	c.ExtLock.RLock()
	defer c.ExtLock.RUnlock()
	if _, ok := c.Extensions["SHAPE"]; !ok {
		panic("Cannot issue request 'Mask' using the uninitialized extension 'SHAPE'. shape.Init(connObj) must be called first.")
	}
	cookie := c.NewCookie(true, false)
	c.NewRequest(maskRequest(c, Operation, DestinationKind, DestinationWindow, XOffset, YOffset, SourceBitmap), cookie)
	return MaskCookie{cookie}
}

// Check returns an error if one occurred for checked requests that are not expecting a reply.
// This cannot be called for requests expecting a reply, nor for unchecked requests.
func (cook MaskCookie) Check() error {
	return cook.Cookie.Check()
}

// Write request to wire for Mask
// maskRequest writes a Mask request to a byte slice.
func maskRequest(c *xgb.Conn, Operation Op, DestinationKind Kind, DestinationWindow xproto.Window, XOffset int16, YOffset int16, SourceBitmap xproto.Pixmap) []byte {
	size := 20
	b := 0
	buf := make([]byte, size)

	c.ExtLock.RLock()
	buf[b] = c.Extensions["SHAPE"]
	c.ExtLock.RUnlock()
	b += 1

	buf[b] = 2 // request opcode
	b += 1

	xgb.Put16(buf[b:], uint16(size/4)) // write request size in 4-byte units
	b += 2

	buf[b] = byte(Operation)
	b += 1

	buf[b] = byte(DestinationKind)
	b += 1

	b += 2 // padding

	xgb.Put32(buf[b:], uint32(DestinationWindow))
	b += 4

	xgb.Put16(buf[b:], uint16(XOffset))
	b += 2

	xgb.Put16(buf[b:], uint16(YOffset))
	b += 2

	xgb.Put32(buf[b:], uint32(SourceBitmap))
	b += 4

	return buf
}

// OffsetCookie is a cookie used only for Offset requests.
type OffsetCookie struct {
	*xgb.Cookie
}

// Offset sends an unchecked request.
// If an error occurs, it can only be retrieved using xgb.WaitForEvent or xgb.PollForEvent.
func Offset(c *xgb.Conn, DestinationKind Kind, DestinationWindow xproto.Window, XOffset int16, YOffset int16) OffsetCookie {
	c.ExtLock.RLock()
	defer c.ExtLock.RUnlock()
	if _, ok := c.Extensions["SHAPE"]; !ok {
		panic("Cannot issue request 'Offset' using the uninitialized extension 'SHAPE'. shape.Init(connObj) must be called first.")
	}
	cookie := c.NewCookie(false, false)
	c.NewRequest(offsetRequest(c, DestinationKind, DestinationWindow, XOffset, YOffset), cookie)
	return OffsetCookie{cookie}
}

// OffsetChecked sends a checked request.
// If an error occurs, it can be retrieved using OffsetCookie.Check()
func OffsetChecked(c *xgb.Conn, DestinationKind Kind, DestinationWindow xproto.Window, XOffset int16, YOffset int16) OffsetCookie {
	c.ExtLock.RLock()
	defer c.ExtLock.RUnlock()
	if _, ok := c.Extensions["SHAPE"]; !ok {
		panic("Cannot issue request 'Offset' using the uninitialized extension 'SHAPE'. shape.Init(connObj) must be called first.")
	}
	cookie := c.NewCookie(true, false)
	c.NewRequest(offsetRequest(c, DestinationKind, DestinationWindow, XOffset, YOffset), cookie)
	return OffsetCookie{cookie}
}

// Check returns an error if one occurred for checked requests that are not expecting a reply.
// This cannot be called for requests expecting a reply, nor for unchecked requests.
func (cook OffsetCookie) Check() error {
	return cook.Cookie.Check()
}

// Write request to wire for Offset
// offsetRequest writes a Offset request to a byte slice.
func offsetRequest(c *xgb.Conn, DestinationKind Kind, DestinationWindow xproto.Window, XOffset int16, YOffset int16) []byte {
	size := 16
	b := 0
	buf := make([]byte, size)

	c.ExtLock.RLock()
	buf[b] = c.Extensions["SHAPE"]
	c.ExtLock.RUnlock()
	b += 1

	buf[b] = 4 // request opcode
	b += 1

	xgb.Put16(buf[b:], uint16(size/4)) // write request size in 4-byte units
	b += 2

	buf[b] = byte(DestinationKind)
	b += 1

	b += 3 // padding

	xgb.Put32(buf[b:], uint32(DestinationWindow))
	b += 4

	xgb.Put16(buf[b:], uint16(XOffset))
	b += 2

	xgb.Put16(buf[b:], uint16(YOffset))
	b += 2

	return buf
}

// QueryExtentsCookie is a cookie used only for QueryExtents requests.
type QueryExtentsCookie struct {
	*xgb.Cookie
}

// QueryExtents sends a checked request.
// If an error occurs, it will be returned with the reply by calling QueryExtentsCookie.Reply()
func QueryExtents(c *xgb.Conn, DestinationWindow xproto.Window) QueryExtentsCookie {
	c.ExtLock.RLock()
	defer c.ExtLock.RUnlock()
	if _, ok := c.Extensions["SHAPE"]; !ok {
		panic("Cannot issue request 'QueryExtents' using the uninitialized extension 'SHAPE'. shape.Init(connObj) must be called first.")
	}
	cookie := c.NewCookie(true, true)
	c.NewRequest(queryExtentsRequest(c, DestinationWindow), cookie)
	return QueryExtentsCookie{cookie}
}

// QueryExtentsUnchecked sends an unchecked request.
// If an error occurs, it can only be retrieved using xgb.WaitForEvent or xgb.PollForEvent.
func QueryExtentsUnchecked(c *xgb.Conn, DestinationWindow xproto.Window) QueryExtentsCookie {
	c.ExtLock.RLock()
	defer c.ExtLock.RUnlock()
	if _, ok := c.Extensions["SHAPE"]; !ok {
		panic("Cannot issue request 'QueryExtents' using the uninitialized extension 'SHAPE'. shape.Init(connObj) must be called first.")
	}
	cookie := c.NewCookie(false, true)
	c.NewRequest(queryExtentsRequest(c, DestinationWindow), cookie)
	return QueryExtentsCookie{cookie}
}

// QueryExtentsReply represents the data returned from a QueryExtents request.
type QueryExtentsReply struct {
	Sequence uint16 // sequence number of the request for this reply
	Length   uint32 // number of bytes in this reply
	// padding: 1 bytes
	BoundingShaped bool
	ClipShaped     bool
	// padding: 2 bytes
	BoundingShapeExtentsX      int16
	BoundingShapeExtentsY      int16
	BoundingShapeExtentsWidth  uint16
	BoundingShapeExtentsHeight uint16
	ClipShapeExtentsX          int16
	ClipShapeExtentsY          int16
	ClipShapeExtentsWidth      uint16
	ClipShapeExtentsHeight     uint16
}

// Reply blocks and returns the reply data for a QueryExtents request.
func (cook QueryExtentsCookie) Reply() (*QueryExtentsReply, error) {
	buf, err := cook.Cookie.Reply()
	if err != nil {
		return nil, err
	}
	if buf == nil {
		return nil, nil
	}
	return queryExtentsReply(buf), nil
}

// queryExtentsReply reads a byte slice into a QueryExtentsReply value.
func queryExtentsReply(buf []byte) *QueryExtentsReply {
	v := new(QueryExtentsReply)
	b := 1 // skip reply determinant

	b += 1 // padding

	v.Sequence = xgb.Get16(buf[b:])
	b += 2

	v.Length = xgb.Get32(buf[b:]) // 4-byte units
	b += 4

	if buf[b] == 1 {
		v.BoundingShaped = true
	} else {
		v.BoundingShaped = false
	}
	b += 1

	if buf[b] == 1 {
		v.ClipShaped = true
	} else {
		v.ClipShaped = false
	}
	b += 1

	b += 2 // padding

	v.BoundingShapeExtentsX = int16(xgb.Get16(buf[b:]))
	b += 2

	v.BoundingShapeExtentsY = int16(xgb.Get16(buf[b:]))
	b += 2

	v.BoundingShapeExtentsWidth = xgb.Get16(buf[b:])
	b += 2

	v.BoundingShapeExtentsHeight = xgb.Get16(buf[b:])
	b += 2

	v.ClipShapeExtentsX = int16(xgb.Get16(buf[b:]))
	b += 2

	v.ClipShapeExtentsY = int16(xgb.Get16(buf[b:]))
	b += 2

	v.ClipShapeExtentsWidth = xgb.Get16(buf[b:])
	b += 2

	v.ClipShapeExtentsHeight = xgb.Get16(buf[b:])
	b += 2

	return v
}

// Write request to wire for QueryExtents
// queryExtentsRequest writes a QueryExtents request to a byte slice.
func queryExtentsRequest(c *xgb.Conn, DestinationWindow xproto.Window) []byte {
	size := 8
	b := 0
	buf := make([]byte, size)

	c.ExtLock.RLock()
	buf[b] = c.Extensions["SHAPE"]
	c.ExtLock.RUnlock()
	b += 1

	buf[b] = 5 // request opcode
	b += 1

	xgb.Put16(buf[b:], uint16(size/4)) // write request size in 4-byte units
	b += 2

	xgb.Put32(buf[b:], uint32(DestinationWindow))
	b += 4

	return buf
}

// QueryVersionCookie is a cookie used only for QueryVersion requests.
type QueryVersionCookie struct {
	*xgb.Cookie
}

// QueryVersion sends a checked request.
// If an error occurs, it will be returned with the reply by calling QueryVersionCookie.Reply()
func QueryVersion(c *xgb.Conn) QueryVersionCookie {
	c.ExtLock.RLock()
	defer c.ExtLock.RUnlock()
	if _, ok := c.Extensions["SHAPE"]; !ok {
		panic("Cannot issue request 'QueryVersion' using the uninitialized extension 'SHAPE'. shape.Init(connObj) must be called first.")
	}
	cookie := c.NewCookie(true, true)
	c.NewRequest(queryVersionRequest(c), cookie)
	return QueryVersionCookie{cookie}
}

// QueryVersionUnchecked sends an unchecked request.
// If an error occurs, it can only be retrieved using xgb.WaitForEvent or xgb.PollForEvent.
func QueryVersionUnchecked(c *xgb.Conn) QueryVersionCookie {
	c.ExtLock.RLock()
	defer c.ExtLock.RUnlock()
	if _, ok := c.Extensions["SHAPE"]; !ok {
		panic("Cannot issue request 'QueryVersion' using the uninitialized extension 'SHAPE'. shape.Init(connObj) must be called first.")
	}
	cookie := c.NewCookie(false, true)
	c.NewRequest(queryVersionRequest(c), cookie)
	return QueryVersionCookie{cookie}
}

// QueryVersionReply represents the data returned from a QueryVersion request.
type QueryVersionReply struct {
	Sequence uint16 // sequence number of the request for this reply
	Length   uint32 // number of bytes in this reply
	// padding: 1 bytes
	MajorVersion uint16
	MinorVersion uint16
}

// Reply blocks and returns the reply data for a QueryVersion request.
func (cook QueryVersionCookie) Reply() (*QueryVersionReply, error) {
	buf, err := cook.Cookie.Reply()
	if err != nil {
		return nil, err
	}
	if buf == nil {
		return nil, nil
	}
	return queryVersionReply(buf), nil
}

// queryVersionReply reads a byte slice into a QueryVersionReply value.
func queryVersionReply(buf []byte) *QueryVersionReply {
	v := new(QueryVersionReply)
	b := 1 // skip reply determinant

	b += 1 // padding

	v.Sequence = xgb.Get16(buf[b:])
	b += 2

	v.Length = xgb.Get32(buf[b:]) // 4-byte units
	b += 4

	v.MajorVersion = xgb.Get16(buf[b:])
	b += 2

	v.MinorVersion = xgb.Get16(buf[b:])
	b += 2

	return v
}

// Write request to wire for QueryVersion
// queryVersionRequest writes a QueryVersion request to a byte slice.
func queryVersionRequest(c *xgb.Conn) []byte {
	size := 4
	b := 0
	buf := make([]byte, size)

	c.ExtLock.RLock()
	buf[b] = c.Extensions["SHAPE"]
	c.ExtLock.RUnlock()
	b += 1

	buf[b] = 0 // request opcode
	b += 1

	xgb.Put16(buf[b:], uint16(size/4)) // write request size in 4-byte units
	b += 2

	return buf
}

// RectanglesCookie is a cookie used only for Rectangles requests.
type RectanglesCookie struct {
	*xgb.Cookie
}

// Rectangles sends an unchecked request.
// If an error occurs, it can only be retrieved using xgb.WaitForEvent or xgb.PollForEvent.
func Rectangles(c *xgb.Conn, Operation Op, DestinationKind Kind, Ordering byte, DestinationWindow xproto.Window, XOffset int16, YOffset int16, Rectangles []xproto.Rectangle) RectanglesCookie {
	c.ExtLock.RLock()
	defer c.ExtLock.RUnlock()
	if _, ok := c.Extensions["SHAPE"]; !ok {
		panic("Cannot issue request 'Rectangles' using the uninitialized extension 'SHAPE'. shape.Init(connObj) must be called first.")
	}
	cookie := c.NewCookie(false, false)
	c.NewRequest(rectanglesRequest(c, Operation, DestinationKind, Ordering, DestinationWindow, XOffset, YOffset, Rectangles), cookie)
	return RectanglesCookie{cookie}
}

// RectanglesChecked sends a checked request.
// If an error occurs, it can be retrieved using RectanglesCookie.Check()
func RectanglesChecked(c *xgb.Conn, Operation Op, DestinationKind Kind, Ordering byte, DestinationWindow xproto.Window, XOffset int16, YOffset int16, Rectangles []xproto.Rectangle) RectanglesCookie {
	c.ExtLock.RLock()
	defer c.ExtLock.RUnlock()
	if _, ok := c.Extensions["SHAPE"]; !ok {
		panic("Cannot issue request 'Rectangles' using the uninitialized extension 'SHAPE'. shape.Init(connObj) must be called first.")
	}
	cookie := c.NewCookie(true, false)
	c.NewRequest(rectanglesRequest(c, Operation, DestinationKind, Ordering, DestinationWindow, XOffset, YOffset, Rectangles), cookie)
	return RectanglesCookie{cookie}
}

// Check returns an error if one occurred for checked requests that are not expecting a reply.
// This cannot be called for requests expecting a reply, nor for unchecked requests.
func (cook RectanglesCookie) Check() error {
	return cook.Cookie.Check()
}

// Write request to wire for Rectangles
// rectanglesRequest writes a Rectangles request to a byte slice.
func rectanglesRequest(c *xgb.Conn, Operation Op, DestinationKind Kind, Ordering byte, DestinationWindow xproto.Window, XOffset int16, YOffset int16, Rectangles []xproto.Rectangle) []byte {
	size := xgb.Pad((16 + xgb.Pad((len(Rectangles) * 8))))
	b := 0
	buf := make([]byte, size)

	c.ExtLock.RLock()
	buf[b] = c.Extensions["SHAPE"]
	c.ExtLock.RUnlock()
	b += 1

	buf[b] = 1 // request opcode
	b += 1

	xgb.Put16(buf[b:], uint16(size/4)) // write request size in 4-byte units
	b += 2

	buf[b] = byte(Operation)
	b += 1

	buf[b] = byte(DestinationKind)
	b += 1

	buf[b] = Ordering
	b += 1

	b += 1 // padding

	xgb.Put32(buf[b:], uint32(DestinationWindow))
	b += 4

	xgb.Put16(buf[b:], uint16(XOffset))
	b += 2

	xgb.Put16(buf[b:], uint16(YOffset))
	b += 2

	b += xproto.RectangleListBytes(buf[b:], Rectangles)

	return buf
}

// SelectInputCookie is a cookie used only for SelectInput requests.
type SelectInputCookie struct {
	*xgb.Cookie
}

// SelectInput sends an unchecked request.
// If an error occurs, it can only be retrieved using xgb.WaitForEvent or xgb.PollForEvent.
func SelectInput(c *xgb.Conn, DestinationWindow xproto.Window, Enable bool) SelectInputCookie {
	c.ExtLock.RLock()
	defer c.ExtLock.RUnlock()
	if _, ok := c.Extensions["SHAPE"]; !ok {
		panic("Cannot issue request 'SelectInput' using the uninitialized extension 'SHAPE'. shape.Init(connObj) must be called first.")
	}
	cookie := c.NewCookie(false, false)
	c.NewRequest(selectInputRequest(c, DestinationWindow, Enable), cookie)
	return SelectInputCookie{cookie}
}

// SelectInputChecked sends a checked request.
// If an error occurs, it can be retrieved using SelectInputCookie.Check()
func SelectInputChecked(c *xgb.Conn, DestinationWindow xproto.Window, Enable bool) SelectInputCookie {
	c.ExtLock.RLock()
	defer c.ExtLock.RUnlock()
	if _, ok := c.Extensions["SHAPE"]; !ok {
		panic("Cannot issue request 'SelectInput' using the uninitialized extension 'SHAPE'. shape.Init(connObj) must be called first.")
	}
	cookie := c.NewCookie(true, false)
	c.NewRequest(selectInputRequest(c, DestinationWindow, Enable), cookie)
	return SelectInputCookie{cookie}
}

// Check returns an error if one occurred for checked requests that are not expecting a reply.
// This cannot be called for requests expecting a reply, nor for unchecked requests.
func (cook SelectInputCookie) Check() error {
	return cook.Cookie.Check()
}

// Write request to wire for SelectInput
// selectInputRequest writes a SelectInput request to a byte slice.
func selectInputRequest(c *xgb.Conn, DestinationWindow xproto.Window, Enable bool) []byte {
	size := 12
	b := 0
	buf := make([]byte, size)

	c.ExtLock.RLock()
	buf[b] = c.Extensions["SHAPE"]
	c.ExtLock.RUnlock()
	b += 1

	buf[b] = 6 // request opcode
	b += 1

	xgb.Put16(buf[b:], uint16(size/4)) // write request size in 4-byte units
	b += 2

	xgb.Put32(buf[b:], uint32(DestinationWindow))
	b += 4

	if Enable {
		buf[b] = 1
	} else {
		buf[b] = 0
	}
	b += 1

	b += 3 // padding

	return buf
}
