# AWS STS Credential Retriever

[![Build Status](https://travis-ci.org/traveloka/lambda-sts-retriever.svg?branch=master)](https://travis-ci.org/traveloka/lambda-sts-retriever)

Using this package will enable sourcing AWS credential from external process. This package will retrieve temporary credential from custom STS endpoint provider.

This version was tested to retrieve STS credentials from [AWS Lambda STS](https://github.com/traveloka/terraform-aws-lambda-sts/releases/tag/v0.3.0)

## Table of Content

- [AWS STS Credential Retriever](#aws-sts-credential-retriever)
  - [Table of Content](#table-of-content)
  - [Dependencies](#dependencies)
  - [How to Build](#how-to-build)
  - [Quick Start](#quick-start)
    - [Usages](#usages)
  - [Todos](#todos)
  - [Contributing](#contributing)
  - [Contributor](#contributor)
  - [License](#license)
  - [Acknowledgments](#acknowledgments)

## Dependencies

- [Go](https://golang.org/dl/)
- [UPX](https://upx.github.io/)
- [awscli](https://github.com/aws/aws-cli)

## How to Build

- Download and configure [dependencies](#dependencies)
- Run `go build -ldflags="-s -w" && upx awscreds`. You may need to adjust `GOARCH` and `GOOS` environment for your build accordingly
- Make it executable `chmod +x awscreds`

## Quick Start

- Build the package or download latest [RELEASES](https://github.com/traveloka/custom-sts-retriever/releases) for your platform.
- Extract and rename the binary file and set in your execution `$PATH` variable. e.g `/opt/sts/awscreds`
- Configure aws cli configuration in `~/.aws/config` and replace the parameters accordingly. e.g:
```
[profile sts]
credential_process = /opt/sts/awscreds -u <URL of custom credential provider> -r <assumed role arn>
```

### Usages

Run `awscreds` to print usage. e.g:

```
NAME:
   awscreds - Set AWS Credential from custom STS provider

USAGE:
   awscreds [global options] command [command options] [arguments...]

VERSION:
   v0.1.0

AUTHOR:
   Frans Caisar Ramadhan <frans.ramadhan@traveloka.com>

COMMANDS:
   help, h  Shows a list of commands or help for one command

GLOBAL OPTIONS:
   --url value, -u value                                             URL to retrieve temporary STS credential. [$STS_API_URL]
   --assumed_role_arn value, -r value, --role_arn value              ARN of assumed IAM Role. [$ASSUMED_ROLE_ARN]
   --duration value, -d value, --token_duration value                Token duration in seconds. (default: 3600) [$STS_TOKEN_DURATION]
   --expiration_window value, -w value, --exp value, --expiry value  Expiry window for STS token duration. (default: 0) [$STS_TOKEN_WINDOW]
   --help, -h                                                        show help (default: false)
   --version, -v                                                     print the version (default: false)
```

## Todos

- Improve test coverage
- Enable CI / CD
- Enable github workflow

## Contributing

Check contribution guide in [CONTRIBUTING.md](https://github.com/traveloka/custom-sts-retriever/blob/master/CONTRIBUTING.md)

## Contributor

For question, issue, and pull request you can contact these people:

- [Frans Caisar Ramadhan](https://github.com/franzramadhan) (**Author**)

## License

See the [LICENSE](https://github.com/traveloka/custom-sts-retriever/blob/master/LICENSE)

## Acknowledgments

This repository was made possible by getting inspirations from below parties:

- [Readme Template](https://gist.github.com/PurpleBooth/109311bb0361f32d87a2)
- [Friendly Readme](https://rowanmanning.com/posts/writing-a-friendly-readme/)
- [Opensource Guide](https://opensource.guide/starting-a-project/)
- [Github Repository Template](https://github.com/traveloka/terraform-aws-modules-template)
- Inspiration from other open source projects
