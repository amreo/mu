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
			"metadata": MAPipeline(
				APCount("totalElements"),
			),
			"content": MAPipeline(
				APSkip(page*size),
				APLimit(size),
			),
		}),
		APSet(bson.M{
			"metadata": APOIfNull(APOArrayElemAt("$metadata", 0), bson.M{
				"totalElements": 0,
			}),
		}),
		APSet(bson.M{
			"metadata.totalPages": "$metadata",
		}),
		APAddFields(bson.M{
			"metadata.totalPages": APOFloor(APODivide("$metadata.totalElements", size)),
			"metadata.size":       APOMin(size, APOSubtract("$metadata.totalElements", size*page)),
			"metadata.number":     page,
		}),
		APAddFields(bson.M{
			"metadata.empty": APOEqual("$metadata.size", 0),
			"metadata.first": page == 0,
			"metadata.last":  APOGreaterOrEqual(page, APOSubtract("$metadata.totalPages", 1)),
		}),
	)
}

// APSearchFilterStage return a aggregation stage that filter the documents when any field match any keyword
func APSearchFilterStage(fields []interface{}, keywords []string) interface{} {
	//Build the search pattern
	quotedKeywords := []string{}
	for _, k := range keywords {
		quotedKeywords = append(quotedKeywords, regexp.QuoteMeta(k))
	}

	//Build the $or conditions
	conditions := []interface{}{}
	for _, q := range quotedKeywords {
		matchKeywordConditions := []interface{}{}
		for _, f := range fields {
			matchKeywordConditions = append(matchKeywordConditions,
				APOCond(
					APOEqual(bson.M{
						"$type": f,
					}, "array"),
					bson.M{
						"$or": APOReduce(f, false,
							APOOr(
								"$$value",
								bson.M{
									"$regexMatch": bson.M{
										"input": "$$this",
										"regex": primitive.Regex{Pattern: q, Options: "i"},
									},
								},
							),
						),
					},
					bson.M{
						"$regexMatch": bson.M{
							"input": f,
							"regex": primitive.Regex{Pattern: q, Options: "i"},
						},
					},
				),
			)
		}

		conditions = append(conditions, APOOr(matchKeywordConditions...))
	}

	// Return the matching stage
	return APMatch(QOExpr(APOAnd(conditions...)))
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
