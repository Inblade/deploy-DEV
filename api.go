
/**
*@api {get} /products/:id  Get list of all Products
*@apiName GetListProducts
*@apiGroup Products
*
*@apiParam {Number} id Products_ID
*@apiParam {String} name Products_Name
*
*@apiSuccess {Object[]} List of all Products
*@apiSuccessExample Success-Response:
      HTTP/1.1 200 OK
      {
      [{
      "id": "123456789",
      "name": "NAME",
      "category": "Category"
      "description": "Products_Description"
      "stock": "stock of products"
    }]
      }

*@apiError ProductNOTExists Product NOT Exists
*@apiErrorExample Error-Response:
      HTTP/1.1 400 NOT Exists
      {
      "error": "Product NOT Exists"
      }
*/
/**
*@api {get} /products/description/:description  Get description about Product
*@apiName GetDescriptionsProducts
*@apiGroup Products
*
*@apiParam {Number} id Products_ID
*@apiParam {String} name Products_Name
*@apiParam {String} category Products_Category
*@apiParam {String} description Product_Description
*@apiParam {Number} stock stock of products
*
*@apiSuccess {Object[]} Desription of Product
*@apiSuccessExample Success-Response:
      HTTP/1.1 200 OK
      {
      [{
      "id": "123456789",
      "name": "NAME",
      "category": "Category"
      "description": "Products_Description"
      "stock": "stock of products"
    }]
      }

*@apiError ProductNOTExists Product NOT Exists
*@apiErrorExample Error-Response:
      HTTP/1.1 403 Forbiden
      {
      "error": "Product NOT Exists"
      }
*/
/**
*@api {post} /products/:id  Update Products info
*@apiPermission admin
*@apiName UpdateProduct
*@apiGroup Products
*
*@apiParam {Number} id Products_ID
*@apiParam {String} name Products_Name
*@apiParam {String} category Products_Category
*@apiParam {String} description Product_Description
*@apiParam {Number} stock stock of products
*
*@apiSuccess {Number} id Product id
*@apiSuccessExample Success-Response:
      HTTP/1.1 200 OK
      {
      "id": "123456789"
      }

*@apiError ProductNOTExists Product NOT Exists
*@apiErrorExample Error-Response:
      HTTP/1.1 403 Forbiden
      {
      "error": "Product NOT Exists"
      }
*/
/**
*@api {post} /orders/shopping_cart/:id  Update Shopping cart
*@apiName UpdateOrders
*@apiGroup Orders
*
*@apiParam {Number} id Products_ID
*@apiParam {String} name Products_Name
*@apiParam {String} category Products_Category
*@apiParam {Number} quantity Product_quantity
*@apiParam {Number} stock stock of products
*
*@apiSuccess {Number} id Order id
*@apiSuccessExample Success-Response:
      HTTP/1.1 200 OK
      {
      "id": "123456789"
      }

*@apiError OrderNOTFound Order NOT Exists
*@apiErrorExample Error-Response:
      HTTP/1.1 403 Forbiden
      {
      "error": "Product NOT Exists"
      }
*/
