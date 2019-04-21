# Go TradeStation Starter

Call Go funcion from EasyLanguage (TradeStation) 🎉🎉🎉

This is just a simple example.

* for more details: https://blog.yat1ma30.com/posts/call-golang-function-from-easylanguage/

## Env

* Windows 10 64-bit
* TradeStation 9.5
* gcc 5.1.0 (tdm-1)
  * http://tdm-gcc.tdragon.net/download
  * append `C:\<your-folder>\TDM-GCC-32\bin\` to path
* go version go1.12.1 windows/amd64
* [Cmder](https://cmder.net/) (whatever)


## What you can do with this

You can call `CalcDouble` function (Go),

```go
//export CalcDouble
func CalcDouble(x float64) float64 {
	return 2 * x
}
```

from EasyLanguage.

```
using elsystem;

external: "yat1ma30\example.dll", double, "CalcDouble", double;

once begin
  ClearPrintLog;
  Print(CalcDouble(2.5)); // 5
end;
```

## Usage

Prepare 32 bit Go first.

```sh
set GOARCH=386
set goroot_bootstrap=%goroot%
cd %goroot%\src
make.bat --no-clean
```

Create a static library. The following command exports `example.a` and `example.h`.

```sh
set GOARCH=386
set CGO_ENABLED=1
set CC=gcc
go build -o libexample.a -buildmode=c-archive main.go
```

Create a dll using the C wrapper.

```sh
gcc -o example.dll example.def -mdll example.cpp -L. -lexample -lwinmm -lws2_32 -lntdll
```

Copy the dll file into TradeStation folder.

e.g. `C:\Program Files (x86)\TradeStation 9.5\Program\yat1ma30\example.dll`

Then call the function from EasyLanguage.

```
using elsystem;

external: "yat1ma30\example.dll", double, "CalcDouble", double;

once begin
  ClearPrintLog;
  Print(CalcDouble(2.5)); // 5
end;
```

## LICENCE

MIT
