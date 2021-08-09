// Code generated by entc, DO NOT EDIT.

package ent

import (
	"entprom/internal/ent/file"
	"entprom/internal/ent/schema"
)

// The init function reads all schema descriptors with runtime code
// (default values, validators, hooks and policies) and stitches it
// to their package variables.
func init() {
	fileFields := schema.File{}.Fields()
	_ = fileFields
	// fileDescDeleted is the schema descriptor for deleted field.
	fileDescDeleted := fileFields[1].Descriptor()
	// file.DefaultDeleted holds the default value on creation for the deleted field.
	file.DefaultDeleted = fileDescDeleted.Default.(bool)
}