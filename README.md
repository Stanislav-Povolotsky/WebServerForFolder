# WebServerForFolder

Simple Web-server to share current folder over the HTTP  
  
Format:  
```shell
	WebServerForFolder [<port> [<folder>]]  
```
where  
>	\<port\>   - HTTP-port (80 by default)  
>	\<folder\> - path to the folder to share (current folder by default)  

Examples:
```shell
	# Share current folder. Web-server will be available http://<computername>:80/
	WebServerForFolder
```

```shell
	# Share folder c:\Temp. Web-server can be reached by http://<computername>:8080/
	WebServerForFolder 8080 c:\Temp
```
  
![Example:](https://user-images.githubusercontent.com/19610545/127737876-3a14ebe2-3cb8-4de9-981c-de3b4586aaa7.gif)
