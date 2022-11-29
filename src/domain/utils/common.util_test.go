package utils

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"testing"
)

type UtilTest struct {
	suite.Suite
}

func TestUtil(t *testing.T) {
	suite.Run(t, new(UtilTest))
}

func (t *UtilTest) TestGetStudyYearSuccess() {
	testGetStudyYearSuccess(t.T(), "62xxxxxxxx", "4")
	testGetStudyYearSuccess(t.T(), "63xxxxxxxx", "3")
	testGetStudyYearSuccess(t.T(), "64xxxxxxxx", "2")
	testGetStudyYearSuccess(t.T(), "65xxxxxxxx", "1")
}

func testGetStudyYearSuccess(t *testing.T, sid string, expect string) {
	want := expect

	actual, err := CalYearFromID(sid)

	assert.Nil(t, err)
	assert.Equal(t, want, actual)
}

func (t *UtilTest) TestCalStudyYearInvalidFormat() {
	testCalStudyYearInvalidInput(t.T(), "")
	testCalStudyYearInvalidInput(t.T(), "65xxx")
	testCalStudyYearInvalidInput(t.T(), "xx24xxxxxx")
	testCalStudyYearInvalidInput(t.T(), "65xxxxxxxxxxx")
}

func (t *UtilTest) TestCalStudyYearInvalidYear() {
	testCalStudyYearInvalidInput(t.T(), "66xxxxxxxxxxx")
	testCalStudyYearInvalidInput(t.T(), "68xxxxxxxxxxx")
	testCalStudyYearInvalidInput(t.T(), "99xxxxxxxxxxx")
}

func testCalStudyYearInvalidInput(t *testing.T, sid string) {
	want := "Invalid student id"

	actual, err := CalYearFromID(sid)

	assert.Equal(t, actual, "")
	assert.Equal(t, want, err.Error())
}
