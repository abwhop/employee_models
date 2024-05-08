package employee_models

type Employee struct {
	PersonnelNumber     string `gorm:"personnel_number" json:"personnelNumber"`
	FullPersonnelNumber string `gorm:"full_personnel_number" json:"fullPersonalNumber"`
	HrSystem            int    `gorm:"personnel_number" json:"hrSystem"`
	HumanId             int    `gorm:"personnel_number" json:"humanId"`
	LastName            string `gorm:"personnel_number" json:"lastName"`
	FirstName           string `gorm:"personnel_number" json:"firstName"`
	MiddleName          string `gorm:"personnel_number" json:"middleName"`
	LastNameForeigner   string `gorm:"last_name_foreigner" json:"lastNameForeigner"`
	FirstNameForeigner  string `gorm:"first_name_foreigner" json:"firstNameForeigner"`
	MiddleNameForeigner string `gorm:"middle_name_foreigner" json:"middleNameForeigner"`
	DateHiring          int    `gorm:"date_hiring" json:"dateHiring"`
	DateResignation     int    `gorm:"date_resignation" json:"dateResignation"`
	DateTransfer        int    `gorm:"date_transfer" json:"dateTransfer"`
	DateStartVacation   int    `gorm:"date_start_vacation" json:"dateStartVacation"`
	DateEndVacation     int    `gorm:"date_end_vacation" json:"dateEndVacation"`
	TypeVacation        int    `gorm:"type_vacation" json:"typeVacation"`
	PositionCode        string `gorm:"position_code" json:"positionCode"`
	PositionName        string `gorm:"position_name" json:"positionName"`
	CompanyCode         string `gorm:"company_code" json:"companyCode"`
	CompanyName         string `gorm:"company_name" json:"companyName"`
	ManufactoryCode     string `gorm:"manufactory_code" json:"manufactoryCode"`
	ManufactoryName     string `gorm:"manufactory_name" json:"manufactoryName"`
	DepartmentId        string `gorm:"department_id" json:"departmentId"`
	DepartmentName      string `gorm:"department_name" json:"departmentName"`
	CategoryCode        int    `gorm:"category_code" json:"categoryCode"`
	DayBirth            int    `gorm:"day_birth" json:"dayBirth"`
	MonthBirth          int    `gorm:"month_birth" json:"monthBirth"`
	ManufactoryAddress  string `gorm:"manufactory_address" json:"manufactoryAddress"`
	DepartmentAddress   string `gorm:"department_address" json:"departmentAddress"`
	ChiefidHuman        int    `gorm:"chiefid_human" json:"chiefidHuman"`
	ListHigherDivisions string `gorm:"list_higher_divisions" json:"listHigherDivisions"`
	EmployeeCity        string `gorm:"employee_city" json:"employeeCity"`
	Isawareib           int    `gorm:"isawareib" json:"isawareib"`
	Isawarekt           int    `gorm:"isawarekt" json:"isawarekt"`
	Isawarepdn          int    `gorm:"isawarepdn" json:"isawarepdn"`
	PhoneInternal       string `gorm:"phone_internal" json:"phoneInternal"`
	PhoneMobile         string `gorm:"phone_mobile" json:"phoneMobile"`
	LoginAd             string `gorm:"login_ad" json:"loginAd"`
	DateStartLoginAd    int    `gorm:"date_start_login_ad" json:"dateStartLoginAd"`
	DateEndLoginAd      int    `gorm:"date_end_login_ad" json:"dateEndLoginAd"`
	LastCompanyCode     string `gorm:"last_company_code" json:"lastCompanyCode"`
	IsLastCompanyNlmk   int    `gorm:"is_last_company_nlmk" json:"isLastCompanyNlmk"`
	Snils               string `gorm:"snils" json:"snils"`
	IsUpdated           int    `gorm:"is_updated" json:"isUpdated"`
	ChangeType          string `gorm:"change_type" json:"changeType"`
	EmployeeId          int    `gorm:"employee_id" json:"employeeId"`
	JobType             int    `gorm:"job_type" json:"jobType"`
	CreateDate          int    `gorm:"create_date" json:"createDate"`
	ChangeDate          string `gorm:"change_date" json:"changeDate"`
	ProcessDate         string `gorm:"process_date" json:"processDate"`
	SubDivision         string `gorm:"sub_division" json:"subDivision"`
	Division            string `gorm:"division" json:"division"`
	FunctionalArea      string `gorm:"functional_area" json:"functionalArea"`
	DepartmentNameSp    string `gorm:"department_name_sp" json:"departmentNameSp"`
	DepartmentNameAd    string `gorm:"department_name_ad" json:"departmentNameAd"`
	DepartmentHouse     string `gorm:"department_house" json:"departmentHouse"`
	Gender              int    `gorm:"gender" json:"gender"`
	Email               string `gorm:"email" json:"email"`
	DateStartMail       int    `gorm:"date_start_mail" json:"dateStartMail"`
	DateEndMail         int    `gorm:"date_end_mail" json:"dateEndMail"`
	QualificationName   string `gorm:"qualification_name" json:"qualificationName"`
	QualificationCode   int    `gorm:"qualification_code" json:"qualificationCode"`
	Persk               string `gorm:"persk" json:"persk"`
	Pernr               string `gorm:"pernr" json:"pernr"`
	Graid               int    `gorm:"graid" json:"graid"`
	EmployeeType        string `gorm:"employee_type" json:"employeeType"`
}

func (Employee) TableName() string {
	return "public.employees"
}
