# goango
## goango is base64encoder/decoder.  
- For example, you can use goango when you write secrets.yaml for Kubernetes.  
## Prerequisites  
- golang => 1.12  
## Getting started  
### Clone project
```
$ git clone https://github.com/takapon564/goango.git
$ cd goango
```
### Execute program  
```
$ go run main.go
```  
or  
```
$ ./goango
```  
### Select encode/decode  
```
Do you want to encode, or decode? Please enter: e/d: e
```  
### If you select "e", enter plain text. Follow, enter "password".
```
enter plain text: password
```  
### You can get eocoded text  
```
result:  cGFzc3dvcmQ=
```
