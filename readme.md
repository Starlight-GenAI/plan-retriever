# Plan retriever service

## To start application
1. Create config.yaml at config folder by using config.example.yaml as reference
2. Add a Google Cloud service account in JSON format at root folder
3. Run
```sh
go build ./cmd/main.go  && ./main
```
4. The application will be running on 127.0.0.1:8080

## To start application using docker
1. Create config.yaml by using config.example.yaml as reference
2. Add a Google Cloud service account in JSON format by updating the Dockerfile at line 17.
3. Run
```sh
docker build . -t plan-retriever
```
4. Run
```sh
docker run -p 8080:8080  -v ./config.yaml:/config plan-retriever -d
```
5. The application will be running on 127.0.0.1:8080

## API document
### 1. POST /plan-status <br />
description: endpoint for checking plan status. There are 3 status which are pending, success, is_not_travel_video and failed
#### Request body
```json
{
    "id": "b9cc5fc8959541c3b9be0c36c653f7f6"
}
```
#### request body detail
```text
id: queue id which getting from uploading video
```
#### Response
```json
{
    "status": "success"
}
```
#### Response detail
```text
status: status of travel plan
```
### 2. POST /get-trip-summary
description: endpoint for getting generated plan.
#### Request body
```json
{
    "id": "b9cc5fc8959541c3b9be0c36c653f7f6"
}
```
id: queue id which getting from uploading video
#### Response
```json
{
    "content": [
        {
            "day": "Day 1",
            "location_with_summary": [
                {
                    "location_name": "Austrian National Library",
                    "summary": "",
                    "place_id": "ChIJ9SMnY5kHbUcRR2-0Bsb53Vc",
                    "lat": 48.2062609,
                    "lng": 16.3668953,
                    "rating": 4.5,
                    "category": "location",
                    "photo": "",
                    "has_recommended_restaurant": false,
                    "recommended_restaurant": {
                        "photo": ""
                    }
                },
                             {
                    "location_name": "Austrian National Library",
                    "summary": "",
                    "place_id": "ChIJ9SMnY5kHbUcRR2-0Bsb53Vc",
                    "lat": 48.2062609,
                    "lng": 16.3668953,
                    "rating": 4.5,
                    "category": "location",
                    "photo": "",
                    "has_recommended_restaurant": false,
                    "recommended_restaurant": {
                        "photo": ""
                    }
                },
            ],
            "count_dining": 0
        },
        {
            "day": "Day 2",
            "location_with_summary": [
                             {
                    "location_name": "Austrian National Library",
                    "summary": "",
                    "place_id": "ChIJ9SMnY5kHbUcRR2-0Bsb53Vc",
                    "lat": 48.2062609,
                    "lng": 16.3668953,
                    "rating": 4.5,
                    "category": "location",
                    "photo": "",
                    "has_recommended_restaurant": false,
                    "recommended_restaurant": {
                        "photo": ""
                    }
                },
                              {
                    "location_name": "Austrian National Library",
                    "summary": "",
                    "place_id": "ChIJ9SMnY5kHbUcRR2-0Bsb53Vc",
                    "lat": 48.2062609,
                    "lng": 16.3668953,
                    "rating": 4.5,
                    "category": "location",
                    "photo": "",
                    "has_recommended_restaurant": false,
                    "recommended_restaurant": {
                        "photo": ""
                    }
                },
            ],
            "count_dining": 0
        }
    ],
    "user_id": ""
}
```
### 3. POST /get-video-summary
description: endpoint for getting all location in video
#### Request body
```json
{
    "id": "b9cc5fc8959541c3b9be0c36c653f7f6"
}
```
id: queue id which getting from uploading video
#### Response
```json
{
        "content": [
        {
            "location_name": "Palace of Justice",
            "start_time": 175,
            "end_time": 191,
            "summary": "",
            "place_id": "ChIJz6MvLWPEw0cRBeLV8QN50AE",
            "lat": 50.8366449,
            "lng": 4.3516068,
            "category": "etc",
            "photo": ""
        },
        {
            "location_name": "Austrian National Library",
            "start_time": 226,
            "end_time": 295,
            "summary": "",
            "place_id": "ChIJ9SMnY5kHbUcRR2-0Bsb53Vc",
            "lat": 48.2062609,
            "lng": 16.3668953,
            "category": "attractions",
            "photo": ""
        },
    ],
    "can_generate_trip": true,
    "user_id": ""
}
```
### 4. POST /get-video-highlight
description: endpoint for listing interesting content in video
#### Request body
```json
{
    "id": "b9cc5fc8959541c3b9be0c36c653f7f6"
}
```
id: queue id which getting from uploading video
#### Response
```json
{
    "content": [
        {
            "hightlight_name": "Palace of Justice",
            "highlight_detail": ""
        },
        {
            "hightlight_name": "Austrian National Library",
            "highlight_detail": ""
        },
    ],
    "queue_id": "",
    "user_id": ""
}
```