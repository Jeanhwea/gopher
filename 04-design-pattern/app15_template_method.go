////////////////////////////////////////////////////////////////////////////////
// 模版方法模式
////////////////////////////////////////////////////////////////////////////////
package main

import "fmt"

type Downloader interface {
	Download(uri string)
}

type template struct {
	implement
	uri string
}

type implement interface {
	download()
	save()
}

func newTemplate(impl implement) *template {
	return &template{implement: impl}
}

func (this *template) Download(uri string) {
	this.uri = uri
	fmt.Println("Prepare downloading")
	this.implement.download()
	this.implement.save()
	fmt.Println("Finish downloading")
}

func (this *template) save() {
	fmt.Println("default save.")
}

////////////////////////////////////////////////////////////////////////////////
type HttpDownloader struct {
	*template
}

func NewHttpDownloader() Downloader {
	downloader := &HttpDownloader{}
	template := newTemplate(downloader)
	downloader.template = template
	return downloader
}

func (this *HttpDownloader) download() {
	fmt.Printf("download %s via http\n", this.uri)
}

func (this *HttpDownloader) save() {
	fmt.Println("Http saved.")
}

////////////////////////////////////////////////////////////////////////////////
type FtpDownloader struct {
	*template
}

func NewFtpDownloader() Downloader {
	downloader := &FtpDownloader{}
	template := newTemplate(downloader)
	downloader.template = template
	return downloader
}

func (this *FtpDownloader) download() {
	fmt.Printf("download %s via ftp\n", this.uri)
}

func main() {
	var downloader Downloader
	downloader = NewHttpDownloader()
	downloader.Download("http://example.com/abc.zip")

	downloader = NewFtpDownloader()
	downloader.Download("ftp://example.com/abc.zip")
}

////////////////////////////////////////////////////////////////////////////////
// Prepare downloading
// download http://example.com/abc.zip via http
// Http saved.
// Finish downloading
// Prepare downloading
// download ftp://example.com/abc.zip via ftp
// default save.
// Finish downloading
