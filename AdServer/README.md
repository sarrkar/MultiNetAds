endpoints:
+ `GET`: `adserver.local/api/ad`
    * `res`: json:
    ```json
    {
        "ad": {
            "Title": "title1",
            "ImageUrl": "image.storage/media/image12.jpg",
            "ImpressionUrl": "eventserver.local/impression/12/EVoSxr4loX",
            "ClickUrl": "eventserver.local/click/12/c6l1kfnery"
        }
    }
    ``` 


expected endpoints:
+ `GET`: `panel.local/api/ads` 
    * `res`: json:
    ```json
    {
        "status": "OK",
		"code":   200,
        "data": [
            {
                "ID": 12,
                "Title": "title1",
                "ImageUrl": "image.storage/media/image12.jpg",
                "BID": 100
            },
            {
                "ID": 14,
                "Title": "title2",
                "ImageUrl": "image.storage/media/image14.jpg",
                "BID": 150
            }
        ]
    }
    ```
+ `GET` `eventserver.local/impression/:id/:OTLkey`
+ `GET` `eventserver.local/click/:id/:OTLkey`

in go:

```go
c.IndentedJSON(http.StatusOK, gin.H{
        "status": "OK",
        "code":   200,
        "data": AdResponse{
        {12, "title1", "image.storage/media/image12.jpg", 100},
        {14, "title2", "image.storage/media/image14.jpg", 150},
    },
})
```