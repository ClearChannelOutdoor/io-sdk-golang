# Change Log

All notable changes in io-sdk-golang will be documented in this file.

The format is based on [Keep a Changelog](http://keepachangelog.com/)
and this project adheres to [Semantic Versioning](http://semver.org/).

## [0.5.5] - 2025-05-20

- fixed minor bug in `AddFilter` where only two filters would be created when using the variadic form of the function (as opposed to all provided parameters)

## [0.5.4] - 2025-02-24

- update Pricing-API's SaleableItem struct with a single PricingGuidance

## [0.5.3] - 2025-02-20

- add support for quotes on pricing-api

## [0.5.2] - 2025-01-02

- Update iab taxonomy model to match industry standard verbiage
- Added /v3/taxa endpoint to iab client

## [0.5.1] - 2024-11-26

- Added a mutex to the underlying client surrounding retrieval of the OAuth v2.0 token

## [0.5.0] - 2024-11-14

- Updating and correcting displays model

## [0.4.9] - 2024-11-05

- Fixing a few remaining details references

## [0.4.8] - 2024-11-05

- Renamed Segment Details to Buy Types

## [0.4.7] - 2024-11-04

- Add Cost field to booking model

## [0.4.6] - 2024-10-28

Updates to the Order model:

- Start and End dates are no longer pointers
- promote Advertiser and Buyer to root level
- remove Employee FullName and add Number
- add Canceled field to Market related data

## [0.4.5] - 2024-10-18

- Start and End dates in the booking model are no longer pointers

## [0.4.4] - 2024-10-16

- Updated photo model to have `submittedBy` field

## [0.4.3] - 2024-10-11

- Update booking model to reflect booking-api

## [0.4.2] - 2024-10-03

- Added customer marketingRestricted boolean field

## [0.4.1] - 2024-09-30

- Added modify scopes to bookings and orders

## [0.4.0] - 2024-09-27

- Provided support for specifying headers when working with API clients

## [0.3.16] - 2024-09-25

- [IOC-875] updating booking model

## [0.3.15] - 2024-08-23

- Fixing disparity between the json declaration for bookingID here versus what's in booking-api

## [0.3.14] - 2024-08-14

- Updated display-api models and added mediaProducts to display model
- Fixed filters in example in readme

## [0.3.13] - 2024-08-14

- Added support for url-api

## [0.3.12] - 2024-08-12

- update booking DigitalDetails model

## [0.3.11] - 2024-08-08

- Remove structure-api since it is being deprecated

## [0.3.10] - 2024-08-07

- Add support for contracts, segment details, and structures

## [0.3.9] - 2024-07-24

- Updated models for orders, orderLines, and bookings

## [0.3.8] - 2024-07-09

- Added support for `photos`

## [0.3.7] - 2024-05-29

- Added support for v2 `creatives`

## [0.3.6] - 2024-05-24

- Updated `account` sub model to include customer number field

## [0.3.5] - 2024-05-07

- Added support for `renewals`, `orders` and `orderlines`, and `bookings`

## [0.3.4] - 2023-11-16

- Added support for `creatives`

## [0.3.3] - 2023-11-15

- Exposing API errors in order to support consumers performing status checks
- Updated build automation manifest for CLI tool

## [0.3.2] - 2023-11-13

- Updated `customer` market sub model to include code field

## [0.3.1] - 2023-10-02

- Updated `customer` model to include customer number and taxonomy details

## [0.3.0] - 2023-08-21

- Modified query options encoding for optimal use with CCO microservices

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
