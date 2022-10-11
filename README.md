# Kitchen simulation server

## About
This project simulates a Kitchen of a restaurant. The Kitchen has a finite order list . This order list is shared across all kitchen instances. All orders which kitchen receives have to be added to a single instance of order-list. Main work unit of the Kitchen are `cooks` . Their job is to take the order and "prepare" the menu item(s) from it, and return the orders as soon and with as little idle time as possible. Kitchen can prepare foods from different orders and it is not mandatory that one cook have to prepare entire order. Order is considered to be prepared when all foods from order list are
prepared. Each cook has the following characteristics:
* rank , which defines the complexity of the food that they can prepare (a cook can only take orders which his current rank or one rank lower that his current one):
* proficiency , it indicates on how may dishes he can work at once. It varies between 1 and 4 (and to follow a bit of logic, the
higher the rank of a cook the higher is the probability that he can work on more dishes at the same time).
* name
* catch phrase

Example of cook object:

```golang
{
"rank": 3,
"proficiency": 3,"name": "Gordon Ramsay",
"catch-phrase": "Hey, panini head, are you listening to me?"
}
```
Another requirement is to implement the cooking apparatus rule. It comprises of the fact that a kitchen has limited space, thus a finite number of ovens, stoves and the likes. The kitchen configuration have to include a limited number of cooking apparatus . For example at kitchen with 3-4
cooks, can have no more than 2 stoves and only one oven.
The Kitchen should handle HTTP (POST) requests of receiving orders from the Dinning Hall and add received order to
order list . For all received orders kitchen have to register time it was received and time is was totally prepared. Cooking
time should be added to order before sending it back to Dinning Hall .
Cooks should be an object instances which run their logic of preparing foods on separate threads , one thread per cook.
The main goal is to reduce preparation time of each order.
Cooking apparatus are sharable resources across all cooks.
Number and types of cooks and cooking apparatus should be configurable.
When order is prepared, meaning that all foods from order are prepared, Kitchen should perform HTTP (POST) request with
prepared order details to Dinning Hall in that way returning prepared order should be served to the table .

## Running the App
To run the App, run in terminal the following command:<br />


 `go run .`


## Running in Docker container
1. To run the app in a docker container, first build the image:<br />

`docker build -t kitchen-go-server .`

2. Then run the container using the created image:<br />

`docker run --name kitchen --network restaurant -it --rm  -p 8080:8080 kitchen-go-server`
To run different kitchens with different configurations, use volume bind. Example:
`docker run --name kitchen1 -v /home/irina/UTM/SEM5/PR/LAB2/First_Checkpoint/kitchen/jsonConfig:/app/jsonConfig --network restaurant -it --rm  -p 8080:8080 kitchen-go-server`
`docker run --name kitchen2 -v /home/irina/UTM/SEM5/PR/LAB2/First_Checkpoint/kitchen/jsonConfig2:/app/jsonConfig --network restaurant -it --rm  -p 8086:8086 kitchen-go-server`

For this you firstly need a created docker network. To create a docker network run:

`docker network create restaurant`

3. To stop the running container:

`docker stop {docker's id}`

## Combining with dining-hall server

The kitchen server listens first for Post request coming from the dining-hall. To run the system correctly, the kitchen server must run first, and after it the dinning-hall 
server should start running. These servers use HTTP Post request for communication.