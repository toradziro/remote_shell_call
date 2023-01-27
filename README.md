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

Also you can test it with curl, example command-line command:
```
curl --insecure --request POST 'https://127.0.0.1:8080/api/v1/remote-execution' --data-raw '{ "cmd": "ls -lah /", "os": "linux", "stdin":"" }'
```
Expected command line responce will look like:

```
{"stdout":"total 0\ndrwxr-xr-x   1 root root  174 Nov 11 10:03 .\ndrwxr-xr-x   1 root root  174 Nov 11 10:03 ..\ndrwxr-xr-x   1 root root 1.8K Nov 11 09:11 bin\ndrwxr-xr-x   1 root root 1.3K Nov 11 09:17 boot\ndrwxr-xr-x  19 root root 4.2K Jan 27 08:46 dev\ndrwxr-xr-x   1 root root 5.2K Jan 13 19:23 etc\ndrwxr-xr-x   1 root root   18 Oct  3 19:13 home\ndrwxr-xr-x   1 root root  100 Oct  3 19:00 lib\ndrwxr-xr-x   1 root root 3.0K Nov 11 09:11 lib64\ndrwxr-xr-x   1 root root    0 Oct  8 11:50 media\ndrwxr-xr-x   1 root root    0 Mar 15  2022 mnt\ndrwxr-xr-x   1 root root    0 Mar 15  2022 opt\ndr-xr-xr-x 290 root root    0 Jan 27 08:45 proc\ndrwx------   1 root root  114 Nov 11 10:51 root\ndrwxr-xr-x  46 root root 1.3K Jan 27 10:10 run\ndrwxr-xr-x   1 root root 3.8K Nov 11 09:11 sbin\ndrwxr-xr-x   1 root root    0 Mar 15  2022 selinux\ndrwxr-xr-x   1 root root   76 Jan 26 15:30 snap\ndrwxr-x---   1 root root   86 Jan 16 15:01 .snapshots\ndrwxr-xr-x   1 root root   28 Oct  3 19:20 srv\ndr-xr-xr-x  13 root root    0 Jan 27 08:45 sys\ndrwxrwxrwt   1 root root 4.7K Jan 27 10:45 tmp\ndrwxr-xr-x   1 root root  110 Oct  3 18:56 usr\ndrwxr-xr-x   1 root root  118 Nov 11 10:04 var\n","stderr":""}
```
