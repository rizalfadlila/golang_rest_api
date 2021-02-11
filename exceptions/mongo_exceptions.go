package exceptions

import "errors"

// ErrMongoNoDocument :nodoc:
var ErrMongoNoDocument = errors.New("mongo: no documents in result")

// ErrMongoInvalidObjectID :nodoc:
var ErrMongoInvalidObjectID = errors.New("the provided hex string is not a valid ObjectID")

// ErrMongoEncodingHEX :nodoc:
var ErrMongoEncodingHEX = errors.New("encoding/hex")
