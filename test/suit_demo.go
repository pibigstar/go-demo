package test

import (
	"testing"
	"github.com/bzy-ai/rpc-mesh/utils"
	"github.com/stretchr/testify/suite"
)

type StringUtilSuit struct {
	suite.Suite
}

func TestStringUtilSuit(t *testing.T) {
	s := &StringUtilSuit{}

	suite.Run(t, s)
}

func (st *StringUtilSuit) TestIsContainUrl() {
	isContainUrl := "hello, https://bzy.ai"
	st.Equal(true, utils.IsContainUrl(isContainUrl))

	notContainUrl := "hello, bzy ai"
	st.Equal(false, utils.IsContainUrl(notContainUrl))
}
