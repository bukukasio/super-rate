# Changelog

## Unreleased

## 2.7.1 - 2020-11-12
* [#81](https://github.com/throttled/throttled/pull/81) Fix statistics calculation when quantity exceeds max burst

## 2.7.0 - 2020-10-09
* [#80](https://github.com/throttled/throttled/pull/80) In `goredisstore`, use `UniversalClient` interface instead of `*Client` implementation (this is backwards compatible)

## 2.6.0 - 2020-08-04
* [#64](https://github.com/throttled/throttled/pull/64) Add `SetTimeNow` to override getting current time to memstore driver
* [#66](https://github.com/throttled/throttled/pull/66) Add `PerDuration` function for getting a perfectly customized `Rate`

## 2.5.0 - 2020-08-02
* [#79](https://github.com/throttled/throttled/pull/79) Import Throttle with `/v2` suffix in the package path

## 2.4.0 - 2020-08-01
* [#78](https://github.com/throttled/throttled/pull/78) Revert upgrade to go-redis V8 (now back on V6)

## 2.3.0 - 2020-08-01
* [#76](https://github.com/throttled/throttled/pull/76) Add basic support for Go Modules

## 2.2.5 - 2020-08-01
* [#67](https://github.com/throttled/throttled/pull/67) Bug fix: Fix TTL in `SetIfNotExistsWithTTL`
* [#74](https://github.com/throttled/throttled/pull/74) Bug fix: Always select DB for Redigo store, even when DB is configured to 0

## 2.2.4 - 2018-11-19
* [#52](https://github.com/throttled/throttled/pull/52) Handle the possibility of `RemoteAddr` without port in `VaryBy`

## 2.2.3 - 2018-11-13
* [#49](https://github.com/throttled/throttled/pull/49) Handle the possibility of an empty `RemoteAddr` in `VaryBy`

## 2.2.2 - 2018-10-18
* [#47](https://github.com/throttled/throttled/pull/47) Don't include origin port in the identifier when using `throttled.VaryBy{RemoteAddr: true}`

## 2.2.1 - 2018-03-21
* [#40](https://github.com/throttled/throttled/pull/40) Replace unmaintained `garyburd/redigo` with `gomodule/redigo`

## 2.2.0 - 2018-03-19
* [#37](https://github.com/throttled/throttled/pull/37) Add `go-redis` support

## 2.1.0 - 2017-10-23
* [#27](https://github.com/throttled/throttled/pull/27) Never assign a Redis key's TTL to zero
* [#32](https://github.com/throttled/throttled/pull/32) Stop using `gopkg.in`

## 2.0.3 - 2015-09-09
* [#15](https://github.com/throttled/throttled/pull/15) Use non-HTTP example for `GCRARateLimiter`

## 2.0.2 - 2015-09-07
* [#14](https://github.com/throttled/throttled/pull/14) Add example demonstrating granular use of `RateLimit`

## 2.0.1 - 2015-09-01
* [#12](https://github.com/throttled/throttled/pull/12) Fix parsing of `TIME` in `redigostore`

## 2.0.0 - 2015-08-28
* [#9](https://github.com/throttled/throttled/pull/9) Substantially rebuild the APIs and implementation of the `throttled` package (wholly not backwards compatible)

(There are other versions, but this is the beginning of `CHANGELOG.md`.)

<!--
# vim: set tw=0:
-->
