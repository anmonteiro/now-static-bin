# Changelog

## [master](https://github.com/anmonteiro/now-static-bin/compare/0.2.2...HEAD)  (unreleased)

## [0.2.2](https://github.com/anmonteiro/now-static-bin/compare/0.2.1...0.2.2) (2018-11-23)

### Bug fixes

- Preserve HTTP Host information when proxying requests
  ([#4](https://github.com/anmonteiro/now-static-bin/pull/4)).

## [0.2.1](https://github.com/anmonteiro/now-static-bin/compare/0.2.0...0.2.1) (2018-11-23)

### Changes

- Don't run UPX on the launcher, as it increases startup time
  ([eff3ad8](https://github.com/anmonteiro/now-static-bin/commit/eff3ad83adf77a8b4b642e2f7009d876e8db0d57)).
  See [zeit/now-builders#63](https://github.com/zeit/now-builders/issues/63) for context.

## [0.2.0](https://github.com/anmonteiro/now-static-bin/compare/0.1.1...0.2.0) (2018-11-22)

### New features

- Allow a `timeout` configuration option to specify how long the launcher should
  wait for the server to start
  ([#1](https://github.com/anmonteiro/now-static-bin/pull/1)). Defaults to 50
  milliseconds.

### Changes

- Rewrite the launcher in Golang for faster startup and native Lambda support
  ([#1](https://github.com/anmonteiro/now-static-bin/pull/1)).

## [0.1.1](https://github.com/anmonteiro/now-static-bin/compare/0.1.0...0.1.1) (2018-11-19)

### Changes

- Allow to configure the maximum lambda size via
  [`maxLambdaSize`](https://zeit.co/docs/v2/deployments/concepts/lambdas/#maximum-bundle-size).
  ([67a5c62f](https://github.com/anmonteiro/now-static-bin/commit/67a5c62f7d86e18e9c5d867c7bf8c11005eebcdd)).
  Defaults to 25MB.

## [0.1.0](https://github.com/anmonteiro/now-static-bin/releases/tag/0.1.0) (2018-11-10)

- Initial version.