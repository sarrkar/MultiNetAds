endpoints:

+ `GET` `eventserver.local/impression/:id/:OTLkey`
    * `res`: `200 ok`
+ `GET` `eventserver.local/click/:id/:OTLkey`
    * `res`: `301 redirect`

expected endpoints:

+ `GET`: `panel.local/api/inc_impression/:id`
+ `GET`: `panel.local/api/inc_click/:id`
