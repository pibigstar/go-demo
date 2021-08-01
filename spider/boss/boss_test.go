package main

import (
	"testing"
)

// 测试招人
func TestHiring(t *testing.T) {
	Hiring()
}

func TestReadSchool(t *testing.T) {
	readSchool()
	for _, s := range school985 {
		t.Log(s)
	}

	for _, s := range school211 {
		t.Log(s)
	}
}

func TestReadJob(t *testing.T) {
	readJobs()
	for jobId, jobName := range jobIds {
		t.Log(jobId, jobName)
	}
}

func TestListRecommend(t *testing.T) {
	for jobId := range jobIds {
		geeks, err := listRecommend(jobId)
		if err != nil {
			t.Error(err)
		}
		for _, geek := range geeks {
			t.Log(geek.GeekCard.GeekName)
		}
	}
}

func TestSetHelloMsg(t *testing.T) {
	setHelloMsg()
}

func TestReadCompany(t *testing.T) {
	readCompany()
	for _, c := range goodCompany {
		t.Log(c)
	}
}

func TestSendEmail(t *testing.T) {
	sendEmail()
}
