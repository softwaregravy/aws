// Copyright 2012 Aaron Jacobs. All Rights Reserved.
// Author: aaronjjacobs@gmail.com (Aaron Jacobs)
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package conn

import (
	. "github.com/jacobsa/oglematchers"
	. "github.com/jacobsa/ogletest"
	"strings"
	"testing"
)

func TestPostBody(t *testing.T) { RunTests(t) }

////////////////////////////////////////////////////////////////////////
// Helpers
////////////////////////////////////////////////////////////////////////

type PostBodyTest struct {
}

func init() { RegisterTestSuite(&PostBodyTest{}) }

////////////////////////////////////////////////////////////////////////
// Tests
////////////////////////////////////////////////////////////////////////

func (t *PostBodyTest) NoParameters() {
	req := Request{
	}

	body := assemblePostBody(req)

	ExpectEq("", body)
}

func (t *PostBodyTest) OneParameter() {
	req := Request{
		"taco": "burrito",
	}

	body := assemblePostBody(req)

	ExpectEq("taco=burrito", body)
}

func (t *PostBodyTest) MultipleParameters() {
	req := Request{
		"taco": "burrito",
		"enchilada": "queso",
		"nachos": "carnitas",
	}

	body := assemblePostBody(req)
	components := strings.Split(string(body), "&")

	AssertThat(components, ElementsAre(Any(), Any(), Any()))
	ExpectThat(components, Contains("taco=burrito"))
	ExpectThat(components, Contains("enchilada=queso"))
	ExpectThat(components, Contains("nachos=carnitas"))
}

func (t *PostBodyTest) EmptyParameterName() {
	ExpectEq("TODO", "")
}

func (t *PostBodyTest) UnreservedCharacters() {
	ExpectEq("TODO", "")
}

func (t *PostBodyTest) StructuralCharacters() {
	ExpectEq("TODO", "")
}

func (t *PostBodyTest) SpaceAndPlus() {
	ExpectEq("TODO", "")
}

func (t *PostBodyTest) KoreanCharacters() {
	ExpectEq("TODO", "")
}

func (t *PostBodyTest) ParameterOrdering() {
	ExpectEq("TODO", "")
}
