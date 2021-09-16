package test

// Code generated by github.com/Khan/genqlient, DO NOT EDIT.

import (
	"encoding/json"
	"fmt"

	"github.com/Khan/genqlient/graphql"
	"github.com/Khan/genqlient/internal/testutil"
)

// Item includes the requested fields of the GraphQL interface Content.
//
// Item is implemented by the following types:
// ItemArticle
// ItemVideo
// ItemTopic
// The GraphQL type's documentation follows.
//
// Content is implemented by various types like Article, Video, and Topic.
type Item interface {
	implementsGraphQLInterfaceItem()
	// GetTypename returns the receiver's concrete GraphQL type-name (see interface doc for possible values).
	GetTypename() string
	// GetId returns the interface-field "id" from its implementation.
	// The GraphQL interface field's documentation follows.
	//
	// ID is the identifier of the content.
	GetId() testutil.ID
	// GetName returns the interface-field "name" from its implementation.
	GetName() string
}

func (v *ItemArticle) implementsGraphQLInterfaceItem() {}

// GetTypename is a part of, and documented with, the interface Item.
func (v *ItemArticle) GetTypename() string { return v.Typename }

// GetId is a part of, and documented with, the interface Item.
func (v *ItemArticle) GetId() testutil.ID { return v.Id }

// GetName is a part of, and documented with, the interface Item.
func (v *ItemArticle) GetName() string { return v.Name }

func (v *ItemVideo) implementsGraphQLInterfaceItem() {}

// GetTypename is a part of, and documented with, the interface Item.
func (v *ItemVideo) GetTypename() string { return v.Typename }

// GetId is a part of, and documented with, the interface Item.
func (v *ItemVideo) GetId() testutil.ID { return v.Id }

// GetName is a part of, and documented with, the interface Item.
func (v *ItemVideo) GetName() string { return v.Name }

func (v *ItemTopic) implementsGraphQLInterfaceItem() {}

// GetTypename is a part of, and documented with, the interface Item.
func (v *ItemTopic) GetTypename() string { return v.Typename }

// GetId is a part of, and documented with, the interface Item.
func (v *ItemTopic) GetId() testutil.ID { return v.Id }

// GetName is a part of, and documented with, the interface Item.
func (v *ItemTopic) GetName() string { return v.Name }

func __unmarshalItem(v *Item, m json.RawMessage) error {
	if string(m) == "null" {
		return nil
	}

	var tn struct {
		TypeName string `json:"__typename"`
	}
	err := json.Unmarshal(m, &tn)
	if err != nil {
		return err
	}

	switch tn.TypeName {
	case "Article":
		*v = new(ItemArticle)
		return json.Unmarshal(m, *v)
	case "Video":
		*v = new(ItemVideo)
		return json.Unmarshal(m, *v)
	case "Topic":
		*v = new(ItemTopic)
		return json.Unmarshal(m, *v)
	case "":
		return fmt.Errorf(
			"Response was missing Content.__typename")
	default:
		return fmt.Errorf(
			`Unexpected concrete type for Item: "%v"`, tn.TypeName)
	}
}

// ItemArticle includes the requested fields of the GraphQL type Article.
type ItemArticle struct {
	Typename string `json:"__typename"`
	// ID is the identifier of the content.
	Id   testutil.ID `json:"id"`
	Name string      `json:"name"`
}

// ItemTopic includes the requested fields of the GraphQL type Topic.
type ItemTopic struct {
	Typename string `json:"__typename"`
	// ID is the identifier of the content.
	Id   testutil.ID `json:"id"`
	Name string      `json:"name"`
}

// ItemVideo includes the requested fields of the GraphQL type Video.
type ItemVideo struct {
	Typename string `json:"__typename"`
	// ID is the identifier of the content.
	Id   testutil.ID `json:"id"`
	Name string      `json:"name"`
}

// Resp is returned by TypeNames on success.
type Resp struct {
	// user looks up a user by some stuff.
	//
	// See UserQueryInput for what stuff is supported.
	// If query is null, returns the current user.
	User       User   `json:"user"`
	RandomItem Item   `json:"-"`
	Users      []User `json:"users"`
}

func (v *Resp) UnmarshalJSON(b []byte) error {

	var firstPass struct {
		*Resp
		RandomItem json.RawMessage `json:"randomItem"`
		graphql.NoUnmarshalJSON
	}
	firstPass.Resp = v

	err := json.Unmarshal(b, &firstPass)
	if err != nil {
		return err
	}

	{
		target := &v.RandomItem
		raw := firstPass.RandomItem
		err = __unmarshalItem(
			target, raw)
		if err != nil {
			return fmt.Errorf(
				"Unable to unmarshal Resp.RandomItem: %w", err)
		}
	}
	return nil
}

// User includes the requested fields of the GraphQL type User.
// The GraphQL type's documentation follows.
//
// A User is a user!
type User struct {
	// id is the user's ID.
	//
	// It is stable, unique, and opaque, like all good IDs.
	Id   testutil.ID `json:"id"`
	Name string      `json:"name"`
}

func TypeNames(
	client graphql.Client,
) (*Resp, error) {
	var err error

	var retval Resp
	err = client.MakeRequest(
		nil,
		"TypeNames",
		`
query TypeNames {
	user {
		id
		name
	}
	randomItem {
		__typename
		id
		name
	}
	users {
		id
		name
	}
}
`,
		&retval,
		nil,
	)
	return &retval, err
}
