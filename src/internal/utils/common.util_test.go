package utils

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"strconv"
	"testing"
	"time"
)

type UtilTest struct {
	suite.Suite
}

func TestUtil(t *testing.T) {
	suite.Run(t, new(UtilTest))
}

func (t *UtilTest) TestGetStudyYearSuccess() {
	currentYear := time.Now().Year()

	testGetStudyYearSuccess(t.T(), "62xxxxxxxx", strconv.Itoa(currentYear-2019))
	testGetStudyYearSuccess(t.T(), "63xxxxxxxx", strconv.Itoa(currentYear-2020))
	testGetStudyYearSuccess(t.T(), "64xxxxxxxx", strconv.Itoa(currentYear-2021))
	testGetStudyYearSuccess(t.T(), "65xxxxxxxx", strconv.Itoa(currentYear-2022))
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
