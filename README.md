# GroupDocs.Viewer Cloud Go SDK

Go package for communicating with the GroupDocs.Viewer Cloud API

## Installation

The package is available at [github.com](https://github.com/groupdocs-viewer-cloud/groupdocs-viewer-cloud-go). You can install it with:

```shell
go get github.com/groupdocs-viewer-cloud/groupdocs-viewer-cloud-go
```

To add dependency to your app copy following into your go.mod and run `go mod tidy`:

```go
require (
 github.com/groupdocs-viewer-cloud/groupdocs-viewer-cloud-go v1.0.0
)
```

## Getting Started

Please follow the [installation](#installation) procedure and then run the following code:

```go
package main

import (
 "context"
 "fmt"

 "github.com/antihax/optional"
 viewer "github.com/groupdocs-viewer-cloud/groupdocs-viewer-cloud-go"
 "golang.org/x/oauth2/clientcredentials"
)

func main() {

 // Create an OAuth2 config using client_credentials grant type
 conf := &clientcredentials.Config{
  ClientID:     "XXXX-XXXX-XXXX-XXXX",      // Your client_id
  ClientSecret: "XXXXXXXXXXXXXXXX",          // Your client_secret
  TokenURL:     "https://api.groupdocs.cloud/connect/token", // The token URL
  Scopes:       []string{},                                  // Optional: specify any required scopes
 }

 // Request the token
 token, err := conf.Token(context.Background())
 if err != nil {
  fmt.Printf("Unable to retrieve token: %v", err)
 }

 // Output the access token
 fmt.Printf("Access Token: %s\n", token.AccessToken)

 cfg := viewer.NewConfiguration()
 client := viewer.NewAPIClient(cfg)
 auth := context.WithValue(context.Background(), viewer.ContextAccessToken, token.AccessToken)

result, _, _ := client.InfoApi.GetSupportedFileFormats(auth)

 for _, format := range result.Formats {
  fmt.Printf("%s (%s)\n", format.FileFormat, format.Extension)
 }
}
```

## Licensing

GroupDocs.Viewer Cloud Go SDK licensed under [MIT License](LICENSE).

## Resources

+ [**Website**](https://www.groupdocs.cloud)
+ [**Product Home**](https://products.groupdocs.cloud/viewer)
+ [**Documentation**](https://docs.groupdocs.cloud/display/viewercloud/Home)
+ [**Free Support Forum**](https://forum.groupdocs.cloud/c/viewer)
+ [**Blog**](https://blog.groupdocs.cloud/category/viewer)

## Contact Us

Your feedback is very important to us. Please feel free to contact us using our [Support Forums](https://forum.groupdocs.cloud/c/viewer).
