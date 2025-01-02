package main

import (
	"fmt"
	"os"
)

type Portfolio struct {
	Name       string
	Title      string
	Github     string
	LinkedIn   string
	About      string
	Skills     []string
	Experience []string
}

func main() {

	if len(os.Args) < 2 {
		help()
		return
	}

	portfolio := Portfolio{
		Name:     "Pratik",
		Title:    "Software Engineer",
		About:    "I'm a software engineer passionate about building cool stuff!",
		Github:   "github.com/prattsvee",
		LinkedIn: "linkedin.com/in/pratik-veera",
		Skills: []string{
			"Go",
			"Python",
			"Docker",
			"Kubernetes",
			"AWS",
		},
		Experience: []string{
			"Software Engineer at NBC Universal (Jun 2024-Sep 2024)",
			"Software Engineer at Tata Communication (Aug 2022-Jul 2023)",
			"Android Developer at Art of Living (Feb 2020-Dec 2020)",
		},
	}

	switch os.Args[1] {
	case "about":
		About(portfolio)
	case "skills":
		Skills(portfolio)
	case "experience": // Fixed typo in "experience"
		Experience(portfolio)
	case "contact": // Fixed typo in "contact"
		Contact(portfolio)
	default:
		help()
	}
}

func About(p Portfolio) {
	fmt.Printf("\n👋 Hi, I'm %s\n", p.Name)
	fmt.Printf("🚀 %s\n", p.Title)
	fmt.Printf("\n%s\n\n", p.About)
}

func Skills(p Portfolio) {
	fmt.Println("\n🛠️  Skills:")
	for _, skill := range p.Skills {
		fmt.Printf("• %s\n", skill)
	}
	fmt.Println()
}

func Experience(p Portfolio) {
	fmt.Println("\n💼 Experience:")
	for _, exp := range p.Experience {
		fmt.Printf("• %s\n", exp)
	}
	fmt.Println()
}

func Contact(p Portfolio) {
	fmt.Println("\n📫 Contact:")
	fmt.Printf("Github: %s\n", p.Github)
	fmt.Printf("LinkedIn: %s\n", p.LinkedIn)
	fmt.Println()
}

func help() {
	fmt.Println(`
Available commands:
  about       - Show introduction
  skills      - List technical skills
  experience  - Show work history
  contact     - Show contact info
    `)
}
