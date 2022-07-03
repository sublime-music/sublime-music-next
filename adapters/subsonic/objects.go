package subsonic

import (
	"encoding/json"
	"fmt"
	"time"
)

// An identifier for a Subsonic resource. Can be either an integer or string
// depending on what the original author of the Subsonic API was feeling that
// day.
type SubsonicIdentifier string

// Custom unmarshaller for SubsonicIdentifier that normalizes to using strings
// to represent identifiers.
func (id *SubsonicIdentifier) UnmarshalJSON(data []byte) error {
	var i int
	if err := json.Unmarshal(data, &i); err != nil {
		var s string
		if err := json.Unmarshal(data, &s); err != nil {
			return err
		}
		*id = SubsonicIdentifier(s)
		return nil
	}
	*id = SubsonicIdentifier(fmt.Sprintf("%d", i))
	return nil
}

func (id *SubsonicIdentifier) String() string {
	return string(*id)
}

type Song struct {
}

type Playlist struct {
	ID        *SubsonicIdentifier `json:"id"`
	Name      string              `json:"name"`
	SongCount *int                `json:"songCount"`
	Duration  *time.Duration      `json:"duration"`
	Songs     []*Song             `json:"entry"`
	Created   *time.Time          `json:"created"`
	Changed   *time.Time          `json:"changed"`
	Comment   *string             `json:"comment"`
	Owner     *string             `json:"owner"`
	Public    *bool               `json:"public"`
	CoverArt  *string             `json:"coverArt"`
}

type Playlists struct {
	Playlist []*Playlist `json:"playlist"`
}

type SubsonicError struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

func (err *SubsonicError) Error() string {
	return fmt.Sprintf("SubsonicError %d: %s", err.Code, err.Message)
}

func (err *SubsonicError) Is(other error) bool {
	otherErr, ok := other.(*SubsonicError)
	if !ok {
		return false
	}
	return otherErr.Code == err.Code && otherErr.Message == err.Message
}

type SubsonicResponse struct {
	Status  string         `json:"status"`
	Error   *SubsonicError `json:"error"`
	Version string         `json:"version"`

	Playlists *Playlists `json:"playlists"`
}

type Response struct {
	SubsonicResponse *SubsonicResponse `json:"subsonic-response"`
}
