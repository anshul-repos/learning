This repository will be used for learning content.



# Directory Structure:

    - learnig:
            - int-prep/coding:
                    - general-topics:
                            - concurrency:
                            - errorhandling:
                            - go-loadbalancer:
                            - oops:
                            - restapi:
                            - workerpool:
                    - go-algos:
                    - go-data-structures:
                            - array:
                            - string:

# 1 REST API CODE:
    - How to run: go run api.go (server will start running on port 8181)
    - How to send post curl request: 
            ansh@Mac golang % curl -X POST http://localhost:8181/employee -H "Content-Type: application/json" -d '{"id":1,"name":"Alice"}'{"id":1,"name":"Alice"}

            ansh@Mac golang % curl -X POST http://localhost:8181/employee -H "Content-Type: application/json" -d '{"id":2,"name":"BOB"}' 
            {"id":2,"name":"BOB"}

    - How to send GET request: http://localhost:8181/employee from browser.

