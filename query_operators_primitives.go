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
)

// QONotEqual return a not equal condition
func QONotEqual(value interface{}) interface{} {
	return bson.M{"$ne": value}
}

// QOLessThan return a less than condition
func QOLessThan(value interface{}) interface{} {
	return bson.M{"$lt": value}
}

// QOLessThanOrEqual return a less than or equal condition
func QOLessThanOrEqual(value interface{}) interface{} {
	return bson.M{"$lte": value}
}

// QOSize return the size operator
func QOSize(value interface{}) interface{} {
	return bson.M{"$size": value}
}

// QOExpr return the expr operator
func QOExpr(value interface{}) interface{} {
	return bson.M{"$expr": value}
}
