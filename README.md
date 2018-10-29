torchprint
---
CLI and library to manage your printer job queue in Pharos Print Center (for GNU/Linux, Windows, MacOS)

> Disclaimer: this project is not affiliated or endorsed by NYU or Pharos.
>
> The API is reverse engineered from NYU Pharos Print Center web interface and may not work on all Pharos Print Center systems.
> (it's fair use)

# CLI

## Why
Because there's no Pharos Popup installer for GNU/Linux. (also it's proprietary)

## Usage
- Login
  - `torchprint login` then follow prompt
- Add a document (pdf, doc, txt, etc) to printer job queue
  - `torchprint add [file-to-print]`
  - `--side single` or `--side double`. Default is double
  - `--color` for color, otherwise monochrome.
- View printer job queue
  - `torchprint ls`
- Delete a job from queue
  - `torchprint rm [job-id]`
  - Delete everything with `--all`
- ~~Edit a job in queue~~ (coming soon)

> Don't worry, you can always use `torchprint help`. (And no it won't call the help desk for you)

## Example
```
$ torchprint login
Enter username: johndoe
Enter password:
success: userid deadbeef balance 4.20

$ torchprint add pirated_textbook.pdf
success: /printjobs/c29tZXRoaW5nLXlvdS13YW50LXRvLXByaW50LW1heWJlLXBpcmF0ZWQtc3R1ZmY= pirated_textbook.pdf Queued

$ torchprint ls
JOB ID                                                             NAME                SUBMISSION TIME             STATE
c29tZXRoaW5nLXlvdS13YW50LXRvLXByaW50LW1heWJlLXBpcmF0ZWQtc3R1ZmY=   test.txt            2018-10-29T18:12:07-04:00   Queued

$ torchprint rm c29tZXRoaW5nLXlvdS13YW50LXRvLXByaW50LW1heWJlLXBpcmF0ZWQtc3R1ZmY=
/printjobs/c29tZXRoaW5nLXlvdS13YW50LXRvLXByaW50LW1heWJlLXBpcmF0ZWQtc3R1ZmY= 200
```

## Building
```
$ go get -u github.com/libertylocked/torchprint/cmd/torchprint
$ $GOPATH/bin/torchprint version
```

## FAQ
Can I use it for my school?
> It has not been tested outside NYU, but if your school uses Pharos, chances are it will work. Just fork it and change the `baseURL` in `api.go`.

Can I save my username and password so I don't have to login every time?
> It is not recommended for security reasons, but yes you can.
> Run login command with `--save` will save your username and password in cleartext in config file.

Where is config file stored?
> It is stored in one of the two places
> - `$HOME/.config/torchprint/.torchprint.json`
> - `$HOME/.torchprint.json`

# Library
## Example
```go
import (
	"fmt"

	"github.com/libertylocked/torchprint"
)

func main() {
	// To get your user ID and token
	logon, token, _ := torchprint.NewAPI("").SetUserPass("username", "password")
	userID := logon.Identifier // User ID

	// To make requests using User ID and Session Token
	api := torchprint.NewAPI(userID).SetToken(token)
	// You can also make requests using User ID and Username:Password (instead of token)
	// api := torchprint.NewAPI(userID).SetUserPass("username", "password")

	// View all printjobs in queue
	api.GetPrintJobs()
	// Upload a document and create a printjob
	api.AddPrintJob("/home/liberty/Documents/pirated_textbook.pdf", torchprint.FinishingOptions{
		Mono:            true,
		Duplex:          true,
		PagesPerSide:    1,
		Copies:          1,
		DefaultPageSize: "Letter",
	})
}
```

## API Documentation
None ;)
