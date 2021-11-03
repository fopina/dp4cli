# dp4cli
Vasco Digipass CLI implementation

This is the result of [this post](https://rpg.skmobi.com/posts/0x478c_dll/), checking how to call an undocumented (proprietary) DLL. [Vasco Digipass](https://www.onespan.com/support/security/product-life-cycle) in this case.

## Usage

One time setup

```
$ dp4cli.exe -setup
Serial: 123456
Activation Code: 111222333445566
```

Generate an OTP

```
$ dp4cli.exe
251923
```
