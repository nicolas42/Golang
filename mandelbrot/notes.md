[https://stackoverflow.com/questions/26744814/serve-image-in-go-that-was-just-created](https://stackoverflow.com/questions/26744814/serve-image-in-go-that-was-just-created)

[https://stackoverflow.com/questions/40316052/in-memory-file-for-testing](https://stackoverflow.com/questions/40316052/in-memory-file-for-testing)

bytes buffers can be used in place of files :) [https://yourbasic.org/golang/io-writer-interface-explained/](https://yourbasic.org/golang/io-writer-interface-explained/)  

Take the address of a bytes buffer for file like behaviour.  This is a good generic reader writer

    var buf bytes.Buffer
    fmt.Fprintf(&buf, "Size: %d MB.", 85)
    s := buf.String()) // s == "Size: 85 MB."


=== More Notes ===

	var buffer bytes.Buffer
	png.Encode(&buffer, img)
	w.Header().Set("Content-Type", "image/jpeg") // <-- set the content-type header
	io.Copy(w, &buffer)


var buffer bytes.Buffer
these little guys are good for when you want a file but don't want a file

to send an image through a responsewriter
set the content type to "image/jpeg"
then write to it with io.Copy
