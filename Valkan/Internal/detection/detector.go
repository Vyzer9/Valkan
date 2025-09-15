package detection

import "strings"

func DetectService(banner string) string {
	banner = strings.ToLower(banner)

	switch {
	case strings.Contains(banner, "ssh"):
		return "SSH"
	case strings.Contains(banner, "http"):
		return "HTTP"
	case strings.Contains(banner, "smtp"):
		return "SMTP"
	case strings.Contains(banner, "ftp"):
		return "FTP"
	case strings.Contains(banner, "mysql"):
		return "MySQL"
	case strings.Contains(banner, "postgres"):
		return "PostgreSQL"
	default:
		return "Desconhecido"
	}
}
