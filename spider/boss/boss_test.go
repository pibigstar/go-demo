package main

import (
	"context"
	"testing"
)

func TestMain(m *testing.M) {
	m.Run()
}

// 测试招人
func TestHiring(t *testing.T) {
	for jobId, jobName := range jobIds {
		Hiring(jobId, jobName)
	}
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

func TestGetQR(t *testing.T) {
	getQRId(context.Background())
}

func TestListJob(t *testing.T) {
	jobs := listJobs()
	for _, j := range jobs {
		t.Log(j.JobName, j.JobId)
	}
}

func TestInputJobs(t *testing.T) {
	inputJobs()
}

func TestSendFeiShu(t *testing.T) {
	sendFeiShu("Boss当前登录登录状态失效")
}
