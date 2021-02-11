package util

import (
	"strconv"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// StringToInt :nodoc:
func StringToInt(value *int, param string, defaultValue ...int) {
	res, err := strconv.Atoi(param)
	if err != nil {
		*value = defaultValue[0]
	} else {
		*value = res
	}
}

// StringToTime :nodoc:
func StringToTime(value *time.Time, param string) {
	t, err := time.Parse(time.RFC3339, param)
	if err != nil {
		*value = time.Date(1900, time.January, 1, 0, 0, 0, 0, time.UTC)
	} else {
		*value = t
	}
}

// StringToObjectID :nodoc:
func StringToObjectID(value *primitive.ObjectID, param string) error {
	id, err := primitive.ObjectIDFromHex(param)
	if err != nil {
		return err
	}
	*value = id
	return nil
}
