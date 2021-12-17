# Goflate

> A simple & small CLI tool to compress and decompress files using Golang

<div align="center">
	<img src="assets/shrink.gif" alt="Compression" height="300px">
</div>

##  Usage

Install the binary to your local machine with the below command

```sh
go get github.com/pedreviljoen/go-flate/goflate
```

Ensure your GOPATH variable is set to your path, if not you can do so with the below command

```sh
export PATH=$PATH:$(go env GOPATH)/bin
```

All set, the binary is downloaded and instaled at your GOPATH, to start using this package simply head to your terminal and execute the below command and a list of instructions will appear

```sh 
goflate
```

## Features

- [x] Compress
- [x] Decompress
- [x] Compression stats

## Contribute

Contributions are welcome!

1. Fork it.
2. Create your feature branch: `git checkout -b my-new-feature`
3. Commit your changes: `git commit -am 'Add some feature'`
4. Push to the branch: `git push origin my-new-feature`
5. Submit a pull request :D

Or open up [a issue](https://github.com/pedreviljoen/go-flate/issues).

## License

Apache License 2.0
