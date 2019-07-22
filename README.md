# directions

<img src="https://img.shields.io/badge/coverage-95%25-brightgreen.svg?style=flat-square" alt="Code coverage">&nbsp;<a href="https://travis-ci.org/schollz/directions"><img src="https://img.shields.io/travis/schollz/directions.svg?style=flat-square" alt="Build Status"></a>&nbsp;<a href="https://godoc.org/github.com/schollz/directions"><img src="http://img.shields.io/badge/godoc-reference-5272B4.svg?style=flat-square" alt="Go Doc"></a> 

This is a Golang library for direction *extraction* for **any recipe on the internet**. This library compartmentalizes and improves aspects of recipe extraction that I did previously with [schollz/meanrecipe](https://github.com/schollz/meanrecipe) and [schollz/extract_recipe](https://github.com/schollz/extract_recipe).

## How does it work?

See my blog post about it: [schollz.com/blog/ingredients](https://schollz.com/blog/ingredients).

## Develop

If you modify the `corpus/` information then you will need to run 

```
$ go generate
```

before using the library again.

## Contributing

Pull requests are welcome. Feel free to...

- Revise documentation
- Add new features
- Fix bugs
- Suggest improvements

## License

MIT
