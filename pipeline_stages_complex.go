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
	"regexp"
	"strings"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// APOptionalSortingStage return a stage that sort documents by the criteria in the params
func APOptionalSortingStage(sortBy string, sortDesc bool) interface{} {
	if sortBy == "" {
		return nil
	}

	sortOrder := 0
	if sortDesc {
		sortOrder = -1
	} else {
		sortOrder = 1
	}

	return APSort(bson.M{
		sortBy: sortOrder,
	})
}

// APOptionalPagingStage return a stage that turn a stream of documents into a page that contains the documents plus some metadata
func APOptionalPagingStage(page int, size int) interface{} {
	if page == -1 || size == -1 {
		return nil
	}

	return MAPipeline(
		APFacet(bson.M{
			"Metadata": MAPipeline(
				APCount("TotalElements"),
			),
			"Content": MAPipeline(
				APSkip(page*size),
				APLimit(size),
			),
		}),
		APSet(bson.M{
			"Metadata": APOIfNull(APOArrayElemAt("$Metadata", 0), bson.M{
				"TotalElements": 0,
			}),
		}),
		APSet(bson.M{
			"Metadata.TotalPages": "$Metadata",
		}),
		APAddFields(bson.M{
			"Metadata.TotalPages": APOFloor(APODivide("$Metadata.TotalElements", size)),
			"Metadata.Size":       APOMin(size, APOSubtract("$Metadata.TotalElements", size*page)),
			"Metadata.Number":     page,
		}),
		APAddFields(bson.M{
			"Metadata.Empty": APOEqual("$Metadata.Size", 0),
			"Metadata.First": page == 0,
			"Metadata.Last":  APOEqual(page, APOSubtract("$Metadata.TotalPages", 1)),
		}),
	)
}

// APSearchFilterStage return a aggregation stage that filter the documents when any field match any keyword
func APSearchFilterStage(fields []string, keywords []string) interface{} {
	//Build the search pattern
	quotedKeywords := []string{}
	for _, k := range keywords {
		quotedKeywords = append(quotedKeywords, regexp.QuoteMeta(k))
	}
	pattern := strings.Join(quotedKeywords, "|")

	//Build the $or conditions
	conditions := []interface{}{}
	for _, f := range fields {
		conditions = append(conditions, bson.M{f: bson.M{
			"$regex": primitive.Regex{Pattern: pattern, Options: "i"},
		}})
	}

	//Return the matching stage
	return APMatch(APOOr(conditions...))
}

// APGroupAndCountStages return some aggregation stagess that group whatFieldName by what and count the documents
func APGroupAndCountStages(whatFieldName string, countFieldName string, what interface{}) interface{} {
	return bson.A{
		APGroup(bson.M{
			"_id":          what,
			countFieldName: APOSum(1),
		}),
		APProject(bson.M{
			"_id":          false,
			whatFieldName:  "$_id",
			countFieldName: true,
		}),
	}
}
