## cross build error

### ld: library not found for -lcrt0.o

```log
/opt/homebrew/opt/go/libexec/pkg/tool/darwin_arm64/link: running cc failed: exit status 1
ld: library not found for -lcrt0.o
clang: error: linker command failed with exit code 1 (use -v to see invocation)
```
env

```log
# macOS Apple M1

ProductName:            macOS
ProductVersion:         13.5
BuildVersion:           22G74
```

- [https://github.com/golang/go/issues/50662](https://github.com/golang/go/issues/50662)
- [https://github.com/golang/go/issues/50669](https://github.com/golang/go/issues/50669)

macOS doesn't support it. also, that's not a Go issue

- [https://stackoverflow.com/questions/5259249/creating-static-mac-os-x-c-build](https://stackoverflow.com/questions/5259249/creating-static-mac-os-x-c-build)