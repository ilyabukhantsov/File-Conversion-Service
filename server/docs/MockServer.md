# Mock documentation

## NO CORS 🔴

### How to start mock server

`go run cmd/server/main.go -mock=true`

### How to run mock http


`curl -X POST http://localhost:8080/files/upload \
  -F "upload=@mock.docx"`

### 🟢 (Always) 200: {

    {
    "fileId":"abc123"
    }
}

### 🔴 (Always) 500: {

    {
    "code": "BAD_REQUEST",
    "details": {
        "additionalProp1": {}
    },
    "message": "Invalid input"
    }
}