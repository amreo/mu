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
	"reflect"

	"go.mongodb.org/mongo-driver/bson"
)

// MAPipeline return a aggregation pipeline joining the steps that could be a single step or a slice of steps
func MAPipeline(steps ...interface{}) bson.A {
	out := bson.A{}
	for _, step := range steps {
		if reflect.TypeOf(step).Kind() == reflect.Slice {
			for _, item := range sliceToSliceOfInterface(step) {
				out = append(out, item)
			}
		} else {
			out = append(out, step)
		}
	}

	return out
}
