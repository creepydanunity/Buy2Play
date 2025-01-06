import "./ProductItem.scss";
import { useState } from "react";

function ProductItem() {
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
      <h4>1200 гемов - Clash Of Clans</h4>
      <p>Доставка за 15 минут</p>
      <p>₽959,00</p>
      <div className="item-controls">
        <button onClick={handleDecrement}> -</button>
        <div>{numberOfItems}</div>
        <button onClick={handleIncrement}>+</button>
      </div>
    </div>
  );
}

export default ProductItem;

