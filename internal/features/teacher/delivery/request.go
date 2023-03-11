package delivery

/*

	Struct Request

*/

type TeacherRequest struct {
	Name       string `json:"name" form:"name"`
	Email      string `json:"email" form:"email"`
	Password   string `json:"password" form:"password"`
	Role       string `json:"role" form:"role"`
	Image      string `json:"image" form:"image"`
	AgencyCode uint   `json:"agency_code" form:"agency_code"`
	Gender     string `json:"gender" form:"gender"`
	Address    string `json:"address" form:"address"`
	Portofolio string `json:"portofolio" form:"portofolio"`
	Expertise  string `json:"expertise" form:"expertise"`
	Material   string `json:"material" form:"material"`
}
