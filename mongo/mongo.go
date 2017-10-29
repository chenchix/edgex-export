//
// Copyright 2017 Mainflux.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//

package mongo

import (
	"gopkg.in/mgo.v2"
)

const (
	DbName         string = "coredata"
	CollectionName string = "exportConfiguration"
)

type MongoRepository struct {
	Session *mgo.Session
}

func NewMongoRepository(ms *mgo.Session) *MongoRepository {
	return &MongoRepository{Session: ms}
}
