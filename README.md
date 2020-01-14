# listenyt
[![Go Report Card](https://goreportcard.com/badge/github.com/ShiraazMoollatjie/listenyt?style=flat-square)](https://goreportcard.com/report/github.com/ShiraazMoollatjie/listenyt)
[![Go Doc](https://img.shields.io/badge/godoc-reference-blue.svg?style=flat-square)](http://godoc.org/github.com/ShiraazMoollatjie/listenyt)
[![Build status](https://ci.appveyor.com/api/projects/status/j8rcy5w3gcvatpxq/branch/master?svg=true)](https://ci.appveyor.com/project/ShiraazMoollatjie/listenyt/branch/master)

> I normally have a workflow that involves getting .m4a versions of youtube links (for offline listening). It's a tedious process to copy your youtube link, then fire up youtube-dl (with extra commands) to download the m4a file.

So `listenyt` (pronounced `listen it`) is a small tool that will monitor the clipboard for youtube links. Once it finds a youtube link, it will attempt to download the link in m4a format.

# Pre requisites
You need `youtube-dl` to be installed on your system. `listenyt` will error if it is not on located on the `PATH`.

In some linux distributions, like Manjaro for example, you may need to install either `xclip` or `xsel` for clipboard management. `listenyt` will try to detect whether you need to do this.

# How to run
To run `listenyt`, simply run:

```sh
go get github.com/ShiraazMoollatjie/listenyt
```

This will run the `listenyt` process. As you copy more youtube links, it will be staged for downloading. Downloading is sequential (for now) because you don't really want too many `youtube-dl` instances running.

# Wishlist
This is simply meant to be used for my usecase, but I can see how others may want to use it at a later stage. So the current wishlist looks something like:

* Configurable `youtube-dl` formats because currently it only downloads in `m4a`.
* Configurable `youtube-dl` sites because it supports not only youtube.
* Maybe publish completed downloads to the system notification bar.
* Maybe persisting the downloaded links so that you can cleanup or continue failed downloads.
* Multiple concurrent downloads. Currently it downloads sequentially. I wasn't sure whether spamming multiple `youtube-dl` instances was a good idea at the time.
* Notify via OS notification that a download is starting.

# Legal stuff
This is only meant for personal use. Please don't use this to make money or something weird like that.