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

// BsonOptionalExtension return a bson with the same key-values pairs as the orig and extension, if extend is true, otherwise return the orig bson
func BsonOptionalExtension(extend bool, orig bson.M, extension bson.M) bson.M {
	if extend {
		for k, v := range extension {
			orig[k] = v
		}
	}

	return orig
}
