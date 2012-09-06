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

package s3

import (
	"errors"
	"github.com/jacobsa/aws/s3/auth/mock"
	"github.com/jacobsa/aws/s3/http"
	"github.com/jacobsa/aws/s3/http/mock"
	. "github.com/jacobsa/oglematchers"
	"github.com/jacobsa/oglemock"
	. "github.com/jacobsa/ogletest"
	"strings"
	"testing"
	"time"
)

func TestBucket(t *testing.T) { RunTests(t) }

////////////////////////////////////////////////////////////////////////
// Helpers
////////////////////////////////////////////////////////////////////////

type fakeClock struct {
	now time.Time
}

func (c *fakeClock) Now() time.Time {
	return c.now
}

type bucketTest struct {
	httpConn mock_http.MockConn
	signer   mock_auth.MockSigner
	bucket   Bucket
	clock    *fakeClock
}

func (t *bucketTest) SetUp(i *TestInfo) {
	var err error

	t.httpConn = mock_http.NewMockConn(i.MockController, "httpConn")
	t.signer = mock_auth.NewMockSigner(i.MockController, "signer")
	t.clock = &fakeClock{}

	t.bucket, err = openBucket("some.bucket", t.httpConn, t.signer, t.clock)
	AssertEq(nil, err)
}

////////////////////////////////////////////////////////////////////////
// GetObject
////////////////////////////////////////////////////////////////////////

type GetObjectTest struct {
	bucketTest
}

func init() { RegisterTestSuite(&GetObjectTest{}) }

func (t *GetObjectTest) DoesFoo() {
	ExpectEq("TODO", "")
}

////////////////////////////////////////////////////////////////////////
// StoreObject
////////////////////////////////////////////////////////////////////////

type StoreObjectTest struct {
	bucketTest
}

func init() { RegisterTestSuite(&StoreObjectTest{}) }

func (t *StoreObjectTest) KeyNotValidUtf8() {
	key := "\x80\x81\x82"
	data := []byte{}

	// Call
	err := t.bucket.StoreObject(key, data)

	ExpectThat(err, Error(HasSubstr("valid")))
	ExpectThat(err, Error(HasSubstr("UTF-8")))
}

func (t *StoreObjectTest) KeyTooLong() {
	key := strings.Repeat("a", 1025)
	data := []byte{}

	// Call
	err := t.bucket.StoreObject(key, data)

	ExpectThat(err, Error(HasSubstr("1024")))
	ExpectThat(err, Error(HasSubstr("bytes")))
}

func (t *StoreObjectTest) CallsSigner() {
	key := "foo/bar/baz"
	data := []byte{0x00, 0xde, 0xad, 0xbe, 0xef}

	// Clock
	t.clock.now = time.Date(1985, time.March, 18, 15, 33, 17, 123, time.Local)

	// Signer
	var httpReq *http.Request
	ExpectCall(t.signer, "Sign")(Any()).
		WillOnce(oglemock.Invoke(func(r *http.Request) error {
		httpReq = r
		return errors.New("")
	}))

	// Call
	t.bucket.StoreObject(key, data)

	AssertNe(nil, httpReq)
	ExpectEq("PUT", httpReq.Verb)
	ExpectEq("/some.bucket/foo/bar/baz", httpReq.Path)
	ExpectEq("Mon, 18 Mar 1985 123:33:17 UTC", httpReq.Headers["Date"])
	ExpectThat(httpReq.Body, DeepEquals(data))
}

func (t *StoreObjectTest) SignerReturnsError() {
	key := ""
	data := []byte{}

	// Signer
	ExpectCall(t.signer, "Sign")(Any()).
		WillOnce(oglemock.Return(errors.New("taco")))

	// Call
	err := t.bucket.StoreObject(key, data)

	ExpectThat(err, Error(HasSubstr("Sign")))
	ExpectThat(err, Error(HasSubstr("taco")))
}

func (t *StoreObjectTest) CallsConn() {
	ExpectEq("TODO", "")
}

func (t *StoreObjectTest) ConnReturnsError() {
	ExpectEq("TODO", "")
}

func (t *StoreObjectTest) ServerReturnsError() {
	ExpectEq("TODO", "")
}

func (t *StoreObjectTest) ServerSaysOkay() {
	ExpectEq("TODO", "")
}
