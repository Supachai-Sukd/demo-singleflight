# demo-singleflight


### Clone project
```
git clone https://github.com/Supachai-Sukd/demo-singleflight.git
```

### Install package
```
go mod tidy
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
-k หมายถึง Keep-Alive คงอยู่ในสถานะเชื่อมต่ออย่างต่อเนื่อง" ซึ่งเป็นการตั้งค่าให้เครื่องผู้ใช้ (client) ใช้การเชื่อมต่อ Keep-Alive เพื่อให้สามารถส่งคำร้องขอ (request) ได้มากกว่าหนึ่งครั้ง โดยไม่ต้องเปิดการเชื่อมต่อใหม่ทุกครั้งที่ส่งคำร้องขอใหม่
### Test with singleflight
```
ab -k -n 100 -c 10 http://localhost:8080/singleflight
```

### Test non singleflight
```
ab -k -n 100 -c 10 http://localhost:8080/normal
```