## Data Collection Tree

### Simple Backend Service using REST API





Run the Application
```

1 -> Runt the Application using CMD : go run main.go
2 -> THe Application will serve over the 8080 port on your local machine

```

Run the Test case

```
1. go test ./... -coverprofile=coverage.out
2. go tool cover -html=coverage.out
```


#### POST request

```
Route -> /v1/insert
Body -> Requried
Sample -> 
{
	"dim": [{
	"key": "device",
	"val": "mobile"
	},
	{
	"key": "country",
	"val": "UK"
	}],
	"metrics": [{
	"key": "webreq",
	"val": 70
	},
	{
	
	"key": "timespent",
	"val": 30
	}]
	}
```


#### Get Request

```
Route -> /v1/query
QueryParameter -> country
Sample Result -> {
    "error": false,
    "result": {
        "dim": [
            {
                "key": "country",
                "val": ""
            }
        ],
        "metrics": [
            {
                "key": "webreq",
                "val": 0
            },
            {
                "key": "timespent",
                "val": 0
            }
        ]
    }
}
```
