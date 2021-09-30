# Inception Quiz 
## Project Structure 
  - data
    
    Store database file 

  - docker
    
    Stores docker files 

  - src

    * app
     
        Stores main file of micro services

    * asset 

    * bin

       Stores binary files

    * common

       Stores functions that multiple services call

    * configuration 

        store parameter configuration 

    * services

        store service files


## How to compile
cd in project 

```sh 
./build.sh
```

## How to run 
cd in project

```sh
./gorun.sh
```
## How to installation 
cd in project 

```sh
docker build -t inception-quiz -f ./docker/Dockerfile .
docker run -p 25671:25671 inception-quiz
```


## API

|  uri                          |  method       |
| ----------------------------- |:-------------:|
| /v1/transaction/get           | POST          |
| /v1/transaction/create        | POST          |

#### get transaction api 
  - request 
     ```
     {
         transaction_id:4
     }
     
     ```

  - response 
    ```
    {
        "result": 1,
        "message": "get transaction",
        "value": [
            {
                "id": 4,
                "token": "src_test_5pc9bto06bh84zo8nuh",
                "amount": 102.3,
                "type": "internet_banking_scb",
                "currentcy": "THB",
                "timestamp": "2021/09/29 22:31:31"
           }
        ]
    }
    ```



#### create transaction api 

  - request 
    ```
     {
         amount:100.00
     }
     
     ```

  - response
    ```
    {
       "result": 1,
       "message": "successfully",
       "value": {
           "transaction": 6,
           "amount": 102.3
       }
    }
    ```
