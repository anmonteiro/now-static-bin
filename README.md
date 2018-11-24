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
  "builds": [
    {
      "src": "*.exe",
      "use": "now-static-bin",
      "config": {
        "port": 4000
      }
    }
  ]
}
```

### Other Examples

Make sure to check the [`examples`](./examples) folder in this repo for
different language

They are also deployed on every commit to master and the latest build lives in
[`now-static-bin-examples.now.sh`](https://now-static-bin-examples.now.sh/).
Details:

| Example    | Demo     | Description     |
|:---------- |:---------|:----------------|
| [OCaml](/examples/bintry) | [https://now-static-bin-examples.now.sh/examples/bintry/main.exe](https://now-static-bin-examples.now.sh/examples/bintry/main.exe) | An OCaml static binary example server (no source available yet) |
| [Rust](/examples/rust) | [https://now-static-bin-examples.now.sh/examples/rust/server.exe](https://now-static-bin-examples.now.sh/examples/rust/server.exe) | A Rust [simple-server](https://github.com/steveklabnik/simple-server) example |
| [Reason](/examples/reason) | [https://now-static-bin-examples.now.sh/examples/reason/main.exe](https://now-static-bin-examples.now.sh/examples/reason/main.exe) | A [Reason](https://reasonml.github.io/) [Cohttp](https://github.com/mirage/ocaml-cohttp) server example that outputs request information |

## Configuration Options

- `port`: the port that the deployed server listens on. Defaults to 8080.
- `timeout`: the timeout that the launcher waits for your server to start
  listening on the specified port. Defaults to 50ms.

## Limitations

- Currently only HTTP servers are supported.
- The
  [`maxLambdaSize`](https://zeit.co/docs/v2/deployments/concepts/lambdas/#maximum-bundle-size)
  setting defaults to 25MB. You can override this (up to a limit of 50MB) in the
  builder config, e.g.:

```json
{
  "builds": [
    {
      "src": "*.exe",
      "use": "now-static-bin",
      "config": {
        "port": 4000,
        "maxLambdaSize": "50mb"
        ^^^ NEW
      }
    }
  ]
}
```

## Copyright and License

Copyright © 2018 António Nuno Monteiro.

Distributed under the MIT License (see [LICENSE](./LICENSE)).
