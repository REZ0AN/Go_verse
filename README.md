# Go Installation

## Download Source
Download source from this link  [Go Binaries](https://go.dev/dl/)
## Unpack Packages 
After downloading the binary file may be in your `~/Downloads` folder from terminal
use this command
```
cd Downloads
```
Then list the items in the directory by this command
```
ls -l | grep go
```
![ls -l | grep go](https://github.com/REZ-OAN/Go_verse/blob/main/images/ls.png)


Extract the binary files from the `go1.22.2.linux-amd64.tar.gz` file in the download folder (you may have download the another version) use this command
```
sudo tar -C /usr/local -xzf go1.22.2.linux-amd64.tar.gz
```
This command will extract the file into `/usr/local` folder

## Add to PATH variable
Where the `go binaries` are you need to find it and then the path should be added to you path variable otherwise you will encounter some problems like `go command not found` by this command will be able to locate the `go binary` where it is 
```
whereis go
```
![whereis go](https://github.com/REZ-OAN/Go_verse/blob/main/images/path.png)
Now copy the path and use below command
```
export PATH=$PATH:{your_copied_path_here}/bin
```
## Verify Installation
Now use below command to verify your Installation
```
go version
```
![go version](https://github.com/REZ-OAN/Go_verse/blob/main/images/version.png)
