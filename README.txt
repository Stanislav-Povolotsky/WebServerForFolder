Simple Web-server to share current folder over the HTTP
Format:
	WebServerForFolder [<port> [<folder>]]
where
	<port>   - HTTP-port (80 by default)
	<folder> - path to the folder to share (current folder by default)

Examples:
	# Share current folder. Web-server will be available http://<computername>:80/
	WebServerForFolder

	# Share folder c:\Temp. Web-server can be reached by http://<computername>:8080/
	WebServerForFolder 8080 c:\Temp
