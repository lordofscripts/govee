# Lord of Scripts&trade; GoveeLux CLI

[![Years](https://badges.pufler.dev/years/lordofscripts)](https://badges.pufler.dev)
[![Go Report Card](https://goreportcard.com/badge/github.com/lordofscripts/govee?style=flat-square)](https://goreportcard.com/report/github.com/lordofscripts/govee)
![Tests](https://github.com/lordofscripts/govee/actions/workflows/go.yml/badge.svg)
[![Coverage](https://coveralls.io/repos/github/lordofscripts/govee/badge.svg?branch=main)](https://coveralls.io/github/lordofscripts/govee?branch=main)
[![Visits](https://badges.pufler.dev/visits/lordofscripts/govee)](https://badges.pufler.dev)
[![Created](https://badges.pufler.dev/created/lordofscripts/govee)](https://badges.pufler.dev)
[![Updated](https://badges.pufler.dev/updated/lordofscripts/govee)](https://badges.pufler.dev)

![Successful](./docs/assets/goveelux.png)

*GoveeLux* is a handy command-line interface application to list all Govee smart
devices in your home network,  control them (On/Off/Brightness/Color), or
query their current state.

The more verbose way requires several **CLI** parameters such as device MAC & Model
in order to control them. Alternatively, you can put all your Govee smart
devices in a configuration file and assign them an easy-to-remember ID (alias).

With this you can use the CLI to perform the actions you desire, or create
desktop shortcuts that invoke the CLI command. This command-line utility has been tested
with Govee **Smart Plugs** and **Cube Wall Sconces** but works with all Govee devices that support
the Govee API, regardless whether they are lighting devices or appliances.

I'm sure you will find this utility quite handy. Feel free to consider a small
donation so that I can continue working on these utilities:


|     | Show your support   |
| --- | :---: |
| [ ![AllMyLinks](./docs/assets/allmylinks.png)](https://allmylinks.com/lordofscripts)      | visit <br> Lord of Scripts&trade; <br> on [AllMyLinks.com](https://allmylinks.com/lordofscripts)                  |
| [ ![Buy me a coffee](./docs/assets/buymecoffee.jpg)](https://allmylinks.com/lordofscripts)|  buy Lord of Scripts&trade; <br> a Capuccino on <br>[BuyMeACoffee.com](https://www.buymeacoffee.com/lostinwriting)|


## Requirements & Dependencies

`GoveeLux` is a pure GO application.

- GO v1.20+
- github.com/loxhill/go-vee
- Request a [Govee API key](https://developer.govee.com/reference/apply-you-govee-api-key)

### Obtaining a GOVEE API key

The keys are for personal use only, if you keep that in mind and suits your needs,
then all is okay.

- Install the Govee Home app on your mobile. You probably already have since
  you have a Govee device. I am using Govee Home app v5.8.30.
- Click on the Cog wheel (`Settings`) on the upper right)
- In the `Settings` page, click on the `Apply for API Key`
- Fill out the requested information (Name, reason, agreement) and send.

You should receive the personal API key within a day.).
[Request](https://developer.govee.com/reference/apply-you-govee-api-key) your
GOVEE API key before using this application.

## Usage

The executable file (deliverable) is `govee`. For details about how to use it and documentation
of all the flags and parameters, please read the short [User Manual](./docs/USER_MANUAL.md).
You may also want to read the [Release Notes](./docs/CHANGELOG.md) to see which
features are supported in your application version.

-----
> All Rights Reserved [LordOfScripts&trade;](https://allmylinks.com/lordofscripts)
