import "./ProductList.scss";
import ProductItem from "./item/ProductItem";


function ProductList() {
  return (
    <div className="product-list">
      <ProductItem/>
      <ProductItem/>
      <ProductItem/>
      <ProductItem/>
      <ProductItem/>
      <ProductItem/>
    </div>
  );
}

export default ProductList;
