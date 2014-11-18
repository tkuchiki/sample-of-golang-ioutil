sample-of-golang-ioutil
=======================

Sample of Golang io/ioutil (http://golang.org/pkg/io/ioutil/)

## 関連ドキュメント、src

- http://golang.org/pkg/io/ioutil/
- http://golang.org/pkg/io/io/
- http://golang.org/pkg/syscall/
- http://golang.org/pkg/os/
- https://golang.org/src/pkg/io/ioutil/ioutil.go
- http://golang.org/src/pkg/io/ioutil/tempfile.go
- http://golang.org/src/pkg/io/ioutil/ioutil_test.go
http://golang.org/src/pkg/io/ioutil/tempfile_test.go

## ioutil.Discard

http://play.golang.org/p/fSK8YRPIvO
  
See: https://blog.golang.org/race-detector


## NopCloser

no-op(ノーオペレーション)な Close メソッドを持つ ReadCloser を返す

~~~~
type nopCloser struct {
	io.Reader
}

func (nopCloser) Close() error { return nil }

// NopCloser returns a ReadCloser with a no-op Close method wrapping
// the provided Reader r.
func NopCloser(r io.Reader) io.ReadCloser {
	return nopCloser{r}
}
~~~~

## ReadAll

エラーまたは EOF まで読み込んだデータを r(io.Reader) として返す。  
成功すれば `err == nil` で、`err == EOF` ではない。  
src から EOF まで読み込むように定義されているので、EOF をエラーとして扱わない。

~~~~
$ go run readall.go test.json
Type: []uint8

{
  "foo": "bar",
  "hoge": "hogehoge"
}
~~~~

~~~~
$ go run readall.go
<!doctype html>
<html>
<head>
    <title>Example Domain</title>

    <meta charset="utf-8" />
    <meta http-equiv="Content-type" content="text/html; charset=utf-8" />
    <meta name="viewport" content="width=device-width, initial-scale=1" />
    <style type="text/css">
    body {
        background-color: #f0f0f2;
        margin: 0;
        padding: 0;
        font-family: "Open Sans", "Helvetica Neue", Helvetica, Arial, sans-serif;

    }
    div {
        width: 600px;
        margin: 5em auto;
        padding: 50px;
        background-color: #fff;
        border-radius: 1em;
    }
    a:link, a:visited {
        color: #38488f;
        text-decoration: none;
    }
    @media (max-width: 700px) {
        body {
            background-color: #fff;
        }
        div {
            width: auto;
            margin: 0 auto;
            border-radius: 0;
            padding: 1em;
        }
    }
    </style>
</head>

<body>
<div>
    <h1>Example Domain</h1>
    <p>This domain is established to be used for illustrative examples in documents. You may use this
    domain in examples without prior coordination or asking for permission.</p>
    <p><a href="http://www.iana.org/domains/example">More information...</a></p>
</div>
</body>
</html>

~~~~

## ReadDir

~~~~
$ ls -l dir
total 8
-rw-r--r--  1 tkuchiki  tkuchiki  40 11 18 07:28 test.json
drwxr-xr-x  2 tkuchiki  tkuchiki  68 11 18 09:13 testdir
~~~~

~~~~~
$ go run readdir.go dir
Name: test.json
Size: 40
Mode: -rw-r--r--
ModTime: 2014-11-18 07:28:06 +0900 JST
IsDir: false

Sys: *syscall.Stat_t
Stat_t.Dev: 16777217
Stat_t.Ino: 71215917
Stat_t.Mode: 33188
Stat_t.Nlink: 1
Stat_t.Uid: 502
Stat_t.Gid: 0
Stat_t.Rdev: 0
Stat_t.Size: 40
Stat_t.Blksize: 4096
Stat_t.Blocks: 8

Name: testdir
Size: 68
Mode: drwxr-xr-x
ModTime: 2014-11-18 09:13:21 +0900 JST
IsDir: true

Sys: *syscall.Stat_t
Stat_t.Dev: 16777217
Stat_t.Ino: 71217047
Stat_t.Mode: 16877
Stat_t.Nlink: 2
Stat_t.Uid: 502
Stat_t.Gid: 0
Stat_t.Rdev: 0
Stat_t.Size: 68
Stat_t.Blksize: 4096
Stat_t.Blocks: 0
~~~~

## ReadFile

~~~~
$ go run readfile.go test.json
Type: []uint8

{
  "foo": "bar",
  "hoge": "hogehoge"
}
~~~~

## TempDir

~~~~
$ go run tempdir.go
/var/folders/j6/msg_dpzn7yjbjrr2yd4x0wkr0000gp/T/579413645
dir/__prefix__928663688
~~~~

## TempFile, WriteFile

~~~~
$ go run tempfile_writefile.go
*os.File
dir/__prefix__509649609

/var/folders/j6/msg_dpzn7yjbjrr2yd4x0wkr0000gp/T/768888212
###############
hoge
foobar
###############

~~~~

