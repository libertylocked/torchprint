torchprint
---
CLI and library to manage your printer job queue at NYU.

> Disclaimer: this project is not affiliated or endorsed by NYU.

# CLI
## Why
Because NYU IT doesn't provide driver installer for GNU/Linux. And with this you can also send docs to printer job queue directly from terminal.

## Usage
- Login
  - `torchprint init` then follow prompt
- Add a document (pdf, doc, txt, etc) to printer job queue
  - `torchprint ./craptoprint.pdf`
  - `--sides single` or `--sides double`. Default is double
  - `--color` for color, otherwise monochrome.
- View printer job queue
  - `torchprint ls`
- Delete a job from queue
  - `torchprint rm <job_d or index_in_queue>`
- ~~Edit a job in queue~~ (coming soon)

> Don't worry, you can always use `torchprint help`. (And no it won't call the help desk for you)

## Building
```
$ go get -u github.com/libertylocked/torchprint/cli/torchprint
```

## Security
Currently the CLI stores the base64 encoded `username:password` on disk in the clear. Although there are plans to replace it with session tokens.
