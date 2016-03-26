# ios-passcode-crack #

A tool for finding a lost iOS restrictions passcode. I wrote this using Go (golang) v1.6

To find your forgotten passcode on iOS, do the following:

1. Backup your device on your Mac, make sure to use a Non-Encrypted backup
2. Navicate to this directory on the filesystem:
   > cd /Users/<username>/Library/Application Support/MobileSync/Backup/<latest backup folder>
3. Open or cat the following file
   > cat 398bc9c2aeeab4cb0c12ada0f52eea12cf14f40b
4. Install this tool or build it from source
5. When the tool asks, enter the values from the RestrictionsPasswordKey field and the RestrictionsPasswordSalt field

## Binary Releases

I have produced binaries for Mac OSX 10.11.4, please look in the releases section for the zip files 

## Installation From Source##

```
go/src/> go get github.com/jordan2175/ios-passcode-crack
go/src/> cd github.com/jordan2175/ios-passcode-crack
go/src/github.com/jordan2175/ios-passcode-crack/> go build iso-passcode-crack.go
```

## Examples ##

```
./ios-passcode-crack
```

## Contributing ##

Contributions welcome! Please fork the repository and open a pull request
with your changes or send me a diff patch file.

## License ##

This is free software, licensed under the Apache License, Version 2.0.

