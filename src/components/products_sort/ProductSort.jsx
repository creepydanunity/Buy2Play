import "./ProductSort.scss"

function ProductSort({number_of_products=8}){

    return(
        <div className="product-sort">
            <div className="product-sort-wrapper">
                <form>
                    <label>Сортировка:</label> 
                    <select >
                        <option value="Рекомендуется">Рекомендуется</option>
                        <option value="Лидер продаж">Лидер продаж</option>
                        <option value="От А до Я">От А до Я</option>
                        <option value="От Я до А">От Я до А</option>
                        <option value="По возрастанию цены">По возрастанию цены</option>
                        <option value="По убыванию цены">По убыванию цены</option>
                        <option value="Сначала старые">Сначала старые</option>
                        <option value="Сначала новые">Сначала новые</option>
                    </select>
                </form>
                <span>Продуктов:{number_of_products}</span>
            </div>
        </div>
    )
}

export default ProductSort;