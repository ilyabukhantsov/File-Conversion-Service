# Mock Documentation

## 🔴 NO CORS

Mock server does not support CORS.

---

## How to start mock server

```bash
go run cmd/server/main.go -mock=true
How to run mock HTTP upload
curl -X POST http://localhost:8080/files/upload \
  -F "upload=@mock.docx"
API Endpoints
Upload file
POST /files/upload
🟢 200 OK (always)
{
  "fileId": "mock200"
}
🔴 500 Internal Server Error (always)
{
  "code": "BAD_REQUEST",
  "details": {
    "additionalProp1": {}
  },
  "message": "Invalid input"
}
Convert file by ID
POST /files/:fileId/convert

Handler: MockConvertFileById

🟢 mock200
Request

POST /files/mock200/convert

Response
{
  "convertedFileId": "def456",
  "status": "success"
}
🔴 mock400
{
  "code": "BAD_REQUEST",
  "message": "Invalid input",
  "details": {
    "additionalProp1": {}
  }
}
🔴 mock404
{
  "code": "NOT_FOUND",
  "message": "File not found",
  "details": {
    "additionalProp1": {}
  }
}
🔴 mock422
{
  "code": "UNPROCESSABLE_ENTITY",
  "message": "Invalid input",
  "details": {
    "additionalProp1": {}
  }
}
🔴 mock500
{
  "code": "INTERNAL_ERROR",
  "message": "Server error",
  "details": {
    "additionalProp1": {}
  }
}
❗ Default (unknown fileId)
{
  "code": "BAD_REQUEST",
  "message": "Invalid input",
  "details": "No validation for other ids in MOCK server, check DOCS for correct ids for tests!"
}
Download converted file
GET /files/:convertedFileId/download

Handler: MockDownloadFileById

🟢 mock200
Request

GET /files/mock200/download

Response
{
  "content": "no file, sry, it's mock server for a reason!"
}
🔴 mock404
{
  "code": "NOT_FOUND",
  "message": "File not found",
  "details": {
    "additionalProp1": {}
  }
}
🔴 mock500
{
  "code": "INTERNAL_ERROR",
  "message": "Server error",
  "details": {
    "additionalProp1": {}
  }
}
❗ Default (unknown convertedFileId)
{
  "code": "BAD_REQUEST",
  "message": "Invalid input",
  "details": "No validation for other ids in MOCK server, check DOCS for correct ids for tests!"
}
Notes!
This is a mock server only
No real file processing happens
mock200, mock400, etc. are test triggers
Any unknown ID returns default error response