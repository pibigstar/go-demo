package main

type GeekListResp struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	ZpData  struct {
		NeedGuide   bool    `json:"needGuide"`
		StartIndex  int     `json:"startIndex"`
		HasMore     bool    `json:"hasMore"`
		Page        int     `json:"page"`
		Tag         int     `json:"tag"`
		ItemCardIdx int     `json:"itemCardIdx"`
		GeekList    []*Geek `json:"geekList"`
		TagDesc     string  `json:"tagDesc"`
	} `json:"zpData"`
}

type Geek struct {
	GeekCard struct {
		SecurityID   string `json:"securityId"`
		GeekID       int    `json:"geekId"`
		GeekName     string `json:"geekName"`
		GeekAvatar   string `json:"geekAvatar"`
		GeekGender   int    `json:"geekGender"`
		GeekWorkYear string `json:"geekWorkYear"`
		GeekDegree   string `json:"geekDegree"`
		GeekDesc     struct {
			Content   string        `json:"content"`
			IndexList []interface{} `json:"indexList"`
		} `json:"geekDesc"`
		ExpectID      int    `json:"expectId"`
		ExpectType    int    `json:"expectType"`
		Salary        string `json:"salary"`
		MiddleContent struct {
			Content   string        `json:"content"`
			IndexList []interface{} `json:"indexList"`
		} `json:"middleContent"`
		JobID          int    `json:"jobId"`
		EncryptJobID   string `json:"encryptJobId"`
		Lid            string `json:"lid"`
		SelectableJob  bool   `json:"selectableJob"`
		ActionDateDesc string `json:"actionDateDesc"`
		ActionDate     int64  `json:"actionDate"`
		GeekWorks      []struct {
			ID                    int         `json:"id"`
			GeekID                int         `json:"geekId"`
			Company               string      `json:"company"`
			IndustryCode          int         `json:"industryCode"`
			Industry              interface{} `json:"industry"`
			IndustryCategory      interface{} `json:"industryCategory"`
			Position              int         `json:"position"`
			PositionCategory      interface{} `json:"positionCategory"`
			BlueCollarPosition    bool        `json:"blueCollarPosition"`
			PositionName          string      `json:"positionName"`
			PositionLv2           int         `json:"positionLv2"`
			IsPublic              int         `json:"isPublic"`
			Department            interface{} `json:"department"`
			Responsibility        interface{} `json:"responsibility"`
			StartDate             string      `json:"startDate"`
			EndDate               string      `json:"endDate"`
			CustomPositionID      int         `json:"customPositionId"`
			CustomIndustryID      int         `json:"customIndustryId"`
			WorkPerformance       interface{} `json:"workPerformance"`
			WorkEmphasisList      interface{} `json:"workEmphasisList"`
			CertStatus            int         `json:"certStatus"`
			WorkType              int         `json:"workType"`
			AddTime               interface{} `json:"addTime"`
			UpdateTime            interface{} `json:"updateTime"`
			CompanyHighlight      interface{} `json:"companyHighlight"`
			PositionNameHighlight interface{} `json:"positionNameHighlight"`
			WorkTime              interface{} `json:"workTime"`
			WorkMonths            int         `json:"workMonths"`
			StillWork             bool        `json:"stillWork"`
			StartYearMonStr       string      `json:"startYearMonStr"`
			EndYearMonStr         string      `json:"endYearMonStr"`
			Current               bool        `json:"current"`
			WorkTypeIntern        bool        `json:"workTypeIntern"`
		} `json:"geekWorks"`
		GeekEdu struct {
			ID             int         `json:"id"`
			UserID         int         `json:"userId"`
			School         string      `json:"school"`
			SchoolID       int         `json:"schoolId"`
			Major          string      `json:"major"`
			Degree         int         `json:"degree"`
			DegreeName     string      `json:"degreeName"`
			EduType        int         `json:"eduType"`
			StartDate      string      `json:"startDate"`
			EndDate        string      `json:"endDate"`
			EduDescription interface{} `json:"eduDescription"`
			AddTime        interface{} `json:"addTime"`
			UpdateTime     interface{} `json:"updateTime"`
			TimeSlot       interface{} `json:"timeSlot"`
			StartYearStr   string      `json:"startYearStr"`
			EndYearStr     string      `json:"endYearStr"`
		} `json:"geekEdu"`
		GeekEdus []struct {
			ID             int         `json:"id"`
			UserID         int         `json:"userId"`
			School         string      `json:"school"`
			SchoolID       int         `json:"schoolId"`
			Major          string      `json:"major"`
			Degree         int         `json:"degree"`
			DegreeName     string      `json:"degreeName"`
			EduType        int         `json:"eduType"`
			StartDate      string      `json:"startDate"`
			EndDate        string      `json:"endDate"`
			EduDescription interface{} `json:"eduDescription"`
			AddTime        interface{} `json:"addTime"`
			UpdateTime     interface{} `json:"updateTime"`
			TimeSlot       interface{} `json:"timeSlot"`
			StartYearStr   string      `json:"startYearStr"`
			EndYearStr     string      `json:"endYearStr"`
		} `json:"geekEdus"`
		ExpectLocation     int           `json:"expectLocation"`
		ExpectPosition     int           `json:"expectPosition"`
		ApplyStatus        int           `json:"applyStatus"`
		ActiveSec          int           `json:"activeSec"`
		Birthday           string        `json:"birthday"`
		ExpectLocationName string        `json:"expectLocationName"`
		ExpectPositionName string        `json:"expectPositionName"`
		ApplyStatusDesc    string        `json:"applyStatusDesc"`
		FreshGraduate      int           `json:"freshGraduate"`
		AgeDesc            string        `json:"ageDesc"`
		GeekSource         int           `json:"geekSource"`
		EncryptGeekID      string        `json:"encryptGeekId"`
		Anonymous          bool          `json:"anonymous"`
		ContactStatus      interface{}   `json:"contactStatus"`
		ExpectInfos        []interface{} `json:"expectInfos"`
		Experiences        []interface{} `json:"experiences"`
		Edus               []interface{} `json:"edus"`
		FeedbackTitle      string        `json:"feedbackTitle"`
		Feedback           []struct {
			Code           int         `json:"code"`
			Memo           string      `json:"memo"`
			ShowType       int         `json:"showType"`
			FeedbackL2List interface{} `json:"feedbackL2List"`
			TitleL2        interface{} `json:"titleL2"`
		} `json:"feedback"`
		DataSource   int `json:"dataSource"`
		InteractDesc struct {
			Content   string `json:"content"`
			IndexList []struct {
				Start int `json:"start"`
				End   int `json:"end"`
			} `json:"indexList"`
		} `json:"interactDesc"`
		Clicked          bool        `json:"clicked"`
		ForHomePage      bool        `json:"forHomePage"`
		HomePageAction   int         `json:"homePageAction"`
		CanUseDirectCall bool        `json:"canUseDirectCall"`
		ExposeEnhanced   interface{} `json:"exposeEnhanced"`
		Matches          interface{} `json:"matches"`
		Viewed           bool        `json:"viewed"`
		Contacting       bool        `json:"contacting"`
	} `json:"geekCard"`
	ActiveTimeDesc string      `json:"activeTimeDesc"`
	TalkTimeDesc   interface{} `json:"talkTimeDesc"`
	PositionName   string      `json:"positionName"`
	Cooperate      int         `json:"cooperate"`
	IsFriend       int         `json:"isFriend"`
	ItemID         int         `json:"itemId"`
	Suid           string      `json:"suid"`
	GeekCallStatus int         `json:"geekCallStatus"`
	Blur           int         `json:"blur"`
	HaveChatted    int         `json:"haveChatted"`
	Tag            int         `json:"tag"`
	MateShareID    int         `json:"mateShareId"`
	MateName       interface{} `json:"mateName"`
	ShareMessage   int         `json:"shareMessage"`
	ShareNote      interface{} `json:"shareNote"`
	EncryptShareID string      `json:"encryptShareId"`
	EncryptGeekID  string      `json:"encryptGeekId"`
	FriendGeek     bool        `json:"friendGeek"`
	UsingGeekCall  bool        `json:"usingGeekCall"`
	BlurGeek       bool        `json:"blurGeek"`

	Weight int // 选择权重，越大越优先考虑
}

type SortGeek []*Geek

func (s SortGeek) Len() int {
	return len(s)
}

func (s SortGeek) Less(i, j int) bool {
	if s[i].Weight > s[j].Weight {
		return true
	}
	return false
}

func (s SortGeek) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

type JobHelloMsg struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	ZpData  struct {
		Jobs []struct {
			JobID           int    `json:"jobId"`
			EncJobID        string `json:"encJobId"`
			JobName         string `json:"jobName"`
			JobCity         string `json:"jobCity"`
			JobSalary       string `json:"jobSalary"`
			Degree          string `json:"degree"`
			Experience      string `json:"experience"`
			Set             bool   `json:"set"`
			EncGreetingID   string `json:"encGreetingId"`
			JobGreeting     string `json:"jobGreeting"`
			JobAddTime      int64  `json:"jobAddTime"`
			GreetingAddTime int64  `json:"greetingAddTime"`
		} `json:"jobs"`
		Greetings []struct {
			JobID           int    `json:"jobId"`
			EncJobID        string `json:"encJobId"`
			JobName         string `json:"jobName"`
			JobCity         string `json:"jobCity"`
			JobSalary       string `json:"jobSalary"`
			Degree          string `json:"degree"`
			Experience      string `json:"experience"`
			Set             bool   `json:"set"`
			EncGreetingID   string `json:"encGreetingId"`
			JobGreeting     string `json:"jobGreeting"`
			JobAddTime      int64  `json:"jobAddTime"`
			GreetingAddTime int64  `json:"greetingAddTime"`
		} `json:"greetings"`
	} `json:"zpData"`
}

type QRMsg struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	ZpData  struct {
		QrID    string `json:"qrId"`
		RandKey string `json:"randKey"`
	} `json:"zpData"`
}

type JobListResp struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	ZpData  struct {
		Page      int  `json:"page"`
		PageSize  int  `json:"pageSize"`
		HasMore   bool `json:"hasMore"`
		StartRow  int  `json:"startRow"`
		TotalSize int  `json:"totalSize"`
		Data      []struct {
			JobID                    int         `json:"jobId"`
			EncryptID                string      `json:"encryptId"`
			JobAuditStatus           int         `json:"jobAuditStatus"`
			JobStatus                int         `json:"jobStatus"`
			JobName                  string      `json:"jobName"`
			ShowType                 int         `json:"showType"`
			LocationName             string      `json:"locationName"`
			ExperienceName           string      `json:"experienceName"`
			DegreeName               string      `json:"degreeName"`
			JobTypeName              string      `json:"jobTypeName"`
			LowSalary                int         `json:"lowSalary"`
			HighSalary               int         `json:"highSalary"`
			ViewCount                int         `json:"viewCount"`
			ConcatCount              int         `json:"concatCount"`
			AddTime                  int64       `json:"addTime"`
			UpdateTime               int64       `json:"updateTime"`
			LastModifyTime           int64       `json:"lastModifyTime"`
			Deleted                  int         `json:"deleted"`
			SalaryMonth              int         `json:"salaryMonth"`
			AddTimeDesc              string      `json:"addTimeDesc"`
			SalaryDesc               string      `json:"salaryDesc"`
			FreeExperience           bool        `json:"freeExperience"`
			Hot                      bool        `json:"hot"`
			RemainDays               int         `json:"remainDays"`
			UnpaidVip                int         `json:"unpaidVip"`
			HotPayStatus             int         `json:"hotPayStatus"`
			CanDelay                 bool        `json:"canDelay"`
			ViewChatCount            int         `json:"viewChatCount"`
			ShowQuickTop             bool        `json:"showQuickTop"`
			QuickTopTime             int         `json:"quickTopTime"`
			QuickTopType             int         `json:"quickTopType"`
			Location                 int         `json:"location"`
			Position                 int         `json:"position"`
			Experience               int         `json:"experience"`
			Degree                   int         `json:"degree"`
			JobType                  int         `json:"jobType"`
			SkillRequire             string      `json:"skillRequire"`
			PositionName             string      `json:"positionName"`
			City                     string      `json:"city"`
			EncryptJobID             string      `json:"encryptJobId"`
			Insurance                int         `json:"insurance"`
			HasDiscount              int         `json:"hasDiscount"`
			PaidJobEndDate           interface{} `json:"paidJobEndDate"`
			AuditingDesc             interface{} `json:"auditingDesc"`
			DaysPerWeekText          interface{} `json:"daysPerWeekText"`
			LeastMonthText           interface{} `json:"leastMonthText"`
			BrandName                string      `json:"brandName"`
			ProxyJob                 int         `json:"proxyJob"`
			ProxyType                int         `json:"proxyType"`
			ComID                    int         `json:"comId"`
			BrandID                  int         `json:"brandId"`
			BrandLogo                string      `json:"brandLogo"`
			Anonymous                int         `json:"anonymous"`
			NewUnpass                bool        `json:"newUnpass"`
			IntermediaryAuditingText interface{} `json:"intermediaryAuditingText"`
			ExtendShowBar            struct {
				Type     int         `json:"type"`
				TipText  string      `json:"tipText"`
				TipDesc  interface{} `json:"tipDesc"`
				LeftTime int         `json:"leftTime"`
			} `json:"extendShowBar"`
			Urged           bool        `json:"urged"`
			UpgradeType     int         `json:"upgradeType"`
			NeedSupplyQuest bool        `json:"needSupplyQuest"`
			JobShowBar      interface{} `json:"jobShowBar"`
			SystemCloseTip  interface{} `json:"systemCloseTip"`
			ShowPriority    bool        `json:"showPriority"`
		} `json:"data"`
		HasNewUnpassJobs   bool `json:"hasNewUnpassJobs"`
		DelayDeadLine      int  `json:"delayDeadLine"`
		HasCertFailAddress bool `json:"hasCertFailAddress"`
		BossJdQmpLibState  int  `json:"bossJdQmpLibState"`
		OperateJobForbid   struct {
			Type    int         `json:"type"`
			Title   interface{} `json:"title"`
			Content interface{} `json:"content"`
		} `json:"operateJobForbid"`
	} `json:"zpData"`
}

type Job struct {
	JobId   string
	JobName string
}

type RequestResumesResp struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	ZpData  struct {
		Type   int `json:"type"`
		Status int `json:"status"`
	} `json:"zpData"`
}
