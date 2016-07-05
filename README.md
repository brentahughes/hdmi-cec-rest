HDMI CEC Rest API
======
Rest api to interact with the consumer electronics control of a hdmi device. Great for turning a raspberry pi into a tv remote.


API Reference
------

### /device
Get list of all active devices


#### GET
* Request
```
curl http://localhost:5000/device
```

* Response
```json
{
  "TV": {
    "PowerStatus": "standby",
    "PhysicalAddress": "0.0.0.0",
    "ActiveSource": false,
    "OSDName": "TV",
    "Vendor": "Samsung",
    "LogicalAddress": 0
  },
  "Playback": {
    "PowerStatus": "on",
    "PhysicalAddress": "1.0.0.0",
    "ActiveSource": false,
    "OSDName": "Roku 3",
    "Vendor": "",
    "LogicalAddress": 4
  },
  "Recording": {
    "PowerStatus": "on",
    "PhysicalAddress": "2.0.0.0",
    "ActiveSource": false,
    "OSDName": "cec.go",
    "Vendor": "Pulse Eight",
    "LogicalAddress": 1
  }
}
```



### /device/[logical_address]
Get details for single device by its physical address

#### GET
* Request
```
curl http://localhost:5000/device/0
```

* Response
```json
{
  "PowerStatus": "standby",
  "PhysicalAddress": "0.0.0.0",
  "ActiveSource": false,
  "OSDName": "TV",
  "Vendor": "Samsung",
  "LogicalAddress": 0
}
```


### /device/[logical_address]/power
Get or set the power status for a device

#### GET
* Request
```
curl http://localhost:5000/device/0/power
```

* Response
```json
{
    "status": "success",
    "message": "standby"
}
```

#### POST
* Request
```
# state can be "on" or "off"
curl -X POST -d '{"state": "on"}' http://localhost:5000/device/0/power
```



### /device/[logical_address]/volume
Set the volume for a amplifier device

This will only work if the device is type Amplifier

#### POST
* Request
```
# state can be "up", "down", or "mute"
curl -X POST -d '{"state": "up"}' http://localhost:5000/device/0/volume
```


Installation
------

This assumes golang has been installed and $GOPATH set

* Install Dependencies
  * Ubuntu/Debian
```bash
apt-get install libcec-dev cec-utils make
```


* Compile Binary
```bash
make
```


* Install Service
```bash
make install
```
