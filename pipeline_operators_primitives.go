// Copyright (c) 2019 Sorint.lab S.p.A.
//
// This program is free software: you can redistribute it and/or modify
// it under the terms of the GNU General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU General Public License for more details.
//
// You should have received a copy of the GNU General Public License
// along with this program.  If not, see <http://www.gnu.org/licenses/>.

package mu

import (
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// APOAdd return a expression that sum the things
func APOAdd(things ...interface{}) interface{} {
	return bson.M{
		"$add": things,
	}
}

// APOSubtract return a expression that subtract a-b
func APOSubtract(a interface{}, b interface{}) interface{} {
	return bson.M{
		"$subtract": bson.A{
			a,
			b,
		},
	}
}

// APODivide return a expression that divide a by b
func APODivide(a interface{}, b interface{}) interface{} {
	return bson.M{
		"$divide": bson.A{
			a,
			b,
		},
	}
}

// APOFloor return a expression that return floor(what)
func APOFloor(what interface{}) interface{} {
	return bson.M{"$floor": what}
}

// APOMin return a expression that return the min of what
func APOMin(what ...interface{}) interface{} {
	return bson.M{"$min": what}
}

// APOMin return a expression that return the max of what
func APOMax(what ...interface{}) interface{} {
	return bson.M{"$max": what}
}

// APOAnd return a expresison that return true if all conds are true
func APOAnd(conds ...interface{}) interface{} {
	return bson.M{
		"$and": conds,
	}
}

// APOOr return a expresison that return true if any conds are true
func APOOr(conds ...interface{}) interface{} {
	return bson.M{
		"$or": conds,
	}
}

// APOEqual return a expression that return true if a and b are equal, otherwise false
func APOEqual(a interface{}, b interface{}) interface{} {
	return bson.M{"$eq": bson.A{a, b}}
}

// APOGreater return a expression that return true if a is greater than b
func APOGreater(a interface{}, b interface{}) interface{} {
	return bson.M{"$gt": bson.A{a, b}}
}

// APORegexFind return a expression that return a regex match
func APORegexFind(input interface{}, regex string, options string) interface{} {
	return bson.M{"$regexFind": bson.M{
		"input":   input,
		"regex":   primitive.Regex{Pattern: regex},
		"options": options,
	}}
}

// APOConcat return a expression that return the concatenation of what
func APOConcat(what ...interface{}) interface{} {
	return bson.M{"$concat": what}
}

// APOSetUnion return a expression that return a set that contains every elements in what
func APOSetUnion(what ...interface{}) interface{} {
	return bson.M{"$setUnion": what}
}

// APOMap return a expression that return a array of mapped value from input to in
func APOMap(input interface{}, as string, in interface{}) interface{} {
	return bson.M{"$map": bson.M{
		"input": input,
		"as":    as,
		"in":    in,
	}}
}

// APOFilter return a expression that return a filtered array of values by cond from input
func APOFilter(input interface{}, as string, cond interface{}) interface{} {
	return bson.M{"$filter": bson.M{
		"input": input,
		"as":    as,
		"cond":  cond,
	}}
}

// APOReduce return a expression that reduce the input array into a value
func APOReduce(input interface{}, initialValue interface{}, in interface{}) interface{} {
	return bson.M{"$reduce": bson.M{
		"input":        input,
		"initialValue": initialValue,
		"in":           in,
	}}
}

// APOSize return a expression that return the size of input
func APOSize(input interface{}) interface{} {
	return bson.M{"$size": input}
}

// APOArrayElemAt return a expression that return the element from the array input with the index index
func APOArrayElemAt(input interface{}, index interface{}) interface{} {
	return bson.M{"$arrayElemAt": bson.A{
		input,
		index,
	}}
}

// APOMergeObjects return a expression that return a merge of what
func APOMergeObjects(what ...interface{}) interface{} {
	return bson.M{"$mergeObjects": what}
}

// APOIfNull return a expression that return what if not null, otherwise altValue
func APOIfNull(what interface{}, altValue interface{}) interface{} {
	return bson.M{"$ifNull": bson.A{
		what,
		altValue,
	}}
}

// APOCond return a expression that return ifTrue if cond is true, otherwise ifFalse
func APOCond(cond interface{}, ifTrue interface{}, ifFalse interface{}) interface{} {
	return bson.M{"$cond": bson.M{
		"if":   cond,
		"then": ifTrue,
		"else": ifFalse,
	}}
}

// APOLet return a expression that calculate the vars and use it in the in expression
func APOLet(vars interface{}, in interface{}) interface{} {
	return bson.M{"$let": bson.M{
		"vars": vars,
		"in":   in,
	}}
}

// APODateFromString return a expression that parse the what into a date
func APODateFromString(what interface{}, format string) interface{} {
	return bson.M{"$dateFromString": bson.M{
		"dateString": what,
		"format":     format,
	}}
}

// APODateFromNullableString return a expression that parse the what into a date
func APODateFromNullableString(what interface{}, format string, onNull interface{}) interface{} {
	return bson.M{"$dateFromString": bson.M{
		"dateString": what,
		"format":     format,
		"onNull":     onNull,
	}}
}

// APOConvert return a expression that convert input to to
func APOConvert(input interface{}, to string) interface{} {
	return bson.M{"$convert": bson.M{
		"input": input,
		"to":    to,
	}}
}

// APOConvertErrorable return a expression that convert input to to
func APOConvertErrorable(input interface{}, to string, onError interface{}) interface{} {
	return bson.M{"$convert": bson.M{
		"input":   input,
		"to":      to,
		"onError": onError,
	}}
}

// APOConvertNullable return a expression that convert input to to
func APOConvertNullable(input interface{}, to string, onNull interface{}) interface{} {
	return bson.M{"$convert": bson.M{
		"input":  input,
		"to":     to,
		"onNull": onNull,
	}}
}

// APOConvertErrorableNullable return a expression that convert input to to
func APOConvertErrorableNullable(input interface{}, to string, onError interface{}, onNull interface{}) interface{} {
	return bson.M{"$convert": bson.M{
		"input":   input,
		"to":      to,
		"onError": onError,
		"onNull":  onNull,
	}}
}

// APOToDouble return a expression that convert input to double
func APOToDouble(input interface{}) interface{} {
	return bson.M{"$toDouble": input}
}

// APOSum return a summing expression of whats
func APOSum(what interface{}) interface{} {
	return bson.M{"$sum": what}
}

// APOMaxAggr return a maximizing expression of whats
func APOMaxAggr(what interface{}) interface{} {
	return bson.M{"$max": what}
}
