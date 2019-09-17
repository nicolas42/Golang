[https://stackoverflow.com/questions/26744814/serve-image-in-go-that-was-just-created](https://stackoverflow.com/questions/26744814/serve-image-in-go-that-was-just-created)

[https://stackoverflow.com/questions/40316052/in-memory-file-for-testing](https://stackoverflow.com/questions/40316052/in-memory-file-for-testing)

bytes buffers can be used in place of files :) [https://yourbasic.org/golang/io-writer-interface-explained/](https://yourbasic.org/golang/io-writer-interface-explained/)  

Take the address of a bytes buffer for file like behaviour.  This is a good generic reader writer

    var buf bytes.Buffer
    fmt.Fprintf(&buf, "Size: %d MB.", 85)
    s := buf.String()) // s == "Size: 85 MB."