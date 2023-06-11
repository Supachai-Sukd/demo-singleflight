# demo-singleflight


### Clone project
```
git clone https://github.com/Supachai-Sukd/demo-singleflight.git
```

### Install Apache benchmark
```
Windows
choco install apache-httpd
source: https://chocolatey.org/packages/apache-httpd

Ubuntu
apt-get install apache2-utils 
source: https://bobcares.com/blog/apache-benchmark-install-ubuntu/

Mac
installed by default in macos
```


-n 100 หมายถึงจำนวนคำขอทั้งหมดที่จะส่ง  
-c 10 หมายถึงจำนวนการร้องขอพร้อมกัน (concurrency)
### Test with singleflight
```
ab -n 100 -c 10 http://localhost:8080/singleflight
```

### Test non singleflight
```
ab -n 100 -c 10 http://localhost:8080/normal
```