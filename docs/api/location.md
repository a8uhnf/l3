# DNS location API

Location API calculates and provides the location of databank to the drone. It expects the drone coordinates and velocity as string
and response back the location as stringified floating point value. 

## API Description

```text
POST /api/v1/location
Request:
    required string x
    required string y
    required string z
    required string vel

Respond:
    string loc    
  
```

- Sample Request Body:
```JSON
    {
      "x": "12.7",
      "y": "13.2",
      "z": "123.2",
      "vel": "1.0"
    }
```

- Success Response Body:
```json
    {
      "loc": 150.1
    }
```
    
- Sample Curl Command:
```bash
      $ curl --location --request POST 'localhost:8080/api/v1/location' \
        --header 'Content-Type: application/json' \
        --data-raw '{
      	  "x": "12.7",
      	  "y": "13.2",
      	  "z": "123.2",
      	  "vel": "1.0"
        }'
```