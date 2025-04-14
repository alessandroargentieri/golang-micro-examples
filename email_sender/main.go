package main

import (
	"flag"
	"fmt"
	"os"
	"strings"
	"time"
	
	"gopkg.in/gomail.v2"
)

// EmailConfig holds all the configuration for sending an email
type EmailConfig struct {
	SMTPHost        string
	SMTPPort        int
	SenderEmail     string
	SenderName      string
	Password        string
	RecipientEmail  string
	Subject         string
	PlainBody       string
	HTMLBody        string
	Attachments     []string
	ReplyTo         string
	ListUnsubscribe string
}

func main() {
	// Parse command-line flags
	smtpHost := flag.String("smtp-host", "", "SMTP server host (e.g., smtp.gmail.com)")
	smtpPort := flag.Int("smtp-port", 587, "SMTP server port")
	senderEmail := flag.String("sender", "", "Sender email address")
	senderName := flag.String("sender-name", "", "Sender name (optional)")
	password := flag.String("password", "", "Sender email password or app-specific password")
	recipient := flag.String("recipient", "", "Recipient email address")
	subject := flag.String("subject", "", "Email subject")
	plainBody := flag.String("plain-body", "", "Email plain text body")
	htmlBody := flag.String("html-body", "", "Email HTML body")
	plainBodyFile := flag.String("plain-body-file", "", "File containing plain text email body")
	htmlBodyFile := flag.String("html-body-file", "", "File containing HTML email body")
	attachmentsList := flag.String("attachments", "", "Comma-separated list of files to attach")
	replyTo := flag.String("reply-to", "", "Reply-To email address (defaults to sender)")
	listUnsubscribe := flag.String("list-unsubscribe", "", "List-Unsubscribe URL or email (for bulk emails)")
	
	flag.Parse()
	
	// Validate required parameters
	if *smtpHost == "" || *senderEmail == "" || *password == "" || *recipient == "" {
		fmt.Println("Error: Missing required parameters")
		fmt.Println("Required parameters: --smtp-host, --sender, --password, --recipient")
		flag.PrintDefaults()
		os.Exit(1)
	}
	
	// Get email bodies from files if specified
	emailPlainBody := *plainBody
	if *plainBodyFile != "" {
		content, err := os.ReadFile(*plainBodyFile)
		if err != nil {
			fmt.Printf("Error reading plain body file: %v\n", err)
			os.Exit(1)
		}
		emailPlainBody = string(content)
	}
	
	emailHTMLBody := *htmlBody
	if *htmlBodyFile != "" {
		content, err := os.ReadFile(*htmlBodyFile)
		if err != nil {
			fmt.Printf("Error reading HTML body file: %v\n", err)
			os.Exit(1)
		}
		emailHTMLBody = string(content)
	}
	
	// Check if we have a subject or body
	if *subject == "" && emailPlainBody == "" && emailHTMLBody == "" {
		fmt.Println("Warning: Subject and both body types are empty. Continuing anyway.")
	}
	
	// Parse attachments
	var attachments []string
	if *attachmentsList != "" {
		attachments = strings.Split(*attachmentsList, ",")
		for i, path := range attachments {
			attachments[i] = strings.TrimSpace(path)
		}
	}
	
	// Set default reply-to if not specified
	replyToEmail := *replyTo
	if replyToEmail == "" {
		replyToEmail = *senderEmail
	}
	
	// Initialize configuration
	config := EmailConfig{
		SMTPHost:        *smtpHost,
		SMTPPort:        *smtpPort,
		SenderEmail:     *senderEmail,
		SenderName:      *senderName,
		Password:        *password,
		RecipientEmail:  *recipient,
		Subject:         *subject,
		PlainBody:       emailPlainBody,
		HTMLBody:        emailHTMLBody,
		Attachments:     attachments,
		ReplyTo:         replyToEmail,
		ListUnsubscribe: *listUnsubscribe,
	}
	
	// Send the email
	err := sendEmail(config)
	if err != nil {
		fmt.Printf("Error sending email: %v\n", err)
		os.Exit(1)
	}
	
	fmt.Println("Email sent successfully!")
}

func sendEmail(config EmailConfig) error {
	m := gomail.NewMessage()
	
	// Set sender
	if config.SenderName != "" {
		m.SetAddressHeader("From", config.SenderEmail, config.SenderName)
	} else {
		m.SetHeader("From", config.SenderEmail)
	}
	
	// Set recipient, subject, and reply-to
	m.SetHeader("To", config.RecipientEmail)
	m.SetHeader("Subject", config.Subject)
	m.SetHeader("Reply-To", config.ReplyTo)
	
	// Set date header to prevent some spam filters
	m.SetHeader("Date", time.Now().Format(time.RFC1123Z))
	
	// Set Message-ID header for tracking and spam prevention
	messageID := fmt.Sprintf("<%d.%s@%s>", time.Now().UnixNano(), 
		strings.Split(config.SenderEmail, "@")[0], strings.Split(config.SenderEmail, "@")[1])
	m.SetHeader("Message-ID", messageID)
	
	// Add List-Unsubscribe header if provided (helps with spam filters)
	if config.ListUnsubscribe != "" {
		m.SetHeader("List-Unsubscribe", fmt.Sprintf("<%s>", config.ListUnsubscribe))
	}
	
	// Set email bodies (plain and/or HTML)
	if config.PlainBody != "" && config.HTMLBody != "" {
		m.SetBody("text/plain", config.PlainBody)
		m.AddAlternative("text/html", config.HTMLBody)
	} else if config.HTMLBody != "" {
		m.SetBody("text/html", config.HTMLBody)
	} else {
		m.SetBody("text/plain", config.PlainBody)
	}
	
	// Add attachments
	for _, file := range config.Attachments {
		m.Attach(file)
	}
	
	// Set up SMTP dialer
	d := gomail.NewDialer(config.SMTPHost, config.SMTPPort, config.SenderEmail, config.Password)
	
	// Send email
	return d.DialAndSend(m)
}

