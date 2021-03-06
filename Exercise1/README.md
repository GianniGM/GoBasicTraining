# tedfeed go package

### Summary
We'll create an application to download the newest videos from the popular website TED.com.
We are in no way associated with TED.com, but we love it! :)
To do this we will use the [atom feed](https://en.wikipedia.org/wiki/Atom_(standard)) the website provides.
We'll learn how to interact with the OS environment, how to deal with HTTP requests and buffers.

### Setup go package
Let's setup the go package directory

    $> mkdir -p $GOPATH/src/tedfeed/
    $> mkdir -p $GOPATH/src/tedfeed/cmd


### Setup application home folder
Create main package and main function

Hint: it is common for go applications to have modules for binaries in the cmd folder.

Write a go application to check the existence of the folder "~/TedFeed", "~/TedFeed/Videos" if they don't exists, create them.

Hint: **user.Current**, **os.Getenv**, **os.Stat & os.IsNotExist**, **os.Stat**, **os.Mkdir**, **os.MkdirAll**, **filepath.Join**


### Download the Ted.com atom feed
Download the atom feed from ted.com: "https://www.ted.com/talks/atom" and print the size of the
feed to the screen.

Hint: **http.Get**, **ioutil.ReadAll**

Hint: don't forget to close the body of the response, defer is your friend.


### References
* [os package](https://golang.org/pkg/os/)
* [user package](https://golang.org/pkg/os/user/)
* [net/http package](https://golang.org/pkg/net/http/)
* [io/ioutil package](https://golang.org/pkg/io/ioutil/)
