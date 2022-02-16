Untuk menjalankan program REST API menggunakan tools POSTMAN

build main.go , atau langsung menjalankan file yang telah di build mncbank.exe
folder documents dan isinya dijadikan satu dengan lokasi mncbank.exe ataupun main.go

URL endpoint : 

1.Login (POST) : localhost:8080/mncbank/login
 parameter : username , password
 *note : username dan password dapat dilihat di dalam file json documents/customer.json

2.Logout (POST) : localhost:8080/mncbank/logout
 parameter : username
headers : Authorization
*note : Authorization didapatkan memalui respon saat melakukan login

3. Pembayaran (POST) : localhost:8080/mncbank/
parameter : username , id_merchant , saldo
headers : Authorization (*token yang didapatkan saat login)
*note : 
- parameter saldo adalah saldo yang akan dibayarkan pada merchant
- Authorization didapatkan memalui respon saat melakukan login
- username dapat dilihat di dalam file json documents/customer.json
- id_merchant dapat dilihat di dalam file json documents/merchant.json

Example :

  1. Login :
  
  ![alt text](https://github.com/denicandraa/testmnc/blob/master/screenshoot/login.jpg?raw=true)
  
  2. Pembayaran :
  
  ![alt text](https://github.com/denicandraa/testmnc/blob/master/screenshoot/pembayaran.jpg?raw=true)
