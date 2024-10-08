// Code generated by ent, DO NOT EDIT.

package ent

import (
	"next_go_blog/ent/article"
	"next_go_blog/ent/schema"
	"next_go_blog/ent/studytime"
	"next_go_blog/ent/user"
	"time"
)

// The init function reads all schema descriptors with runtime code
// (default values, validators, hooks and policies) and stitches it
// to their package variables.
func init() {
	articleFields := schema.Article{}.Fields()
	_ = articleFields
	// articleDescUserID is the schema descriptor for user_id field.
	articleDescUserID := articleFields[0].Descriptor()
	// article.UserIDValidator is a validator for the "user_id" field. It is called by the builders before save.
	article.UserIDValidator = articleDescUserID.Validators[0].(func(int) error)
	// articleDescTitle is the schema descriptor for title field.
	articleDescTitle := articleFields[1].Descriptor()
	// article.TitleValidator is a validator for the "title" field. It is called by the builders before save.
	article.TitleValidator = articleDescTitle.Validators[0].(func(string) error)
	// articleDescBody is the schema descriptor for body field.
	articleDescBody := articleFields[2].Descriptor()
	// article.BodyValidator is a validator for the "body" field. It is called by the builders before save.
	article.BodyValidator = articleDescBody.Validators[0].(func(string) error)
	// articleDescCreatedAt is the schema descriptor for created_at field.
	articleDescCreatedAt := articleFields[3].Descriptor()
	// article.DefaultCreatedAt holds the default value on creation for the created_at field.
	article.DefaultCreatedAt = articleDescCreatedAt.Default.(func() time.Time)
	// articleDescUpdatedAt is the schema descriptor for updated_at field.
	articleDescUpdatedAt := articleFields[4].Descriptor()
	// article.DefaultUpdatedAt holds the default value on creation for the updated_at field.
	article.DefaultUpdatedAt = articleDescUpdatedAt.Default.(func() time.Time)
	// article.UpdateDefaultUpdatedAt holds the default value on update for the updated_at field.
	article.UpdateDefaultUpdatedAt = articleDescUpdatedAt.UpdateDefault.(func() time.Time)
	studytimeFields := schema.StudyTime{}.Fields()
	_ = studytimeFields
	// studytimeDescUserID is the schema descriptor for user_id field.
	studytimeDescUserID := studytimeFields[0].Descriptor()
	// studytime.UserIDValidator is a validator for the "user_id" field. It is called by the builders before save.
	studytime.UserIDValidator = studytimeDescUserID.Validators[0].(func(int) error)
	// studytimeDescTitle is the schema descriptor for title field.
	studytimeDescTitle := studytimeFields[1].Descriptor()
	// studytime.TitleValidator is a validator for the "title" field. It is called by the builders before save.
	studytime.TitleValidator = studytimeDescTitle.Validators[0].(func(string) error)
	// studytimeDescCreatedAt is the schema descriptor for created_at field.
	studytimeDescCreatedAt := studytimeFields[4].Descriptor()
	// studytime.DefaultCreatedAt holds the default value on creation for the created_at field.
	studytime.DefaultCreatedAt = studytimeDescCreatedAt.Default.(func() time.Time)
	// studytimeDescUpdatedAt is the schema descriptor for updated_at field.
	studytimeDescUpdatedAt := studytimeFields[5].Descriptor()
	// studytime.DefaultUpdatedAt holds the default value on creation for the updated_at field.
	studytime.DefaultUpdatedAt = studytimeDescUpdatedAt.Default.(func() time.Time)
	// studytime.UpdateDefaultUpdatedAt holds the default value on update for the updated_at field.
	studytime.UpdateDefaultUpdatedAt = studytimeDescUpdatedAt.UpdateDefault.(func() time.Time)
	userFields := schema.User{}.Fields()
	_ = userFields
	// userDescEmail is the schema descriptor for email field.
	userDescEmail := userFields[0].Descriptor()
	// user.EmailValidator is a validator for the "email" field. It is called by the builders before save.
	user.EmailValidator = userDescEmail.Validators[0].(func(string) error)
	// userDescPassword is the schema descriptor for password field.
	userDescPassword := userFields[1].Descriptor()
	// user.PasswordValidator is a validator for the "password" field. It is called by the builders before save.
	user.PasswordValidator = userDescPassword.Validators[0].(func(string) error)
	// userDescName is the schema descriptor for name field.
	userDescName := userFields[2].Descriptor()
	// user.NameValidator is a validator for the "name" field. It is called by the builders before save.
	user.NameValidator = userDescName.Validators[0].(func(string) error)
	// userDescCreatedAt is the schema descriptor for created_at field.
	userDescCreatedAt := userFields[3].Descriptor()
	// user.DefaultCreatedAt holds the default value on creation for the created_at field.
	user.DefaultCreatedAt = userDescCreatedAt.Default.(func() time.Time)
	// userDescUpdatedAt is the schema descriptor for updated_at field.
	userDescUpdatedAt := userFields[4].Descriptor()
	// user.DefaultUpdatedAt holds the default value on creation for the updated_at field.
	user.DefaultUpdatedAt = userDescUpdatedAt.Default.(func() time.Time)
	// user.UpdateDefaultUpdatedAt holds the default value on update for the updated_at field.
	user.UpdateDefaultUpdatedAt = userDescUpdatedAt.UpdateDefault.(func() time.Time)
}
