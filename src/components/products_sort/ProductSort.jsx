import "./ProductSort.scss"

function ProductSort(){

    return(
        <div>
            <span>Сортировка</span>
            <form>
                <label for="fruits">Выберите фрукт:</label> 
                <select id="fruits" name="fruits">
                    <option value="apple">Яблоко</option>
                    <option value="banana">Банан</option>
                    <option value="orange">Апельсин</option>
                </select>
            </form>
            <span>Продуктов</span>
        </div>
    )
}

export default ProductSort;