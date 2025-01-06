import "./ProductItem.scss";
import { useState } from "react";

function ProductItem({ product }) {
  const [numberOfItems, setNumber] = useState(0);

  const handleDecrement = () => {
    if (numberOfItems > 0) {
      setNumber(numberOfItems - 1);
    }
  };

  const handleIncrement = () => {
    setNumber(numberOfItems + 1);
  };

  return (
    <div className="product-item">
      <img src="clashofclans1200.png" alt="clash" />
      <h4>{product.product_name}</h4>
      <p>{product.product_description}</p>
      <p>{product.product_price}</p>
      <div className="item-controls">
        <button onClick={handleDecrement}> - </button>
        <div>{numberOfItems}</div>
        <button onClick={handleIncrement}>+</button>
      </div>
    </div>
  );
}

export default ProductItem;

