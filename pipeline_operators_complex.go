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

// APOConvertToDoubleOrZero return a expression that convert what to double if it's valid or return zero if invalid or null
func APOConvertToDoubleOrZero(what interface{}) interface{} {
	return APOConvertErrorableNullable(what, "double", 0, 0)
}

//APOJoin return a expression that join the list into a string
func APOJoin(list interface{}, sep interface{}) interface{} {
	return APOReduce(list, "",
		APOConcat("$$value", APOCond(APOEqual("$$value", ""), "", sep), "$$this"),
	)
}

// APOMaxWithCmpExpr return a expression that return the maximium between a and b using the cmpExpressions
func APOMaxWithCmpExpr(cmpExprA interface{}, cmpExprB interface{}, a interface{}, b interface{}) interface{} {
	return APOCond(APOGreater(cmpExprA, cmpExprB), a, b)
}

// APOAny return a expression that return true if any element in input satisfy the cond
// (in the cond the variable this can be used to refer to the current item of the array)
func APOAny(input interface{}, itemName string, cond interface{}) interface{} {
	return APOGreater(APOSize(APOFilter(input, itemName, cond)), 0)
}

// APOGetCaptureFromRegexMatch return a capture group from a regex match of given input and regex
func APOGetCaptureFromRegexMatch(input interface{}, regex string, options string, captureIndex int) {
	return mu.APOLet(
		bson.M{
			"match": mu.APORegexFind(input, regex, options),
		},
		mu.APOConvertToDoubleOrZero(mu.APOArrayElemAt("$$match.captures", captureIndex)),
	),
}
