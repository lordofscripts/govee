# RELEASE NOTES

> Govee API Utility `govee` written in pure GO

Please see [Release Notes](./docs/RELEASE_NOTES.md) instead.

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
