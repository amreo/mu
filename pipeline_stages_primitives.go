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

// MAPipeline return a aggregation pipeline joining the stages that could be a single stage or a slice of multiple stages
func MAPipeline(stages ...interface{}) bson.A {
	out := bson.A{}
	for _, stage := range stages {
		if stage == nil {
			continue
		} else if reflect.TypeOf(stage).Kind() == reflect.Slice {
			for _, item := range sliceToSliceOfInterface(stage) {
				out = append(out, item)
			}
		} else {
			out = append(out, stage)
		}
	}

	return out
}

// APOptionalStage return the stage if optional is true, otherwise return a null stage
func APOptionalStage(optional bool, stage interface{}) interface{} {
	if optional {
		return stage
	}
	return nil
}

// APMatch return a match stage
func APMatch(conditions interface{}) interface{} {
	return bson.M{"$match": conditions}
}

// APSort return a sort stage
func APSort(what interface{}) interface{} {
	return bson.M{"$sort": what}
}

// APLimit return a limit stage
func APLimit(number interface{}) interface{} {
	return bson.M{"$limit": number}
}

// APSkip return a skip stage
func APSkip(number interface{}) interface{} {
	return bson.M{"$skip": number}
}

// APProject return a project stage
func APProject(what interface{}) interface{} {
	return bson.M{"$project": what}
}

// APUnset return a unset stage
func APUnset(what ...interface{}) interface{} {
	return bson.M{"$unset": what}
}

// APSet return a set stage
func APSet(what interface{}) interface{} {
	return bson.M{"$set": what}
}

// APAddFields return a addFields stage
func APAddFields(what interface{}) interface{} {
	return bson.M{"$addFields": what}
}

// APUnwind return a unwind stage
func APUnwind(what string) interface{} {
	return bson.M{"$unwind": what}
}

// APReplaceWith return a replaceWith stage
func APReplaceWith(what interface{}) interface{} {
	return bson.M{"$replaceWith": what}
}

// APLookupSimple return a lookup stage
func APLookupSimple(from string, localField string, foreignField string, as string) interface{} {
	return bson.M{"$lookup": bson.M{
		"from":         from,
		"localField":   localField,
		"foreignField": foreignField,
		"as":           as,
	}}
}

// APLookupPipeline return a lookup stage
func APLookupPipeline(from string, let interface{}, as string, pipeline interface{}) interface{} {
	return bson.M{"$lookup": bson.M{
		"from":     from,
		"let":      let,
		"as":       as,
		"pipeline": pipeline,
	}}
}

// APGroup return a group stage
func APGroup(fields interface{}) interface{} {
	return bson.M{"$group": fields}
}

// APFacet return a facet stage
func APFacet(fields interface{}) interface{} {
	return bson.M{"$facet": fields}
}

// APCount return a count stage
func APCount(field string) interface{} {
	return bson.M{"$count": field}
}
