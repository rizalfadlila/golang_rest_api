package helpers

import (
	"time"

	"github.com/rest_api/pkg/util"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// SetValueFilter :noodc:
func SetValueFilter(params map[string]interface{}, key string, filters *bson.D) {
	if params[key] != nil && params[key] != "" {
		*filters = append(*filters, bson.E{key, params[key]})
	}
}

// SetValueFilterPattern :nodoc:
func SetValueFilterPattern(params map[string]interface{}, key string, filters *bson.D) {
	if params[key] != nil && params[key] != "" {
		*filters = append(*filters, bson.E{key, primitive.Regex{Pattern: params[key].(string), Options: "i"}})
	}
}

// SetValueFilterNumber :nodoc:
func SetValueFilterNumber(params map[string]interface{}, key string, filters *bson.D) {
	if params[key] != nil && params[key] != "" {
		var number int
		util.StringToInt(&number, params[key].(string), 0)
		*filters = append(*filters, bson.E{key, number})
	}
}

// SetValueFilterTimeBetween :nodoc:
func SetValueFilterTimeBetween(params map[string]interface{}, key, val1, val2 string, filters *bson.D) {
	if params[val1] != nil && params[val1] != "" && params[val2] != nil && params[val2] != "" {
		var fromDate time.Time
		var toDate time.Time

		util.StringToTime(&fromDate, params[val1].(string))
		util.StringToTime(&toDate, params[val2].(string))

		*filters = append(*filters, bson.E{key, bson.M{
			"$gte": fromDate,
			"$lte": toDate,
		}})
	}
}

// SetValueFilterOr format [key];[type (string or int or double or etc)] :nodoc:
func SetValueFilterOr(keys []string, value interface{}, filters *bson.D) {
	if value != "" {
		var fields []interface{}
		for i := 0; i < len(keys); i++ {
			fields = append(fields, bson.D{{keys[i], primitive.Regex{Pattern: value.(string), Options: "i"}}})
		}
		*filters = append(*filters, bson.E{"$or", fields})
	}
}

// SetValueFilterIn key must be using plurar
func SetValueFilterIn(params map[string]interface{}, key string, filters *bson.D) {
	if params[key] != nil && params[key] != "" {
		// remove plural last word, ex: ids -> id
		field := key[:len(key)-1]
		*filters = append(*filters, bson.E{field, bson.M{
			"$in": params[key],
		}})
	}
}
