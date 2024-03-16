// Code generated by ogen, DO NOT EDIT.

package api

// DeleteAPIV1NoteBadRequest is response for DeleteAPIV1Note operation.
type DeleteAPIV1NoteBadRequest struct{}

func (*DeleteAPIV1NoteBadRequest) deleteAPIV1NoteRes() {}

// DeleteAPIV1NoteOK is response for DeleteAPIV1Note operation.
type DeleteAPIV1NoteOK struct{}

func (*DeleteAPIV1NoteOK) deleteAPIV1NoteRes() {}

// DeleteAPIV1UserBadRequest is response for DeleteAPIV1User operation.
type DeleteAPIV1UserBadRequest struct{}

func (*DeleteAPIV1UserBadRequest) deleteAPIV1UserRes() {}

// DeleteAPIV1UserOK is response for DeleteAPIV1User operation.
type DeleteAPIV1UserOK struct{}

func (*DeleteAPIV1UserOK) deleteAPIV1UserRes() {}

// GetAPIV1NoteBadRequest is response for GetAPIV1Note operation.
type GetAPIV1NoteBadRequest struct{}

func (*GetAPIV1NoteBadRequest) getAPIV1NoteRes() {}

// GetAPIV1UserBadRequest is response for GetAPIV1User operation.
type GetAPIV1UserBadRequest struct{}

func (*GetAPIV1UserBadRequest) getAPIV1UserRes() {}

// ListAPIV1NoteBadRequest is response for ListAPIV1Note operation.
type ListAPIV1NoteBadRequest struct{}

func (*ListAPIV1NoteBadRequest) listAPIV1NoteRes() {}

// ListAPIV1NoteOK is response for ListAPIV1Note operation.
type ListAPIV1NoteOK struct{}

func (*ListAPIV1NoteOK) listAPIV1NoteRes() {}

// Ref: #/components/schemas/Note
type Note struct {
	Name        OptString `json:"name"`
	Description OptString `json:"description"`
	StartTime   OptString `json:"startTime"`
	EndTime     OptString `json:"endTime"`
}

// GetName returns the value of Name.
func (s *Note) GetName() OptString {
	return s.Name
}

// GetDescription returns the value of Description.
func (s *Note) GetDescription() OptString {
	return s.Description
}

// GetStartTime returns the value of StartTime.
func (s *Note) GetStartTime() OptString {
	return s.StartTime
}

// GetEndTime returns the value of EndTime.
func (s *Note) GetEndTime() OptString {
	return s.EndTime
}

// SetName sets the value of Name.
func (s *Note) SetName(val OptString) {
	s.Name = val
}

// SetDescription sets the value of Description.
func (s *Note) SetDescription(val OptString) {
	s.Description = val
}

// SetStartTime sets the value of StartTime.
func (s *Note) SetStartTime(val OptString) {
	s.StartTime = val
}

// SetEndTime sets the value of EndTime.
func (s *Note) SetEndTime(val OptString) {
	s.EndTime = val
}

// NoteHeaders wraps Note with response headers.
type NoteHeaders struct {
	AccessControlAllowHeaders OptString
	AccessControlAllowMethods OptString
	AccessControlAllowOrigin  OptString
	Response                  Note
}

// GetAccessControlAllowHeaders returns the value of AccessControlAllowHeaders.
func (s *NoteHeaders) GetAccessControlAllowHeaders() OptString {
	return s.AccessControlAllowHeaders
}

// GetAccessControlAllowMethods returns the value of AccessControlAllowMethods.
func (s *NoteHeaders) GetAccessControlAllowMethods() OptString {
	return s.AccessControlAllowMethods
}

// GetAccessControlAllowOrigin returns the value of AccessControlAllowOrigin.
func (s *NoteHeaders) GetAccessControlAllowOrigin() OptString {
	return s.AccessControlAllowOrigin
}

// GetResponse returns the value of Response.
func (s *NoteHeaders) GetResponse() Note {
	return s.Response
}

// SetAccessControlAllowHeaders sets the value of AccessControlAllowHeaders.
func (s *NoteHeaders) SetAccessControlAllowHeaders(val OptString) {
	s.AccessControlAllowHeaders = val
}

// SetAccessControlAllowMethods sets the value of AccessControlAllowMethods.
func (s *NoteHeaders) SetAccessControlAllowMethods(val OptString) {
	s.AccessControlAllowMethods = val
}

// SetAccessControlAllowOrigin sets the value of AccessControlAllowOrigin.
func (s *NoteHeaders) SetAccessControlAllowOrigin(val OptString) {
	s.AccessControlAllowOrigin = val
}

// SetResponse sets the value of Response.
func (s *NoteHeaders) SetResponse(val Note) {
	s.Response = val
}

func (*NoteHeaders) getAPIV1NoteRes() {}

// NewOptNote returns new OptNote with value set to v.
func NewOptNote(v Note) OptNote {
	return OptNote{
		Value: v,
		Set:   true,
	}
}

// OptNote is optional Note.
type OptNote struct {
	Value Note
	Set   bool
}

// IsSet returns true if OptNote was set.
func (o OptNote) IsSet() bool { return o.Set }

// Reset unsets value.
func (o *OptNote) Reset() {
	var v Note
	o.Value = v
	o.Set = false
}

// SetTo sets value to v.
func (o *OptNote) SetTo(v Note) {
	o.Set = true
	o.Value = v
}

// Get returns value and boolean that denotes whether value was set.
func (o OptNote) Get() (v Note, ok bool) {
	if !o.Set {
		return v, false
	}
	return o.Value, true
}

// Or returns value if set, or given parameter if does not.
func (o OptNote) Or(d Note) Note {
	if v, ok := o.Get(); ok {
		return v
	}
	return d
}

// NewOptString returns new OptString with value set to v.
func NewOptString(v string) OptString {
	return OptString{
		Value: v,
		Set:   true,
	}
}

// OptString is optional string.
type OptString struct {
	Value string
	Set   bool
}

// IsSet returns true if OptString was set.
func (o OptString) IsSet() bool { return o.Set }

// Reset unsets value.
func (o *OptString) Reset() {
	var v string
	o.Value = v
	o.Set = false
}

// SetTo sets value to v.
func (o *OptString) SetTo(v string) {
	o.Set = true
	o.Value = v
}

// Get returns value and boolean that denotes whether value was set.
func (o OptString) Get() (v string, ok bool) {
	if !o.Set {
		return v, false
	}
	return o.Value, true
}

// Or returns value if set, or given parameter if does not.
func (o OptString) Or(d string) string {
	if v, ok := o.Get(); ok {
		return v
	}
	return d
}

// NewOptUser returns new OptUser with value set to v.
func NewOptUser(v User) OptUser {
	return OptUser{
		Value: v,
		Set:   true,
	}
}

// OptUser is optional User.
type OptUser struct {
	Value User
	Set   bool
}

// IsSet returns true if OptUser was set.
func (o OptUser) IsSet() bool { return o.Set }

// Reset unsets value.
func (o *OptUser) Reset() {
	var v User
	o.Value = v
	o.Set = false
}

// SetTo sets value to v.
func (o *OptUser) SetTo(v User) {
	o.Set = true
	o.Value = v
}

// Get returns value and boolean that denotes whether value was set.
func (o OptUser) Get() (v User, ok bool) {
	if !o.Set {
		return v, false
	}
	return o.Value, true
}

// Or returns value if set, or given parameter if does not.
func (o OptUser) Or(d User) User {
	if v, ok := o.Get(); ok {
		return v
	}
	return d
}

// PostAPIV1NoteBadRequest is response for PostAPIV1Note operation.
type PostAPIV1NoteBadRequest struct{}

func (*PostAPIV1NoteBadRequest) postAPIV1NoteRes() {}

// PostAPIV1NoteOK is response for PostAPIV1Note operation.
type PostAPIV1NoteOK struct{}

func (*PostAPIV1NoteOK) postAPIV1NoteRes() {}

// PostAPIV1UserBadRequest is response for PostAPIV1User operation.
type PostAPIV1UserBadRequest struct{}

func (*PostAPIV1UserBadRequest) postAPIV1UserRes() {}

// PostAPIV1UserOK is response for PostAPIV1User operation.
type PostAPIV1UserOK struct{}

func (*PostAPIV1UserOK) postAPIV1UserRes() {}

// PutAPIV1NoteBadRequest is response for PutAPIV1Note operation.
type PutAPIV1NoteBadRequest struct{}

func (*PutAPIV1NoteBadRequest) putAPIV1NoteRes() {}

// PutAPIV1NoteOK is response for PutAPIV1Note operation.
type PutAPIV1NoteOK struct{}

func (*PutAPIV1NoteOK) putAPIV1NoteRes() {}

// PutAPIV1UserBadRequest is response for PutAPIV1User operation.
type PutAPIV1UserBadRequest struct{}

func (*PutAPIV1UserBadRequest) putAPIV1UserRes() {}

// PutAPIV1UserOK is response for PutAPIV1User operation.
type PutAPIV1UserOK struct{}

func (*PutAPIV1UserOK) putAPIV1UserRes() {}

// Ref: #/components/schemas/User
type User struct {
	Name     OptString `json:"name"`
	Password OptString `json:"password"`
}

// GetName returns the value of Name.
func (s *User) GetName() OptString {
	return s.Name
}

// GetPassword returns the value of Password.
func (s *User) GetPassword() OptString {
	return s.Password
}

// SetName sets the value of Name.
func (s *User) SetName(val OptString) {
	s.Name = val
}

// SetPassword sets the value of Password.
func (s *User) SetPassword(val OptString) {
	s.Password = val
}

func (*User) getAPIV1UserRes() {}
