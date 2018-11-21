# now-static-bin

This package provides a
[builder](https://zeit.co/docs/v2/deployments/builders/overview#when-to-use-builders)
for Zeit's [Now 2.0](https://zeit.co/blog/now-2) offering that enables running
arbitrary executables in the generated lambdas.

## Usage

Your `now.json` `"builds"` section should look something like this:

## Example

**Note**: don't forget to add `"version": 2` in your `now.json` file to use Now
2.0 explicitly.

```json
{
  "builds": [{
    "src": "*.exe",
    "use": "now-static-bin",
    "config": {
      "port": 4000,
    }
  }]
}
```

## Configuration Options

- `port`: the port that the deployed server listens on. Defaults to 8080.
- `timeout`: the timeout that the launcher waits for your server to start
  listening on the specified port. Defaults to 50ms.

## Limitations

- Currently only HTTP servers are supported.
- `"port"` is mandatory in the configuration, defaults to 8080.
- The
  [`maxLambdaSize`](https://zeit.co/docs/v2/deployments/concepts/lambdas/#maximum-bundle-size)
  setting defaults to 25MB. You can override this (up to a limit of 50MB) in the
  builder config, e.g.:

```json
{
  "builds": [{
    "src": "*.exe",
    "use": "now-static-bin",
    "config": {
      "port": 4000,
      "maxLambdaSize": "50mb"
      ^^^ NEW
    }
  }]
}
```

## Copyright and License

Copyright © 2018 António Nuno Monteiro.

Distributed under the MIT License (see [LICENSE](./LICENSE)).
