import "./ProductList.scss";
import ProductItem from "./item/ProductItem";
import ProductSort from "../products_sort/ProductSort";
import products from "./item/item_data";

function ProductList() {
    const productsCount = products ? products.length : 0; 
    return (
        <div className="product-list">
        <ProductSort number_of_products={productsCount}/>
        <div className="product-list-wrapper">
            {products.map((product) => (
            <ProductItem key={product.product_id} product={product}  />
            ))}
        </div>
        </div>
  );
}

export default ProductList;

