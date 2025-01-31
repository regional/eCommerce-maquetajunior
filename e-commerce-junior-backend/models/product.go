package models

import (
	"fmt"
	"gorm/db"

	"gorm.io/gorm"
)

type Product struct {
	Id          int64   `json:"id"`
	Title       string  `json:"title"`
	Price       float64 `json:"price"`
	Description string  `json:"description"`
	Image       string  `json:"image"`
	CategoryId  int64   `json:"categoryId"`
	Taxes       float64 `json:"taxes"`
	Disccount   float64 `json:"disccount"`
	Inventory   int64   `json:"inventory"`

	Category Category `json:"category" gorm:"foreignKey:CategoryId;references:Id"`
}

type Products []Product

type ShopingCar struct {
	UserId   int64 `json:"user_id"`
	Quantity int64 `json:"quantity"`
	// ProductId int64 `json:"product_id"`

	Product Product `json:"product" gorm:"foreignKey:ProductId;references:Id"`
	// User User `json:"user" gorm:"foreignKey:UserId;references:Id"`
}

type ShopingCars []ShopingCar

func MigrateProduct() {
	products := []Product{
		{Title: "Shoes", Price: 50.0, Description: "Shoes", Image: "https://encrypted-tbn0.gstatic.com/images?q=tbn:ANd9GcQuIGmNvQ-a055sivvGsNg8xy_FB2l0i5Ws2g&s", CategoryId: 1, Taxes: 0.16, Disccount: 0.0, Inventory: 100},
		{Title: "Pants", Price: 20.0, Description: "Pants", Image: "https://w7.pngwing.com/pngs/63/280/png-transparent-jeans-denim-slim-fit-pants-bell-bottoms-jeans-blue-fashion-boy-thumbnail.png", CategoryId: 2, Taxes: 0.16, Disccount: 0.0, Inventory: 100},
		{Title: "T-Shirts", Price: 10.0, Description: "T-Shirts", Image: "https://encrypted-tbn0.gstatic.com/images?q=tbn:ANd9GcQAl8KTbZXd-W2ptU6bvqKxgG67GJHP98emPw&s", CategoryId: 3, Taxes: 0.16, Disccount: 0.0, Inventory: 100},
		{Title: "Sneakers", Price: 60.0, Description: "Sneakers", Image: "https://www.blackbison.co/cdn/shop/products/SNEAKER_CLASSIC_95_AZUL_OSCURO-BLANCO_1_2dd07756-1d74-4634-b522-ad4ab564d2f2.jpg?v=1618348593&width=1087", CategoryId: 1, Taxes: 0.16, Disccount: 0.0, Inventory: 100},
		{Title: "Jeans", Price: 25.0, Description: "Jeans", Image: "https://cdn.koaj.co/131959-big_default/jean-skinny.jpg", CategoryId: 2, Taxes: 0.16, Disccount: 0.0, Inventory: 100},
		{Title: "Hoodies", Price: 30.0, Description: "Hoodies", Image: "https://golty.com.co/wp-content/uploads/2023/07/hoodie-abierto-mujer-golty-rosa-1.png", CategoryId: 3, Taxes: 0.16, Disccount: 0.0, Inventory: 100},
		{Title: "Boots", Price: 70.0, Description: "Boots", Image: "https://thursdayboots.com/cdn/shop/files/1024x1024-Men-Explorer-BlackMatte-101923-2.jpg?v=1698356723&width=1024", CategoryId: 1, Taxes: 0.16, Disccount: 0.0, Inventory: 100},
		{Title: "Shorts", Price: 15.0, Description: "Shorts", Image: "https://xlivejeans.com/wp-content/uploads/2023/09/Sin-titulo-1_0000s_0002_Xlive-10-Agosto-1006-1.jpg", CategoryId: 2, Taxes: 0.16, Disccount: 0.0, Inventory: 100},
		{Title: "Sweaters", Price: 35.0, Description: "Sweaters", Image: "https://cdn11.bigcommerce.com/s-scgdirr/products/17595/images/92077/C1347_-_Moss_Green__69889.1676391063.560.850.jpg?c=2", CategoryId: 3, Taxes: 0.16, Disccount: 0.0, Inventory: 100},
		{Title: "Sandals", Price: 20.0, Description: "Sandals", Image: "https://cdn.media.amplience.net/i/clarks/ss20-journal__img-summer-sandals-guide-1-wk25?w=874&fmt=auto", CategoryId: 1, Taxes: 0.16, Disccount: 0.0, Inventory: 100},
		{Title: "Cargo Pants", Price: 22.0, Description: "Cargo Pants", Image: "https://m.media-amazon.com/images/I/71-z3NvIs0L._AC_UY1000_.jpg", CategoryId: 2, Taxes: 0.16, Disccount: 0.0, Inventory: 100},
		{Title: "Tank Tops", Price: 12.0, Description: "Tank Tops", Image: "https://assets.ajio.com/medias/sys_master/root/20240728/7C46/66a5af8c6f60443f31ced11e/-473Wx593H-465507028-navy-MODEL.jpg", CategoryId: 3, Taxes: 0.16, Disccount: 0.0, Inventory: 100},
		{Title: "Loafers", Price: 55.0, Description: "Loafers", Image: "https://m.media-amazon.com/images/I/71B-Iw66eQL._AC_UY900_.jpg", CategoryId: 1, Taxes: 0.16, Disccount: 0.0, Inventory: 100},
		{Title: "Chinos", Price: 24.0, Description: "Chinos", Image: "https://images.hawesandcurtis.com/tr:q-80/WB/WBPZY004-H20-169590-800px-1040px.jpg", CategoryId: 2, Taxes: 0.16, Disccount: 0.0, Inventory: 100},
		{Title: "Polos", Price: 18.0, Description: "Polos", Image: "https://tennis.vtexassets.com/arquivos/ids/2303011/polos-para-hombre-tennis-azul.jpg?v=638385562327530000", CategoryId: 3, Taxes: 0.16, Disccount: 0.0, Inventory: 100},
		{Title: "Flip Flops", Price: 10.0, Description: "Flip Flops", Image: "https://www.okabashi.com/cdn/shop/products/baha-womens-flip-flops-black-656287.jpg?v=1708448042", CategoryId: 1, Taxes: 0.16, Disccount: 0.0, Inventory: 100},
		{Title: "Joggers", Price: 28.0, Description: "Joggers", Image: "https://m.media-amazon.com/images/I/615DlaG0yfL._AC_SL1500_.jpg", CategoryId: 2, Taxes: 0.16, Disccount: 0.0, Inventory: 100},
		{Title: "Vests", Price: 20.0, Description: "Vests", Image: "https://i0.wp.com/www.theshepherdsknot.com/wp-content/uploads/2022/11/Vest-Trinity-92-scaled.jpg?fit=2560%2C2560&ssl=1", CategoryId: 3, Taxes: 0.16, Disccount: 0.0, Inventory: 100},
		{Title: "Dress Shoes", Price: 80.0, Description: "Dress Shoes", Image: "https://i5.walmartimages.com/asr/a36b3e18-d661-4b01-b811-978da0045adc.c83092048c45887977e63836f5a5340b.jpeg?odnHeight=768&odnWidth=768&odnBg=FFFFFF", CategoryId: 1, Taxes: 0.16, Disccount: 0.0, Inventory: 100},
		{Title: "Sweatpants", Price: 26.0, Description: "Sweatpants", Image: "https://cdni.llbean.net/is/image/wim/512293_32573_44?hei=1095&wid=950&resMode=sharp2&defaultImage=llbprod/512293_0_44", CategoryId: 2, Taxes: 0.16, Disccount: 0.0, Inventory: 100},
		{Title: "Dress", Price: 26.0, Description: "Dress", Image: "https://i.etsystatic.com/8088492/r/il/b2443d/3868645104/il_570xN.3868645104_58y5.jpg", CategoryId: 2, Taxes: 0.16, Disccount: 0.0, Inventory: 100},
		{Title: "Skirt", Price: 26.0, Description: "Skirt", Image: "https://m.media-amazon.com/images/I/71rrRNqvPkL._AC_UY1000_.jpg", CategoryId: 2, Taxes: 0.16, Disccount: 0.0, Inventory: 100},
		{Title: "Blouse", Price: 26.0, Description: "Blouse", Image: "https://gaala.com/cdn/shop/files/stretch-silk-blouse.jpg?v=1704344743&width=2667", CategoryId: 2, Taxes: 0.16, Disccount: 0.0, Inventory: 100},
		{Title: "Heels", Price: 26.0, Description: "Heels", Image: "https://i.ebayimg.com/images/g/LXgAAOSwwchhBoH6/s-l1200.jpg", CategoryId: 2, Taxes: 0.16, Disccount: 0.0, Inventory: 100},
		{Title: "Pantyhose", Price: 26.0, Description: "Pantyhose", Image: "https://m.media-amazon.com/images/I/51rDye26cFS._AC_SL1001_.jpg", CategoryId: 2, Taxes: 0.16, Disccount: 0.0, Inventory: 100},
	}

	err := db.WithDatabaseConnection(func(database *gorm.DB) error {
		database.AutoMigrate(Product{})
		for _, product := range products {
			database.FirstOrCreate(&product, product)
		}

		return nil
	})

	if err != nil {
		fmt.Printf("Error en la migración de productos: %v\n", err)
	}
}
