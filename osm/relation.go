package osm

import (
	"encoding/json"
	"encoding/xml"

	"github.com/paulmach/orb"
)

// xmlNameJSONTypeRel is kind of a hack to encode the proper json
// object type attribute for this struct type.
type xmlNameJSONTypeRel xml.Name

func (x xmlNameJSONTypeRel) MarshalJSON() ([]byte, error) {
	return []byte(`"relation"`), nil
}

// Relation is an collection of nodes, ways and other relations
// with some defining attributes.
type Relation struct {
	XMLName     xmlNameJSONTypeRel `xml:"relation" json:"type"`
	ID          int64              `xml:"id,attr" json:"id"`
	Visible     bool               `xml:"visible,attr" json:"visible"`
	Version     int                `xml:"version,attr" json:"version,omitempty"`
	User        *string            `xml:"user,attr" json:"user,omitempty"`
	UserID      *int64             `xml:"uid,attr" json:"uid,omitempty"`
	ChangesetID int64              `xml:"changeset,attr" json:"changeset,omitempty"`
	Timestamp   TimeOSM            `xml:"timestamp,attr" json:"timestamp,omitempty"`
	Tags        Tags               `xml:"tag" json:"tags,omitempty"`
	Members     Members            `xml:"member" json:"members"`
}

// Members represents an ordered list of relation members.
type Members []Member

// Scan - Implement the database/sql scanner interface
func (ms *Members) Scan(value interface{}) error {
	return json.Unmarshal(value.([]byte), ms)
}

// Member is a member of a relation.
type Member struct {
	Type string `xml:"type,attr" json:"type"`
	Ref  int64  `xml:"ref,attr" json:"ref"`
	Role string `xml:"role,attr" json:"role"`

	Version     int   `xml:"version,attr,omitempty" json:"version,omitempty"`
	ChangesetID int64 `xml:"changeset,attr,omitempty" json:"changeset,omitempty"`
	// Node location if Type == Node
	// Closest vertex to centroid if Type == Way
	// Empty/invalid if Type == Relation
	Lat float64 `xml:"lat,attr,omitempty" json:"lat,omitempty"`
	Lon float64 `xml:"lon,attr,omitempty" json:"lon,omitempty"`

	// Orientation is the direction of the way around a ring of a multipolygon.
	// Only valid for multipolygon or boundary relations.
	Orientation orb.Orientation `xml:"orienation,attr,omitempty" json:"orienation,omitempty"`
}

// ObjectID returns the object id of the relation.
func (r *Relation) ObjectID() int64 {
	return r.ID
}

// Relations is a list of relations with helper functions on top.
type Relations []*Relation

// Scan - Implement the database/sql scanner interface
func (rs *Relations) Scan(value interface{}) error {
	return json.Unmarshal(value.([]byte), rs)
}
