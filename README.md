# Banking

Procedure to run Application.

To add-transactions – URL to post the details of transaction in the request body : [POST] -- http://localhost:8000/transactions

i)
{
"amount":"50.25",
"timestamp": "2023-04-02T18:50:50.834Z"
}

Note : transaction within last 60 seconds will be accepted.
 
EXPECTED RESPONSE: SUCCESS

To check the statistics of transaction, please redirect to URL [GET] –- http://localhost:8000/statistics

EXPECTED RESPONSE: sum, average, max, min, and number of transaction happened in last 60 seconds will be retrieved as response

To add location details : [POST] -- http://localhost:8000/location

Enter the following details in the request body

{
 "city":"bangalore"
}

EXPECTED RESPONSE: Location added Successfully...

To Reset Location details [PUT] -- http://localhost:8000/resetLocation

EXPECTED RESPONSE: location reset Successfull...

To delete all transactions : [DELETE] -- http://localhost:8000/transaction

EXPECTED RESPONSE: 204
