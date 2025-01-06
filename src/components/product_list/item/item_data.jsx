import IMAGES from "../../../images/Images";

const products = [
    {
        "product_id": 1,
        "product_name": "БП",
        "product_price": 599,
        "product_description": "крутотень",
        "product_type": "manual",
        "product_image_url": "линка3",
        "product_subcategory_id": 1,
        "product_sub_category": {
            "subcategory_id": 1,
            "subcategory_name": "Бравлик",
            "subcategory_description": "гемыыы",
            "subcategory_image_url": "линка2",
            "product_category_id": 1,
            "product_category": {
                "category_id": 1,
                "category_name": "Игры",
                "category_description": "анаконда",
                "category_image_url": "линка"
            }
        }
    },
    {
        "product_id": 2,
        "product_name": "Test 1",
        "product_price": 499,
        "product_description": "Test Description 1",
        "product_type": "auto",
        "product_image_url": "http://some_url.png",
        "product_subcategory_id": 1,
        "product_sub_category": {
            "subcategory_id": 1,
            "subcategory_name": "Бравлик",
            "subcategory_description": "гемыыы",
            "subcategory_image_url": "линка2",
            "product_category_id": 1,
            "product_category": {
                "category_id": 1,
                "category_name": "Игры",
                "category_description": "анаконда",
                "category_image_url": "линка"
            }
        }
    }
];


export default products;
