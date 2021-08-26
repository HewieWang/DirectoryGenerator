# DirectoryGenerator
```Go
go run main.go  -l abc -r .com -c pkk -o temp.txt
```
output

> abcpkk.com
> abckpk.com
> abckkp.com

```Go
go run main.go  -l abc,qwe -r .com -c pkk -o temp.txt
```

output

> abcpkk.com
> abckpk.com
> abckkp.com
> qwepkk.com
> qwekpk.com
> qwekkp.com

```Go
go run main.go  -c abcdefghijklmnopqrstuvwxyz
go run main.go  -c aaabb555563 -o xx.txt
```
