package models

type Portfolio struct {
	Profile      Profile           `json:"profile"`
	Statistics   []Statistic       `json:"statistics"`
	Skills       []Skill           `json:"skills"`
	Services     []ServiceOffering `json:"services"`
	Education    []EducationEntry  `json:"education"`
	Certificates []Certificate     `json:"certificates"`
	Projects     []Project         `json:"projects"`
	Contact      ContactInfo       `json:"contact"`
	SocialLinks  []SocialLink      `json:"social_links"`
}

type Profile struct {
	Name         string `json:"name"`
	Title        string `json:"title"`
	Bio          string `json:"bio"`
	Location     string `json:"location"`
	ProfileImage string `json:"profile_image"`
	Availability string `json:"availability"`
}

type Statistic struct {
	Label string `json:"label"`
	Value string `json:"value"`
}

type Skill struct {
	Name        string `json:"name"`
	Category    string `json:"category"`
	Proficiency int    `json:"proficiency"`
}

type ServiceOffering struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	Icon        string `json:"icon"`
}

type EducationEntry struct {
	Institution string `json:"institution"`
	Degree      string `json:"degree"`
	Field       string `json:"field"`
	StartYear   int    `json:"start_year"`
	EndYear     int    `json:"end_year"`
	Description string `json:"description"`
}

type Certificate struct {
	Title           string `json:"title"`
	Issuer          string `json:"issuer"`
	IssueDate       string `json:"issue_date"`
	CredentialID    string `json:"credential_id"`
	Description     string `json:"description"`
	Image           string `json:"image"`
	VerificationURL string `json:"verification_url"`
}

type Project struct {
	Name         string   `json:"name"`
	Description  string   `json:"description"`
	Image        string   `json:"image"`
	URL          string   `json:"url"`
	GitHub       string   `json:"github"`
	Technologies []string `json:"technologies"`
}

type ContactInfo struct {
	Email    string `json:"email"`
	Phone    string `json:"phone"`
	Location string `json:"location"`
}

type SocialLink struct {
	Platform string `json:"platform"`
	URL      string `json:"url"`
	Username string `json:"username"`
}
