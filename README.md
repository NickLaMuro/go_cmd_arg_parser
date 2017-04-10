go_cmd_arg_parser
=================

Parses a string, representative of a shell command, which is split into it's
components and can be passed to `os/exec`'s `Run()` function.

A large amount of inspiration for this code base is taken from the following
stackoverflow post:

http://stackoverflow.com/questions/39151420/golang-executing-command-with-spaces-in-one-of-the-parts

Since I am not much of a go programmer, it was the best I could dig up... if
other examples exist out there, you are probably better off using that...


Usage
-----

```go

import "github.com/NickLaMuro/go_cmd_arg_parser"

func main() {
  parseFromCmdString(`date -j -f "%a %b %d \"%T %Z %Y\"" "01/01/1900" "+%s"`)
  // returns => ["date", "-j", "-f", "%a %b %d \"%T %Z %Y\"", "01/01/1900", "+%s"]
}
```


License
-------

As this code is dirivative of a [stackoverflow](http://stackoverflow.com)
question/answer, the same license that [all answeres are lisenced under][1],
which is the [Creative Commons Attribution Share Alike license][2]

Original credit should go to the author of the [original post][2].


[1]: https://stackexchange.com/legal#3SubscriberContent
[2]: https://creativecommons.org/licenses/by-sa/3.0/
[3]: http://stackoverflow.com/questions/39151420/golang-executing-command-with-spaces-in-one-of-the-parts
