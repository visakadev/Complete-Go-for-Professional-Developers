# Complete-Go-for-Professional-Developers

To run the application, use:
```sh
go run main.go
```

> **Note:**  
> If you see an error like `listen tcp: lookup tcp/%d8080: unknown port`, check your code for incorrect usage of `fmt.Sprintf` or string formatting when specifying the port.  
> The correct way to listen on port 8080 is:
> ```go
> http.ListenAndServe(":8080", nil)
> ```