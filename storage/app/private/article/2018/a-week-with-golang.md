Mình tìm hiểu Golang một cách rất tình cờ và những điều lạ lẵm mà Golang mang lại làm mình quyết định lưu lại vài dòng cảm xúc với nó ở đây.

<br>
### SỰ ĐƠN GIẢN CẦN THIẾT
#### Đơn giản trong cách triển khai hay deployment
Về cơ bản **compiler** dùng để biên dịch ngôn ngữ lập trình cấp cao sang ngôn ngữ cấp thấp như Assembly code của gcc, hay Bytecode của Java hoặc là machine code.

Sở dĩ có sự khác nhau như vậy là do đầu vào của những môi trường là khác nhau, như Java chúng ta đã biết, muốn chạy được Java môi trường phải có Java virtual machine(JVM) và đầu vào của JVM là bytecode

Go compiler chọn option thứ 3, biên dịch từ Golang sang machine code luôn, ngôn ngữ mà CPU có thể hiểu được, và điều này dẫn đến việc triển khai hay deployment sẽ trở nên đơn giản hơn, chỉ cần chọn build ở môi một trường tương ứng như Linux, Window(amd32, amd64), MacOS và dùng file build là có thể sử dụng được ngay luôn.
(khác với PHP ngôn ngữ mình đang dùng là thuộc loại thông dịch nên phải cần cài đặt môi trường và số lượng file deploy rất lớn)

#### Golang không có những thuộc tính của một ngôn ngữ hướng đối tượng OOP
Go không dùng Class, cũng không có tính kế thừa, theo mình thì ưu điểm của việc này là làm code dễ đọc và dễ hình dung hơn. Nhưng riêng mình thì không thích điều này lắm.

Nhưng bù lại việc không dùng Class, Golang lại có:
* _Type_: nó dùng để định nghĩa một loại kiểu dữ liệu nào đó.
* _Reciever_: kiểu để Golang tạo ra một function của một Class vậy.
  
Ví dụ so sánh giữa một file PHP và Golang
```PHP
//file index.php
<?php
class A {
    public $pro1;

    function f1() {
        echo 'Hello world';
    }
}

$a = new A();
$a->f1();
```

```go
//file main.go
type A struct {
    pro1 string 
}

//func ([receiver receiverType]) functionName([paramter1],...) ([returnType1],...)
func (a A) f1 {
    fmt.Println("Hello world");
}

func main() {
    a := a{pro1: "value of property pro1"}
	a.f1()
}
```

#### Format với gofmt
Một ví dụ đơn giản về format code của dấu `{}` và khoảng trắng trong `if` sẽ thấy coding standard trong một project và việc review code sẽ khó khăn đến chừng nào
```php
if (a > b) {
    //To do
}
```
```php
if(a > b){
    //To do
}
```
```php
if (a > b) 
{
    //To do
}
```
```php
if (a > b) //To do
```

Mình luôn ao ước sẽ có 1 tool tự động làm việc format code trước khi code của ai đó được push lên remote để tránh việc reivew code không đáng này phải xảy ra.
*gofmt* đứng ra làm việc này để đảm bảo những điều nhỏ nhặc nhất trong một project chung là giống nhau. Như ví dụ trên, gofmt chỉ chấp nhận một kiểu duy nhất
```go
if a > b {
    //To do
}
```
<br>
#### Concurrency Programming
Concurrency được nói rất nhiều và cũng là một ưu điểm lớn để lập trình viên thiết kế những tác vụ chạy đồng thời một cách đơn giản.

Một ví dụ rất quen thuộc về concurrency là:
>Chúng ta hay thách đố một người nào đó dùng tay trái để vẽ hình tròn còn tay phải vẽ hình vuông, hai việc dường như là xảy ra cùng lúc nhưng về bản chất thì không, não người sẽ luân phiên điều khiển tay trái vẽ một đoạn nhỏ hình tròn sau đó chuyển sang điều khiển tay phải vẽ một đoạn nhỏ hình vuông và cứ tiếp tục luân phiên điều khiển như vậy.

Trong concurrency programming, khi họ dùng tính năng CPU time-slicing của hệ điều hành thì sẽ dễ dàng thấy được khi một tác vụ thứ 1 thực thi một phần của nó và sau đó chuyển trạng thái sang waiting thì CPU sẽ assign đến tác vụ thứ 2 và xử lý một phần của task thứ 2, việc chuyển đổi như vậy xảy ra cho đến khi tất cả chúng được xử lý xong.

![Concurrency Image](https://t-rex.click/storage/article/2018/concurrenct.gif)

Ví dụ đơn giản về concurrency trong Golang, check status từ 5 domain phía dưới
```go
func main() {
	links := []string{
		"http://google.com",
		"http://facebook.com",
		"http://golang.org",
		"http://amazon.com",
		"http://stackoverflow.com",
	}

	for _, link := range links {
        //Concurrency with function checkLink
		go checkLink(link)
    }
    
    time.Sleep(10 * time.Second)
}

func checkLink(link string) {
	_, err := http.Get(link)
	if err != nil {
		fmt.Println(link, "is down")
		return
	}

	fmt.Println(link, "is up")
}
```
[Play](https://play.golang.org/p/dMSfnDYSvdA)

#### Goroutine và Channel
Go cho phép chúng ta viết phần mềm theo kiểu concurrent, vậy công cụ là Goroutine.
**Goroutine** là thứ dùng để tạo ra một tác vụ chạy concurrency trong Golang
**Channel** là thứ dùng để chia sẽ dữ liệu giữa những goroutine với nhau

#### Vì Goroutine và thread gần giống nhau, nên mình thử làm một so sánh nhỏ
**Memory sử dụng ít hơn:** Goroutine được đánh giá sử dụng ít memory hơn thread rất nhiều. Mỗi goroutine chỉ tốn vài kb trong stack space và stack size có thể tăng lên hoặc co lại dựa vào nhu cầu sử dụng của ứng dụng, trong khi đó stack size của threads thì fixed.

Số lượng OS thread để ghép nối với Goroutine ít hơn, thậm chí chỉ cần duy nhất một OS thread cho cả ngàn cả ngàn goroutine. OS thread sẽ được sinh ra thêm khi phát hiện có một goroutine bị block.

**Chi phí của một lần process switching thấp hơn:** trong time-sharing model thì hệ điều hành đảm nhiệm việc chỉ điểm cho CPU là process nào nó cần CPU xử, và việc chuyển qua chuyển lại để xử những process của CPU gọi là **process switching**

### THÚ VỊ KHÁC
#### Multiple Return
Rất tiện lợi khi một function mà có thể trả về nhiều kết quả
```go
func swap(x string, y string) (string, string) {
    return y, x
}

func main() {
    a, b := swap("hello", "world")
    fmt.Println(a, b)
    //Result: world hello
}
```

*References:*

* [http://www.golangbootcamp.com/book/](http://www.golangbootcamp.com/book/)
* [Goroutine under the hood](https://devblog.dwarvesf.com/post/go-routine-under-the-hood/)
* [Concurrency vs. Parallelism](https://howtodoinjava.com/java/multi-threading/concurrency-vs-parallelism/)
* [How goroutines work](https://blog.nindalf.com/posts/how-goroutines-work/)