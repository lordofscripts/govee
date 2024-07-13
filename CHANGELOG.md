# RELEASE NOTES

> Govee API Utility `govee` written in pure GO

## [1.2](https://github.com/lordofscripts/govee/compare/v1.2...v1.1) (2024-07-13)

> A quick maintenance release with support for MacOS & Windows. GoveeLux has
> now gone beyond the Raspberry Pi box boundaries.

### New Features

* Added `-init` command-line flag. Creates sample configuration.
* It should now work on Linux, MacOS & Windows

### Other Changes

* README now has informational badges and Go mascot cheering Govee
* The configuration file name is now `goveelux.json`
* The resulting executable is now `goveelux`
* Coverage improvements: new tests.
* Some refactoring to improve readability and maintenance
* Formatting changes to boost Go Report Card from C to A.
* Implemented GitHub workflow for Go


## [1.1](https://github.com/lordofscripts/govee/compare/v1.1...v1.0) (2024-07-10)

> Added basic support for light control.

### New Features

* Set light brightness (`-brightness`)
* Set light color (`-color`)
* State (`-state`) output now reports light properties if applicable

### Bug Fixes

* no known bugs

### Other Changes

* README split in separate README and USER MANUAL files
* Refactoring

---

## [1.0](https://github.com/lordofscripts/govee/compare/v1.0...v1.0) (2023-11-02)

> Initial version.

### Initial Features

* Configuration file which aliases MAC & device model to an ID with location
* Lists all Govee devices in the local network (`-list`)
* Queries the connection and power state of a device  (`-state`)
* Turns on a device (`-on`)
* Turns off a device (`-off`)
