# Graphmailer

Graphmailer is a simple Go tool for sending HTML emails using the Microsoft Graph API.

## Configuration

The application requires the following environment variables to authenticate with Microsoft Graph:

- `CLIENT_ID`: Your Entra ID Application Client ID.
- `CLIENT_SECRET`: Your Entra ID Application Client Secret.
- `TENANT_ID`: Your Azure Tenant ID.
- `USER_ID`: The Object ID or User Principal Name (email) of the user sending the mail.

### Entra ID App Registration

Graphmailer needs some permissions to use M$ GraphAPI and send mails behalf of users. First it needs an Entra ID App registered, which sets the permissions. To do that, go either in Azure Portal or Entra Admin Center to Microsoft Entra ID view and:

1. Open `App registrations`
2. Select `New registration`
3. Set at least name and organization level, redirect URI is optional 
4. Select `API permissions`
5. Add a permission `Mail.Send` by choosing first GraphAPI, then `application permissions` as a type and then pick `Mail.Send`
6. Next, select `App roles`
7. Create new, set name, choose `Applications` as member types, set `Mail.Send` as value and write a description. Check that `enabled` is checked
8. Then select `Owners` and add owner, e.g., yourself.
9. Finally, create a client secret for the app in `Certificates & secrets`, and you're done

Now you have `CLIENT_ID` and `CLIENT_SECRET` for graphmailer. Put them e.g. in `.env` file:

```.env
CLIENT_ID=123
TENANT_ID=666
CLIENT_SECRET=allyourbasearebelongtous
USER_ID=your-address@example.com
```

To complete the setup:

1.  Create a Microsoft 365 Group in Exchange (e.g., `do-not-reply@example.com`)
2.  Grant the user account specified in `USER_ID` the **"Send As"** permission for that group (Edit delegates)
3.  Configure the mailer to use the group address as the sender.

The application authenticates as itself (using the Client Credentials), but it calls the API endpoint of the `USER_ID` user. Because that user has rights to "Send As" the group, the email is successfully delivered appearing from the group address.

## Usage

Here is a basic example of how to use the library, based on [`cmd/main.go`](cmd/main.go):

```go
package main

import (
    m "github.com/pekk4/graphmailer/pkg/graphmailer"
)

func main() {
    // Initialize session from environment variables
    session := m.InitSession()

    recipients := []string{
        "recipient@example.com",
    }

    subject := "Hello from Graphmailer!"
    // Path to your HTML template
    template := "./static/templates/template.html" 
    sender := "do-not-reply@example.com"
    
    // Create attachments (optional)
    // NewAttachment(name, contentBytes, contentType, contentId, isInline)
    attachments := []m.Attachment{
        m.NewAttachment("logo.png", m.ExampleLogo, "image/png", "logo.png", "true"),
    }

    // Send emails
    // HtmlMailer(subject, sender, templatePath, attachments, recipients, session, batchSize, breakMinutes)
    m.HtmlMailer(subject, sender, template, attachments, recipients, session, 500, 3)
}
```

`batchSize` defines how many emails are sent consecutively before pausing and `breakMinutes` defines how many minutes the application sleeps between batches. The example above sends 500 mails and then sleeps 3 minutes before the next 500 mails.

### Running

Get deps and run mailer with:

```bash
go mod tidy
go run cmd/main.go
```

If you want to write a log file when running, get the output like:

```bash
go run cmd/main.go 2> mail.log
```
