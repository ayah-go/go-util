##/pkg

外部应用程序可以使用的库代码(例如 /pkg/mypubliclib)。其他项目会导入这些库，希望它们能正常工作，所以在这里放东西之前要三思:-)注意，internal 目录是确保私有包不可导入的更好方法，因为它是由 Go 强制执行的。/pkg 目录仍然是一种很好的方式，可以显式地表示该目录中的代码对于其他人来说是安全使用的好方法。由 Travis Jeffery  撰写的 I'll take pkg over internal 博客文章提供了 pkg 和 internal 目录的一个很好的概述，以及什么时候使用它们是有意义的。

当根目录包含大量非 Go 组件和目录时，这也是一种将 Go 代码分组到一个位置的方法，这使得运行各种 Go 工具变得更加容易（正如在这些演讲中提到的那样: 来自 GopherCon EU 2018 的 Best Practices for Industrial Programming , GopherCon 2018: Kat Zien - How Do You Structure Your Go Apps 和 GoLab 2018 - Massimiliano Pippi - Project layout patterns in Go ）。

如果你想查看哪个流行的 Go 存储库使用此项目布局模式，请查看 /pkg 目录。这是一种常见的布局模式，但并不是所有人都接受它，一些 Go 社区的人也不推荐它。

如果你的应用程序项目真的很小，并且额外的嵌套并不能增加多少价值(除非你真的想要:-)，那就不要使用它。当它变得足够大时，你的根目录会变得非常繁琐时(尤其是当你有很多非 Go 应用组件时)，请考虑一下。

Library code that's ok to use by external applications (e.g., /pkg/mypubliclib). Other projects will import these libraries expecting them to work, so think twice before you put something here :-) Note that the internal directory is a better way to ensure your private packages are not importable because it's enforced by Go. The /pkg directory is still a good way to explicitly communicate that the code in that directory is safe for use by others. The I'll take pkg over internal blog post by Travis Jeffery provides a good overview of the pkg and internal directories and when it might make sense to use them.

It's also a way to group Go code in one place when your root directory contains lots of non-Go components and directories making it easier to run various Go tools (as mentioned in these talks: Best Practices for Industrial Programming from GopherCon EU 2018, GopherCon 2018: Kat Zien - How Do You Structure Your Go Apps and GoLab 2018 - Massimiliano Pippi - Project layout patterns in Go).

Note that this is not a universally accepted pattern and for every popular repo that uses it you can find 10 that don't. It's up to you to decide if you want to use this pattern or not. Regardless of whether or not it's a good pattern more people will know what you mean than not. It is a bit confusing for new Go devs, but it's a pretty simple confusion to resolve and that's one of the goals for this project layout repo.

Ok not to use it if your app project is really small and where an extra level of nesting doesn't add much value (unless you really want to). Think about it when it's getting big enough and your root directory gets pretty busy (especially if you have a lot of non-Go app components).

Examples:

https://github.com/prometheus/prometheus/tree/master/pkg
https://github.com/jaegertracing/jaeger/tree/master/pkg
https://github.com/istio/istio/tree/master/pkg
https://github.com/GoogleContainerTools/kaniko/tree/master/pkg
https://github.com/google/gvisor/tree/master/pkg
https://github.com/google/syzkaller/tree/master/pkg
https://github.com/perkeep/perkeep/tree/master/pkg
https://github.com/minio/minio/tree/master/pkg
https://github.com/heptio/ark/tree/master/pkg
https://github.com/argoproj/argo/tree/master/pkg
https://github.com/heptio/sonobuoy/tree/master/pkg
https://github.com/helm/helm/tree/master/pkg
https://github.com/kubernetes/kubernetes/tree/master/pkg
https://github.com/kubernetes/kops/tree/master/pkg
https://github.com/moby/moby/tree/master/pkg