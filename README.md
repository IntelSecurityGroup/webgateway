# Intel/McAfee Web Gateway REST API

## Configuring the McAfee Web Gateway "MWG" REST API access
1. Login to the MWG
2. Go to Configuration --> Appliances --> User Interface --> Click Enable REST-Interface over HTTPS
3. Go to Accounts --> Create Role "REST API" --> Select REST-Interface accessible and any other rights you need
4. Under Internal Adminitrator Accounts create a new user --> Change role to "REST API"

## Running Executeable
mwg -ignoressl=true -user=mwguser -pass=mwgpass -host=127.0.0.1

### Command Options
__ignoressl:__ This will turn off SSL verification.  Works well for self signed certificates.
__user:__ The User configured on the McAfee Web Gateway
__pass:__ The Password configured on the McAfee Web Gateway
__host:__ The IP Address of the McAfee Web Gateway