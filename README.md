# go-yadisk

`go-yadisk` is a Go SDK for interacting with Yandex.Disk using its REST API.  
It provides a simple interface for managing files and folders, uploading, downloading, and retrieving disk information.

## Installation

To install `go-yadisk`, use the following command:

```bash
go get github.com/A-ndrey/go-yadisk
```

## Usage

Here’s a basic example of using `go-yadisk` to retrieve disk information:

```go
package main

import (
    "context"
    "fmt"
    "log"

    "github.com/A-ndrey/go-yadisk"
)

func main() {
    // Create a new client with your OAuth token
    client := yadisk.NewClient("YOUR_OAUTH_TOKEN", yadisk.DefaultHost)

    // Get disk information
    diskInfo, err := client.DiskMetaInfo(context.Background(), yadisk.DiskMetaInfoParams{})()
    if err != nil {
        log.Fatal(err)
    }

    fmt.Printf("Free space: %d\n", diskInfo.UsedSpace)
    fmt.Printf("Total space: %d\n", diskInfo.TotalSpace)
}
```

## Features

`go-yadisk` supports the following features:

- Retrieve disk information.
- Upload files to Yandex.Disk.
- Download files from Yandex.Disk.
- Delete files and folders.
- List files and folders.
- Work with public resources.

## Authorization

To use `go-yadisk`, you need an OAuth token.  
You can obtain it by creating an application in [Yandex.OAuth](https://oauth.yandex.com/).

## Documentation

Detailed API documentation is available on the [Yandex.Disk REST API official site](https://yandex.com/dev/disk-api/doc/en/).

## License

This project is licensed under the MIT License.  
See the [LICENSE](LICENSE) file for details.

## Contributing

Contributions are welcome!  
Feel free to open [issues](https://github.com/A-ndrey/go-yadisk/issues) or submit [pull requests](https://github.com/A-ndrey/go-yadisk/pulls).

---

*Note: This project is not an official SDK by Yandex and is not affiliated with or supported by Yandex.*
