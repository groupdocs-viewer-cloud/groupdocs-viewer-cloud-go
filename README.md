# GroupDocs.Viewer Cloud SDK for Go

This repository contains the GroupDocs.Viewer Cloud SDK for Go source code. This SDK allows you to work with GroupDocs.Viewer Cloud REST APIs in your Go applications, enabling you to render documents in HTML, image, or PDF format, with the flexibility to render the whole document or a custom range of pages.

## How to use the SDK?

The complete source code is available in this repository. You can use it directly in your project or install it via Go modules. For more details, please visit our [documentation website](https://docs.groupdocs.cloud/display/viewercloud/Available+SDKs#AvailableSDKs-Go).

## Installation

Install GroupDocs.Viewer Cloud SDK for Go using Go modules:

```shell
go get github.com/groupdocs-viewer-cloud/groupdocs-viewer-cloud-go
```

To add the dependency to your app, copy the following into your `go.mod` and run `go mod tidy`:

```go
require (
    github.com/groupdocs-viewer-cloud/groupdocs-viewer-cloud-go latest
)
```

## Getting Started

Below is an example demonstrating how to convert and download a document using GroupDocs.Viewer Cloud SDK for Go:

```go
package main

import (
    "context"
    "fmt"
    "os"

    viewer "github.com/groupdocs-viewer-cloud/groupdocs-viewer-cloud-go"
    "golang.org/x/oauth2/clientcredentials"
)

func main() {
    // Get Client Id and Client Secret from https://dashboard.groupdocs.cloud
    conf := &clientcredentials.Config{
        ClientID:     "YOUR_CLIENT_ID",
        ClientSecret: "YOUR_CLIENT_SECRET",
        TokenURL:     "https://api.groupdocs.cloud/connect/token",
    }

    token, err := conf.Token(context.Background())
    if err != nil {
        fmt.Printf("Unable to retrieve token: %v\n", err)
        return
    }

    cfg := viewer.NewConfiguration()
    client := viewer.NewAPIClient(cfg)
    auth := context.WithValue(context.Background(), viewer.ContextAccessToken, token.AccessToken)

    // Convert and download document as JPG
    file, _ := os.Open("myfile.docx")
    convertedFile, _, err := client.ViewApi.ConvertAndDownload(auth, "jpg", file, nil)
    if err != nil {
        fmt.Printf("Conversion error: %v\n", err)
        return
    }
    fmt.Printf("Converted file: %v\n", convertedFile.Name())
}
```

Below is an example demonstrating how to upload a document, render it, and download the result using GroupDocs.Viewer Cloud SDK for Go:

```go
package main

import (
    "context"
    "fmt"
    "os"

    viewer "github.com/groupdocs-viewer-cloud/groupdocs-viewer-cloud-go"
    "golang.org/x/oauth2/clientcredentials"
)

func main() {
    conf := &clientcredentials.Config{
        ClientID:     "YOUR_CLIENT_ID",
        ClientSecret: "YOUR_CLIENT_SECRET",
        TokenURL:     "https://api.groupdocs.cloud/connect/token",
    }

    token, err := conf.Token(context.Background())
    if err != nil {
        fmt.Printf("Unable to retrieve token: %v\n", err)
        return
    }

    cfg := viewer.NewConfiguration()
    client := viewer.NewAPIClient(cfg)
    auth := context.WithValue(context.Background(), viewer.ContextAccessToken, token.AccessToken)

    // Upload a file to cloud storage
    file, _ := os.Open("myfile.docx")
    uploadResult, _, err := client.FileApi.UploadFile(auth, "myfile.docx", file, nil)
    if err != nil {
        fmt.Printf("Upload error: %v\n", err)
        return
    }
    fmt.Printf("Uploaded: %+v\n", uploadResult)

    // Render it to HTML
    viewOptions := viewer.ViewOptions{
        FileInfo: &viewer.FileInfo{
            FilePath: "myfile.docx",
        },
        ViewFormat: "HTML",
        OutputPath: "myfile.html",
    }
    viewResult, _, err := client.ViewApi.CreateView(auth, viewOptions)
    if err != nil {
        fmt.Printf("Render error: %v\n", err)
        return
    }
    fmt.Printf("View created: %+v\n", viewResult)

    // Download the result
    downloadedFile, _, err := client.FileApi.DownloadFile(auth, "myfile.html", nil)
    if err != nil {
        fmt.Printf("Download error: %v\n", err)
        return
    }
    fmt.Printf("Downloaded file: %v\n", downloadedFile.Name())
}
```

## Licensing

GroupDocs.Viewer Cloud Go SDK is licensed under [MIT License](LICENSE).

## Resources

+ [**Website**](https://www.groupdocs.cloud)
+ [**Product Home**](https://products.groupdocs.cloud/viewer)
+ [**Documentation**](https://docs.groupdocs.cloud/display/viewercloud/Home)
+ [**Free Support Forum**](https://forum.groupdocs.cloud/c/viewer)
+ [**Blog**](https://blog.groupdocs.cloud/category/viewer)

## Contact Us

Your feedback is very important to us. Please feel free to contact us using our [Support Forums](https://forum.groupdocs.cloud/c/viewer)