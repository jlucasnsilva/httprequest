# Requestconv

Work in progress...

```go
type (
    Params struct {
        ID int64 `from:"url-param=id"`
        Sort string `from:"url-query=sort"`
        Form Form `from:"request-body"`
    }

    Form struct {
        Name string `json:"name"`
        DateOfBirth time.Time `json:"date_of_birth"`
        Count int `json:"count"`
    }
)

func Handle(w http.ResponseWriter, r *http.Request) {
    var params Params
    err := httprequest.As(r, &params)
}
```
