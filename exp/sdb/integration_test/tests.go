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

package main

import (
	"github.com/jacobsa/aws/exp/sdb"
	. "github.com/jacobsa/ogletest"
	"sync"
)

////////////////////////////////////////////////////////////////////////
// Helpers
////////////////////////////////////////////////////////////////////////

type DomainTest struct {
	domain sdb.Domain

	mutex        sync.Mutex
	itemsToDelete []sdb.ItemName
}

func init() { RegisterTestSuite(&DomainTest{}) }

// Ensure that the given item is deleted before the test finishes.
func (t *DomainTest) ensureDeleted(item sdb.ItemName) {
	t.mutex.Lock()
	defer t.mutex.Unlock()

	t.itemsToDelete = append(t.itemsToDelete, item)
}

func (t *DomainTest) SetUp(i *TestInfo) {
	var err error

	// Open a domain.
	t.domain, err = sdb.OpenDomain(*g_domainName, sdb.Region(*g_region), g_accessKey)
	AssertEq(nil, err)
}

func (t *DomainTest) TearDown() {
	t.mutex.Lock()
	defer t.mutex.Unlock()

	AssertEq("", "TODO: Delete items from itemsToDelete")
}

////////////////////////////////////////////////////////////////////////
// Tests
////////////////////////////////////////////////////////////////////////

func (t *DomainTest) WrongAccessKeySecret() {
	ExpectEq("TODO", "")
}

func (t *DomainTest) InvalidUtf8ItemName() {
	ExpectEq("TODO", "")
}

func (t *DomainTest) InvalidUtf8AttributeName() {
	ExpectEq("TODO", "")
}

func (t *DomainTest) InvalidUtf8AttributeValue() {
	ExpectEq("TODO", "")
}

func (t *DomainTest) LongItemName() {
	ExpectEq("TODO", "")
}

func (t *DomainTest) LongAttributeName() {
	ExpectEq("TODO", "")
}

func (t *DomainTest) LongAttributeValue() {
	ExpectEq("TODO", "")
}

func (t *DomainTest) PutThenGet() {
	ExpectEq("TODO", "")
}

func (t *DomainTest) BatchPutThenGet() {
	ExpectEq("TODO", "")
}

func (t *DomainTest) BatchPutThenBatchGet() {
	ExpectEq("TODO", "")
}

func (t *DomainTest) GetForNonExistentItem() {
	ExpectEq("TODO", "")
}

func (t *DomainTest) GetParticularAttributes() {
	ExpectEq("TODO", "")
}

func (t *DomainTest) BatchGetParticularAttributes() {
	ExpectEq("TODO", "")
}

func (t *DomainTest) BatchGetForNonExistentItems() {
	ExpectEq("TODO", "")
}

func (t *DomainTest) GetNonExistentAttributeName() {
	ExpectEq("TODO", "")
}

func (t *DomainTest) BatchGetNonExistentAttributeName() {
	ExpectEq("TODO", "")
}

func (t *DomainTest) FailedValuePrecondition() {
	ExpectEq("TODO", "")
}

func (t *DomainTest) FailedExistencePrecondition() {
	ExpectEq("TODO", "")
}

func (t *DomainTest) FailedNonExistencePrecondition() {
	ExpectEq("TODO", "")
}

func (t *DomainTest) SuccessfulPreconditions() {
	ExpectEq("TODO", "")
}

func (t *DomainTest) DeleteParticularAttributes() {
	ExpectEq("TODO", "")
}

func (t *DomainTest) DeleteAllAttributes() {
	ExpectEq("TODO", "")
}

func (t *DomainTest) BatchDelete() {
	ExpectEq("TODO", "")
}

func (t *DomainTest) InvalidSelectQuery() {
	ExpectEq("TODO", "")
}

func (t *DomainTest) SelectAll() {
	ExpectEq("TODO", "")
}

func (t *DomainTest) SelectItemName() {
	ExpectEq("TODO", "")
}

func (t *DomainTest) SelectCount() {
	ExpectEq("TODO", "")
}

func (t *DomainTest) SelectWithPredicates() {
	ExpectEq("TODO", "")
}

func (t *DomainTest) SelectWithSortOrder() {
	ExpectEq("TODO", "")
}

func (t *DomainTest) SelectWithLimit() {
	ExpectEq("TODO", "")
}

func (t *DomainTest) SelectEmptyResultSet() {
	ExpectEq("TODO", "")
}

func (t *DomainTest) SelectLargeResultSet() {
	ExpectEq("TODO", "")
}