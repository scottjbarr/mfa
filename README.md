# MFA

This package wraps some common MFA functionality I often need.

All the heavy lifting is done by the [gotp]() package.


## Utilities

### mfa-secret

Generate an MFA secret.

```
$ mfa-secret
7QXC46HFGYZJCWKPWWYA

$ mfa-secret --help
Usage of mfa-secret:
  -length int
         (default 20)
```


## References

- https://www.gojek.io/blog/a-diy-two-factor-authenticator-in-golang
- https://github.com/xlzd/gotp


## License

The MIT License (MIT)

Copyright (c) 2021 Scott Barr

See [LICENSE](LICENSE)
