
# Goods and discount storage REST API

> Simple RESTful API for 3rd party projects, which operate with markets and parsing goods.


## Endpoints

### Get Catalogs
``` bash
GET /api/v1/catalogs?shopId={shopID}

# Response

{
    "shop": {
        "ID": 1
    },
    "catalogs": [
        {
            "categoryID": 741,
            "url": "https://www.service.ru/computers-1/pc?page=",
            "count": 10
        },
         {
            "categoryID": 1,
            "url": "https://www.service.ru/televizory-1/tv?page=",
            "count": 12
        }
    ]
}
```

### Save goods and defining a discount

``` bash
POST /api/v1/goods

# Request sample
{
  "shop": {
    "ID": 1
  },
  "goods": [
    {
      "categoryID": 47,
      "article": "Article_1",
      "name": "Name_1",
      "url": "http://locl.su/id/1231",
      "imageUrl": "https://locl.su/id/1231.jpg",
      "price": 5000
    },
    {
      "categoryID": 47,
      "article": "Article_2",
      "name": "Name_2",
      "url": "http://locl.su/id/1232",
      "imageUrl": "http://locl.su/id/1232.png",
      "price": 12500
    }
  ]
}
```

### Send warning message to administrator
``` bash
POST /api/v1/logs

# Request sample
{
    "text": "ADD_TEXT_HERE"
}

```
