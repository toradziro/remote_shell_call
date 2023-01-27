### Summary
Test task on position of junior Golang developer. Goal of the task is to create simple server on Golang that can run processes remotly and return result of its execution in JSON format. To configure server simply rewrite server_config.yaml file Most important fields are: endpoint - path in your post request that execute command, aka endpoint, wrong endpoind will be cause of internal error server_crt - name of generated sertificate (also has to be changed in configure.sh file) server_key - name of generated ky (also has to be changed in configure.sh file)
### Usage
To generate sertificate, key and to build server you cat run ```./configure.sh``` in current directory, after that just press enter for all key-generation questions

Example of work (can be tested via postman):

```
Request data:
POST  [https://127.0.0.1:8080/api/v1/remote-execution](https://127.0.0.1:8080/api/v1/remote-execution)
[{"cmd": "df -h", "os": "linux", "stdin":""}]
[{"cmd": "ls -lah", "os": "linux", "stdin":""}]
```

Expected respond form: 
```
Body: { "stdout": "result of program execution",
"stderr": "error occured during exetion" }
```