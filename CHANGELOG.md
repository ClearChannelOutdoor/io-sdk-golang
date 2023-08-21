# Change Log

All notable changes in io-sdk-golang will be documented in this file.

The format is based on [Keep a Changelog](http://keepachangelog.com/)
and this project adheres to [Semantic Versioning](http://semver.org/).

## [0.2.1] - 2023-08-10

- Added support for schema updates to `market` package models to include DMA and CBSA details
- Adjusted the `options.AddFilter` method to be variadic to support passing multiple values easily

## [0.2.0] - 2023-08-10

- Added support for schema updates to `display` and `network` package models to include DMA and CBSA details
- Modified `campaign` package models to reflect latest changes

## [0.1.2] - 2023-06-30

- Added `SpotIDs` to the `Frame` model in the `geopath` package

## [0.1.1] - 2023-06-06

- Simplified check for expired or no PASETO

## [0.1.0] - 2023-05-23

- Added support for `context.Context`
- Simplified client create for consumers
- Added API clients for the following:
  - accounts
  - customers
  - geopath frames 
  - geopath measures
  - geopath categorical items (construction types, illumination, media, etc.)
  - advertised products
  - CCO product codes (taxonomy)
  - IAB v1 taxa
  - IAB v2 taxa

## [0.0.2] - 2023-04-07

- Implemented sub-resource routines for making API calls to child resources
- Added networks and network displays to available API clients 

## [0.0.1] - 2023-04-06

- Initial implementation with displays and markets

### Added

### Changed

### Fixed
