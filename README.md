# CryptoJS

This library has been created for the DEDIS projects to use the crypto 
primitives implemented by the Go library in Javascript.

## How to compile

We use GopherJS to transpile Go code in Javascript then you need to set
up your environment by following the instructions [here](https://github.com/gopherjs/gopherjs)

Then you can simply use `gopherjs build` to build the file.

## How to use

### ES6

This example assumes the file is in the current folder but you can
specify the path as usual
```
import './crypto-js'
```

### Others

```
<script src="/crypto-js.js"></script>
```