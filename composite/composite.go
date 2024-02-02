// Code generated by xgbgen from composite.xml. DO NOT EDIT.

// Package composite is the X client API for the Composite extension.
package composite

import (
	"github.com/jezek/xgb"

	"github.com/jezek/xgb/xfixes"
	"github.com/jezek/xgb/xproto"
)

// Init must be called before using the Composite extension.
func Init(c *xgb.Conn) error {
	reply, err := xproto.QueryExtension(c, 9, "Composite").Reply()
	switch {
	case err != nil:
		return err
	case !reply.Present:
		return xgb.Errorf("No extension named Composite could be found on on the server.")
	}

	c.ExtLock.Lock()
	c.Extensions["Composite"] = reply.MajorOpcode
	c.ExtLock.Unlock()
	for evNum, fun := range xgb.NewExtEventFuncs["Composite"] {
		xgb.NewEventFuncs[int(reply.FirstEvent)+evNum] = fun
	}
	for errNum, fun := range xgb.NewExtErrorFuncs["Composite"] {
		xgb.NewErrorFuncs[int(reply.FirstError)+errNum] = fun
	}
	return nil
}

func init() {
	xgb.NewExtEventFuncs["Composite"] = make(map[int]xgb.NewEventFun)
	xgb.NewExtErrorFuncs["Composite"] = make(map[int]xgb.NewErrorFun)
}

const (
	RedirectAutomatic = 0
	RedirectManual    = 1
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

// CreateRegionFromBorderClipCookie is a cookie used only for CreateRegionFromBorderClip requests.
type CreateRegionFromBorderClipCookie struct {
	*xgb.Cookie
}

// CreateRegionFromBorderClip sends an unchecked request.
// If an error occurs, it can only be retrieved using xgb.WaitForEvent or xgb.PollForEvent.
func CreateRegionFromBorderClip(c *xgb.Conn, Region xfixes.Region, Window xproto.Window) CreateRegionFromBorderClipCookie {
	c.ExtLock.RLock()
	defer c.ExtLock.RUnlock()
	if _, ok := c.Extensions["Composite"]; !ok {
		panic("Cannot issue request 'CreateRegionFromBorderClip' using the uninitialized extension 'Composite'. composite.Init(connObj) must be called first.")
	}
	cookie := c.NewCookie(false, false)
	c.NewRequest(createRegionFromBorderClipRequest(c, Region, Window), cookie)
	return CreateRegionFromBorderClipCookie{cookie}
}

// CreateRegionFromBorderClipChecked sends a checked request.
// If an error occurs, it can be retrieved using CreateRegionFromBorderClipCookie.Check()
func CreateRegionFromBorderClipChecked(c *xgb.Conn, Region xfixes.Region, Window xproto.Window) CreateRegionFromBorderClipCookie {
	c.ExtLock.RLock()
	defer c.ExtLock.RUnlock()
	if _, ok := c.Extensions["Composite"]; !ok {
		panic("Cannot issue request 'CreateRegionFromBorderClip' using the uninitialized extension 'Composite'. composite.Init(connObj) must be called first.")
	}
	cookie := c.NewCookie(true, false)
	c.NewRequest(createRegionFromBorderClipRequest(c, Region, Window), cookie)
	return CreateRegionFromBorderClipCookie{cookie}
}

// Check returns an error if one occurred for checked requests that are not expecting a reply.
// This cannot be called for requests expecting a reply, nor for unchecked requests.
func (cook CreateRegionFromBorderClipCookie) Check() error {
	return cook.Cookie.Check()
}

// Write request to wire for CreateRegionFromBorderClip
// createRegionFromBorderClipRequest writes a CreateRegionFromBorderClip request to a byte slice.
func createRegionFromBorderClipRequest(c *xgb.Conn, Region xfixes.Region, Window xproto.Window) []byte {
	size := 12
	b := 0
	buf := make([]byte, size)

	c.ExtLock.RLock()
	buf[b] = c.Extensions["Composite"]
	c.ExtLock.RUnlock()
	b += 1

	buf[b] = 5 // request opcode
	b += 1

	xgb.Put16(buf[b:], uint16(size/4)) // write request size in 4-byte units
	b += 2

	xgb.Put32(buf[b:], uint32(Region))
	b += 4

	xgb.Put32(buf[b:], uint32(Window))
	b += 4

	return buf
}

// GetOverlayWindowCookie is a cookie used only for GetOverlayWindow requests.
type GetOverlayWindowCookie struct {
	*xgb.Cookie
}

// GetOverlayWindow sends a checked request.
// If an error occurs, it will be returned with the reply by calling GetOverlayWindowCookie.Reply()
func GetOverlayWindow(c *xgb.Conn, Window xproto.Window) GetOverlayWindowCookie {
	c.ExtLock.RLock()
	defer c.ExtLock.RUnlock()
	if _, ok := c.Extensions["Composite"]; !ok {
		panic("Cannot issue request 'GetOverlayWindow' using the uninitialized extension 'Composite'. composite.Init(connObj) must be called first.")
	}
	cookie := c.NewCookie(true, true)
	c.NewRequest(getOverlayWindowRequest(c, Window), cookie)
	return GetOverlayWindowCookie{cookie}
}

// GetOverlayWindowUnchecked sends an unchecked request.
// If an error occurs, it can only be retrieved using xgb.WaitForEvent or xgb.PollForEvent.
func GetOverlayWindowUnchecked(c *xgb.Conn, Window xproto.Window) GetOverlayWindowCookie {
	c.ExtLock.RLock()
	defer c.ExtLock.RUnlock()
	if _, ok := c.Extensions["Composite"]; !ok {
		panic("Cannot issue request 'GetOverlayWindow' using the uninitialized extension 'Composite'. composite.Init(connObj) must be called first.")
	}
	cookie := c.NewCookie(false, true)
	c.NewRequest(getOverlayWindowRequest(c, Window), cookie)
	return GetOverlayWindowCookie{cookie}
}

// GetOverlayWindowReply represents the data returned from a GetOverlayWindow request.
type GetOverlayWindowReply struct {
	Sequence uint16 // sequence number of the request for this reply
	Length   uint32 // number of bytes in this reply
	// padding: 1 bytes
	OverlayWin xproto.Window
	// padding: 20 bytes
}

// Reply blocks and returns the reply data for a GetOverlayWindow request.
func (cook GetOverlayWindowCookie) Reply() (*GetOverlayWindowReply, error) {
	buf, err := cook.Cookie.Reply()
	if err != nil {
		return nil, err
	}
	if buf == nil {
		return nil, nil
	}
	return getOverlayWindowReply(buf), nil
}

// getOverlayWindowReply reads a byte slice into a GetOverlayWindowReply value.
func getOverlayWindowReply(buf []byte) *GetOverlayWindowReply {
	v := new(GetOverlayWindowReply)
	b := 1 // skip reply determinant

	b += 1 // padding

	v.Sequence = xgb.Get16(buf[b:])
	b += 2

	v.Length = xgb.Get32(buf[b:]) // 4-byte units
	b += 4

	v.OverlayWin = xproto.Window(xgb.Get32(buf[b:]))
	b += 4

	b += 20 // padding

	return v
}

// Write request to wire for GetOverlayWindow
// getOverlayWindowRequest writes a GetOverlayWindow request to a byte slice.
func getOverlayWindowRequest(c *xgb.Conn, Window xproto.Window) []byte {
	size := 8
	b := 0
	buf := make([]byte, size)

	c.ExtLock.RLock()
	buf[b] = c.Extensions["Composite"]
	c.ExtLock.RUnlock()
	b += 1

	buf[b] = 7 // request opcode
	b += 1

	xgb.Put16(buf[b:], uint16(size/4)) // write request size in 4-byte units
	b += 2

	xgb.Put32(buf[b:], uint32(Window))
	b += 4

	return buf
}

// NameWindowPixmapCookie is a cookie used only for NameWindowPixmap requests.
type NameWindowPixmapCookie struct {
	*xgb.Cookie
}

// NameWindowPixmap sends an unchecked request.
// If an error occurs, it can only be retrieved using xgb.WaitForEvent or xgb.PollForEvent.
func NameWindowPixmap(c *xgb.Conn, Window xproto.Window, Pixmap xproto.Pixmap) NameWindowPixmapCookie {
	c.ExtLock.RLock()
	defer c.ExtLock.RUnlock()
	if _, ok := c.Extensions["Composite"]; !ok {
		panic("Cannot issue request 'NameWindowPixmap' using the uninitialized extension 'Composite'. composite.Init(connObj) must be called first.")
	}
	cookie := c.NewCookie(false, false)
	c.NewRequest(nameWindowPixmapRequest(c, Window, Pixmap), cookie)
	return NameWindowPixmapCookie{cookie}
}

// NameWindowPixmapChecked sends a checked request.
// If an error occurs, it can be retrieved using NameWindowPixmapCookie.Check()
func NameWindowPixmapChecked(c *xgb.Conn, Window xproto.Window, Pixmap xproto.Pixmap) NameWindowPixmapCookie {
	c.ExtLock.RLock()
	defer c.ExtLock.RUnlock()
	if _, ok := c.Extensions["Composite"]; !ok {
		panic("Cannot issue request 'NameWindowPixmap' using the uninitialized extension 'Composite'. composite.Init(connObj) must be called first.")
	}
	cookie := c.NewCookie(true, false)
	c.NewRequest(nameWindowPixmapRequest(c, Window, Pixmap), cookie)
	return NameWindowPixmapCookie{cookie}
}

// Check returns an error if one occurred for checked requests that are not expecting a reply.
// This cannot be called for requests expecting a reply, nor for unchecked requests.
func (cook NameWindowPixmapCookie) Check() error {
	return cook.Cookie.Check()
}

// Write request to wire for NameWindowPixmap
// nameWindowPixmapRequest writes a NameWindowPixmap request to a byte slice.
func nameWindowPixmapRequest(c *xgb.Conn, Window xproto.Window, Pixmap xproto.Pixmap) []byte {
	size := 12
	b := 0
	buf := make([]byte, size)

	c.ExtLock.RLock()
	buf[b] = c.Extensions["Composite"]
	c.ExtLock.RUnlock()
	b += 1

	buf[b] = 6 // request opcode
	b += 1

	xgb.Put16(buf[b:], uint16(size/4)) // write request size in 4-byte units
	b += 2

	xgb.Put32(buf[b:], uint32(Window))
	b += 4

	xgb.Put32(buf[b:], uint32(Pixmap))
	b += 4

	return buf
}

// QueryVersionCookie is a cookie used only for QueryVersion requests.
type QueryVersionCookie struct {
	*xgb.Cookie
}

// QueryVersion sends a checked request.
// If an error occurs, it will be returned with the reply by calling QueryVersionCookie.Reply()
func QueryVersion(c *xgb.Conn, ClientMajorVersion uint32, ClientMinorVersion uint32) QueryVersionCookie {
	c.ExtLock.RLock()
	defer c.ExtLock.RUnlock()
	if _, ok := c.Extensions["Composite"]; !ok {
		panic("Cannot issue request 'QueryVersion' using the uninitialized extension 'Composite'. composite.Init(connObj) must be called first.")
	}
	cookie := c.NewCookie(true, true)
	c.NewRequest(queryVersionRequest(c, ClientMajorVersion, ClientMinorVersion), cookie)
	return QueryVersionCookie{cookie}
}

// QueryVersionUnchecked sends an unchecked request.
// If an error occurs, it can only be retrieved using xgb.WaitForEvent or xgb.PollForEvent.
func QueryVersionUnchecked(c *xgb.Conn, ClientMajorVersion uint32, ClientMinorVersion uint32) QueryVersionCookie {
	c.ExtLock.RLock()
	defer c.ExtLock.RUnlock()
	if _, ok := c.Extensions["Composite"]; !ok {
		panic("Cannot issue request 'QueryVersion' using the uninitialized extension 'Composite'. composite.Init(connObj) must be called first.")
	}
	cookie := c.NewCookie(false, true)
	c.NewRequest(queryVersionRequest(c, ClientMajorVersion, ClientMinorVersion), cookie)
	return QueryVersionCookie{cookie}
}

// QueryVersionReply represents the data returned from a QueryVersion request.
type QueryVersionReply struct {
	Sequence uint16 // sequence number of the request for this reply
	Length   uint32 // number of bytes in this reply
	// padding: 1 bytes
	MajorVersion uint32
	MinorVersion uint32
	// padding: 16 bytes
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

	v.MajorVersion = xgb.Get32(buf[b:])
	b += 4

	v.MinorVersion = xgb.Get32(buf[b:])
	b += 4

	b += 16 // padding

	return v
}

// Write request to wire for QueryVersion
// queryVersionRequest writes a QueryVersion request to a byte slice.
func queryVersionRequest(c *xgb.Conn, ClientMajorVersion uint32, ClientMinorVersion uint32) []byte {
	size := 12
	b := 0
	buf := make([]byte, size)

	c.ExtLock.RLock()
	buf[b] = c.Extensions["Composite"]
	c.ExtLock.RUnlock()
	b += 1

	buf[b] = 0 // request opcode
	b += 1

	xgb.Put16(buf[b:], uint16(size/4)) // write request size in 4-byte units
	b += 2

	xgb.Put32(buf[b:], ClientMajorVersion)
	b += 4

	xgb.Put32(buf[b:], ClientMinorVersion)
	b += 4

	return buf
}

// RedirectSubwindowsCookie is a cookie used only for RedirectSubwindows requests.
type RedirectSubwindowsCookie struct {
	*xgb.Cookie
}

// RedirectSubwindows sends an unchecked request.
// If an error occurs, it can only be retrieved using xgb.WaitForEvent or xgb.PollForEvent.
func RedirectSubwindows(c *xgb.Conn, Window xproto.Window, Update byte) RedirectSubwindowsCookie {
	c.ExtLock.RLock()
	defer c.ExtLock.RUnlock()
	if _, ok := c.Extensions["Composite"]; !ok {
		panic("Cannot issue request 'RedirectSubwindows' using the uninitialized extension 'Composite'. composite.Init(connObj) must be called first.")
	}
	cookie := c.NewCookie(false, false)
	c.NewRequest(redirectSubwindowsRequest(c, Window, Update), cookie)
	return RedirectSubwindowsCookie{cookie}
}

// RedirectSubwindowsChecked sends a checked request.
// If an error occurs, it can be retrieved using RedirectSubwindowsCookie.Check()
func RedirectSubwindowsChecked(c *xgb.Conn, Window xproto.Window, Update byte) RedirectSubwindowsCookie {
	c.ExtLock.RLock()
	defer c.ExtLock.RUnlock()
	if _, ok := c.Extensions["Composite"]; !ok {
		panic("Cannot issue request 'RedirectSubwindows' using the uninitialized extension 'Composite'. composite.Init(connObj) must be called first.")
	}
	cookie := c.NewCookie(true, false)
	c.NewRequest(redirectSubwindowsRequest(c, Window, Update), cookie)
	return RedirectSubwindowsCookie{cookie}
}

// Check returns an error if one occurred for checked requests that are not expecting a reply.
// This cannot be called for requests expecting a reply, nor for unchecked requests.
func (cook RedirectSubwindowsCookie) Check() error {
	return cook.Cookie.Check()
}

// Write request to wire for RedirectSubwindows
// redirectSubwindowsRequest writes a RedirectSubwindows request to a byte slice.
func redirectSubwindowsRequest(c *xgb.Conn, Window xproto.Window, Update byte) []byte {
	size := 12
	b := 0
	buf := make([]byte, size)

	c.ExtLock.RLock()
	buf[b] = c.Extensions["Composite"]
	c.ExtLock.RUnlock()
	b += 1

	buf[b] = 2 // request opcode
	b += 1

	xgb.Put16(buf[b:], uint16(size/4)) // write request size in 4-byte units
	b += 2

	xgb.Put32(buf[b:], uint32(Window))
	b += 4

	buf[b] = Update
	b += 1

	b += 3 // padding

	return buf
}

// RedirectWindowCookie is a cookie used only for RedirectWindow requests.
type RedirectWindowCookie struct {
	*xgb.Cookie
}

// RedirectWindow sends an unchecked request.
// If an error occurs, it can only be retrieved using xgb.WaitForEvent or xgb.PollForEvent.
func RedirectWindow(c *xgb.Conn, Window xproto.Window, Update byte) RedirectWindowCookie {
	c.ExtLock.RLock()
	defer c.ExtLock.RUnlock()
	if _, ok := c.Extensions["Composite"]; !ok {
		panic("Cannot issue request 'RedirectWindow' using the uninitialized extension 'Composite'. composite.Init(connObj) must be called first.")
	}
	cookie := c.NewCookie(false, false)
	c.NewRequest(redirectWindowRequest(c, Window, Update), cookie)
	return RedirectWindowCookie{cookie}
}

// RedirectWindowChecked sends a checked request.
// If an error occurs, it can be retrieved using RedirectWindowCookie.Check()
func RedirectWindowChecked(c *xgb.Conn, Window xproto.Window, Update byte) RedirectWindowCookie {
	c.ExtLock.RLock()
	defer c.ExtLock.RUnlock()
	if _, ok := c.Extensions["Composite"]; !ok {
		panic("Cannot issue request 'RedirectWindow' using the uninitialized extension 'Composite'. composite.Init(connObj) must be called first.")
	}
	cookie := c.NewCookie(true, false)
	c.NewRequest(redirectWindowRequest(c, Window, Update), cookie)
	return RedirectWindowCookie{cookie}
}

// Check returns an error if one occurred for checked requests that are not expecting a reply.
// This cannot be called for requests expecting a reply, nor for unchecked requests.
func (cook RedirectWindowCookie) Check() error {
	return cook.Cookie.Check()
}

// Write request to wire for RedirectWindow
// redirectWindowRequest writes a RedirectWindow request to a byte slice.
func redirectWindowRequest(c *xgb.Conn, Window xproto.Window, Update byte) []byte {
	size := 12
	b := 0
	buf := make([]byte, size)

	c.ExtLock.RLock()
	buf[b] = c.Extensions["Composite"]
	c.ExtLock.RUnlock()
	b += 1

	buf[b] = 1 // request opcode
	b += 1

	xgb.Put16(buf[b:], uint16(size/4)) // write request size in 4-byte units
	b += 2

	xgb.Put32(buf[b:], uint32(Window))
	b += 4

	buf[b] = Update
	b += 1

	b += 3 // padding

	return buf
}

// ReleaseOverlayWindowCookie is a cookie used only for ReleaseOverlayWindow requests.
type ReleaseOverlayWindowCookie struct {
	*xgb.Cookie
}

// ReleaseOverlayWindow sends an unchecked request.
// If an error occurs, it can only be retrieved using xgb.WaitForEvent or xgb.PollForEvent.
func ReleaseOverlayWindow(c *xgb.Conn, Window xproto.Window) ReleaseOverlayWindowCookie {
	c.ExtLock.RLock()
	defer c.ExtLock.RUnlock()
	if _, ok := c.Extensions["Composite"]; !ok {
		panic("Cannot issue request 'ReleaseOverlayWindow' using the uninitialized extension 'Composite'. composite.Init(connObj) must be called first.")
	}
	cookie := c.NewCookie(false, false)
	c.NewRequest(releaseOverlayWindowRequest(c, Window), cookie)
	return ReleaseOverlayWindowCookie{cookie}
}

// ReleaseOverlayWindowChecked sends a checked request.
// If an error occurs, it can be retrieved using ReleaseOverlayWindowCookie.Check()
func ReleaseOverlayWindowChecked(c *xgb.Conn, Window xproto.Window) ReleaseOverlayWindowCookie {
	c.ExtLock.RLock()
	defer c.ExtLock.RUnlock()
	if _, ok := c.Extensions["Composite"]; !ok {
		panic("Cannot issue request 'ReleaseOverlayWindow' using the uninitialized extension 'Composite'. composite.Init(connObj) must be called first.")
	}
	cookie := c.NewCookie(true, false)
	c.NewRequest(releaseOverlayWindowRequest(c, Window), cookie)
	return ReleaseOverlayWindowCookie{cookie}
}

// Check returns an error if one occurred for checked requests that are not expecting a reply.
// This cannot be called for requests expecting a reply, nor for unchecked requests.
func (cook ReleaseOverlayWindowCookie) Check() error {
	return cook.Cookie.Check()
}

// Write request to wire for ReleaseOverlayWindow
// releaseOverlayWindowRequest writes a ReleaseOverlayWindow request to a byte slice.
func releaseOverlayWindowRequest(c *xgb.Conn, Window xproto.Window) []byte {
	size := 8
	b := 0
	buf := make([]byte, size)

	c.ExtLock.RLock()
	buf[b] = c.Extensions["Composite"]
	c.ExtLock.RUnlock()
	b += 1

	buf[b] = 8 // request opcode
	b += 1

	xgb.Put16(buf[b:], uint16(size/4)) // write request size in 4-byte units
	b += 2

	xgb.Put32(buf[b:], uint32(Window))
	b += 4

	return buf
}

// UnredirectSubwindowsCookie is a cookie used only for UnredirectSubwindows requests.
type UnredirectSubwindowsCookie struct {
	*xgb.Cookie
}

// UnredirectSubwindows sends an unchecked request.
// If an error occurs, it can only be retrieved using xgb.WaitForEvent or xgb.PollForEvent.
func UnredirectSubwindows(c *xgb.Conn, Window xproto.Window, Update byte) UnredirectSubwindowsCookie {
	c.ExtLock.RLock()
	defer c.ExtLock.RUnlock()
	if _, ok := c.Extensions["Composite"]; !ok {
		panic("Cannot issue request 'UnredirectSubwindows' using the uninitialized extension 'Composite'. composite.Init(connObj) must be called first.")
	}
	cookie := c.NewCookie(false, false)
	c.NewRequest(unredirectSubwindowsRequest(c, Window, Update), cookie)
	return UnredirectSubwindowsCookie{cookie}
}

// UnredirectSubwindowsChecked sends a checked request.
// If an error occurs, it can be retrieved using UnredirectSubwindowsCookie.Check()
func UnredirectSubwindowsChecked(c *xgb.Conn, Window xproto.Window, Update byte) UnredirectSubwindowsCookie {
	c.ExtLock.RLock()
	defer c.ExtLock.RUnlock()
	if _, ok := c.Extensions["Composite"]; !ok {
		panic("Cannot issue request 'UnredirectSubwindows' using the uninitialized extension 'Composite'. composite.Init(connObj) must be called first.")
	}
	cookie := c.NewCookie(true, false)
	c.NewRequest(unredirectSubwindowsRequest(c, Window, Update), cookie)
	return UnredirectSubwindowsCookie{cookie}
}

// Check returns an error if one occurred for checked requests that are not expecting a reply.
// This cannot be called for requests expecting a reply, nor for unchecked requests.
func (cook UnredirectSubwindowsCookie) Check() error {
	return cook.Cookie.Check()
}

// Write request to wire for UnredirectSubwindows
// unredirectSubwindowsRequest writes a UnredirectSubwindows request to a byte slice.
func unredirectSubwindowsRequest(c *xgb.Conn, Window xproto.Window, Update byte) []byte {
	size := 12
	b := 0
	buf := make([]byte, size)

	c.ExtLock.RLock()
	buf[b] = c.Extensions["Composite"]
	c.ExtLock.RUnlock()
	b += 1

	buf[b] = 4 // request opcode
	b += 1

	xgb.Put16(buf[b:], uint16(size/4)) // write request size in 4-byte units
	b += 2

	xgb.Put32(buf[b:], uint32(Window))
	b += 4

	buf[b] = Update
	b += 1

	b += 3 // padding

	return buf
}

// UnredirectWindowCookie is a cookie used only for UnredirectWindow requests.
type UnredirectWindowCookie struct {
	*xgb.Cookie
}

// UnredirectWindow sends an unchecked request.
// If an error occurs, it can only be retrieved using xgb.WaitForEvent or xgb.PollForEvent.
func UnredirectWindow(c *xgb.Conn, Window xproto.Window, Update byte) UnredirectWindowCookie {
	c.ExtLock.RLock()
	defer c.ExtLock.RUnlock()
	if _, ok := c.Extensions["Composite"]; !ok {
		panic("Cannot issue request 'UnredirectWindow' using the uninitialized extension 'Composite'. composite.Init(connObj) must be called first.")
	}
	cookie := c.NewCookie(false, false)
	c.NewRequest(unredirectWindowRequest(c, Window, Update), cookie)
	return UnredirectWindowCookie{cookie}
}

// UnredirectWindowChecked sends a checked request.
// If an error occurs, it can be retrieved using UnredirectWindowCookie.Check()
func UnredirectWindowChecked(c *xgb.Conn, Window xproto.Window, Update byte) UnredirectWindowCookie {
	c.ExtLock.RLock()
	defer c.ExtLock.RUnlock()
	if _, ok := c.Extensions["Composite"]; !ok {
		panic("Cannot issue request 'UnredirectWindow' using the uninitialized extension 'Composite'. composite.Init(connObj) must be called first.")
	}
	cookie := c.NewCookie(true, false)
	c.NewRequest(unredirectWindowRequest(c, Window, Update), cookie)
	return UnredirectWindowCookie{cookie}
}

// Check returns an error if one occurred for checked requests that are not expecting a reply.
// This cannot be called for requests expecting a reply, nor for unchecked requests.
func (cook UnredirectWindowCookie) Check() error {
	return cook.Cookie.Check()
}

// Write request to wire for UnredirectWindow
// unredirectWindowRequest writes a UnredirectWindow request to a byte slice.
func unredirectWindowRequest(c *xgb.Conn, Window xproto.Window, Update byte) []byte {
	size := 12
	b := 0
	buf := make([]byte, size)

	c.ExtLock.RLock()
	buf[b] = c.Extensions["Composite"]
	c.ExtLock.RUnlock()
	b += 1

	buf[b] = 3 // request opcode
	b += 1

	xgb.Put16(buf[b:], uint16(size/4)) // write request size in 4-byte units
	b += 2

	xgb.Put32(buf[b:], uint32(Window))
	b += 4

	buf[b] = Update
	b += 1

	b += 3 // padding

	return buf
}
