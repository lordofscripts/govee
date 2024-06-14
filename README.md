# Lord of Scripts&trade; Govee CLI

This is a handy command-line interface application to list all Govee smart
devices in your home network, as well as control them (On/Off).

The more verbose way requires several **CLI** parameters such as device MAC & Model
in order to control them. Alternatively, you can put all your Govee smart
devices in a configuration file and assign them an easy-to-remember ID (alias).

With this you can use the CLI to perform the actions you desire, or create
desktop shortcuts that invoke the CLI command.

I'm sure you will find this utility quite handy. Feel free to consider a small
donation so that I can continue working on these utilities:

[Buy LostInWriting a coffee!](https://www.buymeacoffee.com/lostinwriting)

## Requirements

- GO v1.20
- github.com/loxhill/go-vee
- Request a Govee API key

## Using Govee API
This small Go app is conveniently named "govee" which is the name of the
executable.

### Configuration File

The configuration file, if present, should be at `~/.config/govee.json` and it
must be in JSON format. It would look like this:

```
{
	"version": "0.1",
	"apiKey": "YOUR GOVEE API KEY HERE"
	"devices": [
		{
			"model": "H5083",
			"mac": "DEVICE MAC ADDRESS",
			"alias": "SmartPlug1",
			"location": "Recamara #2"
		}
	]
}
```

The MAC address is a series of 8 hexadecimal pairs which is unique to each
device and looks like this: `AA:BB:CC:DD:EE:FF:00:11`

You will need to request a personal API key to Govee, they are quite quick. But
do remember that it is only for **personal** use, not commercial use! If you
use a configuration file, make sure to put your API key where indicated.
Alternatively, you can put the API key in the `GOVEE_API` environment variable.

### Obtaining a GOVEE API key

The keys are for personal use only, if you keep that in mind and suits your needs,
then all is okay.

- Install the Govee Home app on your mobile. You probably already have since
  you have a Govee device. I am using Govee Home app v5.8.30.
- Click on the Cog wheel (`Settings`) on the upper right)
- In the `Settings` page, click on the `Apply for API Key`
- Fill out the requested information (Name, reason, agreement) and send.

You should receive the personal API key within a day.)

### Usage

> govee -list

Uses your WiFi/Internet connection to query your local network and list ALL
Govee smart devices. It will list them giving you their model number (for use
with -model flag), MAC address (-mac flag) and official name.

> govee -mac {MAC_ADDRESS} -model {MODEL_NUMBER} -on

Turns on the device identified by the MAC address and Model number.

> govee -mac {MAC_ADDRESS} -model {MODEL_NUMBER} -off

Turns off the device identified by the MAC address and Model number.

However, it is much easier to assign easy identifications to your device and
even annotate it with their locations within your premises. For that, you need
to use a *configuration file* as described above. Once you do that, you can
easily turn your devices on or off using their alias. Device name aliases are much easier to remember. 

Here are the config-based versions of the
previous two commands:

> govee -id {ALIAS_NAME} -on

Turns on the device identified by the MAC address and Model number. For the
sample configuration above it would become `govee -id SmartPlug1 -on`

> govee -id {ALIAS_NAME} -off

Turns off the device identified by the MAC address and Model number.

-----
https://allmylinks.com/lordofscripts