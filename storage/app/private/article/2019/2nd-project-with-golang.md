Mình tìm hiểu Golang một cách rất tình cờ và những tò mò của mình được giải toả làm nảy sinh ra những điều thú vị từ Golang, và đây là bài tập nhỏ thứ 2 mà mình thử dùng với Golang.

Mình sẽ xây dụng một Restful API đơn giản gồm 2 endpoint: Authorization Server và Product Application Server

## Authorization Server và JWT (JSON Web Tokens) 
Authorization Server sẽ sử dụng một phương thức xác thực đơn giản là JWT với thuật toán hash HS256 để xác thực việc giao tiếp giữa Product Application Server với Client.

![Workflow](http://localhost:8080/storage/article/2019/client-credentials-grant.png)

Package:
```
go get github.com/dgrijalva/jwt-go
```

### JSON Web Tokens
Mục đích của Token là để kiểm tra xem một request từ Client tới Application Server đã được xác thực hay chưa (hay đã đăng nhập với username/password). 
Client sẽ xác thực với Authorization Server, sau đó Authorization Server sẽ gửi cho Client một Token, kể từ đó việc giao tiếp của Client đến những Application Server sẽ gửi kèm theo Token này nhằm để Application Server xác thực.

Cấu trúc của một JSON Web Token sẽ bao gồm 3 thứ header, payload và signature
**Header** 
```json
{
    "typ": "JWT",
    "alg": "HS256"
}
```
Giá trị của Header JSON bao gồm `typ` dùng để xác định loại token là JWT và `alg` viết tắc của algorithm là thuật toán hash

**Payload**
```json
{
    "userId": 12321,
    "userName": "user1",
    "gender": "male",
    "role": "admin"
}
```
Payload là dữ liệu tự định nghĩa được bao gồm trong một token, nó còn hay được gọi là `claims` trong JWT

**Signature**
Signature là một chữ ký số dùng để xác thực và tin tưởng.

Công thức để tạo ra một signature với thuật toán HMAC và một secret key
```
// signature algorithm
data = base64urlEncode( header ) + “.” + base64urlEncode( payload )
hashedData = hash( data, secret )
signature = base64urlEncode( hashedData )
```
Chúng ta cũng có thể tạo ra một signature dùng public key/private key với thuật toán RSA

**JSON Web Token sẽ nối chúng lại với nhau**
Và cuối cùng một token sẽ là một chuỗi nối giữa header, payload và signature
```
token = base64urlEncode( header ) + “.” + base64urlEncode( payload ) + “.” + signature
```

Chú ý: như đã thấy ở trên header và payload chỉ sử dụng `base64urlEncode` chứ không được mã hoá nhé.

#### SSL handshake steps

Symmetric Algorithm (hay được gọi là secret key algorithm) là một thuật toán mã hoá mà nó dùng chung 1 key cho việc mã hoá và giải mã dữ liệu
SH256

*References:*

* [http://www.golangbootcamp.com/book/](http://www.golangbootcamp.com/book/)
* [Goroutine under the hood](https://devblog.dwarvesf.com/post/go-routine-under-the-hood/)
* [Concurrency vs. Parallelism](https://howtodoinjava.com/java/multi-threading/concurrency-vs-parallelism/)
* [How goroutines work](https://blog.nindalf.com/posts/how-goroutines-work/)