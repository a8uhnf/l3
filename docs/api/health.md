# Health Endpoints

There're two endpoint to determine web server is in good condition to serve the request.

- `/ready` determine web server is okay to receive request. Success Http Code: `200`
- `/live` determine whether web server needs to be restarted or not. Success Http Code: `200`