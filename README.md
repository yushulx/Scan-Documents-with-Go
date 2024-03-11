# Web Document Scanning and Uploading with Dynamic Web TWAIN & Go
This sample demonstrates how to implement a simple web document scanning and uploading application using [Dynamic Web TWAIN](https://www.dynamsoft.com/web-twain/overview/?utm_content=nav-products) and Go.

## Prerequisites
- [Go](https://go.dev/dl/)
- [Dynamic Web TWAIN License Key](https://www.dynamsoft.com/customer/license/trialLicense?product=dwt)

## How to Run
1. Replace `LICENSE-KEY` in `index.html` with your own license key.
    ```html
    Dynamsoft.DWT.ProductKey = 'LICENSE-KEY';
	```
2. Run the server:

	```bash
	go run server.go
	```

3. Open `localhost:2024` in your browser.

    ![upload document images to a Go web server](https://www.dynamsoft.com/codepool/img/2024/03/dynamic-web-twain-nodejs-document-scan.png)

## Blog
[How to Scan and Upload Documents with Web TWAIN and Go](https://www.dynamsoft.com/codepool/scan-and-upload-documents-in-web-twain-and-go.html)







